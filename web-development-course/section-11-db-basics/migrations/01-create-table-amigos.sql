CREATE TABLE amigos
(
    id serial NOT NULL,
    name character varying(255) NOT NULL,
    PRIMARY KEY (id)
)
WITH (
    OIDS = FALSE
);

ALTER TABLE amigos
    OWNER to golang;
