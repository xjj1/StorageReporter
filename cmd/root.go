package cmd

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/spf13/cobra"
	"github.com/xjj1/StorageReporter/repository"
)

var selfName string

type App struct {
	c *cobra.Command
}

func (a *App) Execute() error {
	return a.c.Execute()
}

func NewApp(r repository.Repository) *App {
	selfName = filepath.Base(os.Args[0])
	c := newRootCmd()
	c.AddCommand(newEmailCmd(r))
	c.AddCommand(newListCmd(r))
	c.AddCommand(newAddArrayCmd(r))
	c.AddCommand(supportedArrays)
	return &App{c}
}

func newRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   selfName,
		Short: "Reporting tool for storage arrays",
		Long: fmt.Sprintf(`This application generates report for multiple kinds of storage arrays and sends it by email
		
Add arrays : 		

	%s add array[;friendlyname],user,password
		
Configure email settings : 

	%s email rcptTo,mailserver[,mailfrom][,subject]
		
Run "%s makerep" to generate and send the report.

Run "%s help [command]" for more details

		`, selfName, selfName, selfName, selfName),
		SilenceUsage:  true,
		SilenceErrors: true,
	}
	return cmd
}

var supportedArr = `
Supported arrays:
HP 3PAR (via SSH)
HP MSA (via SSH)
HP LeftHand / StoreVirtual (must have cliq installed and in the system PATH)
HP Nimble (via SSH)
PureStorage (via SSH)
Hitachi VSP (must have horcm/horcmstart/raidcom installed, configured and in the system PATH)
`

var supportedArrays = &cobra.Command{
	Use:     "sup",
	Short:   "List supported arrays",
	Long:    supportedArr,
	Example: supportedArr,
}
