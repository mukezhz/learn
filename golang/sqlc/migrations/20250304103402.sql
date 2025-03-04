-- Create "authors" table
CREATE TABLE "public"."authors" ("id" bigserial NOT NULL, "name" text NOT NULL, "bio" text NULL, "created_at" timestamptz NULL DEFAULT now(), "updated_at" timestamptz NULL DEFAULT now(), PRIMARY KEY ("id"));
-- Create "books" table
CREATE TABLE "public"."books" ("id" bigserial NOT NULL, "title" text NOT NULL, "author_id" bigint NOT NULL, "published_date" date NOT NULL, PRIMARY KEY ("id"), CONSTRAINT "books_author_id_fkey" FOREIGN KEY ("author_id") REFERENCES "public"."authors" ("id") ON UPDATE NO ACTION ON DELETE NO ACTION);
