package main

import (
	"fmt"
	"os"
	. "./zhscript"
)

type l___ struct {}

//仅为测试“调用”
func (this *l___) Substr(qv *Qv___, s ...string) (goto1 *Goto___, err1 *Errinfo___, s1, s2 string) {
	var start, stop int
	fmt.Sscan(s[1], &start)
	fmt.Sscan(s[2], &stop)
	s1 = s[0][start:stop]
	s2 = fmt.Sprint(len(s1))
	return
}

func main() {
	z, err := New__(os.Args, nil, &l___{})
	if err == nil {
		_, err = z.Z__()
	}
	if err != nil {
		fmt.Println()
		fmt.Println(err)
	}
}
