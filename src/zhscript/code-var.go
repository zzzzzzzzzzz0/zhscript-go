package zhscript

type code_var___ struct {
	code_i___
	name, val *codes___
}

func (this *code_var___) cls_kw__() *Keyword___ {
	return Kws_.Set
}

func new_var__(kw *Keyword___) *code_var___ {
	return &code_var___{code_i___:code_i___{kw}, name:new(codes___)}
}
