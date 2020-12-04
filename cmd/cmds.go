package cmd

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/pkg/errors"
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

Run "%s makerep_noxls" to generate and send the report.

Run "%s help [command]" for more details

Supported arrays:
HP 3PAR (via SSH)
HP MSA (via SSH)
HP LeftHand / StoreVirtual (must have cliq installed and in the system PATH)
HP Nimble (via SSH)
PureStorage (via SSH)
Hitachi VSP (must have horcm/horcmstart/raidcom installed, configured and in the system PATH)

		`, selfName, selfName, selfName, selfName, selfName),
		SilenceUsage:  true,
		SilenceErrors: true,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		//	Run: func(cmd *cobra.Command, args []string) { },
	}
	return cmd
}

func newEmailCmd(r repository.Repository) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "email",
		Short: "Configure email settings",
		Long: `Configure email settings
		`,
		Example: fmt.Sprintf(`
		%s email rcptTo,mailserver,mailfrom,subject
		`, selfName),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				//fmt.Println("Error: \n >>rep4 add array,user,password")
				return errors.New("error in args")
			}
			x := strings.Split(args[0], ",")
			// dbops.go
			if err := r.AddEmailSettings(x); err != nil {
				return errors.Wrap(err, "add email")
			}
			fmt.Println("added email ", x)

			return nil
		},
	}

	return cmd
}

func newListCmd(r repository.Repository) *cobra.Command {
	var listCmd = &cobra.Command{
		Use:   "list",
		Short: "list configuration",
		Long:  `This command lists the configuration`,
		RunE: func(cmd *cobra.Command, args []string) error {
			// a, err := r.ListArraysNames() //dbops.go
			// if err != nil {
			// 	return err
			// }
			// fmt.Println(strings.Join(a, "\n"))
			m, err := r.GetEmail()
			if err != nil {
				return err
			}

			fmt.Println("mail configration :", m.Rcptto, m.Mailserver, m.Mailfrom, m.Subject)
			return nil
		},
	}
	return listCmd
}
