package zhscript

type Err___ struct {
	s string
}

type Errs___ struct {
	Openfile, Exist, Lock, Annota, Keyword, Fail, Case, Defnamenil *Err___
}

var Errs_ = Errs___ {
	&Err___{"无法打开文件"},
	&Err___{"变量不存在"},
	&Err___{"变量被锁"},
	&Err___{"注解有问题"},
	&Err___{"期待的关键字"},
	&Err___{"调用失败"},
	&Err___{"不合适"},
	&Err___{"定义名不能为空"},
}

type Errinfo___ struct {
	err *Err___
	s string
}

/*func (this *Errinfo___) String() string {
	return this.s
}*/

func (this *Errinfo___) Error() string {
	return this.s
}

func (this *Errinfo___) Add__(i interface{}) {
	if s, ok := i.(string); ok && s != "" {
		this.s += Kws_.Kaiyinhao.s + s + Kws_.Biyinhao.s
		return
	}
	if err, ok := i.(*Err___); ok {
		this.err = err
		this.s += err.s
		return
	}
	if kw, ok := i.(*Keyword___); ok && kw != nil {
		this.s += Kws_.Kaiyinhao.s + kw.s + Kws_.Biyinhao.s
		return
	}
}

func New_errinfo__(a ...interface{}) *Errinfo___ {
	ei := new(Errinfo___)
	ei.s = ""
	for _, i := range a {
		ei.Add__(i)
	}
	return ei
}

func new_errinfo__(err *Err___, s string, code []rune, i, to2 int, kw, kw2 *Keyword___) *Errinfo___ {
	var code2 string
	for _, r := range code[i:to2] {
		code2 += string(r)
	}
	return New_errinfo__(s, code2, err, kw, kw2)
}