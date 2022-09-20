<h1 align="center">
 crack
</h1>

<h4 align="center">常见服务弱口令爆破工具</h4>

<p align="center">
  <a href="https://opensource.org/licenses/MIT">
    <img src="https://img.shields.io/badge/license-MIT-_red.svg">
  </a>
  <a href="https://goreportcard.com/report/github.com/niudaii/crack">
    <img src="https://goreportcard.com/badge/github.com/niudaii/crack?style=flat-square">		
  </a>
  <a href="https://github.com/niudaii/crack/actions">
    <img src="https://img.shields.io/github/workflow/status/niudaii/crack/Release?style=flat-square" alt="Github Actions">
  </a>
  <a href="https://github.com/niudaii/crack/releases">
    <img src="https://img.shields.io/github/release/niudaii/crack/all.svg?style=flat-square">
  </a>
  <a href="https://github.com/niudaii/crack/releases">
  	<img src="https://img.shields.io/github/downloads/niudaii/crack/total">
  </a>
</p>


## 功能

- 支持常见服务口令爆破（未授权检测）
  - ftp
  - ssh
  - wmi
  - wmihash
  - smb
  - mssql
  - oracle
  - mysql
  - rdp
  - postgres
  - redis
  - memcached
  - mongodb
- 多线程爆破，支持进度条
- 全部插件测试用例（[pkg/crack/plugins/plugins_test.go](https://github.com/niudaii/crack/blob/main/pkg/crack/plugins/plugins_test.go)）
- API调用，可参考（[internal/runner/runner.go](https://github.com/niudaii/crack/blob/main/internal/runner/runner.go)）

## 使用

```
➜  crack ./crack -h
Service cracker

Usage:
  ./crack [flags]

Flags:
INPUT:
   -i, -input string       crack service input(example: -i '127.0.0.1:3306', -i '127.0.0.1:3307|mysql')
   -f, -input-file string  crack services file(example: -f 'xxx.txt')
   -m, -module string      choose one module to crack(ftp,ssh,wmi,mssql,oracle,mysql,rdp,postgres,redis,memcached,mongodb) (default "all")
   -user string            user(example: -user 'admin,root')
   -pass string            pass(example: -pass 'admin,root')
   -user-file string       user file(example: -user-file 'user.txt')
   -pass-file string       pass file(example: -pass-file 'pass.txt')

CONFIG:
   -threads int  number of threads (default 1)
   -timeout int  timeout in seconds (default 10)
   -delay int    delay between requests in seconds (0 to disable)
   -crack-all    crack all user:pass

OUTPUT:
   -o, -output string  output file to write found results (default "crack.txt")
   -nc, -no-color      disable colors in output

DEBUG:
   -silent  show only results in output
   -debug   show debug output
```

## 截图

![image-20220903092817097](https://nnotes.oss-cn-hangzhou.aliyuncs.com/notes/image-20220903092817097.png)

## 说明

已经停止更新，该项目作为 [zpscan](https://github.com/niudaii/zpscan) 的模块之一，后续更新参考 zpscan。

## 参考

https://github.com/netxfly/x-crack

https://github.com/shadow1ng/fscan