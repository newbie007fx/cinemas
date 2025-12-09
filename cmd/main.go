package main

import (
	"github.com/newbie007fx/cinemas/cmd/console"
	consoleService "github.com/newbie007fx/cinemas/platform/console"
)

func main() {
	cs := consoleService.New()
	cs.Setup()

	console.InitApp(cs)

	cs.Run()
}
