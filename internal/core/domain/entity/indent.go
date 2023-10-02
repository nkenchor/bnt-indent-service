package entity

import (
	"time"
)

type Indent struct {
	Reference         string       `json:"reference" bson:"reference"`
	IndentRef         string       `json:"indentRef" bson:"indentRef"`
	Location          Location     `json:"location" bson:"location" validate:"required"`
	Year              Year         `json:"year" bson:"year" validate:"required"`
	Denomination      Denomination `json:"denomination" bson:"denomination" validate:"required,min=1"`
	QuantityRequested int          `json:"quantity_requested"   bson:"quantity_requested" validate:"required"`
	QuantityGenerated int          `json:"quantity_generated" bson:"quantity_generated"`
	IsActive          bool 		   `json:"is_active" bson:"is_active"`
	CreatedAt         time.Time    `json:"created_at" bson:"created_at"`
	CreatedBy         User         `json:"created_by" bson:"created_by"`
	UpdatedAt         time.Time    `json:"updated_at" bson:"updated_at"`
	UpdatedBy         User         `json:"updated_by" bson:"updated_by"`
}
