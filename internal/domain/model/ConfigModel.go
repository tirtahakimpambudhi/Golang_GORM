package model

type Server struct {
	Protocol string `config:"Server.Protocol"`
	Host     string `config:"Server.Host"`
	Port     string `config:"Server.Port"`
	Endpoint string `config:"Server.Endpoint"`
}

func NewServer(protocol string, host string, port string, endpoint string) *Server {
	return &Server{Protocol: protocol, Host: host, Port: port, Endpoint: endpoint}
}

type Database struct {
	Dialect  string `config:"Database.Dialect"`
	Host     string `config:"Database.Host"`
	Port     string `config:"Database.Port"`
	User     string `config:"Database.User"`
	Password string `config:"Database.Password"`
	Name     string `config:"Database.Name"`
	Table    string `config:"Database.Table"`
}

func NewDatabase(dialect string, host string, port string, user string, password string, name string, table string) *Database {
	return &Database{Dialect: dialect, Host: host, Port: port, User: user, Password: password, Name: name, Table: table}
}
