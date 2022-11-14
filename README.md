## golang skills
[Hystrix 熔斷]

[gRPC]

## goal
[kafka]

[ELK]

## MacOS Environment
// 初次使用設定開啟 go.mod
go env -w GO111MODULE=on

// 到該專案根目錄執行 下載使用到的包
go mod tidy

// 安裝 brew
/bin/bash -c "$(curl -fsSL https://raw.githubusercontent.com/Homebrew/install/master/install.sh)"

// 安裝 redis
brew install --cask another-redis-desktop-manager

ruby -e "$(curl -fsSL raw.githubusercontent.com/Homebrew/in…)" < /dev/null 2> /dev/null
brew install caskroom/cask/brew-cask 2> /dev/null

// 允許任何來源
sudo spctl --master-disable
sudo spctl --master-enable

[Vscode]
// goformat
/usr/local/go/src/go/format/format.go
tabWidth    = 4
printerMode = printer.UseSpaces

cd /usr/local/go/bin
go install golang.org/x/tools/gopls@latest

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

[GitHub]
// 安裝 git 更新認證
brew tap microsoft/git
brew install --cask git-credential-manager-core
brew upgrade git-credential-manager-core

[Mysql]
brew install mysql
brew services restart mysql

// MySQL 5.7使用的默認爲 utf8mb4_unicode_ci，但是從MySQL8.0開始使用的已經改成 utf8mb4_0900_ai_ci
utf8mb4

[Redis]
// Homebrew 安裝的軟件會默認在 /usr/local/Cellar/
// redis 的配置文件 /usr/local/etc/redis.conf
brew install redis
brew services start redis

[Docker]
// 背景執行
docker-compose up -d

[Heroku]
heroku login

[Tools]
// DB
Navicat Premium
// Redis
Another Desktop Manager
// 截圖
Snipaste
// WS Test
http://www.websocket-test.com/

[Reading]
gin | Light Weight MVC Framework | https://github.com/skyhee/gin-doc-cn
gorm | ORM Framework  | https://github.com/jinzhu/gorm
redis | redis緩存 | https://github.com/go-redis/redis
grpc | grpc微服務 | https://grpc.io
log | 高性能日誌 | https://github.com/uber-go/zap
elasticsearch | 分佈式搜索引擎 | https://www.elastic.co/cn/products/elasticsearch
