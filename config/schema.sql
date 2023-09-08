CREATE TABLE users(
  id serial PRIMARY KEY,
  email varchar(255) NOT NULL UNIQUE,
  phone varchar(20) NOT NULL,
  address varchar(255) NOT NULL DEFAULT '',
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS validation_codes(
  id serial PRIMARY KEY,
  code varchar(20) NOT NULL,
  email varchar(255) NOT NULL,
  useded_at timestamp,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);

