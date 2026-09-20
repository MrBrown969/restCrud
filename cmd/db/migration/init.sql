create table if not exists mobs (
    id varchar,
    name varchar,
    hp int,
    lvl varchar
);

insert into mobs (id,name,hp,lvl)
values
    ('1','Valakas',1000000,'raid boss'),
    ('2','Gremlin',1,'1');