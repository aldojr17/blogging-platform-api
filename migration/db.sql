CREATE database blogging_db;

CREATE TABLE IF NOT EXISTS post_tab (
  id SERIAL PRIMARY KEY,
  title VARCHAR(128) NOT NULL,
  content TEXT NOT NULL,
  category VARCHAR(32) NOT NULL,
  create_time BIGINT DEFAULT 0,
  update_time BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS tag_tab (
  id SERIAL PRIMARY KEY,
  name VARCHAR(32) NOT NULL,
  create_time BIGINT DEFAULT 0
);

CREATE TABLE IF NOT EXISTS post_tag_tab (
  id SERIAL PRIMARY KEY,
  post_id INT NOT NULL,
  tag_id INT NOT NULL,
  create_time BIGINT DEFAULT 0
);