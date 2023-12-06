package repository

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
	_ "github.com/lib/pq"
	"github.com/sirupsen/logrus"
)

type postgresqlUserRepo struct {
	db *sql.DB
}

func NewPostgressqlUserRepo(db *sql.DB) domain.UserRepo {
	return &postgresqlUserRepo{db}
}

func (p *postgresqlUserRepo) GetByID(ctx context.Context, id string) (*swagger.UserData, error) {
	row := p.db.QueryRow("SELECT * FROM user_data WHERE user_id = $1", id)
	s := &swagger.UserData{}
	if err := row.Scan(&s.UserId, &s.UserName, &s.Password, &s.UserEmail, &s.Address, &s.Phone, &s.Birthday, &s.Authority, &s.CartId, &s.Status); err != nil {
		logrus.Error(err)
		return nil, err
	}
	return s, nil
}

func (p *postgresqlUserRepo) GetAllUser(ctx context.Context) ([]swagger.UserData, error) {
	rows, err := p.db.Query("SELECT * FROM user_data")
	if err != nil {
		logrus.Error(err)
	}
	l := []swagger.UserData{}
	for rows.Next() {
		s := swagger.UserData{}
		err := rows.Scan(&s.UserId, &s.UserName, &s.Password, &s.UserEmail, &s.Address, &s.Phone, &s.Birthday, &s.Authority, &s.CartId, &s.Status)
		if err != nil {
			logrus.Error(err)
		}
		l = append(l, s)
	}

	return l, nil
}

func (p *postgresqlUserRepo) LogIn(ctx context.Context, username, password string) (bool, error) {
	var user swagger.UserData

	err1 := p.db.QueryRow("SELECT name FROM user_data WHERE name = $1", username).Scan(&user.UserName)
	err2 := p.db.QueryRow("SELECT name, password FROM user_data WHERE name = $1 AND password =$2", username, password).Scan(&user.UserName, &user.Password)
	if err1 == sql.ErrNoRows {
		return false, nil // 未找到用戶
	} else if err2 == sql.ErrNoRows {
		return false, nil //密碼錯誤

	} else if err1 != nil {

		return false, err1 // 其他錯誤
	}
	fmt.Println("GOOD")
	return true, nil // 認證成功

}
func writeInDB(db *sql.DB, data swagger.UserData) {
	Birthday, _ := time.Parse("2006-01-02", data.Birthday)
	_, err := db.Exec("INSERT INTO user_data (name, password, email, address, phone_number, birthday, authority, current_cart_id, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", data.UserName, data.Password, data.UserEmail, data.Address, data.Phone, Birthday, data.Authority, data.CartId, data.Status)
	if err != nil {
		fmt.Println(err)
	}
}

// 找重複email 有重複false
func searchRepeat(db *sql.DB, address string, phoneNumber string) error {
	rows, err := db.Query("SELECT user_id,phone_number, email FROM user_data")
	if err != nil {
		fmt.Println("Error running query:", err)
		return err
	}
	defer rows.Close()

	for rows.Next() {
		var id int
		var email string
		var phone_number string
		if err := rows.Scan(&id, &phone_number, &email); err != nil {
			return errors.New("Error scanning row")
		}

		if strings.EqualFold(address, email) {
			return errors.New("This Email has been registered")
		}
		if strings.EqualFold(phoneNumber, phone_number) {
			return errors.New("This Phone_number has been registered")
		}
	}

	return nil
}

func updateInDB(emailChanged string, data swagger.UserData) error {

	connStr := "user=orange dbname=user_data sslmode=disable"

	Birthday, _ := time.Parse("2006-01-02", data.Birthday)
	if db, err := sql.Open("postgres", connStr); err == nil {
		if err1 := searchRepeat(db, emailChanged, ""); err1 == nil {
			_, err2 := db.Exec("UPDATE user_data SET name=$1, password=$2, email=$3, address=$4, phone_number=$5, birthday=$6, authority=$7, current_cart_id=$8, status=$9 WHERE email=$10", data.UserName, data.Password, data.UserEmail, data.Address, data.Phone, Birthday, data.Authority, data.CartId, data.Status, emailChanged)
			if err2 != nil {
				fmt.Println(err2)
				return err2
			}
		} else {
			fmt.Println("ERROR")
			fmt.Println(err1)
		}

	} else {
		return err
	}
	return nil
}

func RegisterData(data swagger.UserData) {
	var result string
	connStr := "user=orange dbname=user_data sslmode=disable"
	if db, err := sql.Open("postgres", connStr); err == nil {

		db.QueryRow("SELECT is_valid_password($1, $2, $3)", data.Password, 6, 3).Scan(&result)
		fmt.Println(result)
		defer db.Close()

	}

}

func UpdateData(data swagger.UserData) {
	// var tempData = []swagger.UserData{
	// 	{UserId: 1,
	// 		UserName:  "orange",
	// 		UserEmail: "orange@gmail.com",
	// 		Authority: 1,
	// 		Password:  "Temp123Password123",
	// 		Address:   "Taipei",
	// 		Phone:     "0912345678",
	// 		Status:    1,
	// 		CartId:    1},
	// 	{UserId: 2,
	// 		UserName:  "apple",
	// 		UserEmail: "apple@gmail.com",
	// 		Authority: 0,
	// 		Password:  "orangeOAO123",
	// 		Address:   "Taipei",
	// 		Phone:     "0912345628",
	// 		Status:    1,
	// 		CartId:    1},
	// }
	updateInDB("orange@gmail.com", data)

}
