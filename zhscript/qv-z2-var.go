package zhscript

import (
	"strings"
	"strconv"
)

func var_name3__(s string, qv1 *Qv___, ret2 *Var___, shou *Strings___) (qv *Qv___, err2 *Errinfo___) {
	qv = qv1
	if O_liucheng_ {
		o__('x', "%s", s)
	}
	switch s {
	case Kws_.Top.s:
		qv = top_qv_
	case Kws_.Up.s:
		if qv != nil {
			qv = qv.Up
		}
		if qv == nil {
			err2 = New_errinfo__(Errs_.Annota)
			return
		}
	case Kws_.Lock.s:
		if ret2 != nil {
			ret2.Is_lock = true
		}
	case Kws_.Noarg.s:
		if ret2 != nil {
			ret2.Is_no_arg = true
		}
	default:
		if strings.HasPrefix(s, Kws_.Qianarg.s) {
			if ret2 != nil {
				n := s[len(Kws_.Qianarg.s):]
				if n == "" {
					ret2.Qian_arg = 1
				} else {
					i3, err3 := strconv.Atoi(n)
					if err3 == nil {
						ret2.Qian_arg = i3
					} else {
						err2 = New_errinfo__(s, Errs_.Annota)
						return
					}
				}
			}
			break
		}
		//qv = qv.find2__(s, nil)
		if shou != nil {
			shou.Add__(s)
		}
	}
	return
}

func (this *Qv___) var_name2__(code code___, lvl uint, qv1 *Qv___, ret2 *Var___, buf1 *Buf___) (qv *Qv___, err2 *Errinfo___) {
	kw := code.kw__()
	if O_liucheng_ {
		o__('k', "%s", kw)
	}
	qv = qv1
	switch kw {
	case Kws_.Kaifangkuohao:
		for _, code2 := range code.(*code_1___).codes.a {
			if codet, ok := code2.(*code_text___); ok {
				qv, err2 = var_name3__(codet.s, qv, ret2, buf1.cur.Annota)
				if err2 != nil {
					return
				}
			} else {
				err2 = New_errinfo__(Errs_.Annota)
				return
			}
		}
	case Kws_.Kaidanyinhao:
		_, err2 = this.z2_var__(code.(*code_1___).codes, lvl, kw, nil, buf1)
		if err2 != nil {
			return
		}
	case Kws_.Kaihuakuohao:
		for _, code2 := range code.(*code_1___).codes.a {
			qv, err2 = this.var_name2__(code2, lvl, qv, ret2, buf1)
			if err2 != nil {
				return
			}
		}
	default:
		buf2 := New_buf__()
		_, err2 = this.z_code__(code, lvl, buf2)
		if err2 != nil {
			return
		}
		buf1.Write(buf2.Bytes())
	}
	return
}

func (this *Qv___) var_name__(codes *codes___, lvl uint, buf1 *Buf___) (qv *Qv___, err2 *Errinfo___) {
	qv = this
	for _, code := range codes.a {
		qv, err2 = this.var_name2__(code, lvl, qv, nil, buf1)
		if err2 != nil {
			err2.Add__(codes.String())
			return
		}
	}
	buf1.cur.Annota.Find__(func (s string) bool {
		e := qv.Vars.find__(s)
		if e != nil {
			v := var__(e)
			if v.Val.Type == "q" {
				qv = v.Val.I.(*Qv___)
			}
			return true
		}
		return false
	})
	return
}

func (this *Qv___) z2_var2_z__(kw *Keyword___, f__ func(), buf *Buf___) (exist, is_to_buf bool) {
	switch kw {
	case Kws_.Kaidanyinhao:
		f__()
	case Kws_.Has:
		buf.WriteRune('1')
	}
	exist = true
	is_to_buf = true
	return
}

func (this *Qv___) z2_var2__(name string, annota *Strings___, qv  *Qv___, codes *codes___, lvl uint, kw *Keyword___, load *code_z___, buf *Buf___) (bool, *Errinfo___) {
	if O_liucheng_ {
		o__('g', "%s", name)
	}
	var (
		err2 *Errinfo___
		exist, is_to_buf bool
	)
	if qv.Args != nil {
		switch {
		case strings.HasPrefix(name, Kws_.Args.s):
			n := name[len(Kws_.Args.s):]
			from := 0
			max := len(this.Args.A) - 1
			max2 := max + 2
			to := max
			if n != "" {
				var n2 string
				i4 := strings.Index(n, "/")
				if i4 >= 0 {
					n2 = n[i4 + 1:]
					n = n[0:i4]
				}
				i3, err3 := strconv.Atoi(n)
				if err3 == nil {
					from = i3 - 1
				} else {
					break
				}
				if from < 0 {
					from += max2
					if from < 0 {
						from = 0
					}
				}
				if n2 != "" {
					i3, err3 = strconv.Atoi(n2)
					if err3 == nil {
						to = i3 - 1
					} else {
						break
					}
				}
				if to < 0 {
					to += max2
				}
			}
			switch kw {
			case Kws_.Kaidanyinhao:
				if load != nil {
					for i := from; i <= to; i++ {
						load.args.Add2__(this.Args.A[i])
					}
				} else {
					for i := from; i <= to; i++ {
						buf.add_val__(this.Args.A[i].S, "a")
					}
					is_to_buf = true
				}
			case Kws_.Has:
				buf.WriteRune('1')
				is_to_buf = true
			}
			exist = true
		case name == Kws_.Arg.s:
			exist, is_to_buf = this.z2_var2_z__(kw, func() {
				buf.write__(qv.Args.all__())
			}, buf)
		case name == Kws_.Arg.s + "0":
			exist, is_to_buf = this.z2_var2_z__(kw, func() {
				buf.write__(qv.Args.Src2)
			}, buf)
		case name == Kws_.Arg.s + Kws_.Length.s:
			exist, is_to_buf = this.z2_var2_z__(kw, func() {
				buf.write__(strconv.Itoa(len(qv.Args.A)))
			}, buf)
		case strings.HasPrefix(name, Kws_.Arg.s):
			n := name[len(Kws_.Arg.s):]
			i3, err3 := strconv.Atoi(n)
			if err3 == nil {
				i3--
				if i3 >= 0 && i3 < len(qv.Args.A) {
					exist, is_to_buf = this.z2_var2_z__(kw, func() {
						buf.write__(qv.Args.A[i3].S)
					}, buf)
				} else {
					//参数不存在便做空
					exist, is_to_buf = true, true
				}
			}
		}
		/*if !exist {
			for i3, s := range qv.Args.Names {
				if s == name {
					if i3 < len(qv.Args.A) {
						exist, is_to_buf = this.z2_var2_z__(kw, func() {
							buf.write__(qv.Args.A[i3].S)
						}, buf)
					} else {
						//参数不存在便做空
						exist, is_to_buf = true, true
					}
					break
				}
			}
		}*/
	}
	if !exist {
		switch name {
		case Kws_.Vars.s:
			exist, is_to_buf = this.z2_var2_z__(kw, func() {
				buf.cur.Val = &Val___{I:qv, Type:"q"}
				qv.Up = nil
			}, buf)
		}
	}
	if !exist {
		qv = find_qv__(annota, qv)
		var e *Em___
		for {
			e = qv.Vars.find__(name)
			if e != nil {
				break
			}
			switch kw {
			case Kws_.Kaidanyinhao:
				qv = qv.Up
			case Kws_.Has:
				return true, nil
			case Kws_.Del:
				return false, nil
			}
			if qv == nil {
				break
			}
		}
		if e != nil {
			v := var__(e)
			switch kw {
			case Kws_.Kaidanyinhao, Kws_.Has:
				switch v.Kw {
				default:
					switch kw {
					default:
						switch v.Val.Type {
						case "":
							buf.write__(v.Val.S)
						default:
							buf.cur.Val = v.Val
						}
					case Kws_.Has:
						buf.WriteRune('1')
					}
					is_to_buf = true
				case Kws_.Alias:
					v.Annota_val.Find__(func (s string) bool {
						qv, err2 = var_name3__(s, qv, nil, nil/*buf.cur.Annota*/)
						//O__("%v%v",kw,buf.cur.Annota)
						return err2 != nil
					})
					if err2 != nil {
						err2.Add__(codes.String())
						return false, err2
					}
					is_to_buf, err2 = this.z2_var2__(v.Val.S, annota, qv, codes, lvl + 1, kw, load, buf)
					if err2 != nil {
						err2.Add__(v.Kw)
						return false, err2
					}
				}
			case Kws_.Del:
				qv.Vars.del__(e)
			}
			exist = true
		}
	}
	if !exist {
		/*//类似bash的变量不存在便做空
		switch kw {
		case Kws_.Kaidanyinhao:
			return true, nil
		}*/
		return false, New_errinfo__(name, Errs_.Exist, codes.String())
	}
	return is_to_buf, nil
}

func (this *Qv___) z2_var__(codes *codes___, lvl uint, kw *Keyword___, load *code_z___, buf *Buf___) (bool, *Errinfo___) {
	buf1 := New_buf__()
	qv, err2 := this.var_name__(codes, lvl, buf1)
	if err2 != nil {
		return false, err2
	}
	return this.z2_var2__(buf1.String(), buf1.cur.Annota, qv, codes, lvl, kw, load, buf)
}
