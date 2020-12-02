package cmd

import (

	//"os"

	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/cobra"
	"github.com/xjj1/StorageReporter/db"
)

var selfName string

type App struct {
	c *cobra.Command
}

func (a *App) Execute() error {
	return a.c.Execute()
}

func NewApp(r *db.Repository) *App {

	//var err error
	selfName = filepath.Base(os.Args[0])
	// if err != nil {
	// 	log.Fatalln("A strange error have occured", err)
	// }
	c := NewRootCmd()
	c.AddCommand(NewEmailCmd(r))
	c.AddCommand(NewListCmd(r))
	return &App{c}
}

func NewRootCmd() *cobra.Command {
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

func NewEmailCmd(r *db.Repository) *cobra.Command {
	var cmd = &cobra.Command{
		Use:   "email",
		Short: "Configure email settings",
		Long: `Configure email settings
		`,
		Example: `
		rep4 email rcptTo,mailserver,mailfrom,subject
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				//fmt.Println("Error: \n >>rep4 add array,user,password")
				return errors.New("error in args")
			}
			x := strings.Split(args[0], ",")
			fmt.Println("adding email ", x)
			r.AddEmail(x) // dbops.go
			return nil
		},
	}

	return cmd
}

func NewListCmd(r *db.Repository) *cobra.Command {
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

/*
	var RootCmd = &cobra.Command{
		Use:   "rep4",
		Short: "Report tool for 3PAR, MSA and Lefthand arrays",
		Long: `This application generates report for 3PAR, MSA and Lefthand arrays and sends it by email

		Add arrays :

			rep4 add array[;friendlyname],user,password

		Configure email settings :

			rep4 email rcptTo,mailserver[,subject,mailfrom]

		Run "rep4 makerep" to generate and send the report.

		Run "rep4 makerep_noxls" to generate and send the report.

		Run "rep4 help [command]" for more details

		`,
		SilenceUsage:  true,
		SilenceErrors: true,
		// Uncomment the following line if your bare application
		// has an action associated with it:
		//	Run: func(cmd *cobra.Command, args []string) { },
	}

	var emailCmd = &cobra.Command{
		Use:   "email",
		Short: "Configure email settings",
		Long: `Configure email settings
		`,
		Example: `
		rep4 email rcptTo,mailserver,mailfrom,subject
		`,
		RunE: func(cmd *cobra.Command, args []string) error {
			if len(args) != 1 {
				//fmt.Println("Error: \n >>rep4 add array,user,password")
				return errors.New("error in args")
			}
			x := strings.Split(args[0], ",")
			fmt.Println("adding email ", x)
			db.r.Addemailtodb(x) // dbops.go
			return nil
		},
	}

}

// RootCmd root command

/*

var addCmd = &cobra.Command{
	Use:   "add",
	Short: "adds array to db",
	Long: `
	This subcommand adds array to the config :
	`,
	Example: `	rep4 add array,user,password
	rep4 add array;friendly_name,user,password`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			//fmt.Println("Error: \n >>rep4 add array,user,password")
			return errors.New("error in args")
		}
		var Name, Username, Password string
		Name, Username, Password, err := Split3(args[0], ",")
		if err != nil {
			//fmt.Println("Error: \n rep4 add array,user,password")
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
		a := Device{
			Name:         Name,
			Username:     Username,
			Password:     Password,
			Friendlyname: Friendlyname,
			Type:         UNKNOWN,
		}
		//fmt.Println("add called with args : ", Name, Friendlyname, Username, Password, err)

		// dbops.go
		err = addarraytodb(&a)
		if err != nil {
			//fmt.Println("err= ", err)
			return err
		}
		return nil
	},
}

var delCmd = &cobra.Command{
	Use:     "del",
	Short:   "delete array from db",
	Long:    `This subcommand deletes array from the config :`,
	Example: `rep4 del array`,
	RunE: func(cmd *cobra.Command, args []string) error {
		if len(args) != 1 {
			return errors.New("error in args")
		}
		Name := args[0]
		return deleteArray(Name) // dbops.go
	},
}

var listCmd = &cobra.Command{
	Use:   "list",
	Short: "list configuration",
	Long:  `This command lists the configuration`,
	RunE: func(cmd *cobra.Command, args []string) error {
		a, err := listarrays_names() //dbops.go
		if err != nil {
			return err
		}
		fmt.Println(strings.Join(a, "\n"))

		rcptTo, mailserver, mailfrom, subject, _, _ := getemail()
		fmt.Println("mail configration :", rcptTo, mailserver, mailfrom, subject)
		return nil
	},
}

var makerepCmd = &cobra.Command{
	Use:   "makerep",
	Short: "Make and send the report",
	Long:  `Make and send the report`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//fmt.Println("makerep called with args : ", args)
		fmt.Println("making report")
		return MakeReportAll() // makerep.go
	},
}

//makerep_noxls
var makerepnoxlsCmd = &cobra.Command{
	Use:   "makerep_noxls",
	Short: "Make and send the report without the excel sheet",
	Long:  `Make and send the report without the excel sheet`,
	RunE: func(cmd *cobra.Command, args []string) error {
		//fmt.Println("makerep_noxls called with args : ", args)
		fmt.Println("making report (no xls)")
		return MakeReportAllNOXLS() // makerep.go
	},
}

var verCmd = &cobra.Command{
	Use:   "version",
	Short: "Show software version",
	Long:  `Show software version`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("rep4 version ", VERSION)
		// run autodetect and addarray to db
	},
}



// Split3 splits strings 3-way
func Split3(str, sep string) (string, string, string, error) {
	x := strings.Split(str, sep)
	if len(x) != 3 {
		//fmt.Println("Error in args")
		return "", "", "", errors.New("Error in args")
	}
	return x[0], x[1], x[2], nil
}
*/
