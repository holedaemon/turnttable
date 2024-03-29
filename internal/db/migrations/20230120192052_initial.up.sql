BEGIN;

CREATE TYPE medium AS ENUM (
    'cd',
    'vinyl',
    'cassette'
);

CREATE TABLE records (
    id bigint GENERATED BY DEFAULT AS IDENTITY PRIMARY KEY,
    created_at timestamp DEFAULT NOW() NOT NULL,
    updated_at timestamp DEFAULT NOW() NOT NULL,
    title text NOT NULL,
    artist text NOT NULL,
    label text NOT NULL,
    cn text NOT NULL, -- catalog number
    genre text NOT NULL,
    released timestamptz NOT NULL,
    purchased timestamptz,
    medium medium NOT NULL
);

COMMIT;