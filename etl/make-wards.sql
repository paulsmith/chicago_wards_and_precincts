begin;

drop table if exists wards;

create table wards (
    ward varchar not null,
    primary key (ward)
);

select AddGeometryColumn('wards', 'wkb_geometry', 4326, 'GEOMETRY', 2);

insert into wards (ward, wkb_geometry) (
    select p.ward, st_union(st_buffer(wkb_geometry, 0.000001))
    from precincts p
    group by p.ward
);

commit;

