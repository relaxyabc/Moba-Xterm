package layout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"net/url"
)

func createBottom() *fyne.Container {
	box := container.NewVBox()
	box.Add(widget.NewLabel("使用说明:"))
	box.Add(widget.NewLabel("1:输入用户名,版本,选择 Custom.mxpro 存放路径,点击注册"))
	box.Add(widget.NewLabel("2:手动复制 Custom.mxpro 到 mobaxterm 安装目录即可;"))
	box.Add(widget.NewLabel("  如 D:/Program Files (x86)/Mobatek/MobaXterm"))

	mobaxterm, err := url.Parse("https://mobaxterm.mobatek.net/")
	if err != nil {
		return box
	}

	hBox := container.NewHBox()
	hBox.Add(widget.NewLabel("3:官方网址:"))
	hBox.Add(widget.NewHyperlink("https://mobaxterm.mobatek.net/", mobaxterm))
	box.Add(hBox)
	return box
}
