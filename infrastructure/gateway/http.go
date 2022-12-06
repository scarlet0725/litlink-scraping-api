package gateway

type HTTP interface {
	Get(string, map[string]string, map[string]string) ([]byte, error)
}
