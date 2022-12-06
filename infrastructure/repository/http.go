package repository

type HTTPRepository interface {
	Get(string, map[string]string, map[string]string) ([]byte, error)
}
