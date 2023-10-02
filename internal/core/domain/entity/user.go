package entity



type User struct {
	UserReference   string      `json:"user_reference" bson:"user_reference"`
	UserFullName    string 	  	`json:"user_full_name" bson:"user_full_name"`
	Unit            string 		`json:"unit" bson:"unit"`
	Role            string      `json:"role" bson:"role"`
	Email 			string 		`json:"email" bson:"email"`
	Phone 			string 		`json:"phone" bson:"phone"`
}
