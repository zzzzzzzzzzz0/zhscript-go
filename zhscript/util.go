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

func Startswith__(r1, r2 []rune, i int) (int, bool) {
	i1 := 0
	i2 := i
	for {
		if i1 >= len(r2) {
			return i2, true
		}
		if i2 >= len(r1) {
			break
		}
		if r2[i1] != r1[i2] {
			break
		}
		i1++
		i2++
	}
	return i, false
}
