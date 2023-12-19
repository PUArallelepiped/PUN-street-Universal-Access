package usecase

import (
	"context"
	"errors"
	"time"

	"github.com/PUArallelepiped/PUN-street-Universal-Access/domain"
	"github.com/PUArallelepiped/PUN-street-Universal-Access/swagger"
	"github.com/golang-jwt/jwt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type UserUsecase struct {
	userRepo domain.UserRepo
}

func NewUserUsecase(userRepo domain.UserRepo) domain.UserUsecase {
	return &UserUsecase{
		userRepo: userRepo,
	}
}

type Claims struct {
	Email     string `json:"email"`
	Authority string `json:"authority"`
	jwt.StandardClaims
}

func init() {
	viper.SetConfigFile("../.env")
	viper.SetConfigType("dotenv")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatal(err)
	}
}

func CreateToken(email string, authority string) (string, error) {
	expiresAt := time.Now().Add(24 * time.Hour).Unix()
	issuedAt := time.Now().Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS512, Claims{
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: expiresAt,
			IssuedAt:  issuedAt,
		},
		Email:     email,
		Authority: authority,
	})
	jwtSecret := []byte(viper.GetString("JWT_SECRET"))
	tokenString, err := token.SignedString(jwtSecret)
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func (uu *UserUsecase) Login(ctx context.Context, email string, password string) (string, error) {
	// check if user password is correct
	authority, err := uu.userRepo.Login(ctx, email, password)
	if err != nil {
		return "", err
	}
	// check if user is banned
	isBanned, err := uu.userRepo.IsUserBanned(ctx, email)
	if err != nil {
		return "", err
	}
	if isBanned {
		return "", errors.New("banned")
	}
	token, err := CreateToken(email, authority)
	if err != nil {
		return "", err
	}

	return token, nil
}

func (uu *UserUsecase) ValidateToken(ctx context.Context, token string) error {
	jwtSecret := []byte(viper.GetString("JWT_SECRET"))
	tokenClaims, err := jwt.ParseWithClaims(token, &Claims{}, func(token *jwt.Token) (interface{}, error) {
		return jwtSecret, nil
	})
	_ = tokenClaims
	if err != nil {
		ve, ok := err.(*jwt.ValidationError)
		if ok {
			if ve.Errors&jwt.ValidationErrorMalformed != 0 {
				logrus.Error("That's not even a token")
				return err
			} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
				logrus.Error("Timing is everything")
				return err
			} else {
				logrus.Error("Couldn't handle this token:", err)
				return err
			}
		} else {
			logrus.Error("Couldn't handle this token:", err)
			return err
		}
	}
	return nil
}

func (su *UserUsecase) GetByID(ctx context.Context, id int64) (*swagger.UserData, error) {
	s, err := su.userRepo.GetByID(ctx, id)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return s, nil
}

func (su *UserUsecase) GetAllUser(ctx context.Context) ([]swagger.UserDataShort, error) {
	s, err := su.userRepo.GetAllUser(ctx)
	if err != nil {
		logrus.Error(err)
		return nil, err
	}
	return s, nil
}

func (su *UserUsecase) RegisterUser(ctx context.Context, user *swagger.RegisterInfo) error {
	if user.StoreRegisterInfo != nil {
		// register store
		logrus.Info("register store")
		id, err := su.userRepo.RegisterUser(ctx, user, "011")
		if err != nil {
			logrus.Error(err)
			return err
		}
		err = su.userRepo.RegisterStore(ctx, *user.StoreRegisterInfo, id)
		if err != nil {
			logrus.Error(err)
			return err
		}
	} else {
		// register user
		logrus.Info("register user")
		_, err := su.userRepo.RegisterUser(ctx, user, "001")
		if err != nil {
			logrus.Error(err)
			return err
		}
	}
	return nil
}

func (su *UserUsecase) CheckEmail(ctx context.Context, email string) (bool, error) {
	return su.userRepo.CheckEmail(ctx, email)
}
