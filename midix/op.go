package midix

import (
	"gitlab.com/gomidi/midi/v2"
	"gitlab.com/gomidi/midi/v2/smf"
)

type OpFunc interface {
	Call(t *smf.Track, deltaticks uint32)
}

type MetaTrackSequenceName struct {
	Text string `json:"text"`
}

func (m MetaTrackSequenceName) Call(t *smf.Track, deltaticks uint32) {
	t.Add(deltaticks, smf.MetaTrackSequenceName(m.Text))
}

type MetaInstrument struct {
	Text string `json:"text"`
}

func (m MetaInstrument) Call(t *smf.Track, deltaticks uint32) {
	t.Add(deltaticks, smf.MetaInstrument(m.Text))
}

type ProgramChange struct {
	Channel uint8 `json:"channel"`
	Program uint8 `json:"program"`
}

func (p ProgramChange) Call(t *smf.Track, deltaticks uint32) {
	t.Add(deltaticks, midi.ProgramChange(p.Channel, p.Program))
}

type NoteOn struct {
	Channel  uint8 `json:"channel"`
	Key      uint8 `json:"key"`
	Velocity uint8 `json:"velocity"`
}

func (n NoteOn) Call(t *smf.Track, deltaticks uint32) {
	t.Add(deltaticks, midi.NoteOn(n.Channel, n.Key, n.Velocity))
}

type NoteOff struct {
	Channel uint8 `json:"channel"`
	Key     uint8 `json:"key"`
}

func (n NoteOff) Call(t *smf.Track, deltaticks uint32) {
	t.Add(deltaticks, midi.NoteOff(n.Channel, n.Key))
}

type MetaTempo struct {
	BPM float64 `json:"bpm"`
}

func (m MetaTempo) Call(t *smf.Track, deltaticks uint32) {
	t.Add(deltaticks, smf.MetaTempo(m.BPM))
}

type MetaMeter struct {
	Num   uint8 `json:"num"`
	Denom uint8 `json:"denom"`
}

func (m MetaMeter) Call(t *smf.Track, deltaticks uint32) {
	t.Add(deltaticks, smf.MetaMeter(m.Num, m.Denom))
}

type MetaKey struct {
	Key     uint8 `json:"key"`
	IsMajor bool  `json:"is_major"`
	Num     uint8 `json:"num"`
	IsFlat  bool  `json:"is_flat"`
}

func (m MetaKey) Call(t *smf.Track, deltaticks uint32) {
	t.Add(deltaticks, smf.MetaKey(m.Key, m.IsMajor, m.Num, m.IsFlat))
}

type MetaText struct {
	Text string `json:"text"`
}

func (m MetaText) Call(t *smf.Track, deltaticks uint32) {
	t.Add(deltaticks, smf.MetaText(m.Text))
}

type MetaLyric struct {
	Text string `json:"text"`
}

func (m MetaLyric) Call(t *smf.Track, deltaticks uint32) {
	t.Add(deltaticks, smf.MetaLyric(m.Text))
}

type MetaMarker struct {
	Text string `json:"text"`
}

func (m MetaMarker) Call(t *smf.Track, deltaticks uint32) {
	t.Add(deltaticks, smf.MetaMarker(m.Text))
}

type Close struct {
}

func (Close) Call(t *smf.Track, deltaticks uint32) {
	t.Close(deltaticks)
}
