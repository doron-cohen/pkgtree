package cmd

type Args struct {
	SinceRef     string `default:"HEAD^" help:"The ref to compare against."`
	IncludeDirty bool   `default:"false" help:"Include uncommitted changes."`
	GitDir       string `default:"." type:"existingdir" help:"The git repository to use."`
}
