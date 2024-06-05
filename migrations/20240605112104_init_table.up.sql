
-- public.user_account definition

-- Drop table

-- DROP TABLE public.user_account;

CREATE TABLE public.user_account (
	id uuid NOT NULL DEFAULT gen_random_uuid(),
	email varchar NOT NULL,
	"password" varchar NOT NULL,
	salt varchar NOT NULL,
	"name" varchar NOT NULL,
	phone varchar NOT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	updated_at timestamp NOT NULL DEFAULT now(),
	is_deleted bool NOT NULL DEFAULT false,
	CONSTRAINT user_pkey PRIMARY KEY (id)
);

-- public.authentication_tokens definition

-- Drop table

-- DROP TABLE public.authentication_tokens;

CREATE TABLE public.authentication_tokens (
	token_id uuid NOT NULL DEFAULT gen_random_uuid(),
	user_account_id uuid NOT NULL,
	auth_token varchar NOT NULL,
	generated_at timestamp NULL,
	expires_at timestamp NULL,
	CONSTRAINT authentication_tokens_pkey PRIMARY KEY (token_id)
);


-- public.authentication_tokens foreign keys

ALTER TABLE public.authentication_tokens ADD CONSTRAINT authentication_tokens_fk FOREIGN KEY (user_account_id) REFERENCES public.user_account(id);

-- public.profiles definition

-- Drop table

-- DROP TABLE public.profiles;

CREATE TABLE public.profiles (
	id serial4 NOT NULL,
	first_name varchar NOT NULL,
	last_name varchar NOT NULL,
	gender varchar NOT NULL,
	birth_date timestamp NOT NULL,
	phone_number varchar NOT NULL,
	nickname varchar NOT NULL,
	description varchar NULL,
	photo varchar NULL,
	is_premium bool NOT NULL DEFAULT false,
	updated_at timestamp NOT NULL DEFAULT now(),
	created_at timestamp NOT NULL DEFAULT now(),
	is_deleted bool NOT NULL DEFAULT false,
	user_account_id uuid NOT NULL,
	is_verified bool NOT NULL DEFAULT false,
	CONSTRAINT profile_pk PRIMARY KEY (id)
);
-- public.profiles foreign keys

ALTER TABLE public.profiles ADD CONSTRAINT profiles_fk FOREIGN KEY (user_account_id) REFERENCES public.user_account(id) ON UPDATE CASCADE;

-- public.action_type_swipe definition

-- Drop table

-- DROP TABLE public.action_type_swipe;

CREATE TABLE public.action_type_swipe (
	id serial4 NOT NULL,
	"action" varchar NOT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	created_by varchar NOT NULL,
	updated_at timestamp NOT NULL DEFAULT now(),
	updated_by varchar NOT NULL,
	deleted_at timestamp NULL,
	deleted_by varchar NULL,
	is_deleted bool NOT NULL DEFAULT false,
	description varchar NULL,
	CONSTRAINT action_type_swipe_pk PRIMARY KEY (id)
);

-- public.swipes definition

-- Drop table

-- DROP TABLE public.swipes;

CREATE TABLE public.swipes (
	id serial4 NOT NULL,
	sender_profile_id int4 NOT NULL,
	receiver_profile_id int4 NOT NULL,
	action_type_id int4 NOT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	created_by varchar NOT NULL,
	updated_at timestamp NOT NULL DEFAULT now(),
	updated_by varchar NOT NULL,
	deleted_at timestamp NULL,
	deleted_by varchar NULL,
	is_deleted bool NOT NULL DEFAULT false,
	CONSTRAINT swipes_pk PRIMARY KEY (id)
);


-- public.swipes foreign keys

ALTER TABLE public.swipes ADD CONSTRAINT swipes_fk FOREIGN KEY (sender_profile_id) REFERENCES public.profiles(id);
ALTER TABLE public.swipes ADD CONSTRAINT swipes_fk_1 FOREIGN KEY (receiver_profile_id) REFERENCES public.profiles(id);
ALTER TABLE public.swipes ADD CONSTRAINT swipes_fk_2 FOREIGN KEY (action_type_id) REFERENCES public.action_type_swipe(id);

-- public.premium_packages definition

-- Drop table

-- DROP TABLE public.premium_packages;

CREATE TABLE public.premium_packages (
	id serial4 NOT NULL,
	"name" varchar NOT NULL,
	description varchar NOT NULL,
	price numeric NOT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	created_by varchar NOT NULL,
	updated_at timestamp NOT NULL DEFAULT now(),
	updated_by varchar NOT NULL,
	deleted_at timestamp NULL,
	deleted_by varchar NULL,
	is_deleted bool NOT NULL DEFAULT false,
	CONSTRAINT premium_packages_pk PRIMARY KEY (id)
);

-- public.purchase_histories definition

-- Drop table

-- DROP TABLE public.purchase_histories;

CREATE TABLE public.purchase_histories (
	id serial4 NOT NULL,
	profile_id int4 NOT NULL,
	premium_package_id int4 NOT NULL,
	created_at timestamp NOT NULL DEFAULT now(),
	created_by varchar NOT NULL,
	updated_at timestamp NOT NULL DEFAULT now(),
	updated_by varchar NOT NULL,
	deleted_at timestamp NULL,
	deleted_by varchar NULL,
	is_deleted bool NOT NULL DEFAULT false,
	CONSTRAINT purchase_histories_pk PRIMARY KEY (id)
);


-- public.purchase_histories foreign keys

ALTER TABLE public.purchase_histories ADD CONSTRAINT purchase_histories_fk FOREIGN KEY (profile_id) REFERENCES public.profiles(id);
ALTER TABLE public.purchase_histories ADD CONSTRAINT purchase_histories_fk_1 FOREIGN KEY (premium_package_id) REFERENCES public.premium_packages(id);


-- insert initial data
INSERT INTO public.premium_packages ("name", description, price, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by, is_deleted) VALUES('Unlimited Swipe', 'No swipe quota for user', 10, now(), 'System', now(), 'System', null, null, false);
INSERT INTO public.premium_packages ("name", description, price, created_at, created_by, updated_at, updated_by, deleted_at, deleted_by, is_deleted) VALUES('Verified Account', 'Verified label for user', 10, now(), 'System', now(), 'System', null, null, false);

INSERT INTO public.action_type_swipe ("action", created_at, created_by, updated_at, updated_by, deleted_at, deleted_by, is_deleted,description) VALUES('Swipe Right', now(), 'System', now(), 'System', null, null, false,'Interested');
INSERT INTO public.action_type_swipe ("action", created_at, created_by, updated_at, updated_by, deleted_at, deleted_by, is_deleted, description) VALUES('Swipe Left', now(), 'System', now(), 'System', null, null, false,'Rejected');

INSERT INTO public.user_account (id,email,"password",salt,"name",phone,created_at,updated_at,is_deleted) VALUES
	 ('3471a78b-5488-4f77-821e-f40759da2309','test123@gmail.com','V9DWHrfjQjU2znV24+wI+r+YiyBlYa221BZ90RzVlrtTBoVEGP4UQWz3l+E8Qn1KjITx2r4uHzktZuwWunN7FA==','h/5piuhHxBHDX0IqbZ4ktMKteKET3obmVCCqFrtZENM','test 123','082123123123','2024-06-04 16:56:23.963963','2024-06-04 16:56:23.963963',false),
	 ('0acc58ac-aaa7-4a6b-8a89-e9e999165c8c','test1234@gmail.com','UrK74yNiSnUqKfjU8XrNQ3+/9bSwIlpjIjHzpDw1aps2y2v2rxK8mtkk32EPQXCoesQ+DQaX/hxZV3vA+s8qfg==','eJEcTXhjURasG/lvFnyMsMyg79NvztlF476lx38b484','test 1234','082123123124','2024-06-05 14:52:23.906687','2024-06-05 14:52:23.906687',false),
	 ('fea1749e-883e-4afb-9caf-94a89d570f24','test1235@gmail.com','tTlYKeilfIzUKQ8IN3y3JmAsnP0/kp4OPiKQL0hvUWi3jAS91MC5o9Al9Q1iQf3OfAFQWWXMDRdPCdeP0eA4oQ==','z6.DesvkBdCHj1x7r9bfxh4.mcuu3ajdKlxPSTJtGN0','test 1235','082123123125','2024-06-05 14:52:43.098462','2024-06-05 14:52:43.098462',false),
	 ('7bd3f3a8-dcaf-44c5-9640-ab3fb19f7e3a','test1236@gmail.com','5e1t6LBPjOU7rFIwz8uq4asQE7+O59Doc2xspkxAlgt7Th5yVj27k8c8/em+oj+G81/bMyh0/6jKuZZx9jhPTQ==','F.eB7geI89AeYSQU0Wf/ZYZg2LVFcIAwyaEXylrTyd0','test 1236','082123123126','2024-06-05 14:53:39.38335','2024-06-05 14:53:39.38335',false),
	 ('0956d5f6-77a5-4216-bd4f-997606aea40a','test1237@gmail.com','eWhFW5Txg1KCmB5sjY59HmtKNjhMgIbQv4Je/rWG4VAvlS/F28J6FGc8auCdFEdPSjWqug8RpFR/faGYnaKPMg==','Xn0LVKW0af5w0w5dJ9xqH9RhrHr5NRvl0AUaVcQN9Cw','test 1237','082123123127','2024-06-05 14:53:59.652149','2024-06-05 14:53:59.652149',false),
	 ('3dd260ae-17a2-4a7c-9686-00649ce79293','test1238@gmail.com','l3CN4OYiwJeYNa75tG8VAgsCy0+9n4d5Gx9a6OKeY1rjKBE68+Im1IBBL7v9X4i0JfSAakzajXfmAEW+SEKqhQ==','Pf/VvQ.aij4BWJcVgSb8xvlf9kHr2OQv2KLkwaW6mGI','test 1238','082123123128','2024-06-05 14:54:13.346642','2024-06-05 14:54:13.346642',false),
	 ('089e5949-b31a-4520-b428-627953181586','test1239@gmail.com','6J3QdPDHFT/I+T8q5vmNtqZDels7m6WnY/O9ZqQjlLhBj0KRJEcLnVMg8yc4MlO+O5R/4GJ/aEABYmCNPUobkw==','12LxSCssTqhOS2WM5eKuIP0v/6cyorczwEz4VBAITCo','test 1239','082123123129','2024-06-05 14:54:38.266103','2024-06-05 14:54:38.266103',false),
	 ('aa506c85-6077-4d83-9f2a-c470e4ff13bf','test12391@gmail.com','N/Gk6J70Md/wQ6i9nPh3Qnb2+q5uSsinN/0CFMPQSFRoaPfTtgj30ygWKl4CLLKz0xY++dPai5yP8RgYP1kkdg==','yy.3ZbSFIyMqqO059JcjD5NVhe0YO1VLEhUS/LXkKRo','test 12391','0821231231291','2024-06-05 14:54:50.485598','2024-06-05 14:54:50.485598',false),
	 ('e647df1d-3dd4-48ee-96e9-499dc39fb0df','test12392@gmail.com','LyG0EmE2FH9E7sXL5xlahNMKwFJfMyA9UShSgtNUMkw1Fe1mXvUnaJGxHnCgRO0RQE7iIU/WL0uhPwXNOVunMQ==','ji6/sIxuulDzjsDtUb3g3CCqjDbe3ESCV5mRdc6rD/E','test 12392','0821231231292','2024-06-05 14:55:01.049388','2024-06-05 14:55:01.049388',false),
	 ('fe508b2c-83d5-45fc-88ff-e64f2649b384','test12393@gmail.com','OukAw+07X8HszjUXSF06S4mtgM5EjHxwPo2Ro6qbBO3BWTTYcW0dNTwtCShNyPcW/t2nFhDTkFNRTE0ZX5Yavw==','j5n8Xtn57AtUkt23mos0NdaMR2qTCSnGOyhpCxHyYMs','test 12393','0821231231293','2024-06-05 15:40:52.079793','2024-06-05 15:40:52.079793',false);
INSERT INTO public.user_account (id,email,"password",salt,"name",phone,created_at,updated_at,is_deleted) VALUES
	 ('f3b705fb-0fb1-4df5-8a19-4c701505ba09','test12394@gmail.com','wcxxqflY9R2RP5Nn07xa+DdHMi9GqExx5ofoLqrxIzMu32jZp/uPdLAMKqdAloOcqPKZkvYlXTcVPb9RUIMOwA==','KjvEE764hb8MbsckN.cO.n4KEh/HqFYqcmv9JPMTBXo','test 12394','0821231231294','2024-06-05 15:43:07.79162','2024-06-05 15:43:07.79162',false),
	 ('99ca6391-c400-40d7-8d5b-146819595254','test12395@gmail.com','dt8GdcKoXq5ueUvXCizW0yM263ZqxtO+wDI4nG1xHX6ZbuEQu4YmxLPBxOxLobnI2q+bsYnVmo1j+rPESFOFOQ==','J5Ji3C7mRs8UDkSYb.kTjjBI5zOgijxVB2uSUZrRXDU','test 12395','0821231231295','2024-06-05 15:43:29.293701','2024-06-05 15:43:29.293701',false),
	 ('2e234479-6455-47d1-b0dc-585f3cf2f3eb','test12396@gmail.com','L7ctTVzfTev4A4Lkazxb/M/ssUiWwVc8D0dgQd8uuKxQTkbyI4CelAXafy5dQhoBUAphOCGdV6XyBvkMd6hbSw==','wDoBtVBKQ/uEUZc3hypbp4Q7HX1mrsVCuKTyEDWxUZM','test 12396','0821231231296','2024-06-05 15:43:43.90375','2024-06-05 15:43:43.90375',false);
INSERT INTO public.profiles (first_name,last_name,gender,birth_date,phone_number,nickname,description,photo,is_premium,updated_at,created_at,is_deleted,user_account_id,is_verified) VALUES
	 ('test','123','male','0001-01-01 00:00:00','082123123123','test123','testt','test',true,'2024-06-04 21:31:21.919389','2024-06-04 16:56:23.967965',false,'3471a78b-5488-4f77-821e-f40759da2309',false),
	 ('test','1234','female','0001-01-01 00:00:00','082123123124','test1234','testt','test',false,'2024-06-05 14:52:23.946434','2024-06-05 14:52:23.946434',false,'0acc58ac-aaa7-4a6b-8a89-e9e999165c8c',false),
	 ('test','1235','female','0001-01-01 00:00:00','082123123125','test1235','testt','test',false,'2024-06-05 14:52:43.104072','2024-06-05 14:52:43.104072',false,'fea1749e-883e-4afb-9caf-94a89d570f24',false),
	 ('test','1236','female','1998-01-22 00:00:00','082123123126','test1236','testt','test',false,'2024-06-05 14:53:39.387767','2024-06-05 14:53:39.387767',false,'7bd3f3a8-dcaf-44c5-9640-ab3fb19f7e3a',false),
	 ('test','1237','female','1998-01-24 00:00:00','082123123127','test1237','testt','test',false,'2024-06-05 14:53:59.656462','2024-06-05 14:53:59.656462',false,'0956d5f6-77a5-4216-bd4f-997606aea40a',false),
	 ('test','1238','female','1998-02-24 00:00:00','082123123128','test1238','testt','test',false,'2024-06-05 14:54:13.348732','2024-06-05 14:54:13.348732',false,'3dd260ae-17a2-4a7c-9686-00649ce79293',false),
	 ('test','1239','female','1998-03-27 00:00:00','082123123129','test1239','testt','test',false,'2024-06-05 14:54:38.270521','2024-06-05 14:54:38.270521',false,'089e5949-b31a-4520-b428-627953181586',false),
	 ('test','12391','female','1998-04-27 00:00:00','0821231231291','test12391','testt','test',false,'2024-06-05 14:54:50.489511','2024-06-05 14:54:50.489511',false,'aa506c85-6077-4d83-9f2a-c470e4ff13bf',false),
	 ('test','12392','female','1998-04-27 00:00:00','0821231231292','test12392','testt','test',false,'2024-06-05 14:55:01.054232','2024-06-05 14:55:01.054232',false,'e647df1d-3dd4-48ee-96e9-499dc39fb0df',false),
	 ('test','12393','male','1998-04-27 00:00:00','0821231231293','test12393','testt','test',false,'2024-06-05 15:40:52.082929','2024-06-05 15:40:52.082929',false,'fe508b2c-83d5-45fc-88ff-e64f2649b384',false);
INSERT INTO public.profiles (first_name,last_name,gender,birth_date,phone_number,nickname,description,photo,is_premium,updated_at,created_at,is_deleted,user_account_id,is_verified) VALUES
	 ('test','12394','female','1998-06-27 00:00:00','0821231231294','test12394','testt','test',false,'2024-06-05 15:43:07.797091','2024-06-05 15:43:07.797091',false,'f3b705fb-0fb1-4df5-8a19-4c701505ba09',false),
	 ('test','12395','female','1998-07-27 00:00:00','0821231231295','test12395','testt','test',false,'2024-06-05 15:43:29.298615','2024-06-05 15:43:29.298615',false,'99ca6391-c400-40d7-8d5b-146819595254',false),
	 ('test','12396','female','1998-07-27 00:00:00','0821231231296','test12396','testt','test',false,'2024-06-05 15:43:43.907944','2024-06-05 15:43:43.907944',false,'2e234479-6455-47d1-b0dc-585f3cf2f3eb',false);

INSERT INTO public.purchase_histories (profile_id,premium_package_id,created_at,created_by,updated_at,updated_by,deleted_at,deleted_by,is_deleted) VALUES
	 (1,1,'2024-06-04 21:31:21.919389','System','2024-06-04 21:31:21.919389','System',NULL,NULL,false);

INSERT INTO public.swipes (sender_profile_id,receiver_profile_id,action_type_id,created_at,created_by,updated_at,updated_by,deleted_at,deleted_by,is_deleted) VALUES
	 (11,8,1,'2024-06-05 15:54:04.933999','test 123','2024-06-05 15:54:04.933999','test 123',NULL,NULL,false),
	 (11,4,1,'2024-06-05 15:36:19.610207','test 123','2024-06-05 15:36:19.610207','test 123',NULL,NULL,false),
	 (11,3,1,'2024-06-05 15:54:13.122646','test 123','2024-06-05 15:54:13.122646','test 123',NULL,NULL,false),
	 (11,5,1,'2024-06-05 15:54:18.059539','test 123','2024-06-05 15:54:18.059539','test 123',NULL,NULL,false),
	 (11,6,1,'2024-06-05 15:54:22.65085','test 123','2024-06-05 15:54:22.65085','test 123',NULL,NULL,false),
	 (11,7,1,'2024-06-05 15:54:25.82973','test 123','2024-06-05 15:54:25.82973','test 123',NULL,NULL,false),
	 (11,9,1,'2024-06-05 15:54:30.621046','test 123','2024-06-05 15:54:30.621046','test 123',NULL,NULL,false),
	 (11,10,1,'2024-06-05 15:54:47.003029','test 123','2024-06-05 15:54:47.003029','test 123',NULL,NULL,false),
	 (11,12,1,'2024-06-05 15:57:19.788523','test 123','2024-06-05 15:57:19.788523','test 123',NULL,NULL,false),
	 (11,13,1,'2024-06-05 15:57:24.251714','test 123','2024-06-05 15:57:24.251714','test 123',NULL,NULL,false);



