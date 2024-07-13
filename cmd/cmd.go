package cmd

import (
	"github.com/alecthomas/kong"
)

var cli struct{}

func Run() {
	ctx := kong.Parse(&cli)
	err := ctx.Run()
	ctx.FatalIfErrorf(err)
}
