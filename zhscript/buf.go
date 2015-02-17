package zhscript

import (
	"bytes"
)

type Buf___ struct {
	*bytes.Buffer
	A []*bytes.Buffer
	Annota *List___
	AA []*List___
	kw *Keyword___
}

func New_buf__() *Buf___ {
	buf := new(Buf___)
	buf.get__(0)
	return buf
}

func (this *Buf___) add__() {
	this.Buffer = new(bytes.Buffer)
	this.Annota = new(List___)
	this.A = append(this.A, this.Buffer)
	this.AA = append(this.AA, this.Annota)
}

func (this *Buf___) add2__(buf2 *Buf___) {
	for i, buf := range buf2.A {
		if i == 0 {
			this.Write(buf.Bytes())
		} else {
			this.A = append(this.A, buf)
		}
	}
	for _, a := range buf2.AA {
		this.AA = append(this.AA, a)
	}
}

func (this *Buf___) get__(i int) *bytes.Buffer {
	for i2 := len(this.A); i2 <= i; i2++ {
		this.add__()
	}
	this.Buffer = this.A[i]
	this.Annota = this.AA[i]
	return this.Buffer
}

func (this *Buf___) S__() (s string) {
	for _, buf := range this.A {
		s += buf.String()
	}
	return
}
