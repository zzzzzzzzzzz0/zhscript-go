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
	name := New_buf__()
	if len(codes.a) == 0 {
		if kw == Kws_.Def {
			return New_errinfo__(Errs_.Defnamenil)
		}
		this.set_var2__(name, buf, 0, false, false, kw)
		return nil
	}
	i := 0
	is_lock := false
	is_no_arg := false
	qv := this
	for _, code := range codes.a {
		switch code.kw__() {
		case Kws_.Dunhao:
			err2 = set_var__(qv, name, buf, i, is_lock, is_no_arg, kw)
			if err2 != nil {
				return err2
			}
			name = New_buf__()
			is_lock = false
			qv = this
			i++
			continue
		}
		qv, is_lock, is_no_arg, err2 = this.var_name2__(code, lvl, qv, is_lock, is_no_arg, name)
		if err2 != nil {
			return err2
		}
	}
	return set_var__(qv, name, buf, i, is_lock, is_no_arg, kw)
}

func set_var__(qv *Qv___, name, buf *Buf___, i int, is_lock, is_no_arg bool, kw *Keyword___) *Errinfo___ {
	if buf.Annota.Len() > 0 {
		qv2 := qv
		for {
			if qv2 == nil {
				break
			}
			if qv2.vars.find__(name.String(), name.Annota) != nil {
				qv = qv2
				break
			}
			qv2 = qv2.up
		}
	}
	return qv.set_var2__(name, buf, i, is_lock, is_no_arg, kw)
}

func (this *Qv___) set_var2__(name, buf2 *Buf___, i int, is_lock, is_no_arg bool, kw *Keyword___) *Errinfo___ {
	if buf2.kw == Kws_.Kaiyinhao && i >= len(buf2.A) {
		i = len(buf2.A) - 1
	}
	buf2.get__(i)
	return this.Set_var__(name, buf2, is_lock, is_no_arg, kw)
}

func (this *Qv___) Set_var__(name, val *Buf___, is_lock, is_no_arg bool, kw *Keyword___) *Errinfo___ {
	var2 := this.vars.get__(name.String(), name.Annota)
	
	if var2.Is_lock {
		return New_errinfo__(Errs_.Lock, name)
	}
	var2.Is_lock = is_lock
	var2.Is_no_arg = is_no_arg
	
	var2.Val = val.Buffer.String()
	var2.Annota = name.Annota
	var2.Annota_val = val.Annota
	var2.Kw = kw
	return nil
}
