---
title: go-crack 介绍
date: 2021-04-24 09:37:41
tags: 工具开发

---

我每天都要看妞，没有别的想法，只是为了自己的心情愉悦~~

<!--more-->

<img src="https://tva1.sinaimg.cn/large/008i3skNly1gpulclcazbj30o015cnpd.jpg" style="zoom:30%;" />

---

### 声明

仅限用于技术研究和获得正式授权的测试活动。

### 项目说明

我过去内网渗透中弱口令爆破一般使用超级弱口令(shack2)，但是存在部分问题。一是只支持windows系统+图形化界面，这就意味着要走代理或者 3389 连接上去使用；二是每次爆破都需要选择协议、ip 列表、字典文件，略显麻烦。

然后正好过年的时候看了《白帽子安全开发实战》，发现 go 语言真香，速度快+跨平台编译十分方便，因此花了点时间开发这个项目。

项目大致思路：只需要将端口扫描的结果放入 input.txt，即可启动go-crack，会根据端口对应默认服务，并加载对应服务的爆破字典进行爆破。

### 使用介绍

![](images/008i3skNly1gpuu4bn73tj312u0qona7-20210424150635645.jpg)

- 目前支持的类型（即端口对应的默认服务）

弱口令

```
		21: "ftp",
		22: "ssh",
		135: "wmi",
		161: "snmp",
		445: "smb",
		1433: "mssql",
		//1521: "oracle",
		3306: "mysql",
		//3389: "rdp",
		5432: "postgresql",
		5985: "winrm",
		6379: "redis",
		27017: "mongodb",
```

漏洞（445 端口）

```
		MS17-010
		CVE-2020-0796
```

未授权

```
		9200: "elasticsearch",
		11211: "memcached",
```

- 并发数

启动时我们需要控制的唯一参数就是并发数，`-n`指定即可，不指定的话默认为 10。

- 输入文件

固定输入文件名为：input.txt，每一行的格式为`ip:port`或者`ip:port|porotocol`，后面那种主要是针对修改了默认端口的服务。

135 端口默认对应 wmi 爆破，如果要 hash 爆破的话请指定porotocol 为 wmihash。

445 端口自动会检查 MS-17010和CVE-2020-0796。

- 输出文件

固定输出文件名为：output.txt

- 字典

字典放在/dict 下，根据爆破的服务加载对应的字典，可以自行根据实际情况更新字典。

### 使用演示

执行

```
./go-crack_darwin_amd64 go-crack
./go-crack_darwin_amd64 go-crack -n 15
```

![](images/008i3skNly1gpuu2jwqsxj318c0u0npd.jpg)

可以看到除了弱口令之外，扫描出可能存在 MS-17010，可以进一步确认。

### 后续计划

- [ ] rdp、oracle 协议爆破

- [ ] tomcat、weblogic 等web 弱口令爆破
- [ ] 加入端口扫描+指纹识别（那么只需要输入 IP 即可一键大保健）

### 更新记录

2020.04.20

- 第一版

2020.04.24

- 增加了 wmi爆破和 wmihash 爆破模块

### 参考链接

https://github.com/netxfly/x-crack

https://github.com/k8gege/LadonGo

https://github.com/shadow1ng/fscan

---

喜欢的话给个Star吧，希望你不要不识抬举🐶

