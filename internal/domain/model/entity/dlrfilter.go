package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	mo "github.com/octoposprime/op-be-dlr/internal/domain/model/object"
)

// DlrFilter is a struct that represents the filter of a dlr.
type DlrFilter struct {
	Id        uuid.UUID    `json:"id"`         // Id is the id of the dlr.
	DlrData   string       `json:"dlr_name"`   // DlrData is the dlr name of the dlr.
	DlrType   mo.DlrType   `json:"dlr_type"`   // DlrType is the type of the dlr.
	DlrStatus mo.DlrStatus `json:"dlr_status"` // DlrStatus is the status of the dlr.
	Tags      []string     `json:"tags"`       // Tags is the tags of the dlr.

	CreatedAtFrom time.Time `json:"created_at_from"` // CreatedAt is in the between of CreatedAtFrom and CreatedAtTo.
	CreatedAtTo   time.Time `json:"created_at_to"`   // CreatedAt is in the between of CreatedAtFrom and CreatedAtTo.
	UpdatedAtFrom time.Time `json:"updated_at_from"` // UpdatedAt is in the between of UpdatedAtFrom and UpdatedAtTo.
	UpdatedAtTo   time.Time `json:"updated_at_to"`   // UpdatedAt is in the between of UpdatedAtFrom and UpdatedAtTo.

	SearchText string          `json:"search_text"` // SearchText is the full-text search value.
	SortType   string          `json:"sort_type"`   // SortType is the sorting type (ASC,DESC).
	SortField  mo.DlrSortField `json:"sort_field"`  // SortField is the sorting field of the dlr.

	Limit  int `json:"limit"`  // Limit provides to limitation row size.
	Offset int `json:"offset"` // Offset provides a starting row number of the limitation.
}

// NewDlrFilter creates a new *DlrFilter.
func NewDlrFilter(id uuid.UUID,
	dlrData string,
	dlrType mo.DlrType,
	dlrStatus mo.DlrStatus,
	tags []string,
	createdAtFrom time.Time,
	createdAtTo time.Time,
	updatedAtFrom time.Time,
	updatedAtTo time.Time,
	searchText string,
	sortType string,
	sortField mo.DlrSortField,
	limit int,
	offset int) *DlrFilter {
	return &DlrFilter{
		Id:            id,
		DlrData:       dlrData,
		DlrType:       dlrType,
		DlrStatus:     dlrStatus,
		Tags:          tags,
		CreatedAtFrom: createdAtFrom,
		CreatedAtTo:   createdAtTo,
		UpdatedAtFrom: updatedAtFrom,
		UpdatedAtTo:   updatedAtTo,
		SearchText:    searchText,
		SortType:      sortType,
		SortField:     sortField,
		Limit:         limit,
		Offset:        offset,
	}
}

// NewEmptyDlrFilter creates a new *DlrFilter with empty values.
func NewEmptyDlrFilter() *DlrFilter {
	return &DlrFilter{
		Id:            uuid.UUID{},
		DlrData:       "",
		DlrType:       mo.DlrTypeNONE,
		DlrStatus:     mo.DlrStatusNONE,
		Tags:          []string{},
		CreatedAtFrom: time.Time{},
		CreatedAtTo:   time.Time{},
		UpdatedAtFrom: time.Time{},
		UpdatedAtTo:   time.Time{},
		SearchText:    "",
		SortType:      "",
		SortField:     mo.DlrSortFieldNONE,
		Limit:         0,
		Offset:        0,
	}
}

// String returns a string representation of the DlrFilter.
func (s *DlrFilter) String() string {
	return fmt.Sprintf("Id: %v, "+
		"DlrData: %v, "+
		"DlrType: %v, "+
		"DlrStatus: %v, "+
		"Tags: %v, "+
		"CreatedAtFrom: %v, "+
		"CreatedAtTo: %v, "+
		"UpdatedAtFrom: %v, "+
		"UpdatedAtTo: %v, "+
		"SearchText: %v, "+
		"SortType: %v, "+
		"SortField: %v, "+
		"Limit: %v, "+
		"Offset: %v",
		s.Id,
		s.DlrData,
		s.DlrType,
		s.DlrStatus,
		s.Tags,
		s.CreatedAtFrom,
		s.CreatedAtTo,
		s.UpdatedAtFrom,
		s.UpdatedAtTo,
		s.SearchText,
		s.SortType,
		s.SortField,
		s.Limit,
		s.Offset)
}
