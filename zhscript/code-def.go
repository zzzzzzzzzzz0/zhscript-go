package zhscript

func is_def__(code []rune, up_qv *Qv___, i1 int, buf1 *Buf___, thiz1 *buf_codes___) (i int, kw *Keyword___, thiz *buf_codes___, is_no_arg, ok bool) {
	i = i1
	f1 := func(v *Var___) bool {
		name2 := []rune(v.Name)
		len2 := len(name2)
		if len2 > 0 {
			if _, ok2 := Startswith__(code, name2, i); ok2 {
				thiz = new_buf_codes__(Kws_.Def)
				thiz.add_text2__(v.Name)

				thiz1.add_text__(buf1)
				if v.Qian_arg > 0 {
					qian := v.Qian_arg
					a := thiz1.codes.a
					i2 := len(a) - 1
					for {
						if i2 < 0 {
							i2 = 0
							thiz.separ__()
							break
						}
						code := a[i2]
						if code.kw__() == Kws_.Dunhao {
							qian--
							if qian <= 0 {
								break
							}
						}
						i2--
						if thiz1.kw == nil && qian <= 1 {
							thiz.separ__()
							break
						}
					}
					for i3 := i2; i3 < len(a); i3++ {
						thiz.add__(a[i3])
					}
					thiz1.codes.a = thiz1.codes.a[0:i2]
				}

				is_no_arg = v.Is_no_arg
				if !is_no_arg {
					thiz.separ__()
				}

				i += len2
				kw = Kws_.Interp
				ok = true
				return true
			}
		}
		return false
	}
	f := func(code code___) bool {
		if code.kw__() == Kws_.Def {
			set2 := new(Var___)
			for_codes__(code.(*code_var___).name, func(code code___) bool {
				switch code.kw__() {
					case Kws_.Kaiyinhao:
					set2.Name += code.(*code_text___).s
					case Kws_.Kaifangkuohao:
					for_codes__(code.(*code_1___).codes, func(code2 code___) bool {
						if codet, ok := code2.(*code_text___); ok {
							var_name3__(codet.s, nil, set2, nil)
						}
						return false
					})
				}
				return false
			})
			if f1(set2) {
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

func for_var__(f1 func(*Var___) bool, kw *Keyword___, qv1 *Qv___) (v2 *Var___) {
	qv := qv1
	for {
		if qv.Vars.Ls.Find__(func(e *Em___) bool {
			v := Var__(e)
			if v.Kw == kw {
				if f1(v) {
					v2 = v
					return true
				}
			}
			return false
		}) {
			return
		}
		qv = qv.Up
		if qv == nil {
			break
		}
	}
	return
}