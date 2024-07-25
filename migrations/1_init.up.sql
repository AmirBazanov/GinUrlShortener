CREATE TABLE IF NOT EXISTS url(
                                  id SERIAL,
                                  alias TEXT NOT NULL UNIQUE,
                                d_token TEXT NOT NULL,
                                  url TEXT NOT NULL);
CREATE INDEX IF NOT EXISTS idx_alias ON url(alias);