package layout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/storage"
	"net/url"
	"os"
)

func NewMainMenu(window fyne.Window, app fyne.App) *fyne.MainMenu {
	return fyne.NewMainMenu(newStartMenu(window), newHelpMenu(window, app))
}

func newStartMenu(window fyne.Window) *fyne.Menu {
	return fyne.NewMenu("开始",
		newRegisterItem(window),
		newOpenItem(window),
		newQuitItem(window))
}

func newHelpMenu(window fyne.Window, app fyne.App) *fyne.Menu {
	return fyne.NewMenu("帮助", fyne.NewMenuItem("Github", func() {
		github, err := url.Parse("https://github.com/luckylocode/Moba-Xterm")
		if err != nil {
			return
		}
		app.OpenURL(github)
	}))
}

/**
注册
*/
func newRegisterItem(window fyne.Window) *fyne.MenuItem {
	registerItem := fyne.NewMenuItem("注册", func() {
		validationAndRegister(window)
	})
	registerItem.IsQuit = true
	return registerItem
}

/**
浏览
*/
func newOpenItem(window fyne.Window) *fyne.MenuItem {
	openItem := fyne.NewMenuItem("浏览", func() {

		open := dialog.NewFolderOpen(func(uri fyne.ListableURI, err error) {
			if uri != nil {
				filePath.Text = uri.Path()
			}
			filePath.Refresh()
		}, window)
		// 获取当前路径
		dir, _ := os.Getwd()
		dirUri := storage.NewFileURI(dir)
		openPath, _ := storage.ListerForURI(dirUri)
		open.SetLocation(openPath)
		open.Show()
	})
	return openItem
}

/**
退出
*/
func newQuitItem(window fyne.Window) *fyne.MenuItem {
	quitItem := fyne.NewMenuItem("退出", func() {
		window.Close()
	})
	quitItem.IsQuit = true
	return quitItem
}
