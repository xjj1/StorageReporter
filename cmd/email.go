package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/xjj1/StorageReporter/repository"
)

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
			a, err := r.ListArraysNames() //dbops.go
			if err != nil {
				return err
			}
			if len(a) > 0 {
				fmt.Println("Arrays:")
				fmt.Println(strings.Join(a, "\n"))
			} else {
				log.Println("No arrays configured")
			}
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
