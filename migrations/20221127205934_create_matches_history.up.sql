create table matches_history (
    id bigserial not null primary key,
    user_id bigint not null,
    match_id bigint not null,

    constraint fk_user
        foreign key(user_id)
            references users(id),

    constraint fk_match
        foreign key(match_id)
            references matches(id)
);
