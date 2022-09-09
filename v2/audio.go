package fluidsynth

// #include <fluidsynth.h>
// #include <stdlib.h>
import "C"
import "unsafe"

type AudioDriver struct {
	ptr *C.fluid_audio_driver_t
}

func NewAudioDriver(settings *Settings, synth *Synth) *AudioDriver {
	cPtr := C.new_fluid_audio_driver(settings.ptr, synth.ptr)

	if cPtr == nil {
		return nil
	}

	return &AudioDriver{ptr: cPtr}
}

func (d *AudioDriver) Delete() {
	C.delete_fluid_audio_driver(d.ptr)
}

func AudioDriverRegister(adrivers string) int {
	cAdrivers := C.CString(adrivers)
	defer C.free(unsafe.Pointer(cAdrivers))
	return int(C.fluid_audio_driver_register(&cAdrivers))
}

type FileRenderer struct {
	ptr *C.fluid_file_renderer_t
}

func NewFileRenderer(synth Synth) *FileRenderer {
	cPtr := C.new_fluid_file_renderer(synth.ptr)
	if cPtr == nil {
		return nil
	}
	return &FileRenderer{ptr: cPtr}
}

func (r *FileRenderer) Delete() {
	C.delete_fluid_file_renderer(r.ptr)
}

func (r *FileRenderer) ProcessBlock() bool {
	return C.fluid_file_renderer_process_block(r.ptr) == C.FLUID_OK
}
