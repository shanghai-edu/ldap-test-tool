 ## ldap-test-tool

一个轻量级的 ldap 测试工具

支持：
- ldap 认证
- ldap 查询(默认基于用户)
- 自定义 filter 的 ldap 查询
- 多用户的批量 ldap 认证
- 多用户的批量 ldap 查询
- 支持批量查询结果输出到 csv
- REST API

#### 编译
```
go get ./...
go build
```

#### release
可以直接下载编译好的 release 版本

提供 win64 和 linux64 两个平台的可执行文件

https://github.com/shanghai-edu/ldap-test-tool/releases/

#### 配置文件
默认配置文件为目录下的 `cfg.json`，也可以使用 `-c` 或 `--config` 来加载自定义的配置文件。

openldap 配置示例
```
{
    "ldap": {
        "addr": "ldap.example.org:389",
        "baseDn": "dc=example,dc=org",
        "bindDn": "cn=manager,dc=example,dc=org",
        "bindPass": "password",
        "authFilter": "(&(uid=%s))",
        "attributes": ["uid", "cn", "mail"],
        "tls":        false,
        "startTLS":   false
    },
    "http": {
        "listen": "0.0.0.0:8888"
    }
}
```
AD 配置示例
```
{
    "ldap": {
        "addr": "ad.example.org:389",
        "baseDn": "dc=example,dc=org",
        "bindDn": "manager@example.org",
        "bindPass": "password",
        "authFilter": "(&(sAMAccountName=%s))",
        "attributes": ["sAMAccountName", "displayName", "mail"],
        "tls":        false,
        "startTLS":   false
    },
    "http": {
        "listen": "0.0.0.0:8888"
    }
}
```


#### 命令体系
命令行部分使用 [cobra](github.com/spf13/cobra) 框架，可以使用 `help` 命令查看命令的使用方式
```
# ./ldap-test-tool help
ldap-test-tool is a simple tool for ldap test
build by shanghai-edu.
Complete documentation is available at github.com/shanghai-edu/ldap-test-tool

Usage:
  ldap-test-tool [flags]
  ldap-test-tool [command]

Available Commands:
  auth        Auth Test
  help        Help about any command
  http        Enable a http server for ldap-test-tool
  search      Search Test
  version     Print the version number of ldap-test-tool

Flags:
  -c, --config string   load config file. default cfg.json (default "cfg.json")
  -h, --help            help for ldap-test-tool

Use "ldap-test-tool [command] --help" for more information about a command.
```

#### 认证
```
./ldap-test-tool auth -h
Auth Test

Usage:
  ldap-test-tool auth [flags]
  ldap-test-tool auth [command]

Available Commands:
  multi       Multi Auth Test
  single      Single Auth Test

Flags:
  -h, --help   help for auth

Global Flags:
  -c, --config string   load config file. default cfg.json (default "cfg.json")

Use "ldap-test-tool auth [command] --help" for more information about a command.
```
##### 单用户测试
命令行说明
```
Single Auth Test

Usage:
  ldap-test-tool auth single [username] [password] [flags]

Flags:
  -h, --help   help for single

Global Flags:
  -c, --config string   load config file. default cfg.json (default "cfg.json")
```
示例
```
./ldap-test-tool auth single qfeng 123456
LDAP Auth Start 
==================================

qfeng auth test successed 

==================================
LDAP Auth Finished, Time Usage 47.821884ms 
```
##### 批量测试
命令行说明
```
# ./ldap-test-tool auth multi -h
Multi Auth Test

Usage:
  ldap-test-tool auth multi [filename] [flags]

Flags:
  -h, --help   help for multi

Global Flags:
  -c, --config string   load config file. default cfg.json (default "cfg.json")
```
示例
```
# cat authusers.txt 
qfeng,123456
qfengtest,111111
```
用户名和密码以逗号分隔(csv风格)
authusers.txt 中有两个用户，密码正确的 qfeng 和密码错误的 qfengtest
```
# ./ldap-test-tool auth multi authusers.txt 
LDAP Multi Auth Start 
==================================

Successed count 1 
Failed count 1 
Failed users:
 -- User: qfengtest , Msg: Cannot find such user 

==================================
LDAP Multi Auth Finished, Time Usage 49.582994ms 
```
#### 查询
```
# ./ldap-test-tool search -h
Search Test

Usage:
  ldap-test-tool search [flags]
  ldap-test-tool search [command]

Available Commands:
  filter      Search By Filter
  multi       Search Multi Users
  user        Search Single User

Flags:
  -h, --help   help for search

Global Flags:
  -c, --config string   load config file. default cfg.json (default "cfg.json")

Use "ldap-test-tool search [command] --help" for more information about a command.
[root@wiki-qfeng ldap-test-tool]# 
```
##### 单用户查询
命令行说明
```
# ./ldap-test-tool search user -h
Search Single User

Usage:
  ldap-test-tool search user [username] [flags]

Flags:
  -h, --help   help for user

Global Flags:
  -c, --config string   load config file. default cfg.json (default "cfg.json")
[root@wiki-qfeng ldap-test-tool]# 
```
示例
```
# ./ldap-test-tool search user qfeng
LDAP Search Start 
==================================


DN: uid=qfeng,ou=people,dc=example,dc=org
Attributes:
 -- uid  : qfeng 
 -- cn   : 冯骐测试 
 -- mail : qfeng@example.org


==================================
LDAP Search Finished, Time Usage 44.711268ms 
```
PS: 如果属性有多值，将以 `;` 分割
##### LDAP Filter 查询
```
# ./ldap-test-tool search filter -h
Search By Filter

Usage:
  ldap-test-tool search filter [searchFilter] [flags]

Flags:
  -h, --help   help for filter

Global Flags:
  -c, --config string   load config file. default cfg.json (default "cfg.json")
```
示例
```
# ./ldap-test-tool search filter "(cn=*测试)"
LDAP Search By Filter Start 
==================================


DN: uid=test1,ou=people,dc=example,dc=org
Attributes:
 -- uid  : test1 
 -- cn   : 一号测试 
 -- mail : test1@example.org 


DN: uid=test2,ou=people,dc=example,dc=org
Attributes:
 -- uid  : test2 
 -- cn   : 二号测试 
 -- mail : test2@example.org 


DN: uid=test3,ou=people,dc=example,dc=org
Attributes:
 -- uid  : test3
 -- cn   : 三号测试 
 -- mail : test3@example.org 

results count  3

==================================
LDAP Search By Filter Finished, Time Usage 46.071833ms 
```
##### 批量查询测试
命令行说明
```
# ./ldap-test-tool search multi -h
Search Multi Users

Usage:
  ldap-test-tool search multi [filename] [flags]

Flags:
  -f, --file   output search to users.csv, failed search to failed.csv
  -h, --help   help for multi

Global Flags:
  -c, --config string   load config file. default cfg.json (default "cfg.json")
```
示例
```
# cat searchusers.txt 
qfeng
qfengtest
nofounduser
```
searchuser.txt 中有三个用户，其中 nofounduser 是不存在的用户
```
# ldap-test-tool.exe search multi .\searchusers.txt
LDAP Multi Search Start
==================================

Successed users:

DN: uid=qfeng,ou=people,dc=example,dc=org
Attributes:
 -- uid  : qfeng
 -- cn   : 冯骐
 -- mail : qfeng@example.org


DN: uid=qfengtest,ou=people,dc=example,dc=org
Attributes:
 -- uid  : qfengtest
 -- cn   : 冯骐测试
 -- mail : qfeng@example.org

nofounduser : Cannot find such user

Successed count 2
Failed count 1

==================================
LDAP Multi Search Finished, Time Usage 134.744ms
```
当使用 `-f` 选项时，查询的结果将输出到 `csv` 中。`csv` 将以配置文件中 `attributes` 的属性作为 title。因此当使用 `-f` 选项时，`attributes` 不得为空。
```
# ./ldap-test-tool search multi searchusers.txt -f
LDAP Multi Search Start 
==================================

OutPut to csv successed

==================================
LDAP Multi Search Finished, Time Usage 88.756956ms 

# ls | grep csv
failed.csv
users.csv
```

#### HTTP API
HTTP API 部分使用 [beego](https://github.com/astaxie/beego) 框架
使用如下命令开启 HTTP API
```
# ldap-test-tool.exe http
2018/03/12 14:30:25 [I] http server Running on http://0.0.0.0:8888
```
##### 健康状态
检测 ldap 健康状态
```
# curl http://127.0.0.1:8888/api/v1/ldap/health   
{
  "msg": "ok",
  "success": true
}
```
##### 查询用户
查询单个用户信息
```
# curl  http://127.0.0.1:8888/api/v1/ldap/search/user/qfeng
{
  "user": {
    "dn": "uid=qfeng,ou=people,dc=example,dc=org",
    "attributes": {
      "cn": [
        "冯骐"
      ],
      "mail": [
        "qfeng"
      ],
      "uid": [
        "qfeng"
      ]
    }
  },
  "success": true
}
```
##### Filter 查询
根据 LDAP Filter 查询
```
# curl  http://127.0.0.1:8888/api/v1/ldap/search/filter/\(cn=*测试\)
{
  "results": [
    {
      "dn": "uid=test1,ou=people,dc=example,dc=org",
      "attributes": {
        "cn": [
          "一号测试"
        ],
        "mail": [
          "test1@example.org"
        ],
        "uid": [
          "test1"
        ]
      }
    },
    {
      "dn": "uid=test2,ou=people,dc=example,dc=org",
      "attributes": {
        "cn": [
          "二号测试"
        ],
        "mail": [
          "test2@example.org"
        ],
        "uid": [
          "test2"
        ]
      }
    },
    {
      "dn": "uid=test3,ou=people,dc=example,dc=org",
      "attributes": {
        "cn": [
          "三号测试"
        ],
        "mail": [
          "test3@example.org"
        ],
        "uid": [
          "test3"
        ]
      }
    },
  ],
  "success": true
}
```
##### 多用户查询
同时查询多个用户，以 `application/json` 方式发送请求数据，请求数据示例
```
["qfeng","qfengtest","nofounduser"]
```
curl 示例
```
# curl -X POST  -H 'Content-Type:application/json' -d '["qfeng","qfengtest","nofounduser"]' http://127.0.0.1:8888/api/v1/ldap/search/multi
{
  "success": true,
  "result": {
    "successed": 2,
    "failed": 1,
    "users": [
      {
        "dn": "uid=qfeng,ou=people,dc=example,dc=org",
        "attributes": {
          "cn": [
            "冯骐"
          ],
          "mail": [
            "qfeng@example.org"
          ],
          "uid": [
            "qfeng"
          ]
        }
      },
      {
        "dn": "uid=qfengtest,ou=people,dc=example,dc=org",
        "attributes": {
          "cn": [
            "冯骐测试"
          ],
          "mail": [
            "qfeng@example.org"
          ],
          "uid": [
            "qfengtest"
          ]
        }
      }
    ],
    "failed_messages": [
      {
        "username": "nofounduser",
        "message": "Cannot find such user"
      }
    ]
  }
}
```
#### 认证
##### 单用户认证
单个用户认证测试，以 `application/json` 方式发送请求数据，请求数据示例
```
{
	"username": "qfeng",
	"password": "123456"
}
```
curl 示例
```
# curl -X POST  -H 'Content-Type:application/json' -d '{"username":"qfeng","password":"123456"}' http://127.0.0.1:8888/api/v1/ldap/auth/single
{
  "msg": "user 20150073 Auth Successed",
  "success": true
}
```
##### 多用户认证
同时发起多个用户认证测试，以 `application/json` 方式发送请求数据，请求数据示例
```
[{
	"username": "qfeng",
	"password": "123456"
}, {
	"username": "qfengtest",
	"password": "1111111"
}]
```
curl 示例
```
# curl -X POST  -H 'Content-Type:application/json' -d '[{"username":"qfeng","password":"123456"},{"username":"qfengtest","password":"1111111"}]' http://127.0.0.1:8888/api/v1/ldap/auth/multi
{
  "success": true,
  "result": {
    "successed": 1,
    "failed": 1,
    "failed_messages": [
      {
        "username": "qfengtest",
        "message": "LDAP Result Code 49 \"Invalid Credentials\": "
      }
    ]
  }
}
```
#### LICENSE
Apache License 2.0