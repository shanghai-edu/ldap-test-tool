package models

import (
	"testing"
)

var openldap = &LDAP_CONFIG{
	Addr:       "ldap.example.org:389",
	BaseDn:     "dc=example,dc=org",
	BindDn:     "cn=manager,dc=example,dc=org",
	BindPass:   "password",
	AuthFilter: "(&(uid=%s))",
	Attributes: []string{},
	TLS:        false,
	StartTLS:   false,
}

var ad = &LDAP_CONFIG{
	Addr:       "ldap.example.org:636",
	BaseDn:     "dc=example.dc=org",
	BindDn:     "manager@example.org",
	BindPass:   "password",
	AuthFilter: "(&(sAMAccountName=%s))",
	Attributes: []string{"sAMAccountName", "displayName", "mail"},
	TLS:        true,
	StartTLS:   false,
}

func Test_ldap_auth_ad(t *testing.T) {
	err := ad.Connect()
	defer ad.Close()
	if err != nil {
		t.Error(err)
		return
	}

	success, err := ad.Auth("user1", "pass")
	t.Log(success, err)
	success, err = ad.Auth("user2", "pass")
	t.Log(success, err)

}

func Test_ldap_auth_openldap(t *testing.T) {
	err := openldap.Connect()
	defer openldap.Close()
	if err != nil {
		t.Error(err)
		return
	}
	success, err := openldap.Auth("user1", "pass")
	t.Log(success, err)
	success, err = openldap.Auth("user2", "pass")
	t.Log(success, err)
	openldap.Close()
}

func Test_ldap_search_ad(t *testing.T) {
	err := ad.Connect()
	defer ad.Close()
	if err != nil {
		t.Error(err)
		return
	}
	user, err := ad.Search_User("user1")
	t.Log(user, err)
	user, err = ad.Search_User("user2")
	t.Log(user, err)
	res, err := ad.Search("(mail=user1@example.org)")
	t.Log(res, err)

	ad.Close()
}

func Test_ldap_search_openldap(t *testing.T) {
	err := openldap.Connect()
	defer openldap.Close()
	if err != nil {
		t.Error(err)
		return
	}
	user, err := openldap.Search_User("user1")
	t.Log(user, err)
	user, err = openldap.Search_User("user2")
	t.Log(user, err)
	res, err := openldap.Search("(sn=冯冯)")
	t.Log(res, err)

	openldap.Close()
}

var user1 = User{
	Username: "user1",
	Password: "pass",
}

var user2 = User{
	Username: "user2",
	Password: "pass",
}

var user3 = User{
	Username: "cannotsuchuser",
	Password: "123456",
}

func Test_multi_auth(t *testing.T) {
	userlist := []User{user1, user2, user3}
	res, err := Multi_Auth(openldap, userlist)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(res)

}

func Test_multi_search_user(t *testing.T) {
	userlist := []string{"user1", "user2", "user3", "cannotsuchuser"}
	res, err := Multi_Search_User(ad, userlist)
	if err != nil {
		t.Error(err)
		return
	}
	t.Log(res)
}
