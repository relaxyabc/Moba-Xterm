package layout

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/data/validation"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/storage"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"image/color"
	"os"
)

var LicenseOption = []string{"专业版", "教育版", "个人版"}

var userName *widget.Entry

var licenseVersion *widget.Select

var targetVersion *widget.Entry

var game *widget.Check

var plugin *widget.Check

var filePath *widget.Entry

func NewBodyLayout(window fyne.Window, app fyne.App) *fyne.Container {
	form := createCenter(window, app)

	return container.NewBorder(makeCell(30, 10), createBottom(), makeCell(30, 30), makeCell(30, 30), form)
}

func createCenter(window fyne.Window, app fyne.App) *widget.Form {
	// 用户名
	userName = widget.NewEntry()
	userName.SetPlaceHolder("username")
	userName.Validator = NewBlankValidation("blank is not allowed")

	// 许可协议
	licenseVersion = widget.NewSelect(LicenseOption, func(license string) {
	})
	licenseVersion.SetSelectedIndex(0)

	// 使用版本
	targetVersion = widget.NewEntry()
	targetVersion.SetPlaceHolder("11.0")
	targetVersion.Validator = validation.NewRegexp("^\\d+\\.\\d+$", "version must be x.x")

	// 保存路径
	filePath = widget.NewEntry()

	// 插件
	game = widget.NewCheck("游戏", func(b bool) {})
	plugin = widget.NewCheck("插件", func(b bool) {})

	box := container.NewHBox(game, container.NewPadded(), plugin)
	buttonBox := operationButtonBox(window)
	form := &widget.Form{
		Items: []*widget.FormItem{
			{Text: "用户名称", Widget: userName},
			{Text: "授权版本", Widget: licenseVersion},
			{Text: "", Widget: makeCell(30, 0.05)},
			{Text: "目标版本", Widget: targetVersion},
			{Text: "保存路径", Widget: createSavePathRow(window)},
		},
	}
	// 保存路径
	//form.Append("保存路径", createSavePathRow(window))
	//form.Append("", makeCell(30, 0.05))
	form.Append("", box)
	form.Append("", buttonBox)
	return form
}

func createSavePathRow(window fyne.Window) *fyne.Container {
	return container.NewBorder(nil, nil, nil,
		widget.NewButtonWithIcon("浏览", theme.FileIcon(), func() {
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
		}),
		filePath)
}

func operationButtonBox(window fyne.Window) *fyne.Container {
	buttonBox := container.NewHBox()
	buttonBox.Add(layout.NewSpacer())
	buttonBox.Add(widget.NewButton("注册", func() {
		validationAndRegister(window)
	}))
	return buttonBox
}

func makeCell(width, height float32) fyne.CanvasObject {
	//rect := canvas.NewRectangle(&color.NRGBA{R: 128, G: 128, B: 128, A: 255})
	//  显示白色
	rect := canvas.NewRectangle(&color.NRGBA{R: 0, G: 0, B: 0, A: 0})
	rect.SetMinSize(fyne.NewSize(width, height))
	return rect
}
