package datasetutil

import (
	"fmt"
	"os"
	"os/exec"
	"time"
)

func Extract7Zip(file string, password string) error {
	command := fmt.Sprintf("/usr/bin/7z x -o'files/' %s -p'%s'", file, password)
	cmd := exec.Command("sh", "-c", command)
	err := cmd.Run()
	if err != nil {
		return err
	}

	// Remove 7z file
	err = os.Remove(file)
	if err != nil {
		return err
	}

	return nil
}

func Compress7Zip(file string) error {
	command := fmt.Sprintf("/usr/bin/7z a -t7z -m0=lzma2 -mx=9 -mfb=64 -md=32m -ms=on -mhe=on %s.7z %s", file, file)
	cmd := exec.Command("sh", "-c", command)
	err := cmd.Run()
	if err != nil {
		return err
	}
	return nil
}

func TimedPruneDirectory(directory string, seconds int) {
	// Wait for s seconds
	time.Sleep(time.Duration(seconds) * time.Second)
	
	// Remove all files in the directory
	err := os.RemoveAll(directory)
	if err != nil {
		fmt.Println(err)
	}
}