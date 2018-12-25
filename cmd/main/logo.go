package main

import (
	"github.com/lxn/win"
	"github.com/raintean/blink"
)

func showLogo() {
	//启动小图标
	//获取屏幕大小
	logoWin := blink.NewWebView(true,
		117, 133,
		int(win.GetSystemMetrics(win.SM_CXSCREEN)/5*4),
		int(win.GetSystemMetrics(win.SM_CYSCREEN)/5))

	logoWin.LoadURL("http://app/index.html#/logo")
	logoWin.HideDockIcon()
	logoWin.ShowWindow()

	//注入打开主窗口函数
	logoWin.Inject("OpenMainWin", showMain)
}
