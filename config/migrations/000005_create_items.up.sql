BEGIN;
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
COMMIT;

