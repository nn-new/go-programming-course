package privilege

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Privilege struct {
	ID               primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Image            string             `bson:"image" json:"image" validate:"required"`
	Title            string             `bson:"title" json:"title" validate:"required"`
	CatalogId        string             `bson:"catalog_id" json:"catalogId"`
	CategoryId       string             `bson:"category_id" json:"categoryId"`
	StartDate        time.Time          `bson:"start_date" json:"startDate"`
	EndDate          time.Time          `bson:"end_date" json:"endDate"`
	TypeId           string             `bson:"type_id" json:"typeId"`
	Point            int                `bson:"point" json:"point"`
	Eligibility      []string           `bson:"eligibility" json:"eligibility"`
	Active           bool               `bson:"active" json:"active"`
	PrivilegeId      string             `bson:"privilege_id" json:"privilegeId"`
	PromotionText    string             `bson:"promotion_text" json:"promotionText"`
	DiscountText     string             `bson:"discount_text" json:"discountText"`
	ShortDescription string             `bson:"short_description" json:"shortDescription"`
	FullDescription  string             `bson:"full_description" json:"fullDescription"`
	BrachInformation string             `bson:"brach_information" json:"brachInformation"`
	Condition        string             `bson:"condition" json:"condition"`
	BrandTitle       string             `bson:"brand_title" json:"brandTitle"`
	BrandText        string             `bson:"brand_text" json:"brandText"`
}
