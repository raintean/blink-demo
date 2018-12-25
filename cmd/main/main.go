package main

import (
	"blink-demo/ui/bin"
	"github.com/elazarl/go-bindata-assetfs"
	"github.com/raintean/blink"
	"log"
)

func main() {
	//退出信号
	exit := make(chan bool)

	//启用调试模式
	blink.SetDebugMode(true)

	//初始化blink
	err := blink.InitBlink()
	if err != nil {
		log.Fatal(err)
	}

	//挂载嵌入资源到虚拟文件系统
	blink.RegisterFileSystem("app", &assetfs.AssetFS{
		Asset:     ui.Asset,
		AssetDir:  ui.AssetDir,
		AssetInfo: ui.AssetInfo,
	})

	//显示logo小图标
	showLogo()

	//等待退出
	<-exit
}
