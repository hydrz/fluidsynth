package fluidsynth

// #include <fluidsynth.h>
// #include <stdlib.h>
import "C"
import "unsafe"

type FluidType int

const (
	FLUID_NO_TYPE  FluidType = C.FLUID_NO_TYPE
	FLUID_NUM_TYPE FluidType = C.FLUID_NUM_TYPE
	FLUID_INT_TYPE FluidType = C.FLUID_INT_TYPE
	FLUID_STR_TYPE FluidType = C.FLUID_STR_TYPE
	FLUID_SET_TYPE FluidType = C.FLUID_SET_TYPE
)

type Settings struct {
	ptr *C.fluid_settings_t
}

func NewSettings() *Settings {
	cPtr := C.new_fluid_settings()

	if cPtr == nil {
		return nil
	}

	return &Settings{ptr: cPtr}
}

func (s *Settings) Delete() {
	C.delete_fluid_settings(s.ptr)
}

func (s *Settings) GetType(name string) FluidType {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	switch C.fluid_settings_get_type(s.ptr, cname) {
	case C.FLUID_NO_TYPE:
		return FLUID_NO_TYPE
	case C.FLUID_NUM_TYPE:
		return FLUID_NUM_TYPE
	case C.FLUID_INT_TYPE:
		return FLUID_INT_TYPE
	case C.FLUID_STR_TYPE:
		return FLUID_STR_TYPE
	case C.FLUID_SET_TYPE:
		return FLUID_SET_TYPE
	default:
		return FLUID_NO_TYPE
	}
}

func (s *Settings) GetHints(name string, val int) bool {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cVal := C.int(val)
	return C.fluid_settings_get_hints(s.ptr, cname, &cVal) == FLUID_OK
}

func (s *Settings) IsRealtime(name string) bool {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return C.fluid_settings_is_realtime(s.ptr, cname) != 0
}

func (s *Settings) Setstr(name string, str string) bool {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return C.fluid_settings_setstr(s.ptr, cname, C.CString(str)) == FLUID_OK
}

func (s *Settings) Copystr(name string, str string, len int) bool {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return C.fluid_settings_copystr(s.ptr, cname, C.CString(str), C.int(len)) == FLUID_OK
}

func (s *Settings) Dupstr(name string, str string) bool {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cStr := C.CString(str)
	defer C.free(unsafe.Pointer(cStr))
	return C.fluid_settings_dupstr(s.ptr, cname, &cStr) == FLUID_OK
}

func (s *Settings) GetstrDefault(name string, def string) bool {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cDef := C.CString(def)
	defer C.free(unsafe.Pointer(cDef))
	return C.fluid_settings_getstr_default(s.ptr, cname, &cDef) == FLUID_OK
}

func (s *Settings) StrEqual(name string, value string) bool {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return C.fluid_settings_str_equal(s.ptr, cname, C.CString(value)) == FLUID_OK
}

func (s *Settings) Setnum(name string, val float64) bool {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return C.fluid_settings_setnum(s.ptr, cname, C.double(val)) == FLUID_OK
}

func (s *Settings) Getnum(name string, val float64) bool {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cVal := C.double(val)
	return C.fluid_settings_getnum(s.ptr, cname, &cVal) == FLUID_OK
}

func (s *Settings) GetnumDefault(name string, val float64) bool {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cVal := C.double(val)
	return C.fluid_settings_getnum_default(s.ptr, cname, &cVal) == FLUID_OK
}

func (s *Settings) GetnumRange(name string,
	min float64, max float64) bool {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cMin := C.double(min)
	cMax := C.double(max)
	return C.fluid_settings_getnum_range(s.ptr, cname, &cMin, &cMax) == FLUID_OK
}

func (s *Settings) Setint(name string, val int) bool {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	return C.fluid_settings_setint(s.ptr, cname, C.int(val)) == FLUID_OK
}

func (s *Settings) Getint(name string, val int) bool {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cVal := C.int(val)
	return C.fluid_settings_getint(s.ptr, cname, &cVal) == FLUID_OK
}

func (s *Settings) GetintDefault(name string, val int) bool {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cVal := C.int(val)
	return C.fluid_settings_getint_default(s.ptr, cname, &cVal) == FLUID_OK
}

func (s *Settings) GetintRange(name string, min int, max int) bool {
	cname := C.CString(name)
	defer C.free(unsafe.Pointer(cname))
	cMin := C.int(min)
	cMax := C.int(max)
	return C.fluid_settings_getint_range(s.ptr, cname, &cMin, &cMax) == FLUID_OK
}
