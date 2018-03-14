package cmd

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
	"time"

	"github.com/shanghai-edu/ldap-test-tool/g"
	"github.com/shanghai-edu/ldap-test-tool/models"
)

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

func inArray(str string, array []string) bool {
	for _, s := range array {
		if s == str {
			return true
		}
	}
	return false
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

func WriteUsersToCsv(users []models.LDAP_RESULT, filename string) (err error) {
	csvFile, err := os.Create(filename)
	if err != nil {
		return
	}
	defer csvFile.Close()
	csvFile.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	writer := csv.NewWriter(csvFile)
	title := make([]string, len(g.Config().Ldap.Attributes)+1)
	title[0] = "DN"
	for index, value := range g.Config().Ldap.Attributes {
		title[index+1] = value
	}
	writer.Write(title)
	for _, user := range users {
		s := make([]string, len(g.Config().Ldap.Attributes)+1)
		s[0] = user.DN
		for key, value := range user.Attributes {
			valueString := strings.Join(value, ";")
			for i, v := range title {
				if key == v {
					s[i] = valueString
				}
			}
		}
		writer.Write(s)
	}
	writer.Flush()
	return
}

func WriteFailsToCsv(msgs []models.Failed_Message, filename string) (err error) {
	csvFile, err := os.Create(filename)
	if err != nil {
		return
	}
	defer csvFile.Close()
	csvFile.WriteString("\xEF\xBB\xBF") // 写入UTF-8 BOM
	writer := csv.NewWriter(csvFile)
	title := make([]string, 2)
	title[0] = "username"
	title[1] = "message"
	writer.Write(title)
	for _, msg := range msgs {
		s := make([]string, 2)
		s[0] = msg.Username
		s[1] = msg.Message
		writer.Write(s)
	}
	writer.Flush()
	return
}
