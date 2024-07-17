package cmd

import (
	"context"
	"path/filepath"
	"slices"

	"github.com/doron-cohen/pkgtree/core"
)

type AffectedCmd struct {
	Args

	IncludeChanged bool `default:"true" help:"include changed packages"`
}

func (c *AffectedCmd) Run() error {
	ctx := context.Background()

	gitDir, err := filepath.Abs(c.GitDir)
	if err != nil {
		return err
	}

	affected, err := core.GetAffectedPackages(ctx, c.SinceRef, c.IncludeDirty, gitDir, c.IncludeChanged)
	if err != nil {
		return err
	}

	slices.Sort(affected)
	for _, path := range affected {
		println(path)
	}

	return nil
}
