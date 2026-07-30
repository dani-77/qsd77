package main

import (
	"fmt"
	"os"
	"os/exec"

	"github.com/spf13/cobra"
)

var (
	version = "1.3.3"
	config  string
)

func quickshellEnv() []string {
	xdgConfigDirs := os.Getenv("XDG_CONFIG_DIRS")
	if xdgConfigDirs == "" {
		xdgConfigDirs = "/etc/xdg"
	}
	return append(os.Environ(), "XDG_CONFIG_DIRS=/usr/share:"+xdgConfigDirs)
}

func runQuickshell(args []string) {
	fullArgs := append([]string{"-c", config}, args...)
	cmd := exec.Command("qs", fullArgs...)
	cmd.Env = quickshellEnv()

	saida, err := cmd.CombinedOutput()
	if err != nil {
		fmt.Printf("Error: %v\n%s", err, string(saida))
		return
	}
	fmt.Print(string(saida))
}

func launchQuickshell(cfg string) {
	cmd := exec.Command("qs", "-c", cfg)
	cmd.Env = quickshellEnv()
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr
	cmd.Stdin = os.Stdin
	if err := cmd.Run(); err != nil {
		fmt.Printf("Error: %v\n", err)
		os.Exit(1)
	}
}

func main() {
	var rootCmd = &cobra.Command{
		Use:   "qsd77",
		Short: "qsd77, a bridge to quickshell-d77",
	}

	rootCmd.PersistentFlags().StringVarP(&config, "config", "c", "quickshell-d77", "quickshell shell/config name (e.g. utumno, quickshell-d77)")

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
	
var wallpaperCmd = &cobra.Command{
		Use:   "wallpaper",
		Short: "toggle wallpaper menu",
		Run: func(cmd *cobra.Command, args []string) {
			runQuickshell([]string{"ipc", "call", "wallpaper", "toggle"})
		},
	}

var dashboardCmd = &cobra.Command{
		Use:   "dashboard",
		Short: "toggle dashboard",
		Run: func(cmd *cobra.Command, args []string) {
			runQuickshell([]string{"ipc", "call", "dashboard", "toggle"})
		},
	}

var ollamaCmd = &cobra.Command{
		Use:   "ollama",
		Short: "toggle ollama chat popup",
		Run: func(cmd *cobra.Command, args []string) {
			runQuickshell([]string{"ipc", "call", "ollamachat", "toggle"})
		},
	}

	var runCmd = &cobra.Command{
		Use:   "run [config]",
		Short: "launch a quickshell config (defaults to -c/--config, e.g. quickshell-d77)",
		Args:  cobra.MaximumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			cfg := config
			if len(args) > 0 {
				cfg = args[0]
			}
			launchQuickshell(cfg)
		},
	}

	var ipcCmd = &cobra.Command{
		Use:   "ipc",
		Short: "raw passthrough to quickshell IPC",
	}

	var ipcCallCmd = &cobra.Command{
		Use:   "call <target> <action> [args...]",
		Short: "call any quickshell IPC target/action directly",
		Args:  cobra.MinimumNArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			runQuickshell(append([]string{"ipc", "call"}, args...))
		},
	}

	ipcCmd.AddCommand(ipcCallCmd)

	rootCmd.AddCommand(runCmd)
	rootCmd.AddCommand(launcherCmd)
	rootCmd.AddCommand(lockerCmd)
	rootCmd.AddCommand(sessionCmd)
	rootCmd.AddCommand(wallpaperCmd)
	rootCmd.AddCommand(dashboardCmd)
	rootCmd.AddCommand(ollamaCmd)
	rootCmd.AddCommand(ipcCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
