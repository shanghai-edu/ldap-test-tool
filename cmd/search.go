package cmd

import (
	"fmt"
	"os"
	"time"

	"github.com/shanghai-edu/ldap-test-tool/g"
	"github.com/shanghai-edu/ldap-test-tool/models"
	"github.com/spf13/cobra"
)

var csvFile bool

func init() {
	searchMultiCmd.Flags().BoolVarP(&csvFile, "file", "f", false, "output search to users.csv, failed search to failed.csv")
	rootCmd.AddCommand(searchCmd)
	searchCmd.AddCommand(searchUserCmd)
	searchCmd.AddCommand(searchMultiCmd)
	searchCmd.AddCommand(searchFilterCmd)
}

var searchCmd = &cobra.Command{
	Use:       "search",
	Short:     "Search Test",
	Args:      cobra.OnlyValidArgs,
	ValidArgs: []string{searchUserCmd.Use, searchFilterCmd.Use, searchMultiCmd.Use},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`
  filter      Search By Filter
  multi       Search Multi Users
  user        Search Single User
`)
	},
}

var searchUserCmd = &cobra.Command{
	Use:   "user [username]",
	Short: "Search Single User",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		action := "Search"

		username := args[0]
		startTime := time.Now()
		PrintStart(action)

		result, err := models.Single_Search_User(g.Config().Ldap, username)

		if err != nil {
			fmt.Fprintf(os.Stderr, "%s Search failed: %s \n", username, err.Error())
			PrintEnd(action, startTime)
			return
		}
		PrintSearchResult(result)

		PrintEnd(action, startTime)
	},
}

var searchMultiCmd = &cobra.Command{
	Use:   "multi [filename]",
	Short: "Search Multi Users",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		action := "Multi Search"

		userlist := args[0]
		searchUsers, err := g.GetLines(userlist)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Read file %s failed: %s \n", userlist, err.Error())
			return
		}
		startTime := time.Now()
		PrintStart(action)

		res, err := models.Multi_Search_User(g.Config().Ldap, searchUsers)

		if err != nil {
			fmt.Fprintf(os.Stderr, "Multi Search failed: %s \n", err.Error())
			PrintEnd(action, startTime)
			return
		}
		if csvFile {
			err = WriteUsersToCsv(res.Users, g.USERS_CSV)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Open file %s failed: \n", g.USERS_CSV)
				return
			}
			if len(res.Failed_Messages) > 0 {
				err = WriteFailsToCsv(res.Failed_Messages, g.FAILED_CSV)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Open file %s failed: \n", g.FAILED_CSV)
					return
				}
			}
			fmt.Println("OutPut to csv successed")
			PrintEnd(action, startTime)
			return
		}
		fmt.Println("Successed users:")
		for _, user := range res.Users {
			PrintSearchResult(user)
		}

		for _, failed_Message := range res.Failed_Messages {
			fmt.Printf("%s : %s \n", failed_Message.Username, failed_Message.Message)
		}
		fmt.Println("")
		fmt.Printf("Successed count %d \n", res.Successed)
		fmt.Printf("Failed count %d \n", res.Failed)

		PrintEnd(action, startTime)

	},
}

var searchFilterCmd = &cobra.Command{
	Use:   "filter [searchFilter]",
	Short: "Search By Filter",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		action := "Search By Filter"

		filter := args[0]
		startTime := time.Now()
		PrintStart(action)

		results, err := models.Single_Search(g.Config().Ldap, filter)

		if err != nil {
			fmt.Printf("%s Search failed: %s \n", filter, err.Error())
			PrintEnd(action, startTime)
			return
		}
		for _, res := range results {
			PrintSearchResult(res)
		}
		fmt.Println("results count ", len(results))
		PrintEnd(action, startTime)
	},
}
