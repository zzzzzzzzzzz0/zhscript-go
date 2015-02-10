package zhscript

import (
	"os/exec"
	"strconv"
	"syscall"
)

func (this *Qv___) z2_exec__(s string, buf *Buf___) {
	c := exec.Command("/bin/sh", "-c", s)
	err := c.Start()
	if err != nil {
		buf.get__(1).WriteString(err.Error())
	}
	if err := c.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				buf.WriteString(strconv.Itoa(status.ExitStatus()))
			}
		} else {
			buf.get__(1).WriteString(err.Error())
		}
	}
}