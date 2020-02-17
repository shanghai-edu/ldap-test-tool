package cmd

import (
	"fmt"
	"os"

	"github.com/shanghai-edu/ldap-test-tool/g"
	"github.com/spf13/cobra"
)

var cfg string

func init() {
	cobra.OnInitialize(initConfig)
	rootCmd.PersistentFlags().StringVarP(&cfg, "config", "c", "cfg.json", "load config file. default cfg.json")
}

func initConfig() {
	g.ParseConfig(cfg)
}

var rootCmd = &cobra.Command{
	Use:   "ldap-test-tool",
	Short: "ldap-test-tool is a simple tool for ldap test",
	Long: `ldap-test-tool is a simple tool for ldap test
build by shanghai-edu.
Complete documentation is available at github.com/shanghai-edu/ldap-test-tool`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
  auth        Auth Test
  check	      Check Cdap Connectivity
  help        Help about any command
  http        Enable a http server for ldap-test-tool
  search      Search Test
  version     Print the version number of ldap-test-tool
		`)
	},
}

func Execute() {

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
