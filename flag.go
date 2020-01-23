package main

import "flag"

var (
	help string
)

func init() {
	flag.StringVar(&help, "h", "false", "This is help")
}
func main() {
	flag.Parse()

	if help != "false" {
		flag.Usage()
	}
}

func usage() {
	flag.PrintDefaults()
}
