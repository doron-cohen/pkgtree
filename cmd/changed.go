package cmd

import (
	"context"

	"github.com/doron-cohen/pkgtree/core"
)

type ChangedCmd struct{}

func (c *ChangedCmd) Run() error {
	ctx := context.Background()
	files, err := core.GetChangedPackages(ctx, core.DifferTypeGit)
	if err != nil {
		return err
	}

	for _, file := range files {
		println(file)
	}

	return nil
}
