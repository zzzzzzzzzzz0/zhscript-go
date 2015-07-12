package zhscript

import "fmt"

func (this *Qv___) z_code__(code code___, lvl uint, buf *Buf___) (*Goto___, *Errinfo___) {
	kw := code.kw__()
	if o_liucheng_ {
		o__('k', "%s", kw)
	}
	switch kw {
	case Kws_.Kaiyinhao:
		t := code.(*code_text___)
		buf.WriteString(t.s)
		return nil, nil
	case Kws_.Dunhao:
		buf.add__()
		return nil, nil
	}
	var codes2 *codes___
	cls_kw := code.cls_kw__()
	switch cls_kw {
	case Kws_.Kaidanyinhao, Kws_.Has, Kws_.Del:
		_, err2 := this.z2_var__(code.(*code_1___).codes, lvl, kw, nil, buf)
		return nil, err2
	case Kws_.If:
		var err2 *Errinfo___
		codes2, err2 = code.(*code_logic___).z__(this, lvl)
		if err2 != nil {
			return nil, err2
		}
		if codes2 == nil {
			return nil, nil
		}
	case Kws_.For:
		return this.z2_for__(code.(*code_1___).codes, lvl, buf)
	case Kws_.Set:
		return nil, this.z2_set__(code.(*code_var___), lvl, kw)
	case Kws_.Load, Kws_.Call:
		var code2 code_z___
		switch cls_kw {
		default:
			code2 = code.(*code_load___).code_z___
		case Kws_.Call:
			code2 = code.(*code_call___).code_z___
		}
		return code2.z__(this, lvl, buf)
	default:
		code1, ok := code.(*code_1___)
		if !ok {
			return nil, New_errinfo__(s__(code), Errs_.Case)
		}
		codes2 = code1.codes
	}

	buf2 := New_buf__()
	goto2, err2 := this.z2__(codes2, lvl + 1, buf2)
	switch kw {
	case Kws_.Kaihuakuohao, Kws_.If:
		buf.add2__(buf2)
	}
	if err2 != nil || goto2 != nil {
		return goto2, err2
	}
	if o_liucheng_ {
		o__('K', "(%d) %s", lvl, kw)
	}
	switch kw {
	case Kws_.Eval:
		err3 := this.z2_eval__(buf2, buf, codes2)
		if err3 != nil {
			return nil, err3
		}
	default:
		s := buf2.S__()
		if o_liucheng_ {
			o__('g', "%s", Replace_crlf__(s))
		}
		switch kw {
		case Kws_.Kaifangkuohao:
			buf.cur.Annota.Add__(s)
			
		case Kws_.Break:
			return &Goto___{Kws_.Break, s}, nil
		case Kws_.Continue:
			return &Goto___{Kws_.Continue, s}, nil
			
		case Kws_.Quit:
			return &Goto___{Kws_.Quit, s}, nil
		case Kws_.Return:
			return &Goto___{Kws_.Return, s}, nil
			
		case Kws_.Echo:
			fmt.Print(s)
		case Kws_.Exec:
			this.z2_exec__(s, buf)
		}
	}
	return nil, nil
}

func (this *Qv___) z2__(codes *codes___, lvl uint, buf *Buf___) (*Goto___, *Errinfo___) {
	if o_liucheng_ {
		o_n__()
		o__('n', "(%d)", lvl)
		o__(0, "%s", codes.String())
		o_n__()
	}
	for _, code := range codes.a {
		goto2, err2 := this.z_code__(code, lvl, buf)
		if err2 != nil || goto2 != nil {
			return goto2, err2
		}
	}
	return nil, nil
}
