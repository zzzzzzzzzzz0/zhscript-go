package zhscript

type code___ interface {
	kw__() *Keyword___
	cls_kw__() *Keyword___
}

func s__(code code___) (string) {
	return s2__(code, nil)
}

func s2__(code code___, for_kw *Keyword___) (s string) {
	kw := code.kw__()
	switch kw {
	case Kws_.Kaiyinhao:
		v := code.(*code_text___)
		switch for_kw {
		case Kws_.Def:
			s += v.s
		default:
			s += Replace_crlf__(v.s)
		}
		return
	}
	s += kw.s
	cls_kw := code.cls_kw__()
	switch cls_kw {
	case Kws_.If:
		v := code.(*code_logic___)
		s += v.logic.String()
		if v.then != nil {
			s += Kws_.Then.s + v.then.String()
		}
		if v.else1 != nil {
			s += Kws_.Else.s + v.else1.String()
		}
	case Kws_.Set:
		v := code.(*code_var___)
		s += v.name.String() + Kws_.Equ.s + v.val.String()
	/*case Kws_.Load, Kws_.Call:
		s += code.(*code_z___).codes.String()*/
	case Kws_.Load:
		s += code.(*code_load___).codes.String()
	case Kws_.Call:
		s += code.(*code_call___).codes.String()
	default:
		if v, ok := code.(*code_1___); ok {
			s += v.codes.String()
			if v.kw2 != nil {
				s += v.kw2.s
			}
		} else {
			return
		}
	}
	switch kw {
	case Kws_.Kaihuakuohao, Kws_.Kaidanyinhao, Kws_.Kaifangkuohao:
		return
	}
	s += Kws_.Juhao.s
	return
}

//inc
type code_i___ struct {
	kw *Keyword___
}

func (this *code_i___) kw__() *Keyword___ {
	return this.kw
}
func (this *code_i___) cls_kw__() *Keyword___ {
	return this.kw
}
