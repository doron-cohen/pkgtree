package cmd

import (
	"context"
	"log/slog"
	"path/filepath"
	"slices"

	"github.com/doron-cohen/pkgtree/internal/core"
	"github.com/doron-cohen/pkgtree/internal/logger"
)

type AffectedCmd struct {
	CommonArgs
	ChangeArgs

	IncludeChanged bool `default:"true" help:"include changed packages"`
}

func (c *AffectedCmd) Run() error {
	ctx := context.Background()
	ctx = logger.NewContext(ctx, logger.NewConsoleLogger(slog.LevelInfo))

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
