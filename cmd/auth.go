package cmd

import (
	"fmt"
	"time"

	"github.com/shanghai-edu/ldap-test-tool/g"
	"github.com/shanghai-edu/ldap-test-tool/models"
	"github.com/spf13/cobra"
)

var username, password, users string

func PrintStart(action string) {
	fmt.Printf("LDAP %s Start \n", action)
	fmt.Println("==================================")
	fmt.Println("")
}

func PrintEnd(action string, startTime time.Time) {
	endTime := time.Now()
	fmt.Println("")
	fmt.Println("==================================")
	fmt.Printf("LDAP %s Finished, Time Usage %s \n", action, endTime.Sub(startTime))
}
func init() {
	authUserCmd.Flags().StringVarP(&username, "username", "u", "", "the username for auth(required)")
	authUserCmd.Flags().StringVarP(&password, "password", "p", "", "the password for auth(required)")
	authUserCmd.MarkFlagRequired("username")
	authUserCmd.MarkFlagRequired("password")
	authMultiCmd.Flags().StringVarP(&users, "users", "u", "", `users for auth, split username and password by "," (required)`)
	authMultiCmd.MarkFlagRequired("users")
	rootCmd.AddCommand(authCmd)
	authCmd.AddCommand(authUserCmd)
	authCmd.AddCommand(authMultiCmd)
}

var authCmd = &cobra.Command{
	Use:   "auth",
	Short: "Auth test",
	Long:  `usage, ldap-test-tool auth [command]`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(`
  multi       Multi auth test
  single      Single auth test
`)
	},
}

var authUserCmd = &cobra.Command{
	Use:   "single",
	Short: "Single auth test",
	Long:  `usage, ldap-test-tool auth single -u useranem -p password`,
	Run: func(cmd *cobra.Command, args []string) {
		action := "Auth"
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
	Long:  `usage, ldap-test-tool auth multi -u users.txt`,
	Run: func(cmd *cobra.Command, args []string) {
		authUsers, err := g.GetUsers(users)
		if err != nil {
			fmt.Printf("Read file %s failed: %s \n", users, err.Error())
			return
		}
		action := "Multi Auth"
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
