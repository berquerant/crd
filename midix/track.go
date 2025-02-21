package midix

import (
	"fmt"
	"log/slog"

	"github.com/berquerant/crd/errorx"
	"github.com/berquerant/crd/logx"
	"gitlab.com/gomidi/midi/v2/smf"
)

type TrackNoSelector interface {
	Select(opType OpType) int
}

var (
	_ TrackNoSelector = &TrackNoSelectorImpl{}
)

type TrackNoSelectorImpl struct {
	trackNum int
}

func NewTrackNoSelector(trackNum int) (*TrackNoSelectorImpl, error) {
	if trackNum < 1 {
		return nil, errorx.Invalid("TrackNoSelector requires positive trackNum, %d", trackNum)
	}
	return &TrackNoSelectorImpl{
		trackNum: trackNum,
	}, nil
}

func (t TrackNoSelectorImpl) Select(opType OpType) int {
	switch opType := opType.(type) {
	case *MetaTrack:
		return 0
	case *FixedTrack:
		switch t.trackNum {
		case 1:
			return 0
		default:
			n := opType.TrackNo
			return n%(t.trackNum-1) + 1
		}
	default:
		logx.Panic(errorx.Unexpected("TrackOp: %#v", opType))
		return -1
	}
}

type TrackSetController struct {
	set      *TrackSet
	selector TrackNoSelector
}

func NewTrackSetControllerFromTrackNum(trackNum int) (*TrackSetController, error) {
	selector, err := NewTrackNoSelector(trackNum)
	if err != nil {
		return nil, err
	}
	set := NewTrackSetFromTrackNum(trackNum)
	return &TrackSetController{
		set:      set,
		selector: selector,
	}, nil
}

func NewTrackSetController(set *TrackSet, selector TrackNoSelector) *TrackSetController {
	return &TrackSetController{
		set:      set,
		selector: selector,
	}
}

func (c *TrackSetController) Add(op *TrackOp) {
	n := c.selector.Select(op.Type)
	c.set.Add(n, op)
}

func (c *TrackSetController) Distribute(op *TrackOp) {
	for i := range c.set.Len() {
		c.set.Add(i, op)
	}
}

func (c TrackSetController) Set() *TrackSet { return c.set }

type TrackSet struct {
	list []*Track
}

func NewTrackSetFromTrackNum(trackNum int) *TrackSet {
	if trackNum < 0 {
		trackNum = 0
	}
	list := make([]*Track, trackNum)
	for i := range trackNum {
		list[i] = NewTrack()
	}
	return &TrackSet{
		list: list,
	}
}

func NewTrackSet(list []*Track) *TrackSet {
	return &TrackSet{
		list: list,
	}
}

func (ts TrackSet) Len() int         { return len(ts.list) }
func (ts TrackSet) Get(i int) *Track { return ts.list[i] }

func (ts *TrackSet) Add(trackNo int, op *TrackOp) {
	slog.Debug("TrackSet",
		slog.Int("track", trackNo),
		slog.String("op_type", fmt.Sprintf("%T", op.Func)),
		logx.JSON("op", op),
	)

	delta := op.TickDelta
	t := ts.list[trackNo]
	t.Add(op)
	for i, x := range ts.list {
		if i == trackNo {
			continue
		}
		x.AddTickDelta(delta)
	}
}

type Track struct {
	tickDelta uint32
	ops       []*TrackOp
}

func NewTrack() *Track {
	return &Track{}
}

func (t Track) Len() int                       { return len(t.ops) }
func (t *Track) AddTickDelta(tickDelta uint32) { t.tickDelta += tickDelta }
func (t *Track) Add(op *TrackOp) {
	op.TickDelta += t.tickDelta
	t.ops = append(t.ops, op)
	t.tickDelta = 0
}

func (t Track) Apply(tt *smf.Track) {
	for _, x := range t.ops {
		x.Call(tt)
	}
}

type OpType interface {
	IsOpType()
}

//go:generate go tool marker -output track_marker_generated.go -method IsOpType -type FixedTrack,MetaTrack

func NewFixedTrack(no int) *FixedTrack {
	return &FixedTrack{
		TrackNo: no,
	}
}

type FixedTrack struct {
	TrackNo int
}

func NewMetaTrack() *MetaTrack {
	return &MetaTrack{}
}

type MetaTrack struct{}

type TrackOp struct {
	TickDelta uint32 `json:"tickdelta"`
	Func      OpFunc `json:"func"`
	Type      OpType `json:"type"`
}

func (op TrackOp) Call(t *smf.Track) {
	op.Func.Call(t, op.TickDelta)
}

func NewTrackOp(tickDelta uint32, opType OpType, opFunc OpFunc) *TrackOp {
	return &TrackOp{
		TickDelta: tickDelta,
		Func:      opFunc,
		Type:      opType,
	}
}
