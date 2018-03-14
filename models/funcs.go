package models

type User struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type Multi_Auth_Result struct {
	Successed       int              `json:"successed"`
	Failed          int              `json:"failed"`
	Failed_Messages []Failed_Message `json:"failed_messages"`
}

type Multi_Search_User_Result struct {
	Successed       int              `json:"successed"`
	Failed          int              `json:"failed"`
	Users           []LDAP_RESULT    `json:"users"`
	Failed_Messages []Failed_Message `json:"failed_messages"`
}
type Failed_Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

func Multi_Auth(lc *LDAP_CONFIG, userlist []User) (result Multi_Auth_Result, err error) {
	err = lc.Connect()
	defer lc.Close()

	if err != nil {
		return
	}
	var Failed_Msg Failed_Message
	for _, user := range userlist {
		success, err := lc.Auth(user.Username, user.Password)
		if success {
			result.Successed++
		} else {
			result.Failed++
			Failed_Msg.Username = user.Username
			Failed_Msg.Message = err.Error()
			result.Failed_Messages = append(result.Failed_Messages, Failed_Msg)
		}
	}
	return
}

func Multi_Search_User(lc *LDAP_CONFIG, userlist []string) (result Multi_Search_User_Result, err error) {
	err = lc.Connect()
	defer lc.Close()

	if err != nil {
		return
	}
	var Failed_Msg Failed_Message
	for _, username := range userlist {
		user, err := lc.Search_User(username)
		if err == nil {
			result.Successed++
			result.Users = append(result.Users, user)
		} else {
			result.Failed++
			Failed_Msg.Username = username
			Failed_Msg.Message = err.Error()
			result.Failed_Messages = append(result.Failed_Messages, Failed_Msg)
		}
	}
	return
}

func Single_Search(lc *LDAP_CONFIG, SearchFilter string) (results []LDAP_RESULT, err error) {
	err = lc.Connect()
	defer lc.Close()

	if err != nil {
		return
	}
	results, err = lc.Search(SearchFilter)

	return
}

func Single_Auth(lc *LDAP_CONFIG, username, password string) (success bool, err error) {
	err = lc.Connect()
	defer lc.Close()

	if err != nil {
		return
	}
	success, err = lc.Auth(username, password)

	return
}

func Single_Search_User(lc *LDAP_CONFIG, username string) (user LDAP_RESULT, err error) {
	err = lc.Connect()
	defer lc.Close()

	if err != nil {
		return
	}
	user, err = lc.Search_User(username)

	return
}

func Health_Check(lc *LDAP_CONFIG) (success bool, err error) {
	err = lc.Connect()
	defer lc.Close()

	if err != nil {
		return
	}
	return true, nil
}
