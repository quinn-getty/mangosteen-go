CREATE TABLE users(
  id serial PRIMARY KEY,
  email varchar(255) NOT NULL UNIQUE,
  phone varchar(20) NOT NULL,
  address varchar(255) NOT NULL DEFAULT '',
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);

