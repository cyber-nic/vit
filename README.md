vit is a simple and tiny filesystem navigator

The term 'vit' is a play 'git' and 'vite', the French word for fast.

vit was initially imagined to help navigate amongst the may git repos scattered on the author's filesystem, but vite isn't git specific; it supports aliasing and navigating to any path.

The techical challenge with this idea is simply that no application may alter your shell's current path.
See 'you can do this' https://stackoverflow.com/questions/52435908/how-to-change-the-shells-current-working-directory-in-go

This tool's solution is to combine the 'vit' bin with a tiny bash function 'vd' which combines 'vit' and 'cd'.

```
function v() { if [ "$#" -eq 0 ]; then vit; elif [ "$#" -eq 1 ]; then cd `vit alias get $1`; else echo "Invalid args."; fi }

```

--help : view this message\n
vit init -> create a vit config file (typically ~/.vit/config)\n
vit list -> list current config\n

// vit alias set foo .

// vit init
// vit alias set . // set current dir and use default name
// vit alias set .. // set parent dir and use default name
// vit alias set foo .. // alias 'foo' to parent dir
// vit alias set 0 . // alias '0' to current dir

# WARNING

vit supports navigating to paths using either the alias index or name. The index takes PRECEDENSE.

// vit alias set 0 /path/to/bar

// vit alias get foo
// vit alias rm foo
// vit alias set 0 /other/path/to/bar -f

// vit cd bar
// vit cd 0

`ln -s /code/vit/vit /home/me/go/bin/vit`
`go build -o vit *.go`

add this to .bashrc

[alias]
0 = /code/4k2
foo = ~/test

go run \*.go init

# adding multiple folders
```
for d in `find .  -maxdepth 1 | sort` ; do vit alias add "$d"; done
```