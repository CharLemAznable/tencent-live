### tencent-live

[![Build Status](https://travis-ci.org/CharLemAznable/tencent-live.svg?branch=master)](https://travis-ci.org/CharLemAznable/tencent-live)
![GitHub release (latest by date)](https://img.shields.io/github/v/release/CharLemAznable/tencent-live)
[![MIT Licence](https://badges.frapsoft.com/os/mit/mit.svg?v=103)](https://opensource.org/licenses/mit-license.php)
![GitHub code size](https://img.shields.io/github/languages/code-size/CharLemAznable/tencent-live)

腾讯云直播演示工程.

#### 配置文件

1. ```appConfig.toml```

```toml
Port = 20514
ContextPath = ""

PushServer = ""     # 推流地址
PushSecretKey = ""  # 推流密钥
LiveServer = ""     # 直播地址
LiveSecretKey = ""  # 直播密钥

StreamName = ""     # 流名称
Timeout = 3600      # 流过期时间(单位: 秒)
```

2. ```logback.xml```

```xml
<logging>
    <filter enabled="true">
        <tag>file</tag>
        <type>file</type>
        <level>INFO</level>
        <property name="filename">tencent-live.log</property>
        <property name="format">[%D %T] [%L] (%S) %M</property>
        <property name="rotate">false</property>
        <property name="maxsize">0M</property>
        <property name="maxlines">0K</property>
        <property name="daily">false</property>
    </filter>
</logging>
```

#### 部署执行

1. 下载最新的可执行文件压缩包并解压

    下载地址: [tencent-live release](https://github.com/CharLemAznable/tencent-live/releases)

```bash
$ tar -xvJf tencent-live-[version].[arch].[os].tar.xz
```

2. 新建/编辑配置文件, 启动运行

```bash
$ nohup ./tencent-live-[version].[arch].[os].bin &
```
