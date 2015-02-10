package zhscript

func (this *codes___) z__(code []rune, up_qv *Qv___, from, to, end_mask, mask, lvl int, thiz *buf_codes___) (int, *Keyword___, *Errinfo___) {
	if o_liucheng2_ {
		o_n__()
		o__('n', "(%d)%d-%d,%b", lvl, from, to, end_mask)
		//o__('n', " %p/%p", thiz.codes, this)
		o__(0, "%v", string(code[from:to]))
		o_n__()
	}
	var (
		i = from
		yinhaonei, kuohaonei, yuanyangnei, code_nei int
		kw *Keyword___
		ok bool
	)
	buf := New_buf__()
	mask_bak := new(List___)

	for i < to {
		codei := code[i]
		if codei >= 0 && codei <= ' ' {
			if yuanyangnei == 0 {
				if yinhaonei == 0 {
					if code_nei == 0 {
						i++
						continue
					}
				}
				switch codei {
				case ' ':
				default:
					i++
					continue
				}
			}
		}
		i, kw, ok = is_kw__(code, i, mask)
		if ok {
			if o_liucheng2_ {
				o__('k', "%s", kw)
			}
			switch kw {
			case Kws_.Kaiyinhao:
				if yinhaonei == 0 {
					mask_bak.PushBack(mask)
					mask = m_yinhao_ | m_kaidanyinhao_
				} else {
					buf.WriteString(kw.s)
				}
				yinhaonei++
				continue
			case Kws_.Biyinhao:
				yinhaonei--
				if yinhaonei == 0 {
					mask = mask_bak.Back_i__()
				} else {
					buf.WriteString(kw.s)
				}
				continue
			case Kws_.Begin_yuanyang:
				if yuanyangnei == 0 {
					if code_nei != 0 {
						buf.WriteString(kw.s)
					}
					mask_bak.PushBack(mask)
					mask = m_yuanyang_
				} else {
					buf.WriteString(kw.s)
				}
				yuanyangnei++
				continue
			case Kws_.End_yuanyang:
				yuanyangnei--
				if yuanyangnei == 0 {
					mask = mask_bak.Back_i__()
					if code_nei != 0 {
						buf.WriteString(kw.s)
					}
				} else {
					buf.WriteString(kw.s)
				}
				continue
			case Kws_.Begin_code:
				if code_nei == 0 {
					mask_bak.PushBack(mask)
					mask = m_code_ | m_yuanyang_ | m_kuohao_
				} else {
					buf.WriteString(kw.s)
				}
				code_nei++
				continue
			case Kws_.End_code:
				code_nei--
				if code_nei == 0 {
					mask = mask_bak.Back_i__()
				} else {
					buf.WriteString(kw.s)
				}
				continue
			case Kws_.Kaikuohao:
				if kuohaonei == 0 {
					mask_bak.PushBack(mask)
					mask = m_kuohao_
				}
				kuohaonei++
				continue
			case Kws_.Bikuohao:
				kuohaonei--
				if kuohaonei == 0 {
					mask = mask_bak.Back_i__()
				}
				continue
			case Kws_.LF:
				buf.WriteRune('\n')
				continue
			case Kws_.CR:
				buf.WriteRune('\r')
				continue
			}
			thiz.add_text__(buf)
			switch {
			case kw.is2__(m_ret_lvl_):
				if lvl > 0 {
					if o_liucheng2_ {
						o__('r', "lvl %s", kw)
					}
					return i, kw, nil
				}
			case kw.is2__(m_logic_):
				if lvl > 0 {
					if o_liucheng2_ {
						o__('r', "if %s", kw)
					}
					return i, kw, nil
				}
			}
		}
		if kuohaonei != 0 {
			if !ok {
				i++
			}
			continue
		}
		var thiz2 *buf_codes___
		if !ok {
			if yinhaonei == 0 && yuanyangnei == 0 && code_nei == 0 && (end_mask & m_bidanyinhao_ == 0) {
				var is_no_arg bool
				i, kw, thiz2, is_no_arg, ok = is_def__(code, up_qv, i, thiz)
				if ok {
					thiz.add_text__(buf)
					if is_no_arg {
						thiz2.to_load__(thiz.add_load__(kw, thiz2.kw))
						continue
					}
				}
			}
		}
		if ok {
			switch {
			case kw.is2__(m_juhao_):
				break
			default:
				end_mask2 := m_juhao_
				if thiz2 == nil {
					thiz2 = new_buf_codes__(kw)
				}
				from2 := i
				var (
					i2 int
					kw2 *Keyword___
					err2 *Errinfo___
					if2 *code_logic___
					var2 *code_var___
				)
				switch kw {
				case Kws_.Kaihuakuohao:
					end_mask2 = m_huakuohao_
				case Kws_.Kaidanyinhao:
					end_mask2 |= m_bidanyinhao_ | m_kaifangkuohao_
				case Kws_.Kaifangkuohao:
					end_mask2 = m_bifangkuohao_
				case Kws_.If:
					end_mask2 |= m_logic_ | m_then_
					if2 = new_logic__()
				case Kws_.Has:
					end_mask2 |= m_kaifangkuohao_
				case Kws_.Set, Kws_.Alias, Kws_.Def:
					end_mask2 |= m_equ_ | m_kaifangkuohao_ | m_dunhao_
					var2 = new_var__(kw)
				case Kws_.Call, Kws_.Load, Kws_.Interp:
					end_mask2 |= m_dunhao_
				}
				thiz2.up = thiz
				for {
					i2, kw2, err2 = this.z__(code, up_qv, from2, to, end_mask2, mask | end_mask2, lvl + 1, thiz2)
					if err2 != nil {
						if o_liucheng2_ {
							o__('r', "err2 (%d)%s", lvl, err2)
						}
						return from2, kw2, err2
					}
					if o_liucheng2_ {
						o__('K', "(%d)%d %s %s", lvl, i2, kw, kw2)
						//o__('K', " %p/%p", thiz2.codes, thiz.codes)
					}
					b := false
					switch kw2 {
					case nil, Kws_.Kaihuakuohao, Kws_.Kaifangkuohao:
					case Kws_.Dunhao, Kws_.Maohao:
						thiz2.separ__()
					default:
						switch {
						case kw2.is2__(m_logic_):
							thiz2.add_1__(kw2)
						case kw2 == Kws_.Then || kw == Kws_.Then || kw2 == Kws_.Else || kw == Kws_.Else:
							if if2 == nil {
								if end_mask & m_then_ != 0 {
									err3 := this.z3__(thiz, thiz2, if2, var2, kw, i, i2, code)
									if err3 != nil {
										return i, kw2, err3
									}
									if o_liucheng2_ {
										o__('r', "m_then %s", kw2)
									}
									return i2, kw2, nil
								}
								return i, kw2, new_errinfo__(Errs_.Keyword, "", code, i, i2, Kws_.If, nil)
							}
							switch thiz2.kw {
							case Kws_.Then:
								thiz2.to__(if2.then)
							case Kws_.Else:
								thiz2.to__(if2.else1)
							default:
								thiz2.to__(if2.logic)
							}
							is_then := kw2 == Kws_.Then || kw == Kws_.Then
							end_mask2 &= ^m_logic_
							if is_then {
								if if2.then == nil {
									if2.then = new(codes___)
								}
								thiz2.kw = Kws_.Then
							} else {
								if if2.else1 == nil {
									if2.else1 = new(codes___)
								}
								thiz2.kw = Kws_.Else
							}
						case kw2 == Kws_.Equ || kw == Kws_.Equ:
							if var2 == nil {
								if end_mask & m_equ_ != 0 {
									if o_liucheng2_ {
										o__('r', "m_equ %s", kw2)
									}
									return i2, kw2, nil
								}
								return i, kw2, new_errinfo__(Errs_.Keyword, "", code, i, i2, Kws_.Set, nil)
							}
							thiz2.to__(var2.name)
							thiz2.kw = Kws_.Equ
							//end_mask2 &= ^m_kaifangkuohao_
							if var2.val == nil {
								var2.val = new(codes___)
							}
						default:
							b = true
						}
					}
					if from2 == i2 {
						break
					}
					from2 = i2
					if b {
						break
					}
				}
				err3 := this.z3__(thiz, thiz2, if2, var2, kw, i, i2, code)
				if err3 != nil {
					return i, kw2, err3
				}
				switch kw2 {
				case Kws_.Bihuakuohao:
					if kw != Kws_.Kaihuakuohao && (end_mask & m_huakuohao_ != 0) {
						if o_liucheng2_ {
							o__('r', "bihua %s", kw)
						}
						return i2, kw2, nil
					}
				}
				i = i2
			}
			if end_mask & kw.mask != 0 {
				if o_liucheng2_ {
					o__('e', "%s", kw)
				}
				switch kw {
				case Kws_.Juhao:
					i -= len(kw.s2)
				}
				return i, kw, nil
			}
			continue
		}
		buf.WriteRune(codei)
		i++
	}
	thiz.add_text__(buf)
	if o_liucheng2_ {
		o__('r', "e")
		o_n__()
	}
	return i, nil, nil
}

func (this *codes___) z3__(thiz, thiz2 *buf_codes___, if2 *code_logic___, var2 *code_var___, kw *Keyword___, i, i2 int, code []rune) (*Errinfo___) {
	if o_liucheng2_ {
		o__('x', "%s", kw)
	}
	switch kw {
	case Kws_.If:
		switch thiz2.kw {
		case Kws_.Then:
			thiz2.to__(if2.then)
		case Kws_.Else:
			thiz2.to__(if2.else1)
		default:
			return new_errinfo__(Errs_.Keyword, "", code, i, i2, Kws_.Then, Kws_.Else)
		}
		thiz.add__(if2)
	case Kws_.Set, Kws_.Alias, Kws_.Def:
		if var2.val == nil {
			return new_errinfo__(Errs_.Keyword, var2.name.String(), code, i, i2, Kws_.Equ, nil)
		}
		thiz2.to__(var2.val)
		thiz.add__(var2)
	case Kws_.Interp, Kws_.Load:
		thiz2.to_load__(thiz.add_load__(kw, thiz2.kw))
	case Kws_.Call:
		thiz2.to_load__(thiz.add_call__(kw))
	default:
		var kw2 *Keyword___
		switch kw {
		case Kws_.Kaidanyinhao:
			kw2 = Kws_.Bidanyinhao
		case Kws_.Kaifangkuohao:
			kw2 = Kws_.Bifangkuohao
		case Kws_.Kaihuakuohao:
			kw2 = Kws_.Bihuakuohao
		}
		thiz2.to__(thiz.add_2__(kw, kw2))
	}
	return nil
}

func (this *codes___) z2__(code []rune, up_qv *Qv___) *Errinfo___ {
	buf := new_buf_codes__(nil)
	_, _, err := this.z__(code, up_qv, 0, len(code), 0, m_all_, 0, buf)
	if err == nil {
		buf.to__(this)
		if o_tree_ {
			for_o_codes__(this, 0)
		}
	}
	return err
}
