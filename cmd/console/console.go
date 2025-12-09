package console

import "github.com/newbie007fx/cinemas/platform/console"

type Console struct {
	ConsoleService *console.ConsoleService
}

func InitApp(consoleService *console.ConsoleService) {
	cs := &Console{
		ConsoleService: consoleService,
	}

	cs.initServeCommand()
	cs.initRunMigrateCommand()
	cs.initRunForceMigrateCommand()
	cs.initRunRollbackCommand()
}
