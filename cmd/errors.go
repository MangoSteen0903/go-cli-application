package cmd

import "errors"

var ErrInvalidNumofArgs = errors.New("invalid number of arguments")
var ErrInvalidSubCommand = errors.New("invalid sub command")
