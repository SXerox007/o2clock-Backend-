create table all_users (
id uuid NOT NULL DEFAULT uuid_generate_v4() primary key,
phone text not null,
country_code text not null,
first_name text not null,
last_name text not null,
company_name text not null,
user_name text not null,
password text not null,
email text not null,
lat float,
lan float,
version text not null,
capture_time timestamp default current_timestamp
);