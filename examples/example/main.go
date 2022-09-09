package main

import (
	"math"
	"math/rand"
	"time"

	"github.com/hydrz/fluidsynth.go/v2"
)

func main() {
	// Create the settings
	settings := fluidsynth.NewSettings()
	if settings == nil {
		panic("Failed to create the settings!")
	}
	defer settings.Delete()

	// Change the settings
	settings.Setstr("audio.driver", "alsa")

	// Create the synthesizer
	synth := fluidsynth.NewSynth(settings)
	if synth == nil {
		panic("Failed to create the synth!")
	}
	defer synth.Delete()

	// Load a SoundFont and reset presets (so that new instruments get used from the SoundFont)
	// Depending on the size of the SoundFont, this will take some time to complete...
	sFontId := synth.SFLoad("example.sf2", true)

	if sFontId == fluidsynth.FLUID_FAILED {
		panic("Loading the SoundFont failed!")
	}

	audioDriver := fluidsynth.NewAudioDriver(settings, synth)
	if audioDriver == nil {
		panic("Failed to create the audio driver!")
	}
	defer audioDriver.Delete()

	for i := 0; i < 12; i++ {
		/* Generate a random key */
		key := 60 + uint8(math.Ceil(12*rand.Float64()))

		/* Play a note */
		synth.NoteOn(0, key, 80)

		/* Sleep for 1 second */
		time.Sleep(1 * time.Second)

		/* Stop the note */
		synth.NoteOff(0, key)
	}
}
