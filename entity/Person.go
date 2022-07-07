package entity

type Person struct{
	ID int64 `json:"id"`
	FirstName string `json:"firstName"`
	LastName string `json:"lastName"`
	Street string `json:"street"`
	City string `json:"city"`
	State string `json:"state"`
	PostalCode string `json:"postalCode"`
}