package main

import (
	"errors"

	"github.com/berquerant/crd/errorx"
	"github.com/berquerant/crd/op"
	"github.com/spf13/cobra"
)

func overrideInstanceFromFlags(cmd *cobra.Command, instance *op.Instance) error {
	{
		x, err := getBPM(cmd)
		if err == nil {
			instance.BPM = &x
		} else if !errors.Is(err, errorx.ErrOK) {
			return err
		}
	}
	{
		x, err := getVelocity(cmd)
		if err == nil {
			instance.Velocity = &x
		} else if !errors.Is(err, errorx.ErrOK) {
			return err
		}
	}
	{
		x, err := getMeter(cmd)
		if err == nil {
			instance.Meter = &x
		} else if !errors.Is(err, errorx.ErrOK) {
			return err
		}
	}
	{
		x, err := getKey(cmd)
		if err == nil {
			instance.Key = &x
		} else if !errors.Is(err, errorx.ErrOK) {
			return err
		}
	}
	return nil
}
