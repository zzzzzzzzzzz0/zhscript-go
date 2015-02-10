package zhscript

import (
	"strconv"
)

type code_logic___ struct {
	code_i___
	logic, then, else1 *codes___
	result bool
	left_val string
}

func (this *code_logic___) z2__(codes *codes___, qv *Qv___, lvl, lvl2 uint) (bool, *Errinfo___) {
	var (
		left_kw *Keyword___
		result, has_result, is_not, is_right bool
		s string
	)
	for i := 0; i < len(codes.a); i++ {
		code := codes.a[i]
		kw := code.kw__()
		if o_liucheng_ {
			o_n__()
			o__('n', "if(%d)%d)%d)%s", lvl, lvl2, i, kw)
			o_n__()
		}
		b := true
		if kw.is2__(m_logic_) {
			switch kw {
			case Kws_.Dengyu, Kws_.Notdengyu, Kws_.Xiaoyudengyu, Kws_.Xiaoyu, Kws_.Dayudengyu, Kws_.Dayu:
				left_kw = kw
				is_right = true
			case Kws_.And:
				result = this.result__(left_kw, s, is_not)
				has_result = true
				if !result {
					b = false
				}
			case Kws_.Or:
				result = this.result__(left_kw, s, is_not)
				has_result = true
				if result {
					b = false
				}
			case Kws_.Not:
				is_not = !is_not
			}
			s = ""
			if has_result {
				left_kw = nil
				is_not = false
				is_right = false
			}
			if b {
				continue
			}
			break
		}
		if kw == Kws_.Kaihuakuohao {
			var err *Errinfo___
			result, err = this.z2__(code.(*code_1___).codes, qv, lvl, lvl2 + 1)
			if err != nil {
				return false, err
			}
			if result {
				break
			}
		} else {
			var err *Errinfo___
			buf := New_buf__()
			_, err = qv.z_code__(code, lvl + lvl2, buf)
			if err != nil {
				return false, err
			}
			s += buf.String()
		}
		if !is_right {
			if s != "" {
				this.left_val = s
			}
		}
		has_result = false
	}
	if !has_result {
		result = this.result__(left_kw, s, is_not)
	}
	return result, nil
}

func (this *code_logic___) z__(qv *Qv___, lvl uint) (*codes___, *Errinfo___) {
	var codes *codes___
	this.reset__()
	result, err := this.z2__(this.logic, qv, lvl, 0)
	if err != nil {
		return nil, err
	}
	this.result = result
	if this.result {
		codes = this.then
	} else {
		codes = this.else1
	}
	return codes, nil
}

func (this *code_logic___) result__(left_kw *Keyword___, s string, is_not bool) bool {
	b := true
	if left_kw == nil {
		b = Bool__(s)
	} else {
		var f1, f2 float64
		var err error
		f1, err = strconv.ParseFloat(this.left_val, 0)
		if err == nil {
			f2, err = strconv.ParseFloat(s, 0)
		}
		switch left_kw {
		case Kws_.Dengyu:
			if err == nil {
				b = f1 == f2
			} else {
				b = this.left_val == s
			}
		case Kws_.Notdengyu:
			if err == nil {
				b = f1 != f2
			} else {
				b = this.left_val != s
			}
		case Kws_.Xiaoyudengyu:
			if err == nil {
				b = f1 <= f2
			} else {
				b = this.left_val <= s
			}
		case Kws_.Xiaoyu:
			if err == nil {
				b = f1 < f2
			} else {
				b = this.left_val < s
			}
		case Kws_.Dayudengyu:
			if err == nil {
				b = f1 >= f2
			} else {
				b = this.left_val >= s
			}
		case Kws_.Dayu:
			if err == nil {
				b = f1 > f2
			} else {
				b = this.left_val > s
			}
		}
	}
	if is_not {
		b = !b
	}
	if o_liucheng_ {
		o__('K', "%s %v %s %v", this.left_val, left_kw, s, b)
	}
	return b
}

func Bool__(s string) (b bool) {
	switch s {
	case "", "0", "false":
	default:
		b = true
	}
	return
}

func (this *code_logic___) codes__(i int) *codes___ {
	return this.code_1__(i).codes
}
func (this *code_logic___) code_1__(i int) *code_1___ {
	return this.logic.a[i].(*code_1___)
}

func (this *code_logic___) reset__() {
	this.result = true
	this.left_val = ""
}

func new_logic__() *code_logic___ {
	return &code_logic___{code_i___:code_i___{Kws_.If}, logic:new(codes___)}
}
