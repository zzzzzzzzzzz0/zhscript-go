package zhscript

import (
	"container/list"
)

type List___ struct {
	list.List
}
type Em___ struct {
	list.Element
}

func (this *List___) Back_i__() int {
	e := this.Back()
	this.Remove(e)
	return e.Value.(int)
}

func (this *List___) Find__(f func (*list.Element) bool) bool {
	e := this.Front()
	for e != nil {
		if f(e) {
			return true
		}
		e = e.Next()
	}
	return false
}
