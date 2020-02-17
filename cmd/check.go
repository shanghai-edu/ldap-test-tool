package cmd

import (
	"fmt"
	"os"

	"github.com/shanghai-edu/ldap-test-tool/g"
	"github.com/shanghai-edu/ldap-test-tool/models"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(checkCmd)
}

var checkCmd = &cobra.Command{
	Use:   "check",
	Short: "Check Cdap Connectivity",
	Long:  `Check Cdap Connectivity`,
	Run: func(cmd *cobra.Command, args []string) {
		_, err := models.Health_Check(g.Config().Ldap)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed: %v", err)
			return
		}
		fmt.Println("Successed")
	},
}
