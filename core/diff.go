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
	includeDirty bool,
	repoRoot string,
) (files []string, err error) {
	var out bytes.Buffer

	args := []string{"diff", "--name-only", ref}
	if !includeDirty {
		args = append(args, "HEAD")
	}

	cmd := exec.CommandContext(ctx, "git", args...)
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
