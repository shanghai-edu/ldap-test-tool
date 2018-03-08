package cmd

import (
	"github.com/shanghai-edu/ldap-test-tool/http"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(httpCmd)
}

var httpCmd = &cobra.Command{
	Use:   "http",
	Short: "Enable a http server for ldap-test-tool",
	Long:  `ldap-test-tool provide a restful api for ldap test`,
	Run: func(cmd *cobra.Command, args []string) {
		http.Start()
	},
}
