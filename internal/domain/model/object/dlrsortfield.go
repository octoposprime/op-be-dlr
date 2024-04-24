package domain

// DlrSortField is a type that represents the sort fields of a dlr.
type DlrSortField int8

const (
	DlrSortFieldNONE DlrSortField = iota
	DlrSortFieldId
	//DlrSortFieldName
	DlrSortFieldCreatedAt
	DlrSortFieldUpdatedAt
)
