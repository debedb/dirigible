package dirigible

import (
	"strings"

	"github.com/abiosoft/ishell"
)

func RunShell(exec *Executor) {
	shell := ishell.New()
	shell.SetPrompt("dirigible> ")

	// display welcome info.
	shell.Println("Hello.\n")

	// register a function for "greet" command.
	shell.AddCmd(&ishell.Cmd{
		Name: "greet",
		Help: "greet user",
		Func: func(c *ishell.Context) {
			c.Println("Hello", strings.Join(c.Args, " "))
		},
	})

	// run shell
	shell.Run()
}
