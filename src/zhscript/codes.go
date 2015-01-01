package zhscript

import (
	"strings"
)

type codes___ struct {
	a []code___
}

func (this *codes___) String() (s string) {
	for _, code := range this.a {
		s += s__(code, nil)
	}
	return
}

func for_codes__(codes *codes___, f func (code code___) bool) bool {
	if codes == nil {
		return false
	}
	for _, code := range codes.a {
		if code.kw__() == Kws_.Kaihuakuohao {
			if for_codes__(code.(*code_1___).codes, f) {
				return true
			}
		} else {
			if f(code) {
				return true
			}
		}
	}
	return false
}

func for_o_codes__(codes *codes___, lvl uint) {
	var head2 string
	for i := uint(0); i < lvl; i++ {
		head2 += "|---"
	}
	for _, code := range codes.a {
		var head string
		head += head2
		kw := code.kw__()
		o__(' ', "%s", head)
		o__(' ', "%s", kw.s)
		cls_kw := code.cls_kw__()
		switch(cls_kw) {
		case Kws_.Kaiyinhao:
			v := code.(*code_text___)
			o__(' ', "%s\n", Replace_crlf__(v.s) + "|")
		case Kws_.If:
			v := code.(*code_logic___)
			o_n__()
			for_o_codes__(v.logic, lvl + 1)
			if v.then != nil {
				o__(' ', "%s|%s\n", head, Kws_.Then)
				for_o_codes__(v.then, lvl + 1)
			}
			if v.else1 != nil {
				o__(' ', "%s|%s\n", head, Kws_.Else)
				for_o_codes__(v.else1, lvl + 1)
			}
		case Kws_.Set:
			v := code.(*code_var___)
			o_n__()
			for_o_codes__(v.name, lvl + 1)
			o__(' ', "%s|%s\n", head, Kws_.Equ)
			for_o_codes__(v.val, lvl + 1)
		default:
			var codes *codes___
			switch(cls_kw) {
			case Kws_.Load:
				code2 := code.(*code_load___)
				if code2.kw2 != kw {
					o__(' ', "/%s", code2.kw2)
				}
				codes = code2.codes
			case Kws_.Call:
				codes = code.(*code_call___).codes
			default:
				if code1, ok := code.(*code_1___); ok {
					codes = code1.codes
				}
			}
			o_n__()
			if codes != nil {
				for_o_codes__(codes, lvl + 1)
			}
		}
	}
}

func Replace_crlf__(s string) string {
	r := strings.NewReplacer("\r", Kws_.CR.s, "\n", Kws_.LF.s, "\t", ".")
	return r.Replace(s)
}