package storage

type JwtStorage interface {
	Store(bearerToken string, jwt string) error
	Get(bearerToken string) (string, error)
}
