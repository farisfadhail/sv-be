create table articles
(
    id         int auto_increment not null primary key,
    title   varchar(200) not null,
    content text not null,
    category varchar(100) not null,
    status  varchar(50) not null default 'draft',
    created_date timestamp default now(),
    updated_date timestamp default now()
) engine = InnoDB;