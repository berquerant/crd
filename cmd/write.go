package main

import (
	"bytes"
	"fmt"
	"io"
	"log/slog"

	"github.com/berquerant/crd/chord"
	"github.com/berquerant/crd/errorx"
	"github.com/berquerant/crd/input"
	"github.com/berquerant/crd/midix"
	"github.com/berquerant/crd/op"
	"github.com/berquerant/crd/play"
	"github.com/berquerant/crd/util"
	"github.com/spf13/cobra"
	"gopkg.in/yaml.v3"
)

func init() {
	rootCmd.AddCommand(writeCmd)
	writeCmd.AddCommand(writeCmdParse, writeCmdEvent, writeCmdPlay, writeCmdConv)
	setBPMPersistentFlag(writeCmd)
	setVelocityPersistentFlag(writeCmd)
	setMeterPersistentFlag(writeCmd)
	setKeyPersistentFlag(writeCmd)
	setTrackFlag(writeCmd)
	writeCmdPlay.Flags().StringP("port", "p", "", "midi out port name")
	writeCmdConv.Flags().StringSliceP("command", "c", nil, "input instances conversions")
}

var writeCmd = &cobra.Command{
	Use:   "write [FILE]",
	Short: "write midi",
	Long: `generate midi from yaml

Examples:
crd write midi some.yml -o out.midi

YAML format:

# I7/III
- chord:
    # chord degree, see: crd info attr
    # I
    degree: "1"
    # chord symbol
    # name or meta.display
    # see: crd info chord
    # seventh
    name: "7"
  # optional, slash chord
  # major third
  base: "3"
  values:
    # 1 + 1/4 beat
    - "1"
    - "1/4"
  # optional, tempo
  bpm: 120
  # optional, velocity
  # available values: pp, p, mp, mf, f, ff
  # forte
  velocity: f
  # optional, meter
  meter: "4/4"
  # optional, key
  # see: crd key list
  # C minor
  key: "Cm"
  # metadata
  meta:
    txt: "text message"
    lic: "lyric message"
    mrk: "marker message"

# Rest
# 2 beats
- values:
    - 2
`,
	RunE: func(cmd *cobra.Command, args []string) error {
		wArgs, err := newWriteCmdArgs(cmd, args)
		if err != nil {
			return err
		}
		out, err := getOutput(cmd)
		if err != nil {
			return err
		}
		defer out.Close()
		return wArgs.writeMIDITo(out)
	},
}

var writeCmdPlay = &cobra.Command{
	Use:   "play [FILE]",
	Short: `play midi`,
	RunE: func(cmd *cobra.Command, args []string) error {
		wArgs, err := newWriteCmdArgs(cmd, args)
		if err != nil {
			return err
		}
		outPortName, _ := cmd.Flags().GetString("port")
		var buf bytes.Buffer
		if err := wArgs.writeMIDITo(&buf); err != nil {
			return err
		}

		return midix.NewReader().Play(bytes.NewReader(buf.Bytes()), outPortName)
	},
}

var writeCmdEvent = &cobra.Command{
	Use:   "event [FILE]",
	Short: `write midi events`,
	RunE: func(cmd *cobra.Command, args []string) error {
		wArgs, err := newWriteCmdArgs(cmd, args)
		if err != nil {
			return err
		}
		out, err := getOutput(cmd)
		if err != nil {
			return err
		}
		defer out.Close()

		var buf bytes.Buffer
		if err := wArgs.writeMIDITo(&buf); err != nil {
			return err
		}

		for ev := range midix.NewReader().Events(bytes.NewReader(buf.Bytes())) {
			fmt.Fprintf(out, "Track %d\t@%d(%d)\t%s\n",
				ev.TrackNo, ev.AbsTicks, ev.AbsTicks/midix.DefaultTicksPerQuoaterNote, ev.Message)
		}

		return nil
	},
}

var writeCmdConv = &cobra.Command{
	Use:   "conv [FILE]",
	Short: `convert instances yaml`,
	Long: `convert instances yaml

Examples:
# Write chord name to meta.txt.
crd write conv --command cmt`,
	RunE: func(cmd *cobra.Command, args []string) error {
		commands, _ := cmd.Flags().GetStringSlice("command")
		if len(commands) == 0 {
			return errorx.Invalid("require command")
		}
		instances, err := parseInstancesFromArgs(args)
		if err != nil {
			return err
		}

		getModifier := func(c string) (input.Modifier, bool) {
			switch c {
			case "cmt":
				slog.Debug("modifier", slog.String("command", c))
				return input.NewChordMetaTextMofidier(".", " on "), true
			default:
				return nil, true
			}
		}
		modifiers := []input.Modifier{}
		for i, c := range commands {
			x, ok := getModifier(c)
			if !ok {
				return errorx.NotFound("Modifier[%d] %s", i, c)
			}
			modifiers = append(modifiers, x)
		}

		for i, m := range modifiers {
			for j, x := range instances {
				if err := m.Modify(x); err != nil {
					return fmt.Errorf("%w: modifier[%d] instances[%d]", err, i, j)
				}
			}
		}

		wArgs, err := newWriteCmdArgsFromInputInstances(cmd, instances)
		if err != nil {
			return err
		}
		return writeYamlOutput(cmd, wArgs.instances)
	},
}

var writeCmdParse = &cobra.Command{
	Use:   "parse [FILE]",
	Short: `parse input instances yaml`,
	RunE: func(cmd *cobra.Command, args []string) error {
		wArgs, err := newWriteCmdArgs(cmd, args)
		if err != nil {
			return err
		}

		return writeYamlOutput(cmd, wArgs.instances)
	},
}

type writeCmdArgs struct {
	cmap      chord.Mapper
	instances []op.Instance
	trackSet  *midix.TrackSetController
}

func (w writeCmdArgs) writeMIDITo(wr io.Writer) error {
	writer, err := w.writeToPlay()
	if err != nil {
		return err
	}
	_, err = writer.WriteTo(wr)
	return err
}

func (w writeCmdArgs) writeToPlay() (midix.Writer, error) {
	mWriter := midix.NewWriter(midix.DefaultTicksPerQuoaterNote, w.trackSet)
	writer := play.NewWriter(w.cmap, func(k op.Key) play.Key {
		return play.NewKey(k, w.cmap)
	})
	if err := writer.Write(mWriter, w.instances); err != nil {
		return nil, err
	}
	return mWriter, nil
}

func newWriteCmdArgs(cmd *cobra.Command, args []string) (*writeCmdArgs, error) {
	inputInstances, err := parseInstancesFromArgs(args)
	if err != nil {
		return nil, err
	}
	return newWriteCmdArgsFromInputInstances(cmd, inputInstances)
}

func newWriteCmdArgsFromInputInstances(cmd *cobra.Command, inputInstances []*input.Instance) (*writeCmdArgs, error) {
	set, err := getTrackSetController(cmd)
	if err != nil {
		return nil, err
	}

	cmap, err := newChordMap(cmd)
	if err != nil {
		return nil, err
	}
	instances := make([]op.Instance, len(inputInstances))
	for i, x := range inputInstances {
		v := op.Instance{
			Values:   x.Values,
			BPM:      x.BPM,
			Velocity: x.Velocity,
			Meter:    x.Meter,
			Key:      x.Key,
			Meta:     x.Meta,
		}
		if i == 0 {
			if err := overrideInstanceFromFlags(cmd, &v); err != nil {
				return nil, err
			}
		}
		if c := x.Chord; c != nil {
			x, ok := cmap.GetChord(c.Chord)
			if !ok {
				return nil, errorx.NotFound("Chord %s", c.Chord)
			}
			y := op.NewChord(c.Degree, x, c.Base)
			v.Chord = &y
		}
		instances[i] = v
	}
	return &writeCmdArgs{
		cmap:      cmap,
		instances: instances,
		trackSet:  set,
	}, nil
}

func parseInstancesFromArgs(args []string) ([]*input.Instance, error) {
	var result []*input.Instance
	if err := readFileOrStdinFromArgs(args, func(r io.ReadCloser) error {
		instances, err := util.ReadAndParse(r, parseInstances)
		if err != nil {
			return err
		}
		result = instances
		return nil
	}); err != nil {
		return nil, err
	}
	return result, nil
}

func parseInstances(buf []byte) ([]*input.Instance, error) {
	var r []*input.Instance
	if err := yaml.Unmarshal(buf, &r); err != nil {
		return nil, err
	}
	return r, nil
}
