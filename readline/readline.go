/*
Package readline implements a readline(3)-style editable command line interface
with programmable tab completion, history and syntax highlighting.
*/
package readline

import "github.com/lwithers/terminal-go/keyboard"

type TabCompleter interface {
	NextWord(words []string) []string
	CompleteWord(words []string, partial string) []string
}

type Historian interface {
	Last(offset int) string
	LastMatch(partial string) string
	Search(substr string) string
}

type SyntaxHighlighter interface {
	Colours(words []string, rgbOut []int32)
}

func Readline(k chan keyboard.KeyPress, t TabCompleter, h Historian,
	s SyntaxHighlighter) string {
	var line []rune
	for kp := range k {
		if kp.IsRune() {
			line = append(line, rune(kp))
			continue
		}
		switch kp {
		case keyboard.Key_Enter:
			return string(line)
		case keyboard.Key_Backspace:
			if len(line) > 0 {
				line = line[:len(line)-1]
			}
		}
	}
	return ""
}
