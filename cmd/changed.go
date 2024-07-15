package cmd

import (
	"context"
	"path/filepath"

	"github.com/doron-cohen/pkgtree/core"
)

type ChangedCmd struct {
	SinceRef     string `default:"HEAD^" help:"The ref to compare against."`
	IncludeDirty bool   `default:"false" help:"Include dirty files in the output."`
	GitDir       string `default:"." type:"existingdir" help:"The git repository to use."`
}

func (c *ChangedCmd) Run() error {
	ctx := context.Background()

	gitDir, err := filepath.Abs(c.GitDir)
	if err != nil {
		return err
	}

	files, err := core.GetChangedPackages(ctx, c.SinceRef, c.IncludeDirty, gitDir)
	if err != nil {
		return err
	}

	for _, file := range files {
		println(file)
	}

	return nil
}
