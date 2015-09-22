package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

func installR(installerPath, rootPath string) (installDir, rVersion string) {
	// Create windows-friendly path to R installation directory for a specific version of R
	installDir, rVersion = createInstallDir(installerPath, rootPath)
	// Build Command struct
	cmd := exec.Command(installerPath,
		"/SILENT",
		fmt.Sprintf(`/DIR=%v`, installDir))

	// Execute the Command struct and handle errors
	err := cmd.Run()
	errCheck(err)
	// Return forward-slash installDir path
	installDir = filepath.ToSlash(installDir)
	return
}

// A function for parsing the R version from the installer file and creating the correct
// installation directory for R
func createInstallDir(installerPath string, rootPath string) (installDir, rVersionNum string) {
	fileName := filepath.Base(installerPath)
	rVersion := strings.TrimSuffix(fileName, "-win.exe")
	installDir = convertToWindowsPath(filepath.Join(rootPath, rVersion))
	err := os.MkdirAll(installDir, 0755)
	errCheck(err)
	rVersionNum = strings.TrimPrefix(rVersion, "R-")
	return
}
