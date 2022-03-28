CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE order_lists
(
    id          serial       not null unique,
    title       varchar(255) not null,
    description varchar(255),
    time varchar(255),
    done boolean not null default false
);

CREATE TABLE users_lists
(
    id      serial                                           not null unique,
    user_id int references users (id) on delete cascade      not null,
    list_id int references order_lists (id) on delete cascade not null
);