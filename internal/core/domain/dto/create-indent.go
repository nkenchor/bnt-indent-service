package dto

type CreateIndent struct {
	Year              int16  `json:"year" bson:"year" validate:"required"`
	Denomination      string `json:"denomination" bson:"denomination" validate:"required,min=1"`
	QuantityRequested int    `json:"quantity_requested"   bson:"quantity_requested" validate:"required"`
}
