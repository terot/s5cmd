package command

import (
	"context"
	"fmt"

	"github.com/urfave/cli/v2"

	errorpkg "github.com/peak/s5cmd/error"
	"github.com/peak/s5cmd/log"
	"github.com/peak/s5cmd/objurl"
	"github.com/peak/s5cmd/storage"
	"github.com/peak/s5cmd/strutil"
)

var sizeCommandFlags = []cli.Flag{
	&cli.BoolFlag{Name: "humanize", Aliases: []string{"H"}},
	&cli.BoolFlag{Name: "group", Aliases: []string{"g"}},
}

var SizeCommand = &cli.Command{
	Name:     "du",
	HelpName: "disk-usage",
	Usage:    "TODO",
	Flags:    append(sizeCommandFlags, globalFlags...),
	Before: func(c *cli.Context) error {
		validate := func() error {
			if err := validateGlobalFlags(c); err != nil {
				return err
			}
			if c.Args().Len() != 1 {
				return fmt.Errorf("expected only 1 argument")
			}
			return nil
		}
		if err := validate(); err != nil {
			printError(givenCommand(c), c.Command.Name, err)
			return err
		}

		setGlobalFlags(c)
		return nil
	},
	Action: func(c *cli.Context) error {
		groupByClass := c.Bool("group")
		humanize := c.Bool("humanize")

		err := Size(
			c.Context,
			givenCommand(c),
			c.Args().First(),
			groupByClass,
			humanize,
		)
		if err != nil {
			printError(givenCommand(c), c.Command.Name, err)
			return err
		}

		return nil
	},
}

func Size(
	ctx context.Context,
	fullCommand string,
	src string,
	groupByClass bool,
	humanize bool,
) error {
	srcurl, err := objurl.New(src)
	if err != nil {
		return err
	}

	client, err := storage.NewClient(srcurl)
	if err != nil {
		return err
	}

	storageTotal := map[string]sizeAndCount{}
	total := sizeAndCount{}

	for object := range client.List(ctx, srcurl, true, storage.ListAllItems) {
		if object.Type.IsDir() || errorpkg.IsCancelation(object.Err) {
			continue
		}

		if err := object.Err; err != nil {
			printError(fullCommand, "list", err)
			continue
		}
		storageClass := string(object.StorageClass)
		s := storageTotal[storageClass]
		s.addObject(object)
		storageTotal[storageClass] = s

		total.addObject(object)
	}

	if !groupByClass {
		msg := SizeMessage{
			Source:        srcurl.String(),
			Count:         total.count,
			Size:          total.size,
			showHumanized: humanize,
		}
		log.Info(msg)
		return nil
	}

	for k, v := range storageTotal {
		msg := SizeMessage{
			Source:        srcurl.String(),
			StorageClass:  k,
			Count:         v.count,
			Size:          v.size,
			showHumanized: humanize,
		}
		log.Info(msg)
	}

	return nil
}

type sizeAndCount struct {
	size  int64
	count int64
}

func (s *sizeAndCount) addObject(obj *storage.Object) {
	s.size += obj.Size
	s.count++
}

// SizeMessage is the structure for logging disk usage.
type SizeMessage struct {
	Source       string `json:"source"`
	StorageClass string `json:"storage_class,omitempty"`
	Count        int64  `json:"count"`
	Size         int64  `json:"size"`

	showHumanized bool
}

// humanize is a helper function to humanize bytes.
func (s SizeMessage) humanize() string {
	var size string
	if s.showHumanized {
		size = humanizeBytes(s.Size)
	} else {
		size = fmt.Sprintf("%d", s.Size)
	}
	return size
}

// String returns the string representation of SizeMessage.
func (s SizeMessage) String() string {
	storageCls := ""
	if s.StorageClass != "" {
		storageCls = fmt.Sprintf(" [%s]", s.StorageClass)
	}
	return fmt.Sprintf(
		"%s bytes in %d objects: %s%s",
		s.humanize(),
		s.Count,
		s.Source,
		storageCls,
	)
}

// JSON returns the JSON representation of SizeMessage.
func (s SizeMessage) JSON() string {
	return strutil.JSON(s)
}
