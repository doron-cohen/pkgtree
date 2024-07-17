package cmd

import (
	"context"
	"path/filepath"
	"slices"

	"github.com/doron-cohen/pkgtree/core"
)

type ChangedCmd struct {
	Args
}

func (c *ChangedCmd) Run() error {
	ctx := context.Background()

	gitDir, err := filepath.Abs(c.GitDir)
	if err != nil {
		return err
	}

	changed, err := core.GetChangedPackages(ctx, c.SinceRef, c.IncludeDirty, gitDir)
	if err != nil {
		return err
	}

	slices.Sort(changed)
	for _, path := range changed {
		println(path)
	}

	return nil
}
