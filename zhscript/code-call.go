package zhscript

import "reflect"

type code_call___ struct {
	code_z___
}

var call_t_ reflect.Value
func call_cfg__(i interface{}) {
	call_t_ = reflect.ValueOf(i)
}

var callcache_ map[string] reflect.Value = make(map[string] reflect.Value)

func call__(args *Args___, qv *Qv___, buf2 *Buf___) (goto1 *Goto___, err1 *Errinfo___) {
	m, ok := callcache_[args.Src]
	if !ok {
		m = call_t_.MethodByName(args.Src)
		ok = m.IsValid()
		if !ok {
			err1 = New_errinfo__(args.Src, Errs_.Fail)
			return
		}
		callcache_[args.Src] = m
	}

	a := []reflect.Value {}
	a = append(a, reflect.ValueOf(qv))
	for i, v := range args.A {
		if O_call_ {
			o__('n', ">%d) %v", i, v)
			O_n__()
		}
		switch v.Type {
		case "":
			a = append(a, reflect.ValueOf(v.S))
		default:
			a = append(a, reflect.ValueOf(v.I))
		}
	}
	defer func() {
		if err := recover(); err != nil {
			err1 = New_errinfo__(err, Errs_.Fail)
		}
	}()
	i2 := 0
	for i, v := range m.Call(a) {
		v2 := v.Interface()
		if O_call_ {
			o__('n', "<%d) %T %v", i, v2, v2)
			O_n__()
		}
		switch i {
		case 0:
			goto1 = v2.(*Goto___)
			continue
		case 1:
			err1 = v2.(*Errinfo___)
			continue
		}
		switch reflect.TypeOf(v2).Kind() {
		case reflect.String:
			buf2.get__(i2).WriteString(v.String())
			i2++
			continue
		case reflect.Slice:
			if s2, ok2 := v2.([]string); ok2 {
				for _, s := range s2 {
					buf2.get__(i2).WriteString(s)
					i2++
				}
				continue
			}
			if s2, ok2 := v2.([]interface{}); ok2 {
				for _, s := range s2 {
					if s3, ok3 := s.(string); ok3 {
						buf2.get__(i2).WriteString(s3)
					} else {
						buf2.set_i__(i2, s)
					}
					i2++
				}
				continue
			}
		}
		buf2.set_i__(i2, v2)
		i2++
	}
	return
}

func (this *buf_codes___) add_call__(kw *Keyword___) *codes___ {
	v := code_call___{code_z___{code_i___:code_i___{kw}, codes:new(codes___)}}
	this.add__(&v)
	return v.codes
}
