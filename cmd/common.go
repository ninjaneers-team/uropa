package cmd

import (
	"os"

	"github.com/fatih/color"
	"github.com/ninjaneers-team/uropa/diff"
	"github.com/ninjaneers-team/uropa/dump"
	"github.com/ninjaneers-team/uropa/file"
	"github.com/ninjaneers-team/uropa/solver"
	"github.com/ninjaneers-team/uropa/state"
	"github.com/ninjaneers-team/uropa/utils"
)

var stopChannel chan struct{}

// SetStopCh sets the stop channel for long running commands.
// This is useful for cases when a process needs to be cancelled gracefully
// before it can complete to finish. Example: SIGINT
func SetStopCh(stopCh chan struct{}) {
	stopChannel = stopCh
}

var dumpConfig dump.Config

func syncMain(filename string, dry bool, parallelism int) error {
	// read target file
	targetContent, err := file.GetContentFromFile(filename)
	if err != nil {
		return err
	}

	client, err := utils.GetOpaClient(config)
	if err != nil {
		return err
	}

	// read the current state
	rawState, err := dump.Get(client, dumpConfig)
	if err != nil {
		return err
	}
	currentState, err := state.Get(rawState)
	if err != nil {
		return err
	}

	// read the target state
	rawState, err = file.Get(targetContent, currentState)
	if err != nil {
		return err
	}
	targetState, err := state.Get(rawState)
	if err != nil {
		return err
	}

	s, _ := diff.NewSyncer(currentState, targetState)
	stats, errs := solver.Solve(stopChannel, s, client, parallelism, dry)
	if errs != nil {
		return utils.ErrArray{Errors: errs}
	}
	printFn := color.New(color.FgGreen, color.Bold).PrintfFunc()
	printFn("Summary:\n")
	printFn("  Created: %v\n", stats.CreateOps)
	printFn("  Updated: %v\n", stats.UpdateOps)
	printFn("  Deleted: %v\n", stats.DeleteOps)
	if diffCmdNonZeroExitCode &&
		stats.CreateOps+stats.UpdateOps+stats.DeleteOps != 0 {
		os.Exit(2)
	}
	return nil
}
