create table microposts (
  id serial primary key,
  uuid varchar(64) not null unique,
  title varchar(255),
  content text,
  created_at timestamp not null
)