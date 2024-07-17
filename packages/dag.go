package pkgs

import (
	"context"
	"fmt"
	"strings"

	"github.com/heimdalr/dag"
	"golang.org/x/tools/go/packages"
)

type DependencyGraph struct {
	d *dag.DAG
}

func BuildDependencyGraph(ctx context.Context, moduleDir string) (*DependencyGraph, error) {
	cfg := &packages.Config{
		Context: ctx,
		Mode:    packages.NeedName | packages.NeedImports | packages.NeedDeps | packages.NeedModule,
		Dir:     moduleDir,
	}

	pkgs, err := packages.Load(cfg, "./...")
	if err != nil {
		return nil, fmt.Errorf("failed to load packages: %w", err)
	}

	if len(pkgs) == 0 {
		return nil, fmt.Errorf("no packages found")
	}

	modulePath := pkgs[0].Module.Path

	d := dag.NewDAG()
	for _, pkg := range pkgs {
		err = d.AddVertexByID(pkg.PkgPath, pkg)
		if err != nil {
			return nil, fmt.Errorf("failed to add vertex: %w", err)
		}
	}

	for _, pkg := range pkgs {
		for _, dep := range pkg.Imports {
			if strings.HasPrefix(dep.PkgPath, modulePath) {
				// We are adding an edge from the dependant to the source to quickly
				// identify all packages which are dependent on certain packages.
				err = d.AddEdge(dep.PkgPath, pkg.PkgPath)
				if err != nil {
					return nil, fmt.Errorf("failed to add edge: %w", err)
				}
			}
		}
	}

	return &DependencyGraph{d}, nil
}

func (d *DependencyGraph) GetImporters(pkgName string) ([]string, error) {
	descendants, err := d.d.GetDescendants(pkgName)
	if err != nil {
		return nil, fmt.Errorf("failed to get package descendants: %w", err)
	}

	paths := make([]string, 0, len(descendants))
	for id, desc := range descendants {
		descPkg, ok := desc.(*packages.Package)
		if !ok {
			return nil, fmt.Errorf("package %s is not a package", id)
		}

		paths = append(paths, descPkg.PkgPath)
	}

	return paths, nil
}