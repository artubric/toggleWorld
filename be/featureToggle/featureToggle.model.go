package featuretoggle

import "go.mongodb.org/mongo-driver/bson/primitive"

// FeatureToggle ...
type FeatureToggle struct {
	ID            primitive.ObjectID `json:"id" bson:"_id"`
	IsActive      bool               `json:"isActive" bson:"isActive"`
	DisplayName   string             `json:"displayName" bson:"displayName"`
	TechnicalName string             `json:"technicalName" bson:"technicalName"`
	ExpiresOn     string             `json:"expiresOn" bson:"expiresOn"`
	Description   string             `json:"description" bson:"description"`
	IsInverted    bool               `json:"isInverted" bson:"isInverted"`
	CustomerIds   []string           `json:"customerIds" bson:"customerIds"`
}
