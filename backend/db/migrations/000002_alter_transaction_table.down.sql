ALTER TABLE transactions
DROP COLUMN IF EXISTS description,
DROP COLUMN IF EXISTS category_id,
DROP COLUMN IF EXISTS date,
DROP COLUMN IF EXISTS amount; 

DROP TABLE categories;