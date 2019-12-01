package main

import (
	"github.com/alexandre-normand/bobino"
	"github.com/zchee/color"
)

func main() {
	a := bobino.AskForBool("Do you consider yourself old?")
	color.Magenta("Answer: %t", a)
}
