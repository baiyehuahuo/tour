package main

import (
	"errors"
	"flag"
	"log"
)

type Name string

func (n *Name) String() string {
	return string(*n)
}

func (n *Name) Set(s string) error {
	if len(*n) > 0 {
		return errors.New("name flag already set")
	}
	*n = Name("Hello, " + s)
	return nil
}

var _ flag.Value = (*Name)(nil)

func main() {
	var name Name
	flag.Var(&name, "name", "帮助信息")
	flag.Parse()
	log.Printf("name is: %s", name)
}

func SubCommand() {
	var name string
	flag.Parse()
	args := flag.Args()
	if len(args) == 0 {
		return
	}

	switch args[0] {
	case "go":
		goCmd := flag.NewFlagSet("go", flag.ExitOnError)
		goCmd.StringVar(&name, "name", "Go 语言", "帮助信息")
		_ = goCmd.Parse(args[1:])
	case "php":
		phpCmd := flag.NewFlagSet("php", flag.ExitOnError)
		phpCmd.StringVar(&name, "n", "PHP 语言", "帮助信息")
		_ = phpCmd.Parse(args[1:])
	}

	log.Printf("name: %s", name)
}
