package core

import (
	"bytes"
	"context"
	"fmt"
	"os/exec"
	"strings"
)

func GetChangedFiles(
	ctx context.Context,
	ref string,
	repoRoot string,
) (files []string, err error) {
	var out bytes.Buffer
	cmd := exec.CommandContext(ctx, "git", "diff", "--name-only", ref, "HEAD")
	cmd.Stdout = &out
	cmd.Dir = repoRoot

	err = cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to run git diff: %w", err)
	}

	fileList := strings.TrimSpace(out.String())
	if fileList == "" {
		return nil, nil
	}

	files = strings.Split(fileList, "\n")
	return files, nil
}
