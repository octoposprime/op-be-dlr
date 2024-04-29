package infrastructure

import (
	mo "github.com/octoposprime/op-be-dlr/internal/domain/model/object"
)

var DlrSortMap map[mo.DlrSortField]string = map[mo.DlrSortField]string{
	mo.DlrSortFieldId:        "id",
	mo.DlrSortFieldCreatedAt: "created_at",
	mo.DlrSortFieldUpdatedAt: "updated_at",
}
