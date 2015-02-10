package zhscript

func is_def__(code []rune, up_qv *Qv___, i1 int, thiz1 *buf_codes___) (i int, kw *Keyword___, thiz *buf_codes___, is_no_arg, ok bool) {
	i = i1
	f1 := func(name string, is_no_arg1 bool) bool {
		name2 := []rune(name)
		if _, ok2 := Startswith__(code, name2, i); ok2 {
			thiz = new_buf_codes__(Kws_.Def)
			thiz.add_text2__(name)
			if !is_no_arg1 {
				thiz.separ__()
			}
			i += len(name2)
			is_no_arg = is_no_arg1
			kw = Kws_.Interp
			ok = true
			return true
		}
		return false
	}
	f := func(code code___) bool {
		if code.kw__() == Kws_.Def {
			var (
				name string
				is_no_arg1 bool
			)
			for_codes__(code.(*code_var___).name, func(code code___) bool {
				switch code.kw__() {
					case Kws_.Kaiyinhao:
					name += code.(*code_text___).s
					case Kws_.Kaifangkuohao:
					for_codes__(code.(*code_1___).codes, func(code2 code___) bool {
						if codet, ok := code2.(*code_text___); ok {
							switch codet.s {
								case Kws_.Noarg.s:
								is_no_arg1 = true
							}
						}
						return false
					});
				}
				return false
			})
			if f1(name, is_no_arg1) {
				return true
			}
		}
		return false
	}
	thiz3 := thiz1
	for {
		if thiz3 == nil {
			break
		}
		if for_codes__(thiz3.codes, f) {
			return
		}
		thiz3 = thiz3.up
	}
	v := for_var__(f1, Kws_.Def, up_qv)
	if v != nil {
		return
	}
	return
}

func for_var__(f1 func(name string, is_no_arg bool) bool, kw *Keyword___, qv1 *Qv___) (v2 *Var___) {
	qv := qv1
	for {
		if qv.vars.ls.Find__(func(e *Em___) bool {
			v := var__(e)
			if v.Kw == kw {
				if f1(v.Name, v.Is_no_arg) {
					v2 = v
					return true
				}
			}
			return false
		}) {
			return
		}
		qv = qv.up
		if qv == nil {
			break
		}
	}
	return
}