package cmd

import (
	"fmt"

	"github.com/shanghai-edu/ldap-test-tool/g"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(versionCmd)
}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "Print the version number of ldap-test-tool",
	Long:  `All software has versions. This is ldap-test-tool's`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("ldap-test-tool version : ", g.VERSION)
	},
}
