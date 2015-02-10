package zhscript


type buf_codes___ struct {
	codes *codes___
	kw *Keyword___
	up *buf_codes___
}

func (this *buf_codes___) add__(c code___) {
	this.codes.a = append(this.codes.a, c)
}

func (this *buf_codes___) add_1__(kw *Keyword___) {
	this.add__(&code_i___{kw})
}

func (this *buf_codes___) separ__() {
	this.add_1__(Kws_.Dunhao)
}

func (this *buf_codes___) String() (s string) {
	s += this.codes.String()
	return
}

func (this *buf_codes___) to_load__(codes *codes___) {
	id := -1
	for i, code := range this.codes.a {
		if code.kw__() == Kws_.Dunhao {
			id = i
			break
		}
	}
	if id < 0 {
		if len(this.codes.a) <= 1 {
			codes.a = append(codes.a, this.codes.a...)
		} else {
			v := kaihuakuohao__()
			codes.a = append(codes.a, v)
			v.codes.a = append(v.codes.a, this.codes.a...)
		}
	} else if id == 1 {
		codes.a = append(codes.a, this.codes.a...)
	} else {
		v := kaihuakuohao__()
		codes.a = append(codes.a, v)
		for i, code := range this.codes.a {
			if i < id {
				v.codes.a = append(v.codes.a, code)
			} else {
				codes.a = append(codes.a, code)
			}
		}
	}
	this.clear__()
}

func (this *buf_codes___) to__(codes *codes___) {
	codes.a = append(codes.a, this.codes.a...)
	this.clear__()
}

func (this *buf_codes___) clear__() {
	//this.codes.a = this.codes.a[:0]
	this.codes = new(codes___)
}

func new_buf_codes__(kw *Keyword___) *buf_codes___ {
	return &buf_codes___{codes:new(codes___), kw:kw}
}
