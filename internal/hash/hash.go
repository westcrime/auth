package hash

type HashService interface {
	HashPassword(password string) (string, error)
	CheckPassword(password, hash string) bool
}
