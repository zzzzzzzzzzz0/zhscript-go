package zhscript

type Var___ struct {
	Name, Val string
	Is_lock, Is_no_arg bool
	Annota, Annota_val *List___
	Kw *Keyword___
}

func var__(e *Em___) *Var___ {
	return e.Value.(*Var___)
}

type vars___ struct {
	ls *List___
}

func (this *vars___) init__() {
	this.ls = new(List___)
}

func (this *vars___) get__(name string, annota *List___) *Var___ {
	e := this.find__(name, annota)
	if e != nil {
		return var__(e)
	}
	v := new(Var___)
	v.Name = name
	v.Annota = annota
	this.ls.PushBack(v)
	return v
}

func (this *vars___) find__(name string, annota *List___) *Em___ {
	var e2 *Em___
	this.ls.Find__(func(e *Em___) bool {
		v := var__(e)
		if v.Name == name {
			for ea2 := annota.Front(); ea2 != nil; ea2 = ea2.Next() {
				b := true
				for ea := v.Annota.Front(); ea != nil; ea = ea.Next() {
					if ea.Value.(string) == ea2.Value.(string) {
						b = false
						break
					}
				}
				if b {
					return false
				}
			}
			e2 = e
			return true
		}
		return false
	})
	return e2
}

func (this *vars___) del__(e *Em___) {
	this.ls.Remove(e.Element)
}

func For_var__(qv *Qv___, f func(v *Var___, i int) bool) {
	if f == nil {
		return
	}
	var i int
	for {
		if qv == nil {
			break
		}
		qv.vars.ls.Find__(func(e *Em___) bool {
			return f(var__(e), i)
		});
		qv = qv.up
		i++
	}
}