package log

import (
	"database/sql"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
)

func authenticateUser(username, password string) (bool, error) {
	var user swagger.UserData
	connStr := "user=orange dbname=user_data sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err == nil {
		err1 := db.QueryRow("SELECT * FROM users WHERE username = $1 AND password = $2", username, password).Scan(
			&user.UserId, &user.UserName, &user.UserEmail, &user.Password)

		if err1 == sql.ErrNoRows {
			return false, nil // 未找到用戶
		} else if err1 != nil {
			return false, err // 其他錯誤
		}

		return true, nil // 認證成功
	} else {
		return false, err
	}

}
