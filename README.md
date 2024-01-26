# Go Miauth Package for Misskey Authentication 

# WIP!!!!!!!!!!!!!!!

## Overview

MiauthのGo言語用パッケージです。

## Usage

### セッションIDを取得する

UUIDの生成にgithub.com/google/uuidを使用しています。
go.modファイルに追加しておいてください。

```go
sessionID := miauth.GenerateSessionID()
```

### MiauthURLの構造を作成する

```go
appName := "MyApp"
callbackURL := "https://myapp.example.pigi/callback"
permission := "write:notes,write:following,read:drive"

miauthURL := miauth.ConstructMiauthURL(sessionID, appName, callbackURL, permission)
```
### 認証を行う

```go
checkURL := fmt.Sprintf("https://{host}/api/miauth/%s/check", sessionID)

accessTokenResponse, err := miauth.PerformMiauthAuthentication(sessionID)
if err != nil {
	fmt.Println("Error:", err)
	return
}

token := accessTokenResponse.Token
user := accessTokenResponse.User
```
### 大まかな実装

 ```go
package main

import (
	"fmt"
	"github.com/CAT5NEKO/miauth"
)

func main() {

	sessionID := miauth.GenerateSessionID()
	appName := "MyApp"
	callbackURL := "https://example.pigi/callback"
	permission := "write:notes,write:following,read:drive"
	miauthURL := miauth.ConstructMiauthURL(sessionID, appName, callbackURL, permission)
	
	accessTokenResponse, err := miauth.PerformMiauthAuthentication(sessionID)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	
	token := accessTokenResponse.Token
	user := accessTokenResponse.User

	fmt.Println("Session ID:", sessionID)
	fmt.Println("Miauth URL:", miauthURL)
	fmt.Println("Access Token:", token)
	fmt.Println("User ID:", user.ID)
	fmt.Println("User Name:", user.Name)
}
```