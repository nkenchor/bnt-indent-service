package services

import (
	"bnt/bnt-indent-service/internal/core/domain/dto"
	"bnt/bnt-indent-service/internal/core/domain/mapper"
	logger "bnt/bnt-indent-service/internal/core/helper/log-helper"
	ports "bnt/bnt-indent-service/internal/port"

	"github.com/pkg/errors"
)

type indentService struct {
	indentRepository ports.IndentRepository
}

func (service *indentService) DisableIndent(reference string) (interface{}, error) {
	logger.LogEvent("INFO", "Disabling indent with reference: "+reference)
	_, err := service.GetIndentByRef(reference)

	if err != nil {
		return nil, err
	}
	return service.indentRepository.DisableIndent(reference)
}

func (service *indentService) GetIndentList(page string) (interface{}, error) {
	logger.LogEvent("INFO", "Getting all indents...")
	indent, err := service.indentRepository.GetIndentList(page)

	if err != nil {
		return nil, err
	}
	return indent, nil
}

func NewIndent(indentRepository ports.IndentRepository) *indentService {
	return &indentService{
		indentRepository: indentRepository,
	}
}

func (service *indentService) CreateIndent(createIndentDto dto.CreateIndent) (interface{}, error) {

	indent := mapper.MapCreateDto(createIndentDto)

	activeIndent,_:= service.GetLastIndentByDenominationYear(indent.Denomination.Denomination, indent.Year.Year)
	if activeIndent!= nil {	
		return nil, errors.New("Sorry, an indent with the same year and denomination is already active. Please disable currently active indent before creating.")
	}

	return service.indentRepository.CreateIndent(indent)
}

func (service *indentService) UpdateIndent(reference string, updateIndentDto dto.UpdateIndent) (interface{}, error) {
	indent := mapper.MapUpdateDto(updateIndentDto)
	return service.indentRepository.UpdateIndent(reference, indent)
}

func (service *indentService) DeleteIndent(reference string) (interface{}, error) {
	logger.LogEvent("INFO", "Deleting indent with reference: "+reference)
	_, err := service.GetIndentByRef(reference)
	if err != nil {
		return nil, err
	}
	return service.indentRepository.DeleteIndent(reference)
}
func (service *indentService) EnableIndent(reference string) (interface{}, error) {
	logger.LogEvent("INFO", "Enabling indent with reference: "+reference)
	_, err := service.GetIndentByRef(reference)
	if err != nil {
		return nil, err
	}
	return service.indentRepository.EnableIndent(reference)
}

func (service *indentService) GetIndentByRef(reference string) (interface{}, error) {
	logger.LogEvent("INFO", "Getting indent with reference: "+reference)
	indent, err := service.indentRepository.GetIndentByRef(reference)

	if err != nil {
		return nil, err
	}
	return indent, nil
}

func (service *indentService) GetYearList(count string) (interface{}, error) {
	logger.LogEvent("INFO", "Getting all year list...")
	years, err := service.indentRepository.GetYearList(count)
	if err != nil {
		return nil, err
	}
	return years, nil
}

func (service *indentService) GetLocationList() (interface{}, error) {
	logger.LogEvent("INFO", "Getting all location list...")
	locations, err := service.indentRepository.GetLocationList()
	if err != nil {
		return nil, err
	}
	return locations, nil
}
func (service *indentService) GetDenominationList() (interface{}, error) {
	logger.LogEvent("INFO", "Getting all denomination list...")
	denominations, err := service.indentRepository.GetDenominationList()
	if err != nil {
		return nil, err
	}
	return denominations, nil
}
func (service *indentService) GetIndentByYear(year string, page string) (interface{}, error) {
	logger.LogEvent("INFO", "Getting all indent by year...")
	indent, err := service.indentRepository.GetIndentByYear(year, page)

	if err != nil {
		return nil, err
	}
	return indent, nil
}
func (service *indentService) GetIndentByDenomination(denomination string, page string) (interface{}, error) {
	logger.LogEvent("INFO", "Getting all indent by denomination...")
	indent, err := service.indentRepository.GetIndentByYear(denomination, page)

	if err != nil {
		return nil, err
	}
	return indent, nil
}
func (service *indentService) GetLastIndentByDenominationYear(denomination string, year int16) (interface{}, error) {
	logger.LogEvent("INFO", "Getting all indent by denomination...")
	indent, err := service.indentRepository.GetLastIndentByDenominationYear(denomination, year)

	if err != nil {
		return nil, err
	}
	return indent, nil
}