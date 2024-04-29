package presentation

import (
	"fmt"

	me "github.com/octoposprime/op-be-dlr/internal/domain/model/entity"
	mo "github.com/octoposprime/op-be-dlr/internal/domain/model/object"
	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/dlr"
	tuuid "github.com/octoposprime/op-be-shared/tool/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// Dlr is a struct that represents the dto of a dlr basic values.
type Dlr struct {
	proto *pb.Dlr
}

// NewDlr creates a new *Dlr.
func NewDlr(pb *pb.Dlr) *Dlr {
	return &Dlr{
		proto: pb,
	}
}

// String returns a string representation of the Dlr.
func (s *Dlr) String() string {
	return fmt.Sprintf("Id: %v, "+
		"DlrData: %v, "+
		"DlrType: %v, "+
		"DlrStatus: %v, "+
		"Tags: %v",
		s.proto.Id,
		s.proto.DlrData,
		s.proto.DlrType,
		s.proto.DlrStatus,
		s.proto.Tags)
}

// NewDlrFromEntity creates a new *Dlr from entity.
func NewDlrFromEntity(entity me.Dlr) *Dlr {
	return &Dlr{
		&pb.Dlr{
			Id:        entity.Id.String(),
			DlrData:   entity.DlrData,
			DlrType:   pb.DlrType(entity.DlrType),
			DlrStatus: pb.DlrStatus(entity.DlrStatus),
			Tags:      entity.Tags,

			// Only for view
			CreatedAt: timestamppb.New(entity.CreatedAt),
			UpdatedAt: timestamppb.New(entity.UpdatedAt),
		},
	}
}

// ToPb returns a protobuf representation of the Dlr.
func (s *Dlr) ToPb() *pb.Dlr {
	return s.proto
}

// ToEntity returns a entity representation of the Dlr.
func (s *Dlr) ToEntity() *me.Dlr {
	return &me.Dlr{
		Id: tuuid.FromString(s.proto.Id),
		Dlr: mo.Dlr{
			DlrData:   s.proto.DlrData,
			DlrType:   mo.DlrType(s.proto.DlrType),
			DlrStatus: mo.DlrStatus(s.proto.DlrStatus),
			Tags:      s.proto.Tags,
		},
	}
}

type Dlrs struct {
	Dlrs      []*Dlr `json:"dlrs"`
	TotalRows int64  `json:"total_rows"`
}

// NewDlrsFromEntities creates a new []*Dlr from entities.
func NewDlrFromEntities(entities me.Dlrs) Dlrs {
	dlrs := make([]*Dlr, len(entities.Dlrs))
	for i, entity := range entities.Dlrs {
		dlrs[i] = NewDlrFromEntity(entity)
	}

	return Dlrs{
		Dlrs:      dlrs,
		TotalRows: entities.TotalRows,
	}
}

// ToPbs returns a protobuf representation of the Dlrs.
func (s Dlrs) ToPbs() *pb.Dlrs {
	dlrs := make([]*pb.Dlr, len(s.Dlrs))
	for i, dlr := range s.Dlrs {
		dlrs[i] = dlr.proto
	}
	return &pb.Dlrs{
		Dlrs:      dlrs,
		TotalRows: s.TotalRows,
	}
}
