-- Adminer 4.8.1 PostgreSQL 15.1 (Debian 15.1-1.pgdg110+1) dump

CREATE SEQUENCE analise_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."analise" (
    "id" integer DEFAULT nextval('analise_id_seq') NOT NULL,
    "iten_id" integer NOT NULL,
    "fornecedor_id" integer NOT NULL,
    "atende" boolean,
    "obs" character varying(256),
    "analista" character varying(45),
    CONSTRAINT "analise_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


CREATE SEQUENCE componente_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."componente" (
    "id" integer DEFAULT nextval('componente_id_seq') NOT NULL,
    "nome" character varying(45) NOT NULL,
    "modelo_id" integer NOT NULL,
    CONSTRAINT "componente_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


CREATE SEQUENCE equipamento_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."equipamento" (
    "id" integer DEFAULT nextval('equipamento_id_seq') NOT NULL,
    "nome" character varying(45) NOT NULL,
    "tipo_id" integer NOT NULL,
    "modelo_id" integer NOT NULL,
    CONSTRAINT "equipamento_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


CREATE SEQUENCE fornecedor_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."fornecedor" (
    "id" integer DEFAULT nextval('fornecedor_id_seq') NOT NULL,
    "nome" character varying(45) NOT NULL,
    CONSTRAINT "fornecedor_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


CREATE SEQUENCE iten_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."iten" (
    "id" integer DEFAULT nextval('iten_id_seq') NOT NULL,
    "equipamento_id" integer NOT NULL,
    "pregao_numero" integer NOT NULL,
    "pregao_ano" integer NOT NULL,
    "solicitante_id" integer NOT NULL,
    CONSTRAINT "iten_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


CREATE SEQUENCE marca_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."marca" (
    "id" integer DEFAULT nextval('marca_id_seq') NOT NULL,
    "nome" character varying(45) NOT NULL,
    CONSTRAINT "marca_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


CREATE SEQUENCE modelo_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."modelo" (
    "id" integer DEFAULT nextval('modelo_id_seq') NOT NULL,
    "nome" character varying(45) NOT NULL,
    "marca_id" integer NOT NULL,
    CONSTRAINT "modelo_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


CREATE TABLE "public"."peca" (
    "equipamento_id" integer NOT NULL,
    "componente_id" integer NOT NULL
) WITH (oids = false);


CREATE TABLE "public"."pregao" (
    "username" character varying(16) NOT NULL,
    "create_time" timestamp DEFAULT CURRENT_TIMESTAMP,
    "update_time" timestamp,
    "numero" integer NOT NULL,
    "ano" integer NOT NULL,
    CONSTRAINT "pregao_pkey" PRIMARY KEY ("numero", "ano")
) WITH (oids = false);


CREATE SEQUENCE solicitante_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."solicitante" (
    "id" integer DEFAULT nextval('solicitante_id_seq') NOT NULL,
    "nome" character varying(45) NOT NULL,
    CONSTRAINT "solicitante_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


CREATE SEQUENCE tipo_id_seq INCREMENT 1 MINVALUE 1 MAXVALUE 2147483647 CACHE 1;

CREATE TABLE "public"."tipo" (
    "id" integer DEFAULT nextval('tipo_id_seq') NOT NULL,
    "nome" character varying(45) NOT NULL,
    CONSTRAINT "tipo_pkey" PRIMARY KEY ("id")
) WITH (oids = false);


ALTER TABLE ONLY "public"."analise" ADD CONSTRAINT "analise_fornecedor_id_fkey" FOREIGN KEY (fornecedor_id) REFERENCES fornecedor(id) NOT DEFERRABLE;
ALTER TABLE ONLY "public"."analise" ADD CONSTRAINT "analise_iten_id_fkey" FOREIGN KEY (iten_id) REFERENCES iten(id) NOT DEFERRABLE;

ALTER TABLE ONLY "public"."componente" ADD CONSTRAINT "componente_modelo_id_fkey" FOREIGN KEY (modelo_id) REFERENCES modelo(id) NOT DEFERRABLE;

ALTER TABLE ONLY "public"."equipamento" ADD CONSTRAINT "equipamento_modelo_id_fkey" FOREIGN KEY (modelo_id) REFERENCES modelo(id) NOT DEFERRABLE;
ALTER TABLE ONLY "public"."equipamento" ADD CONSTRAINT "equipamento_tipo_id_fkey" FOREIGN KEY (tipo_id) REFERENCES tipo(id) NOT DEFERRABLE;

ALTER TABLE ONLY "public"."iten" ADD CONSTRAINT "iten_equipamento_id_fkey" FOREIGN KEY (equipamento_id) REFERENCES equipamento(id) NOT DEFERRABLE;
ALTER TABLE ONLY "public"."iten" ADD CONSTRAINT "iten_pregao_numero_pregao_ano_fkey" FOREIGN KEY (pregao_numero, pregao_ano) REFERENCES pregao(numero, ano) NOT DEFERRABLE;
ALTER TABLE ONLY "public"."iten" ADD CONSTRAINT "iten_solicitante_id_fkey" FOREIGN KEY (solicitante_id) REFERENCES solicitante(id) NOT DEFERRABLE;

ALTER TABLE ONLY "public"."modelo" ADD CONSTRAINT "modelo_marca_id_fkey" FOREIGN KEY (marca_id) REFERENCES marca(id) NOT DEFERRABLE;

ALTER TABLE ONLY "public"."peca" ADD CONSTRAINT "peca_componente_id_fkey" FOREIGN KEY (componente_id) REFERENCES componente(id) NOT DEFERRABLE;
ALTER TABLE ONLY "public"."peca" ADD CONSTRAINT "peca_equipamento_id_fkey" FOREIGN KEY (equipamento_id) REFERENCES equipamento(id) NOT DEFERRABLE;

-- 2022-12-23 13:08:29.655407+00