
create table all_tokens (
id uuid NOT NULL DEFAULT uuid_generate_v4() primary key,
user_id uuid REFERENCES all_users(id),
access_token text not null,
version text not null,
capture_time timestamp default current_timestamp
);