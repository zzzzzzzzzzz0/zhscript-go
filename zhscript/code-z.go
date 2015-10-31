package zhscript

type code_z___ struct {
	code_i___
	codes *codes___
	args Args___
	kw11 *Keyword___
}

func (this *code_z___) z2__(code code___, qv *Qv___, lvl uint, buf *Buf___) (s string, is_to_buf bool, err *Errinfo___) {
	kw := code.kw__()
	switch kw {
	case Kws_.Kaidanyinhao:
		is_to_buf, err = qv.z2_var__(code.(*code_1___).codes, lvl, kw, this, buf)
		if err != nil {
			return
		}
		if !is_to_buf {
			return
		}
		s = buf.String()
	default:
		buf2 := New_buf__()
		_, err = qv.z_code__(code, lvl, buf2)
		if err != nil {
			return
		}
		s = buf2.String()
		is_to_buf = true
	}
	if O_liucheng_ {
		o__('x', "%s", s)
	}
	return
}

func (this *code_z___) z__(qv *Qv___, lvl uint, buf *Buf___) (goto1 *Goto___, err *Errinfo___) {
	l := len(this.codes.a)
	if l == 0 {
		return
	}
	this.args.Reset__()
	var (
		s string
		is_to_buf bool
	)
	switch this.args.Src_type {
	case Src_is_varname_:
		s = s2__(this.codes.a[0], this.kw11)
	default:
		buf2 := New_buf__()
		s, is_to_buf, err = this.z2__(this.codes.a[0], qv, lvl, buf2)
		if err != nil {
			return
		}
	}
	this.args.Src = s
	if l > 2 {
		var (
			is_eoc bool
			s2 string
		)
		for i := 2; !is_eoc; {
			val := &Val___{}
			var has_to_buf, has_noto_buf bool
			for {
				if i >= l {
					is_eoc = true
					if !has_to_buf && has_noto_buf {
					} else {
						has_to_buf = true
					}
					break
				}
				code := this.codes.a[i]
				i++
				if code.kw__() == Kws_.Dunhao {
					has_to_buf = true
					break
				}
				buf2 := New_buf__()
				s2, is_to_buf, err = this.z2__(code, qv, lvl, buf2)
				if err != nil {
					return
				}
				if !is_to_buf {
					has_noto_buf = true
					continue
				}
				val.S += s2
				val.copy_i__(buf2.cur.Val)
				has_to_buf = true
			}
			if has_to_buf {
				this.args.Add2__(val)
			}
		}
	}
	switch this.kw {
	case Kws_.Call:
		goto1, err = call__(&this.args, qv, buf)
		if err != nil {
			err.Add__(this.codes.String())
			return
		}
	default:
		var qv2 *Qv___
		qv2, err = New_qv__(&this.args, qv)
		if err == nil {
			goto1, err = qv2.Z__(lvl, buf)
			if goto1 != nil && goto1.Kw == Kws_.Return && !qv2.is_through {
				goto1 = nil
			}
			return
		}
	}
	return
}
