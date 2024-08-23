package main

const exampleHive = "CREATE TABLE `default.mytable`(" +
	"  `id` bigint, " +
	"  `name` varchar(50), " +
	"  `birthday` date, " +
	"  `created` timestamp, " +
	"  `is_manager` boolean, " +
	"  `salary` decimal(9,2), " +
	"  `working_year` int, " +
	"  `ts` timestamp, " +
	"  `t` timestamp)" +
	"ROW FORMAT SERDE " +
	"  'org.apache.hadoop.hive.serde2.lazy.LazySimpleSerDe' " +
	"STORED AS INPUTFORMAT " +
	"  'org.apache.hadoop.mapred.TextInputFormat' " +
	"OUTPUTFORMAT " +
	"  'org.apache.hadoop.hive.ql.io.HiveIgnoreKeyTextOutputFormat'" +
	"LOCATION" +
	"  'file:/opt/hive/data/warehouse/mytable'" +
	"TBLPROPERTIES (" +
	"  'bucketing_version'='2', " +
	"  'last_modified_by'='hive', " +
	"  'last_modified_time'='1702023836', " +
	"  'transient_lastDdlTime'='1720950051');"

const exampleMySQL = `CREATE TABLE test.a1 (
  id bigint NOT NULL AUTO_INCREMENT COMMENT 'id',
  c1 varchar(10) COLLATE utf8mb4_general_ci NOT NULL COMMENT 'c1',
  c2 int DEFAULT NULL COMMENT 'c2',
  c3 integer DEFAULT NULL COMMENT 'c3',
  c4 bigint(11) DEFAULT NULL COMMENT 'c4',
  c5 bigint unsigned DEFAULT NULL COMMENT 'c5',
  c6 tinyint(1) DEFAULT NULL COMMENT 'c6',
  c7 char(5) COLLATE utf8mb4_general_ci DEFAULT NULL COMMENT 'c7',
  c8 date DEFAULT NULL COMMENT 'c8',
  c9 datetime DEFAULT NULL COMMENT 'c9',
  c10 decimal(11,2) DEFAULT NULL COMMENT 'c10',
  c11 double(4,2) DEFAULT NULL COMMENT 'c11',
  c12 double DEFAULT NULL,
  c13 double unsigned DEFAULT NULL COMMENT 'c13',
  c14 float DEFAULT NULL COMMENT 'c14',
  Column2 tinyint(1) DEFAULT NULL COMMENT 'Column2',
  Column3 TINYTEXT DEFAULT NULL COMMENT 'Column3',
  PRIMARY KEY (id)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci COMMENT 'a1';`

const examplePLSQL = `CREATE TABLE "ZT_0731"."MYTABLE" 
   (	"ID" NUMBER(19,0) NOT NULL ENABLE, 
	"NAME" VARCHAR2(50), 
	"BIRTHDAY" DATE, 
	"CREATED" DATE, 
	"IS_MANAGER" NUMBER(1,0), 
	"SALARY" NUMBER(9,2), 
	"WORKING_YEAR" NUMBER(10,0), 
	"TS" TIMESTAMP (6), 
	-- "COLUMN1" LONG, 
	-- "COLUMN2" INTERVAL DAY (2) TO SECOND (0), 
	-- "COLUMN3" INTERVAL YEAR (2) TO MONTH, 
	-- "COLUMN4" TIMESTAMP (6) WITH LOCAL TIME ZONE, 
	-- "COLUMN6" TIMESTAMP (6) WITH TIME ZONE, 
	 CONSTRAINT "MYTABLE_PK" PRIMARY KEY ("ID")
  USING INDEX PCTFREE 10 INITRANS 2 MAXTRANS 255 COMPUTE STATISTICS 
  STORAGE(INITIAL 65536 NEXT 1048576 MINEXTENTS 1 MAXEXTENTS 2147483645
  PCTINCREASE 0 FREELISTS 1 FREELIST GROUPS 1 BUFFER_POOL DEFAULT FLASH_CACHE DEFAULT CELL_FLASH_CACHE DEFAULT)
  TABLESPACE "SYSTEM"  ENABLE
   ) SEGMENT CREATION IMMEDIATE 
  PCTFREE 10 PCTUSED 40 INITRANS 1 MAXTRANS 255 NOCOMPRESS LOGGING
  STORAGE(INITIAL 65536 NEXT 1048576 MINEXTENTS 1 MAXEXTENTS 2147483645
  PCTINCREASE 0 FREELISTS 1 FREELIST GROUPS 1 BUFFER_POOL DEFAULT FLASH_CACHE DEFAULT CELL_FLASH_CACHE DEFAULT)
  TABLESPACE "SYSTEM";

CREATE UNIQUE INDEX "MYTABLE_PK" ON "MYTABLE" ("ID") 
  PCTFREE 10 INITRANS 2 MAXTRANS 255 COMPUTE STATISTICS 
  STORAGE(INITIAL 65536 NEXT 1048576 MINEXTENTS 1 MAXEXTENTS 2147483645
  PCTINCREASE 0 FREELISTS 1 FREELIST GROUPS 1 BUFFER_POOL DEFAULT FLASH_CACHE DEFAULT CELL_FLASH_CACHE DEFAULT)
  TABLESPACE "SYSTEM" ;

COMMENT ON COLUMN TEST.MYTABLE.NAME IS '姓名';
COMMENT ON COLUMN TEST.MYTABLE.BIRTHDAY IS '生日';
COMMENT ON COLUMN TEST.MYTABLE.CREATED IS '创建时间';
COMMENT ON COLUMN TEST.MYTABLE.IS_MANAGER IS '是否为管理者';
COMMENT ON COLUMN TEST.MYTABLE.SALARY IS '薪水';
COMMENT ON COLUMN TEST.MYTABLE.WORKING_YEAR IS '工作年限';
COMMENT ON COLUMN TEST.MYTABLE.TS IS '时间戳';

COMMENT ON TABLE TEST.MYTABLE IS 'Hello';
`

const examplePg = `CREATE TABLE public.mytable (
	id int8 NOT NULL,
	"name" varchar(50) NULL,
	birthday date NULL,
	created timestamp(0) NULL,
	is_manager bool NULL,
	salary numeric(9, 2) NULL,
	working_year int4 NULL,
	ts timestamp NULL,
	t time NULL, 
	CONSTRAINT newtable_pk PRIMARY KEY (id)
);

COMMENT ON COLUMN public.mytable.id IS 'id';
COMMENT ON COLUMN public.mytable."name" IS '姓名';
COMMENT ON COLUMN public.mytable.birthday IS '生日';
COMMENT ON COLUMN public.mytable.created IS '创建时间';
COMMENT ON COLUMN public.mytable.is_manager IS '是否为管理者';
COMMENT ON COLUMN public.mytable.salary IS '薪水';
COMMENT ON COLUMN public.mytable.working_year IS '工作年限';
COMMENT ON COLUMN public.mytable.ts IS '时间Ts';
COMMENT ON COLUMN public.mytable.t IS '时间T';

COMMENT ON TABLE public.mytable IS '娃哈哈';
`

const exampleSqlite = `CREATE TABLE NewTable (
	Column1 INTEGER,
	Column2 NUMERIC,
	Column3 REAL,
	Column4 TEXT
);`

const exampleTSQL = `CREATE TABLE test.dbo.mytable (
	id bigint NOT NULL,
	name varchar(50) COLLATE SQL_Latin1_General_CP1_CI_AS NULL,
	birthday date NULL,
	created datetime NULL,
	is_manager bit NULL,
	salary decimal(9,2) NULL,
	working_year int NULL,
	-- ts timestamp NULL,
	t time(0) NULL,
	Column1 date NULL,
	Column2 datetime2(0) NULL,
	Column3 money NULL,
	[date] date NULL,
	CONSTRAINT mytable_PK PRIMARY KEY (id)
);

EXEC test.sys.sp_updateextendedproperty 'MS_Description', N'name', 'schema', N'dbo', 'table', N'mytable', 'column', N'name';
EXEC test.sys.sp_addextendedproperty 'MS_Description', N'birthday', 'schema', N'dbo', 'table', N'mytable', 'column', N'birthday';
EXEC test.sys.sp_addextendedproperty 'MS_Description', N'created', 'schema', N'dbo', 'table', N'mytable', 'column', N'created';
EXEC test.sys.sp_addextendedproperty 'MS_Description', N'is_manager', 'schema', N'dbo', 'table', N'mytable', 'column', N'is_manager';
EXEC test.sys.sp_addextendedproperty 'MS_Description', N'salary', 'schema', N'dbo', 'table', N'mytable', 'column', N'salary';
EXEC test.sys.sp_addextendedproperty 'MS_Description', N'working_year', 'schema', N'dbo', 'table', N'mytable', 'column', N'working_year';
EXEC test.sys.sp_addextendedproperty 'MS_Description', N'ts', 'schema', N'dbo', 'table', N'mytable', 'column', N'ts';
EXEC test.sys.sp_addextendedproperty 'MS_Description', N't', 'schema', N'dbo', 'table', N'mytable', 'column', N't';
EXEC test.sys.sp_addextendedproperty 'MS_Description', N'Column1', 'schema', N'dbo', 'table', N'mytable', 'column', N'Column1';
EXEC test.sys.sp_addextendedproperty 'MS_Description', N'Column2', 'schema', N'dbo', 'table', N'mytable', 'column', N'Column2';

EXEC test.sys.sp_addextendedproperty 'MS_Description', N'mytable hahh', 'schema', N'dbo', 'table', N'mytable';
EXEC test.sys.sp_addextendedproperty @name=N'MS_Description', @value=N'这里是ID' , @level0type=N'SCHEMA',@level0name=N'dbo', @level1type=N'TABLE',@level1name=N'Student', @level2type=N'COLUMN',@level2name=N'id';
`
