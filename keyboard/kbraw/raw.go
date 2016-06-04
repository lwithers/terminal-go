/*
Package kbraw provides low-level keyboard access.
*/
package kbraw

/*
#include <stdlib.h>
#include <string.h>
#include <termios.h>
#include <unistd.h>

void *set_raw_mode(void)
{
	struct termios *t;
	struct termios tnew;

	t = malloc(sizeof(*t));
	if (tcgetattr(STDIN_FILENO, t)) {
		goto err_out;
	}

	memcpy(&tnew, t, sizeof(tnew));
	cfmakeraw(&tnew);

	if (tcsetattr(STDIN_FILENO, TCSANOW, &tnew)) {
		goto err_out;
	}

	return t;

err_out:
	free(t);
	return NULL;
}

void clear_raw_mode(void *tok)
{
	tcsetattr(STDIN_FILENO, TCSANOW, (const struct termios*)tok);
	free(tok);
}
*/
import "C"

import "unsafe"

var (
	rawModeRestore unsafe.Pointer
)

func RawSetup() {
	if rawModeRestore != nil {
		return
	}

	rawModeRestore = C.set_raw_mode()
	if rawModeRestore == nil {
		panic("unable to set raw mode")
	}
}

func RawTeardown() {
	if rawModeRestore == nil {
		return
	}

	C.clear_raw_mode(rawModeRestore)
	rawModeRestore = nil
}
