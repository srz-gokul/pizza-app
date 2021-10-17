package data

import (
	"errors"
	"testing"

	"gopkg.in/DATA-DOG/go-sqlmock.v1"
)

func Test_repo_AddOrderDetails(t *testing.T) {
	type args struct {
		order *OrderData
	}
	tests := []struct {
		name    string
		given   func(db sqlmock.Sqlmock)
		args    args
		wantErr bool
	}{
		{
			name: "Normal case",
			given: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(".+").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit()
			},
			args: args{
				&OrderData{
					UserID:    1,
					PizzaID:   1,
					PizzaSize: "medium",
				},
			},
		},
		{
			name: "Tx begin failed",
			given: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin().WillReturnError(errors.New("error"))
			},
			wantErr: true,
		},
		{
			name: "query execution failed",
			given: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(".+").WillReturnError(errors.New("error"))
			},
			args: args{
				&OrderData{
					UserID:    1,
					PizzaID:   1,
					PizzaSize: "medium",
				},
			},
			wantErr: true,
		},
		{
			name: "Tx commit failed",
			given: func(mock sqlmock.Sqlmock) {
				mock.ExpectBegin()
				mock.ExpectExec(".+").WillReturnResult(sqlmock.NewResult(1, 1))
				mock.ExpectCommit().WillReturnError(errors.New("error"))
			},
			args: args{
				&OrderData{
					UserID:    1,
					PizzaID:   1,
					PizzaSize: "medium",
				},
			},
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			db, mock, err := sqlmock.New()
			if err != nil {
				t.Errorf("unable initialize mock database: %v", err)
				return
			}
			r := New(db)
			tt.given(mock)
			if err := r.AddOrderDetails(tt.args.order); (err != nil) != tt.wantErr {
				t.Errorf("repo.AddOrderDetails() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
