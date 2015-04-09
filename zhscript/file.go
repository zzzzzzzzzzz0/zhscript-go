package zhscript

import (
	"io/ioutil"
	"strings"
	"os"
)

var known_path_ = new(List___)

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
	if !known_path_.Find__(func(e *Em___) bool {return e.Value.(string) == dir}) {
		if o_path_ {
			o__(0, "known_path %s|", dir)
			o_n__()
		}
		known_path_.PushBack(dir)
	}
}

func ReadFile__(src string) ([]byte, bool) {
	if src2, base, ok := Get_path__(src); ok {
		b, err := ioutil.ReadFile(src2)
		if err == nil {
			if base == "" {
				Known_path_add__(src2)
			}
			return b, true
		}
	}
	return nil, false
}

func Get_path__(src string) (src2, dir string, ok bool) {
	e2 := known_path_.Front()
	src2 = src
	for {
		_, err := os.Stat(src2)
		if err == nil {
			ok = true
			if dir == "" {
				dir = Get_dir__(src)
			}
			return
		}
		if e2 == nil {
			break
		}
		dir = e2.Value.(string)
		src2 = dir + "/" + src
		if o_path_ {
			o__(0, "src %s dir %s", src2, dir)
			o_n__()
		}
		e2 = e2.Next()
	}
	return
}

func Get_dir__(src string) (base string) {
	e2 := known_path_.Front()
	for {
		if e2 == nil {
			break
		}
		base2 := e2.Value.(string)
		if strings.HasPrefix(src, base2) && len(base2) > len(base) {
			base = base2
		}
		e2 = e2.Next()
	}
	return
}