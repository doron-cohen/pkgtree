package cmd

import (
	"context"
	"log/slog"
	"path/filepath"
	"slices"

	"github.com/doron-cohen/pkgtree/core"
	"github.com/doron-cohen/pkgtree/logger"
)

type ChangedCmd struct {
	Args
}

func (c *ChangedCmd) Run() error {
	ctx := context.Background()
	ctx = logger.NewContext(ctx, logger.NewConsoleLogger(slog.LevelInfo))

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
