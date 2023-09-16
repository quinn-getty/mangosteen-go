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

CREATE TYPE kind AS ENUM(
  'expenses',
  'in_come'
);

CREATE TABLE items(
  id serial PRIMARY KEY,
  user_id serial NOT NULL,
  amount integer NOT NULL,
  tag_ids integer[] NOT NULL,
  kind kind NOT NULL,
  happened_at timestamp NOT NULL DEFAULT now(),
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);

CREATE TABLE IF NOT EXISTS tags(
  id serial PRIMARY KEY,
  user_id serial NOT NULL,
  name varchar(50) NOT NULL,
  sign varchar(10) NOT NULL,
  deleted_at timestamp,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);

