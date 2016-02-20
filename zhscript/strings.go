package zhscript

type Strings___ struct {
	A []string
}

func New_strings__() *Strings___ {
	return &Strings___{[]string{}}
}

func(this *Strings___) Add__(s ...string) {
	this.A = append(this.A, s...)
}

func(this *Strings___) Len__() int {
	return len(this.A)
}

func (this *Strings___) Find__(f__ func (string) bool) bool {
	for _, s := range this.A {
		if f__(s) {
			return true
		}
	}
	return false
}
