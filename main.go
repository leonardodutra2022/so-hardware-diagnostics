package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed embeded/USBDeview.exe
//go:embed embeded/displayplacer
//go:embed embeded/DriverView.exe
//go:embed embeded/MonitorInfoView.exe
//go:embed embeded/MultiMonitorTool.exe
//go:embed embeded/DumpEDID.exe
//go:embed all:frontend/dist
var assets embed.FS

//go:embed build/appicon.png
var icon []byte

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// Create application with options
	err := wails.Run(&options.App{
		Title:            "Diagnostics",
		Width:            550,
		Height:           700,
		Assets:           assets,
		MinWidth:         540,
		MinHeight:        680,
		BackgroundColour: &options.RGBA{R: 62, G: 62, B: 116, A: 1},
		DisableResize:    true,
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		Linux: &linux.Options{
			Icon:                icon,
			WindowIsTranslucent: false,
		},
		Windows: &windows.Options{
			DisableWindowIcon: true,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
