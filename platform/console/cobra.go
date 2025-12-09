package console

import (
	"github.com/spf13/cobra"
)

type ConsoleService struct {
	cmdList []*cobra.Command
}

func New() *ConsoleService {
	return &ConsoleService{}
}

func (cs *ConsoleService) Setup() {
	cs.cmdList = []*cobra.Command{}
}

func (cs *ConsoleService) RegisterCommand(cmd *cobra.Command) {
	cs.cmdList = append(cs.cmdList, cmd)
}

func (ConsoleService) GetCommandInstance() *cobra.Command {
	return &cobra.Command{}
}

func (cs *ConsoleService) Run() error {
	rootCmd := &cobra.Command{}

	for _, cmd := range cs.cmdList {
		rootCmd.AddCommand(cmd)
	}

	return rootCmd.Execute()
}
