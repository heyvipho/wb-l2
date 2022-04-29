package main

import "errors"

var ErrNotEnoughArguments error = errors.New("Not enough arguments")
var ErrIncorrectArgument error = errors.New("Incorrect argument")

var ErrCanNotConvertTimeString error = errors.New("Can not convert time string")
