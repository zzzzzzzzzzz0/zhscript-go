package zhscript


type buf_codes___ struct {
	codes *codes___
	up *buf_codes___
	kw10 *Keyword___
	names []string
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
			this.to2__(codes)
		} else {
			v := kaihuakuohao__()
			codes.a = append(codes.a, v)
			this.to2__(v.codes)
		}
	} else if id == 1 {
		this.to2__(codes)
	} else {
		v := kaihuakuohao__()
		codes.a = append(codes.a, v)
		this.to3__(codes, v.codes, id)
	}
	this.clear__()
}

func (this *buf_codes___) to_logic__(codes *codes___) {
	this.to2__(codes)
	this.clear__()
}

func (this *buf_codes___) to3__(codes, codes2 *codes___, id int) {
	for i, code := range this.codes.a {
		switch code.kw__() {
		case Kws_.Juhao:
			continue
		}
		if codes2 != nil && i < id {
			codes2.a = append(codes2.a, code)
			continue
		}
		codes.a = append(codes.a, code)
	}
}

func (this *buf_codes___) to2__(codes *codes___) {
	this.to3__(codes, nil, 0)
}

func (this *buf_codes___) to__(codes *codes___) {
	this.to2__(codes)
	this.clear__()
}

func (this *buf_codes___) clear__() {
	//this.codes.a = this.codes.a[:0]
	this.codes = new(codes___)
}

func new_buf_codes__(kw *Keyword___) *buf_codes___ {
	return &buf_codes___{codes:new(codes___), kw10:kw}
}
