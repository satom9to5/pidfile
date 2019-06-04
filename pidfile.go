package pidfile

import (
	"bytes"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strconv"
)

var (
	pidfilePath     = ""
	notPidfileError = errors.New("pidfile not configured.")
)

func Initialize(varPidfilePath string) {
	pidfilePath = varPidfilePath
}

func GetPidfilePath() string {
	return pidfilePath
}

func Read() (int, error) {
	if pidfilePath == "" {
		return 0, notPidfileError
	}

	b, err := ioutil.ReadFile(pidfilePath)
	if err != nil {
		return 0, err
	}

	pid, err := strconv.Atoi(string(bytes.TrimSpace(b)))
	if err != nil {
		return 0, fmt.Errorf("pidfile parse error from %s: %s", pidfilePath, err)
	}

	return pid, nil
}

func Write() error {
	if pidfilePath == "" {
		return notPidfileError
	}

	if err := os.MkdirAll(filepath.Dir(pidfilePath), os.FileMode(0755)); err != nil {
		return err
	}

	file, err := os.Create(pidfilePath)
	if err != nil {
		return fmt.Errorf("open pidfile error %s: %s", pidfilePath, err)
	}
	defer file.Close()

	_, err = fmt.Fprintf(file, "%d", os.Getpid())

	return err
}

func Remove() error {
	if pidfilePath == "" {
		return notPidfileError
	} else {
		return os.Remove(pidfilePath)
	}
}

func GetProcess() (*os.Process, error) {
	if pid, err := Read(); err == nil {
		return os.FindProcess(pid)
	} else {
		return nil, err
	}
}
