package main

import (
	"bufio"
	"fmt"
	"os"
)

type Macro struct {
	name    string
	version Version
}

type Version int

const (
	Version202507 Version = 202507
)

var macros = []Macro{
	{name: "types", version: Version202507},
	{name: "memory", version: Version202507},
	{name: "result", version: Version202507},
}

func filterByVersion(version Version) []Macro {
	var s []Macro
	for _, macro := range macros {
		if macro.version == version {
			s = append(s, macro)
		}
	}
	return s
}

func writeMacros(w *bufio.Writer, version Version, m []Macro) {
	fmt.Fprintf(w, "#if _DYWOQCXX_VERSION >= %d\n", version)
	for _, macro := range m {
		_, err := fmt.Fprintf(w, "#    define __dywoqcxx_%s_version %d\n", macro.name, macro.version)
		if err != nil {
			panic(err)
		}
	}
	w.WriteString("#endif")
	w.WriteString("\n")
}

func main() {
	macros202507 := filterByVersion(Version202507)

	f, err := os.Create("../include/version.hxx")
	if err != nil {
		panic(err)
	}
	defer f.Close()

	w := bufio.NewWriter(f)
	defer func() {
		err = w.Flush()
		if err != nil {
			panic(err)
		}
	}()

	w.WriteString("#ifndef _DYWOQCXX_VERSION_HXX\n")
	w.WriteString("#define _DYWOQCXX_VERSION_HXX\n")
	w.WriteString("\n")
	w.WriteString("#include \"__config.hxx\"\n")
	w.WriteString("\n")

	writeMacros(w, Version202507, macros202507)

	w.WriteString("\n")
	w.WriteString("#endif\n")

	fmt.Println("succesfully generated support.hxx header")
}
