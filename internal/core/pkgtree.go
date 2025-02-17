package core

import (
	"context"
	"fmt"

	"github.com/doron-cohen/pkgtree/internal/pkgs"
)

func GetChangedPackages(ctx context.Context, ref string, includeDirty bool, repoRoot string) ([]string, error) {
	files, err := GetChangedFiles(ctx, ref, includeDirty, repoRoot)
	if err != nil {
		return nil, err
	}

	pkgNames, err := pkgs.GetFilesPackagePaths(ctx, repoRoot, files...)
	if err != nil {
		return nil, err
	}

	return unique(pkgNames), nil
}

func GetAffectedPackages(
	ctx context.Context, ref string, includeDirty bool, repoRoot string, includeChanged bool,
) ([]string, error) {
	changed, err := GetChangedPackages(ctx, ref, includeDirty, repoRoot)
	if err != nil {
		return nil, fmt.Errorf("failed to get changed packages: %w", err)
	}

	depGraph, err := pkgs.BuildDependencyGraph(ctx, repoRoot, pkgs.EdgeDirectionFromDependency)
	if err != nil {
		return nil, fmt.Errorf("failed to build dependency graph: %w", err)
	}

	affected := make([]string, 0, len(changed))
	if includeChanged {
		affected = append(affected, changed...)
	}

	for _, pkgName := range changed {
		paths, err := depGraph.GetImporters(pkgName)
		if err != nil {
			return nil, err
		}

		affected = append(affected, unique(paths)...)
	}

	return unique(affected), nil
}

func GetPackageTree(ctx context.Context, repoRoot string) (*pkgs.DependencyGraph, error) {
	return pkgs.BuildDependencyGraph(ctx, repoRoot, pkgs.EdgeDirectionToDependency)
}

func unique(pkgs []string) []string {
	seen := make(map[string]struct{})
	uniquePkgs := make([]string, 0, len(pkgs))
	for _, pkg := range pkgs {
		if _, ok := seen[pkg]; ok {
			continue
		}

		seen[pkg] = struct{}{}
		uniquePkgs = append(uniquePkgs, pkg)
	}

	return uniquePkgs
}
