package zhscript

type code_load___ struct {
	code_z___
}

func (this *code_load___) cls_kw__() *Keyword___ {
	return Kws_.Load
}

func (this *buf_codes___) add_load__(kw, kw2 *Keyword___, argnames []string) *codes___ {
	v := code_load___{code_z___{code_i___:code_i___{kw}, codes:new(codes___), kw11:kw2}}
	switch kw2 {
	case Kws_.Load:
		v.args.Src_type = Src_is_file_
	case Kws_.Interp:
		v.args.Src_type = Src_is_code_
	case Kws_.Def:
		v.args.Src_type = Src_is_varname_
	}
	v.args.Names = argnames
	this.add__(&v)
	return v.codes
}
