package zhscript

import "strings"

type codes___ struct {
	a []code___
}

func (this *codes___) String() (s string) {
	for _, code := range this.a {
		s += s__(code)
	}
	return
}

func for_codes__(codes *codes___, f func (code code___) bool) bool {
	return for_codes2__(codes, f, false)
}

func for_codes2__(codes *codes___, f func (code code___) bool, reverse bool) bool {
	if codes == nil {
		return false
	}
	var (
		i int
		c func() code___
	)
	l := len(codes.a)
	if reverse {
		i = l
		c = func() code___ {
			i--
			if i < 0 {
				return nil
			}
			return codes.a[i]
		}
	} else {
		i = -1
		c = func() code___ {
			i++
			if i >= l {
				return nil
			}
			return codes.a[i]
		}
	}
	for {
		code := c()
		if code == nil {
			break
		}
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

func o_codes__(codes *codes___, lvl uint) {
	var head, head2 string
	for i := uint(0); i < lvl; i++ {
		head2 += "|---"
	}
	head_kw := func(kw *Keyword___) {
		o__(' ', "%s|%s\n", head, kw)
	}
	for _, code := range codes.a {
		head = head2
		kw := code.kw__()
		o__(' ', "%s%s", head, kw.s)
		cls_kw := code.cls_kw__()
		switch cls_kw {
		case Kws_.Kaiyinhao:
			v := code.(*code_text___)
			o__(' ', "%s\n", Replace_crlf__(v.s) + "|")
		case Kws_.If:
			v := code.(*code_logic___)
			o_n__()
			o_codes__(v.logic, lvl + 1)
			if v.then != nil {
				head_kw(Kws_.Then)
				o_codes__(v.then, lvl + 1)
			}
			if v.else1 != nil {
				head_kw(Kws_.Else)
				o_codes__(v.else1, lvl + 1)
			}
		case Kws_.Set:
			v := code.(*code_var___)
			o_n__()
			o_codes__(v.name, lvl + 1)
			head_kw(Kws_.Equ)
			o_codes__(v.val, lvl + 1)
		default:
			var codes *codes___
			switch cls_kw {
			case Kws_.Load:
				code2 := code.(*code_load___)
				if code2.kw11 != kw {
					o__(' ', "%s", code2.kw11)
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
				o_codes__(codes, lvl + 1)
			}
		}
	}
}

func Replace_crlf__(s string) string {
	r := strings.NewReplacer("\r", Kws_.CR.s, "\n", Kws_.LF.s, "\t", ".")
	return r.Replace(s)
}