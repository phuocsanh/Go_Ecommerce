package impl

import (
	"context"
	"go_ecommerce/internal/database"
)

type sUserLogin struct {
	r *database.Queries
}

func NewUserLoginImpl(r *database.Queries) *sUserLogin {
	return &sUserLogin{
		r : r,
	}
}

func (s *sUserLogin) Login(ctx context.Context) ( error) {
	return nil
}
func (s *sUserLogin) Register(ctx context.Context) ( error) {
	return nil
}
func (s *sUserLogin) VerifyOTP(ctx context.Context) ( error) {
	return nil
}
func (s *sUserLogin) UpdatePasswordRegister(ctx context.Context) ( error) {
	return nil
}