package usecase

import (
	"context"
	"log"
	"os"
	"testing"
	"time"

	bootDB "git.ecobin.ir/ecomicro/bootstrap/db"
	"git.ecobin.ir/ecomicro/template/app/user/domain"
	"git.ecobin.ir/ecomicro/template/app/user/repository"
	"git.ecobin.ir/ecomicro/x"
	"github.com/sony/sonyflake"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func insertUserToDB(ctx context.Context, repo domain.Repository, user domain.User) (*domain.User, error) {
	return repo.Create(ctx, user)
}

type repositoryMock struct {
	db *gorm.DB
	domain.Repository
	mock.Mock
}

var _ domain.Repository = &repositoryMock{}

func (m *repositoryMock) Create(ctx context.Context, user domain.User) (*domain.User, error) {
	args := m.Called(ctx, user)
	return args.Get(0).(*domain.User), args.Error(1)
}

func (m *repositoryMock) Update(ctx context.Context, condition domain.User, user domain.User) ([]domain.User, error) {
	args := m.Called(ctx, condition, user)
	return args.Get(0).([]domain.User), args.Error(1)
}

func (m *repositoryMock) GetUserById(ctx context.Context, id uint64) (*domain.User, error) {
	args := m.Called(ctx, id)
	return args.Get(0).(*domain.User), args.Error(1)
}

// foo adapter
type fooAdapterMock struct {
	domain.FooAdapter
	mock.Mock
}

var _ domain.FooAdapter = &fooAdapterMock{}

func (m *fooAdapterMock) Bar(ctx context.Context, user domain.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
}

// baz adapter
type sonyflakeMock struct {
	x.Sonyflake
	mock.Mock
}

var _ x.Sonyflake = &sonyflakeMock{}

func (m *sonyflakeMock) NextID() (uint64, error) {
	args := m.Called()
	return args.Get(0).(uint64), args.Error(1)
}

type bazAdapterMock struct {
	domain.BazAdapter
	mock.Mock
}

var _ domain.BazAdapter = &bazAdapterMock{}

func (m *bazAdapterMock) Create(ctx context.Context, user domain.User) error {
	args := m.Called(ctx, user)
	return args.Error(0)
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

func assertCreatedUser(t *testing.T, got *domain.User, want domain.User) {
	assert.NotNil(t, got, "User")
	assert.ElementsMatch(t, want.Allow, got.Allow, "Allow")
	assert.ElementsMatch(t, want.Deny, got.Deny, "Deny")
	assert.ElementsMatch(t, want.Roles, got.Roles, "Roles")
	assert.Equalf(t, want.Id, got.Id, "Id")
}

func newUsecase() (*usecase, domain.Repository, *gorm.DB) {
	os.Remove("gorm.db")
	db, err := gorm.Open(sqlite.Open("gorm.db"), &gorm.Config{})
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

	sonyflake := sonyflake.NewSonyflake(sonyflake.Settings{StartTime: time.Now()})

	return &usecase{
		userRepo:   r,
		sf:         sonyflake,
		fooAdapter: &fooAdapterMock{},
		bazAdapter: &bazAdapterMock{},
	}, r, db
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
		ctx  context.Context
		user domain.User
	}
	type assertionArgs struct {
		user domain.User
		errorAssertion
	}
	var tests = []struct {
		name    string
		arrange func(t *testing.T)
		args    args
		want    assertionArgs
		assert  func(t *testing.T, user *domain.User, got error, want assertionArgs)
	}{
		{
			name: "duplicate_id",
			arrange: func(t *testing.T) {
				usecase, repository, _ := newUsecase()
				u = usecase

				// mock adapters
				mockFoo := new(fooAdapterMock)
				mockFoo.On("Bar", mock.Anything, mock.Anything).Return(nil)
				u.fooAdapter = mockFoo

				mockBaz := new(bazAdapterMock)
				mockBaz.On("Create", mock.Anything, mock.Anything).Return(nil)
				u.bazAdapter = mockBaz

				mockSony := new(sonyflakeMock)
				mockSony.On("NextID").Return(userId, nil)
				u.sf = mockSony

				_, err = insertUserToDB(ctx, repository, domain.User{
					Id:    userId,
					Roles: make([]string, 0),
					Deny:  make([]string, 0),
					Allow: make([]string, 0),
				})
				if err != nil {
					t.Fatal(err)
				}
			},
			args: args{
				ctx: ctx,
				user: domain.User{
					Roles: []string{"role1", "role2"},
					Allow: []string{"allow1", "allow2"},
					Deny:  []string{"deny1", "deny2"},
				},
			},
			want: assertionArgs{
				errorAssertion: errorAssertion{
					hasErr: true,
				},
			},
			assert: func(t *testing.T, user *domain.User, err error, want assertionArgs) {
				assertErr(t, err, want.errorAssertion)
			},
		},
		{
			name: "successful",
			arrange: func(t *testing.T) {
				usecase, _, _ := newUsecase()
				u = usecase

				// mock adapters
				mockFoo := new(fooAdapterMock)
				mockFoo.On("Bar", mock.Anything, mock.Anything).Return(nil)
				u.fooAdapter = mockFoo

				mockBaz := new(bazAdapterMock)
				mockBaz.On("Create", mock.Anything, mock.Anything).Return(nil)
				u.bazAdapter = mockBaz

				mockSony := new(sonyflakeMock)
				mockSony.On("NextID").Return(userId, nil)
				u.sf = mockSony
			},
			args: args{
				ctx: ctx,
				user: domain.User{
					Roles: []string{"role1", "role2"},
					Allow: []string{"allow1", "allow2"},
					Deny:  []string{"deny1", "deny2"},
				},
			},
			want: assertionArgs{
				errorAssertion: errorAssertion{
					hasErr: false,
				},
				user: domain.User{
					Id:    userId,
					Roles: []string{"role1", "role2"},
					Allow: []string{"allow1", "allow2"},
					Deny:  []string{"deny1", "deny2"},
				},
			},
			assert: func(t *testing.T, user *domain.User, err error, want assertionArgs) {
				assertErr(t, err, want.errorAssertion)
				assertCreatedUser(t, user, want.user)
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
			dbUser, err := u.Create(tt.args.ctx, tt.args.user)
			// Assert
			if tt.assert != nil {
				tt.assert(t, dbUser, err, tt.want)
			}

		})
	}
}

func Test_usecase_Update(t *testing.T) {
	u := &usecase{}
	var err error
	var userId uint64 = 654654654
	ctx := context.Background()

	type args struct {
		ctx  context.Context
		user domain.User
	}
	type assertionArgs struct {
		user domain.User
		errorAssertion
	}
	var tests = []struct {
		name    string
		arrange func(t *testing.T)
		args    args
		want    assertionArgs
		assert  func(t *testing.T, user *domain.User, got error, want assertionArgs)
	}{
		{
			name: "not found",
			arrange: func(t *testing.T) {
				usecase, _, _ := newUsecase()
				u = usecase
			},
			args: args{
				ctx: ctx,
				user: domain.User{
					Id:    userId,
					Roles: []string{"role1", "role2"},
					Allow: []string{"allow1", "allow2"},
					Deny:  []string{"deny1", "deny2"},
				},
			},
			want: assertionArgs{
				errorAssertion: errorAssertion{
					hasErr:  true,
					errType: domain.ErrNotFound,
				},
			},
			assert: func(t *testing.T, user *domain.User, err error, want assertionArgs) {
				assertErr(t, err, want.errorAssertion)
			},
		},
		{
			name: "successful",
			arrange: func(t *testing.T) {
				usecase, repository, _ := newUsecase()
				u = usecase
				_, err = insertUserToDB(ctx, repository, domain.User{
					Id:    userId,
					Roles: make([]string, 0),
					Deny:  make([]string, 0),
					Allow: make([]string, 0),
				})
				if err != nil {
					t.Fatal(err)
				}

			},
			args: args{
				ctx: ctx,
				user: domain.User{
					Id:    userId,
					Roles: []string{"role1", "role2"},
					Allow: []string{"allow1", "allow2"},
					Deny:  []string{"deny1", "deny2"},
				},
			},
			want: assertionArgs{
				errorAssertion: errorAssertion{
					hasErr: false,
				},
				user: domain.User{
					Id:    userId,
					Roles: []string{"role1", "role2"},
					Allow: []string{"allow1", "allow2"},
					Deny:  []string{"deny1", "deny2"},
				},
			},
			assert: func(t *testing.T, user *domain.User, err error, want assertionArgs) {
				assertErr(t, err, want.errorAssertion)
				assertCreatedUser(t, user, want.user)
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
			dbUser, err := u.Update(tt.args.ctx, tt.args.user)
			// Assert
			if tt.assert != nil {
				tt.assert(t, dbUser, err, tt.want)
			}

		})
	}
}

func Test_usecase_GetUserById(t *testing.T) {
	u := &usecase{}
	var err error
	var userId uint64 = 654654654
	ctx := context.Background()

	type args struct {
		ctx context.Context
		id  uint64
	}
	type assertionArgs struct {
		user domain.User
		errorAssertion
	}
	var tests = []struct {
		name    string
		arrange func(t *testing.T)
		args    args
		want    assertionArgs
		assert  func(t *testing.T, user *domain.User, got error, want assertionArgs)
	}{
		{
			name: "not found",
			arrange: func(t *testing.T) {
				usecase, _, _ := newUsecase()
				u = usecase
			},
			args: args{
				ctx: ctx,
				id:  userId,
			},
			want: assertionArgs{
				errorAssertion: errorAssertion{
					hasErr: true,
				},
			},
			assert: func(t *testing.T, user *domain.User, err error, want assertionArgs) {
				assertErr(t, err, want.errorAssertion)
			},
		},
		{
			name: "successful",
			arrange: func(t *testing.T) {
				usecase, repository, _ := newUsecase()
				u = usecase
				_, err = insertUserToDB(ctx, repository, domain.User{
					Id:    userId,
					Roles: []string{"role1", "role2"},
					Allow: []string{"allow1", "allow2"},
					Deny:  []string{"deny1", "deny2"},
				})
				if err != nil {
					t.Fatal(err)
				}

			},
			args: args{
				ctx: ctx,
				id:  userId,
			},
			want: assertionArgs{
				errorAssertion: errorAssertion{
					hasErr: false,
				},
				user: domain.User{
					Id:    userId,
					Roles: []string{"role1", "role2"},
					Allow: []string{"allow1", "allow2"},
					Deny:  []string{"deny1", "deny2"},
				},
			},
			assert: func(t *testing.T, user *domain.User, err error, want assertionArgs) {
				assertErr(t, err, want.errorAssertion)
				assertCreatedUser(t, user, want.user)
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
			dbUser, err := u.GetUserById(tt.args.ctx, tt.args.id)
			log.Println("---------------------,", tt.name, dbUser, err)
			// Assert
			if tt.assert != nil {
				tt.assert(t, dbUser, err, tt.want)
			}

		})
	}
}
