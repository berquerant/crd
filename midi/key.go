package midi

import (
	"github.com/berquerant/crd/note"
	"gitlab.com/gomidi/midi/midimessage/meta"
	"gitlab.com/gomidi/midi/midimessage/meta/key"
)

type (
	keyMaker func() meta.Key

	keyTuple struct {
		major, minor keyMaker
	}
)

var (
	keyMap = map[note.Name]map[note.Accidental]keyTuple{
		note.C: {
			note.Flat: {
				minor: key.BMin,
				major: key.BMaj,
			},
			note.Natural: {
				minor: key.CMin,
				major: key.CMaj,
			},
			note.Sharp: {
				minor: key.CSharpMin,
				major: key.DFlatMaj,
			},
		},
		note.D: {
			note.Flat: {
				minor: key.CSharpMin,
				major: key.DFlatMaj,
			},
			note.Natural: {
				minor: key.DMin,
				major: key.DMaj,
			},
			note.Sharp: {
				minor: key.DSharpMin,
				major: key.EFlatMaj,
			},
		},
		note.E: {
			note.Flat: {
				minor: key.EFlatMin,
				major: key.EFlatMaj,
			},
			note.Natural: {
				minor: key.EMin,
				major: key.EMaj,
			},
			note.Sharp: {
				minor: key.FMin,
				major: key.FMaj,
			},
		},
		note.F: {
			note.Flat: {
				minor: key.FMin,
				major: key.FMaj,
			},
			note.Natural: {
				minor: key.FMin,
				major: key.FMaj,
			},
			note.Sharp: {
				minor: key.FSharpMin,
				major: key.FSharpMaj,
			},
		},
		note.G: {
			note.Flat: {
				minor: key.FSharpMin,
				major: key.GFlatMaj,
			},
			note.Natural: {
				minor: key.GMin,
				major: key.GMaj,
			},
			note.Sharp: {
				minor: key.GSharpMin,
				major: key.AFlatMaj,
			},
		},
		note.A: {
			note.Flat: {
				minor: key.GSharpMin,
				major: key.AFlatMaj,
			},
			note.Natural: {
				minor: key.AMin,
				major: key.AMaj,
			},
			note.Sharp: {
				minor: key.BFlatMin,
				major: key.BFlatMaj,
			},
		},
		note.B: {
			note.Flat: {
				minor: key.BFlatMin,
				major: key.BFlatMaj,
			},
			note.Natural: {
				minor: key.BMin,
				major: key.BMaj,
			},
			note.Sharp: {
				minor: key.CMin,
				major: key.CMaj,
			},
		},
	}
)

func getKey(name note.Name, accidental note.Accidental, isMinor bool) keyMaker {
	t := keyMap[name][accidental]
	if isMinor {
		return t.minor
	}
	return t.major
}
