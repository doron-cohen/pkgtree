package pkgs

import (
	"context"
	"fmt"
	"log/slog"
	"path/filepath"

	"github.com/doron-cohen/pkgtree/logger"
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
		logger.FromContext(ctx).Info("no packages found for file", slog.String("path", filePath))
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
