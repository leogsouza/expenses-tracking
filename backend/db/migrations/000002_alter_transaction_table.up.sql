CREATE TABLE IF NOT EXISTS categories (
  id pk PRIMARY KEY,
  name VARCHAR(45) NOT NULL,
  created_at TIMESTAMPTZ NOT NULL DEFAULT now()
);

INSERT INTO categories (id, name, created_at) VALUES ('123456', 'Outros', now());

ALTER TABLE transactions
ADD COLUMN description VARCHAR(60) NOT NULL,
ADD COLUMN category_id pk NOT NULL REFERENCES categories,
ADD COLUMN date TIMESTAMPTZ NOT NULL,
ADD COLUMN amount MONEY NOT NULL DEFAULT 0.00;