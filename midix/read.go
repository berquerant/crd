package midix

import (
	"io"
	"iter"

	"gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/smf"
)

func GetOutPortNames() []string {
	ports := midi.GetOutPorts()
	names := make([]string, len(ports))
	for i, p := range ports {
		names[i] = p.String()
	}
	return names
}

func GetInPortNames() []string {
	ports := midi.GetInPorts()
	names := make([]string, len(ports))
	for i, p := range ports {
		names[i] = p.String()
	}
	return names
}

type Reader interface {
	Events(r io.Reader) iter.Seq[smf.TrackEvent]
	Play(r io.Reader, outPortName string) error
}

var (
	_ Reader = &MIDIReader{}
)

func NewReader() *MIDIReader {
	return &MIDIReader{}
}

type MIDIReader struct{}

func (r MIDIReader) read(rd io.Reader) *smf.TracksReader {
	return smf.ReadTracksFrom(rd)
}

func (r MIDIReader) Play(rd io.Reader, outPortName string) error {
	defer midi.CloseDriver()
	out, err := midi.FindOutPort(outPortName)
	if err != nil {
		return err
	}
	return r.read(rd).Play(out)
}

func (r MIDIReader) Events(rd io.Reader) iter.Seq[smf.TrackEvent] {
	return func(yield func(smf.TrackEvent) bool) {
		var done bool
		r.read(rd).Do(func(ev smf.TrackEvent) {
			if done {
				return
			}
			if !yield(ev) {
				done = true
			}
		})
	}
}
