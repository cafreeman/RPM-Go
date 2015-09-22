package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
)

func installRPackages(rootPath, installDir, version string) {
	packageDir := packageRepoDir(rootPath, version)
	dirContents, err := ioutil.ReadDir(packageDir)
	errCheck(err)
	// Initialize the cmdArgs slice with the two static arguments we need to pass to R.exe, then
	// iterate over the directory contents and append all of the winbows binaries to cmdArgs
	cmdArgs := []string{"CMD", "INSTALL"}
	for _, elem := range dirContents {
		// Extract filename from os.FileInfo struct
		name := elem.Name()
		// Filter out only the files that end in `.zip` (since these will be the Windows binaries), and
		// append to the `filtered` slice
		if filepath.Ext(name) == ".zip" {
			fullPath := filepath.Join(packageDir, name)
			cmdArgs = append(cmdArgs, fullPath)
		}
	}

	// Construct path to R executable
	rCMDPath := filepath.Join(convertToWindowsPath(installDir), "bin", "R.exe")
	// Create Command struct with arguments from the cmdArgs slice
	cmd := exec.Command(rCMDPath, cmdArgs...)
	// Execute R CMD INSTALL for all packages and print the stdout
	out, err := cmd.Output()
	errCheck(err)
	fmt.Println(string(out))
}

func packageRepoDir(rootPath, version string) string {
	rDir := filepath.Dir(convertToWindowsPath(rootPath))
	return filepath.Join(rDir, "localRepo", "bin", "windows", "contrib", version[:3])
}
