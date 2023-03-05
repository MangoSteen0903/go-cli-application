package cmd

import "errors"

var ErrNotServerSpecified = errors.New("server should be specified")
var ErrInvalidSubCommand = errors.New("invalid sub-command")
var ErrInvalidMethods = errors.New("methods invalid")
