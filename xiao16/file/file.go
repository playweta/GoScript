package file

import (
	"io/ioutil"
	"sync"
)

var lock = sync.Mutex{}

func ReadFile(path string) ([]byte, error) {
	file, err := ioutil.ReadFile(path)
	return file, err
}
func WriteFileAppend(path string, data []byte) error {
	lock.Lock()
	file, _ := ReadFile(path)
	err := ioutil.WriteFile(path, append(file, data...), 0777)
	lock.Unlock()
	return err
}
func WriteFile(path string, data []byte) error {
	lock.Lock()
	err := ioutil.WriteFile(path, data, 0777)
	lock.Unlock()
	return err
}
