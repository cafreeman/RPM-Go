package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"runtime"
)

var (
	urls = []string{"http://cran.cnr.berkeley.edu/bin/windows/base/old/3.1.3/R-3.1.3-win.exe",
		"https://cran.r-project.org/bin/windows/base/R-3.2.2-win.exe"}
)

func main() {
	var root string
	var rInstallDir string
	if runtime.GOOS == "windows" {
		root = svnRoot()
		rInstallDir = filepath.ToSlash(filepath.Join(root, "3rdParty", "R", "R_Installed_Files"))
	} else {
		rInstallDir = `/Volumes/C/SVN/trunk/3rdParty/R/R_Installed_Files`
	}

	fmt.Print("Press enter to continue...")
	bufio.NewReader(os.Stdin).ReadBytes('\n')

	for _, url := range urls {
		installerPath := downloadR(url, rInstallDir)
		installDir, rVersion := installR(installerPath, rInstallDir)
		installRPackages(rInstallDir, installDir, rVersion)
	}
}
