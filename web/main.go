package main

import (
	"fmt"
	"time"
	"web/utils"
)

func main() {
	userToken, _ := utils.GenToken("user", utils.FrontUserExpireDuration, utils.FrontUserSecret)
	fmt.Println(userToken)
	token, _ := utils.AuthToken(userToken, utils.FrontUserSecret)
	fmt.Println(token.UserName)
	// unix-->标准时间
	fmt.Println(time.Unix(token.ExpiresAt, 0).Format("2006-01-02 15:04:05"))
}
