package cmd

import (
	"fmt"
	"lab/internal/lab"
	"lab/internal/remote"

	"github.com/spf13/cobra"
)

var l *lab.Lab

var cmd_remote = &cobra.Command{
	Use:   "remote",
	Short: "list all remote labs",
	Long: `by default remote only lists the remote labs.
	for detail in each sub command run [lab <command> <subcommand> --help]`,
	Run: list_remote,
}

var cmd_remote_add = &cobra.Command{
	Use:   "add [commands]",
	Short: "add remote to lab",
	Long: `by default remote only lists the remote labs.
	for detail in each sub command run [lab <command> <subcommand> --help]`,
	Run: add_remote,
}

func CMD_init(lab *lab.Lab) {
	l = lab
	var rootCmd = &cobra.Command{Use: "app"}
	rootCmd.AddCommand(cmd_remote)
	cmd_remote.AddCommand(cmd_remote_add)
	rootCmd.Execute()
}

func list_remote(cmd *cobra.Command, args []string) {
	for i, r := range l.GetAllRemote() {
		fmt.Printf("remote no.[%d]:\t%s", i, r.Name)
	}
}

func add_remote(cmd *cobra.Command, args []string) {
	fmt.Println(args)
	l.AddRemote(remote.New(args[0], args[1], args[2], args[3]))
}
