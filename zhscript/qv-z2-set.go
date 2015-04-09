package zhscript

import "strings"

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
		return this.set_var2__(name, buf, 0, nil, kw)
	}
	i := 0
	qv := this
	ret2 := new(Var___)
	for _, code := range codes.a {
		switch code.kw__() {
		case Kws_.Dunhao:
			err2 = set_var__(qv, name, buf, i, ret2, kw)
			if err2 != nil {
				return err2
			}
			name = New_buf__()
			qv = this
			ret2 = new(Var___)
			i++
			continue
		}
		qv, err2 = this.var_name2__(code, lvl, qv, ret2, name)
		if err2 != nil {
			return err2
		}
	}
	return set_var__(qv, name, buf, i, ret2, kw)
}

func set_var__(qv *Qv___, name, buf *Buf___, i int, set2 *Var___, kw *Keyword___) *Errinfo___ {
	qv = find_qv__(name.Annota, qv)
	return qv.set_var2__(name, buf, i, set2, kw)
}

func (this *Qv___) set_var2__(name, buf2 *Buf___, i int, set2 *Var___, kw *Keyword___) *Errinfo___ {
	if buf2.kw == Kws_.Kaiyinhao && i >= len(buf2.A) {
		i = len(buf2.A) - 1
	}
	buf2.get__(i)
	return this.Set_var__(name, buf2, set2, kw)
}

func (this *Qv___) Set_var__(name, val *Buf___, set2 *Var___, kw *Keyword___) (err *Errinfo___) {
	name2 := name.String()
	if kw == Kws_.Def {
		len2 := len(name2)
		if len2 == 0 {
			err = New_errinfo__(name2, Errs_.Defnamenil)
		} else {
			for_var__(func(v *Var___) bool {
				if len2 > len(v.Name) && strings.HasPrefix(name2, v.Name) {
					err = New_errinfo__(name2, Errs_.Defnamenil, v.Name)
					return true
				}
				return false
			}, Kws_.Def, this)
		}
		if err != nil {
			return
		}
	}

	var2 := this.Vars.get__(name2)
	if var2.Is_lock {
		err = New_errinfo__(Errs_.Lock, name)
		return
	}
	if set2 != nil {
		var2.Is_lock = set2.Is_lock
		var2.Is_no_arg = set2.Is_no_arg
		var2.Qian_arg = set2.Qian_arg
	}
	var2.Val = val.String()
	var2.Annota_val = val.Annota
	var2.Kw = kw
	this.find__(name.Annota, func(s string) {
		this.Annota.PushBack(s)
	})
	return
}
