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

	// Construct path to R executable
	rCMDPath := filepath.Join(convertToWindowsPath(installDir), "bin", "R.exe")

	for _, elem := range dirContents {
		name := elem.Name()
		if filepath.Ext(name) == ".zip" {
			fullpath := filepath.Join(packageDir, name)
			cmdString := exec.Command(rCMDPath, "CMD", "INSTALL", fullpath)
			out, err := cmdString.Output()
			fmt.Println(string(out))
			fmt.Println(err)
			// cmdArgs = append(cmdArgs, cmdString)
		}
	}
}

func packageRepoDir(rootPath, version string) string {
	rDir := filepath.Dir(convertToWindowsPath(rootPath))
	return filepath.Join(rDir, "localRepo", "bin", "windows", "contrib", version[:3])
}
