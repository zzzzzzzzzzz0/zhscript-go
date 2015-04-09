package zhscript

type Qv___ struct {
	args *Args___
	Vars *Vars___
	Annota *List___
	codes *codes___
	Up *Qv___
	Not_my interface{}
}

func (this *Qv___) find2__(s string, on_no func(string)) *Qv___ {
	b := true
	for ea := this.Annota.Front(); ea != nil; ea = ea.Next() {
		if ea.Value.(string) == s {
			b = false
			break
		}
	}
	if b {
		if on_no != nil {
			on_no(s)
		}
		return nil
	}
	return this
}

func (this *Qv___) find__(annota *List___, on_no func(string)) *Qv___ {
	for ea2 := annota.Front(); ea2 != nil; ea2 = ea2.Next() {
		if this.find2__(ea2.Value.(string), on_no) == nil {
			return nil
		}
	}
	return this
}

func find_qv__(annota *List___, qv *Qv___) *Qv___ {
	if annota.Len() > 0 {
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
	if o_args_ {
		switch this.args.Src_type {
		case Src_is_code_:
			o__('n', "...")
		default:
			o__('n', "%s", this.args.Src)
		}
		o_n__()
		for i, s := range this.args.A {
			o__('n', "%d) %s", i, s)
			o_n__()
		}
	}

	return this.z2__(this.codes, lvl, buf)
}

var content_convert_ func([]byte, string) []byte

func New_qv__(args *Args___, up_qv *Qv___) (*Qv___, *Errinfo___) {
	var code []byte
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
			if content_convert_ != nil {
				code = content_convert_(content, args.Src)
			} else {
				code = content
			}
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
			code = []byte(v.Val)
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
	qv.args = args
	qv.codes = codes
	qv.Up = up_qv
	if up_qv != nil {
		qv.Not_my = up_qv.Not_my
	}
	qv.Annota = new(List___)
	
	qv.Vars = new(Vars___)
	qv.Vars.init__()
	
	return qv, nil
}
