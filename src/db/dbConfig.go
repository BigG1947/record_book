package db

type Config struct {
	User     string `json:"user"`
	Password string `json:"password"`
	Protocol string `json:"protocol"`
	Host     string `json:"host"`
	Port     int    `json:"port"`
	DbName   string `json:"db_name"`
}

var DefaultConfig = Config{
	User: "root",
	Password: "",
	Protocol: "tcp",
	Host: "127.0.0.1",
	Port:  3306,
	DbName: "record_book",
}

var TestConfig = Config{
	User: "root",
	Password: "",
	Protocol: "tcp",
	Host: "127.0.0.1",
	Port:  3306,
	DbName: "test_record_book",
}