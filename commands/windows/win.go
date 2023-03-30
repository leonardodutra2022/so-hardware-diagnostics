package windows

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/leonardodutra9898/tools-display-info/utils"
)

func GetReportWindowsCommand() {
	msg := "Getting full report..."

	utils.AddLogText(msg)
	command := exec.Command("cmd", "/C", "msinfo32", "/report", utils.PATH_SAVE+utils.REPORT_FULL)
	command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := command.Start()
	if err != nil {
		log.Fatal(err)
	}
	command.Wait()
}

func ExtraReportUSBWithEnergy() {
	exeUSB := string(utils.PATH_SAVE + "USBDeview.exe")
	utils.TIME_5s()

	command := exec.Command(exeUSB, "/shtml", utils.PATH_SAVE+utils.REPORT_EXTRA_USB)
	command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := command.Start()
	if err != nil {
		log.Fatal(err)
	}
	command.Wait()
	os.Remove(exeUSB)
}

func ExtraReportBatery() {
	exeExtra := string(utils.PATH_SAVE + "BatteryInfoView.exe")
	utils.TIME_5s()

	command := exec.Command(exeExtra, "/shtml", utils.PATH_SAVE+utils.REPORT_EXTRA_BATERY)
	command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := command.Start()
	if err != nil {
		log.Fatal(err)
	}
	command.Wait()
	os.Remove(exeExtra)
}

func ExtraReportDriver() {
	exeExtra := string(utils.PATH_SAVE + "DriverView.exe")
	utils.TIME_5s()

	command := exec.Command(exeExtra, "/shtml", utils.PATH_SAVE+utils.REPORT_EXTRA_DRIVER)
	command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := command.Start()
	if err != nil {
		log.Fatal(err)
	}
	command.Wait()
	os.Remove(exeExtra)
}

func ExtraReportMonitor() {
	exeExtra := string(utils.PATH_SAVE + "MonitorInfoView.exe")
	utils.TIME_5s()

	command := exec.Command(exeExtra, "/shtml", utils.PATH_SAVE+utils.REPORT_EXTRA_MONITOR)
	command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := command.Start()
	if err != nil {
		log.Fatal(err)
	}
	command.Wait()
	os.Remove(exeExtra)
}

func ExtraReportMonitorII() {
	exeExtra := string(utils.PATH_SAVE + "MultiMonitorTool.exe")
	utils.TIME_5s()

	command := exec.Command(exeExtra, "/shtml", utils.PATH_SAVE+utils.REPORT_EXTRA_MONITOR_II)
	command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := command.Start()
	if err != nil {
		log.Fatal(err)
	}
	command.Wait()
	os.Remove(exeExtra)
}

func ExtraReportEDID() {
	msg := ""
	fmt.Println("")
	exeExtra := string(utils.PATH_SAVE + "DumpEDID.exe")
	utils.TIME_5s()
	utils.AddLogText(msg)
	command := exec.Command(exeExtra)
	command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := command.Output()
	utils.Execute(output, err, utils.REPORT_EXTRA_EDID)
	if err != nil {
		log.Fatal(err)
	}
	command.Wait()
	os.Remove(exeExtra)
}

func EnergyReport() {
	msg := "Getting energy report..."

	utils.AddLogText(msg)
	command := exec.Command("powercfg", "ENERGY", "DURATION", "30", "OUTPUT", utils.PATH_SAVE+utils.REPORT_ENERGY_WINDOWS)
	command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := command.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func GetExtraReport() {
	msg := "Getting extra report..."

	utils.AddLogText(msg)
	command := exec.Command("dxdiag", "/dontskip", "/whql:off", "/64bit", "/t", utils.PATH_SAVE+utils.REPORT_EXTRA)
	command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	err := command.Start()
	if err != nil {
		log.Fatal(err)
	}
}

func GetSysInfoCommand() {
	msg := "Getting general system information..."

	utils.AddLogText(msg)
	command := exec.Command("systeminfo")
	command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := command.Output()
	utils.Execute(output, err, utils.REPORT_SYS)
}

func GetPrograms() {
	msg := "Getting information from system programs..."

	utils.AddLogText(msg)
	command := exec.Command("wmic", "product")
	command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := command.Output()
	utils.Execute(output, err, utils.REPORT_PROGRAMS)
	command.Wait()
}

func Monitors() {
	msg := "Getting information monitors..."

	utils.AddLogText(msg)
	command := exec.Command("wmic", "desktopmonitor")
	command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := command.Output()
	utils.Execute(output, err, utils.REPORT_MONITOR)
}

func Domain() {
	msg := "Getting information domain..."

	utils.AddLogText(msg)
	command := exec.Command("wmic", "ntdomain")
	command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := command.Output()
	utils.Execute(output, err, utils.REPORT_DOMAIN)
}

func MemoryRAM() {
	msg := "Getting memory report..."

	utils.AddLogText(msg)
	command := exec.Command("wmic", "MEMORYCHIP")
	command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := command.Output()
	utils.Execute(output, err, utils.REPORT_MEMORY)
}

func BIOS() {
	msg := "Getting bios report..."

	utils.AddLogText(msg)
	command := exec.Command("wmic", "bios", "get", "Version,BIOSVersion,CurrentLanguage,Manufacturer,ReleaseDate,SerialNumber,Status")
	command.SysProcAttr = &syscall.SysProcAttr{HideWindow: true}
	output, err := command.Output()
	utils.Execute(output, err, utils.REPORT_BIOS)
}

func DeleteFiles() {
	os.Remove(utils.PATH_SAVE + utils.REPORT_EXTRA)
	os.Remove(utils.PATH_SAVE + utils.REPORT_DOMAIN)
	os.Remove(utils.PATH_SAVE + utils.REPORT_MONITOR)
	os.Remove(utils.PATH_SAVE + utils.REPORT_MEMORY)
	os.Remove(utils.PATH_SAVE + utils.REPORT_PROGRAMS)
	os.Remove(utils.PATH_SAVE + utils.REPORT_SYS)
	os.Remove(utils.PATH_SAVE + utils.REPORT_FULL)
	os.Remove(utils.PATH_SAVE + utils.REPORT_EXTRA_DRIVER)
	os.Remove(utils.PATH_SAVE + utils.REPORT_EXTRA_MONITOR)
	os.Remove(utils.PATH_SAVE + utils.REPORT_EXTRA_MONITOR_II)
	os.Remove(utils.PATH_SAVE + utils.REPORT_EXTRA_USB)
	os.Remove(utils.PATH_SAVE + utils.REPORT_EXTRA_EDID)
	existReportEnergyExtra, _ := utils.Exists(utils.PATH_SAVE + utils.REPORT_EXTRA_BATERY)
	if existReportEnergyExtra {
		os.Remove(utils.PATH_SAVE + utils.REPORT_EXTRA_BATERY)
	}
	os.Remove(utils.PATH_SAVE + utils.REPORT_BIOS)
	existReportEnergy, _ := utils.Exists(utils.PATH_SAVE + utils.REPORT_ENERGY_WINDOWS)
	if existReportEnergy {
		os.Remove(utils.PATH_SAVE + utils.REPORT_ENERGY_WINDOWS)
	}
}
