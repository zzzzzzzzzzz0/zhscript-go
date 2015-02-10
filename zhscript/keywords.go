package zhscript

const (
	m_juhao_ = 1 << iota
	m_yinhao_
	m_yuanyang_
	m_code_
	m_kaidanyinhao_
	m_bidanyinhao_
	m_kuohao_
	m_kaifangkuohao_
	m_bifangkuohao_
	m_huakuohao_
	m_dunhao_
	m_logic_
	m_then_
	m_equ_
	m_1_
	m_ret_lvl_
	m_all_no_see_ = m_ret_lvl_ + m_dunhao_ + m_logic_ +
		m_bidanyinhao_ + m_kaifangkuohao_ + m_bifangkuohao_
	m_all_ = 0x7fffffff - m_all_no_see_
)

type Keyword___ struct {
	s string
	s2 []rune
	mask int
}

func (this *Keyword___) String() (s string) {
	if this != nil {
		s = this.s
	}
	return
}

func (this *Keyword___) is__(code []rune, i int) (int, bool) {
	return Startswith__(code, this.s2, i)
}

func (this *Keyword___) is2__(i int) bool {
	return this.mask & i != 0
}

var Kws_ = struct {
	Juhao,
	Kaiyinhao, Biyinhao,
	Kaidanyinhao, Bidanyinhao,
	Kaikuohao, Bikuohao,
	Kaifangkuohao, Bifangkuohao,
	Kaihuakuohao, Bihuakuohao,
	Dunhao, Maohao,
	
	Begin_code, End_code,
	Begin_yuanyang, End_yuanyang,
	
	Dengyu, Notdengyu, Xiaoyudengyu, Xiaoyu, Dayudengyu, Dayu,
	Not, And, Or,
	If, Then, Else,

	For, Break, Continue,
	Quit, Return,

	Set, Top, Up, Lock, Equ, Del, Has,
	Alias,
	Def, Noarg,

	Arg, Args, Length,
	CR, LF,

	Eval, Interp, Load,
	
	Call,

	Echo, Exec *Keyword___
} {
	new_kw__("。", m_juhao_),
	new_kw__("“", m_yinhao_),
	new_kw__("”", m_yinhao_),
	new_kw__("‘", m_kaidanyinhao_),
	new_kw__("’", m_bidanyinhao_ | m_ret_lvl_),
	new_kw__("（", m_kuohao_),
	new_kw__("）", m_kuohao_),
	new_kw__("【", m_kaifangkuohao_),
	new_kw__("】", m_bifangkuohao_ | m_ret_lvl_),
	new_kw__("先", m_1_),
	new_kw__("了", m_1_ | m_ret_lvl_),
	new_kw__("、", m_dunhao_ | m_ret_lvl_),
	new_kw__("：", m_dunhao_ | m_ret_lvl_),

	new_kw__("下代码", m_code_),
	new_kw__("上代码", m_code_),
	new_kw__("下原样", m_yuanyang_),
	new_kw__("上原样", m_yuanyang_),
	
	new_kw__("等于",	m_logic_),
	new_kw__("不等于",	m_logic_),
	new_kw__("小于等于",m_logic_),
	new_kw__("小于",	m_logic_),
	new_kw__("大于等于",m_logic_),
	new_kw__("大于",	m_logic_),
	new_kw__("不",		m_logic_),
	new_kw__("并且",	m_logic_),
	new_kw__("或者",	m_logic_),
	new_kw__("如果", m_1_),
	new_kw__("那么", m_then_ | m_ret_lvl_),
	new_kw__("否则", m_then_ | m_ret_lvl_),

	new_kw__("循环", m_1_),
	new_kw__("跳出", m_1_),
	new_kw__("再来", m_1_),

	new_kw__("结束", m_1_),
	new_kw__("退出", m_1_),

	new_kw__("赋予", m_1_),
	new_kw__("顶", 0),
	new_kw__("上", 0),
	new_kw__("锁", 0),
	new_kw__("以", m_equ_ | m_ret_lvl_),
	new_kw__("删除", m_1_),
	new_kw__("存在", m_1_),
	new_kw__("别名", m_1_),
	new_kw__("定义", m_1_),
	new_kw__("无参", 0),

	new_kw__("参数", 0),
	new_kw__("参数栈", 0),
	new_kw__("数目", 0),

	new_kw__("回车", m_1_),
	new_kw__("换行", m_1_),

	new_kw__("算术", m_1_),
	new_kw__("解释", m_1_),
	new_kw__("加载", m_1_),
	
	new_kw__("调用", m_1_),
	
	new_kw__("显示", m_1_),
	new_kw__("执行", m_1_),
}

var all_kw_ []*Keyword___

func new_kw__(s string, mask int) *Keyword___ {
	kw := new(Keyword___)
	kw.s = s
	kw.s2 = []rune(s)
	kw.mask = mask

	//可“遇”关键字调整到最前
	if false {
		/*all_kw := make([]*Keyword___, 0)
		all_kw = append(all_kw, kw)
		all_kw = append(all_kw, all_kw_...)*/
		all_kw := make([]*Keyword___, 1 + len(all_kw_))
		all_kw[0] = kw
		copy(all_kw[1:], all_kw_)
		all_kw_ = all_kw
	} else {
		all_kw_ = append(all_kw_, kw)
	}
	
	return kw
}

func is_kw__(code []rune, i int, mask int) (int, *Keyword___, bool) {
	for _, kw := range all_kw_ {
		if kw.mask & mask != 0 {
			i1, ok := kw.is__(code, i)
			if ok {
				return i1, kw, true
			}
		}
	}
	return i, nil, false
}

func is_kw2__(code []rune, i int, kws ...*Keyword___) (int, *Keyword___, bool) {
	for _, kw := range kws {
		i1, ok := kw.is__(code, i)
		if ok {
			return i1, kw, true
		}
	}
	return i, nil, false
}
