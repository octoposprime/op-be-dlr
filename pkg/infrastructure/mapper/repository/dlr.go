package infrastructure

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/lib/pq"
	me "github.com/octoposprime/op-be-dlr/internal/domain/model/entity"
	mo "github.com/octoposprime/op-be-dlr/internal/domain/model/object"
	tgorm "github.com/octoposprime/op-be-shared/tool/gorm"
)

// Dlr is a struct that represents the db mapper of a dlr basic values.
type Dlr struct {
	tgorm.Model

	DlrData   string         `json:"dlr_data" gorm:"not null;default:''"`  // DlrData is the dlr data of the dlr.
	DlrType   int            `json:"dlr_type" gorm:"not null;default:0"`   // DlrType is the type of the dlr.
	DlrStatus int            `json:"dlr_status" gorm:"not null;default:0"` // DlrStatus is the status of the dlr.
	Tags      pq.StringArray `json:"tags" gorm:"type:text[]"`              // Tags is the tags of the dlr.
}

// NewDlr creates a new *Dlr.
func NewDlr(id uuid.UUID,
	dlrData string,
	dlrType int,
	dlrStatus int,
	tags pq.StringArray) *Dlr {
	return &Dlr{
		Model:     tgorm.Model{ID: id},
		DlrData:   dlrData,
		DlrType:   dlrType,
		DlrStatus: dlrStatus,
		Tags:      tags,
	}
}

// String returns a string representation of the Dlr.
func (s *Dlr) String() string {
	return fmt.Sprintf("Id: %v, "+
		"DlrData: %v, "+
		"DlrType: %v, "+
		"DlrStatus: %v, "+
		"Tags: %v",
		s.ID,
		s.DlrData,
		s.DlrType,
		s.DlrStatus,
		s.Tags)
}

// NewDlrFromEntity creates a new *Dlr from entity.
func NewDlrFromEntity(entity me.Dlr) *Dlr {
	return &Dlr{
		Model:     tgorm.Model{ID: entity.Id},
		DlrData:   entity.DlrData,
		DlrType:   int(entity.DlrType),
		DlrStatus: int(entity.DlrStatus),
		Tags:      entity.Tags,
	}
}

// ToEntity returns a entity representation of the Dlr.
func (s *Dlr) ToEntity() *me.Dlr {
	return &me.Dlr{
		Id: s.ID,
		Dlr: mo.Dlr{
			DlrData:   s.DlrData,
			DlrType:   mo.DlrType(s.DlrType),
			DlrStatus: mo.DlrStatus(s.DlrStatus),
			Tags:      s.Tags,
		},
		CreatedAt: s.CreatedAt,
		UpdatedAt: s.UpdatedAt,
	}
}

type Dlrs []*Dlr

// NewDlrsFromEntities creates a new []*Dlr from entities.
func NewDlrFromEntities(entities []me.Dlr) Dlrs {
	dlrs := make([]*Dlr, len(entities))
	for i, entity := range entities {
		dlrs[i] = NewDlrFromEntity(entity)
	}
	return dlrs
}

// ToEntities creates a new []me.Dlr entity.
func (s Dlrs) ToEntities() []me.Dlr {
	dlrs := make([]me.Dlr, len(s))
	for i, dlr := range s {
		dlrs[i] = *dlr.ToEntity()
	}
	return dlrs
}
