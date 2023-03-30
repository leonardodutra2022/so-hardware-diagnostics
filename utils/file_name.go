package utils

import (
	"fmt"
	"log"
	"os/exec"
	"strings"
	"syscall"
)

var REPORT_MONITOR string = "report_graphic.txt"
var REPORT_MONITOR_II string = "report_monitors.txt"
var REPORT_MONITOR_III string = "report_monitors_III.txt"
var REPORT_MONITOR_IV string = "report_monitors_IV.txt"
var REPORT_PROGRAMS string = "report_programs.txt"
var REPORT_SYS string = "report_sysinfo.txt"
var REPORT_SYS_II string = "report_sysinfo_II.txt"
var REPORT_SYS_III string = "report_sysinfo_III.txt"
var REPORT_SYS_IV string = "report_sysinfo_IV.txt"
var REPORT_FULL string = "report_full.txt"
var REPORT_DOMAIN string = "report_domain.txt"
var REPORT_ENERGY string = "report_energy.txt"
var REPORT_ENERGY_WINDOWS string = "report_energy.html"
var REPORT_MEMORY string = "report_memory.txt"
var REPORT_MEMORY_II string = "report_memory_II.txt"
var REPORT_MEMORY_III string = "report_memory_III.txt"
var REPORT_USB string = "report_usb.txt"
var REPORT_USB_II string = "report_usb_II.txt"
var REPORT_USB_III string = "report_usb_III.txt"
var REPORT_PCI string = "report_pci.txt"
var REPORT_DISC string = "report_disc.txt"
var REPORT_DISC_II string = "report_disc_II.txt"
var REPORT_DISC_III string = "report_disc_III.txt"
var REPORT_DISC_IV string = "report_disc_IV.txt"
var REPORT_DISC_V string = "report_disc_V.txt"
var REPORT_DISC_VI string = "report_disc_VI.txt"
var REPORT_DISC_VII string = "report_disc_VII.txt"
var REPORT_CPU string = "report_cpu.txt"
var REPORT_CPU_II string = "report_cpu_II.txt"
var REPORT_BIOS string = "report_bios.txt"
var REPORT_SO string = "report_so.txt"
var REPORT_SO_KERNEL string = "report_so_kernel.txt"
var REPORT_ACPI string = "report_acpi.txt"
var REPORT_MOUNT string = "report_mount.txt"
var REPORT_EXTRA string = "report_more.txt"
var REPORT_EXTRA_USB string = "report_extra_usb.html"
var REPORT_EXTRA_DEVMAN string = "report_extra_devman.html"
var REPORT_EXTRA_BATERY string = "report_extra_batery.html"
var REPORT_EXTRA_DRIVER string = "report_extra_driver.html"
var REPORT_EXTRA_MONITOR string = "report_extra_monitor.html"
var REPORT_EXTRA_MONITOR_II string = "report_extra_monitorII.html"
var REPORT_EXTRA_EDID string = "report_extra_edid.txt"
var REPORT_DISPLAY_PLACER string = "report_extra_displayp.txt"
var REPORT_NETWORK_I string = "report_network_I.txt"
var REPORT_NETWORK_II string = "report_network_II.txt"
var REPORT_NETWORK_III string = "report_network_III.txt"
var REPORT_NETWORK_IV string = "report_network_IV.txt"
var REPORT_NETWORK_V string = "report_network_V.txt"
var LOGS string = "logs.log"
var PATH_SAVE string
var LogsTexts []string = []string{}

func AddLogText(log string) {
	LogsTexts = append(LogsTexts, log)
	LogsTexts = append(LogsTexts, "\n")
}

func SaveLogs() {
	SaveInDiskStrings(LogsTexts, PATH_SAVE+LOGS)
}

func SetPath(pathFile string) {
	if WhatsOS() == "windows" {
		if pathFile == "" {
			PATH_SAVE = "C:\\"
		}
		if !strings.HasSuffix(pathFile, `\`) {
			pathFile = string(pathFile + `\`)
		}
		PATH_SAVE = strings.ReplaceAll(pathFile, `\`, `\\`)

	} else if WhatsOS() == "darwin" {
		if pathFile == "" {
			PATH_SAVE = `/`
		}
		PATH_SAVE = string(pathFile + "/")
		PermissaoExtra(PATH_SAVE)

	} else {
		if pathFile == "" {
			PATH_SAVE = `/`
		}
		PATH_SAVE = string(pathFile + "/")
	}
}

func PermissaoExtra(pathfile string) {
	params := string(pathfile)
	cmd := string("chmod 777 " + "'" + params + "'")
	command := exec.Command("sh", "-c", cmd)
	command.SysProcAttr = &syscall.SysProcAttr{}
	output, err := command.Output()

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(output)
}
