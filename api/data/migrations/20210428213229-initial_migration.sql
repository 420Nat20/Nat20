-- +migrate Up
create table campaigns (
    id serial not null constraint campaigns_pk primary key,
    created_at timestamp,
    updated_at timestamp,
    deleted_at timestamp,
    dm varchar(50) not null
);
alter table campaigns owner to doadmin;
create table locations (
    id serial not null constraint location_pk primary key,
    updated_at timestamp not null,
    created_at timestamp not null,
    deleted_at timestamp,
    name varchar(100) not null,
    event_description varchar(255),
    visited boolean default false,
    campaign_id integer not null constraint location_campaigns_id_fk references campaigns on update cascade on delete cascade
);
alter table locations owner to doadmin;
create table sub_locations (
    id serial not null constraint sub_locations_pk primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    deleted_at timestamp,
    name varchar(50) not null,
    event_description varchar(255),
    visited boolean default false,
    location_id integer not null constraint sub_locations_location_id_fk references locations on update cascade on delete cascade
);
alter table sub_locations owner to doadmin;
create table users (
    id serial not null constraint users_pk primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    deleted_at timestamp,
    discord_id varchar(100) not null
);
alter table users owner to doadmin;
create unique index users_discord_id_uindex on users (discord_id);
create table player_characters (
    id serial not null constraint player_characters_pk primary key,
    created_at timestamp not null,
    updated_at timestamp not null,
    deleted_at timestamp,
    campaign_id integer constraint player_characters_campaigns_id_fk references campaigns on update cascade on delete cascade,
    user_id integer not null constraint player_characters_users_id_fk references users on update cascade on delete cascade,
    name varchar(50) not null,
    class varchar(50) not null,
    background varchar(255) not null,
    race varchar(30) not null,
    alignment varchar(50) not null,
    strength integer not null constraint strength_valid check (
        (strength > 8)
        AND (strength < 15)
    ),
    dexterity integer not null constraint dexterity_valid check (
        (dexterity > 8)
        AND (dexterity < 15)
    ),
    constitution integer not null constraint constitution_valid check (
        (constitution > 8)
        AND (constitution < 15)
    ),
    intelligence integer not null constraint intelligence_valid check (
        (intelligence > 8)
        AND (intelligence < 15)
    ),
    wisdom integer not null constraint wisdom_valid check (
        (wisdom > 8)
        AND (wisdom < 15)
    ),
    charisma integer not null constraint charisma_valid check (
        (charisma > 8)
        AND (charisma < 15)
    ),
    trait_one varchar(255),
    trait_two varchar(255),
    ideal varchar(255),
    bond varchar(255),
    flaw varchar(255)
);
alter table player_characters owner to doadmin;
-- +migrate Down
DROP SCHEMA public CASCADE;
CREATE SCHEMA public;