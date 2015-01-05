package zhscript

import (
	"container/list"
)

func is_def__(code []rune, up_qv *Qv___, i int, thiz, thiz1 *buf_codes___) (int, *Keyword___, *buf_codes___, bool) {
	var thiz2 *buf_codes___
	f1 := func(name string) bool {
		name2 := []rune(name)
		if _, ok := Startswith__(code, name2, i); ok {
			thiz2 = new_buf_codes__(Kws_.Def)
			thiz2.add_text2__(name)
			thiz2.separ__()
			i += len(name2)
			return true
		}
		return false
	}
	f := func(code code___) bool {
		if code.kw__() == Kws_.Def {
			name := "" 
			for_codes__(code.(*code_var___).name, func(code code___) bool {
				if code.kw__() == Kws_.Kaiyinhao {
					name += code.(*code_text___).s
				}
				return false
			})
			if f1(name) {
				return true
			}
		}
		return false
	}
	if for_codes__(thiz.codes, f) {
		return i, Kws_.Interp, thiz2, true
	}
	if thiz1 != nil {
		if for_codes__(thiz1.codes, f) {
			return i, Kws_.Interp, thiz2, true
		}
	}
	v := for_var__(f1, Kws_.Def, up_qv)
	if v != nil {
		return i, Kws_.Interp, thiz2, true
	}
	return i, nil, thiz2, false
}

func for_var__(f1 func(name string) bool, kw *Keyword___, qv1 *Qv___) (v2 *var___) {
	qv := qv1
	for {
		if qv.vars.ls.Find__(func(e *list.Element) bool {
			v := var__(e)
			if v.kw == kw {
				if f1(v.name) {
					v2 = v
					return true
				}
			}
			return false
		}) {
			return
		}
		qv = qv.up_qv
		if qv == nil {
			break
		}
	}
	return
}