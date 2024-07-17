# pkgtree

Go packages dependency analysis tool

## Why?

When working with large Go repositories running all tests or other tasks for every change in CI is costly.
With `pkgtree` it is possible to identify which packages changed or affected according to some commit.

## Examples

Get changed packages from last commit to current commit (HEAD^ to HEAD):
```
❯ pkgtree changed
```

Get changed packages from a certain commit to current commit including uncommited changes:
```
❯ pkgtree changed --since-ref=75b634b30f4d6705433d916b2f55aa233dc9be55 --include-dirty
```

Get affected (including changed) packages from last commit to current commit (HEAD^ to HEAD):
```
❯ pkgtree affected
```

Get affected packages from certain reference excluding the changed packages themselves:
```
❯ pkgtree affected --since-ref=origin/main~4 --include-changed=false
```

Pass a different git dir:
```
❯ pkgtree affected --git-dir=~/dev/pkgtree
```
