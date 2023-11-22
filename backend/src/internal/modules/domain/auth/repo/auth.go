package repo

import (
	"context"
	"golang.org/x/crypto/bcrypt"
	"hospital/internal/modules/db"
	"hospital/internal/modules/db/ent"
	"hospital/internal/modules/db/ent/account"
	"hospital/internal/modules/domain/auth/dto"
)

type AccountRepo struct {
	client *ent.Client
}

func NewAccountRepo(client *ent.Client) *AccountRepo {
	return &AccountRepo{
		client: client,
	}
}

func (r *AccountRepo) GetById(ctx context.Context, id int) (*dto.Account, error) {
	Account, err := r.client.Account.Get(ctx, id)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToAccountDTO(Account), nil
}

func (r *AccountRepo) List(ctx context.Context) (dto.Accounts, error) {
	Accounts, err := r.client.Account.Query().All(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToAccountDTOs(Accounts), nil
}

func (r *AccountRepo) Create(ctx context.Context, dtm *dto.CreateAccount) (*dto.Account, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(dtm.Password), bcrypt.DefaultCost)
	if err != nil {
		return nil, db.WrapError(err)
	}
	Account, err := r.client.Account.Create().
		SetLogin(dtm.Login).
		SetPasswordHash(string(hashedPassword)).
		Save(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToAccountDTO(Account), nil
}

func (r *AccountRepo) Login(ctx context.Context, dtm *dto.CreateAccount) (*dto.Account, error) {

	Account, err := r.client.Account.Query().
		Where(account.LoginEQ(dtm.Login)).
		Only(ctx)
	if bcrypt.CompareHashAndPassword([]byte(Account.PasswordHash), []byte(dtm.Password)) != nil {
		return nil, db.WrapError(err)
	}
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToAccountDTO(Account), nil
}

func (r *AccountRepo) Update(ctx context.Context, id int, dtm *dto.UpdateAccount) (*dto.Account, error) {
	Account, err := r.client.Account.UpdateOneID(id).
		SetLogin(dtm.Login).
		SetPasswordHash(dtm.Password).
		Save(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToAccountDTO(Account), nil
}

func (r *AccountRepo) Delete(ctx context.Context, id int) error {
	err := r.client.Account.DeleteOneID(id).Exec(ctx)
	if err != nil {
		return db.WrapError(err)
	}

	return nil
}

func (r *AccountRepo) Restore(ctx context.Context, id int) (*dto.Account, error) {
	Account, err := r.client.Account.UpdateOneID(id).Save(ctx)
	if err != nil {
		return nil, db.WrapError(err)
	}

	return ToAccountDTO(Account), nil
}

func ToAccountDTO(model *ent.Account) *dto.Account {
	if model == nil {
		return nil
	}
	return &dto.Account{
		Id:       model.ID,
		Password: model.PasswordHash,
		Login:    model.Login,
	}
}

func ToAccountDTOs(models ent.Accounts) dto.Accounts {
	if models == nil {
		return nil
	}
	dtms := make(dto.Accounts, len(models))
	for i := range models {
		dtms[i] = ToAccountDTO(models[i])
	}
	return dtms
}
