package zhscript

import (
	"io/ioutil"
	"strings"
	"os"
)

var known_path_ = New_strings__()

func Known_path_add__(path string) {
	i := strings.LastIndex(path, "/")
	if i <= 0 {
		return
	}
	dir := path[0:i]
	for {
		i = len(dir)
		if i == 0 {
			return
		}
		i--
		if dir[i] == '/' {
			dir = dir[0:i]
		} else {
			break
		}
	}
	if dir == "" {
		return
	}
	if !known_path_.Find__(func(s string)bool {return s == dir}) {
		if O_path_ {
			o__(0, "known_path %s|", dir)
			O_n__()
		}
		known_path_.Add__(dir)
	}
}

func ReadFile__(src string) ([]byte, bool) {
	if src2, _, ok := Get_path__(src); ok {
		b, err := ioutil.ReadFile(src2)
		if err == nil {
			Known_path_add__(src2)
			return b, true
		}
	}
	return nil, false
}

func Get_path__(src string) (src2, dir string, ok bool) {
	var i int
	src2 = src
	for {
		if O_path_ {
			o__(0, "src %s dir %s", src2, dir)
			O_n__()
		}
		_, err := os.Stat(src2)
		if err == nil {
			ok = true
			if dir == "" {
				known_path_.Find__(func(dir2 string)bool {
					if strings.HasPrefix(src, dir2) && len(dir2) > len(dir) {
						dir = dir2
					}
					return false
				})
			}
			return
		}
		if i >= known_path_.Len__() {
			break
		}
		dir = known_path_.A[i]
		src2 = dir + "/" + src
		i++
	}
	return
}
