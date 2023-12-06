package pkg

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
	"strings"
	"syscall"
)

func Cd(currentDir *string, dir string) error {
	if dir == "" {
		*currentDir = "/home/rus/"
		return nil
	}

	dirEntries, err := os.ReadDir(*currentDir)

	if err != nil {
		return err
	}

	for _, entry := range dirEntries {
		if dir == entry.Name() {
			*currentDir = *currentDir + dir + "/"
			return nil
		}
	}

	return fmt.Errorf("Directory isn't found")
}

func Kill(pid int) error {
	process, err := os.FindProcess(pid)

	if err != nil {
		return err
	}

	err = process.Signal(syscall.SIGTERM)

	if err != nil {
		return err
	}

	return nil
}

func Ps() error {
	procDir := "/proc/"

	files, err := ioutil.ReadDir(procDir)

	if err != nil {
		return err
	}

	for _, file := range files {
		pid, err := strconv.Atoi(file.Name())

		if err != nil || !file.IsDir() {
			continue
		}

		cmdLineFile := filepath.Join(procDir, strconv.Itoa(pid), "cmdline")

		cmdLineData, err := ioutil.ReadFile(cmdLineFile)

		if err != nil {
			fmt.Println(err)
			continue
		}

		procName := strings.Replace(string(cmdLineData), "\x00", " ", -1)

		fmt.Printf("PID: %d, Name: %s\n", pid, procName)
	}

	return nil
}
