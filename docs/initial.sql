-- Database: orders_by

-- DROP DATABASE IF EXISTS orders_by;

CREATE DATABASE orders_by
    WITH
    OWNER = postgres
    ENCODING = 'UTF8'
    LC_COLLATE = 'English_United States.1252'
    LC_CTYPE = 'English_United States.1252'
    TABLESPACE = pg_default
    CONNECTION LIMIT = -1;

-- Table: public.orders

-- DROP TABLE IF EXISTS public.orders;

CREATE TABLE IF NOT EXISTS public.orders
(
    id bigint NOT NULL DEFAULT nextval('orders_id_seq'::regclass),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone,
    ordered_at timestamp with time zone NOT NULL,
    customer_name text COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT orders_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.orders
    OWNER to postgres;

-- SEQUENCE: public.orders_id_seq

-- DROP SEQUENCE IF EXISTS public.orders_id_seq;

CREATE SEQUENCE IF NOT EXISTS public.orders_id_seq
    INCREMENT 1
    START 1
    MINVALUE 1
    MAXVALUE 9223372036854775807
    CACHE 1
    OWNED BY orders.id;

ALTER SEQUENCE public.orders_id_seq
    OWNER TO postgres;
    
-- Table: public.items

-- DROP TABLE IF EXISTS public.items;

CREATE TABLE IF NOT EXISTS public.items
(
    id integer NOT NULL,
    created_at timestamp with time zone NOT NULL,
    updated_at time with time zone,
    order_id integer NOT NULL,
    item_code character varying COLLATE pg_catalog."default" NOT NULL,
    description character varying COLLATE pg_catalog."default" NOT NULL,
    quantity integer NOT NULL,
    CONSTRAINT items_pkey PRIMARY KEY (id)
)

TABLESPACE pg_default;

ALTER TABLE IF EXISTS public.items
    OWNER to postgres;