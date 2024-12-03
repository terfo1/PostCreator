ALTER TABLE "comments" DROP CONSTRAINT IF EXISTS "comments_post_id_fkey";
ALTER TABLE "comments" DROP CONSTRAINT IF EXISTS "comments_user_id_fkey";
ALTER TABLE "posts" DROP CONSTRAINT IF EXISTS "posts_user_id_fkey";
ALTER TABLE "admins" DROP CONSTRAINT IF EXISTS "admins_user_id_fkey";

DROP TABLE IF EXISTS "comments";
DROP TABLE IF EXISTS "posts";
DROP TABLE IF EXISTS "admins";
DROP TABLE IF EXISTS "users";
