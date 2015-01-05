package zhscript

func (this *Qv___) z2_set__(var1 *code_var___, lvl uint, kw *Keyword___) *Errinfo___ {
	buf := New_buf__()
	a := var1.val.a
	for {
		if len(a) == 0 {
			break
		}
		kw := a[0].kw__()
		if kw == Kws_.Kaihuakuohao {
			a = a[0].(*code_1___).codes.a
		} else {
			buf.kw = kw
			break
		}
	}
	_, err2 := this.z2__(var1.val, lvl, buf)
	if err2 != nil {
		return err2
	}
	codes := var1.name
	if len(codes.a) == 0 {
		if kw == Kws_.Def {
			return New_errinfo__(Errs_.Defnamenil)
		}
		this.set_var2__("", buf, 0, false, kw)
		return nil
	}
	i := 0
	name := ""
	is_lock := false
	qv := this
	for _, code := range codes.a {
		switch code.kw__() {
		case Kws_.Dunhao:
			err2 = qv.set_var2__(name, buf, i, is_lock, kw)
			if err2 != nil {
				return err2
			}
			name = ""
			is_lock = false
			qv = this
			i++
			continue
		}
		buf2 := New_buf__()
		qv, is_lock, err2 = this.var_name2__(code, lvl, qv, is_lock, codes, buf2)
		if err2 != nil {
			return err2
		}
		name += buf2.String()
	}
	return qv.set_var2__(name, buf, i, is_lock, kw)
}

func (this *Qv___) set_var2__(name string, buf2 *Buf___, i int, is_lock bool, kw *Keyword___) *Errinfo___ {
	var val string
	if buf2.kw == Kws_.Kaiyinhao && i >= len(buf2.A) {
		i = len(buf2.A) - 1
	}
	if i >= 0 && i < len(buf2.A) {
		val = buf2.get__(i).String()
	} else {
		val = ""
	}
	return this.Set_var__(name, val, is_lock, kw)
}

func (this *Qv___) Set_var__(name, val string, is_lock bool, kw *Keyword___) *Errinfo___ {
	var2 := this.vars.get__(name)
	
	if var2.is_lock {
		return New_errinfo__(Errs_.Lock, name)
	}
	var2.is_lock = is_lock
	
	var2.kw = kw
	var2.val = val
	return nil
}
