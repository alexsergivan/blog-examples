package repository

// Repository interface to handle users data.
type Repository interface {
	GetUsers() map[int]string
}