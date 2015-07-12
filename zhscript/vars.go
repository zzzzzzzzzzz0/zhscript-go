package zhscript

type Var___ struct {
	Name string
	Val *Val___
	Is_lock bool
	Annota_val *Strings___
	Kw *Keyword___

	Is_no_arg, Is_through bool
	Qian_arg int
	Argnames []string
}

func (this *Var___) argnames_add__(s string) {
	this.Argnames = append(this.Argnames, s)
}

func Var__(e *Em___) *Var___ {
	return e.Value.(*Var___)
}

type Val___ struct {
	S string
	I interface{}
	Type rune
}

func (this *Val___) copy_i__(val *Val___) {
	if val != nil {
		this.I = val.I
		this.Type = val.Type
	}
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
		v := Var__(e)
		if v.Name == name {
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
