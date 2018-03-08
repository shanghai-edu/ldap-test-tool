package models

import (
	"crypto/tls"
	"errors"
	"fmt"

	ldap "gopkg.in/ldap.v2"
)

type LDAP_CONFIG struct {
	Addr       string   `json:"addr"`
	BaseDn     string   `json:"baseDn"`
	BindDn     string   `json:"bindDn`
	BindPass   string   `json:"bindPass"`
	AuthFilter string   `json:"authFilter"`
	Attributes []string `json:"attributes"`
	TLS        bool     `json:"tls"`
	StartTLS   bool     `json:"startTLS"`
	Conn       *ldap.Conn
}

type LDAP_RESULT struct {
	DN         string              `json:"dn"`
	Attributes map[string][]string `json:"attributes"`
}

func (lc *LDAP_CONFIG) Close() {
	if lc.Conn != nil {
		lc.Conn.Close()
		lc.Conn = nil
	}
}

func (lc *LDAP_CONFIG) Connect() (err error) {
	if lc.TLS {
		lc.Conn, err = ldap.DialTLS("tcp", lc.Addr, &tls.Config{InsecureSkipVerify: true})
	} else {
		lc.Conn, err = ldap.Dial("tcp", lc.Addr)
	}
	if err != nil {
		return err
	}
	if !lc.TLS && lc.StartTLS {
		err = lc.Conn.StartTLS(&tls.Config{InsecureSkipVerify: true})
		if err != nil {
			lc.Conn.Close()
			return err
		}
	}

	err = lc.Conn.Bind(lc.BindDn, lc.BindPass)
	if err != nil {
		lc.Conn.Close()
		return err
	}
	return err
}

func (lc *LDAP_CONFIG) Auth(username, password string) (success bool, err error) {
	searchRequest := ldap.NewSearchRequest(
		lc.BaseDn, // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf(lc.AuthFilter, username), // The filter to apply
		lc.Attributes,                        // A list attributes to retrieve
		nil,
	)
	sr, err := lc.Conn.Search(searchRequest)
	if err != nil {
		return
	}
	if len(sr.Entries) == 0 {
		err = errors.New("Cannot find such user")
		return
	}
	if len(sr.Entries) > 1 {
		err = errors.New("Multi users in search")
		return
	}
	err = lc.Conn.Bind(sr.Entries[0].DN, password)
	if err != nil {
		return
	}
	//Rebind as the search user for any further queries
	err = lc.Conn.Bind(lc.BindDn, lc.BindPass)
	if err != nil {
		return
	}
	success = true
	return
}

func (lc *LDAP_CONFIG) Search_User(username string) (user LDAP_RESULT, err error) {
	searchRequest := ldap.NewSearchRequest(
		lc.BaseDn, // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		fmt.Sprintf(lc.AuthFilter, username), // The filter to apply
		lc.Attributes,                        // A list attributes to retrieve
		nil,
	)
	sr, err := lc.Conn.Search(searchRequest)
	if err != nil {
		return
	}
	if len(sr.Entries) == 0 {
		err = errors.New("Cannot find such user")
		return
	}
	if len(sr.Entries) > 1 {
		err = errors.New("Multi users in search")
		return
	}

	attributes := make(map[string][]string)
	for _, attr := range sr.Entries[0].Attributes {
		attributes[attr.Name] = attr.Values
	}

	user.DN = sr.Entries[0].DN
	user.Attributes = attributes
	return
}

func (lc *LDAP_CONFIG) Search(SearchFilter string) (results []LDAP_RESULT, err error) {
	searchRequest := ldap.NewSearchRequest(
		lc.BaseDn, // The base dn to search
		ldap.ScopeWholeSubtree, ldap.NeverDerefAliases, 0, 0, false,
		SearchFilter,  // The filter to apply
		lc.Attributes, // A list attributes to retrieve
		nil,
	)
	sr, err := lc.Conn.Search(searchRequest)
	if err != nil {
		return
	}
	if len(sr.Entries) == 0 {
		err = errors.New("Cannot find such dn")
		return
	}
	results = []LDAP_RESULT{}
	var result LDAP_RESULT
	for _, entry := range sr.Entries {
		attributes := make(map[string][]string)
		for _, attr := range entry.Attributes {
			attributes[attr.Name] = attr.Values
		}
		result.DN = entry.DN
		result.Attributes = attributes
		results = append(results, result)
	}
	return
}
