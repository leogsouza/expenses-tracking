CREATE TYPE PK VARCHAR(32)

CREATE TABLE IF NOT EXISTS users (
  id PK PRIMARY KEY,
  name VARCHAR(45) NOT NULL,
  email VARCHAR(60) NOT NULL,
  password VARCHAR(45) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()  
)

INSERT INTO users (id, name, email, password, created_at) VALUES ("123456", "Leonardo", "teste@teste.com", "12345678", now());

CREATE TABLE IF NOT EXISTS accounts (
  id PK PRIMARY KEY,
  name VARCHAR(45) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()  
)

INSERT INTO accounts (id, name, created_at) VALUES ("123456", "Carteira", now());

CREATE TYPE transaction_type AS ENUM ('Expense', 'Income')
CREATE TYPE transaction_status AS ENUM('Pending', 'Outdated', 'Paid')

CREATE TABLE IF NOT EXISTS transactions (
  id PK PRIMARY KEY,
  user_id PK NOT NULL REFERENCES users,
  type transaction_type NOT NULL,  
  account_id PK NOT NULL REFERENCES accounts
  transaction_date TIMESTAMPTZ
  status transaction_status NOT NULL DEFAULT 'Pending'
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
)