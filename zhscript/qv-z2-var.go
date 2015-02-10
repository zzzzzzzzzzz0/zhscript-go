package zhscript

import (
	"strings"
	"strconv"
)

func (this *Qv___) var_name3__(s string, qv1 *Qv___, is_lock1, is_no_arg1 bool, shou *List___) (qv *Qv___, is_lock, is_no_arg bool, err2 *Errinfo___) {
	qv = qv1
	is_lock = is_lock1
	is_no_arg = is_no_arg1
	if o_liucheng_ {
		o__('x', "%s", s)
	}
	switch s {
	case Kws_.Up.s:
		qv = qv.up
		if qv == nil {
			err2 = New_errinfo__(Errs_.Annota)
			return
		}
	case Kws_.Lock.s:
		is_lock = true
	case Kws_.Noarg.s:
		is_no_arg = true
	case Kws_.Top.s:
		qv = top_qv_
	default:
		if shou != nil {
			shou.PushBack(s)
		}
	}
	return
}

func (this *Qv___) var_name2__(code code___, lvl uint, qv1 *Qv___, is_lock1, is_no_arg1 bool, buf1 *Buf___) (qv *Qv___, is_lock, is_no_arg bool, err2 *Errinfo___) {
	kw := code.kw__()
	if o_liucheng_ {
		o__('k', "%s", kw)
	}
	qv = qv1
	is_lock = is_lock1
	is_no_arg = is_no_arg1
	switch(kw) {
	case Kws_.Kaifangkuohao:
		for _, code2 := range code.(*code_1___).codes.a {
			if codet, ok := code2.(*code_text___); ok {
				qv, is_lock, is_no_arg, err2 = this.var_name3__(codet.s, qv, is_lock, is_no_arg, buf1.Annota)
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
			qv, is_lock, is_no_arg, err2 = this.var_name2__(code2, lvl, qv, is_lock, is_no_arg, buf1)
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

func (this *Qv___) var_name__(codes *codes___, lvl uint, buf1 *Buf___) (qv *Qv___, is_lock, is_no_arg bool, err2 *Errinfo___) {
	qv = this
	for _, code := range codes.a {
		qv, is_lock, is_no_arg, err2 = this.var_name2__(code, lvl, qv, is_lock, is_no_arg, buf1)
		if err2 != nil {
			err2.Add__(codes.String())
			return
		}
	}
	return
}

func (this *Qv___) z2_var2__(name string, annota *List___, qv  *Qv___, codes *codes___, lvl uint, kw *Keyword___, load *code_z___, buf *Buf___) (bool, *Errinfo___) {
	if o_liucheng_ {
		o__('g', "%s", name)
	}
	var (
		err2 *Errinfo___
		exist, is_to_buf bool
	)
	if qv.args != nil {
		switch {
		case strings.HasPrefix(name, Kws_.Args.s):
			n := name[len(Kws_.Args.s):]
			from := 0
			if n != "" {
				i3, err3 := strconv.Atoi(n)
				if err3 == nil {
					from = i3 - 1
				} else {
					break
				}
			}
			switch kw {
			case Kws_.Kaidanyinhao:
				if load != nil {
					for i := from; i < len(this.args.A); i++ {
						load.args.Add__(this.args.A[i])
					}
				} else {
					for i := from; i < len(this.args.A); i++ {
						if i > from {
							buf.WriteString(" ")
						}
						buf.WriteString(this.args.A[i])
					}
					is_to_buf = true
				}
			case Kws_.Has:
				buf.WriteRune('1')
				is_to_buf = true
			}
			exist = true
		case name == Kws_.Arg.s:
			switch kw {
			case Kws_.Kaidanyinhao:
				buf.WriteString(qv.args.all__())
			case Kws_.Has:
				buf.WriteRune('1')
			}
			exist = true
			is_to_buf = true
		case name == Kws_.Arg.s + "0":
			switch kw {
			case Kws_.Kaidanyinhao:
				switch qv.args.Src_type {
				case Src_is_code_:
				default:
					buf.WriteString(qv.args.Src)
				}
			case Kws_.Has:
				buf.WriteRune('1')
			}
			exist = true
			is_to_buf = true
		case name == Kws_.Arg.s + Kws_.Length.s:
			switch kw {
			case Kws_.Kaidanyinhao:
				buf.WriteString(strconv.Itoa(len(qv.args.A)))
			case Kws_.Has:
				buf.WriteRune('1')
			}
			exist = true
			is_to_buf = true
		case strings.HasPrefix(name, Kws_.Arg.s):
			n := name[len(Kws_.Arg.s):]
			i3, err3 := strconv.Atoi(n)
			if err3 == nil {
				i3--
				if i3 >= 0 && i3 < len(qv.args.A) {
					switch kw {
					case Kws_.Kaidanyinhao:
						buf.WriteString(qv.args.A[i3])
					case Kws_.Has:
						buf.WriteRune('1')
					}
					exist = true
					is_to_buf = true
				}
			}
		}
	}
	if !exist {
		var e *Em___
		for {
			e = qv.vars.find__(name, annota)
			if e != nil {
				break
			}
			switch kw {
			case Kws_.Kaidanyinhao:
				qv = qv.up
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
						buf.WriteString(v.Val)
					case Kws_.Has:
						buf.WriteRune('1')
					}
					is_to_buf = true
				case Kws_.Alias:
					v.Annota_val.Find__(func (e *Em___) bool {
						qv, _, _, err2 = this.var_name3__(e.String(), qv, false, false, nil)
						return err2 != nil
					})
					if err2 != nil {
						err2.Add__(codes.String())
						return false, err2
					}
					is_to_buf, err2 = this.z2_var2__(v.Val, annota, qv, codes, lvl + 1, kw, load, buf)
					if err2 != nil {
						err2.Add__(v.Kw)
						return false, err2
					}
				}
			case Kws_.Del:
				this.vars.del__(e)
			}
			exist = true
		}
	}
	if !exist {
		return false, New_errinfo__(name, Errs_.Exist, codes.String())
	}
	return is_to_buf, nil
}

func (this *Qv___) z2_var__(codes *codes___, lvl uint, kw *Keyword___, load *code_z___, buf *Buf___) (bool, *Errinfo___) {
	buf1 := New_buf__()
	qv, _, _, err2 := this.var_name__(codes, lvl, buf1)
	if err2 != nil {
		return false, err2
	}
	return this.z2_var2__(buf1.String(), buf1.Annota, qv, codes, lvl, kw, load, buf)
}
