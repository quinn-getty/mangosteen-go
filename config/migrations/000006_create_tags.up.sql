CREATE TABLE IF NOT EXISTS tags(
  id serial PRIMARY KEY,
  user_id serial NOT NULL,
  name varchar(50) NOT NULL,
  sign varchar(10) NOT NULL,
  deleted_at timestamp,
  created_at timestamp NOT NULL DEFAULT now(),
  updated_at timestamp NOT NULL DEFAULT now()
);

