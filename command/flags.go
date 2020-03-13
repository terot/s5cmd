package command

import (
	"fmt"

	"github.com/peak/s5cmd/log"
	"github.com/peak/s5cmd/parallel"
	"github.com/peak/s5cmd/storage"
	"github.com/urfave/cli/v2"
)

const (
	defaultWorkerCount         = 256
	defaultUploadConcurrency   = 5
	defaultDownloadConcurrency = 5
	defaultChunkSize           = 50
	defaultRetryCount          = 10

	megabytes = 1024 * 1024
)

var globalFlags = []cli.Flag{
	&cli.IntFlag{Name: "download-concurrency", Aliases: []string{"dw"}, Value: defaultDownloadConcurrency},
	&cli.IntFlag{Name: "upload-concurrency", Aliases: []string{"uw"}, Value: defaultUploadConcurrency},
	&cli.IntFlag{Name: "download-chunk-size", Aliases: []string{"ds"}, Value: defaultChunkSize},
	&cli.IntFlag{Name: "upload-chunk-size", Aliases: []string{"us"}, Value: defaultChunkSize},
	&cli.IntFlag{Name: "retry-count", Aliases: []string{"r"}, Value: defaultRetryCount},
	&cli.BoolFlag{Name: "no-verify-ssl"},
	&cli.StringFlag{Name: "endpoint-url"},
	&cli.IntFlag{Name: "numworkers", Value: defaultWorkerCount},
	&cli.BoolFlag{Name: "json"},
	&cli.StringFlag{Name: "log", Value: "info"},
}

func validateGlobalFlags(c *cli.Context) error {
	downloadConcurrency := c.Int("download-concurrency")
	downloadChunkSize := c.Int64("download-chunk-size")
	uploadConcurrency := c.Int("upload-concurrency")
	uploadChunkSize := c.Int64("upload-chunk-size")
	retryCount := c.Int("retry-count")

	if uploadChunkSize < 5 {
		return fmt.Errorf("upload chunk size should be greater than 5 MB")
	}

	if downloadChunkSize < 5 {
		return fmt.Errorf("download chunk size should be greater than 5 MB")
	}

	if downloadConcurrency < 1 || uploadConcurrency < 1 {
		return fmt.Errorf("download/upload concurrency should be greater than 1")
	}

	if retryCount < 1 {
		return fmt.Errorf("retry count must be a positive value")
	}

	return nil
}

var setGlobals = true

func setGlobalFlags(c *cli.Context) {
	if !setGlobals {
		return
	}

	if isSet(c, "log", "json") {
		log.SetJSON(c.Bool("json"))
		log.SetLevel(c.String("log"))
	}

	if isSet(c, "numworkers") {
		parallel.Close()
		parallel.Init(c.Int("numworkers"))
	}

	if isSet(
		c,
		"retry-count",
		"endpoint-url",
		"no-verify-ssl",
		"upload-chunk-size",
		"upload-concurrency",
		"download-chunk-size",
		"download-concurrency",
	) {
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
	}

	setGlobals = false
}

func isSet(c *cli.Context, flags ...string) bool {
	for _, flag := range flags {
		if c.IsSet(flag) {
			return true
		}
	}
	return false
}
