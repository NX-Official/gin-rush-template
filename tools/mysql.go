package tools

import (
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func IsDuplicateKeyError(err error) bool {
	var mysqlErr *mysql.MySQLError
	if errors.As(err, &mysqlErr) {
		return mysqlErr.Number == 1062
	}
	return false
}
