package zhscript

import (
	"github.com/soniah/evaler"
)

func (this *Qv___) z2_eval__(s string, buf *Buf___, codes *codes___) *Errinfo___ {
	res, err := evaler.Eval(s)
	if err != nil {
		buf.get__(1).WriteString(err.Error())
		return nil
	}
	buf.WriteString(res.FloatString(0))
	return nil
}