package pkgs

import (
	"context"
	"fmt"
	"slices"

	"github.com/doron-cohen/pkgtree/logger"
	"golang.org/x/tools/go/packages"
)

func GetFilesPackagePaths(ctx context.Context, moduleDir string, filePaths ...string) ([]string, error) {
	cfg := &packages.Config{
		Context: ctx,
		Mode:    packages.NeedName | packages.NeedFiles,
		Dir:     moduleDir,
	}

	patterns := make([]string, len(filePaths))
	for i, filePath := range filePaths {
		patterns[i] = fmt.Sprintf("file=%s", filePath)
	}

	pkgs, err := packages.Load(cfg, patterns...)
	if err != nil {
		return nil, fmt.Errorf("failed to load package: %w", err)
	}

	if len(pkgs) == 0 {
		logger.FromContext(ctx).Info("no packages found for files")
		return nil, nil
	}

	pkgNames := make([]string, 0, len(pkgs))
	for _, pkg := range pkgs {
		for _, file := range pkg.GoFiles {
			if slices.Contains(filePaths, file) {
				pkgNames = append(pkgNames, pkg.PkgPath)
			}
		}
	}

	return pkgNames, nil
}
