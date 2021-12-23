package midi

import (
	"fmt"

	mw "gitlab.com/gomidi/midi/writer"
)

type (
	// Factory applies operations, creates a new midi file.
	Factory interface {
		WriteSMF(operations []Operation) error
	}

	factory struct {
		file string
	}
)

func NewFactory(file string) Factory {
	return &factory{
		file: file,
	}
}

func (s *factory) WriteSMF(operations []Operation) error {
	return mw.WriteSMF(s.file, 1, func(w *mw.SMF) error {
		for i, op := range operations {
			if err := op(w); err != nil {
				return fmt.Errorf("failed to write smf at %d op %w", i, err)
			}
		}
		return nil
	})
}
