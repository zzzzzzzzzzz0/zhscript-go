package zhscript

import "bytes"

type Buf___ struct {
	*bytes.Buffer
	cur *Buf_item___
	a []*Buf_item___
	
	kw *Keyword___
}

type Buf_item___ struct {
	*bytes.Buffer
	Annota *Strings___
	Val *Val___
}

func New_buf__() *Buf___ {
	buf := new(Buf___)
	buf.get__(0)
	return buf
}

func (this *Buf___) add__() {
	this.Buffer = new(bytes.Buffer)
	this.cur = &Buf_item___{Buffer:this.Buffer, Annota:New_strings__()} 
	this.a = append(this.a, this.cur)
}

func (this *Buf___) add2__(buf2 *Buf___) {
	for i, buf := range buf2.a {
		if i == 0 {
			this.Write(buf.Bytes())
		} else {
			this.a = append(this.a, buf)
		}
	}
}

func (this *Buf___) get__(i int) *Buf_item___ {
	for i2 := len(this.a); i2 <= i; i2++ {
		this.add__()
	}
	this.cur = this.a[i]
	this.Buffer = this.cur.Buffer
	return this.cur
}

func (this *Buf___) set_i__(i int, i2 interface{}) {
	this.get__(i)
	this.cur.Val = &Val___{I:i2, Type:"i"}
}

func (this *Buf___) S__() (s string) {
	for _, buf := range this.a {
		s += buf.String()
	}
	return
}
