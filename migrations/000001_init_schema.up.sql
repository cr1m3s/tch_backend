CREATE TABLE "users" (
    "id" BIGSERIAL NOT NULL,
    "name" VARCHAR NOT NULL,
    "email" VARCHAR NOT NULL,
    "photo" VARCHAR NOT NULL,
    "verified" BOOLEAN NOT NULL,
    "password" VARCHAR NOT NULL,
    "role" VARCHAR NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    "updated_at" TIMESTAMP NOT NULL,

    CONSTRAINT "users_pkey" PRIMARY KEY ("id")
);

CREATE UNIQUE INDEX "users_email_key" ON "users"("email");

CREATE TABLE "advertisements" (
    "id" BIGSERIAL NOT NULL,
    "title" VARCHAR NOT NULL,
    "provider" VARCHAR NOT NULL,
    "attachment" VARCHAR NOT NULL,
    "experience"  VARCHAR NOT NULL, 
    "category" VARCHAR NOT NULL,
    "time" VARCHAR NOT NULL,
    "price" SERIAL NOT NULL,
    "format" VARCHAR NOT NULL,
    "language" VARCHAR NOT NULL,
    "description" VARCHAR NOT NULL,
    "mobile_phone" VARCHAR NOT NULL,
    "email" VARCHAR NOT NULL,
    "telegram" VARCHAR NOT NULL,
    "created_at" TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,

    CONSTRAINT "advertisement_pkey" PRIMARY KEY ("id")
);