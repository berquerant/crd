package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/berquerant/crd/cc"
	"github.com/berquerant/crd/logger"
	"github.com/berquerant/crd/midi"
)

var (
	out           = flag.String("o", "crd.out.mid", "Output filepath.")
	bpm           = flag.Int("b", 100, "BPM.")
	meter         = flag.String("m", "4/4", "Time signature.")
	verbose       = flag.Int("v", 0, "Logging verbosity.")
	doLex         = flag.Bool("lex", false, "Only print the result of lex.")
	doParse       = flag.Bool("p", false, "Only print the result of parse.")
	doUnparse     = flag.Bool("unparse", false, "Only print the result of unparse, string expression of the result of parse.")
	showSemitones = flag.Bool("s", false, "Only print chord semitones. 48 means C4.")
)

const usage = `Usage of crd:
  echo SCORE | crd [flags] -o FILE

Example of SCORE:
  C[1/2] G[1/2] Am[1/2] Em[1/2] F[1/2] C[1/2] F[1/2] G[1/2]

SCORE is a sequence of chords, chord V forms SCORE like
V V | V V V
Vertical bars, spaces and newlines are only for readability.

A chord is a string formatted like NOTE OPTION [VALUE].
NOTE is a note name, C, D#, Gb, ...
VALUE ia a relative note value, for example 1 means a whole note, 1/2 means a half note.
OPTION is a chord option, like m (minor). Available options are below:

  m    minor triad
  aug  augmented triad
  dim  diminished triad
  7    dominant seventh
  m7   minor seventh
  M7   major seventh
  mM7  minor major seventh
  dim7 diminished seventh
  m7-5 half diminished seventh
  aug7 augmented seventh
  6    add sixth
  m6   add minor sixth
  sus4 suspended forth

Flags:`

func main() {
	flag.Usage = func() {
		fmt.Fprintln(os.Stderr, usage)
		flag.PrintDefaults()
	}
	flag.Parse()
	initLogger()
	if isDebug() {
		doDebug()
		return
	}

	mn, md, mok := parseMeter()
	if !mok {
		log.Panicf("invalid time signature %s", *meter)
	}
	// parse input and construct AST
	lexer := cc.NewLexer(os.Stdin)
	lexer.Debug(*verbose)
	status := cc.Parse(lexer)
	logger.Get().Info("parser exit with %d", status)
	if err := lexer.Err(); err != nil {
		log.Panicf("parser got error %v", err)
	}
	// generate midi operations based on AST
	w := midi.NewASTWriter(midi.NewWriter())
	w.Writer().BPM(*bpm)
	w.Writer().Meter(mn, md)
	for _, n := range lexer.Result().NodeList {
		w.WriteNode(n)
	}
	// write midi file
	f := midi.NewFactory(*out)
	if err := f.WriteSMF(w.Writer().Operations()); err != nil {
		log.Panic(err)
	}
}

func parseMeter() (uint8, uint8, bool) {
	v := strings.Split(*meter, "/")
	if len(v) != 2 {
		return 0, 0, false
	}
	x, err := strconv.Atoi(v[0])
	if err != nil {
		return 0, 0, false
	}
	y, err := strconv.Atoi(v[1])
	if err != nil {
		return 0, 0, false
	}
	return uint8(x), uint8(y), true
}

func initLogger() {
	if *verbose < 1 {
		logger.Get().SetLevel(logger.Linfo)
	} else {
		logger.Get().SetLevel(logger.Ldebug)
	}
}

func isDebug() bool {
	return *doLex || *doParse || *doUnparse || *showSemitones
}

func doDebug() {
	l := cc.NewLexer(os.Stdin)
	l.Debug(*verbose)
	debugger := cc.NewDebugger(l)
	if *doLex {
		debugger.Lex()
		return
	}
	if *doUnparse {
		debugger.Unparse()
		return
	}
	if *showSemitones {
		debugger.Semitones()
		return
	}
	debugger.Parse()
}
