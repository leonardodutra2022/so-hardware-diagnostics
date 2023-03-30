package utils

import (
	"archive/zip"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

var pathSave string

func Zip() {
	pathSave = PATH_SAVE
	var reportFull string = REPORT_FULL
	var sysInfo string = REPORT_SYS
	var programs string = REPORT_PROGRAMS
	var monitor string = REPORT_MONITOR
	var memory string = REPORT_MEMORY
	var energy string = REPORT_ENERGY_WINDOWS
	var domain string = REPORT_DOMAIN
	var more string = REPORT_EXTRA
	var extraDriver string = REPORT_EXTRA_DRIVER
	var extraMonitor string = REPORT_EXTRA_MONITOR
	var extraMonitorII string = REPORT_EXTRA_MONITOR_II
	var extraUsb string = REPORT_EXTRA_USB
	var extraEdid string = REPORT_EXTRA_EDID
	var bios string = REPORT_BIOS

	fmt.Println("Creating zip archive...")
	archive, err := os.Create(pathSave + "DiagnosticLogs.zip")
	if err != nil {
		panic(err)
	}
	defer archive.Close()
	zipWriter := zip.NewWriter(archive)

	f1, err := os.Open(pathSave + reportFull)
	if err != nil {
		panic(err)
	}
	defer f1.Close()

	w1, err := zipWriter.Create(reportFull)
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(w1, f1); err != nil {
		panic(err)
	}

	f2, err := os.Open(pathSave + sysInfo)
	if err != nil {
		panic(err)
	}
	defer f2.Close()

	w2, err := zipWriter.Create(sysInfo)
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(w2, f2); err != nil {
		panic(err)
	}

	f3, err := os.Open(pathSave + programs)
	if err != nil {
		panic(err)
	}
	defer f3.Close()

	w3, err := zipWriter.Create(programs)
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(w3, f3); err != nil {
		panic(err)
	}

	f4, err := os.Open(pathSave + monitor)
	if err != nil {
		panic(err)
	}
	defer f4.Close()

	w4, err := zipWriter.Create(monitor)
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(w4, f4); err != nil {
		panic(err)
	}

	f5, err := os.Open(pathSave + memory)
	if err != nil {
		panic(err)
	}
	defer f5.Close()

	w5, err := zipWriter.Create(memory)
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(w5, f5); err != nil {
		panic(err)
	}

	checkReportEnergy, _ := Exists(pathSave + energy)
	if checkReportEnergy {
		f6, err := os.Open(pathSave + energy)
		if err != nil {
			panic(err)
		}
		defer f6.Close()

		w6, err := zipWriter.Create(energy)
		if err != nil {
			panic(err)
		}
		if _, err := io.Copy(w6, f6); err != nil {
			panic(err)
		}
	}

	f7, err := os.Open(pathSave + domain)
	if err != nil {
		panic(err)
	}
	defer f7.Close()

	w7, err := zipWriter.Create(domain)
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(w7, f7); err != nil {
		panic(err)
	}

	f8, err := os.Open(pathSave + more)
	if err != nil {
		panic(err)
	}
	defer f8.Close()

	w8, err := zipWriter.Create(more)
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(w8, f8); err != nil {
		panic(err)
	}

	f10, err := os.Open(pathSave + extraDriver)
	if err != nil {
		panic(err)
	}
	defer f10.Close()

	w10, err := zipWriter.Create(extraDriver)
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(w10, f10); err != nil {
		panic(err)
	}

	f11, err := os.Open(pathSave + extraMonitor)
	if err != nil {
		panic(err)
	}
	defer f11.Close()

	w11, err := zipWriter.Create(extraMonitor)
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(w11, f11); err != nil {
		panic(err)
	}

	f12, err := os.Open(pathSave + extraMonitorII)
	if err != nil {
		panic(err)
	}
	defer f12.Close()

	w12, err := zipWriter.Create(extraMonitorII)
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(w12, f12); err != nil {
		panic(err)
	}

	f13, err := os.Open(pathSave + extraUsb)
	if err != nil {
		panic(err)
	}
	defer f13.Close()

	w13, err := zipWriter.Create(extraUsb)
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(w13, f13); err != nil {
		panic(err)
	}

	f14, err := os.Open(pathSave + extraEdid)
	if err != nil {
		panic(err)
	}
	defer f14.Close()

	w14, err := zipWriter.Create(extraEdid)
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(w14, f14); err != nil {
		panic(err)
	}

	f15, err := os.Open(pathSave + bios)
	if err != nil {
		panic(err)
	}
	defer f15.Close()

	w15, err := zipWriter.Create(bios)
	if err != nil {
		panic(err)
	}
	if _, err := io.Copy(w15, f15); err != nil {
		panic(err)
	}

	zipWriter.Close()
}

func ZipFolder(pathReports string) {
	zipFile := string(pathReports + "DiagnosticLogs.zip")
	// Get a Buffer to Write To
	outFile, err := os.Create(zipFile)
	if err != nil {
		fmt.Println(err)
	}
	defer outFile.Close()

	// Create a new zip archive.
	w := zip.NewWriter(outFile)

	// Add some files to the archive.
	addFiles(w, pathReports, "")

	if err != nil {
		fmt.Println(err)
	}

	// Make sure to check the error on Close.
	err = w.Close()
	if err != nil {
		fmt.Println(err)
	}
}

func addFiles(w *zip.Writer, basePath, baseInZip string) {
	// Open the Directory
	files, err := ioutil.ReadDir(basePath)
	if err != nil {
		fmt.Println(err)
	}

	for _, file := range files {
		if !strings.HasSuffix(file.Name(), ".zip") {
			if strings.HasPrefix(file.Name(), "report_") {
				if !file.IsDir() {
					dat, err := ioutil.ReadFile(basePath + file.Name())
					if err != nil {
						fmt.Println(err)
					}
					// Add some files to the archive.
					f, err := w.Create(baseInZip + file.Name())
					if err != nil {
						fmt.Println(err)
					}
					_, err = f.Write(dat)
					if err != nil {
						fmt.Println(err)
					}

				} else if file.IsDir() {
					newBase := basePath + file.Name() + "/"
					addFiles(w, newBase, baseInZip+file.Name()+"/")
				}
			}
		}
	}
}
