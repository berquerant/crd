package chord

import (
	_ "embed"
	"fmt"

	"github.com/berquerant/crd/errorx"
	"github.com/berquerant/crd/logx"
	"github.com/berquerant/crd/note"
	"gopkg.in/yaml.v3"
)

type (
	Semitone = note.Semitone
	Degree   = note.Degree
)

// Named degree.
type Attribute struct {
	Name   string `yaml:"name"`
	Degree Degree `yaml:"degree,omitempty"`
}

func (a Attribute) validate() error {
	if a.Name == "" {
		return errorx.Invalid("Attribute should have name")
	}
	return nil
}

func (a Attribute) Semitone() (Semitone, bool) {
	return a.Degree.Semitone()
}

//go:generate go run github.com/berquerant/crd/cmd gen attr -d 20 -o attribute.yml
//go:embed attribute.yml
var basicAttributes []byte

func BasicAttributes() []Attribute {
	attrs, err := ParseAttributes(basicAttributes)
	if err != nil {
		logx.Panic(err)
	}
	return attrs
}

func ParseAttributes(b []byte) ([]Attribute, error) {
	var attrs []Attribute
	if err := yaml.Unmarshal(b, &attrs); err != nil {
		return nil, err
	}
	for _, a := range attrs {
		if err := a.validate(); err != nil {
			return nil, err
		}
	}
	return attrs, nil
}

var (
	genAttrDegreeNamePrefix = map[note.DegreeName]string{
		note.MajorDegree:      "Major",
		note.MinorDegree:      "Minor",
		note.PerfectDegree:    "Perfect",
		note.AugmentedDegree:  "Augmented",
		note.DiminishedDegree: "Diminished",
	}
)

func GenerateAttributes(maxDegree uint) []Attribute {
	attrs := []Attribute{}
	for d := range note.GenerateDegrees(maxDegree) {
		prefix, ok := genAttrDegreeNamePrefix[d.Name]
		if !ok {
			continue
		}

		name := fmt.Sprintf("%s%d", prefix, d.Value)
		a := Attribute{
			Name:   name,
			Degree: d,
		}
		attrs = append(attrs, a)
	}

	return attrs
}
