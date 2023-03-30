package utils

import (
	"errors"
	"log"
	"os"

	"github.com/leonardodutra9898/tools-display-info/model"
	"golang.org/x/text/encoding/charmap"
)

func SaveInfoSO(info *model.SysInfo) {

	var output string = "Information System\n"
	output += string("Hostname: \t" + info.Hostname + "\n")
	output += string("CPU: \t\t" + info.CPU + "\n")
	output += string("RAM: \t\t" + IntToStr(int(info.RAM)) + "GB\n")
	output += string("Disk: \t\t" + IntToStr(int(info.Disk)) + "GB\n")
	output += string("Plataform: \t" + info.Platform + "\n")
	SaveInDisk(output, "system_info.txt")
}

func SaveInDisk(data string, nameFile string) {
	dataFile := []byte(data)
	if err := os.WriteFile(nameFile, dataFile, 0755); err != nil {
		panic(err)
	}
}

func SaveInDiskStrings(data []string, nameFile string) {
	dataFile := []byte(JoinStrings(data))
	if err := os.WriteFile(nameFile, dataFile, 0755); err != nil {
		panic(err)
	}
}

func Execute(command []byte, err error, reportName string) {
	if err != nil {
		log.Fatal(err)
	}
	decoder := charmap.CodePage850.NewDecoder()
	out, err := decoder.Bytes(command)

	if err != nil {
		log.Fatal(err)
	}

	// fmt.Println(string(out))
	f, err := os.Create(PATH_SAVE + reportName)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.Write(out)
}

func Exists(namePath string) (bool, error) {
	_, err := os.Stat(namePath)
	if err == nil {
		return true, nil
	}
	if errors.Is(err, os.ErrNotExist) {
		return false, nil
	}
	return false, err
}
