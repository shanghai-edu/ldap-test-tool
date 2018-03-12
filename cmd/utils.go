package cmd

import (
	"fmt"
	"strings"
	"time"

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
