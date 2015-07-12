package zhscript

import (
	"strconv"
)

func (this *Qv___) z2_for__(codes *codes___, lvl uint, buf *Buf___) (*Goto___, *Errinfo___) {
	i2 := len(codes.a) - 1
	if i2 < 0 {
		for {}
		return nil, nil
	}
	var (
		name string
		has_max bool
		max int
		cnt_name, cnt_val *Buf___
	)
	{
		var (
			is_buf2 bool
			i3 int
		)
		buf2 := New_buf__()
		f := func() {
			if is_buf2 {
				if i3 == 0 {
					s2 := buf2.String()
					i := 0
					for ; i < len(s2); i++ {
						r := s2[i]
						if r >= '0' && r <= '9' {
						} else {
							break
						}
					}
					if i > 0 {
						max, _ = strconv.Atoi(s2[0:i])
						has_max = true
						s2 = s2[i:]
					}
					name = s2
				}
				i3++
				is_buf2 = false
				buf2 = New_buf__()
			}
		}
		for i := 0; i < i2; i++ {
			code := codes.a[i]
			if code.kw__() == Kws_.Dunhao {
				is_buf2 = true
				f()
				continue
			}
			goto2, err2 := this.z_code__(code, lvl, buf2)
			if err2 != nil || goto2 != nil {
				return goto2, err2
			}
			is_buf2 = true
		}
		f()
		if name != "" {
			cnt_name = New_buf__()
			cnt_name.WriteString(name)
			cnt_val = New_buf__()
		}
	}
	code := codes.a[i2]
	for i := 1;; i++ {
		if has_max && i > max {
			break
		}
		if cnt_name != nil {
			cnt_val.Reset()
			cnt_val.WriteString(strconv.Itoa(i))
			this.Set_var__(cnt_name, cnt_val, nil, Kws_.Set)
		}
		goto2, err2 := this.z_code__(code, lvl, buf)
		if err2 != nil {
			return goto2, err2
		}
		if goto2 != nil {
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
			if o_liucheng_ {
				o__('r', "goto2 (%d)%s", lvl, goto2.S)
			}
			return goto2, err2
		}
	}
	return nil, nil
}
