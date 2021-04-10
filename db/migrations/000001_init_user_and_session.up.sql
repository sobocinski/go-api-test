CREATE TABLE "users" (
    id serial PRIMARY KEY,
    "password" varchar(64) NOT NULL,
    "email" varchar(300) NOT NULL,
    "created_at" timestamptz NOT NULL DEFAULT now(),
    "updated_at" timestamptz
);

CREATE TABLE IF NOT EXISTS "cars"(
   id serial PRIMARY KEY,
   "user_id" int4 NOT NULL,
   "name" varchar(300) NOT NULL,
   "manufacturer" varchar(60) NOT NULL,
   "model" varchar(60) NOT NULL,
   "registration_number" varchar(16),
   "created_at" timestamptz NOT NULL DEFAULT now(),
   "updated_at" timestamptz,
   CONSTRAINT fb_user
      FOREIGN KEY(user_id) 
	   REFERENCES users(id)
);

