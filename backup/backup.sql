--
-- PostgreSQL database dump
--

\restrict NW6t73jc0tt09fxm5DsSY8dyiILWLFIH9BgOEx9DRBDSd1FvOc91ENBDRWk2v2f

-- Dumped from database version 18.1
-- Dumped by pg_dump version 18.1

-- Started on 2025-12-14 22:55:25

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
SET transaction_timeout = 0;
SET client_encoding = 'UTF8';
SET standard_conforming_strings = on;
SELECT pg_catalog.set_config('search_path', '', false);
SET check_function_bodies = false;
SET xmloption = content;
SET client_min_messages = warning;
SET row_security = off;

SET default_tablespace = '';

SET default_table_access_method = heap;

--
-- TOC entry 220 (class 1259 OID 17142)
-- Name: categories; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.categories (
    id integer NOT NULL,
    name character varying(150) NOT NULL,
    description text,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.categories OWNER TO postgres;

--
-- TOC entry 219 (class 1259 OID 17141)
-- Name: categories_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.categories_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.categories_id_seq OWNER TO postgres;

--
-- TOC entry 5049 (class 0 OID 0)
-- Dependencies: 219
-- Name: categories_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.categories_id_seq OWNED BY public.categories.id;


--
-- TOC entry 224 (class 1259 OID 17184)
-- Name: item_activities; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.item_activities (
    id integer NOT NULL,
    item_id integer NOT NULL,
    activity_type character varying(50) NOT NULL,
    start_date date NOT NULL,
    end_date date,
    days integer,
    note text,
    created_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone
);


ALTER TABLE public.item_activities OWNER TO postgres;

--
-- TOC entry 223 (class 1259 OID 17183)
-- Name: item_activities_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.item_activities_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.item_activities_id_seq OWNER TO postgres;

--
-- TOC entry 5050 (class 0 OID 0)
-- Dependencies: 223
-- Name: item_activities_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.item_activities_id_seq OWNED BY public.item_activities.id;


--
-- TOC entry 222 (class 1259 OID 17157)
-- Name: items; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.items (
    id integer NOT NULL,
    category_id integer NOT NULL,
    name character varying(255) NOT NULL,
    sku character varying(100) NOT NULL,
    purchase_price numeric(14,2) NOT NULL,
    purchase_date date NOT NULL,
    depreciation_rate numeric(5,4) DEFAULT 0.20 NOT NULL,
    useful_life_days integer,
    current_status character varying(50) DEFAULT 'active'::character varying,
    note text,
    created_at timestamp with time zone DEFAULT now(),
    updated_at timestamp with time zone DEFAULT now(),
    deleted_at timestamp with time zone,
    CONSTRAINT items_purchase_price_check CHECK ((purchase_price >= (0)::numeric))
);


ALTER TABLE public.items OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 17156)
-- Name: items_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.items_id_seq
    AS integer
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER SEQUENCE public.items_id_seq OWNER TO postgres;

--
-- TOC entry 5051 (class 0 OID 0)
-- Dependencies: 221
-- Name: items_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.items_id_seq OWNED BY public.items.id;


--
-- TOC entry 4866 (class 2604 OID 17145)
-- Name: categories id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories ALTER COLUMN id SET DEFAULT nextval('public.categories_id_seq'::regclass);


--
-- TOC entry 4874 (class 2604 OID 17187)
-- Name: item_activities id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.item_activities ALTER COLUMN id SET DEFAULT nextval('public.item_activities_id_seq'::regclass);


--
-- TOC entry 4869 (class 2604 OID 17160)
-- Name: items id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items ALTER COLUMN id SET DEFAULT nextval('public.items_id_seq'::regclass);


--
-- TOC entry 5039 (class 0 OID 17142)
-- Dependencies: 220
-- Data for Name: categories; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.categories (id, name, description, created_at, updated_at, deleted_at) FROM stdin;
1	Laptop & Komputer	Perangkat komputasi utama untuk karyawan	2025-12-13 10:26:47.132962+07	2025-12-13 10:26:47.132962+07	\N
2	Monitor	Monitor kerja kantor	2025-12-13 10:26:47.132962+07	2025-12-13 10:26:47.132962+07	\N
3	Printer & Scanner	Perangkat cetak dan pindai dokumen	2025-12-13 10:26:47.132962+07	2025-12-13 10:26:47.132962+07	\N
4	Peralatan Jaringan	Router, switch, access point dll	2025-12-13 10:26:47.132962+07	2025-12-13 10:26:47.132962+07	\N
5	Elektronik Kantor	TV, projector, CCTV, speaker, dll	2025-12-13 10:26:47.132962+07	2025-12-13 10:26:47.132962+07	\N
6	Furniture Kantor	Kursi, meja, lemari arsip, dll.	2025-12-13 10:26:47.132962+07	2025-12-13 10:26:47.132962+07	\N
7	Peralatan Kebersihan	Vacuum, alat pel listrik, dll.	2025-12-13 10:26:47.132962+07	2025-12-13 10:26:47.132962+07	\N
8	Kendaraan Operasional	Motor dan mobil kantor	2025-12-13 10:26:47.132962+07	2025-12-13 10:26:47.132962+07	\N
9	Peralatan Umum Kantor	Tool kit, tangga, perlengkapan kerja	2025-12-13 10:26:47.132962+07	2025-12-13 10:26:47.132962+07	\N
10	Aksesoris	Mouse, keyboard, kabel	2025-12-13 10:26:47.132962+07	2025-12-13 17:13:49.858531+07	\N
18	Data Updated	Deskripsi Baru	2025-12-13 21:27:13.323525+07	2025-12-13 23:25:12.167711+07	2025-12-13 23:25:12.167711+07
19	Data Baru yang baru\r\n	Deskripsi yang paling baru\r\n	2025-12-13 21:30:14.81134+07	2025-12-13 23:30:14.806318+07	2025-12-13 23:30:14.806318+07
17	Testing Name Terbaru	Testing Description\r\n	2025-12-13 21:20:39.855378+07	2025-12-13 23:42:26.826386+07	2025-12-13 23:42:26.826386+07
20	Pakaian\r\n	Seragam Kantor, Seragam Karyawan\r\n	2025-12-13 21:32:56.239017+07	2025-12-13 23:42:39.053549+07	2025-12-13 23:42:39.053549+07
\.


--
-- TOC entry 5043 (class 0 OID 17184)
-- Dependencies: 224
-- Data for Name: item_activities; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.item_activities (id, item_id, activity_type, start_date, end_date, days, note, created_at, deleted_at) FROM stdin;
1	1	assigned	2024-02-15	\N	\N	Dipakai CTO sejak pembelian	2025-12-13 10:30:33.04143+07	\N
2	2	assigned	2023-09-05	\N	\N	Dipakai staf IT	2025-12-13 10:30:33.04143+07	\N
3	3	returned	2024-01-10	2024-01-10	1	Dikembalikan dan tidak lagi digunakan	2025-12-13 10:30:33.04143+07	\N
4	5	maintenance	2024-07-01	2024-07-03	2	Kalibrasi ulang warna	2025-12-13 10:30:33.04143+07	\N
5	6	maintenance	2024-04-11	2024-04-11	1	Ganti selang tinta	2025-12-13 10:30:33.04143+07	\N
6	7	relocated	2024-03-01	2024-03-01	1	Dipindahkan ke ruang server	2025-12-13 10:30:33.04143+07	\N
7	13	assigned	2024-05-05	2024-05-10	5	Dipakai untuk perbaikan lampu	2025-12-13 10:30:33.04143+07	\N
8	15	maintenance	2024-12-01	2024-12-01	1	Monitor sedang dalam perbaikan	2025-12-13 10:30:33.04143+07	\N
\.


--
-- TOC entry 5041 (class 0 OID 17157)
-- Dependencies: 222
-- Data for Name: items; Type: TABLE DATA; Schema: public; Owner: postgres
--

COPY public.items (id, category_id, name, sku, purchase_price, purchase_date, depreciation_rate, useful_life_days, current_status, note, created_at, updated_at, deleted_at) FROM stdin;
1	1	MacBook Pro 14 M1	LTP-101	28000000.00	2024-02-10	0.2000	1095	active	Dipakai CTO	2025-12-13 10:28:52.320542+07	2025-12-13 10:28:52.320542+07	\N
2	1	Lenovo ThinkPad T14	LTP-102	16000000.00	2023-09-01	0.2000	1095	active	Dipakai staf IT	2025-12-13 10:28:52.320542+07	2025-12-13 10:28:52.320542+07	\N
3	1	ASUS ExpertBook B5	LTP-103	14000000.00	2022-11-20	0.2000	1095	retired	Performa melemah	2025-12-13 10:28:52.320542+07	2025-12-13 10:28:52.320542+07	\N
4	2	Dell Ultrasharp 24 Inch	MON-201	4200000.00	2023-03-01	0.2000	1825	active	Monitor premium	2025-12-13 10:28:52.320542+07	2025-12-13 10:28:52.320542+07	\N
5	2	Xiaomi 27 Inch 2K	MON-202	3500000.00	2024-06-10	0.2000	1825	maintenance	Alur warna tidak stabil	2025-12-13 10:28:52.320542+07	2025-12-13 10:28:52.320542+07	\N
6	3	Canon G3010 Ink Tank	PRT-301	2500000.00	2023-01-12	0.2000	1825	active	Sering dipakai divisi admin	2025-12-13 10:28:52.320542+07	2025-12-13 10:28:52.320542+07	\N
7	4	Mikrotik RB4011 Router	NET-401	4800000.00	2023-04-01	0.2000	1825	active	Router utama kantor	2025-12-13 10:28:52.320542+07	2025-12-13 10:28:52.320542+07	\N
8	4	TP-Link Switch 16-Port	NET-402	900000.00	2022-07-15	0.2000	1825	active	Switch lantai 1	2025-12-13 10:28:52.320542+07	2025-12-13 10:28:52.320542+07	\N
9	5	Panasonic Projector XGA	ELK-501	8000000.00	2023-05-20	0.2000	1825	active	Dipakai ruang meeting A	2025-12-13 10:28:52.320542+07	2025-12-13 10:28:52.320542+07	\N
10	5	LED TV Samsung 43 Inch	ELK-502	5200000.00	2022-09-09	0.2000	1825	active	Dipakai ruang training	2025-12-13 10:28:52.320542+07	2025-12-13 10:28:52.320542+07	\N
11	6	Kursi Kantor Ergotron	FUR-601	1500000.00	2023-02-20	0.2000	3650	active	Dipakai staf HR	2025-12-13 10:28:52.320542+07	2025-12-13 10:28:52.320542+07	\N
12	6	Meja Kerja Premium	FUR-602	3000000.00	2022-01-10	0.2000	3650	active	Meja CEO	2025-12-13 10:28:52.320542+07	2025-12-13 10:28:52.320542+07	\N
14	8	Motor Honda Beat 2023	VEH-801	18000000.00	2023-01-15	0.2000	2555	active	Motor kurir internal	2025-12-13 10:28:52.320542+07	2025-12-13 10:28:52.320542+07	\N
15	2	LG Monitor 22 Inch	MON-203	1800000.00	2021-12-10	0.2000	1825	active	jarang digunakan	2025-12-13 10:28:52.320542+07	2025-12-13 10:28:52.320542+07	\N
17	1	invalid name	PCS-101	15000000.00	2025-09-11	0.2000	1095	active		2025-12-14 19:02:11.467335+07	2025-12-14 19:23:15.26028+07	2025-12-14 19:23:15.26028+07
13	9	Tangga Aluminium 2m	GEN-901	600000.00	2023-04-04	0.2000	3650	maintenance	Kebutuhan Teknisi untuk memanjat	2025-12-13 10:28:52.320542+07	2025-12-14 20:15:17.724036+07	\N
\.


--
-- TOC entry 5052 (class 0 OID 0)
-- Dependencies: 219
-- Name: categories_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.categories_id_seq', 20, true);


--
-- TOC entry 5053 (class 0 OID 0)
-- Dependencies: 223
-- Name: item_activities_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.item_activities_id_seq', 1, false);


--
-- TOC entry 5054 (class 0 OID 0)
-- Dependencies: 221
-- Name: items_id_seq; Type: SEQUENCE SET; Schema: public; Owner: postgres
--

SELECT pg_catalog.setval('public.items_id_seq', 17, true);


--
-- TOC entry 4878 (class 2606 OID 17155)
-- Name: categories categories_name_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_name_key UNIQUE (name);


--
-- TOC entry 4881 (class 2606 OID 17153)
-- Name: categories categories_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.categories
    ADD CONSTRAINT categories_pkey PRIMARY KEY (id);


--
-- TOC entry 4888 (class 2606 OID 17196)
-- Name: item_activities item_activities_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.item_activities
    ADD CONSTRAINT item_activities_pkey PRIMARY KEY (id);


--
-- TOC entry 4883 (class 2606 OID 17175)
-- Name: items items_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_pkey PRIMARY KEY (id);


--
-- TOC entry 4885 (class 2606 OID 17177)
-- Name: items items_sku_key; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_sku_key UNIQUE (sku);


--
-- TOC entry 4879 (class 1259 OID 17204)
-- Name: categories_name_unique; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX categories_name_unique ON public.categories USING btree (lower((name)::text)) WHERE (deleted_at IS NULL);


--
-- TOC entry 4886 (class 1259 OID 17207)
-- Name: items_sku_unique; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX items_sku_unique ON public.items USING btree (sku) WHERE (deleted_at IS NULL);


--
-- TOC entry 4890 (class 2606 OID 17197)
-- Name: item_activities item_activities_item_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.item_activities
    ADD CONSTRAINT item_activities_item_id_fkey FOREIGN KEY (item_id) REFERENCES public.items(id) ON DELETE CASCADE;


--
-- TOC entry 4889 (class 2606 OID 17178)
-- Name: items items_category_id_fkey; Type: FK CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.items
    ADD CONSTRAINT items_category_id_fkey FOREIGN KEY (category_id) REFERENCES public.categories(id) ON DELETE RESTRICT;


-- Completed on 2025-12-14 22:55:25

--
-- PostgreSQL database dump complete
--

\unrestrict NW6t73jc0tt09fxm5DsSY8dyiILWLFIH9BgOEx9DRBDSd1FvOc91ENBDRWk2v2f

