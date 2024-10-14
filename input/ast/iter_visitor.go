package ast

import (
	"iter"
	"sync"
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
	VisitSwitch(s, v.Meta)
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
	VisitSwitch(s, v.Meta)
}

func (s *IterVisitor) VisitChordMeta(v *ChordMeta) {
	if v == nil || !s.send(v) {
		return
	}
	for _, x := range v.Data {
		VisitSwitch(s, x)
	}
}

func (s *IterVisitor) VisitChordMetadata(v *ChordMetadata) {
	if v == nil || !s.send(v) {
		return
	}
}
