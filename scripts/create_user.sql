CREATE SEQUENCE IF NOT EXISTS users_id_seq;

-- Table Definition
CREATE TABLE "users" (
    "id" int4 NOT NULL DEFAULT nextval('users_id_seq'::regclass),
    "username" varchar(100),
    "password" varchar(100),
    PRIMARY KEY ("id")
);

INSERT INTO "users" ("id", "username", "password") VALUES
(1, 'john', 'cGFzc3dvcmQ=');

