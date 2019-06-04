package pidfile

import (
	"io/ioutil"
	"os"
	"strconv"
	"testing"
)

func TempFile() (*os.File, error) {
	if f, err := ioutil.TempFile("", "pidfile-test-"); err == nil {
		return f, nil
	} else {
		return nil, err
	}
}

func TempFileName() (string, error) {
	if f, err := TempFile(); err == nil {
		f.Close()
		return f.Name(), nil
	} else {
		return "", err
	}
}

func TestInitialize(t *testing.T) {
	newPath := "/tmp/pidfilePath"

	Initialize(newPath)

	if newPath != pidfilePath {
		t.Errorf("cannot assign pidfile")
	}
}

func TestGetPidfilePath(t *testing.T) {
	pidfilePath = "/tmp/pidfilePath"

	if GetPidfilePath() != pidfilePath {
		t.Errorf("not equal pidfile path")
	}
}

func TestRead(t *testing.T) {
	var err error

	pidfile, err := TempFile()
	if err != nil {
		t.Error(err)
	}

	pidfilePath = pidfile.Name()
	defer pidfile.Close()

	_, err = pidfile.WriteString(strconv.FormatInt((int64)(os.Getpid()), 10))
	if err != nil {
		t.Error(err)
	}

	_, err = Read()
	if err != nil {
		t.Error(err)
	}
}

func TestWrite(t *testing.T) {
	var err error

	pidfilePath, err = TempFileName()
	if err != nil {
		t.Error(err)
	}

	err = Write()
	if err != nil {
		t.Error(err)
	}
}

func TestRemove(t *testing.T) {
	var err error

	pidfilePath, err = TempFileName()
	if err != nil {
		t.Error(err)
	}

	err = Remove()
	if err != nil {
		t.Error(err)
	}
}

func TestGetProcess(t *testing.T) {
	var err error

	pidfile, err := TempFile()
	if err != nil {
		t.Error(err)
	}

	pidfilePath = pidfile.Name()
	defer pidfile.Close()

	_, err = pidfile.WriteString(strconv.FormatInt((int64)(os.Getpid()), 10))
	if err != nil {
		t.Error(err)
	}

	process, err := GetProcess()
	if err != nil {
		t.Error(err)
	}

	if process == nil {
		t.Errorf("process not found.")
	}

	if process.Pid != os.Getpid() {
		t.Errorf("pid is not equal.")
	}
}
