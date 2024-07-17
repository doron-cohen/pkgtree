# pkgtree

Go packages dependency analysis tool

## Why?

When working with large Go repositories running all tests or other tasks for every change in CI is costly.
With `pkgtree` it is possible to identify which packages changed or affected according to some commit.

## Examples

Get changed packages from last commit to current commit (HEAD^ to HEAD):
```
❯ go run main.go changed
github.com/doron-cohen/pkgtree/core
github.com/doron-cohen/pkgtree/packages
```

Get changed packages from a certain commit to current commit including uncommited changes:
```
❯ go run main.go changed --since-ref=75b634b30f4d6705433d916b2f55aa233dc9be55 --include-dirty
github.com/doron-cohen/pkgtree/cmd
github.com/doron-cohen/pkgtree/core
github.com/doron-cohen/pkgtree/logger
github.com/doron-cohen/pkgtree/packages
```

Get affected (including changed) packages from last commit to current commit (HEAD^ to HEAD):
```
❯ go run main.go affected
github.com/doron-cohen/pkgtree
github.com/doron-cohen/pkgtree/cmd
github.com/doron-cohen/pkgtree/core
github.com/doron-cohen/pkgtree/packages
```

Get affected packages from certain reference excluding the changed packages themselves:
```
❯ go run main.go affected --since-ref=origin/main~4 --include-changed=false
github.com/doron-cohen/pkgtree
github.com/doron-cohen/pkgtree/cmd
github.com/doron-cohen/pkgtree/core
github.com/doron-cohen/pkgtree/packages
```

Pass a different git dir:
```
❯ go run main.go affected --git-dir=~/dev/pkgtree
github.com/doron-cohen/pkgtree
github.com/doron-cohen/pkgtree/cmd
github.com/doron-cohen/pkgtree/core
github.com/doron-cohen/pkgtree/packages
```
