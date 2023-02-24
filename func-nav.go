package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
)

func validateAlphaNumeric(id string) {
	isAlphaNumeric := regexp.MustCompile(`^[a-zA-Z0-9.]*$`).MatchString(id)
	if !isAlphaNumeric {
		exitGracefully(fmt.Sprintf("Invalid alias '%s'. Must be alphanumeric or '.'", id))
	}
}

func getAlias(id string) Alias {
	validateAlphaNumeric(id)
	alias := getConfig()

	if index, err := strconv.Atoi(id); err == nil {
		if index < len(alias) {
			return alias[index]
		}
	}

	for i := range alias {
		if alias[i].name == id {
			return alias[i]
		}
	}

	fmt.Fprintf(os.Stderr, "vite alias `%s` not found\n", id)

	return Alias{
		index: 0,
		name:  "current",
		path:  ".",
	}
}

func getAliasPath(id string) {
	a := getAlias(id)
	fmt.Println(a.path)
}

func settAliasPath(id string) {
	a := getAlias(id)
	fmt.Println(a.path)
}
