package zhscript

import (
	"io/ioutil"
	"strings"
	"container/list"
	"os"
)

var known_path_ = new(List___)

func Known_path_add__(path string) {
	i := strings.LastIndex(path, "/")
	if i > 0 {
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
		if !known_path_.Find__(func(e *list.Element) bool {return e.Value.(string) == dir}) {
			if o_path_ {
				o__(0, "known_path %s|", dir)
				o_n__()
			}
			known_path_.PushBack(dir)
		}
	}
}

func ReadFile__(src string) (b []byte, err error) {
	src2, _ := Get_path__(src)
	b, err = ioutil.ReadFile(src2)
	if err == nil {
		Known_path_add__(src2)
	}
	return
}

func Get_path__(src string) (string, bool) {
	e2 := known_path_.Front()
	src2 := src
	for {
		_, err := os.Stat(src2)
		if err == nil {
			return src2, true
		}
		if e2 == nil {
			break
		}
		src2 = e2.Value.(string) + "/" + src
		if o_path_ {
			o__(0, "src %s", src2)
			o_n__()
		}
		e2 = e2.Next()
	}
	return src, false
}