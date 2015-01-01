package zhscript

import (
	"bytes"
)

type Buf___ struct {
	*bytes.Buffer
	A []*bytes.Buffer
	kw *Keyword___
}

func New_buf__() *Buf___ {
	buf := new(Buf___)
	buf.get__(0)
	return buf
}

func (this *Buf___) add__() {
	this.Buffer = new(bytes.Buffer)
	this.A = append(this.A, this.Buffer)
}

func (this *Buf___) add2__(buf2 *Buf___) {
	for i, buf := range buf2.A {
		if i == 0 {
			this.Write(buf.Bytes())
		} else {
			this.A = append(this.A, buf)
		}
	}
}

func (this *Buf___) get__(i int) *bytes.Buffer {
	for i2 := len(this.A); i2 <= i; i2++ {
		this.add__()
	}
	return this.A[i]
}

func (this *Buf___) String() (s string) {
	for _, buf := range this.A {
		s += buf.String()
	}
	return
}
