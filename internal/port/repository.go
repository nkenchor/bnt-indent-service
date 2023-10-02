package ports

import "bnt/bnt-indent-service/internal/core/domain/entity"

type IndentRepository interface {
	CreateIndent(indent entity.Indent) (interface{}, error)
	UpdateIndent(ref string, indent entity.Indent) (interface{}, error)
	DeleteIndent(ref string) (interface{}, error)
	EnableIndent(ref string) (interface{}, error)
	DisableIndent(ref string) (interface{}, error)
	GetIndentByRef(CountryCode string) (interface{}, error)
	GetDenominationList() (interface{}, error)
	GetYearList(count string) (interface{}, error)
	GetLocationList() (interface{}, error)
	GetIndentList(page string) (interface{}, error)
	GetIndentByYear(year string, page string) (interface{}, error)
	GetIndentByDenomination(denomination string, page string) (interface{}, error)
	GetLastIndentByDenominationYear(denomination string, year int16) (interface{}, error)
}
