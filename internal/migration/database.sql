create table users
(
    id varchar(255) NOT NULL,
    full_name varchar(255) NOT NULL,
    password varchar(255) NOT NULL,
    email varchar(255) NOT NULL,
    created_at timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP
    updated_at timestamp NOT NULL DEFAULt CURRENT_TIMESTAMP on UPDATE CURRENT_TIMESTAMP
    primary key (id)
)engine=InnoDB;