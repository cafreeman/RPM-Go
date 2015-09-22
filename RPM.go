package main

import (
	"runtime"
)

var (
	urls = []string{"http://cran.cnr.berkeley.edu/bin/windows/base/old/3.1.3/R-3.1.3-win.exe",
		"https://cran.r-project.org/bin/windows/base/R-3.2.2-win.exe"}
)

func main() {
	var rootPath string
	if runtime.GOOS == "windows" {
		rootPath = `C:/SVN/trunk/3rdParty/R/R_Installed_Files`
	} else {
		rootPath = `/Volumes/C/SVN/trunk/3rdParty/R/R_Installed_Files`
	}

	for _, url := range urls {
		installerPath := downloadR(url, rootPath)
		installDir, rVersion := installR(installerPath, rootPath)
		installRPackages(rootPath, installDir, rVersion)
	}
}
