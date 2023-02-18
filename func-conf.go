package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getConfPath() string {
	home, err := os.UserHomeDir()
	if err != nil {
		exitGracefully("invalid config path")
	}
	path := home + conf
	return path
}

func readFileLines(path string) []string {
	f, err := os.Open(path)
	if err != nil {
		exitGracefully(fmt.Sprintf("%s : failed to open", path))
	}

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanLines)

	var txtlines []string
	for scanner.Scan() {
		l := scanner.Text()
		// skip empty lines
		t := strings.TrimSpace(l)
		if t == "" {
			continue
		}
		txtlines = append(txtlines, l)
	}

	f.Close()
	return txtlines
}

func parseConfFileLines(lines []string) []Alias {
	if lines[0] != "[alias]" {
		exitGracefully(fmt.Sprintf("%s : invalid vit config format (header)", getConfPath()))
	}

	alias := []Alias{}

	for i, l := range lines[1:] {
		res := strings.Split(l, "=")
		if len(res) != 2 {
			exitGracefully(fmt.Sprintf("%s : invalid vit config format (%s)", getConfPath(), l))
		}
		alias = append(alias, Alias{
			index: i,
			name:  strings.TrimSpace(res[0]),
			path:  strings.TrimSpace(res[1]),
		})
	}

	return alias
}

func getConfig() []Alias {
	path := getConfPath()
	if _, err := os.Stat(path); err != nil {
		exitGracefully(fmt.Sprintf("%s : no such file or directory. Run `vit init`", path))
	}

	ls := readFileLines(path)
	return parseConfFileLines(ls)

}

func printConfig() {
	alias := getConfig()
	for _, a := range alias {
		fmt.Printf("%d / %s = %s\n", a.index, a.name, a.path)
	}
}

func createVitConfig() {
	fmt.Println(">> createVitConfig")

	path := getConfPath()
	f, err := os.Create(path)
	if err != nil {
		exitGracefully(fmt.Sprintf("%s", err))
	}
	defer f.Close()

	_, err = f.WriteString("[alias]\n")
	if err != nil {
		exitGracefully(fmt.Sprintf("%s", err))
	}
}
