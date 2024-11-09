create table player_slots
(
    id            uuid primary key,
    serial_number serial unique
);

create table players
(
    id             uuid primary key
        references player_slots (id)
            on delete cascade,
    player_name    text not null,
    contact_number text
);

create table player_events
(
    id         serial primary key,
    user_id    uuid
        references player_slots
            on delete cascade,
    event_type varchar(10) not null, -- win , lose
    amount     int         not null
);

create or replace view player_stats as
select user_id,
       sum(case when event_type = 'win' then amount else -amount end) as current_chips,
       sum(case when event_type = 'win' then 1 else 0 end)            as win,
       sum(case when event_type = 'lose' then 1 else 0 end)           as lose
       -- implement player tiers
from player_events
group by user_id;