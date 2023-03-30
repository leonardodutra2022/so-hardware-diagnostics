package linux

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"syscall"

	"github.com/leonardodutra9898/tools-display-info/utils"
)

func EnergyReport() {
	msg := "Getting energy information..."
	utils.AddLogText(msg)
	command := exec.Command("tlp-stat")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_ENERGY)
	} else {
		utils.AddLogText("Report Energy with error in your system..")
	}
}

func GetFullInfo() {
	msg := "Getting general information..."
	utils.AddLogText(msg)
	command := exec.Command("dmidecode")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_FULL)
	} else {
		utils.AddLogText("Report Full with error in your system..")
	}
}

func MountsReport() {
	msg := "Getting mounts information..."
	utils.AddLogText(msg)
	command := exec.Command("mount")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_MOUNT)
	} else {
		utils.AddLogText("Report Mounts with error in your system..")
	}
}

func GetACPIInfo() {
	msg := "Getting general information..."
	utils.AddLogText(msg)
	command := exec.Command("acpi", "-Via")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_ACPI)
	} else {
		utils.AddLogText("Report ACPI with error in your system..")
	}
}

func GetSOCommand() {
	msg := "Getting SO information..."
	utils.AddLogText(msg)
	command := exec.Command("uname", "-srio")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_SO)
	} else {
		utils.AddLogText("Report SO with error in your system..")
	}
}

func GetSysInfoCommand() {
	msg := "Getting general system information..."
	utils.AddLogText(msg)
	command := exec.Command("lscpu")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_SYS)
	} else {
		utils.AddLogText("Report Sys Info with error in your system..")
	}
}

func GetSysInfoIICommand() {
	msg := "Getting general system II information..."
	utils.AddLogText(msg)
	command := exec.Command("lshw", "-short")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_SYS_II)
	} else {
		utils.AddLogText("Report Sys Info II (lshw) with error in your system..")
	}

}

func GetSysInfoIIICommand() {
	msg := "Getting general system III information..."
	utils.AddLogText(msg)
	command := exec.Command("hwinfo", "--short")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_SYS_III)
	} else {
		utils.AddLogText("Report Sys Info III with error in your system..")
	}
}

func GetPCICommand() {
	msg := "Getting PCI information..."
	utils.AddLogText(msg)
	command := exec.Command("lspci")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_PCI)
	} else {
		utils.AddLogText("Report PCI with error in your system..")
	}
}

func GetDiskICommand() {
	msg := "Getting Disk information..."
	utils.AddLogText(msg)
	command := exec.Command("lsblk")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_DISC)
	} else {
		utils.AddLogText("Report DISK I with error in your system..")
	}
}

func GetDiskIICommand() {
	msg := "Getting Disk II information..."
	utils.AddLogText(msg)
	command := exec.Command("df", "-H")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_DISC_II)
	} else {
		utils.AddLogText("Report DISK II with error in your system..")
	}
}

func GetDiskIIICommand() {
	msg := "Getting Disk III information..."
	utils.AddLogText(msg)
	command := exec.Command("fdisk", "-l")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_DISC_III)
	} else {
		utils.AddLogText("Report DISK III (fdisk) not allowed as non sudo in your system..")
	}
}

func GetDiskIVCommand() {
	msg := "Getting Disk IV information..."
	utils.AddLogText(msg)
	command := exec.Command("hdparm")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_DISC_IV)
	} else {
		utils.AddLogText("Report DISK IV with error in your system..")
	}
}

func GetPrograms() {
	msg := "Getting information programs..."
	utils.AddLogText(msg)
	command := exec.Command("dpkg", "--get-selections")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_PROGRAMS)
	} else {
		utils.AddLogText("Report PROGRAMS with error in your system..")
	}
	command.Wait()
}

func ListUSBs() {
	msg := "Getting information USBs..."
	utils.AddLogText(msg)
	command := exec.Command("lsusb")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_USB)
	} else {
		utils.AddLogText("Report USB I with error in your system..")
	}
}

func Monitors() {
	msg := "Getting information monitors..."
	utils.AddLogText(msg)
	command := exec.Command("lshw", "-c", "display")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_MONITOR)
	} else {
		utils.AddLogText("Report Monitor (Graphic) I with error in your system..")
	}
}

func Monitors_II() {
	msg := "Getting information graphic..."
	utils.AddLogText(msg)
	command := exec.Command("xrandr")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_MONITOR_II)
	} else {
		utils.AddLogText("Report Monitor II with error in your system..")
	}
}

func MemoryRAM() {
	msg := "Getting memory report..."
	utils.AddLogText(msg)
	command := exec.Command("free", "-mh")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_MEMORY)
	} else {
		utils.AddLogText("Report Memory I with error in your system..")
	}
}

func MemoryRAM_II() {
	msg := "Getting memory II report..."
	utils.AddLogText(msg)
	command := exec.Command("cat", "/proc/meminfo")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_MEMORY_II)
	} else {
		utils.AddLogText("Report Memory II with error in your system..")
	}
}

func MemoryRAM_III() {
	msg := "Getting memory III report..."
	utils.AddLogText(msg)
	command := exec.Command("dmidecode", "-t", "memory")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_MEMORY_III)
	} else {
		utils.AddLogText("Report Memory III with error in your system..")
	}
}

func CPU() {
	msg := "Getting CPU report..."
	utils.AddLogText(msg)
	command := exec.Command("cat", "/proc/cpuinfo")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_CPU)
	} else {
		utils.AddLogText("Report CPU I with error in your system..")
	}
}

func Network_I() {
	msg := "Getting Network I report..."
	utils.AddLogText(msg)
	command := exec.Command("ip", "addr")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_NETWORK_I)
	} else {
		utils.AddLogText("Report Network I with error in your system..")
	}
}

func Network_II() {
	msg := "Getting Network II report..."
	utils.AddLogText(msg)
	command := exec.Command("netstat", "-s")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_NETWORK_II)
	} else {
		utils.AddLogText("Report Network II with error in your system..")
	}
}

func Network_III() {
	msg := "Getting Network III report..."
	utils.AddLogText(msg)
	command := exec.Command("netstat", "-r")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_NETWORK_III)
	} else {
		utils.AddLogText("Report Network III with error in your system..")
	}

}

func Network_IV() {
	msg := "Getting Network IV report..."
	utils.AddLogText(msg)
	command := exec.Command("ss", "-s")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_NETWORK_IV)
	} else {
		utils.AddLogText("Report Network IV with error in your system..")
	}
}

func Network_V() {
	msg := "Getting Network V report..."
	utils.AddLogText(msg)
	command := exec.Command("ifconfig")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_NETWORK_V)
	} else {
		utils.AddLogText("Report Network V with error in your system..")
	}
}

func PermissaoDisplayExtra() {
	msg := "Permission display extra..."
	utils.AddLogText(msg)
	params := string(utils.PATH_SAVE + "displayplacer")
	cmd := string("chmod +x " + "'" + params + "'")
	command := exec.Command("sh", "-c", cmd)
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
}

func DisplayExtraReport() {
	msg := ""
	fmt.Println("")
	exeExtra := string(utils.PATH_SAVE + "displayplacer")
	utils.TIME_5s()
	utils.AddLogText(msg)
	command := exec.Command(exeExtra, "list")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	utils.Execute(output, err, utils.REPORT_DISPLAY_PLACER)
	if err != nil {
		log.Fatal(err)
	}
	command.Wait()
	os.Remove(exeExtra)
}

func CPU_II() {
	msg := "Getting CPU II report..."
	utils.AddLogText(msg)
	command := exec.Command("dmidecode", "-t", "processor")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()

	if err == nil {
		utils.Execute(output, err, utils.REPORT_CPU_II)
	} else {
		utils.AddLogText("Report CPU II with error in your system..")
	}
}

func BIOS() {
	msg := "Getting CPU II report..."
	utils.AddLogText(msg)
	command := exec.Command("dmidecode", "-t", "bios")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()

	if err == nil {
		utils.Execute(output, err, utils.REPORT_BIOS)
	} else {
		utils.AddLogText("Report BIOS with error in your system..")
	}
}

func DeleteFiles() {
	os.Remove(utils.PATH_SAVE + utils.REPORT_CPU)
	os.Remove(utils.PATH_SAVE + utils.REPORT_DISC)
	os.Remove(utils.PATH_SAVE + utils.REPORT_DISC_II)
	os.Remove(utils.PATH_SAVE + utils.REPORT_DISC_III)
	os.Remove(utils.PATH_SAVE + utils.REPORT_DISC_IV)
	os.Remove(utils.PATH_SAVE + utils.REPORT_DISC_V)
	os.Remove(utils.PATH_SAVE + utils.REPORT_DISC_VI)
	os.Remove(utils.PATH_SAVE + utils.REPORT_DISC_VII)
	os.Remove(utils.PATH_SAVE + utils.REPORT_MONITOR)
	os.Remove(utils.PATH_SAVE + utils.REPORT_MONITOR_II)
	os.Remove(utils.PATH_SAVE + utils.REPORT_MONITOR_III)
	os.Remove(utils.PATH_SAVE + utils.REPORT_MEMORY)
	os.Remove(utils.PATH_SAVE + utils.REPORT_MEMORY_II)
	os.Remove(utils.PATH_SAVE + utils.REPORT_MOUNT)
	os.Remove(utils.PATH_SAVE + utils.REPORT_PCI)
	os.Remove(utils.PATH_SAVE + utils.REPORT_PROGRAMS)
	os.Remove(utils.PATH_SAVE + utils.REPORT_SO)
	os.Remove(utils.PATH_SAVE + utils.REPORT_SYS)
	os.Remove(utils.PATH_SAVE + utils.REPORT_SYS_II)
	os.Remove(utils.PATH_SAVE + utils.REPORT_SYS_IV)
	os.Remove(utils.PATH_SAVE + utils.REPORT_USB)
	os.Remove(utils.PATH_SAVE + utils.REPORT_USB_II)
	os.Remove(utils.PATH_SAVE + utils.REPORT_USB_III)
	os.Remove(utils.PATH_SAVE + utils.REPORT_DISPLAY_PLACER)
	os.Remove(utils.PATH_SAVE + utils.REPORT_ENERGY)
	os.Remove(utils.PATH_SAVE + utils.REPORT_SO_KERNEL)
	os.Remove(utils.PATH_SAVE + utils.REPORT_NETWORK_I)
	os.Remove(utils.PATH_SAVE + utils.REPORT_NETWORK_II)
	os.Remove(utils.PATH_SAVE + utils.REPORT_NETWORK_III)
	os.Remove(utils.PATH_SAVE + utils.REPORT_NETWORK_IV)
	os.Remove(utils.PATH_SAVE + utils.REPORT_NETWORK_V)
}
