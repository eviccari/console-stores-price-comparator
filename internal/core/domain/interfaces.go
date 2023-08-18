package domain

type WebStore interface {
	GetItems() (items []interface{}, err error)
}
