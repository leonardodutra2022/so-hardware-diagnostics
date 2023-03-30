package macos

import (
	"os/exec"
	"syscall"

	"github.com/leonardodutra9898/tools-display-info/utils"
)

func EnergyReport() {
	msg := "Getting energy information..."

	utils.AddLogText(msg)
	command := exec.Command("system_profiler", "SPPowerDataType", "-detailLevel", "full")
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
	command := exec.Command("uname", "-sr")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_SO)
	} else {
		utils.AddLogText("Report SO with error in your system..")
	}
}

func GetSOKernel() {
	msg := "Getting SO Kernel information..."
	utils.AddLogText(msg)
	command := exec.Command("sysctl", "kern.version")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_SO_KERNEL)
	} else {
		utils.AddLogText("Report SO (Kernel) with error in your system..")
	}
}

func GetSysInfoCommand() {
	msg := "Getting general system information..."
	utils.AddLogText(msg)
	command := exec.Command("system_profiler", "SPHardwareDataType", "-detailLevel", "full")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_SYS)
	} else {
		utils.AddLogText("Report System Info with error in your system..")
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
		utils.AddLogText("Report Sys Info II (lshw) not found in your system..")
	}

}

func GetSysInfoIVCommand() {
	msg := "Getting general system IV information..."
	utils.AddLogText(msg)
	command := exec.Command("system_profiler", "SPHardwareDataType", "SPSoftwareDataType", "-detailLevel", "full")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_SYS_IV)
	} else {
		utils.AddLogText("Report SysInfo IV with error in your system..")
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
		utils.AddLogText("Report SysInfo III with error in your system..")
	}
}

func GetPCICommand() {
	msg := "Getting PCI information..."
	utils.AddLogText(msg)
	command := exec.Command("system_profiler", "SPPCIDataType")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_PCI)
	} else {
		utils.AddLogText("Report PCI I with error in your system..")
	}
}

func GetDiskICommand() {
	msg := "Getting Disk information..."
	utils.AddLogText(msg)
	command := exec.Command("system_profiler", "SPStorageDataType", "-detailLevel", "full")
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
		utils.AddLogText("REPORT DISK III - fdisk not allowed as non sudo in your system..")
	}
}

func GetDiskIVCommand() {
	msg := "Getting Disk IV (Serial/ATA) information..."
	utils.AddLogText(msg)
	command := exec.Command("system_profiler", "SPSerialATADataType", "-detailLevel", "full")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_DISC_IV)
	} else {
		utils.AddLogText("Report DISK IV with error in your system..")
	}
}

func GetDiskVCommand() {
	msg := "Getting Disk V (SCSI) information..."
	utils.AddLogText(msg)
	command := exec.Command("system_profiler", "SPParallelSCSIDataType")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_DISC_V)
	} else {
		utils.AddLogText("Report DISK V with error in your system..")
	}
}

func GetDiskVICommand() {
	msg := "Getting Disk VI (NVMe) information..."
	utils.AddLogText(msg)
	command := exec.Command("system_profiler", "SPNVMeDataType", "-detailLevel", "full")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()

	if err == nil {
		utils.Execute(output, err, utils.REPORT_DISC_VI)
	} else {
		utils.AddLogText("Report DISK VI with error in your system..")
	}
}

func GetDiskVIICommand() {
	msg := "Getting Disk VII information..."
	utils.AddLogText(msg)
	command := exec.Command("diskutil", "list")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_DISC_VII)
	} else {
		utils.AddLogText("Report DISK VII with error in your system..")
	}
}

func GetPrograms() {
	msg := "Getting information programs..."
	utils.AddLogText(msg)
	command := exec.Command("system_profiler", "SPApplicationsDataType")
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
	command := exec.Command("system_profiler", "SPUSBDataType")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_USB)
	} else {
		utils.AddLogText("Report USB I with error in your system..")
	}
}

func ListUSBs_II() {
	msg := "Getting information USBs II..."
	utils.AddLogText(msg)
	command := exec.Command("ioreg", "-p", "IOUSB")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_USB_II)
	} else {
		utils.AddLogText("Report USB II with error in your system..")
	}
}

func ListUSBs_III() {
	msg := "Getting information USBs III..."
	utils.AddLogText(msg)
	command := exec.Command("ioreg", "-p", "IOUSB", "-w0", "-l")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_USB_III)
	} else {
		utils.AddLogText("Report USB III with error in your system..")
	}
}

func Monitors() {
	msg := "Getting information monitors..."
	utils.AddLogText(msg)
	command := exec.Command("system_profiler", "SPDisplaysDataType")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_MONITOR)
	} else {
		utils.AddLogText("Report Monitor (Graphic) with error in your system..")
	}
}

func Monitors_II() {
	msg := "Getting information graphic..."
	utils.AddLogText(msg)
	command := exec.Command("ioreg", "-w0", "-c", "IOHIDInterface")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_MONITOR_II)
	} else {
		utils.AddLogText("Report Monitor (Graphic) II with error in your system..")
	}
}

func Monitors_III() {
	msg := "Getting information graphic III..."
	utils.AddLogText(msg)
	command := exec.Command("ioreg", "-lw0", "|", "grep", "EDID")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_MONITOR_III)
	} else {
		utils.AddLogText("Report Monitor (Graphic) III with error in your system..")
	}
}

func Monitors_IV() {
	msg := "Getting more information graphic..."
	utils.AddLogText(msg)
	command := exec.Command("defaults", "read", "/Library/Preferences/com.apple.windowserver.plist")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_MONITOR_IV)
	} else {
		utils.AddLogText("Report Monitor (Graphic) IV with error in your system..")
	}
}

func MemoryRAM() {
	msg := "Getting memory report..."
	utils.AddLogText(msg)
	command := exec.Command("system_profiler", "SPMemoryDataType", "-detailLevel", "full")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_MEMORY)
	} else {
		utils.AddLogText("Report Memory RAM I with error in your system..")
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
		utils.AddLogText("Report Memory RAM II with error in your system..")
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
		utils.AddLogText("Report Memory RAM III with error in your system..")
	}
}

func CPU() {
	msg := "Getting CPU report..."
	utils.AddLogText(msg)
	command := exec.Command("sysctl", "-a")
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()
	if err == nil {
		utils.Execute(output, err, utils.REPORT_CPU)
	} else {
		utils.AddLogText("Report Memory CPU I with error in your system..")
	}
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
		utils.AddLogText("Report Memory CPU II with error in your system..")
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
		utils.AddLogText("Report Memory BIOS with error in your system..")
	}
}

func DeleteLogs() {
	msg := "Deleting log files..."
	utils.AddLogText(msg)
	command := exec.Command("rm", string(utils.PATH_SAVE+"/"), "*.txt")
	command.SysProcAttr = &syscall.SysProcAttr{}

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
