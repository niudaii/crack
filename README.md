<h1 align="center">
 crack
</h1>

<h4 align="center">常见服务弱口令爆破工具</h4>

<p align="center">
  <a href="https://opensource.org/licenses/MIT">
    <img src="https://img.shields.io/badge/license-MIT-_red.svg">
  </a>
  <a href="https://github.com/niudaii/crack/actions">
    <img src="https://img.shields.io/github/workflow/status/niudaii/crack/Go?style=flat-square" alt="Github Actions">
  </a>
  <a href="https://goreportcard.com/badge/github.com/niudaii/crack">
    <img src="https://goreportcard.com/badge/github.com/niudaii/crack">		
  </a>
  <a href="https://github.com/niudaii/crack/releases">
    <img src="https://img.shields.io/github/release/niudaii/crack/all.svg?style=flat-square">
  </a>
</p>

## 功能

- 支持常见服务口令爆破（未授权检测）
  - ftp
  - ssh
  - wmi
  - smb
  - mssql
  - oracle
  - mysql
  - rdp
  - postgres
  - redis
  - memcached
  - mongodb

- 支持彩色输出
- 全平台支持

## 使用

```
➜  crack git:(main) ✗ ./crack -h                  
Cracker

Usage:
  ./crack [flags]

Flags:
INPUT:
   -i, -input string       crack service input(example: -i '127.0.0.1:3306', -i '127.0.0.1:3307|mysql')
   -f, -input-file string  crack service file(example: -f 'xxx.txt')
   -m, -module string      choose module to crack(ftp,ssh,wmi,mssql,oracle,mysql,rdp,postgres,redis,memcached,mongodb) (default "all")
   -user string            user(example: -user 'admin,root')
   -pass string            pass(example: -pass 'admin,root')
   -user-file string       user file(example: -user-file 'user.txt')
   -pass-file string       pass file(example: -pass-file 'pass.txt')
   -crack-all              crack all user:pass

CONFIG:
   -threads int  number of threads (default 1)
   -delay int    delay between requests in seconds (0 to disable)
   -timeout int  timeout in seconds (default 10)

OUTPUT:
   -o, -output string  output file to write found results (default "crack.txt")
   -nc, -no-color      disable colors in output

DEBUG:
   -silent  show only results in output
   -debug   show debug output
```



## 参考

https://github.com/netxfly/x-crack

https://github.com/shadow1ng/fscan