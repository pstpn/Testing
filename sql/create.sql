create table if not exists company
(
    id serial primary key,
    name text,
    city text
);

create table if not exists employee
(
    id serial primary key,
    phone_number text unique,
    full_name text,
    company_id int references company(id) on delete cascade,
    post text,
    password text,
    refresh_token text,
    token_expired_at timestamp,
    date_of_birth date
);

create table if not exists info_card
(
    id serial primary key,
    created_employee_id int references employee(id) on delete cascade,
    is_confirmed boolean,
    created_date date default now()
);

create table if not exists document
(
    id serial primary key,
    serial_number text,
    info_card_id int references info_card(id) on delete cascade,
    type text
);

create table if not exists photo
(
    id serial primary key,
    document_id int references document(id) on delete cascade,
    key text
);

create table if not exists field
(
    id serial primary key,
    document_id int references document(id) on delete cascade,
    type text,
    value text,

    unique (document_id, type)
);

create table if not exists checkpoint
(
    id serial primary key,
    phone_number text
);

create table if not exists passage
(
    id serial primary key,
    checkpoint_id int references checkpoint(id) on delete cascade,
    document_id int references document(id) on delete cascade,
    type text,
    time timestamp
);