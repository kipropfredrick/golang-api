CREATE TABLE "users" (
    "id" bigserial PRIMARY KEY,
    "username" varchar(100) NOT NULL,
    "email" varchar(100) NOT NULL,
    "password" text NOT NULL
)