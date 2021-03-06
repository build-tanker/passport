CREATE TABLE person (
  id UUID NOT NULL PRIMARY KEY,
  source VARCHAR(128),
  name VARCHAR(128),
  email VARCHAR(128),
  picture_url VARCHAR(512),
  gender VARCHAR(32),
  source_id VARCHAR(256),
  deleted BOOLEAN DEFAULT false,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);

CREATE TABLE token (
  id UUID NOT NULL PRIMARY KEY,
  person UUID NOT NULL REFERENCES person(id),
  source VARCHAR(128),
  access_token UUID,
  expires_in INT DEFAULT 2592000,
  external_access_token VARCHAR(256),
  external_refresh_token VARCHAR(256),
  external_expires_in INT,
  external_token_type VARCHAR(128),
  deleted BOOLEAN DEFAULT false,
  created_at TIMESTAMP DEFAULT NOW(),
  updated_at TIMESTAMP DEFAULT NOW()
);