package cmd

import (
	"context"
	"fmt"

	"github.com/doron-cohen/pkgtree/internal/core"
	"github.com/doron-cohen/pkgtree/internal/pkgs"
	"github.com/pterm/pterm"
)

type TreeCmd struct {
	CommonArgs
}

func (c *TreeCmd) Run() error {
	ctx := context.Background()

	dag, err := core.GetPackageTree(ctx, c.GitDir)
	if err != nil {
		return err
	}

	fmt.Println(dag.String())
	return nil
}

func formatTree(dag *pkgs.DependencyGraph) string {
	tree := pterm.Tree{}

}
