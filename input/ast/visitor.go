package ast

import (
	"iter"
	"sync"

	"github.com/berquerant/ybase"
)

type IterVisitor struct {
	nodeC chan Node
	mux   sync.Mutex
}

func NewIterVisitor() *IterVisitor {
	var v IterVisitor
	return &v
}

var (
	_ Visitor = &IterVisitor{}
)

func (s *IterVisitor) All(root Node) iter.Seq[Node] {
	s.mux.Lock()
	defer s.mux.Unlock()

	s.nodeC = make(chan Node, 100)
	go func() {
		VisitSwitch(s, root)
		close(s.nodeC)
	}()

	return func(yield func(Node) bool) {
		for x := range s.nodeC {
			if !yield(x) {
				break
			}
		}
		for range s.nodeC {
		}
	}
}

func (s *IterVisitor) send(v Node) bool {
	s.nodeC <- v
	return true
}

func (s *IterVisitor) VisitChordList(v *ChordList) {
	if v == nil || !s.send(v) {
		return
	}
	for _, x := range v.List {
		VisitSwitch(s, x)
	}
}
func (s *IterVisitor) VisitChord(v *Chord) {
	if v == nil || !s.send(v) {
		return
	}
	VisitSwitch(s, v.Degree)
	VisitSwitch(s, v.Symbol)
	VisitSwitch(s, v.Base)
	VisitSwitch(s, v.Values)
}
func (s *IterVisitor) VisitChordDegree(v *ChordDegree) {
	if v == nil || !s.send(v) {
		return
	}
}
func (s *IterVisitor) VisitChordSymbol(v *ChordSymbol) {
	if v == nil || !s.send(v) {
		return
	}
}
func (s *IterVisitor) VisitChordBase(v *ChordBase) {
	if v == nil || !s.send(v) {
		return
	}
	VisitSwitch(s, v.Degree)
}
func (s *IterVisitor) VisitChordValues(v *ChordValues) {
	if v == nil || !s.send(v) {
		return
	}
	for _, x := range v.Values {
		VisitSwitch(s, x)
	}
}
func (s *IterVisitor) VisitChordValue(v *ChordValue) {
	if v == nil || !s.send(v) {
		return
	}
}
func (s *IterVisitor) VisitRest(v *Rest) {
	if v == nil || !s.send(v) {
		return
	}
	VisitSwitch(s, v.Values)
}

// MapVisitor maps Token value and AST into dictionary.
type MapVisitor struct {
	extract func(ybase.Token) (any, bool)
	result  map[string]any
}

var (
	_ Visitor = &MapVisitor{}
)

func NewMapVisitor(extractor func(ybase.Token) (any, bool)) *MapVisitor {
	return &MapVisitor{
		extract: extractor,
	}
}

func (s MapVisitor) Result() map[string]any {
	return s.result
}

func (s *MapVisitor) VisitChordList(v *ChordList) {
	if v == nil {
		s.result = nil
		return
	}
	xs := make([]map[string]any, len(v.List))

	for i, x := range v.List {
		VisitSwitch(s, x)
		xs[i] = s.result
	}

	s.result = map[string]any{
		"list": xs,
	}
}

func (s *MapVisitor) VisitChord(v *Chord) {
	if v == nil {
		s.result = nil
		return
	}
	d := map[string]any{}

	VisitSwitch(s, v.Degree)
	d["degree"] = s.result

	VisitSwitch(s, v.Symbol)
	if x := s.result; x != nil {
		d["symbol"] = x
	}

	VisitSwitch(s, v.Base)
	if x := s.result; x != nil {
		d["base"] = x
	}

	VisitSwitch(s, v.Values)
	d["values"] = s.result

	s.result = d
}

func (s *MapVisitor) VisitChordDegree(v *ChordDegree) {
	if v == nil {
		s.result = nil
		return
	}
	d := map[string]any{}
	s.setWhenOK(d, "degree")(v.Degree)
	s.setWhenOK(d, "accidental")(v.Accidental)
	s.result = d
}

func (s *MapVisitor) VisitChordSymbol(v *ChordSymbol) {
	if v == nil {
		s.result = nil
		return
	}

	d := map[string]any{}
	s.setWhenOK(d, "symbol")(v.Symbol)
	s.result = d
}

func (s *MapVisitor) VisitChordBase(v *ChordBase) {
	if v == nil {
		s.result = nil
		return
	}
	if v == nil {
		s.result = nil
		return
	}

	d := map[string]any{}
	VisitSwitch(s, v.Degree)
	d["degree"] = s.result
	s.result = d
}

func (s *MapVisitor) VisitChordValues(v *ChordValues) {
	if v == nil {
		s.result = nil
		return
	}
	xs := make([]map[string]any, len(v.Values))
	for i, x := range v.Values {
		VisitSwitch(s, x)
		xs[i] = s.result
	}
	s.result = map[string]any{
		"values": xs,
	}
}

func (s *MapVisitor) VisitChordValue(v *ChordValue) {
	if v == nil {
		s.result = nil
		return
	}
	d := map[string]any{}
	s.setWhenOK(d, "num")(v.Num)
	s.setWhenOK(d, "denom")(v.Denom)
	s.result = d
}

func (s *MapVisitor) VisitRest(v *Rest) {
	if v == nil {
		s.result = nil
		return
	}
	VisitSwitch(s, v.Values)
	s.result = map[string]any{
		"values": s.result,
	}
}

func (s MapVisitor) setWhenOK(d map[string]any, key string) func(ybase.Token) {
	return func(token ybase.Token) {
		if x, ok := s.extract(token); ok {
			d[key] = x
		}
	}
}
