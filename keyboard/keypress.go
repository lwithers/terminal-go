package keyboard

import "fmt"

// KeyPress represents a single keypress. If the value is > 0 then it
// represents a rune.
type KeyPress int32

const (
	Key_Escape    = KeyPress(-1)
	Key_Enter     = KeyPress(-2)
	Key_Backspace = KeyPress(-3)
)

func (k KeyPress) IsRune() bool {
	return k > 0
}

func (k KeyPress) String() string {
	if k > 0 {
		return string([]rune{rune(k)})
	}

	name, ok := KeyNames[k]
	if !ok {
		return fmt.Sprintf("[keypress %d]", k)
	}
	return name
}
