package main

import (
	"log"
	"os"
)

//把UI的地址独立出来,方便开发的时候能够指定为开发地址,方便实时刷新,默认使用内嵌资源域名(app)
var uiAddress = "http://app/"

func init() {
	//如果环境变量中指定了ui地址,则使用环境变量中的
	if address, exist := os.LookupEnv("BLINK_UI"); exist {
		log.Println("使用指定的UI地址:", address)
		uiAddress = address
	}
}
