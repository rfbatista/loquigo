package domain

type User struct {
	ID         string `bson:"_id"`
	ExternalId string `json:"id" bson:"external_id"`
}
