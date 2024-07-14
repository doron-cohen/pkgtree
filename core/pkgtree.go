package core

import (
	"context"
)

func GetChangedPackages(ctx context.Context, differType DifferType) ([]string, error) {
	return GetChangedFiles(ctx, differType, ".")
}
