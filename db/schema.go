package db

// Set up the schema for the database

var schema = `

CREATE TABLE IF NOT EXISTS public.users
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    username character varying(50) COLLATE pg_catalog."default" NOT NULL,
    password character varying(50) COLLATE pg_catalog."default" NOT NULL,
    user_type character varying(50) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT user_id PRIMARY KEY (id)
)

CREATE TABLE IF NOT EXISTS public.apps
(
    id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
    durationmins numeric,
    doc_id integer,
    pat_id integer,
    CONSTRAINT apps_pkey PRIMARY KEY (id),
    CONSTRAINT doc_id FOREIGN KEY (doc_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION,
    CONSTRAINT pat_id FOREIGN KEY (pat_id)
        REFERENCES public.users (id) MATCH SIMPLE
        ON UPDATE NO ACTION
        ON DELETE NO ACTION
)


`

// CREATE TABLE IF NOT EXISTS public.apps
// (
//     id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
//     duration numeric,
//     doc_id integer,
//     pat_id integer,
//     CONSTRAINT apps_pkey PRIMARY KEY (id),
//     CONSTRAINT doc_id FOREIGN KEY (doc_id)
//         REFERENCES public.doctors (id) MATCH SIMPLE
//         ON UPDATE NO ACTION
//         ON DELETE NO ACTION,
//     CONSTRAINT pat_id FOREIGN KEY (pat_id)
//         REFERENCES public.patients (id) MATCH SIMPLE
//         ON UPDATE NO ACTION
//         ON DELETE NO ACTION
// )

// CREATE TABLE IF NOT EXISTS public.admin
// (
//     id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
//     username character varying(50) COLLATE pg_catalog."default" NOT NULL,
//     password character varying(50) COLLATE pg_catalog."default" NOT NULL,
//     CONSTRAINT admin_pkey PRIMARY KEY (id)
// )

// CREATE TABLE IF NOT EXISTS public.doctors
// (
//     id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
//     username character varying(50) COLLATE pg_catalog."default" NOT NULL,
//     password character varying(50) COLLATE pg_catalog."default" NOT NULL,
//     time_left double precision NOT NULL DEFAULT 8,
//     CONSTRAINT doctors_pkey PRIMARY KEY (id)
// )

// CREATE TABLE IF NOT EXISTS public.patients
// (
//     id integer NOT NULL GENERATED ALWAYS AS IDENTITY ( INCREMENT 1 START 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1 ),
//     username character varying(50) COLLATE pg_catalog."default" NOT NULL,
//     password character varying(50) COLLATE pg_catalog."default" NOT NULL,
//     CONSTRAINT patients_pkey PRIMARY KEY (id)
// )
