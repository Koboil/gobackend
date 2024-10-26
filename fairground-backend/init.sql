create table users
(
    user_id    serial
        primary key,
    first_name varchar(100) not null,
    last_name  varchar(100) not null,
    email      varchar(100) not null
        unique,
    password   varchar(100) not null,
    role       varchar(50)  not null
);

alter table users
    owner to admin;

create table parents
(
    parent_id      integer not null
        primary key
        references users,
    tokens_balance integer default 0,
    points_earned  integer default 0
);

alter table parents
    owner to admin;

create table students
(
    student_id     integer not null
        primary key
        references users,
    tokens_balance integer default 0,
    points_earned  integer default 0
);

alter table students
    owner to admin;

create table standholders
(
    stand_holder_id integer not null
        primary key
        references users
);

alter table standholders
    owner to admin;

create table organizers
(
    organizer_id   integer not null
        primary key
        references users,
    tokens_balance integer
);

alter table organizers
    owner to admin;

create table fairgroundmanagement
(
    management_id integer not null
        primary key
        references users
);

alter table fairgroundmanagement
    owner to admin;

create table tokens
(
    token_id      serial
        primary key,
    purchase_date timestamp default CURRENT_TIMESTAMP,
    amount        integer not null,
    purchased_by  integer
        references parents
            on delete cascade
);

alter table tokens
    owner to admin;

create index idx_token_parent
    on tokens (purchased_by);

create table transactions
(
    transaction_id serial
        primary key,
    from_user_id   integer
        references users
            on delete cascade,
    to_user_id     integer
        references users
            on delete cascade,
    tokens         integer not null,
    transfer_date  timestamp default CURRENT_TIMESTAMP
);

alter table transactions
    owner to admin;

create index idx_transaction_user
    on transactions (from_user_id, to_user_id);

create table activity
(
    activity_id    serial
        primary key,
    user_id        integer
        references users
            on delete cascade,
    activity_type  varchar(50) not null,
    tokens_spent   integer   default 0,
    points_awarded integer   default 0,
    activity_date  timestamp default CURRENT_TIMESTAMP
);

alter table activity
    owner to admin;

create index idx_activity_user
    on activity (user_id);

create table parent_student
(
    parent_student_id serial
        primary key,
    parent_id         integer
        references parents,
    student_id        integer
        references students
);

alter table parent_student
    owner to admin;

create table raffles
(
    raffle_id     serial
        primary key,
    organizer_id  integer      not null
        references organizers
            on delete cascade,
    raffle_name   varchar(100) not null,
    ticket_price  integer      not null,
    prize_details varchar(255),
    start_date    timestamp default CURRENT_TIMESTAMP,
    end_date      timestamp
);

alter table raffles
    owner to admin;

create table raffletickets
(
    raffle_ticket_id serial
        primary key,
    purchase_date    timestamp default CURRENT_TIMESTAMP,
    user_id          integer
        references users
            on delete cascade,
    raffle_id        integer
        constraint raffle_id_fk
            references raffles
);

alter table raffletickets
    owner to admin;

create index idx_user_raffle
    on raffletickets (user_id);


INSERT INTO public.users (user_id, first_name, last_name, email, password, role) VALUES (1000, 'Harry', 'Mike', 'admin@gmail.com', '$2a$14$Wa7FiIGt3X6IayD1La9az.6Fbu0xFDzijlumFzTj7RovIXD7/6yGm', 'fairground-management');
INSERT INTO public.users (user_id, first_name, last_name, email, password, role) VALUES (1001, 'James', 'Bond', 'organizer@gmail.com', '$2a$14$ZvDEmWS4b6p8sUa6F1mM/uHf0YaXrnoKmNBQr9y3Qs1Ufqz9iFHkO', 'organizer');

INSERT INTO public.fairgroundmanagement (management_id) VALUES (1000);
INSERT INTO public.organizers (organizer_id,tokens_balance) VALUES (1001,0);


INSERT INTO public.raffles (raffle_id, organizer_id, raffle_name, ticket_price, prize_details, start_date, end_date) VALUES (1, 1001, 'Charity', 50, 'Bunny', '2024-10-25 06:54:15.181690', '2024-11-30 11:52:57.000000');
INSERT INTO public.raffles (raffle_id, organizer_id, raffle_name, ticket_price, prize_details, start_date, end_date) VALUES (2, 1001, 'Lucky Draw', 100, '1000 Tokens', '2024-10-25 06:54:15.181690', '2024-12-25 11:54:06.000000');
