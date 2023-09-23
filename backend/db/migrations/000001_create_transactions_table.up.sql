CREATE DOMAIN pk AS 
    VARCHAR(32) NOT NULL CHECK (value !~ '\s');

CREATE TABLE IF NOT EXISTS users (
  id pk PRIMARY KEY,
  name VARCHAR(45) NOT NULL,
  email VARCHAR(60) NOT NULL,
  password VARCHAR(45) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()  
);

CREATE TABLE IF NOT EXISTS accounts (
  id pk PRIMARY KEY,
  name VARCHAR(45) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()  
);

CREATE TYPE transaction_type AS ENUM ('Expense', 'Income');
CREATE TYPE transaction_status AS ENUM('Pending', 'Outdated', 'Paid');

CREATE TABLE IF NOT EXISTS transactions (
  id pk PRIMARY KEY,
  user_id pk NOT NULL REFERENCES users,
  type transaction_type NOT NULL,  
  account_id pk NOT NULL REFERENCES accounts,
  transaction_date TIMESTAMPTZ,
  status transaction_status NOT NULL DEFAULT 'Pending',
  created_at TIMESTAMPTZ NOT NULL DEFAULT now(),
  updated_at TIMESTAMPTZ NOT NULL DEFAULT now()
);