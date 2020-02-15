CREATE TABLE public.resultados (
	id integer NOT NULL,
	"data" date null, 
	dezena_1 smallint NOT null,
	dezena_2 smallint NOT null, 
	dezena_3 smallint NOT null, 
	dezena_4 smallint NOT null,
	dezena_5 smallint NOT null,
	dezena_6 smallint NOT null
	
)
WITH (
	OIDS=FALSE
);
