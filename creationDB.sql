DROP TABLE album

CREATE TABLE album(
    id serial not null,
    titolo text,
    artista_id int,
    prezzo numeric(5,2),
    constraint album_pk primary key (id),
    constraint album_artista_fk foreign key (artista_id) references artista(id) on delete cascade
);

CREATE TABLE artista(
    id serial not null,
    nome text,
    cognome text,
    data_nascita int,
    casa_discografica_id int,
    constraint artista_pk primary key (id),
    constraint artista_casa_fk foreign key (casa_discografica_id) references casaDiscografica(id)on delete cascade
);

CREATE TABLE casaDiscografica(
    id serial not null,
    nome text,
    data_fondazione int,
    constraint casa_pk primary key (id)
);