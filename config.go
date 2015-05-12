package ledisdbcli

type Config struct {
	Addr         string
	Password     string
	MaxIdleConns int
	DBIndex      int
}
