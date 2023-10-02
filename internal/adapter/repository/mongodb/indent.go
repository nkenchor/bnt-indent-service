package repository

import (
	"bnt/bnt-indent-service/internal/core/domain/entity"
	shared "bnt/bnt-indent-service/internal/core/domain/shared"

	//errorhelper "bnt/bnt-indent-service/internal/core/helper/error-helper"
	logger "bnt/bnt-indent-service/internal/core/helper/log-helper"
	ports "bnt/bnt-indent-service/internal/port"
	"context"
	"reflect"
	"strconv"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type IndentInfra struct {
	Collection *mongo.Collection
}

func NewIndent(Collection *mongo.Collection) *IndentInfra {
	return &IndentInfra{Collection}
}

//UserRepo implements the repository.UserRepository interface
var _ ports.IndentRepository = &IndentInfra{}

func (r *IndentInfra) CreateIndent(indent entity.Indent) (interface{}, error) {
	logger.LogEvent("INFO", "Persisting indent configurations with reference: "+indent.Reference)
	
	_, err := r.Collection.InsertOne(context.TODO(),indent)
	if err != nil {
		return nil, err
	}
	
	logger.LogEvent("INFO", "Persisting indent configurations with reference: "+indent.Reference+" completed successfully...")
	return indent.Reference, nil
}

func (r *IndentInfra) UpdateIndent(reference string, indent entity.Indent) (interface{}, error) {
	logger.LogEvent("INFO", "Persisting indent configurations with reference: "+reference)
	_, err := r.Collection.ReplaceOne(
		context.TODO(),
		bson.M{"reference": bson.M{"$eq": reference}},indent)
	if err != nil {
		return nil, err
	}
	logger.LogEvent("INFO", "Persisting indent configurations with reference: "+reference+" completed successfully. ")
	return indent.Reference, nil
}
func (r *IndentInfra) EnableIndent(reference string) (interface{}, error) {
	logger.LogEvent("INFO", "Enabling indent configurations with reference: "+reference)

	_, err := r.Collection.UpdateOne(
		context.TODO(),
		bson.M{"reference": bson.M{"$eq": reference}},
		bson.M{"$set": bson.M{"is_active": true}},
	)
	if err != nil {
		return nil, err
	}
	logger.LogEvent("INFO", "Enabling indent configurations with reference: "+reference+" completed successfully. ")
	return reference, nil
}
func (r *IndentInfra) DisableIndent(reference string) (interface{}, error) {
	logger.LogEvent("INFO", "Disabling indent configurations with reference: "+reference)

	_, err := r.Collection.UpdateOne(
		context.TODO(),
		bson.M{"reference": bson.M{"$eq": reference}},
		bson.M{"$set": bson.M{"is_active": false}},
	)
	if err != nil {
		return nil, err
	}
	logger.LogEvent("INFO", "Disabling indent configurations with reference: "+reference+" completed successfully. ")
	return reference, nil
}
func (r *IndentInfra) DeleteIndent(reference string) (interface{}, error) {
	logger.LogEvent("INFO", "Deleting indent configurations with reference: "+reference)
	_, err := r.Collection.DeleteOne(
		context.TODO(),
		bson.M{"reference": bson.M{"$eq": reference}},
	)
	if err != nil {
		return nil, err
	}
	logger.LogEvent("INFO", "Deleting indent configurations with reference: "+reference+" completed successfully. ")
	return reference, nil
}
func (r *IndentInfra) GetIndentByRef(reference string) (interface{}, error) {
	logger.LogEvent("INFO", "Retrieving indent configurations with reference: "+reference)
	indent := entity.Indent{}
	filter := bson.M{"reference": reference}
	err := r.Collection.FindOne(context.TODO(), filter).Decode(&indent)
	if err != nil || indent == (entity.Indent{}) {
		return nil, err
	}
	logger.LogEvent("INFO", "Retrieving indent configurations with reference: "+reference+" completed successfully. ")
	return indent, nil
}

func (r *IndentInfra) GetIndentList(page string) (interface{}, error) {
	logger.LogEvent("INFO", "Retrieving all indent configuration entries...")
	var indents []entity.Indent
	var indent entity.Indent
	findOptions, err := GetPage(page)
	if err != nil {
		return nil, err
	}
	cursor, err := r.Collection.Find(context.TODO(), bson.M{}, findOptions)
	if err != nil {
		return []entity.Indent{}, err
	}
	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&indent)
		if err != nil {

			return nil, err
		}
		indents = append(indents, indent)
	}
	if reflect.ValueOf(indents).IsNil() {
		logger.LogEvent("INFO", "There are no results in this collection...")
		return []entity.Indent{}, nil
	}
	logger.LogEvent("INFO", "Retrieving all indent configuration entries completed successfully")
	return indents, nil
}

func (r *IndentInfra) GetYearList(count string) (interface{}, error) {
	logger.LogEvent("INFO", "Retrieving all year entries...")
	var years []entity.Year
	var upperIndex, err = strconv.Atoi(count)
	if err != nil || upperIndex == 0 {
		upperIndex = 1
	}
	for i := 2020; i < (2020 + upperIndex); i++ {

		years = append(years, entity.Year{Year: 2020, Code: strconv.Itoa(i)})

	}

	logger.LogEvent("INFO", "Retrieving all indent configuration entries completed successfully")
	return years, nil
}

func (r *IndentInfra) GetDenominationList() (interface{}, error) {
	logger.LogEvent("INFO", "Retrieving all denomination entries...")
	var denominations []entity.Denomination
	var denominationArray = shared.Denomination
    
	for value, code := range denominationArray {
		denominations = append(denominations, entity.Denomination{Denomination: strconv.Itoa(value) + " Naira", Code: code, Value: value})
	}

	logger.LogEvent("INFO", "Retrieving all denomination entries completed successfully")
	return denominations, nil
}
func (r *IndentInfra) GetLocationList() (interface{}, error) {
	logger.LogEvent("INFO", "Retrieving all location entries...")
	var locations []entity.Location
	var locationArray = shared.Location

	for location, code := range locationArray {
		locations = append(locations, entity.Location{Location: location, Code: code})
	}

	logger.LogEvent("INFO", "Retrieving all location entries completed successfully")
	return locations, nil
}
func (r *IndentInfra) GetIndentByDenomination(denomination string, page string) (interface{}, error) {
	logger.LogEvent("INFO", "Retrieving indent configurations with denomination: "+ denomination)
	indent := entity.Indent{}
	var indents []entity.Indent
	filter := bson.M{"denomination.denomination": denomination}
	findOptions, err := GetPage(page)
	if err != nil {
		return nil, err
	}
	
	cursor, err := r.Collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return []entity.Indent{}, err
	}
	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&indent)
	
		if err != nil {

			return nil, err
		}
		indents = append(indents, indent)
	}
	if reflect.ValueOf(indents).IsNil() {
		logger.LogEvent("INFO", "There are no results in this collection...")
		return []entity.Indent{}, nil
	}
	logger.LogEvent("INFO", "Retrieving indent configurations with denomination: "+denomination+" completed successfully. ")
	return indent, nil
}
func (r *IndentInfra) GetIndentByYear(year string, page string) (interface{}, error) {
	logger.LogEvent("INFO", "Retrieving indent configurations with year: "+ year)
	indent := entity.Indent{}
	var indents []entity.Indent
	filter := bson.M{"year.year": year}
	findOptions, err := GetPage(page)
	if err != nil {
		return nil, err
	}
	
	cursor, err := r.Collection.Find(context.TODO(), filter, findOptions)
	if err != nil {
		return []entity.Indent{}, err
	}
	for cursor.Next(context.TODO()) {
		err := cursor.Decode(&indent)
		if err != nil {

			return nil, err
		}
		indents = append(indents, indent)
	}
	if reflect.ValueOf(indents).IsNil() {
		logger.LogEvent("INFO", "There are no results in this collection...")
		return []entity.Indent{}, nil
	}
	logger.LogEvent("INFO", "Retrieving indent configurations with year: "+year+" completed successfully. ")
	return indent, nil
}
func (r *IndentInfra) GetLastIndentByDenominationYear(denomination string, year int16) (interface{}, error) {
	logger.LogEvent("INFO", "Retrieving indent configurations with denomination: "+ denomination)
	indent := entity.Indent{}
	filter := bson.M{"denomination.denomination": denomination, "year.year": year, "is_active": true}
	
	err := r.Collection.FindOne(context.TODO(), filter).Decode(&indent)
	if err != nil || indent == (entity.Indent{}) {
		return nil, err
	}
	logger.LogEvent("INFO", "Retrieving indent configurations with denomination: "+denomination+" completed successfully. ")
	return indent.Reference, nil
}