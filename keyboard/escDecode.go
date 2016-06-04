package keyboard

type tree struct {
	kp KeyPress
	c  map[uint8]*tree
}

func decodeEscapeSequence(buf []byte) (KeyPress, int) {
	p := seqMap

	for n := 1; n < len(buf); n++ {
		if p.kp != 0 {
			return p.kp, n
		}
		p = p.c[buf[n]]
		if p == nil {
			return Key_Escape, 1
		}
	}

	if p.kp != 0 {
		return p.kp, len(buf)
	}

	return 0, 0
}
