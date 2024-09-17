
ALTER TABLE "token" DROP CONSTRAINT "token_user_id_fkey";

DROP TABLE IF EXISTS "token";

DROP TABLE IF EXISTS "account";

DROP TYPE IF EXISTS "token_type";
DROP TYPE IF EXISTS "Role";