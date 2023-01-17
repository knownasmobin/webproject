package redis

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"git.ecobin.ir/ecomicro/bootstrap/store/redis"
	"git.ecobin.ir/ecomicro/template/app/baz/domain"
	"git.ecobin.ir/ecomicro/tooty"
	"git.ecobin.ir/ecomicro/x"
	gredis "github.com/go-redis/redis/v9"
)

type repository struct {
	r   domain.Repository
	rdb *redis.Redis
}

var _ domain.Repository = &repository{}

func New(r domain.Repository, rdb *redis.Redis) *repository {
	return &repository{r: r, rdb: rdb}
}

func (u *repository) persist(ctx context.Context, baz domain.Baz) {
	d, err := json.Marshal(baz)
	if err != nil {
		x.LogError(err, ctx)
	} else {
		err = u.rdb.Set(ctx, fmt.Sprint(baz.UserId), string(d), time.Hour).Err()
		if err != nil {
			x.LogError(err, ctx)
		}
	}
}

func (u *repository) Create(ctx context.Context, baz domain.Baz) (*domain.Baz, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[C] [R] Create baz", "cache")
	defer tooty.CloseTheAPMSpan(span)

	createdUser, err := u.r.Create(ctx, baz)
	if err != nil {
		return nil, err
	}
	u.persist(ctx, baz)
	return createdUser, err
}

func (u *repository) Get(ctx context.Context, userId uint64) (*domain.Baz, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[C] [R] Get user", "cache")
	defer tooty.CloseTheAPMSpan(span)

	rawBaz, err := u.rdb.Get(ctx, fmt.Sprint(userId)).Result()
	if err == nil {
		var baz domain.Baz
		err = json.Unmarshal([]byte(rawBaz), &baz)
		if err == nil {
			return &baz, nil
		} else {
			x.LogError(err, ctx)
		}
	} else {
		if !errors.Is(err, gredis.Nil) {
			x.LogError(err, ctx)
		}
	}

	user, err := u.r.Get(ctx, userId)
	if user != nil {
		u.persist(ctx, *user)
	}
	return user, err
}

func (u *repository) Update(ctx context.Context, condition, baz domain.Baz) ([]domain.Baz, error) {
	span := tooty.OpenAnAPMSpan(ctx, "[C] [R] update baz", "cache")
	defer tooty.CloseTheAPMSpan(span)

	bazs, err := u.r.Update(ctx, condition, baz)
	if err != nil {
		return nil, err
	}
	for _, baz := range bazs {
		u.persist(ctx, baz)
	}
	return nil, err
}
