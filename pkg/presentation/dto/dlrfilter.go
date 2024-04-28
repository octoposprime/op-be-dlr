package presentation

import (
	"fmt"
	"time"

	"github.com/google/uuid"
	me "github.com/octoposprime/op-be-dlr/internal/domain/model/entity"
	mo "github.com/octoposprime/op-be-dlr/internal/domain/model/object"
	pb "github.com/octoposprime/op-be-shared/pkg/proto/pb/dlr"
	tuuid "github.com/octoposprime/op-be-shared/tool/uuid"
	"google.golang.org/protobuf/types/known/timestamppb"
)

// DlrFilter is a struct that represents the filter dto of a dlr.
type DlrFilter struct {
	proto *pb.DlrFilter
}

// NewDlrFilter creates a new *DlrFilter.
func NewDlrFilter(pb *pb.DlrFilter) *DlrFilter {
	return &DlrFilter{
		proto: pb,
	}
}

// String returns a string representation of the DlrFilter.
func (s *DlrFilter) String() string {
	return fmt.Sprintf("Id: %v, "+
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
		s.proto.Id,
		s.proto.DlrType,
		s.proto.DlrStatus,
		s.proto.Tags,
		s.proto.CreatedAtFrom,
		s.proto.CreatedAtTo,
		s.proto.UpdatedAtFrom,
		s.proto.UpdatedAtTo,
		s.proto.SearchText,
		s.proto.SortType,
		s.proto.SortField,
		s.proto.Limit,
		s.proto.Offset)
}

// NewDlrFilterFromEntity creates a new *DlrFilter from entity.
func NewDlrFilterFromEntity(entity me.DlrFilter) *DlrFilter {
	id := entity.Id.String()
	dlrType := pb.DlrType(entity.DlrType)
	dlrStatus := pb.DlrStatus(entity.DlrStatus)
	tags := entity.Tags
	createdAtFrom := timestamppb.New(entity.CreatedAtFrom)
	createdAtTo := timestamppb.New(entity.CreatedAtTo)
	updatedAtFrom := timestamppb.New(entity.UpdatedAtFrom)
	updatedAtTo := timestamppb.New(entity.UpdatedAtTo)
	searchText := entity.SearchText
	sortType := entity.SortType
	sortField := pb.DlrSortField(entity.SortField)
	limit := int32(entity.Limit)
	offset := int32(entity.Offset)
	return &DlrFilter{
		&pb.DlrFilter{
			Id:            &id,
			DlrType:       &dlrType,
			DlrStatus:     &dlrStatus,
			Tags:          tags,
			CreatedAtFrom: createdAtFrom,
			CreatedAtTo:   createdAtTo,
			UpdatedAtFrom: updatedAtFrom,
			UpdatedAtTo:   updatedAtTo,
			SearchText:    &searchText,
			SortType:      &sortType,
			SortField:     &sortField,
			Limit:         &limit,
			Offset:        &offset,
		},
	}
}

// ToEntity returns a entity representation of the DlrFilter.
func (s *DlrFilter) ToEntity() *me.DlrFilter {
	id := uuid.UUID{}
	if s.proto.Id != nil {
		id = tuuid.FromString(*s.proto.Id)
	}
	dlrType := 0
	if s.proto.DlrType != nil {
		dlrType = int(*s.proto.DlrType)
	}
	dlrStatus := 0
	if s.proto.DlrStatus != nil {
		dlrStatus = int(*s.proto.DlrStatus)
	}
	tags := []string{}
	if s.proto.Tags != nil {
		tags = s.proto.Tags
	}
	createdAtFrom := time.Time{}
	if s.proto.CreatedAtFrom != nil {
		createdAtFrom = s.proto.CreatedAtFrom.AsTime()
	}
	createdAtTo := time.Time{}
	if s.proto.CreatedAtTo != nil {
		createdAtTo = s.proto.CreatedAtTo.AsTime()
	}
	updatedAtFrom := time.Time{}
	if s.proto.UpdatedAtFrom != nil {
		updatedAtFrom = s.proto.UpdatedAtFrom.AsTime()
	}
	updatedAtTo := time.Time{}
	if s.proto.UpdatedAtTo != nil {
		updatedAtTo = s.proto.UpdatedAtTo.AsTime()
	}
	searchText := ""
	if s.proto.SearchText != nil {
		searchText = string(*s.proto.SearchText)
	}
	sortType := ""
	if s.proto.SortType != nil {
		sortType = string(*s.proto.SortType)
	}
	sortField := 0
	if s.proto.SortField != nil {
		sortField = int(*s.proto.SortField)
	}
	limit := 0
	if s.proto.Limit != nil {
		limit = int(*s.proto.Limit)
	}
	offset := 0
	if s.proto.Offset != nil {
		offset = int(*s.proto.Offset)
	}
	return &me.DlrFilter{
		Id:            id,
		DlrType:       mo.DlrType(dlrType),
		DlrStatus:     mo.DlrStatus(dlrStatus),
		Tags:          tags,
		CreatedAtFrom: createdAtFrom,
		CreatedAtTo:   createdAtTo,
		UpdatedAtFrom: updatedAtFrom,
		UpdatedAtTo:   updatedAtTo,
		SearchText:    searchText,
		SortType:      sortType,
		SortField:     mo.DlrSortField(sortField),
		Limit:         limit,
		Offset:        offset,
	}
}
