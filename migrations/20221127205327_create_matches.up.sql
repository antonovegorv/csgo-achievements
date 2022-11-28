create table matches (
    id bigserial not null primary key,
    share_code varchar not null unique,
    match_id bigint not null,
    outcome_id bigint not null,
    token_id bigint not null
);
