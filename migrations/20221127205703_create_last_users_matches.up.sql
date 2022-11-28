create table last_users_matches (
    user_id bigint not null primary key,
    match_id bigint not null,

    constraint fk_user
        foreign key(user_id)
            references users(id),
    
    constraint fk_match
        foreign key(match_id)
            references matches(id)
);
