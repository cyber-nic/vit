package main

import (
	"fmt"
	"log"
	"os"
)

const conf = "/.vitconfig"

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func exitGracefully(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	fmt.Fprintln(os.Stdout, ".")
	os.Exit(1)
}

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		printConfig()
		os.Exit(0)
	}

	if len(args) == 1 && args[0] == "--help" {
		fmt.Printf(`
		vit is a tiny and simple filesystem navigation helper.

		The term 'vit' is a play 'git' and 'vite', the French word for fast.

		vit was initially imagined to help navigate amongst the may git repos scattered on the author's filesystem, but vite isn't git specific; it supports aliasing and navigating to any path.

		The techical challenge with this idea is simply that no application may alter your shell's current path.
		See 'you can do this' https://stackoverflow.com/questions/52435908/how-to-change-the-shells-current-working-directory-in-go
		
		This tool's solution is to combine the 'vit' bin with a tiny bash function 'vd' which combines 'vit' and 'cd'.

		--help : view this message\n
		vit init -> create a vit config file (typically ~/.vit/config)\n
		vit list -> list current config\n

		`)
		
	}

	// todo: support -c arg to specify a custom config file -- used for testing
	if len(args) == 1 && args[0] == "init" {
		path := getConfPath()
		if f, _ := os.Stat(path); f != nil {
			fmt.Printf("%s: File already exists. Skipping\n", path)
			os.Exit(0)
		}

		createVitConfig()
		os.Exit(0)
	}
	if len(args) == 3 && args[0] == "alias" {
		// add foo, get foo, rm foo
		if args[1] == "add" {
			printConfig()
			os.Exit(0)
		}
		if args[1] == "get" {
			getAliasPath(args[2])
			os.Exit(0)
		}
		if args[1] == "rm" {
			printConfig()
			os.Exit(0)
		}
		log.Fatal(fmt.Errorf("alias: invalid arguments"))
	}

	if len(args) == 1 {
		getAliasPath(args[0])
		os.Exit(0)
	}

	log.Fatal(fmt.Errorf("cmd: invalid arguments"))
}
