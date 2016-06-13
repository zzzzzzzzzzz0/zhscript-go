package zhscript

import (
	"os/exec"
	"strconv"
	"syscall"
	"os"
)

func (this *Qv___) z2_exec__(s string, buf *Buf___) {
	c := exec.Command("/bin/sh", "-c", s)
	c.Stdout = os.Stdout
	c.Stderr = os.Stderr
	c.Stdin = os.Stdin
	err := c.Start()
	if err != nil {
		buf.get__(1).WriteString(err.Error())
	}
	if err := c.Wait(); err != nil {
		if exiterr, ok := err.(*exec.ExitError); ok {
			if status, ok := exiterr.Sys().(syscall.WaitStatus); ok {
				buf.get__(0).WriteString(strconv.Itoa(status.ExitStatus()))
			}
		} else {
			buf.get__(1).WriteString(err.Error())
		}
	}
	buf.get__(2).WriteString(strconv.Itoa(c.Process.Pid))
}