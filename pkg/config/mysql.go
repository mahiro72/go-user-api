package config

import "fmt"

var MySQL mySQL

type mySQL struct {
	_DATABASE      string
	_USER          string
	_PASSWORD      string
	_ROOT_PASSWORD string
	_HOST          string
	DSN            string
}

func initMySQL() error {
	var err error
	if MySQL._DATABASE, err = getEnv("MYSQL_DATABASE"); err != nil {
		return err
	}

	if MySQL._USER, err = getEnv("MYSQL_USER"); err != nil {
		return err
	}

	if MySQL._PASSWORD, err = getEnv("MYSQL_PASSWORD"); err != nil {
		return err
	}

	if MySQL._ROOT_PASSWORD, err = getEnv("MYSQL_ROOT_PASSWORD"); err != nil {
		return err
	}

	if MySQL._HOST, err = getEnv("MYSQL_HOST"); err != nil {
		return err
	}
	
	MySQL.DSN = fmt.Sprintf(
		"%s:%s@tcp(%s:3306)/%s?charset=utf8mb4&parseTime=true",
		MySQL._USER,
		MySQL._PASSWORD,
		MySQL._HOST,
		MySQL._DATABASE,
	)
	return nil
}
