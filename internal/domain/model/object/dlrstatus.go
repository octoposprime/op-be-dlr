package domain

// DlrStatus is a status that represents the status of a dlr.
type DlrStatus int8

const (
	DlrStatusNONE DlrStatus = iota
	DlrStatusACTIVE
	DlrStatusINACTIVE
)
