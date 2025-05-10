-- +migrate Up
CREATE TABLE "testing_create_tb" (
  "test_id" text PRIMARY KEY,
  "full_name" text,
  "email" text UNIQUE,
  "password" text,
  "role" text,
  "created_at" TIMESTAMPTZ NOT NULL,
  "updated_at" TIMESTAMPTZ NOT NULL
);

-- +migrate Down
DROP TABLE testing_create_tb;