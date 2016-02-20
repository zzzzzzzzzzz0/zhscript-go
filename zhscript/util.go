package zhscript

func Fields__(s string) []string {
	buf := ""
	ret := []string {}
	var has, in_s1, in_s2 bool
	for _, r := range s {
		switch r {
		case '"':
			if !in_s2 {
				in_s1 = !in_s1
				continue
			}
		case '\'':
			if !in_s1 {
				in_s2 = !in_s2
				continue
			}
		}
		if in_s1 || in_s2 {
		} else {
			switch r {
			case ' ':
				ret = append(ret, buf)
				buf = ""
				has = false
				continue
			}
		}
		buf += string(r)
		has = true
	}
	if has {
		ret = append(ret, buf)
	}
	return ret
}

func Startswith2__(r1, r2 []rune, i int, skip_blank bool) (int, bool) {
	i1 := i
	i2 := 0
	for {
		if i2 >= len(r2) {
			return i1, true
		}
		if i1 >= len(r1) {
			break
		}
		r3 := r1[i1]
		r4 := r2[i2]
		i1++
		i2++
		if skip_blank {
			if r3 >= 0 && r3 <= ' ' {
				continue
			}
			if r4 >= 0 && r4 <= ' ' {
				continue
			}
		}
		if r3 != r4 {
			break
		}
	}
	return i, false
}

func Startswith__(r1, r2 []rune, i int) (int, bool) {
	return Startswith2__(r1, r2, i, false)
}
