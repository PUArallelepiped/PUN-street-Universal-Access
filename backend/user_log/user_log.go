package main

import (
    "database/sql"
    "strings"
    "fmt"
    _ "github.com/lib/pq"
    "net/mail"
)

func validMailAddress(address string) (string, bool) {
    addr, err := mail.ParseAddress(address)
    if err != nil {
    return "", false
    }
    return addr.Address, true
}
//找重複email 有重複false
func searchMailAddress(db *sql.DB, address string) bool {
    rows, err := db.Query("SELECT id, email FROM UserData")
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
        if strings.EqualFold(email, address) {
            fmt.Println("This Email has been registered")
            return false
        }
    }

    return true
}

func main() {
    // 建立数据库连接
    connStr := "user=postgres dbname=postgres sslmode=disable"
    db, err := sql.Open("postgres", connStr)
    if err != nil {
        fmt.Println("Error opening database:", err)
        return
    }
    defer db.Close()

	var _name string
    var _account string
	var _address string
	var _birthday string
	var _phoneNumber string
    var _passwd string
	var _authority int
	_birthday = "2002/12/12"
	_name = "orange"
    _account = "k98006@gmail.com"
	_address ="Taiwan"
	_phoneNumber ="091232242"
    _passwd = "orange"
    _authority =0;
    // db.QueryRow("SELECT COUNT(*) FROM users").Scan(&_usercount)
   
    if addr, ok := validMailAddress(_account); ok {
        fmt.Printf("value: %-30s valid email: %-10t address: %s\n", _account, ok, addr)
    } else {
        fmt.Printf("value: %-30s valid email: %-10t\n", _account, ok)
    }
    //搜尋
    _TFwrite:=searchMailAddress(db,_account);
    if _TFwrite{
        fmt.Println("Can be register :D");
    }
    // 开始事务
    tx, err := db.Begin()
    if err != nil {
        fmt.Println("Error starting transaction:", err)
        return
    }
    
    // 执行插入操作
    if _TFwrite {
		_, err = tx.Exec("INSERT INTO UserData (Name, Password, Email, Address, PhoneNumber, Birthday, Authority) VALUES ($1, $2, $3, $4, $5, $6, $7)", _name, _passwd, _account, _address, _phoneNumber, _birthday, _authority)

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
