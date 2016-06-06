# Terminal handling for Go

This repo contains packages allowing pseudo-graphical terminal control for
xterms (perhaps similar to curses, except only targetting xterms and probably
only 1% of its functionality).

## Keyboard package

Provides simple keypress recognition, for example shift+enter or left arrow.
Allows for UTF-8 rune input. Turns off screen echo.

Possible future TODO items:
 - investigate key up / key down events
 - maybe echo mode would help with readline?

## Screen package

Buffered screen output, allowing full RGB foreground/background control on a
character-by-character basis.

Possible future TODO items:
 - GUI / widget library

## Other TODO items

- readline
