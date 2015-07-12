package zhscript

import (
	"fmt"
	"os/exec"
	"path/filepath"
)

type Zhscript___ struct {
	Args Args___
}

func New__(a []string, cc func([]byte, string) []byte, i interface{}) (z *Zhscript___, err *Errinfo___) {
	z = new(Zhscript___)
	z.Args.Parse__(a, 1, 0)
	
	call_cfg__(i)
	content_convert_ = cc
	
	file, _ := exec.LookPath(a[0])
	path, _ := filepath.Abs(file)
	Known_path_add__(path)
	
	if top_qv_ == nil {
		top_qv_, err = New_qv__(nil, nil)
		if err != nil {
			return
		}
	}

	return
}

func (this *Zhscript___) Z__() (s string, err *Errinfo___) {
	var qv *Qv___

	qv, err = this.New_main_qv__(&this.Args)
	if err != nil {
		return
	}
	
	buf := New_buf__()
	_, err = qv.Z__(0, buf)
	s = buf.String()
	return
}

func (this *Zhscript___) New_main_qv__(args *Args___) (*Qv___, *Errinfo___) {
	return New_qv__(args, top_qv_)
}

var	top_qv_ *Qv___

func o__(r rune, format string, a ...interface{}) (n int, err error) {
	if o_ansi_ {
		head := "\x1b[0;3"
		switch r {
		case 'k':
			head += "7;42"
		case 'K':
			head += "3;42"
		case 'r':
			head += "0;42"
		case 'e':
			head += "1;42"
		case 'x':
			head += "7;41"
		case 'g':
			head += "1"
		case 'n':
			head += "4"
		default:
			head += "2"
		}
		format = head + "m" + format + "\x1b[0m"
	} else {
		if r != ' ' {
			head := "{"
			if r > 0 {
				head += string(r) + " "
			}
			format = head + format + "}"
		}
	}
	return fmt.Printf(format, a...)
}
func O__(format string, a ...interface{}) (n int, err error) {
	return o__(0, format, a...)
}
func o_n__() {
	fmt.Println()
}
func o_t__() {
	fmt.Print("\t")
}