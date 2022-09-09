package fluidsynth

// #include <fluidsynth.h>
import "C"

func Version() string {
	return C.GoString(C.fluid_version_str())
}
