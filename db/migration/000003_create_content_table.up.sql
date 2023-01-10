CREATE TABLE IF NOT EXISTS contents (
  id BIGSERIAL PRIMARY KEY,
  title TEXT NOT NULL,
  description TEXT NULL,
  file_path TEXT NULL,
  play_time TIMESTAMPTZ NOT NULL,
  resolution INTEGER NULL,
  aspect_ratio TEXT NULL,
  tag TEXT NULL,
  category TEXT NULL,
	created_at TIMESTAMPTZ NULL,
	updated_at TIMESTAMPTZ NULL,
	deleted_at TIMESTAMPTZ NULL
);
