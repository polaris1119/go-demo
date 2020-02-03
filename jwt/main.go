// Copyright 2020 The StudyGolang Authors. All rights reserved.
// Use of self source code is governed by a MIT
// license that can be found in the LICENSE file.
// https://studygolang.com
// Author: polaris	polaris@studygolang.com

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/dgrijalva/jwt-go"
)


var htmlTPL = `
<div>
<form action="/" method="post">
<input type="text" name="username" placeholder="请输入用户名">
<input type="password" name="passwd" placeholder="请输入用密码">
<input type="submit" value="登录">
</form>
</div>
`

var hmacSecret = []byte(`uU74PHZLaY0m`)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) {
		if req.Method != "POST" {
			cookie, err := req.Cookie("jwt")
			if err == http.ErrNoCookie {
				fmt.Fprint(w, htmlTPL)
			} else {
				userData, err := parseAndvalidateToken(cookie.Value)
				if err != nil {
					fmt.Fprint(w, `<font color="red">用户解析异常，请重新登录！</font>` + htmlTPL)
					return
				}

				hadLoginAndWelcome(w, userData["username"], userData["login-time"])
			}
			return
		}

		username := req.FormValue("username")
		passwd := req.FormValue("passwd")
		if username != "polaris" || passwd != "studygolang" {
			fmt.Fprint(w, `<font color="red">用户名或密码错误</font>` + htmlTPL)
			return
		}

		curTime := time.Now().Format("2006-01-02 15:04:05")
		jwtToken, err := genToken(username, curTime)
		if err != nil {
			fmt.Println("生成 token 出错：", err)
			fmt.Fprint(w, `<font color="red">服务端错误</font>` + htmlTPL)
			return
		}

		cookie := &http.Cookie{
			Name: "jwt",
			Value: jwtToken,
			MaxAge: 86400,
		}
		http.SetCookie(w, cookie)

		hadLoginAndWelcome(w, username, curTime)
	})

	http.ListenAndServe(":2020", nil)
}

func hadLoginAndWelcome(w http.ResponseWriter, username, loginTime string) {
	fmt.Fprintf(w, `<div><h2>欢迎你，%s，登录时间：%s</h2></div>`, username, loginTime)
}


func genToken(username, curTime string) (string, error) {
	// Claims 用于存放数据
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username": username,
		"login-time": curTime,
	})

	return token.SignedString(hmacSecret)
}

func parseAndvalidateToken(jwtToken string) (map[string]string, error) {
	// 解析 token
	token, err := jwt.Parse(jwtToken, func(token *jwt.Token) (interface{}, error) {
		// 进行 alg 即签名算法校验
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return hmacSecret, nil
	})

	// 校验有效性，并获取 Claims 中的值
	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		username := claims["username"].(string)
		loginTime := claims["login-time"].(string)

		return map[string]string{
			"username": username,
			"login-time": loginTime,
		}, nil
	}

	return nil, err
}