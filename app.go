package main

import (
	"context"
	"fmt"

	"github.com/leonardodutra9898/tools-display-info/commands/linux"
	"github.com/leonardodutra9898/tools-display-info/commands/macos"
	"github.com/leonardodutra9898/tools-display-info/commands/windows"
	"github.com/leonardodutra9898/tools-display-info/utils"
	"github.com/wailsapp/wails/v2/pkg/runtime"
)

func extractWin() {
	utils.ExtractUSBApp(&assets)
	utils.ExtractDriverView(&assets)
	utils.ExtractMonitorInfoView(&assets)
	utils.ExtractMonitorInfoIIView(&assets)
	utils.ExtractEDID(&assets)
}

func extractMac() {
	utils.ExtractDisplayPlacer(&assets)
}

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(pathFile string) string {
	utils.SetPath(pathFile)

	if utils.WhatsOS() == "windows" {
		extractWin()
		go windows.EnergyReport()
		go windows.Monitors()
		go windows.Domain()
		go windows.MemoryRAM()
		go windows.GetExtraReport()
		go windows.GetSysInfoCommand()
		go windows.BIOS()
		windows.ExtraReportUSBWithEnergy()
		windows.ExtraReportDriver()
		windows.ExtraReportMonitor()
		windows.ExtraReportMonitorII()
		windows.ExtraReportEDID()
		windows.GetReportWindowsCommand()
		windows.GetPrograms()
		utils.TIME_5s()
		utils.Zip()
		windows.DeleteFiles()
	} else if utils.WhatsOS() == "linux" {
		linux.EnergyReport()
		linux.GetFullInfo()
		linux.MountsReport()
		linux.GetACPIInfo()
		linux.GetSOCommand()
		linux.ListUSBs()
		linux.GetSysInfoCommand()
		linux.GetSysInfoIICommand()
		linux.GetSysInfoIIICommand()
		linux.GetDiskICommand()
		linux.GetDiskIICommand()
		linux.GetDiskIIICommand()
		linux.GetDiskIVCommand()
		linux.GetPCICommand()
		linux.GetPrograms()
		linux.Monitors()
		linux.Monitors_II()
		linux.MemoryRAM()
		linux.MemoryRAM_II()
		linux.MemoryRAM_III()
		linux.CPU()
		linux.CPU_II()
		linux.BIOS()
		linux.Network_I()
		linux.Network_IV()
		linux.Network_III()
		linux.Network_II()
		linux.Network_V()

		utils.ZipFolder(utils.PATH_SAVE)
		linux.DeleteFiles()
	} else {
		extractMac()
		macos.EnergyReport()
		macos.MountsReport()
		macos.GetSOCommand()
		macos.GetSOKernel()
		macos.ListUSBs()
		macos.ListUSBs_II()
		macos.ListUSBs_III()
		macos.GetSysInfoCommand()
		macos.GetSysInfoIICommand()
		macos.GetSysInfoIVCommand()
		macos.GetDiskICommand()
		macos.GetDiskIICommand()
		macos.GetDiskIIICommand()
		macos.GetDiskIVCommand()
		macos.GetDiskVCommand()
		macos.GetDiskVICommand()
		macos.GetDiskVIICommand()
		macos.GetPCICommand()
		macos.GetPrograms()
		macos.Monitors()
		macos.Monitors_II()
		macos.Monitors_III()
		macos.MemoryRAM()
		macos.CPU()
		linux.PermissaoDisplayExtra()
		utils.TIME_5s()
		linux.DisplayExtraReport()
		macos.Network_I()
		macos.Network_II()
		macos.Network_III()
		linux.Network_V()
		linux.Network_IV()
		utils.ZipFolder(utils.PATH_SAVE)
		linux.DeleteFiles()
	}

	return fmt.Sprintf("Status: %s", "Process Complete.")
}

func (a *App) SelectFolder() string {
	result, _ := runtime.OpenDirectoryDialog(a.ctx, runtime.OpenDialogOptions{
		Title: "Select Dir to save reports...",
	})
	return result
}

func (a *App) ArchApp() string {
	if utils.WhatsOS() == "windows" {
		return "windows"
	} else if utils.WhatsOS() == "linux" {
		return "linux"
	}
	return "macos"
}
