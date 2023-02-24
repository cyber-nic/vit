package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"text/tabwriter"
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

func normalizePath(path string) string {
	abs, err := filepath.Abs(path)
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("%s : failed to compute absolute path", path))
		os.Exit(1)
	}

	fileInfo, err := os.Stat(abs)
	if err != nil {
		fmt.Fprintln(os.Stderr, fmt.Errorf("%s : invalid path", abs))
		os.Exit(1)
	}

	if fileInfo.IsDir() {
		return abs
	}

	fmt.Fprintln(os.Stderr, fmt.Errorf("Invalid path"))
	os.Exit(1)
	return ""
}

func getAliasFromPath(path string) string {
	if path == "/" {
		return "root"
	}

	res := strings.Split(path, "/")
	r := res[len(res)-1]
	return r
}

func testAlias(a string) {
	validateAlphaNumeric(a)

	conf := getConfig()
	for _, c := range conf {
		if a == c.name {
			fmt.Fprintln(os.Stderr, fmt.Errorf("Alias '%s' already in use", c.name))
			os.Exit(1)
		}
	}
}

func addAliasCurrentPath(p string) {
	n := normalizePath(p)
	a := getAliasFromPath(n)
	testAlias(a)
	appendAlias(a, n)
}

func addAliasNamedPath(p string, a string) {
	n := normalizePath(p)
	testAlias(a)
	appendAlias(a, n)
}

func appendAlias(a string, p string) {
	confPath := getConfPath()
	f, err := os.OpenFile(confPath, os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	if _, err = f.WriteString(fmt.Sprintf("  %s = %s\n", a, p)); err != nil {
		panic(err)
	}

}

func printConfig() {
	conf := getConfig()
	w := new(tabwriter.Writer)
	// minwidth, tabwidth, padding, padchar, flags
	w.Init(os.Stdout, 3, 8, 0, '\t', 0)

	for _, a := range conf {
		// fmt.Printf("%d / %s = %s\n\n", a.index, a.name, a.path)
		fmt.Fprintln(w, fmt.Sprintf("%d\t%s\t%s", a.index, a.name, a.path))

	}
	w.Flush()
}

func createVitConfig() {
	// fmt.Println(">> createVitConfig")

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

// sort
func sortConfItems() {
	// todo
}
