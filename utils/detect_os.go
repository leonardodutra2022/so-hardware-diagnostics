package utils

import (
	"runtime"
)

func DetectOS() string {
	if runtime.GOOS == "linux" {
		return "linux"
	} else if runtime.GOOS == "darwin" {
		return "darwin"
	}
	return "windows"
}

func WhatsOS() string {
	output := DetectOS()
	return output

}
