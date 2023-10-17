--
-- PostgreSQL database dump
--

-- Dumped from database version 13.12
-- Dumped by pg_dump version 13.12

-- Started on 2023-09-29 21:00:27

SET statement_timeout = 0;
SET lock_timeout = 0;
SET idle_in_transaction_session_timeout = 0;
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
-- TOC entry 200 (class 1259 OID 16490)
-- Name: api_queue_details; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.api_queue_details (
    id bigint NOT NULL,
    thenscope text,
    customer_id bigint,
    cid character varying(20),
    status character(1),
    "flowEnded" boolean,
    "ReddisKey" character varying(20),
    "storeResponse" text,
    "TimePicked" timestamp with time zone,
    "TimeFired" timestamp with time zone,
    "timeEnded" timestamp with time zone
);


ALTER TABLE public.api_queue_details OWNER TO postgres;

--
-- TOC entry 201 (class 1259 OID 16496)
-- Name: api_queue_details_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.api_queue_details_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.api_queue_details_id_seq OWNER TO postgres;

--
-- TOC entry 3137 (class 0 OID 0)
-- Dependencies: 201
-- Name: api_queue_details_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.api_queue_details_id_seq OWNED BY public.api_queue_details.id;


--
-- TOC entry 202 (class 1259 OID 16498)
-- Name: buyer_stats; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.buyer_stats (
    id bigint NOT NULL,
    buyer_id bigint,
    per_hour_count bigint,
    per_hour_date timestamp with time zone,
    per_day_count bigint,
    per_day_date date,
    per_week_count bigint,
    per_week_date date,
    per_month_count bigint,
    per_month_date date,
    per_year_count bigint,
    per_year_date date,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.buyer_stats OWNER TO postgres;

--
-- TOC entry 203 (class 1259 OID 16501)
-- Name: buyer_stats_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.buyer_stats_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.buyer_stats_id_seq OWNER TO postgres;

--
-- TOC entry 3138 (class 0 OID 0)
-- Dependencies: 203
-- Name: buyer_stats_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.buyer_stats_id_seq OWNED BY public.buyer_stats.id;


--
-- TOC entry 204 (class 1259 OID 16503)
-- Name: buyergroups; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.buyergroups (
    id bigint NOT NULL,
    "groupName" character varying(100) NOT NULL,
    "buyerIds" character varying(100),
    priority character varying(100),
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.buyergroups OWNER TO postgres;

--
-- TOC entry 205 (class 1259 OID 16506)
-- Name: buyergroups_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.buyergroups_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.buyergroups_id_seq OWNER TO postgres;

--
-- TOC entry 3139 (class 0 OID 0)
-- Dependencies: 205
-- Name: buyergroups_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.buyergroups_id_seq OWNED BY public.buyergroups.id;


--
-- TOC entry 206 (class 1259 OID 16508)
-- Name: buyers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.buyers (
    id bigint NOT NULL,
    name character varying(255) NOT NULL,
    price numeric,
    "webhookUrl" text,
    "webhookMethod" character varying(20),
    "webhookTimeout" character varying(30),
    headers text,
    "requestBody" text,
    "httpAuthUsername" character varying(20),
    "httpAuthPassword" character varying(20),
    "responseModel" text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    "contractName" character varying(100),
    "contractPrice" numeric,
    "contractPriority" bigint,
    "contractCap" bigint,
    "contractCapFrequency" text,
    "contractStatus" boolean,
    "httpAthPassword" character varying(20)
);


ALTER TABLE public.buyers OWNER TO postgres;

--
-- TOC entry 207 (class 1259 OID 16514)
-- Name: buyers_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.buyers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.buyers_id_seq OWNER TO postgres;

--
-- TOC entry 3140 (class 0 OID 0)
-- Dependencies: 207
-- Name: buyers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.buyers_id_seq OWNED BY public.buyers.id;


--
-- TOC entry 208 (class 1259 OID 16516)
-- Name: customer_profiles; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.customer_profiles (
    customer_id bigint NOT NULL,
    user_id bigint,
    edu_level character varying(255),
    edu_interest character varying(255),
    edu_updated_at timestamp with time zone,
    deleted_at timestamp with time zone,
    test_id bigint,
    health_insurance character varying(255),
    health_multiple_prescription character varying(255),
    health_diabetic character varying(255),
    health_cpap character varying(255),
    health_low_income character varying(255),
    health_back_pain character varying(255),
    health_updated_at timestamp with time zone,
    finance_checking_account character varying(255),
    finance_debt_amount character varying(255),
    finance_debt boolean,
    finance_updated_at timestamp with time zone,
    car_owner character varying(255),
    car_updated_at timestamp with time zone,
    car_insurance character varying(255),
    car_license text,
    work_status character varying(255),
    work_updated_at timestamp with time zone,
    edu_jornaya_token character varying(255),
    captcha_score numeric,
    captcha_updated_at timestamp with time zone,
    navy_program character varying(255),
    us_citizen boolean,
    gpa_average numeric,
    fq_score bigint,
    fq_updated_at timestamp with time zone,
    ranch character varying(255),
    brta_filt character varying(255),
    glad1 character varying(255),
    glad2 character varying(255),
    home_owner character varying(255),
    home_updated_at timestamp with time zone,
    debt_type character varying(255),
    edu_grant character varying(255),
    finex character varying(255),
    demply character varying(255),
    clej character varying(255),
    auto_accident character varying(255),
    accident character varying(50),
    medicaid character varying(50),
    medicare character varying(50),
    auto_year character varying(50),
    auto_make character varying(50),
    auto_model character varying(50),
    additional_fields_db text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.customer_profiles OWNER TO postgres;

--
-- TOC entry 209 (class 1259 OID 16522)
-- Name: customer_profiles_customer_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.customer_profiles_customer_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.customer_profiles_customer_id_seq OWNER TO postgres;

--
-- TOC entry 3141 (class 0 OID 0)
-- Dependencies: 209
-- Name: customer_profiles_customer_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.customer_profiles_customer_id_seq OWNED BY public.customer_profiles.customer_id;


--
-- TOC entry 210 (class 1259 OID 16524)
-- Name: customers; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.customers (
    id bigint NOT NULL,
    first_name character varying(255),
    last_name character varying(255),
    age bigint,
    email character varying(50) NOT NULL,
    gender character varying(10),
    address1 character varying(255),
    address2 character varying(255),
    city character varying(255),
    state character varying(100),
    country character varying(100),
    zipcode character varying(50),
    subscribed boolean,
    subscribed_at timestamp with time zone,
    s1 bigint,
    s2 bigint,
    s3 character varying(50),
    user_agent character varying(100),
    optin_ip character varying(50),
    optin_url character varying(100),
    optin_time timestamp with time zone,
    utm_source character varying(100),
    utm_medium character varying(100),
    utm_campaign character varying(100),
    created_at timestamp with time zone NOT NULL,
    updated_at timestamp with time zone,
    phone character varying(20),
    q character varying(255),
    l character varying(255),
    ongage_contact_id character varying(255),
    ongage_status bigint,
    oversight_contact_id character varying(255),
    oversight_status bigint,
    validation character varying(255),
    validation_date timestamp with time zone,
    processed boolean,
    processed_date timestamp with time zone,
    tz text,
    opt_in_phone boolean,
    dob text,
    click_ip text,
    pushnami_id character varying(255),
    optin_phone bigint,
    last_visit timestamp with time zone
);


ALTER TABLE public.customers OWNER TO postgres;

--
-- TOC entry 211 (class 1259 OID 16530)
-- Name: customers_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.customers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.customers_id_seq OWNER TO postgres;

--
-- TOC entry 3142 (class 0 OID 0)
-- Dependencies: 211
-- Name: customers_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.customers_id_seq OWNED BY public.customers.id;


--
-- TOC entry 212 (class 1259 OID 16532)
-- Name: lead_post_history; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.lead_post_history (
    id bigint NOT NULL,
    buyer_id bigint,
    trace_id text,
    buyer_group_id bigint,
    customer_id bigint,
    knowledge_base_name character varying(100),
    rule_manifest_id character varying(255),
    rulemanifest_status character varying(100),
    response_code bigint,
    status character varying(255),
    response text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.lead_post_history OWNER TO postgres;

--
-- TOC entry 213 (class 1259 OID 16538)
-- Name: lead_post_history_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.lead_post_history_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.lead_post_history_id_seq OWNER TO postgres;

--
-- TOC entry 3143 (class 0 OID 0)
-- Dependencies: 213
-- Name: lead_post_history_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.lead_post_history_id_seq OWNED BY public.lead_post_history.id;


--
-- TOC entry 214 (class 1259 OID 16540)
-- Name: national_audience; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.national_audience (
    zeta_id text
);


ALTER TABLE public.national_audience OWNER TO postgres;

--
-- TOC entry 215 (class 1259 OID 16546)
-- Name: post_history; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.post_history (
    id bigint NOT NULL,
    buyer_id bigint,
    buyer_group_id bigint,
    customer_id bigint,
    knowledgebase_name character varying(100),
    rulemanifest_id character varying(255),
    rulemanifest_status character varying(100),
    responsecode bigint,
    resp_msg character varying(255),
    resp_result text,
    created_at timestamp with time zone
);


ALTER TABLE public.post_history OWNER TO postgres;

--
-- TOC entry 216 (class 1259 OID 16552)
-- Name: post_history_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.post_history_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.post_history_id_seq OWNER TO postgres;

--
-- TOC entry 3144 (class 0 OID 0)
-- Dependencies: 216
-- Name: post_history_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.post_history_id_seq OWNED BY public.post_history.id;


--
-- TOC entry 217 (class 1259 OID 16554)
-- Name: rule_queue_details; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.rule_queue_details (
    id bigint NOT NULL,
    "actionName" character varying(100),
    customer_id bigint,
    cid character varying(20),
    "ReddisKey" character varying(20),
    status text,
    endpoint character varying(100),
    method character varying(10),
    "queryParams" text,
    "requestBody" text,
    "goRoutine" character varying(50),
    "TimePicked" timestamp with time zone,
    "TimeFired" timestamp with time zone,
    "timeEnded" timestamp with time zone,
    trace_id text
);


ALTER TABLE public.rule_queue_details OWNER TO postgres;

--
-- TOC entry 218 (class 1259 OID 16560)
-- Name: rule_queue_details_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.rule_queue_details_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.rule_queue_details_id_seq OWNER TO postgres;

--
-- TOC entry 3145 (class 0 OID 0)
-- Dependencies: 218
-- Name: rule_queue_details_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.rule_queue_details_id_seq OWNED BY public.rule_queue_details.id;


--
-- TOC entry 219 (class 1259 OID 16562)
-- Name: rules; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.rules (
    id bigint NOT NULL,
    rule text,
    "ruleName" character varying(100),
    "ruleGroup" character varying(100),
    priority bigint
);


ALTER TABLE public.rules OWNER TO postgres;

--
-- TOC entry 220 (class 1259 OID 16568)
-- Name: rules_Info; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public."rules_Info" (
    id bigint NOT NULL,
    "flowName" character varying(255) NOT NULL,
    manifest text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    name character varying(255)
);


ALTER TABLE public."rules_Info" OWNER TO postgres;

--
-- TOC entry 221 (class 1259 OID 16574)
-- Name: rules_Info_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public."rules_Info_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public."rules_Info_id_seq" OWNER TO postgres;

--
-- TOC entry 3146 (class 0 OID 0)
-- Dependencies: 221
-- Name: rules_Info_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public."rules_Info_id_seq" OWNED BY public."rules_Info".id;


--
-- TOC entry 222 (class 1259 OID 16576)
-- Name: rules_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.rules_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.rules_id_seq OWNER TO postgres;

--
-- TOC entry 3147 (class 0 OID 0)
-- Dependencies: 222
-- Name: rules_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.rules_id_seq OWNED BY public.rules.id;


--
-- TOC entry 223 (class 1259 OID 16578)
-- Name: start_point_config; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.start_point_config (
    id bigint NOT NULL,
    name character varying(255) NOT NULL,
    start_point bigint,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.start_point_config OWNER TO postgres;

--
-- TOC entry 224 (class 1259 OID 16581)
-- Name: start_point_config_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.start_point_config_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.start_point_config_id_seq OWNER TO postgres;

--
-- TOC entry 3148 (class 0 OID 0)
-- Dependencies: 224
-- Name: start_point_config_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.start_point_config_id_seq OWNED BY public.start_point_config.id;


--
-- TOC entry 225 (class 1259 OID 16583)
-- Name: userdetails; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.userdetails (
    id bigint NOT NULL,
    "userName" character varying(100),
    email character varying(50),
    fid character varying(100),
    age bigint,
    phone character varying(13),
    address character varying(500),
    "carBuyingInterest" boolean,
    "homeBuyInterest" boolean,
    password character varying(100)
);


ALTER TABLE public.userdetails OWNER TO postgres;

--
-- TOC entry 226 (class 1259 OID 16589)
-- Name: userdetails_id_seq; Type: SEQUENCE; Schema: public; Owner: postgres
--

CREATE SEQUENCE public.userdetails_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;


ALTER TABLE public.userdetails_id_seq OWNER TO postgres;

--
-- TOC entry 3149 (class 0 OID 0)
-- Dependencies: 226
-- Name: userdetails_id_seq; Type: SEQUENCE OWNED BY; Schema: public; Owner: postgres
--

ALTER SEQUENCE public.userdetails_id_seq OWNED BY public.userdetails.id;


--
-- TOC entry 227 (class 1259 OID 16591)
-- Name: work_flow_json_details; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.work_flow_json_details (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name text NOT NULL,
    json text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);


ALTER TABLE public.work_flow_json_details OWNER TO postgres;

--
-- TOC entry 228 (class 1259 OID 16598)
-- Name: workflow_execution; Type: TABLE; Schema: public; Owner: postgres
--

CREATE TABLE public.workflow_execution (
    trace_id uuid,
    workflow_id uuid,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    payload text
);


ALTER TABLE public.workflow_execution OWNER TO postgres;

--
-- TOC entry 2947 (class 2604 OID 16604)
-- Name: api_queue_details id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.api_queue_details ALTER COLUMN id SET DEFAULT nextval('public.api_queue_details_id_seq'::regclass);


--
-- TOC entry 2948 (class 2604 OID 16605)
-- Name: buyer_stats id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.buyer_stats ALTER COLUMN id SET DEFAULT nextval('public.buyer_stats_id_seq'::regclass);


--
-- TOC entry 2949 (class 2604 OID 16606)
-- Name: buyergroups id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.buyergroups ALTER COLUMN id SET DEFAULT nextval('public.buyergroups_id_seq'::regclass);


--
-- TOC entry 2950 (class 2604 OID 16607)
-- Name: buyers id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.buyers ALTER COLUMN id SET DEFAULT nextval('public.buyers_id_seq'::regclass);


--
-- TOC entry 2951 (class 2604 OID 16608)
-- Name: customer_profiles customer_id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.customer_profiles ALTER COLUMN customer_id SET DEFAULT nextval('public.customer_profiles_customer_id_seq'::regclass);


--
-- TOC entry 2952 (class 2604 OID 16609)
-- Name: customers id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.customers ALTER COLUMN id SET DEFAULT nextval('public.customers_id_seq'::regclass);


--
-- TOC entry 2953 (class 2604 OID 16610)
-- Name: lead_post_history id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.lead_post_history ALTER COLUMN id SET DEFAULT nextval('public.lead_post_history_id_seq'::regclass);


--
-- TOC entry 2954 (class 2604 OID 16611)
-- Name: post_history id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.post_history ALTER COLUMN id SET DEFAULT nextval('public.post_history_id_seq'::regclass);


--
-- TOC entry 2955 (class 2604 OID 16612)
-- Name: rule_queue_details id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.rule_queue_details ALTER COLUMN id SET DEFAULT nextval('public.rule_queue_details_id_seq'::regclass);


--
-- TOC entry 2956 (class 2604 OID 16613)
-- Name: rules id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.rules ALTER COLUMN id SET DEFAULT nextval('public.rules_id_seq'::regclass);


--
-- TOC entry 2957 (class 2604 OID 16614)
-- Name: rules_Info id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."rules_Info" ALTER COLUMN id SET DEFAULT nextval('public."rules_Info_id_seq"'::regclass);


--
-- TOC entry 2958 (class 2604 OID 16615)
-- Name: start_point_config id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.start_point_config ALTER COLUMN id SET DEFAULT nextval('public.start_point_config_id_seq'::regclass);


--
-- TOC entry 2959 (class 2604 OID 16616)
-- Name: userdetails id; Type: DEFAULT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.userdetails ALTER COLUMN id SET DEFAULT nextval('public.userdetails_id_seq'::regclass);


--
-- TOC entry 2962 (class 2606 OID 16618)
-- Name: api_queue_details api_queue_details_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.api_queue_details
    ADD CONSTRAINT api_queue_details_pkey PRIMARY KEY (id);


--
-- TOC entry 2964 (class 2606 OID 16620)
-- Name: buyer_stats buyer_stats_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.buyer_stats
    ADD CONSTRAINT buyer_stats_pkey PRIMARY KEY (id);


--
-- TOC entry 2966 (class 2606 OID 16622)
-- Name: buyergroups buyergroups_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.buyergroups
    ADD CONSTRAINT buyergroups_pkey PRIMARY KEY (id);


--
-- TOC entry 2969 (class 2606 OID 16624)
-- Name: buyers buyers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.buyers
    ADD CONSTRAINT buyers_pkey PRIMARY KEY (id);


--
-- TOC entry 2972 (class 2606 OID 16626)
-- Name: customer_profiles customer_profiles_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.customer_profiles
    ADD CONSTRAINT customer_profiles_pkey PRIMARY KEY (customer_id);


--
-- TOC entry 2974 (class 2606 OID 16628)
-- Name: customers customers_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.customers
    ADD CONSTRAINT customers_pkey PRIMARY KEY (id);


--
-- TOC entry 2983 (class 2606 OID 16630)
-- Name: lead_post_history lead_post_history_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.lead_post_history
    ADD CONSTRAINT lead_post_history_pkey PRIMARY KEY (id);


--
-- TOC entry 2985 (class 2606 OID 16632)
-- Name: post_history post_history_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.post_history
    ADD CONSTRAINT post_history_pkey PRIMARY KEY (id);


--
-- TOC entry 2987 (class 2606 OID 16634)
-- Name: rule_queue_details rule_queue_details_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.rule_queue_details
    ADD CONSTRAINT rule_queue_details_pkey PRIMARY KEY (id);


--
-- TOC entry 2992 (class 2606 OID 16636)
-- Name: rules_Info rules_Info_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public."rules_Info"
    ADD CONSTRAINT "rules_Info_pkey" PRIMARY KEY (id);


--
-- TOC entry 2989 (class 2606 OID 16638)
-- Name: rules rules_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.rules
    ADD CONSTRAINT rules_pkey PRIMARY KEY (id);


--
-- TOC entry 2995 (class 2606 OID 16640)
-- Name: start_point_config start_point_config_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.start_point_config
    ADD CONSTRAINT start_point_config_pkey PRIMARY KEY (id);


--
-- TOC entry 2997 (class 2606 OID 16642)
-- Name: userdetails userdetails_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.userdetails
    ADD CONSTRAINT userdetails_pkey PRIMARY KEY (id);


--
-- TOC entry 3001 (class 2606 OID 16644)
-- Name: work_flow_json_details work_flow_json_details_pkey; Type: CONSTRAINT; Schema: public; Owner: postgres
--

ALTER TABLE ONLY public.work_flow_json_details
    ADD CONSTRAINT work_flow_json_details_pkey PRIMARY KEY (id);


--
-- TOC entry 2967 (class 1259 OID 16645)
-- Name: idx_buyergroups_group_name; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_buyergroups_group_name ON public.buyergroups USING btree ("groupName");


--
-- TOC entry 2970 (class 1259 OID 16646)
-- Name: idx_buyers_name; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_buyers_name ON public.buyers USING btree (name);


--
-- TOC entry 2975 (class 1259 OID 16647)
-- Name: idx_created_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_created_at ON public.customers USING btree (created_at);


--
-- TOC entry 2978 (class 1259 OID 16648)
-- Name: idx_created_at_buyer_id_status; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_created_at_buyer_id_status ON public.lead_post_history USING btree (buyer_id, status, created_at);


--
-- TOC entry 2979 (class 1259 OID 16649)
-- Name: idx_customer_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_customer_id ON public.lead_post_history USING btree (customer_id);


--
-- TOC entry 2976 (class 1259 OID 16650)
-- Name: idx_customers_email; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_customers_email ON public.customers USING btree (email);


--
-- TOC entry 2998 (class 1259 OID 16651)
-- Name: idx_name_created_at; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_name_created_at ON public.work_flow_json_details USING btree (name, created_at);


--
-- TOC entry 2977 (class 1259 OID 16652)
-- Name: idx_phone; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_phone ON public.customers USING btree (phone);


--
-- TOC entry 2980 (class 1259 OID 16653)
-- Name: idx_rule_manifest_id_customer_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_rule_manifest_id_customer_id ON public.lead_post_history USING btree (customer_id, rule_manifest_id);


--
-- TOC entry 2990 (class 1259 OID 16654)
-- Name: idx_rules_Info_flow_name; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX "idx_rules_Info_flow_name" ON public."rules_Info" USING btree ("flowName");


--
-- TOC entry 2993 (class 1259 OID 16655)
-- Name: idx_start_point_config_name; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_start_point_config_name ON public.start_point_config USING btree (name);


--
-- TOC entry 2981 (class 1259 OID 16656)
-- Name: idx_trace_id; Type: INDEX; Schema: public; Owner: postgres
--

CREATE INDEX idx_trace_id ON public.lead_post_history USING btree (trace_id);


--
-- TOC entry 2999 (class 1259 OID 16657)
-- Name: idx_work_flow_json_details_name; Type: INDEX; Schema: public; Owner: postgres
--

CREATE UNIQUE INDEX idx_work_flow_json_details_name ON public.work_flow_json_details USING btree (name);


-- Completed on 2023-09-29 21:00:28

--
-- PostgreSQL database dump complete
--

