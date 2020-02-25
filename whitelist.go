package main

import (
	"net"
	"strconv"
	"whitelist/ipset"

	"github.com/gin-gonic/gin"
)

const (
	ipsetName   = "scp-whitelist"
	ipTimeout   = 86400
	protectPort = 7777
	webPort     = 2333
)

func main() {
	//ipset
	whitelist, err := ipset.New(ipsetName, "hash:ip", &ipset.Params{})
	if err != nil {
		panic(err)
	}

	//web
	r := gin.Default()
	r.GET("/whitelist", func(c *gin.Context) {
		ip := c.ClientIP()
		if checkIP(ip) {
			err := whitelist.Add(ip, ipTimeout)
			if err == nil {
				c.String(200, "您的IP %s 成功加入白名单，可以开始玩耍啦~", ip)
				return
			}
		}

		c.String(403, "看起来出现了问题，要不咱们重新打开试试？")
	})
	r.Run(":" + strconv.Itoa(webPort))
}

func checkIP(ip string) bool {
	address := net.ParseIP(ip)
	if address == nil {
		return false
	}
	return true
}
