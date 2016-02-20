package zhscript

import "time"

type Var___ struct {
	Name string
	Val *Val___
	Is_lock bool
	Annota_val *Strings___
	Kw *Keyword___
	Timestamp int64

	Is_no_arg, Is_through bool
	Qian_arg int
	Argnames []string
}

func (this *Var___) argnames_add__(s string) {
	this.Argnames = append(this.Argnames, s)
}

func (this *Var___) use__() {
	this.Timestamp = time.Now().Unix()
}

func var__(e *Em___) *Var___ {
	return e.Value.(*Var___)
}

type Val___ struct {
	S string
	I interface{}
	Type string
}

func (this *Val___) copy_i__(val *Val___) {
	if val != nil {
		this.I = val.I
		this.Type = val.Type
	}
}

type Vars___ struct {
	ls *List___
}

func (this *Vars___) init__() {
	this.ls = new(List___)
}

func (this *Vars___) get__(name string) *Var___ {
	e := this.find__(name)
	if e != nil {
		return var__(e)
	}
	v := new(Var___)
	v.Name = name
	v.use__()
	this.ls.PushBack(v)
	return v
}

func (this *Vars___) find__(name string) (e2 *Em___) {
	this.ls.Find__(func(e *Em___) bool {
		v := var__(e)
		if v.Name == name {
			e2 = e
			v.use__()
			return true
		}
		return false
	})
	return
}

func (this *Vars___) Find__(f__ func(*Var___) bool) (ok bool) {
	this.ls.Find__(func(e *Em___) bool {
		v := var__(e)
		if f__(v) {
			ok = true
			return true
		}
		return false
	})
	return
}

func (this *Vars___) del__(e *Em___) {
	this.ls.Remove(e.Element)
}
