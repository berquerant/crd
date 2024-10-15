package example_test

import (
	"bytes"
	"testing"
)

func TestExample(t *testing.T) {
	e := &executor{}
	if !e.init(t) {
		return
	}
	defer e.close()

	for _, tc := range []struct {
		title  string
		args   []string
		input  string
		output string
	}{
		{
			title: "write midi event overriding meta",
			args: []string{"write", "event",
				"--key", "A", "--bpm", "80", "--meter", "3/4", "--velocity", "mf"},
			input: `# Write
# - I, 1 beat
# - V7, 1 beat
# - IV, 2 beats
# in A major as midi events.
- chord:
    name: MajorTriad
    degree: "1"
  values:
    - "1"
  key: C # overrided by --key
- chord:
    name: DominantSeventh
    degree: "5"
  values:
    - "1"
- chord:
    name: MajorTriad
    degree: "4"
  values:
    - "2"`,
			output: `Track 0	@0(0)	MetaTrackName text: "Piano"
Track 0	@0(0)	MetaInstrument text: "Piano"
Track 0	@0(0)	ProgramChange channel: 0 program: 0
Track 0	@0(0)	MetaTempo bpm: 80.00
Track 0	@0(0)	MetaTimeSig meter: 3/4
Track 0	@0(0)	MetaKeySig key: AMaj
Track 0	@0(0)	NoteOn channel: 0 key: 57 velocity: 85
Track 0	@0(0)	NoteOn channel: 0 key: 69 velocity: 85
Track 0	@0(0)	NoteOn channel: 0 key: 73 velocity: 85
Track 0	@0(0)	NoteOn channel: 0 key: 76 velocity: 85
Track 0	@960(1)	NoteOff channel: 0 key: 57
Track 0	@960(1)	NoteOff channel: 0 key: 69
Track 0	@960(1)	NoteOff channel: 0 key: 73
Track 0	@960(1)	NoteOff channel: 0 key: 76
Track 0	@960(1)	NoteOn channel: 0 key: 64 velocity: 85
Track 0	@960(1)	NoteOn channel: 0 key: 76 velocity: 85
Track 0	@960(1)	NoteOn channel: 0 key: 80 velocity: 85
Track 0	@960(1)	NoteOn channel: 0 key: 83 velocity: 85
Track 0	@960(1)	NoteOn channel: 0 key: 86 velocity: 85
Track 0	@1920(2)	NoteOff channel: 0 key: 64
Track 0	@1920(2)	NoteOff channel: 0 key: 76
Track 0	@1920(2)	NoteOff channel: 0 key: 80
Track 0	@1920(2)	NoteOff channel: 0 key: 83
Track 0	@1920(2)	NoteOff channel: 0 key: 86
Track 0	@1920(2)	NoteOn channel: 0 key: 62 velocity: 85
Track 0	@1920(2)	NoteOn channel: 0 key: 74 velocity: 85
Track 0	@1920(2)	NoteOn channel: 0 key: 78 velocity: 85
Track 0	@1920(2)	NoteOn channel: 0 key: 81 velocity: 85
Track 0	@3840(4)	NoteOff channel: 0 key: 62
Track 0	@3840(4)	NoteOff channel: 0 key: 74
Track 0	@3840(4)	NoteOff channel: 0 key: 78
Track 0	@3840(4)	NoteOff channel: 0 key: 81
Track 0	@3840(4)	MetaEndOfTrack
`,
		},
		{
			title: "write midi event meta",
			args:  []string{"write", "event"},
			input: `# Write
# - I, 1 beat
# - V7, 1 beat
# - IV, 2 beats
# in C major as midi events.
- chord:
    name: MajorTriad
    degree: "1"
  values:
    - "1"
  key: C
  meta:
    # txt is text
    txt: "I"
    # lic is lyric
    lic: "Start"
- chord:
    name: DominantSeventh
    degree: "5"
  values:
    - "1"
  meta:
    txt: "V7"
    mrk: "dominant"
- chord:
    name: MajorTriad
    degree: "4"
  values:
    - "2"
  meta:
    txt: "IV"`,
			output: `Track 0	@0(0)	MetaTrackName text: "Piano"
Track 0	@0(0)	MetaInstrument text: "Piano"
Track 0	@0(0)	ProgramChange channel: 0 program: 0
Track 0	@0(0)	MetaTempo bpm: 100.00
Track 0	@0(0)	MetaTimeSig meter: 4/4
Track 0	@0(0)	MetaKeySig key: CMaj
Track 0	@0(0)	MetaText text: "I"
Track 0	@0(0)	MetaLyric text: "Start"
Track 0	@0(0)	NoteOn channel: 0 key: 48 velocity: 64
Track 0	@0(0)	NoteOn channel: 0 key: 60 velocity: 64
Track 0	@0(0)	NoteOn channel: 0 key: 64 velocity: 64
Track 0	@0(0)	NoteOn channel: 0 key: 67 velocity: 64
Track 0	@960(1)	NoteOff channel: 0 key: 48
Track 0	@960(1)	NoteOff channel: 0 key: 60
Track 0	@960(1)	NoteOff channel: 0 key: 64
Track 0	@960(1)	NoteOff channel: 0 key: 67
Track 0	@960(1)	MetaText text: "V7"
Track 0	@960(1)	MetaMarker text: "dominant"
Track 0	@960(1)	NoteOn channel: 0 key: 55 velocity: 64
Track 0	@960(1)	NoteOn channel: 0 key: 67 velocity: 64
Track 0	@960(1)	NoteOn channel: 0 key: 71 velocity: 64
Track 0	@960(1)	NoteOn channel: 0 key: 74 velocity: 64
Track 0	@960(1)	NoteOn channel: 0 key: 77 velocity: 64
Track 0	@1920(2)	NoteOff channel: 0 key: 55
Track 0	@1920(2)	NoteOff channel: 0 key: 67
Track 0	@1920(2)	NoteOff channel: 0 key: 71
Track 0	@1920(2)	NoteOff channel: 0 key: 74
Track 0	@1920(2)	NoteOff channel: 0 key: 77
Track 0	@1920(2)	MetaText text: "IV"
Track 0	@1920(2)	NoteOn channel: 0 key: 53 velocity: 64
Track 0	@1920(2)	NoteOn channel: 0 key: 65 velocity: 64
Track 0	@1920(2)	NoteOn channel: 0 key: 69 velocity: 64
Track 0	@1920(2)	NoteOn channel: 0 key: 72 velocity: 64
Track 0	@3840(4)	NoteOff channel: 0 key: 53
Track 0	@3840(4)	NoteOff channel: 0 key: 65
Track 0	@3840(4)	NoteOff channel: 0 key: 69
Track 0	@3840(4)	NoteOff channel: 0 key: 72
Track 0	@3840(4)	MetaEndOfTrack
`,
		},
		{
			title: "write midi",
			args:  []string{"write"},
			input: `# Write
# - I, 1 beat
# - V7, 1 beat
# - IV, 2 beats
# in C major to out.midi.
- chord:
    name: MajorTriad
    degree: "1"
  values:
    - "1"
  key: C
- chord:
    name: DominantSeventh
    degree: "5"
  values:
    - "1"
- chord:
    name: MajorTriad
    degree: "4"
  values:
    - "2"`,
			output: ignoreStdout,
		},
		{
			title: "syllable text into instances yaml key changes",
			args:  []string{"text", "conv", "syllable", "--key", "D"},
			input: `D[1] ; D, 1 beat
A_7/E[1]{key=C} ; A7/E, 1 beat, key: D major to C major
E[2] ; E, 2 beats
R[1] ; Rest, 1 beat
`,
			output: `- chord:
    degree: "1"
    name: ""
  values:
    - "1"
- chord:
    degree: "6"
    name: "7"
    base: "5"
  values:
    - "1"
  key: C
  meta:
    key: C
- chord:
    degree: "3"
    name: ""
  values:
    - "2"
- values:
    - "1"
`,
		},
		{
			title: "syllable text into instances yaml",
			args:  []string{"text", "conv", "syllable", "--key", "D"},
			input: `D[1] ; D, 1 beat
A_7/E[1] ; A7/E, 1 beat
E[2] ; E, 2 beats
R[1] ; Rest, 1 beat
`,
			output: `- chord:
    degree: "1"
    name: ""
  values:
    - "1"
- chord:
    degree: "5"
    name: "7"
    base: "5"
  values:
    - "1"
- chord:
    degree: "2"
    name: ""
  values:
    - "2"
- values:
    - "1"
`,
		},
		{
			title: "degree text into instances yaml meta",
			args:  []string{"text", "conv", "degree"},
			input: `1[1] ; I, 1 beat
5_7/5[1] ; V7/V, 1 beat
2[2]{key=Cm} ; II, 2 beats, modulation
R[1] ; Rest, 1 beat
`,
			output: `- chord:
    degree: "1"
    name: ""
  values:
    - "1"
- chord:
    degree: "5"
    name: "7"
    base: "5"
  values:
    - "1"
- chord:
    degree: "2"
    name: ""
  values:
    - "2"
  key: Cm
  meta:
    key: Cm
- values:
    - "1"
`,
		},
		{
			title: "degree text into instances yaml",
			args:  []string{"text", "conv", "degree"},
			input: `1[1] ; I, 1 beat
5_7/5[1] ; V7/V, 1 beat
2[2] ; II, 2 beats
R[1] ; Rest, 1 beat
`,
			output: `- chord:
    degree: "1"
    name: ""
  values:
    - "1"
- chord:
    degree: "5"
    name: "7"
    base: "5"
  values:
    - "1"
- chord:
    degree: "2"
    name: ""
  values:
    - "2"
- values:
    - "1"
`,
		},
		{
			title:  "info attr list",
			args:   []string{"info", "attr", "list"},
			output: ignoreStdout,
		},
		{
			title: "info attr describe perfect4",
			args: []string{"info", "attr", "describe",
				"--target", "Perfect4", "--root", "D"},
			output: `attribute:
    name: Perfect4
    degree: "4"
semitone: 5
semitone_without_octave: 5
root: D
applied: G
`,
		},
		{
			title: "info attr describe augmented4 using sharp",
			args: []string{"info", "attr", "describe",
				"--target", "Augmented4", "--root", "C", "--precedeSharp"},
			output: `attribute:
    name: Augmented4
    degree: '#4'
semitone: 6
semitone_without_octave: 6
root: C
applied: F#
`,
		},
		{
			title:  "info chord list",
			args:   []string{"info", "chord", "list"},
			output: ignoreStdout,
		},
		{
			title: "info chord describe Eb",
			args: []string{"info", "chord", "describe",
				"--target", "Eb"},
			output: `chord:
    name: MajorTriad
    meta:
        display: ""
    attributes:
        - Perfect1
        - Major3
        - Perfect5
root: Eb
attributes:
    - attribute:
        name: Perfect1
        degree: "1"
      semitone: 0
      semitone_without_octave: 0
      root: Eb
      applied: Eb
    - attribute:
        name: Major3
        degree: "3"
      semitone: 4
      semitone_without_octave: 4
      root: Eb
      applied: G
    - attribute:
        name: Perfect5
        degree: "5"
      semitone: 7
      semitone_without_octave: 7
      root: Eb
      applied: Bb
`,
		},
		{
			title:  "info key list",
			args:   []string{"info", "key", "list"},
			output: ignoreStdout,
		},
		{
			title: "info key describe Cm",
			args: []string{"info", "key", "describe",
				"--key", "Cm"},
			output: `scale:
    key: Cm
    notes:
        - C
        - D
        - Eb
        - F
        - G
        - Ab
        - Bb
    flat: 3
diatonic:
    triads:
        - Cm
        - Ddim
        - Eb
        - Fm
        - Gm
        - Ab
        - Bb
    sevenths:
        - Cm7
        - Dm7b5
        - Ebmaj7
        - Fm7
        - Gm7
        - Abmaj7
        - Bb_7
`,
		},
		{
			title: "info key conv parallel of C",
			args: []string{"info", "key", "conv",
				"--key", "C", "--command", "p"},
			output: `Cm
`,
		},
		{
			title: "info key conv subdominant of parallel of C",
			args: []string{"info", "key", "conv",
				"--key", "C", "--command", "ps"},
			output: `Fm
`,
		},
	} {
		t.Run(tc.title, func(t *testing.T) {
			stdin := bytes.NewBufferString(tc.input)
			e.do(t, tc.output, stdin, tc.args...)
		})
	}
}
