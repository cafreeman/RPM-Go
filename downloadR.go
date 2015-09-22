package main

import (
	"fmt"
	"github.com/cheggaaa/pb"
	"io"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"time"
)

// DownloadR download an R installer from a specified URL
func downloadR(url string, rootPath string) (installerPath string) {
	// Parse URL and create filename from last element
	tokens := strings.Split(url, "/")
	fileName := tokens[len(tokens)-1]

	// Check if file has already been downloaded. If not, create the file at the specified path
	installerPath = createDownloadPath(rootPath, fileName)
	fmt.Println(installerPath)
	// Check to see if the installer has already been downloaded
	if _, err := os.Stat(installerPath); err == nil {
		fmt.Println(fileName, "already exists!")
		return
	}
	output, err := os.Create(installerPath)
	errCheck(err)
	defer output.Close()

	// Start download process
	fmt.Println("Downloading", url, "to", fileName)

	// Download file from URL
	response, err := http.Get(url)
	errCheck(err)
	defer response.Body.Close()

	// Print http response status to console
	fmt.Println(response.Status)
	// Get the response size from the HTTP header for progress bar
	responseSize, _ := strconv.Atoi(response.Header.Get("Content-Length"))

	// Create progress bar
	bar := pb.New(int(responseSize)).SetUnits(pb.U_BYTES).SetRefreshRate(time.Millisecond * 10)
	bar.ShowSpeed = true
	bar.SetWidth(120)
	bar.Start()

	// Create multi-writer for output destination and progress bar
	writer := io.MultiWriter(output, bar)

	// Copy to output
	_, err = io.Copy(writer, response.Body)
	errCheck(err)
	bar.Finish()

	fmt.Printf("%s with %v bytes downloaded\n", fileName, responseSize)

	// return os-compatible installer path
	return
}

func createDownloadPath(rootPath string, fileName string) string {
	absPath, err := filepath.Abs(rootPath)
	errCheck(err)
	return convertToWindowsPath(filepath.Join(absPath, fileName))
}
