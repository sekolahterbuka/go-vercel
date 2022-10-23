CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE
    "auth" (
        "id" uuid PRIMARY KEY NOT NULL,
        "username" varchar UNIQUE NOT NULL,
        "password" varchar NOT NULL,
        "email" VARCHAR NOT NULL
    );