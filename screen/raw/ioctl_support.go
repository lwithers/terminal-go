package raw

/*
#include <string.h>
#include <sys/ioctl.h>
#include <termios.h>
#include <unistd.h>

struct screen_size {
	int rows, cols;
};

struct screen_size query_screen_size(void)
{
	int ret;
	struct winsize raw;
	struct screen_size sz;

	memset(&sz, 0, sizeof(sz));

	ret = ioctl(STDOUT_FILENO, TIOCGWINSZ, &raw);
	if (ret != -1) {
		sz.rows = raw.ws_row;
		sz.cols = raw.ws_col;
	}
	return sz;
}
*/
import "C"

func GetWinSize() (width, height int) {
	sz := C.query_screen_size()
	return int(sz.cols), int(sz.rows)
}
