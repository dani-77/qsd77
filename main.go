package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

func runQuickshell(args []string) {
	cmd := exec.Command("qs", args...)
	saida, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %v\n%s", err, string(saida))
		return
	}
	fmt.Print(string(saida))
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "qsd77",
		Short: "qsd77, a bridge to quickshell-d77",
	}

var launcherCmd = &cobra.Command{
		Use:   "launcher",
		Short: "toggle launcher",
		Run: func(cmd *cobra.Command, args []string) {
			runQuickshell([]string{"ipc", "call", "launcher", "toggle"})
		},
	}

var lockerCmd = &cobra.Command{
		Use:   "locker",
		Short: "lock screen",
		Run: func(cmd *cobra.Command, args []string) {
			runQuickshell([]string{"ipc", "call", "lockscreen", "lock"})
		},
	}

var sessionCmd = &cobra.Command{
		Use:   "session",
		Short: "toggle session menu",
		Run: func(cmd *cobra.Command, args []string) {
			runQuickshell([]string{"ipc", "call", "session", "toggle"})
		},
	}
	rootCmd.AddCommand(launcherCmd)
	rootCmd.AddCommand(lockerCmd)
	rootCmd.AddCommand(sessionCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
