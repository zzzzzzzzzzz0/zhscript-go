package zhscript

type code_load___ struct {
	code_z___
}

func (this *code_load___) cls_kw__() *Keyword___ {
	return Kws_.Load
}

func (this *buf_codes___) add_load__(kw, kw2 *Keyword___) *codes___ {
	v := code_load___{code_z___{code_i___:code_i___{kw}, codes:new(codes___), kw2:kw2}}
	this.add__(&v)
	return v.codes
}
