package main

import (
	"database/sql"
	"fmt"
	"net/mail"
	"regexp"
	"strings"
	"time"

	_ "github.com/lib/pq"
)

// 手機號碼格式
func validPhoneNumber(phoneNumber string) bool {
	regex := `^8869[0-9]{8}$|^09[0-9]{8}$`
	match, _ := regexp.MatchString(regex, phoneNumber)
	return match

}

// email格式
func validMailAddress(address string) (string, bool) {
	addr, err := mail.ParseAddress(address)
	if err != nil {
		return "", false
	}
	return addr.Address, true
}

// 找重複email 有重複false
func searchMailAddress(db *sql.DB, address string) bool {
	rows, err := db.Query("SELECT id, Email FROM UserData")
	if err != nil {
		fmt.Println("Error running query:", err)
		return false
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var email string
		if err := rows.Scan(&id, &email); err != nil {
			fmt.Println("Error scanning row:", err)
			return false
		}

		if strings.EqualFold(address, email) {
			fmt.Println("This Email has been registered")
			return false
		}
	}

	return true
}

func main() {
	// 建立数据库连接
	connStr := "user=orange dbname=user_data sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		fmt.Println("Error opening database:", err)
		return
	}
	defer db.Close()
	type UserData struct {
		Name, Account, Address, PhoneNumber, Password string
		Authority                                     int

		Birthday time.Time
	}
	//stirng2Time
	date, err := time.Parse("2006/01/02", "2002/12/17")
	if err != nil {
		fmt.Println("日期解析错误:", err)
		return
	}
	data := UserData{
		Name:        "XDD",
		Account:     "k98006@gmail.com",
		Address:     "Taiwan",
		Birthday:    date,
		PhoneNumber: "0912455678",
		Password:    "orange",
		Authority:   0,
	}

	// db.QueryRow("SELECT COUNT(*) FROM users").Scan(&_usercount)
	if validPhoneNumber(data.PhoneNumber) {
		fmt.Println("Correct PhoneNumber")
	} else {
		fmt.Println("Wrong PhoneNumber")
	}

	if addr, ok := validMailAddress(data.Account); ok {
		fmt.Printf("value: %-30s valid email: %-10t address: %s\n", data.Account, ok, addr)
	} else {
		fmt.Printf("value: %-30s valid email: %-10t\n", data.Account, ok)
	}
	//搜尋
	_TFwrite := searchMailAddress(db, data.Account)
	if _TFwrite {
		fmt.Println("Can be register :D")
	}
	// 开始事务
	tx, err := db.Begin()
	if err != nil {
		fmt.Println("Error starting transaction:", err)
		return
	}

	// 执行插入操作
	if _TFwrite {
		_, err = tx.Exec("INSERT INTO UserData (Name, Password, Email, Address, PhoneNumber, Birthday, Authority) VALUES ($1, $2, $3, $4, $5, $6, $7)", data.Name, data.Password, data.Account, data.Address, data.PhoneNumber, data.Birthday, data.Authority)

		if err != nil {
			fmt.Println("Error inserting data:", err)
			tx.Rollback()
			return
		}
		// 提交事务
		err = tx.Commit()
		if err != nil {
			fmt.Println("Error committing transaction:", err)
			return
		}
	}

}
