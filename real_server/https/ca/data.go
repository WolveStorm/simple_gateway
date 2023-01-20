package ca

import (
	"path/filepath"
	"runtime"
)

var dir string

func init() {
	_, file, _, _ := runtime.Caller(0)
	dir = filepath.Dir(file)
}

func Path(rel string) string {
	//如果传入的是绝对路径
	if filepath.IsAbs(rel) {
		return rel
	}
	join := filepath.Join(dir, rel)
	//否则将拼接上在的目录，变成绝对路径
	return join
}
