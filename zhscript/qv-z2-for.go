package zhscript

func (this *Qv___) z2_for__(codes *codes___, lvl uint, buf *Buf___) (*Goto___, *Errinfo___) {
	i2 := len(codes.a) - 1
	var name string
	for i := 0; i <= i2; {
		buf2 := New_buf__()
		goto2, err2 := this.z_code__(codes.a[i], lvl, buf2)
		if err2 != nil {
			return goto2, err2
		}
		if i == 0 && i2 > 0 {
			name = buf2.String()
		} else {
			buf.Write(buf2.Bytes())
		}
		if i == i2 {
			if goto2 == nil {
				continue
			}
			if o_liucheng_ {
				o__('r', "goto2 %s %s", name, goto2.S)
			}
			if goto2.S == "" || name == goto2.S {
				if goto2.Kw == Kws_.Continue {
					continue
				}
				if goto2.Kw == Kws_.Break {
					break
				}
			}
		}
		if goto2 != nil {
			if o_liucheng_ {
				o__('r', "goto2 (%d)%s", lvl, goto2.S)
			}
			return goto2, err2
		}
		i++
	}
	return nil, nil
}
