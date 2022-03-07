package users

type UserFilter struct {
	Id          string		`json:"id" gorm:"column:id;primary_key" bson:"_id" dynamodbav:"id" firestore:"id" match:"equal" validate:"required,max=40"`
	Username    string		`json:"username" gorm:"column:username" bson:"username" dynamodbav:"username" firestore:"username" match:"prefix" validate:"required,username,max=100"`
	Email       string		`json:"email" gorm:"column:email" bson:"email" dynamodbav:"email" firestore:"email" match:"prefix" validate:"email,max=100"`
	Phone       string		`json:"phone" gorm:"column:phone" bson:"phone" dynamodbav:"phone" firestore:"phone" validate:"required,phone,max=18"`
	PageIndex 	int			`mapstructure:"pageIndex" json:"pageIndex,omitempty" gorm:"column:pageIndex" bson:"pageIndex,omitempty" dynamodbav:"pageIndex,omitempty" firestore:"pageIndex,omitempty"`
	PageSize 	int			`mapstructure:"pageSize" json:"pageSize,omitempty" gorm:"column:pageSize" bson:"pageSize,omitempty" dynamodbav:"pageSize,omitempty" firestore:"pageSize,omitempty"`
}

type SearchResult struct {
	List 		[]User	`mapstructure:"list" json:"list,omitempty" gorm:"column:list" bson:"list,omitempty" dynamodbav:"list,omitempty" firestore:"list,omitempty"`
	Total 		int		`mapstructure:"total" json:"total,omitempty" gorm:"column:total" bson:"total,omitempty" dynamodbav:"total,omitempty" firestore:"total,omitempty"`
	PageIndex 	int		`mapstructure:"pageIndex" json:"pageIndex,omitempty" gorm:"column:pageIndex" bson:"pageIndex,omitempty" dynamodbav:"pageIndex,omitempty" firestore:"pageIndex,omitempty"`
	PageSize 	int		`mapstructure:"pageSize" json:"pageSize,omitempty" gorm:"column:pageSize" bson:"pageSize,omitempty" dynamodbav:"pageSize,omitempty" firestore:"pageSize,omitempty"`
}