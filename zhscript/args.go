package zhscript

import (
	"strings"
	"os"
)

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

var O_ansi_, O_liucheng_, O_liucheng2_, O_tree_, O_path_, O_args_, O_call_ bool

type o_set___ struct {
	tag string
	b *bool
}
var o_set_ = []o_set___ {
	{"-zhscript-o-ansi", &O_ansi_},
	{"-zhscript-o-lc", &O_liucheng_},
	{"-zhscript-o-lc2", &O_liucheng2_},
	{"-zhscript-o-tree", &O_tree_},
	{"-zhscript-o-path", &O_path_},
	{"-zhscript-o-args", &O_args_},
	{"-zhscript-o-call", &O_call_},
}

func (this *Args___) Parse__(a []string, from int, lvl uint) {
	all_is := false
	for1:
	for i, s := range a {
		if O_args_ {
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
			for _, o_set := range o_set_ {
				if s == o_set.tag {
					*o_set.b = true
					continue for1
				}
			}
			if s == "-zhscript-help" {
				for _, o_set := range o_set_ {
					println(o_set.tag)
				}
				os.Exit(250)
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
