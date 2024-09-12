-- +goose Up
-- +goose StatementBegin
CREATE TYPE "token_type" AS ENUM (
  'refresh_token',
  'access_token',
  'magic_link'
);

CREATE TYPE "Role" AS ENUM (
  'user',
  'subscriber',
  'moderator',
  'creator',
  'admin',
  'organizer',
  'recruiter'
);

CREATE TABLE "account" (
  "id" uuid PRIMARY KEY,
  "username" varchar(50) UNIQUE NOT NULL,
  "email" VARCHAR(100) UNIQUE NOT NULL,
  "password_hash" varchar,
  "bio" text,
  "profile_image" VARCHAR(255),
  "position" varchar(100),
  "organisation" varchar(100) NOT NULL,
  "birthday" date
);

CREATE TABLE "token" (
  "id" uuid PRIMARY KEY,
  "token" varchar NOT NULL,
  "user_id" int NOT NULL,
  "expiry" timestamptz NOT NULL DEFAULT 'now()'
);

ALTER TABLE "token" ADD FOREIGN KEY ("user_id") REFERENCES "account" ("id");
-- +goose StatementEnd


-- +goose Up
-- +goose StatementBegin
ALTER TABLE "token" DROP CONSTRAINT "token_user_id_fkey";

DROP TABLE IF EXISTS "token";

DROP TABLE IF EXISTS "account";

DROP TYPE IF EXISTS "token_type";
DROP TYPE IF EXISTS "Role";
-- +goose StatementEnd