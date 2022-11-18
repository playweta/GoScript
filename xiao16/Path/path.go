package Path

import (
	"log"
	"os"
	"path"
	"path/filepath"
	"runtime"
)

// 运行时项目当前地址
func GetCurrentAbPathByCaller(index int) string {
	var abPath string
	_, filename, _, ok := runtime.Caller(index)
	if ok {
		abPath = path.Dir(filename)
	}
	return abPath
}

// 运行时候AppData\Local\Temp\GoLand 地址 //   打包exe后当前地址
func GetCurrentAbPathByExecutable() string {
	exePath, err := os.Executable()
	if err != nil {
		log.Fatal(err)
	}
	res, _ := filepath.EvalSymlinks(filepath.Dir(exePath))
	return res
}
