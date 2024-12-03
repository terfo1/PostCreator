CREATE TABLE "admins" (
                          "id" integer PRIMARY KEY NOT NULL,
                          "user_id" integer UNIQUE NOT NULL,
                          "created_at" timestamp
);

CREATE TABLE "users" (
                         "id" integer PRIMARY KEY NOT NULL,
                         "email" varchar UNIQUE NOT NULL,
                         "username" varchar UNIQUE NOT NULL,
                         "password" varchar NOT NULL,
                         "created_at" timestamp
);

CREATE TABLE "posts" (
                         "id" integer PRIMARY KEY NOT NULL,
                         "title" varchar NOT NULL,
                         "body" text,
                         "user_id" integer NOT NULL,
                         "created_at" timestamp
);

CREATE TABLE "comments" (
                            "id" integer PRIMARY KEY NOT NULL,
                            "post_id" integer NOT NULL,
                            "user_id" integer NOT NULL,
                            "content" text,
                            "created_at" timestamp
);

COMMENT ON COLUMN "posts"."body" IS 'Content of the post';

ALTER TABLE "posts" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id");

ALTER TABLE "comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");

ALTER TABLE "admins" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id");
