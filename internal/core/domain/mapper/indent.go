package mapper

import (
	"bnt/bnt-indent-service/internal/core/domain/dto"
	"bnt/bnt-indent-service/internal/core/domain/entity"
	"fmt"
	"time"

	"github.com/google/uuid"
)

func MapCreateDto(createIndentDto dto.CreateIndent) entity.Indent {

	return entity.Indent{
		Reference:         uuid.New().String(),
		IndentRef:         fmt.Sprint(createIndentDto.Year) + " | " + createIndentDto.Denomination,
		Location:          entity.Location{Location: "Abuja", Code: "ABJ"},
		Year:              entity.Year{Year: createIndentDto.Year, Code: fmt.Sprint(createIndentDto.Year)},
		Denomination:      entity.Denomination{Denomination: createIndentDto.Denomination, Value: 5, Code: "5"},
		QuantityRequested: createIndentDto.QuantityRequested,
		CreatedAt:         time.Now().UTC(),
		CreatedBy: entity.User{UserReference: uuid.New().String(), UserFullName: "Nkenchor Osemeke",
			Unit: "ICT", Role: "Admin", Email: "nkenchor@osemeke.com", Phone: "+2347032127545"},
		UpdatedAt: time.Now().UTC(),
		UpdatedBy: entity.User{UserReference: uuid.New().String(), UserFullName: "Nkenchor Osemeke",
			Unit: "ICT", Role: "Admin", Email: "nkenchor@osemeke.com", Phone: "+2347032127545"},
	}

}

func MapUpdateDto(createIndentDto dto.UpdateIndent) entity.Indent {

	return entity.Indent{
		Reference:         uuid.New().String(),
		IndentRef:         fmt.Sprint(createIndentDto.Year) + " | " + createIndentDto.Denomination,
		Location:          entity.Location{Location: "Abuja", Code: "ABJ"},
		Year:              entity.Year{Year: createIndentDto.Year, Code: fmt.Sprint(createIndentDto.Year)},
		Denomination:      entity.Denomination{Denomination: createIndentDto.Denomination, Value: 5, Code: "5"},
		QuantityRequested: createIndentDto.QuantityRequested,
		CreatedAt:         time.Now().UTC(),
		CreatedBy: entity.User{UserReference: uuid.New().String(), UserFullName: "Nkenchor Osemeke",
			Unit: "ICT", Role: "Admin", Email: "nkenchor@osemeke.com", Phone: "+2347032127545"},
		UpdatedAt: time.Now().UTC(),
		UpdatedBy: entity.User{UserReference: uuid.New().String(), UserFullName: "Nkenchor Osemeke",
			Unit: "ICT", Role: "Admin", Email: "nkenchor@osemeke.com", Phone: "+2347032127545"},
	}

}
