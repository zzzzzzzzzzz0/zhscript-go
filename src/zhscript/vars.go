package zhscript

import (
	"container/list"
)

type var___ struct {
	name, val string
	is_lock bool
	kw *Keyword___
}

func var__(e *list.Element) *var___ {
	return e.Value.(*var___)
}

type vars___ struct {
	ls *List___
}

func (this *vars___) init__() {
	this.ls = new(List___)
}

func (this *vars___) get__(name string) *var___ {
	e := this.find__(name)
	if e != nil {
		return var__(e)
	}
	v := new(var___)
	v.name = name
	this.ls.PushBack(v)
	return v
}

func (this *vars___) find__(name string) *list.Element {
	var e2 *list.Element
	this.ls.Find__(func(e *list.Element) bool {
		v := var__(e)
		if v.name == name {
			e2 = e
			return true
		}
		return false
	})
	return e2
}

func (this *vars___) del__(e *list.Element) {
	this.ls.Remove(e)
}