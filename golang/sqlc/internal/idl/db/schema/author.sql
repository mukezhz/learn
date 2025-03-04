CREATE TABLE authors (
  id BIGSERIAL PRIMARY KEY,
  name text NOT NULL,
  bio text,
  age VARCHAR(10),
  created_at TIMESTAMPTZ DEFAULT NOW(),
  updated_at TIMESTAMPTZ DEFAULT NOW()
);
