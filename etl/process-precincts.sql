begin;

alter table precincts rename to precincts_orig;

create table precincts (
    id varchar not null,
    ward varchar not null,
    precinct varchar not null,
    primary key (id)
);

select AddGeometryColumn('precincts', 'wkb_geometry', 4326, 'GEOMETRY', 2);

insert into precincts (id, ward, precinct, wkb_geometry) (
    select id, substring(id for 2)::int::varchar as ward, substring(id from 3 for 3)::int::varchar as precinct, wkb_geometry
    from precincts_orig
);

drop table precincts_orig;

commit;
