package command

import (
	"context"

	"github.com/urfave/cli/v2"

	"github.com/peak/s5cmd/log"
	"github.com/peak/s5cmd/parallel"
	"github.com/peak/s5cmd/storage"
)

var app = &cli.App{
	Flags: globalFlags,
	Before: func(c *cli.Context) error {
		if err := validateGlobalFlags(c); err != nil {
			return err
		}

		s3opts := storage.S3Opts{
			MaxRetries:             c.Int("retry-count"),
			EndpointURL:            c.String("endpoint-url"),
			NoVerifySSL:            c.Bool("no-verify-ssl"),
			UploadChunkSizeBytes:   c.Int64("upload-chunk-size") * megabytes,
			UploadConcurrency:      c.Int("upload-concurrency"),
			DownloadChunkSizeBytes: c.Int64("download-chunk-size") * megabytes,
			DownloadConcurrency:    c.Int("download-concurrency"),
		}

		storage.SetS3Options(s3opts)
		log.Init(c.String("log"), c.Bool("json"))
		parallel.Init(c.Int("numworkers"))

		return nil
	},
	After: func(c *cli.Context) error {
		parallel.Close()
		log.Close()
		return nil
	},
	Action: func(c *cli.Context) error {
		return cli.ShowAppHelp(c)
	},
}

func Main(ctx context.Context, args []string) error {
	app.Commands = []*cli.Command{
		ListCommand,
		SizeCommand,
		MakeBucketCommand,
		DeleteCommand,
		CopyCommand,
		MoveCommand,
		GetCommand,
		RunCommand,
		VersionCommand,
	}

	return app.RunContext(ctx, args)
}
