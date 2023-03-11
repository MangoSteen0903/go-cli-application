package cmd

import "errors"

var ErrInvalidNumofArgs = errors.New("invalid number of arguments")
var ErrInvalidSubCommand = errors.New("invalid sub command")
var ErrInvalidNumOfPositionalArgs = errors.New("invalid number of positional argument")
var ErrHelpRequest = errors.New("flag: help requested")
var ErrFlagNotDefiend = errors.New("flag provided but not defined: %s")
