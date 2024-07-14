package core

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
)

type DifferType string

const (
	DifferTypeGit DifferType = "git"
)

func GetChangedFiles(
	ctx context.Context,
	differType DifferType,
	repoRoot string,
) (files []string, err error) {
	switch differType {
	case DifferTypeGit:
		files, err = GetGitDiff(ctx, repoRoot)
		if err != nil {
			return nil, fmt.Errorf("failed to get git diff: %w", err)
		}
	default:
		return nil, fmt.Errorf("unsupported diff type: %s", differType)
	}

	return files, nil
}

func GetGitDiff(ctx context.Context, repoRoot string) ([]string, error) {
	var out bytes.Buffer
	cmd := exec.CommandContext(ctx, "git", "diff", "--name-only", "HEAD^", "HEAD")
	cmd.Stdout = &out
	cmd.Dir = repoRoot

	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to run git diff: %w", err)
	}

	changedFiles := strings.Split(strings.TrimSpace(out.String()), "\n")
	return changedFiles, nil
}
