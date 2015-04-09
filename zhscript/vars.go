package zhscript

type Var___ struct {
	Name, Val string
	Is_lock, Is_no_arg bool
	Qian_arg int
	Annota_val *List___
	Kw *Keyword___
}

func Var__(e *Em___) *Var___ {
	return e.Value.(*Var___)
}

type Vars___ struct {
	Ls *List___
}

func (this *Vars___) init__() {
	this.Ls = new(List___)
}

func (this *Vars___) get__(name string) *Var___ {
	e := this.find__(name)
	if e != nil {
		return Var__(e)
	}
	v := new(Var___)
	v.Name = name
	this.Ls.PushBack(v)
	return v
}

func (this *Vars___) find__(name string) *Em___ {
	var e2 *Em___
	this.Ls.Find__(func(e *Em___) bool {
		if Var__(e).Name == name {
			e2 = e
			return true
		}
		return false
	})
	return e2
}

func (this *Vars___) del__(e *Em___) {
	this.Ls.Remove(e.Element)
}
