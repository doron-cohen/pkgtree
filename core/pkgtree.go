package core

import (
	"context"

	"github.com/doron-cohen/pkgtree/packages"
	"golang.org/x/exp/slices"
)

func GetChangedPackages(ctx context.Context, ref string, includeDirty bool, repoRoot string) ([]string, error) {
	files, err := GetChangedFiles(ctx, ref, includeDirty, repoRoot)
	if err != nil {
		return nil, err
	}

	pkgs := make([]string, 0, len(files))
	for _, file := range files {
		if file == "" {
			continue
		}

		pkg, err := packages.GetFilePackageName(ctx, file)
		if err != nil {
			return nil, err
		}

		pkgs = append(pkgs, pkg)
	}

	return uniqueAndSort(pkgs), nil
}

func uniqueAndSort(pkgs []string) []string {
	seen := make(map[string]struct{})
	uniquePkgs := make([]string, 0, len(pkgs))
	for _, pkg := range pkgs {
		if _, ok := seen[pkg]; ok {
			continue
		}

		seen[pkg] = struct{}{}
		uniquePkgs = append(uniquePkgs, pkg)
	}

	slices.Sort(uniquePkgs)
	return uniquePkgs
}
