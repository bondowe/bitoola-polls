package models

type Country struct {
	Code  string `bson:"_id"`
	Names map[string]string
}
