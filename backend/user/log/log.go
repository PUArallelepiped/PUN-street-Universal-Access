package log

import (
	"database/sql"
	"fmt"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
	_ "github.com/lib/pq"
)

func LogIn(username, password string) (bool, error) {
	var user swagger.UserData
	connStr := "user=orange dbname=user_data sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err == nil {

		err1 := db.QueryRow("SELECT name FROM user_data WHERE name = $1", username).Scan(&user.UserName)
		err2 := db.QueryRow("SELECT name, password FROM user_data WHERE name = $1 AND password =$2", username, password).Scan(&user.UserName, &user.Password)
		if err1 == sql.ErrNoRows {
			fmt.Println("未找到用戶")
			return false, nil // 未找到用戶
		} else if err2 == sql.ErrNoRows {
			fmt.Println("密碼錯誤")
			return false, nil //密碼錯誤

		} else if err1 != nil {

			return false, err // 其他錯誤
		}
		fmt.Println("GOOD")
		return true, nil // 認證成功
	} else {

		return false, err
	}

}
