-- Table: public.groups

-- DROP TABLE IF EXISTS public.groups;

CREATE TABLE IF NOT EXISTS public.groups
(
    group_id integer NOT NULL DEFAULT nextval('groups_group_id_seq'::regclass),
    group_name character varying(10) COLLATE pg_catalog."default" NOT NULL,
    direction character varying COLLATE pg_catalog."default",
    CONSTRAINT groups_pkey PRIMARY KEY (group_id),
    CONSTRAINT groups_direction_check CHECK (direction::text = ANY (ARRAY['Engineers'::character varying, 'Humanities'::character varying]::text[]))
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.groups
    OWNER to postgres;