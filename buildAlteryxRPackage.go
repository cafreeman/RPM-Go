package main

import (
	// "os"
	// "fmt"
	"os/exec"
	"path/filepath"
)

func buildAlteryxRPackage(root string) {
	alteryxRPackageClean(root)
}

func alteryxRPackageClean(root string) {
	root = convertToWindowsPath(root)
	alteryxRPkgPath := filepath.Join(root, "Alteryx", "Plugins", "AlteryxRPackage")

	var commands []*exec.Cmd

	delCmds := delPatternCmd(alteryxRPkgPath)
	rmdirCmd := exec.Command("cmd", "/C", "/y", "/d", filepath.Join(alteryxRPkgPath, "alteryxrdatax.Rcheck"))
	xcopyCmds := xcopyPatternCmd(root, alteryxRPkgPath)

	commands = append(commands, append(delCmds, append(xcopyCmds, rmdirCmd)...)...)

	for _, cmd := range commands {
		err := cmd.Run()
		errCheck(err)
	}
}

// func alteryxRPackageCheck() {
//
// }
//
// func alteryxRPackageInstall() {
//
// }

// func expandPathAndGlob(dirPath, pattern string) []string {
// 	fullPathPattern := filepath.Join(dirPath, pattern)
// 	matches, err := filepath.Glob(fullPathPattern)
// 	errCheck(err)
// 	return matches
// }

func delPatternCmd(alteryxRPkgPath string) []*exec.Cmd {
	delPatterns := []string{
		"build_done",
		"alteryx*.gz",
		"alteryx*.zip",
	}

	var delCommands []*exec.Cmd

	for _, delPattern := range delPatterns {
		delPath := filepath.Join(alteryxRPkgPath, delPattern)
		cmd := exec.Command("cmd", "/C", "del", delPath)
		delCommands = append(delCommands, cmd)
	}

	return delCommands
}

func xcopyPatternCmd(root, alteryxRPkgPath string) []*exec.Cmd {
	xcopyPatterns := [][]string{
		[]string{root, "XSRCLib", "SrcDataWrap2", "RecordLib"},
		[]string{root, "XSRCLib", "inc", "DateTimeValidate.h"},
		[]string{root, "Alteryx", "AlteryxSDKBuilder", "SrcLib_Replacement.h"},
		[]string{root, "3rdParty", "Utility", "liblzf-1.51", "lzf_src.h"},
		[]string{root, "Alteryx", "Open_AlteryxYXDB", "liblzf-3.6", "lzf*.h"},
		[]string{root, "Alteryx", "Open_AlteryxYXDB", "liblzf-3.6", "lzf_*.c"},
		[]string{root, "Alteryx", "Open_AlteryxYXDB", "liblzf-3.6", "config.h"},
		[]string{root, "Alteryx", "Open_AlteryxYXDB", "Open_AlteryxYXDB.h"},
		[]string{root, "Alteryx", "Open_AlteryxYXDB", "Open_AlteryxYXDB.cpp"},
	}

	var xcopyCmds []*exec.Cmd

	for _, pattern := range xcopyPatterns {
		xcopyPath := filepath.Join(pattern...)
		cmd := exec.Command("cmd", "/c", "/y", "/d", xcopyPath, filepath.Join(alteryxRPkgPath, "src"))
		xcopyCmds = append(xcopyCmds, cmd)
	}

	return xcopyCmds
}
