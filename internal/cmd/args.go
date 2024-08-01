package cmd

// CommonArgs are args common to all commands
type CommonArgs struct {
	GitDir string `default:"." type:"existingdir" help:"The git repository to use."`
}

// ChangeArgs are args common to git change related commands
type ChangeArgs struct {
	SinceRef     string `default:"HEAD^" help:"The ref to compare against."`
	IncludeDirty bool   `default:"false" help:"Include uncommitted changes."`
}
