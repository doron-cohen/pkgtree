package cmd

import (
	"github.com/alecthomas/kong"
)

var cli struct {
	Changed  ChangedCmd  `cmd:"" help:"List packages that have changed in the working tree."`
	Affected AffectedCmd `cmd:"" help:"List affected packages."`
	Tree     TreeCmd     `cmd:"" help:"Print packages tree."`
}

func Run() {
	ctx := kong.Parse(&cli)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
