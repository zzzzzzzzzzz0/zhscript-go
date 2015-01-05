package zhscript

type code_1___ struct {
	code_i___
	kw2 *Keyword___
	codes *codes___
}

func kaihuakuohao__() *code_1___ {
	return &code_1___{code_i___:code_i___{Kws_.Kaihuakuohao}, codes:new(codes___)}
}

func (this *buf_codes___) add_2__(kw, kw2 *Keyword___) *codes___ {
	v := code_1___{code_i___{kw}, kw2, new(codes___)}
	this.add__(&v)
	return v.codes
}
