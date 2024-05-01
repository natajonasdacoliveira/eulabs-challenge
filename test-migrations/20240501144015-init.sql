-- +migrate Up
CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE EXTENSION IF NOT EXISTS "pgcrypto";

CREATE TABLE product(
  id UUID PRIMARY KEY DEFAULT uuid_generate_v4(),
  name VARCHAR(255) NOT NULL,
  price DECIMAL(10, 2) NOT NULL,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

-- +migrate Down
DROP EXTENSION IF EXISTS "pgcrypto";

DROP EXTENSION IF EXISTS "uuid-ossp";