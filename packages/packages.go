package packages

import (
	"context"
	"fmt"
	"path/filepath"

	"golang.org/x/tools/go/packages"
)

func GetFilePackageName(ctx context.Context, filePath string) (string, error) {
	cfg := &packages.Config{
		Context: ctx,
		Mode:    packages.NeedName | packages.NeedFiles,
		Dir:     filepath.Dir(filePath),
	}

	pkgs, err := packages.Load(cfg, ".")
	if err != nil {
		return "", fmt.Errorf("failed to load package: %w", err)
	}

	if len(pkgs) == 0 {
		return "", fmt.Errorf("no packages found")
	}

	return pkgs[0].PkgPath, nil
}

func getFilePackageName(filePath string, pkgs []*packages.Package) (string, error) {
	for _, pkg := range pkgs {
		for _, file := range pkg.GoFiles {
			if file == filePath {
				return pkg.PkgPath, nil
			}
		}
	}

	return "", fmt.Errorf("package not found")
}
