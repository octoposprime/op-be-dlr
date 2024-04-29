package domain

import (
	"errors"

	smodel "github.com/octoposprime/op-be-shared/pkg/model"
)

var ERRORS []error = []error{
	ErrorNone,
	ErrorDlrIsEmpty,
	ErrorDlrDlrDataIsEmpty,
}

const (
	ErrId      string = "id"
	ErrDlr     string = "dlr"
	ErrDlrData string = "dlrdata"
)

const (
	ErrEmpty         string = "empty"
	ErrTooShort      string = "tooshort"
	ErrTooLong       string = "toolong"
	ErrNotValid      string = "notvalid"
	ErrInactive      string = "inactive"
	ErrAlreadyExists string = "alreadyexists"
)

var (
	ErrorNone              error = nil
	ErrorDlrIsEmpty        error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrDlr + smodel.ErrSep + ErrEmpty)
	ErrorDlrDlrDataIsEmpty error = errors.New(smodel.ErrBase + smodel.ErrSep + ErrDlr + smodel.ErrSep + ErrDlrData + smodel.ErrSep + ErrEmpty)
)

func GetErrors() []error {
	return ERRORS
}
