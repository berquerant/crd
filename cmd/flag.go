package main

import (
	"errors"
	"fmt"
	"strings"

	"github.com/berquerant/crd/errorx"
	"github.com/berquerant/crd/op"
	"github.com/berquerant/crd/util"
	"github.com/spf13/cobra"
)

//
// common flag definitions
//

func setAttributePersistentFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringSlice("attr", nil, "additional attribute definitions")
}

func getAttributeFlag(cmd *cobra.Command) []string {
	v, _ := cmd.Flags().GetStringSlice("attr")
	return v
}

func setChordPersistentFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringSlice("chord", nil, "additional chord definitions")
}

func getChordFlag(cmd *cobra.Command) []string {
	v, _ := cmd.Flags().GetStringSlice("chord")
	return v
}

func setBPMPersistentFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().Uint("bpm", 0, "override bpm")
}

func getBPM(cmd *cobra.Command) (op.BPM, error) {
	v, _ := cmd.Flags().GetUint("bpm")
	if v == 0 {
		var d op.BPM
		return d, errorx.ErrOK
	}
	return op.BPM(v), nil
}

func setVelocityPersistentFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().String("velocity", "",
		fmt.Sprintf("override velocity: %s", strings.Join(op.GetDynamicSignStrings(), ",")))
}

func getVelocity(cmd *cobra.Command) (op.DynamicSign, error) {
	v, _ := cmd.Flags().GetString("velocity")
	if v == "" {
		return op.UnknownDynamicSign, errorx.ErrOK
	}
	d := op.NewDynamicSign(v)
	if d == op.UnknownDynamicSign {
		return d, errorx.Invalid("velocity %s", v)
	}
	return d, nil
}

func setMeterPersistentFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().String("meter", "", "override meter")
}

func getMeter(cmd *cobra.Command) (op.Meter, error) {
	v, _ := cmd.Flags().GetString("meter")
	var d op.Meter
	if v == "" {
		return d, errorx.ErrOK
	}
	r, err := util.ParseRat(v)
	if err != nil {
		return d, err
	}
	return op.NewMeter(r.Num, r.Denom)
}

func setKeyPersistentFlag(cmd *cobra.Command) {
	cmd.PersistentFlags().StringP("key", "k", "", "override key")
}

func getKey(cmd *cobra.Command) (op.Key, error) {
	v, _ := cmd.Flags().GetString("key")
	if v == "" {
		var d op.Key
		return d, errorx.ErrOK
	}
	return op.ParseKey(v)
}

func getScale(cmd *cobra.Command) (*op.Scale, error) {
	key, err := getKey(cmd)
	if err != nil {
		if !errors.Is(err, errorx.ErrOK) {
			return nil, err
		}
		key = op.MustParseKey("C")
	}

	scale, err := op.NewScale(key)
	if err != nil {
		return nil, err
	}
	return scale, nil
}
