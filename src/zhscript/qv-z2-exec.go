package zhscript

import (
	"fmt"
	"os/exec"
)

func (this *Qv___) z2_exec__(s string, buf *Buf___) {
	o, err := exec.Command("/bin/sh", "-c", s).Output()
	if err != nil {
		fmt.Println(s, err)
	} else {
		buf.Write(o)
	}
}