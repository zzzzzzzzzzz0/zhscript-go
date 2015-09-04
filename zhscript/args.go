package zhscript

import "strings"

var Shebang_flag_ = "--" + Kws_.Juhao.s
const (
	Src_is_code_ = iota
	Src_is_file_
	Src_is_varname_
)

type Args___ struct {
	Src, Src2 string
	Src_type int
	A []*Val___
	Names []string

	is_mk_all bool
	all string
}

func (this *Args___) Reset__() {
	this.A = []*Val___ {}
	this.is_mk_all = false
	this.all = ""
}

func (this *Args___) Src_file__(s string) {
	this.Src = s
	this.Src_type = Src_is_file_
}

func (this *Args___) Src_code__(s string) {
	this.Src = s
	this.Src_type = Src_is_code_
}

func (this *Args___) Add__(s ...string) {
	for _, s2 := range s {
		this.A = append(this.A, &Val___{S:s2})
	}
}

func (this *Args___) Add2__(v *Val___) {
	this.A = append(this.A, v)
}

func (this *Args___) A__() (a []string) {
	a = []string {}
	for _, v := range this.A {
		a = append(a, v.S)
	}
	return
}

func (this *Args___) all__() string {
	if !this.is_mk_all {
		this.is_mk_all = true
		for i, v := range this.A {
			s := v.S
			if i > 0 {
				this.all += " "
			}
			if s == "" {
				this.all += `""`
				continue
			}
			b := false
			for _, r := range s {
				if !(r >= '0' && r <= '9' || r >= 'a' && r <= 'z' || r >= 'A' && r <= 'Z' || r == '.' || r == '-' || r == '_') {
					b = true
					break
				}
			}
			if b {
				this.all += "\""
				s = strings.Replace(s, "\"", "\\\"", -1)
			}
			this.all += s
			if b {
				this.all += "\""
			}
		}
	}
	return this.all
}

var o_liucheng_, o_liucheng2_, o_ansi_, o_tree_, o_path_, o_args_ bool

func (this *Args___) Parse__(a []string, from int, lvl uint) {
	all_is := false
	for i, s := range a {
		if o_args_ {
			var i1 uint = 0
			for ; i1 < lvl; i1++ {
				o_t__()
			}
			o__('n', "%d) %s", i, s)
			o_n__()
		}
		if i < from {
			continue
		}
		if !all_is {
			if s == "-zhscript-o-lc" {
				o_liucheng_ = true
				continue
			}
			if s == "-zhscript-o-lc2" {
				o_liucheng2_ = true
				continue
			}
			if s == "-zhscript-o-ansi" {
				o_ansi_ = true
				continue
			}
			if s == "-zhscript-o-tree" {
				o_tree_ = true
				continue
			}
			if s == "-zhscript-o-path" {
				o_path_ = true
				continue
			}
			if s == "-zhscript-o-args" {
				o_args_ = true
				continue
			}
			if i == from && strings.HasSuffix(s, " " + Shebang_flag_) {
				this.Parse__(Fields__(s), 0, lvl + 1)
				continue
			}
			if s == Shebang_flag_ {
				continue
			}
			if lvl == 0 && this.Src == "" && !strings.HasPrefix(s, "-") {
				this.Src_file__(s)
				continue
			}
			if s == "----" {
				all_is = true
				continue
			}
		}
		this.Add__(s)
	}
}
