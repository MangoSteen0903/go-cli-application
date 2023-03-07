package cmd

import "errors"

var ErrInvalidArguments = errors.New("invalid arguments number")
var ErrInvalidSubCommand = errors.New("invalid sub-command")
var ErrNotURLSpecified = errors.New("URL not specified")
