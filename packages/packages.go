package pkgs

import (
	"context"
	"fmt"
	"path/filepath"

	"golang.org/x/tools/go/packages"
)

func GetFilePackagePath(ctx context.Context, filePath string) (string, error) {
	cfg := &packages.Config{
		Context: ctx,
		Mode:    packages.NeedName | packages.NeedFiles,
		Dir:     filepath.Dir(filePath),
	}

	pkgs, err := packages.Load(cfg, "file="+filePath)
	if err != nil {
		return "", fmt.Errorf("failed to load package: %w", err)
	}

	if len(pkgs) == 0 {
		// TODO: debug log about skipping non go file
		return "", nil
	}

	for _, pkg := range pkgs {
		for _, file := range pkg.GoFiles {
			if file == filePath {
				return pkg.PkgPath, nil
			}
		}
	}

	return "", nil
}
