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
CREATE SEQUENCE public.api_queue_details_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.api_queue_details_id_seq OWNED BY public.api_queue_details.id;
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
CREATE SEQUENCE public.buyer_stats_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.buyer_stats_id_seq OWNED BY public.buyer_stats.id;
CREATE TABLE public.buyergroups (
    id bigint NOT NULL,
    "groupName" character varying(100) NOT NULL,
    "buyerIds" character varying(100),
    priority character varying(100),
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);
CREATE SEQUENCE public.buyergroups_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.buyergroups_id_seq OWNED BY public.buyergroups.id;
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
CREATE SEQUENCE public.buyers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.buyers_id_seq OWNED BY public.buyers.id;
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
CREATE SEQUENCE public.customer_profiles_customer_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.customer_profiles_customer_id_seq OWNED BY public.customer_profiles.customer_id;
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
CREATE SEQUENCE public.customers_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.customers_id_seq OWNED BY public.customers.id;
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
CREATE SEQUENCE public.lead_post_history_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.lead_post_history_id_seq OWNED BY public.lead_post_history.id;
CREATE TABLE public.national_audience (
    zeta_id text
);
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
CREATE SEQUENCE public.post_history_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.post_history_id_seq OWNED BY public.post_history.id;
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
CREATE SEQUENCE public.rule_queue_details_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.rule_queue_details_id_seq OWNED BY public.rule_queue_details.id;
CREATE TABLE public.rules (
    id bigint NOT NULL,
    rule text,
    "ruleName" character varying(100),
    "ruleGroup" character varying(100),
    priority bigint
);
CREATE TABLE public."rules_Info" (
    id bigint NOT NULL,
    "flowName" character varying(255) NOT NULL,
    manifest text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    name character varying(255)
);
CREATE SEQUENCE public."rules_Info_id_seq"
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public."rules_Info_id_seq" OWNED BY public."rules_Info".id;
CREATE SEQUENCE public.rules_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.rules_id_seq OWNED BY public.rules.id;
CREATE TABLE public.start_point_config (
    id bigint NOT NULL,
    name character varying(255) NOT NULL,
    start_point bigint,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);
CREATE SEQUENCE public.start_point_config_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.start_point_config_id_seq OWNED BY public.start_point_config.id;
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
CREATE SEQUENCE public.userdetails_id_seq
    START WITH 1
    INCREMENT BY 1
    NO MINVALUE
    NO MAXVALUE
    CACHE 1;
ALTER SEQUENCE public.userdetails_id_seq OWNED BY public.userdetails.id;
CREATE TABLE public.work_flow_json_details (
    id uuid DEFAULT gen_random_uuid() NOT NULL,
    name text NOT NULL,
    json text,
    created_at timestamp with time zone,
    updated_at timestamp with time zone
);
CREATE TABLE public.workflow_execution (
    trace_id uuid,
    workflow_id uuid,
    created_at timestamp with time zone,
    updated_at timestamp with time zone,
    payload text
);
ALTER TABLE ONLY public.api_queue_details ALTER COLUMN id SET DEFAULT nextval('public.api_queue_details_id_seq'::regclass);
ALTER TABLE ONLY public.buyer_stats ALTER COLUMN id SET DEFAULT nextval('public.buyer_stats_id_seq'::regclass);
ALTER TABLE ONLY public.buyergroups ALTER COLUMN id SET DEFAULT nextval('public.buyergroups_id_seq'::regclass);
ALTER TABLE ONLY public.buyers ALTER COLUMN id SET DEFAULT nextval('public.buyers_id_seq'::regclass);
ALTER TABLE ONLY public.customer_profiles ALTER COLUMN customer_id SET DEFAULT nextval('public.customer_profiles_customer_id_seq'::regclass);
ALTER TABLE ONLY public.customers ALTER COLUMN id SET DEFAULT nextval('public.customers_id_seq'::regclass);
ALTER TABLE ONLY public.lead_post_history ALTER COLUMN id SET DEFAULT nextval('public.lead_post_history_id_seq'::regclass);
ALTER TABLE ONLY public.post_history ALTER COLUMN id SET DEFAULT nextval('public.post_history_id_seq'::regclass);
ALTER TABLE ONLY public.rule_queue_details ALTER COLUMN id SET DEFAULT nextval('public.rule_queue_details_id_seq'::regclass);
ALTER TABLE ONLY public.rules ALTER COLUMN id SET DEFAULT nextval('public.rules_id_seq'::regclass);
ALTER TABLE ONLY public."rules_Info" ALTER COLUMN id SET DEFAULT nextval('public."rules_Info_id_seq"'::regclass);
ALTER TABLE ONLY public.start_point_config ALTER COLUMN id SET DEFAULT nextval('public.start_point_config_id_seq'::regclass);
ALTER TABLE ONLY public.userdetails ALTER COLUMN id SET DEFAULT nextval('public.userdetails_id_seq'::regclass);
ALTER TABLE ONLY public.api_queue_details
    ADD CONSTRAINT api_queue_details_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.buyer_stats
    ADD CONSTRAINT buyer_stats_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.buyergroups
    ADD CONSTRAINT buyergroups_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.buyers
    ADD CONSTRAINT buyers_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.customer_profiles
    ADD CONSTRAINT customer_profiles_pkey PRIMARY KEY (customer_id);
ALTER TABLE ONLY public.customers
    ADD CONSTRAINT customers_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.lead_post_history
    ADD CONSTRAINT lead_post_history_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.post_history
    ADD CONSTRAINT post_history_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.rule_queue_details
    ADD CONSTRAINT rule_queue_details_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public."rules_Info"
    ADD CONSTRAINT "rules_Info_pkey" PRIMARY KEY (id);
ALTER TABLE ONLY public.rules
    ADD CONSTRAINT rules_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.start_point_config
    ADD CONSTRAINT start_point_config_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.userdetails
    ADD CONSTRAINT userdetails_pkey PRIMARY KEY (id);
ALTER TABLE ONLY public.work_flow_json_details
    ADD CONSTRAINT work_flow_json_details_pkey PRIMARY KEY (id);
CREATE UNIQUE INDEX idx_buyergroups_group_name ON public.buyergroups USING btree ("groupName");
CREATE UNIQUE INDEX idx_buyers_name ON public.buyers USING btree (name);
CREATE INDEX idx_created_at ON public.customers USING btree (created_at);
CREATE INDEX idx_created_at_buyer_id_status ON public.lead_post_history USING btree (buyer_id, status, created_at);
CREATE INDEX idx_customer_id ON public.lead_post_history USING btree (customer_id);
CREATE UNIQUE INDEX idx_customers_email ON public.customers USING btree (email);
CREATE INDEX idx_name_created_at ON public.work_flow_json_details USING btree (name, created_at);
CREATE INDEX idx_phone ON public.customers USING btree (phone);
CREATE INDEX idx_rule_manifest_id_customer_id ON public.lead_post_history USING btree (customer_id, rule_manifest_id);
CREATE UNIQUE INDEX "idx_rules_Info_flow_name" ON public."rules_Info" USING btree ("flowName");
CREATE UNIQUE INDEX idx_start_point_config_name ON public.start_point_config USING btree (name);
CREATE INDEX idx_trace_id ON public.lead_post_history USING btree (trace_id);
CREATE UNIQUE INDEX idx_work_flow_json_details_name ON public.work_flow_json_details USING btree (name);
