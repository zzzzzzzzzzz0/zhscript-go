package zhscript

import (
	"container/list"
	"fmt"
)

type List___ struct {
	list.List
}
type Em___ struct {
	*list.Element
}

func (this *List___) Back_i__() int {
	e := this.Back()
	this.Remove(e)
	return e.Value.(int)
}

func (this *List___) Find__(f__ func (*Em___) bool) bool {
	var e2 Em___
	for e := this.Front(); e != nil; e = e.Next() {
		e2.Element = e
		if f__(&e2) {
			return true
		}
	}
	return false
}

func (this *List___) o__(r rune, s string) {
	i := 0
	this.Find__(func(e *Em___) bool {
		s += fmt.Sprintf(" %d)%v", i, e.Value)
		i++
		return false
	})
	o__(r, "%s", s)
}

func (this *Em___) String() (s string) {
	if s1, ok := this.Value.(string); ok {
		s = s1
	}
	return
}