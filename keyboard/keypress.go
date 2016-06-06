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

// IsRune returns true if k represents a Unicode rune. This is true for all
// positive values of k.
func (k KeyPress) IsRune() bool {
	return k > 0
}

// String returns the string representation of the keypress: either the
// Unicode rune it represents, or the key's name from KeyNames.
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
