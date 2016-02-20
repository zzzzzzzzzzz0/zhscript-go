package zhscript

type Qv___ struct {
	Args *Args___
	Vars *Vars___
	Annota *Strings___
	codes *codes___
	Up *Qv___
	is_through bool
	Not_my interface{}
}

func (this *Qv___) find2__(s string, on_no__ func(string)) *Qv___ {
	for _, s2 := range this.Annota.A {
		if s2 == s {
			return this
		}
	}
	if on_no__ != nil {
		on_no__(s)
	}
	return nil
}

func (this *Qv___) find__(annota *Strings___, on_no__ func(string)) *Qv___ {
	for _, s := range annota.A {
		if this.find2__(s, on_no__) == nil {
			return nil
		}
	}
	return this
}

func find_qv__(annota *Strings___, qv *Qv___) *Qv___ {
	if annota.Len__() > 0 {
		qv2 := qv
		for {
			if qv2 == nil {
				break
			}
			if qv2.find__(annota, nil) != nil {
				qv = qv2
				break
			}
			qv2 = qv2.Up
		}
	}
	return qv
}

func (this *Qv___) Z__(lvl uint, buf *Buf___) (*Goto___, *Errinfo___) {
	if O_args_ {
		switch this.Args.Src_type {
		case Src_is_code_:
			o__('n', "...")
		default:
			o__('n', "%s", this.Args.Src)
		}
		O_n__()
		o__('n', "%v %v", this.Args.Src2, this.is_through)
		O_n__()
		for i, s := range this.Args.A {
			o__('n', "%d) %s", i, s)
			O_n__()
		}
	}

	return this.z2__(this.codes, lvl, buf)
}

var content_convert__ func([]byte, string) []byte

func New_qv__(args *Args___, up_qv *Qv___) (*Qv___, *Errinfo___) {
	var (
		code []byte
		is_through bool
	)
	if args != nil {
		switch args.Src_type {
		case Src_is_file_:
			content, ok := ReadFile__(args.Src)
			if !ok {
				return nil, New_errinfo__(args.Src, Errs_.Openfile)
			}
			l := len(content)
			if l > 2 && content[0] == '#' && content[1] == '!' {
				for i := 2; i < l; i++ {
					if content[i] == '\n' {
						content = content[i + 1:]
						break
					}
				}
			}
			if content_convert__ != nil {
				code = content_convert__(content, args.Src)
			} else {
				code = content
			}
			args.Src2 = args.Src
		case Src_is_varname_:
			v := for_var__(func(v *Var___) bool {
				if v.Name == args.Src {
					return true
				}
				return false
			}, Kws_.Def, up_qv)
			if v == nil {
				return nil, New_errinfo__(args.Src, Errs_.Exist, Kws_.Def)
			}
			code = []byte(v.Val.S)
			args.Src2 = args.Src
			is_through = v.Is_through
		default:
			code = []byte(args.Src)
		}
	}
	
	codes, id := codecache__(code)
	if codes == nil {
		var err2 *Errinfo___
		codes, err2 = codecache2__(code, id, up_qv)
		if err2 != nil {
			return nil, err2
		}
	}
	
	qv := new(Qv___)
	qv.Args = args
	qv.codes = codes
	qv.Up = up_qv
	qv.is_through = is_through
	if up_qv != nil {
		qv.Not_my = up_qv.Not_my
	}
	qv.Annota = New_strings__()
	
	qv.Vars = new(Vars___)
	qv.Vars.init__()
	if args != nil {
		switch args.Src_type {
		case Src_is_varname_:
			for i, s := range args.Names {
				if i >= len(args.A) {
					break
				}
				name := New_buf__()
				name.WriteString(s)
				val := New_buf__()
				val.WriteString(args.A[i].S)
				qv.Set_var__(name, val, nil, Kws_.Set)
			}
		}
	}
	
	return qv, nil
}
