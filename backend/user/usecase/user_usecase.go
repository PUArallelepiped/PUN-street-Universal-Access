package usecase

import (
	"context"
	"errors"
	"regexp"
	"strings"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
	"github.com/go-playground/validator/v10"
	"github.com/sirupsen/logrus"

	_ "github.com/lib/pq"
)

type UserUsecase struct {
	userRepo domain.UserRepo
}

func NewUserUsecase(userRepo domain.UserRepo) domain.UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

func (su *UserUsecase) GetByID(ctx context.Context, id string) (*swagger.UserData, error) {
	s, err := su.userRepo.GetByID(ctx, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return s, nil
}

func (su *UserUsecase) GetAllUser(ctx context.Context) ([]swagger.UserData, error) {
	s, err := su.userRepo.GetAllUser(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return s, nil
}

type CustomValidator struct {
	validator *validator.Validate
}

func (cv *CustomValidator) Validate(i interface{}) error {
	return cv.validator.Struct(i)
}

func validateE164(fl validator.FieldLevel) bool {
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
func ValidPhoneNumber(phoneNumber string) error {
	v := validator.New()
	v.RegisterValidation("phone_number", validateE164)
	cv := &CustomValidator{validator: v}
	Number := PhoneNumber{Number: phoneNumber}
	err := cv.Validate(Number)
	return err
}

func (su *UserUsecase) LogIn(ctx context.Context, userName string, userPassword string) (bool, error) {
	correctLogIn, err := su.userRepo.LogIn(ctx, userName, userPassword)
	return correctLogIn, err
}
