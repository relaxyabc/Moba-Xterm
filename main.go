package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"github.com/flopp/go-findfont"
	"mobaxterm/asset"
	"mobaxterm/layout"
	"os"
	"strings"
)

func init() {
	fontPaths := findfont.List()
	for _, path := range fontPaths {
		// 楷体:simkai.ttf, 黑体:simhei.ttf
		if strings.Contains(path, "simkai.ttf") {
			// 设置环境变量 ; 取消环境变量 os.Unsetenv("FYNE_FONT")
			os.Setenv("FYNE_FONT", path)
			break
		}
	}
}

func main() {
	mobaxterm := app.New()
	// 图标
	mobaxterm.SetIcon(asset.GetLogoResource())
	window := mobaxterm.NewWindow("MobaXterm")

	// 创建菜单
	window.SetMainMenu(layout.NewMainMenu(window, mobaxterm))
	// body 布局
	body := layout.NewBodyLayout(window, mobaxterm)
	window.SetContent(body)
	// 禁止调整大小
	window.SetFixedSize(true)
	window.Resize(fyne.NewSize(630, 520))
	window.CenterOnScreen()
	window.ShowAndRun()
}
