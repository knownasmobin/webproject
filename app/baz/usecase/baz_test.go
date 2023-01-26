package usecase

import (
	"context"
	"os"
	"testing"

	bootDB "git.ecobin.ir/ecomicro/bootstrap/db"
	bootStoreRedis "git.ecobin.ir/ecomicro/bootstrap/store/redis"
	"git.ecobin.ir/ecomicro/template/app/baz/domain"
	repository "git.ecobin.ir/ecomicro/template/app/baz/repository/psql"
	redisCache "git.ecobin.ir/ecomicro/template/app/baz/repository/redis"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func insertBazToDB(ctx context.Context, repo domain.Repository, baz domain.Baz) (*domain.Baz, error) {
	return repo.Create(ctx, baz)
}

type repositoryMock struct {
	db *gorm.DB
	domain.Repository
	mock.Mock
}

var _ domain.Repository = &repositoryMock{}

func (m *repositoryMock) Create(ctx context.Context, Baz domain.Baz) (*domain.Baz, error) {
	args := m.Called(ctx, Baz)
	return args.Get(0).(*domain.Baz), args.Error(1)
}

func (m *repositoryMock) Update(ctx context.Context, condition domain.Baz, Baz domain.Baz) ([]domain.Baz, error) {
	args := m.Called(ctx, condition, Baz)
	return args.Get(0).([]domain.Baz), args.Error(1)
}

func (m *repositoryMock) Get(ctx context.Context, id uint64) (*domain.Baz, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.Baz), args.Error(1)
}

// Utils \\
type errorAssertion struct {
	hasErr  bool
	errType error
}

func assertErr(t *testing.T, got error, want errorAssertion) {
	if want.hasErr {
		assert.Error(t, got)
	} else {
		assert.NoError(t, got)
	}
	if want.errType != nil {
		assert.ErrorIs(t, got, want.errType)
	}
}

func assertCreatedBaz(t *testing.T, got *domain.Baz, want domain.Baz) {
	assert.NotNil(t, got, "Baz")
	assert.Equal(t, want.UserId, got.UserId, "UserId")
}

func newUsecase() (*usecase, domain.Repository, *gorm.DB, domain.Repository) {
	os.Remove("gorm.db")
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	redis, err := bootStoreRedis.New(bootStoreRedis.Config{
		Addr: "0.0.0.0:6363",
	})
	if err != nil {
		panic(err)
	}
	autoMigrate := true
	dataBase := &bootDB.Database{
		Config: bootDB.DatabaseConfiguration{
			AutoMigrate: &autoMigrate,
		},
		GormDB: db,
	}
	r := repository.NewRepository(dataBase.GormDB)
	redisCacheRepo := redisCache.New(r, redis)

	return &usecase{
		bazRepo: &repositoryMock{},
	}, r, db, redisCacheRepo
}

// Testing \\
func TestMain(m *testing.M) {
	res := m.Run()
	os.Remove("gorm.db")
	os.Exit(res)
}
func Test_usecase_Create(t *testing.T) {
	u := &usecase{}
	var err error
	var userId uint64 = 654654654
	ctx := context.Background()

	type args struct {
		ctx context.Context
		Baz domain.Baz
	}
	type assertionArgs struct {
		Baz domain.Baz
		errorAssertion
	}
	var tests = []struct {
		name    string
		arrange func(t *testing.T)
		args    args
		want    func() assertionArgs
		assert  func(t *testing.T, Baz *domain.Baz, got error, want assertionArgs)
	}{
		{
			name: "duplicate_id",
			arrange: func(t *testing.T) {
				usecase, repository, _, _ := newUsecase()
				u = usecase
				u.bazRepo = repository
				_, err = insertBazToDB(ctx, repository, domain.Baz{
					UserId: userId,
				})
				if err != nil {
					t.Fatal(err)
				}
			},
			args: args{
				ctx: ctx,
				Baz: domain.Baz{
					UserId: userId,
				},
			},
			want: func() assertionArgs {
				return assertionArgs{
					errorAssertion: errorAssertion{
						hasErr: true,
					}}
			},
			assert: func(t *testing.T, Baz *domain.Baz, err error, want assertionArgs) {
				assertErr(t, err, want.errorAssertion)
			},
		},
		{
			name: "successful",
			arrange: func(t *testing.T) {
				usecase, repository, _, _ := newUsecase()
				u = usecase
				u.bazRepo = repository
			},
			args: args{
				ctx: ctx,
				Baz: domain.Baz{
					UserId: userId,
				},
			},
			want: func() assertionArgs {
				return assertionArgs{
					errorAssertion: errorAssertion{
						hasErr: false,
					},
					Baz: domain.Baz{
						UserId: userId,
					}}
			},
			assert: func(t *testing.T, Baz *domain.Baz, err error, want assertionArgs) {
				assertErr(t, err, want.errorAssertion)
				assertCreatedBaz(t, Baz, want.Baz)
			},
		},
		{
			name: "duplicate id with redis cache",
			arrange: func(t *testing.T) {
				usecase, _, _, redisCache := newUsecase()
				u = usecase
				u.bazRepo = redisCache
				_, err = insertBazToDB(ctx, redisCache, domain.Baz{
					UserId: userId,
				})
				if err != nil {
					t.Fatal(err)
				}
			},
			args: args{
				ctx: ctx,
				Baz: domain.Baz{
					UserId: userId,
				},
			},
			want: func() assertionArgs {
				return assertionArgs{
					errorAssertion: errorAssertion{
						hasErr: true,
					}}
			},
			assert: func(t *testing.T, Baz *domain.Baz, err error, want assertionArgs) {
				assertErr(t, err, want.errorAssertion)
			},
		},
		{
			name: "successful with redis cache",
			arrange: func(t *testing.T) {
				usecase, _, _, redisCache := newUsecase()
				u = usecase
				u.bazRepo = redisCache
			},
			args: args{
				ctx: ctx,
				Baz: domain.Baz{
					UserId: userId,
				},
			},
			want: func() assertionArgs {
				return assertionArgs{
					errorAssertion: errorAssertion{
						hasErr: false,
					},
					Baz: domain.Baz{
						UserId: userId,
					}}
			},
			assert: func(t *testing.T, Baz *domain.Baz, err error, want assertionArgs) {
				assertErr(t, err, want.errorAssertion)
				assertCreatedBaz(t, Baz, want.Baz)
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Arrange
			if tt.arrange != nil {
				tt.arrange(t)
			}
			// Act
			dbBaz, err := u.Create(tt.args.ctx, tt.args.Baz)
			// Assert
			if tt.assert != nil {
				tt.assert(t, dbBaz, err, tt.want())
			}

		})
	}
}
