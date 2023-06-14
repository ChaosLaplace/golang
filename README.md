## golang skills
1.[Hystrix 熔斷]

2.[gRPC]

3.[kafka]

## goal
[ELK]

## MacOS Environment
### 初次使用設定開啟 go.mod
```sh
go env -w GO111MODULE=on
```

### 到該專案根目錄執行 下載使用到的包
```sh
go mod tidy
```

### 安裝 brew
```sh
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"
```

### 安裝 redis
```sh
brew install --cask another-redis-desktop-manager
ruby -e "$(curl -fsSL raw.githubusercontent.com/Homebrew/in…)" < /dev/null 2> /dev/null
brew install caskroom/cask/brew-cask 2> /dev/null
```

### 允許任何來源
```sh
sudo spctl --master-disable
sudo spctl --master-enable
```

## Vscode
### goformat
```sh
/usr/local/go/src/go/format/format.go
tabWidth    = 4
printerMode = printer.UseSpaces

cd /usr/local/go/bin
go install golang.org/x/tools/gopls@latest
```

```go
"[go]": {
    "editor.insertSpaces": true,
    "editor.snippetSuggestions": "none",
    "editor.formatOnSave": true,
    "editor.codeActionsOnSave": {
        "source.organizeImports": true
    }
},
"editor.renderControlCharacters": true,
"editor.renderWhitespace": "all",
"go.formatTool": "goformat",
```

## GitHub
### 安裝 git 更新認證
```sh
brew tap microsoft/git
brew install --cask git-credential-manager-core
brew upgrade git-credential-manager-core
```

## Mysql
```sh
// MySQL 5.7使用的默認爲 utf8mb4_unicode_ci，但是從MySQL8.0開始使用的已經改成 utf8mb4_0900_ai_ci
utf8mb4
brew install mysql
brew services restart mysql
```

## Redis
```sh
// Homebrew 安裝的軟件會默認在 /usr/local/Cellar/
// redis 的配置文件 /usr/local/etc/redis.conf
brew install redis
brew services start redis
```
## Kafka
### 安裝
```sh
brew install kafka

安裝kafka是需要依賴於zookeeper的，所以安裝kafka的時候也會包含zookeeper 
kafka的安裝目錄：/usr/local/Cellar/kafka 
kafka的配置文件目錄：/usr/local/etc/kafka 
kafka服務的配置文件：/usr/local/etc/kafka/server.properties 
zookeeper配置文件： /usr/local/etc/kafka/zookeeper.properties

# server.properties
broker.id=0
listeners=PLAINTEXT://:9092
advertised.listeners=PLAINTEXT://127.0.0.1:9092
log.dirs=/usr/local/var/lib/kafka-logs

# zookeeper.properties
dataDir=/usr/local/var/lib/zookeeper
clientPort=2181
maxClientCnxns=0
```

### 啟動 Zookeeper & Kafka
```sh
# # 啟動 Kafka 之前先啟動 Zookeeper
cd /usr/local/Cellar/kafka/3.3.1_1
./bin/zookeeper-server-start /usr/local/etc/kafka/zookeeper.properties
./bin/kafka-server-start /usr/local/etc/kafka/server.properties
```

## Docker
### 背景執行
```sh
docker-compose up -d
```

## Heroku
```sh
heroku login
```

## Tools
### DB
Navicat Premium

### Redis
Another Desktop Manager

### 截圖
Snipaste

### WS Test
http://www.websocket-test.com/

## Reading
### gin
Light Weight MVC Framework | https://github.com/skyhee/gin-doc-cn

### gorm
ORM Framework  | https://github.com/jinzhu/gorm

### redis
redis緩存 | https://github.com/go-redis/redis

### grpc
grpc微服務 | https://grpc.io

### log
高性能日誌 | https://github.com/uber-go/zap

### elasticsearch
分佈式搜索引擎 | https://www.elastic.co/cn/products/elasticsearch


## Jenkins
```sh
更新 apt
sudo apt update -y

安裝 java 11
sudo apt install openjdk-11-jdk -y

安裝 Jenkins
密鑰添加到系統
curl -fsSL https://pkg.jenkins.io/debian-stable/jenkins.io-2023.key | sudo tee \
    /usr/share/keyrings/jenkins-keyring.asc > /dev/null

加一個 Jenkins apt 存儲庫條目
echo deb [signed-by=/usr/share/keyrings/jenkins-keyring.asc] \
    https://pkg.jenkins.io/debian-stable binary/ | sudo tee \
    /etc/apt/sources.list.d/jenkins.list > /dev/null

sudo apt-get update
sudo apt-get install fontconfig openjdk-11-jre
sudo apt-get install jenkins

修改預設 Port 8880
sudo vim /usr/lib/systemd/system/jenkins.service
sudo ufw allow 8880
sudo systemctl daemon-reload
sudo service jenkins restart

首次訪問
sudo cat /var/lib/jenkins/secrets/initialAdminPassword
```
