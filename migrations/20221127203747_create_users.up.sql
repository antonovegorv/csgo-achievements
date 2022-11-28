create table users (
    id bigserial not null primary key,
    steam_id bigint not null,
    game_authentication_code varchar not null
);
