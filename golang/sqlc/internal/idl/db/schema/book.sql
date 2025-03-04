CREATE TABLE books (
    id BIGSERIAL PRIMARY KEY,
    title text NOT NULL,
    author_id BIGINT NOT NULL REFERENCES authors (id),
    published_date DATE NOT NULL
);