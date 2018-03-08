package cmd

import (
	"fmt"
	"strings"
	"time"

	"github.com/shanghai-edu/ldap-test-tool/g"
	"github.com/shanghai-edu/ldap-test-tool/models"
	"github.com/spf13/cobra"
)

var userlist, filter string

func init() {
	searchUserCmd.Flags().StringVarP(&username, "username", "u", "", "the username for search (required)")
	searchUserCmd.MarkFlagRequired("username")
	searchMultiCmd.Flags().StringVarP(&userlist, "userlist", "u", "", "the userlist for search (required)")
	searchMultiCmd.MarkFlagRequired("userlist")
	searchFilterCmd.Flags().StringVarP(&filter, "filter", "f", "", "the filter for search (required)")
	searchFilterCmd.MarkFlagRequired("filter")
	rootCmd.AddCommand(searchCmd)
	searchCmd.AddCommand(searchUserCmd)
	searchCmd.AddCommand(searchMultiCmd)
	searchCmd.AddCommand(searchFilterCmd)
}
func getLongestKeyLen(m map[string][]string) int {
	l := 0
	for key, _ := range m {
		if len(key) > l {
			l = len(key)
		}
	}
	return l
}

func addSpace(s string, l int) string {
	for i := 0; i < l; i++ {
		s = s + " "
	}
	return s
}

func PrintSearchResult(result models.LDAP_RESULT) {
	fmt.Println("")
	fmt.Printf("DN: %s \n", result.DN)
	fmt.Println("Attributes:")
	longestKeyLenth := getLongestKeyLen(result.Attributes)
	for key, value := range result.Attributes {
		if len(key) < longestKeyLenth {
			key = addSpace(key, (longestKeyLenth - len(key)))
		}
		valueString := strings.Join(value, ";")
		fmt.Printf(" -- %s : %s \n", key, valueString)
	}
	fmt.Println("")

}

var searchCmd = &cobra.Command{
	Use:   "search",
	Short: "Search test",
	Long:  `usage, ldap-test-tool search [command]`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf(`
  filter      Search by filter
  multi       Search multi users
  user        Search single user
`)
	},
}

var searchUserCmd = &cobra.Command{
	Use:   "user",
	Short: "Search single user",
	Long:  `usage, ldap-test-tool search user -u useranem`,
	Run: func(cmd *cobra.Command, args []string) {
		action := "Search"
		startTime := time.Now()
		PrintStart(action)

		result, err := models.Single_Search_User(g.Config().Ldap, username)

		if err != nil {
			fmt.Printf("%s Search failed: %s \n", username, err.Error())
			PrintEnd(action, startTime)
			return
		}
		PrintSearchResult(result)

		PrintEnd(action, startTime)
	},
}

var searchMultiCmd = &cobra.Command{
	Use:   "multi",
	Short: "Search multi users",
	Long:  `usage, ldap-test-tool search multi -u userlist`,
	Run: func(cmd *cobra.Command, args []string) {
		searchUsers, err := g.GetLines(userlist)
		if err != nil {
			fmt.Printf("Read file %s failed: %s \n", userlist, err.Error())
			return
		}
		action := "Multi Search"
		startTime := time.Now()
		PrintStart(action)

		res, err := models.Multi_Search_User(g.Config().Ldap, searchUsers)

		if err != nil {
			fmt.Printf("%s Search failed: %s \n", username, err.Error())
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
	Use:   "filter",
	Short: "Search by filter",
	Long:  `usage, ldap-test-tool search filter -f (cn=测试*)`,
	Run: func(cmd *cobra.Command, args []string) {
		action := "Search By Filter"
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
