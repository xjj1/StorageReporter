package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/pkg/errors"
	"github.com/spf13/cobra"
	"github.com/xjj1/StorageReporter/arrays"
	"github.com/xjj1/StorageReporter/connector/sshcon"
	"github.com/xjj1/StorageReporter/devices"
	"github.com/xjj1/StorageReporter/repository"
)

func newAddArrayCmd(r repository.Repository) *cobra.Command {
	var addCmd = &cobra.Command{
		Use:   "add",
		Short: "adds array to db",
		Long: `
		This subcommand adds array to the config :
		`,
		Example: fmt.Sprintf(`	
		%s add array,user,password
		%s add array;friendly_name,user,password`, selfName, selfName),
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				return errors.New("error in args")
			}
			var Name, Username, Password string
			Name, Username, Password, err := split3(args[0], ",")
			if err != nil {
				return errors.New("error in args")
			}
			x := strings.Split(Name, `;`)
			Friendlyname := ""

			//fmt.Println("len x=", len(x))
			if len(x) != 1 && len(x) != 2 {
				//fmt.Println("Error: \n rep4 add array,user,password")
				return errors.New("Invalid name")
			}
			if len(x) == 2 {
				Name = x[0]
				Friendlyname = x[1]
			}
			a := devices.Device{
				Name:         Name,
				Username:     Username,
				Password:     Password,
				Friendlyname: Friendlyname,
				//Type:         devices.UNKNOWN,
			}

			err = arrays.AutoDetect(sshcon.NewSSH(), &a)
			if err != nil {
				return errors.New("Cannot autodetect / unknown array")
			}
			log.Println("Detected", a.Type.String())

			err = r.AddArray(a)
			if err != nil {
				return errors.Wrap(err, "add array sql:")
			}
			log.Println(a.String(), "Added")
			return nil
		},
	}
	return addCmd
}

// Split3 splits a string to 3 strings or returns error
func split3(str, sep string) (string, string, string, error) {
	x := strings.Split(str, sep)
	if len(x) != 3 {
		//fmt.Println("Error in args")
		return "", "", "", errors.New("Error in args")
	}
	return x[0], x[1], x[2], nil
}
