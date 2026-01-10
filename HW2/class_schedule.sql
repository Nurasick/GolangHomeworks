-- Table: public.class_schedule

-- DROP TABLE IF EXISTS public.class_schedule;

CREATE TABLE IF NOT EXISTS public.class_schedule
(
    schedule_id integer NOT NULL DEFAULT nextval('class_schedule_schedule_id_seq'::regclass),
    subject character varying(50) COLLATE pg_catalog."default" NOT NULL,
    lesson_time character varying(30) COLLATE pg_catalog."default",
    group_id integer,
    CONSTRAINT class_schedule_pkey PRIMARY KEY (schedule_id),
    CONSTRAINT class_schedule_group_id_fkey FOREIGN KEY (group_id)
        REFERENCES public.groups (group_id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.class_schedule
    OWNER to postgres;