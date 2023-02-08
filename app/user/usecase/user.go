package usecase

import (
	"context"
	"time"

	"git.ecobin.ir/ecomicro/template/app/user/domain"
	"git.ecobin.ir/ecomicro/tooty"
	"github.com/golang-jwt/jwt/v4"
)

const jWTSecret = "some jwt secret"

type usecase struct {
	userRepo   domain.Repository
	bazAdapter domain.BazAdapter
}

var _ domain.Usecase = &usecase{}
var _ domain.Adapter = &usecase{}

func NewUserUsecase(userRepo domain.Repository) *usecase {
	return &usecase{
		userRepo: userRepo,
	}
}
func (uu *usecase) SetAdapters(bazAdapter domain.BazAdapter) {
	uu.bazAdapter = bazAdapter
}
func (uu *usecase) Create(
	ctx context.Context,
	user domain.User,
) (*domain.User, error) {

	dbUser, err := uu.userRepo.Create(ctx, user)
	if err != nil {
		return nil, err
	}
	return dbUser, nil
}

func (uu *usecase) Update(ctx context.Context, user domain.User) (*domain.User, error) {
	userArray, err := uu.userRepo.Update(ctx, domain.User{
		Id: user.Id,
	}, user)
	if err != nil {
		return nil, err
	}
	if len(userArray) == 0 {
		return nil, domain.ErrNotFound
	}
	return &userArray[0], nil
}
func (uu *usecase) GetUserById(ctx context.Context, id int) (*domain.User, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[U] get user by id", "usecase")
	defer tooty.CloseTheAPMSpan(span)
	user, err := uu.userRepo.GetUserById(ctx, id)
	if err != nil {
		return nil, err
	}
	return user, nil
}
func (uu *usecase) GetByCondition(ctx context.Context, user domain.User) ([]domain.User, error) {
	dbUser, err := uu.userRepo.GetByCondition(ctx, user)
	if err != nil {
		return nil, err
	}
	return dbUser, nil
}

func (uu *usecase) validateJWTToken(token string, secret []byte) (userId int, err error) {
	claims := &Claims{}
	_, err = jwt.ParseWithClaims(token, claims, func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	if err != nil {
		return 0, err
	}
	return claims.UserId, nil
}

func (uu *usecase) generateJwtToken(userId int, secret []byte, expirationSeconds int) (string, error) {
	expirationTime := time.Now().Add(time.Duration(expirationSeconds) * time.Second)
	// Create the JWT claims, which includes the username and expiry time
	claims := &Claims{
		UserId: userId,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(expirationTime),
		},
	}
	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	return token.SignedString(secret)
}

type Claims struct {
	UserId int `json:"userId"`
	jwt.RegisteredClaims
}

func (uu *usecase) LoginUserCredential(
	ctx context.Context,
	user domain.User,
) (*domain.User, string, error) {

	dbUser, err := uu.userRepo.GetUserByPassword(ctx, user.Username, user.Password)
	if err != nil {
		return nil, "", err
	}
	token, err := uu.generateJwtToken(dbUser.Id, []byte(jWTSecret), 3600) // one hour
	if err != nil {
		return nil, "", err
	}
	return dbUser, token, nil

}

func (uu *usecase) ValidateToken(ctx context.Context, tokenStr string) (*domain.User, error) {

	userId, err := uu.validateJWTToken(tokenStr, []byte(jWTSecret))
	if err != nil {
		return nil, err
	}

	dbUser, err := uu.userRepo.GetUserById(ctx, userId)
	if err != nil {
		if err != domain.ErrNotFound {
			return nil, err
		}
		return nil, err
	}
	return dbUser, nil

}
