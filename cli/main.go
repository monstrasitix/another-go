package main

import (
	"os"

	"github.com/monstrasitix/finance/cli/commands"
)

var (
	cmds = map[string]func([]string){
		"widgets": commands.RunWidgets,
	}
)

func main() {
	command := os.Args[1]
	params := os.Args[1:]

	if callback, exists := cmds[command]; exists {
		callback(params)
	}
}
