package register

import (
	"database/sql"
	"errors"
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
	"github.com/go-playground/validator/v10"

	_ "github.com/lib/pq"
)

func init() {

}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func ValidateE164(fl validator.FieldLevel) bool {
	// 使用正則表達式確認是否符合 E.164 格式
	e164Pattern := `^(\+8869[0-9]{8}|09[0-9]{8})$`
	return regexp.MustCompile(e164Pattern).MatchString(fl.Field().String())
}

type PhoneNumber struct {
	Number string `validate:"required,phone_number"`
}

// 只少符合三種(數字、符號、英文大寫、英文小寫)
func PasswordCheck(passwd string) error {
	if len(passwd) < 6 {
		return errors.New("password too short")
	}

	indNum := [4]bool{false, false, false, false}
	spCode := "!@#$%^&*_-"

	for _, char := range passwd {
		switch {
		case 'A' <= char && char <= 'Z':
			indNum[0] = true
		case 'a' <= char && char <= 'z':
			indNum[1] = true
		case '0' <= char && char <= '9':
			indNum[2] = true
		case strings.ContainsRune(spCode, char):
			indNum[3] = true
		default:
			return errors.New("Unsupported character")
		}
	}

	codeCount := 0
	for _, i := range indNum {
		if i {
			codeCount++
		}
	}

	if codeCount < 3 {
		return errors.New("Too simple password")
	}

	return nil
}

// 手機號碼格式
func validPhoneNumber(phoneNumber string) error {
	v := validator.New()
	v.RegisterValidation("phone_number", ValidateE164)
	cv := &CustomValidator{validator: v}
	Number := PhoneNumber{Number: phoneNumber}
	err := cv.Validate(Number)
	return err
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
func DecideWrite(tmp [4]bool) bool {

	var ans bool = true
	for i := range tmp {
		ans = ans && tmp[i]
	}
	return ans
}

var data = []swagger.UserData{
	{UserId: 1,
		UserName:  "orange",
		UserEmail: "orange@gmail.com",
		Authority: 1,
		Password:  "123",
		Address:   "Taipei",
		Phone:     "0912345678",
		Status:    1,
		CartId:    1},
	{UserId: 2,
		UserName:  "apple",
		UserEmail: "apple@gmail.com",
		Authority: 0,
		Password:  "orangeOAO123",
		Address:   "Taipei",
		Phone:     "0912345628",
		Status:    1,
		CartId:    1},
}

func WriteInDB(db *sql.DB, data swagger.UserData) {
	Birthday, _ := time.Parse("2006-01-02", data.Birthday)
	_, err := db.Exec("INSERT INTO user_data (name, password, email, address, phone_number, birthday, authority, current_cart_id, status) VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)", data.UserName, data.Password, data.UserEmail, data.Address, data.Phone, Birthday, data.Authority, data.CartId, data.Status)
	if err != nil {
		fmt.Println(err)
	}
}

func UpdateInDB(emailChanged string, data swagger.UserData) error {
	connStr := "user=orange dbname=user_data sslmode=disable"
	Birthday, _ := time.Parse("2006-01-02", data.Birthday)
	if db, err := sql.Open("postgres", connStr); err == nil {
		_, err1 := db.Exec("UPDATE user_data SET name=$1, password=$2, email=$3, address=$4, phone_number=$5, birthday=$6, authority=$7, current_cart_id=$8, status=$9 WHERE email=$10", data.UserName, data.Password, data.UserEmail, data.Address, data.Phone, Birthday, data.Authority, data.CartId, data.Status, emailChanged)
		if err1 != nil {
			fmt.Println(err1)
			return err1
		}

	} else {
		return err
	}
	return nil
}

func Register() {
	connStr := "user=orange dbname=user_data sslmode=disable"
	var _allErr [3]error
	if db, err := sql.Open("postgres", connStr); err == nil {
		_flag := true
		_allErr[0] = PasswordCheck(data[1].Password)
		_allErr[1] = validPhoneNumber(data[1].Phone)
		_allErr[2] = searchRepeat(db, data[1].UserEmail, data[1].Phone)
		for i := 0; i < 3; i++ {
			if _allErr[i] != nil {
				fmt.Println(_allErr[i])
				_flag = false
				break
			}
		}
		if _flag {
			WriteInDB(db, data[1])
			fmt.Println("OAO")
		}
		defer db.Close()

	}

}

func UpdateData() {
	var tempData = []swagger.UserData{
		{UserId: 1,
			UserName:  "orange",
			UserEmail: "orange@gmail.com",
			Authority: 1,
			Password:  "Temp123Password123",
			Address:   "Taipei",
			Phone:     "0912345678",
			Status:    1,
			CartId:    1},
		{UserId: 2,
			UserName:  "apple",
			UserEmail: "apple@gmail.com",
			Authority: 0,
			Password:  "orangeOAO123",
			Address:   "Taipei",
			Phone:     "0912345628",
			Status:    1,
			CartId:    1},
	}
	UpdateInDB("orange@gmail.com", tempData[0])

}
