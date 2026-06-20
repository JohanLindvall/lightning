// Code generated from instruments.json (minio/simdjson-go testdata); root renamed Benchmark.
package bench

type Benchmark struct {
	Graphstate  any `json:"graphstate"`
	Instruments []struct {
		DefaultFilterCutoff           int64  `json:"default_filter_cutoff"`
		DefaultFilterCutoffEnabled    bool   `json:"default_filter_cutoff_enabled"`
		DefaultFilterMode             int64  `json:"default_filter_mode"`
		DefaultFilterResonance        int64  `json:"default_filter_resonance"`
		DefaultFilterResonanceEnabled bool   `json:"default_filter_resonance_enabled"`
		DefaultPan                    int64  `json:"default_pan"`
		DuplicateCheckType            int64  `json:"duplicate_check_type"`
		DuplicateNoteAction           int64  `json:"duplicate_note_action"`
		Fadeout                       int64  `json:"fadeout"`
		GlobalVolume                  int64  `json:"global_volume"`
		GraphInsert                   int64  `json:"graph_insert"`
		LegacyFilename                string `json:"legacy_filename,nocopy"`
		MidiBank                      int64  `json:"midi_bank"`
		MidiChannel                   int64  `json:"midi_channel"`
		MidiDrumSet                   int64  `json:"midi_drum_set"`
		MidiProgram                   int64  `json:"midi_program"`
		Name                          string `json:"name,nocopy"`
		NewNoteAction                 int64  `json:"new_note_action"`
		NoteMap                       any    `json:"note_map"`
		PanningEnvelope               struct {
			LoopEnd   int64 `json:"loop_end"`
			LoopStart int64 `json:"loop_start"`
			Nodes     []struct {
				Tick  int64 `json:"tick"`
				Value int64 `json:"value"`
			} `json:"nodes"`
			ReleaseNode  int64 `json:"release_node"`
			SustainEnd   int64 `json:"sustain_end"`
			SustainStart int64 `json:"sustain_start"`
		} `json:"panning_envelope"`
		PitchEnvelope struct {
			LoopEnd   int64 `json:"loop_end"`
			LoopStart int64 `json:"loop_start"`
			Nodes     []struct {
				Tick  int64 `json:"tick"`
				Value int64 `json:"value"`
			} `json:"nodes"`
			ReleaseNode  int64 `json:"release_node"`
			SustainEnd   int64 `json:"sustain_end"`
			SustainStart int64 `json:"sustain_start"`
		} `json:"pitch_envelope"`
		PitchPanCenter        int64 `json:"pitch_pan_center"`
		PitchPanSeparation    int64 `json:"pitch_pan_separation"`
		PitchToTempoLock      int64 `json:"pitch_to_tempo_lock"`
		RandomCutoffWeight    int64 `json:"random_cutoff_weight"`
		RandomPanWeight       int64 `json:"random_pan_weight"`
		RandomResonanceWeight int64 `json:"random_resonance_weight"`
		RandomVolumeWeight    int64 `json:"random_volume_weight"`
		SampleMap             any   `json:"sample_map"`
		Tuning                any   `json:"tuning"`
		VolumeEnvelope        struct {
			LoopEnd   int64 `json:"loop_end"`
			LoopStart int64 `json:"loop_start"`
			Nodes     []struct {
				Tick  int64 `json:"tick"`
				Value int64 `json:"value"`
			} `json:"nodes"`
			ReleaseNode  int64 `json:"release_node"`
			SustainEnd   int64 `json:"sustain_end"`
			SustainStart int64 `json:"sustain_start"`
		} `json:"volume_envelope"`
		VolumeRampDown int64 `json:"volume_ramp_down"`
		VolumeRampUp   int64 `json:"volume_ramp_up"`
	} `json:"instruments"`
	Message   any    `json:"message"`
	Name      string `json:"name,nocopy"`
	Orderlist any    `json:"orderlist"`
	Patterns  []struct {
		Data           any    `json:"data"`
		Name           string `json:"name,nocopy"`
		Rows           int64  `json:"rows"`
		RowsPerBeat    int64  `json:"rows_per_beat"`
		RowsPerMeasure int64  `json:"rows_per_measure"`
	} `json:"patterns"`
	Pluginstate any `json:"pluginstate"`
	Samples     []struct {
		C5Samplerate   int64  `json:"c5_samplerate"`
		GlobalVolume   int64  `json:"global_volume"`
		LegacyFilename string `json:"legacy_filename,nocopy"`
		Length         int64  `json:"length"`
		LoopEnd        int64  `json:"loop_end"`
		LoopStart      int64  `json:"loop_start"`
		Name           string `json:"name,nocopy"`
		Pan            int64  `json:"pan"`
		SustainEnd     int64  `json:"sustain_end"`
		SustainStart   int64  `json:"sustain_start"`
		VibratoDepth   int64  `json:"vibrato_depth"`
		VibratoRate    int64  `json:"vibrato_rate"`
		VibratoSweep   int64  `json:"vibrato_sweep"`
		VibratoType    int64  `json:"vibrato_type"`
		Volume         int64  `json:"volume"`
	} `json:"samples"`
	Version int64 `json:"version"`
}
