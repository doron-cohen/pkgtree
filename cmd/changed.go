package cmd

import (
	"context"

	"github.com/doron-cohen/pkgtree/core"
)

type ChangedCmd struct {
	SinceRef string `default:"HEAD^" help:"The ref to compare against."`
}

func (c *ChangedCmd) Run() error {
	ctx := context.Background()
	files, err := core.GetChangedPackages(ctx, c.SinceRef)
	if err != nil {
		return err
	}

	for _, file := range files {
		println(file)
	}

	return nil
}
