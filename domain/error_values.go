package domain

import "errors"

var (
	ErrUrlParse= errors.New("cannot parse url. check given url\n")
	ErrRegexpCompile = errors.New("cannot compile regexp pattern. check given regexp pattern\n")
	ErrMakeDirAll = errors.New("cannnot make directory structure\n")
	ErrSaveFile = errors.New("cannnot save a html file\n")
	ErrEmptyParams = errors.New("required parameters are empty.\n")
)