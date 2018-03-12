package cmd

import (
	"fmt"
	"time"

	"github.com/shanghai-edu/ldap-test-tool/g"
	"github.com/shanghai-edu/ldap-test-tool/models"
	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(authCmd)
	authCmd.AddCommand(authUserCmd)
	authCmd.AddCommand(authMultiCmd)
}

var authCmd = &cobra.Command{
	Use:       "auth",
	Short:     "Auth Test",
	Long:      `Usage: ldap-test-tool auth [command]`,
	Args:      cobra.OnlyValidArgs,
	ValidArgs: []string{authUserCmd.Use, authMultiCmd.Use},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
  multi       Multi auth test
  single      Single auth test
`)
	},
}

var authUserCmd = &cobra.Command{
	Use:   "single",
	Short: "Single auth test",
	Long:  `Usage: ldap-test-tool auth single [username] [password]`,
	Args:  cobra.ExactArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		action := "Auth"

		username := args[0]
		password := args[1]

		startTime := time.Now()
		PrintStart(action)

		_, err := models.Single_Auth(g.Config().Ldap, username, password)

		if err != nil {
			fmt.Printf("%s auth test failed: %s \n", username, err.Error())
			PrintEnd(action, startTime)
			return
		}
		fmt.Printf("%s auth test successed \n", username)
		PrintEnd(action, startTime)
	},
}

var authMultiCmd = &cobra.Command{
	Use:   "multi",
	Short: "Multi auth test",
	Long:  `Usage: ldap-test-tool auth multi [filename]`,
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		action := "Multi Auth"

		userlist := args[0]
		authUsers, err := g.GetUsers(userlist)
		if err != nil {
			fmt.Printf("Read file %s failed: %s \n", userlist, err.Error())
			return
		}
		startTime := time.Now()
		PrintStart(action)

		res, err := models.Multi_Auth(g.Config().Ldap, authUsers)
		if err != nil {
			fmt.Printf("Multi Auth  failed: %s \n", err.Error())
			PrintEnd(action, startTime)
			return
		}
		fmt.Printf("Successed count %d \n", res.Successed)
		fmt.Printf("Failed count %d \n", res.Failed)
		fmt.Println("Failed users:")
		for _, failed_Message := range res.Failed_Messages {
			fmt.Printf(" -- User: %s , Msg: %s \n", failed_Message.Username, failed_Message.Message)
		}
		PrintEnd(action, startTime)
	},
}
