package utils

import (
	"embed"
	"fmt"
	"os"
)

func ExtractUSBApp(assets *embed.FS) {
	exe0, err0 := assets.ReadFile("embeded/USBDeview.exe")
	exeUSB := string(PATH_SAVE + "USBDeview.exe")
	fileExists, _ := Exists(exeUSB)
	if !fileExists {
		os.Remove(exeUSB)
	}
	writeFile(exeUSB, exe0)
	fmt.Println(err0)
}

func writeFile(fileName string, bin []byte) {
	out1, _ := os.Create(fileName)
	out1.Write(bin)
	defer out1.Close()
}

func ExtractBateryInfo(assets *embed.FS) {
	exe2, err0 := assets.ReadFile("embeded/BatteryInfoView.exe")
	exeExtra := string(PATH_SAVE + "BatteryInfoView.exe")
	fileExists, _ := Exists(exeExtra)
	if !fileExists {
		os.Remove(exeExtra)
	}
	writeFile(exeExtra, exe2)
	fmt.Println(err0)
}

func ExtractDriverView(assets *embed.FS) {
	exe3, err0 := assets.ReadFile("embeded/DriverView.exe")
	exeExtra := string(PATH_SAVE + "DriverView.exe")
	fileExists, _ := Exists(exeExtra)
	if !fileExists {
		os.Remove(exeExtra)
	}
	writeFile(exeExtra, exe3)
	fmt.Println(err0)
}

func ExtractMonitorInfoView(assets *embed.FS) {
	exe4, err0 := assets.ReadFile("embeded/MonitorInfoView.exe")
	exeExtra := string(PATH_SAVE + "MonitorInfoView.exe")
	fileExists, _ := Exists(exeExtra)
	if !fileExists {
		os.Remove(exeExtra)
	}
	writeFile(exeExtra, exe4)
	fmt.Println(err0)
}

func ExtractMonitorInfoIIView(assets *embed.FS) {
	exe1, err0 := assets.ReadFile("embeded/MultiMonitorTool.exe")
	exeExtra := string(PATH_SAVE + "MultiMonitorTool.exe")
	fileExists, _ := Exists(exeExtra)
	if !fileExists {
		os.Remove(exeExtra)
	}
	writeFile(exeExtra, exe1)
	fmt.Println(err0)
}

func ExtractEDID(assets *embed.FS) {
	exe1, err0 := assets.ReadFile("embeded/DumpEDID.exe")
	exeExtra := string(PATH_SAVE + "DumpEDID.exe")
	fileExists, _ := Exists(exeExtra)
	if !fileExists {
		os.Remove(exeExtra)
	}
	writeFile(exeExtra, exe1)
	fmt.Println(err0)
}

func ExtractDisplayPlacer(assets *embed.FS) {
	exe1, err0 := assets.ReadFile("embeded/displayplacer")
	exeExtra := string(PATH_SAVE + "displayplacer")
	fileExists, _ := Exists(exeExtra)
	if !fileExists {
		os.Remove(exeExtra)
	}
	writeFile(exeExtra, exe1)
	fmt.Println(err0)
}
