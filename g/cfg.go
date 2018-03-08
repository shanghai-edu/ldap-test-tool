package g

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"strings"
	"sync"

	"github.com/shanghai-edu/ldap-test-tool/models"
	"github.com/toolkits/file"
)

type GlobalConfig struct {
	Ldap *models.LDAP_CONFIG `jspn:"ldap"`
	Http *HttpConfig         `json:"http"`
}

type HttpConfig struct {
	Listen string `json:"listen"`
}

var (
	ConfigFile string
	config     *GlobalConfig
	lock       = new(sync.RWMutex)
)

func Config() *GlobalConfig {
	lock.RLock()
	defer lock.RUnlock()
	return config
}

func GetLines(filePath string) ([]string, error) {
	result := []string{}
	b, err := ioutil.ReadFile(filePath)
	if err != nil {
		return result, err
	}
	s := string(b)
	for _, lineStr := range strings.Split(s, "\n") {
		lineStr = strings.TrimSpace(lineStr)
		if lineStr == "" {
			continue
		}
		result = append(result, lineStr)
	}
	return result, nil
}

func GetUsers(filePath string) ([]models.User, error) {
	var user models.User
	users := []models.User{}

	lines, err := GetLines(filePath)
	if err != nil {
		return users, err
	}

	for _, line := range lines {
		if strings.Contains(line, ",") {
			strList := strings.Split(line, ",")
			user.Username = strList[0]
			user.Password = strList[1]
			users = append(users, user)
		}
	}
	return users, nil
}
func ParseConfig(cfg string) {
	if cfg == "" {
		log.Fatalln("use -c to specify configuration file")
	}

	if !file.IsExist(cfg) {
		log.Fatalln("config file:", cfg, "is not existent. maybe you need `mv cfg.example.json cfg.json`")
	}

	ConfigFile = cfg

	configContent, err := file.ToTrimString(cfg)
	if err != nil {
		log.Fatalln("read config file:", cfg, "fail:", err)
	}

	var c GlobalConfig
	err = json.Unmarshal([]byte(configContent), &c)
	if err != nil {
		log.Fatalln("parse config file:", cfg, "fail:", err)
	}

	lock.Lock()
	defer lock.Unlock()

	config = &c

}
