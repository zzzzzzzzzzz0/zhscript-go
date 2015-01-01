package zhscript

type code_z___ struct {
	code_i___
	codes *codes___
	args Args___
	kw2 *Keyword___
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
	if o_liucheng_ {
		o__('x', "%s", s)
	}
	return
}

func (this *code_z___) z__(qv *Qv___, lvl uint, buf *Buf___) (goto1 *Goto___, err *Errinfo___) {
	l := len(this.codes.a)
	if l == 0 {
		return
	}
	this.args.reset__()
	var (
		s string
		is_to_buf bool
	)
	switch this.kw2 {
	case Kws_.Def:
		s = s__(this.codes.a[0], this.kw2)
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
			b2 bool
			s2 string
		)
		for i := 2; !b2; {
			s = ""
			b := false
			for {
				if i >= l {
					b2 = true
					break
				}
				code := this.codes.a[i]
				i++
				if code.kw__() == Kws_.Dunhao {
					break
				}
				buf2 := New_buf__()
				s2, is_to_buf, err = this.z2__(code, qv, lvl, buf2)
				if err != nil {
					return
				}
				if !is_to_buf {
					continue
				}
				s += s2
				b = true
			}
			if b {
				this.args.Add__(s)
			}
		}
	}
	switch(this.kw) {
	case Kws_.Call:
		goto1, err = call__(&this.args, qv, buf)
		if err != nil {
			err.Add__(this.codes.String())
			return
		}
	default:
		switch this.kw2 {
		case Kws_.Load:
			this.args.Src_type = Src_is_file_
		case Kws_.Interp:
			this.args.Src_type = Src_is_code_
		case Kws_.Def:
			this.args.Src_type = Src_is_varname_
		} 
		var qv2 *Qv___
		qv2, err = New_qv__(&this.args, qv)
		if err == nil {
			goto1, err = qv2.Z__(lvl, buf)
			if goto1 != nil && goto1.I == G_return_ {
				goto1 = nil
			}
			return
		}
	}
	return
}
