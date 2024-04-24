package domain

import (
	"fmt"
)

// Dlr is a struct that represents the object of a dlr basic values.
type Dlr struct {
	DlrData   string    `json:"dlr_data"`   // DlrData is the data of the dlr.
	DlrType   DlrType   `json:"dlr_type"`   // DlrType is the type of the dlr.
	DlrStatus DlrStatus `json:"dlr_status"` // DlrStatus is the status of the dlr.
	Tags      []string  `json:"tags"`       // Tags is the tags of the dlr.
}

// NewDlr creates a new *Dlr.
func NewDlr(dlrData string,
	dlrType DlrType,
	dlrStatus DlrStatus,
	tags []string) *Dlr {
	return &Dlr{
		DlrData:   dlrData,
		DlrType:   dlrType,
		DlrStatus: dlrStatus,
		Tags:      tags,
	}
}

// NewEmptyDlr creates a new *Dlr with empty values.
func NewEmptyDlr() *Dlr {
	return &Dlr{
		DlrData:   "",
		DlrType:   DlrTypeNONE,
		DlrStatus: DlrStatusNONE,
		Tags:      []string{},
	}
}

// String returns a string representation of the Dlr.
func (s *Dlr) String() string {
	return fmt.Sprintf("DlrData: %v, "+
		"DlrType: %v, "+
		"DlrStatus: %v, "+
		"Tags: %v",
		s.DlrData,
		s.DlrType,
		s.DlrStatus,
		s.Tags)
}

// Equals returns true if the Dlr is equal to the other Dlr.
func (s *Dlr) Equals(other *Dlr) bool {
	if s.DlrData != other.DlrData {
		return false
	}
	if s.DlrType != other.DlrType {
		return false
	}
	if s.DlrStatus != other.DlrStatus {
		return false
	}
	for i := range s.Tags {
		if s.Tags[i] != other.Tags[i] {
			return false
		}
	}
	return true
}

// Clone returns a clone of the Dlr.
func (s *Dlr) Clone() *Dlr {
	return &Dlr{
		DlrData:   s.DlrData,
		DlrType:   s.DlrType,
		DlrStatus: s.DlrStatus,
		Tags:      s.Tags,
	}
}

// IsEmpty returns true if the Dlr is empty.
func (s *Dlr) IsEmpty() bool {
	if s.DlrData != "" {
		return false
	}
	if s.DlrType != DlrTypeNONE {
		return false
	}
	if s.DlrStatus != DlrStatusNONE {
		return false
	}
	if len(s.Tags) != 0 {
		return false
	}
	return true
}

// IsNotEmpty returns true if the Dlr is not empty.
func (s *Dlr) IsNotEmpty() bool {
	return !s.IsEmpty()
}

// Clear clears the Dlr.
func (s *Dlr) Clear() {
	s.DlrData = ""
	s.DlrType = DlrTypeNONE
	s.DlrStatus = DlrStatusNONE
	s.Tags = []string{}
}

// Validate validates the Dlr.
func (s *Dlr) Validate() error {
	if s.IsEmpty() {
		return ErrorDlrIsEmpty
	}
	if s.DlrData == "" {
		return ErrorDlrDlrDataIsEmpty
	}
	return nil
}
