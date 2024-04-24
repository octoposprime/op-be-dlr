package domain

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	mo "github.com/octoposprime/op-be-dlr/internal/domain/model/object"
)

// Dlr is a struct that represents the entity of a dlr basic values.
type Dlr struct {
	Id     uuid.UUID `json:"id"` // Id is the id of the dlr.
	mo.Dlr           // Dlr is the basic values of the dlr.

	// Only for view
	CreatedAt time.Time `json:"created_at"` // CreatedAt is the create time.
	UpdatedAt time.Time `json:"updated_at"` // UpdatedAt is the update time.
}

// NewDlr creates a new *Dlr.
func NewDlr(id uuid.UUID,
	dlr mo.Dlr) *Dlr {
	return &Dlr{
		Id:  id,
		Dlr: dlr,
	}
}

// NewEmptyDlr creates a new *Dlr with empty values.
func NewEmptyDlr() *Dlr {
	return &Dlr{
		Id:  uuid.UUID{},
		Dlr: *mo.NewEmptyDlr(),
	}
}

// String returns a string representation of the Dlr.
func (s *Dlr) String() string {
	return fmt.Sprintf("Id: %v, "+
		"Dlr: %v",
		s.Id,
		s.Dlr)
}

// Equals returns true if the Dlr is equal to the other Dlr.
func (s *Dlr) Equals(other *Dlr) bool {
	if s.Id != other.Id {
		return false
	}
	if !s.Dlr.Equals(&other.Dlr) {
		return false
	}
	return true
}

// Clone returns a clone of the Dlr.
func (s *Dlr) Clone() *Dlr {
	return &Dlr{
		Id:  s.Id,
		Dlr: *s.Dlr.Clone(),
	}
}

// IsEmpty returns true if the Dlr is empty.
func (s *Dlr) IsEmpty() bool {
	if s.Id.String() != "" && s.Id != (uuid.UUID{}) {
		return false
	}
	if !s.Dlr.IsEmpty() {
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
	s.Id = uuid.UUID{}
	s.Dlr.Clear()
}

// Validate validates the Dlr.
func (s *Dlr) Validate() error {
	if s.IsEmpty() {
		return mo.ErrorDlrIsEmpty
	}
	if err := s.Dlr.Validate(); err != nil {
		return err
	}
	return nil
}

// Dlrs contains a slice of *Dlr and total number of dlrs.
type Dlrs struct {
	Dlrs      []Dlr `json:"dlrs"`       // Dlrs is the slice of *Dlr.
	TotalRows int64 `json:"total_rows"` // TotalRows is the total number of rows.
}
