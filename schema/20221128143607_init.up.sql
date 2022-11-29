CREATE TABLE IF NOT EXISTS accounts (
    id serial not null unique,
    user_id int unique,
    balance int default 0
);

CREATE TABLE IF NOT EXISTS transactions (
    id serial not null unique,
    pub_id int not null,
    sub_id int not null,
    value int not null,
    status varchar(255) not null,
    created_at timestamp default current_timestamp
);

INSERT INTO accounts (user_id, balance) VALUES (1, 100), (2, 200), (3, 0), (4, 5000)