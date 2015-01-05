package zhscript

type code_text___ struct {
	code_i___
	s string
}

func (this *code_text___) kw2__() *Keyword___ {
	return Kws_.Biyinhao
}

func (this *buf_codes___) add_text__(buf *Buf___) {
	if buf.Len() <= 0 {
		return
	}
	this.add_text2__(buf.String())
	buf.Reset()
}

func (this *buf_codes___) add_text2__(s string) {
	v := code_text___{code_i___{Kws_.Kaiyinhao}, s}
	this.add__(&v)
}
