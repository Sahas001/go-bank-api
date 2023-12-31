package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

type Storage interface {
	CreateAccount(*Account) error
	DeleteAccount(int) error
	UpdateAccount(*Account) error
  GetAccounts() ([]*Account, error)
	GetAccountById(int) (*Account, error)
}

type PostgresStore struct {
	db *sql.DB
}

func NewPostgresStore() (*PostgresStore, error) {
	connStr := "user=postgres dbname=postgres password=gobank sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}
	if err := db.Ping(); err != nil {
		return nil, err
	}
	return &PostgresStore{
		db: db,
	}, nil
}

func (s *PostgresStore) Init() error {
	return s.createAccountTable()
}

func (s *PostgresStore) createAccountTable() error {
	query := `create table if not exists account(
    id serial primary key,
    firstname varchar(50),
    lastname varchar(50),
    number serial,
    balance serial,
    created_at timestamp
  )`

	_, err := s.db.Exec(query)
	return err

}

func (s *PostgresStore) CreateAccount(acc *Account) error {
	query := (`
  insert into account 
  (firstname, lastname, number, balance, created_at) 
  values ('John', 'Doe', 1234567890, 1000, '2017-01-01')`)
  resp, err := s.db.Query(query, acc.Firstname, acc.Lastname, acc.Number, acc.Balance, acc.CreatedAt)

  if err != nil {
    return err
  }
  fmt.Println(resp)
	return nil
}

func (s *PostgresStore) UpdateAccount(*Account) error {
	return nil
}

func (s *PostgresStore) DeleteAccount(id int) error {
	return nil
}

func (s *PostgresStore) GetAccountById(id int) (*Account, error) {
	return nil, nil
}

func (s *PostgresStore) GetAccounts() ([]*Account, error) {
  rows, err := s.db.Query(`select * from account`)
  if err != nil {
    return nil, err
  }

  accounts := []*Account{}
  for rows.Next() {
  account, err := scanIntoAccounts(rows)
  if err != nil {
    return nil, err
  }
  accounts = append(accounts, account)
  }
  return accounts, nil
}

func scanIntoAccounts(rows *sql.Rows) (*Account, error) {
    account := new(Account)
    err := rows.Scan(&account.ID, &account.Firstname, &account.Lastname, &account.Number, &account.Balance, &account.CreatedAt)  
  return account, err
}











