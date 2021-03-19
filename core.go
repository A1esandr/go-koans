package gokoans

import (
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"runtime"
	"strings"
)

const (
	__true__  = true
	__false__ = false
)

func assert(o bool, text string) {
	if !o {
		fmt.Printf("\n%s\n\n", __getRecentLine())
		os.Exit(1)
	}
}

func __getRecentLine() string {
	_, file, line, _ := runtime.Caller(2)
	buf, _ := ioutil.ReadFile(file)
	code := strings.TrimSpace(strings.Split(string(buf), "\n")[line-1])
	return fmt.Sprintf("%v:%d\n%s", path.Base(file), line, code)
}
