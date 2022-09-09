package fluidsynth

// #include <stdlib.h>
import "C"

func cBool(b bool) C.int {
	if b {
		return 1
	}
	return 0
}
