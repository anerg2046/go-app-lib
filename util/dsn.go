package util

import "fmt"

func MysqlDSN(host, user, pass, port, dbname string) string {
	return fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", user, pass, host, port, dbname)
}

func PostgresDSN(host, user, pass, port, dbname string) string {
	return fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Shanghai", host, user, pass, dbname, port)
}

func MssqlDSN(host, user, pass, port, dbname string) string {
	return fmt.Sprintf("sqlserver://%s:%s@%s:%s?database=%s", user, pass, host, port, dbname)
}

func SqliteDSN(path string) string {
	return path
}
