package zhscript

import "github.com/soniah/evaler"
import "strconv"

func (this *Qv___) z2_eval__(buf2, buf *Buf___, codes *codes___) *Errinfo___ {
	s := buf2.get__(0).String()
	var n int
	if len(buf2.a) > 1 {
		var err error
		n, err = strconv.Atoi(buf2.get__(1).String())
		if err != nil {
			buf.get__(1).WriteString(err.Error())
			return nil
		}
	}
	res, err := evaler.Eval(s)
	if err != nil {
		buf.get__(1).WriteString(err.Error())
		return nil
	}
	buf.write__(res.FloatString(n))
	return nil
}