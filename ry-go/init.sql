
-- ----------------------------
-- Sequence structure for gen_table_column_column_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."gen_table_column_column_id_seq";
CREATE SEQUENCE "public"."gen_table_column_column_id_seq"
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for gen_table_table_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."gen_table_table_id_seq";
CREATE SEQUENCE "public"."gen_table_table_id_seq"
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for sys_config_config_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."sys_config_config_id_seq";
CREATE SEQUENCE "public"."sys_config_config_id_seq"
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 5
CACHE 1;

-- ----------------------------
-- Sequence structure for sys_dept_dept_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."sys_dept_dept_id_seq";
CREATE SEQUENCE "public"."sys_dept_dept_id_seq"
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;


-- ----------------------------
-- Sequence structure for sys_dict_data_dict_code_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."sys_dict_data_dict_code_seq";
CREATE SEQUENCE "public"."sys_dict_data_dict_code_seq"
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 29
CACHE 1;

-- ----------------------------
-- Sequence structure for sys_dict_type_dict_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."sys_dict_type_dict_id_seq";
CREATE SEQUENCE "public"."sys_dict_type_dict_id_seq"
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 14
CACHE 1;

-- ----------------------------
-- Sequence structure for sys_job_job_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."sys_job_job_id_seq";
CREATE SEQUENCE "public"."sys_job_job_id_seq"
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 5
CACHE 1;

-- ----------------------------
-- Sequence structure for sys_job_log_job_log_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."sys_job_log_job_log_id_seq";
CREATE SEQUENCE "public"."sys_job_log_job_log_id_seq"
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for sys_login_record_info_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."sys_login_record_info_id_seq";
CREATE SEQUENCE "public"."sys_login_record_info_id_seq"
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for sys_menu_menu_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."sys_menu_menu_id_seq";
CREATE SEQUENCE "public"."sys_menu_menu_id_seq"
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 20
CACHE 1;

-- ----------------------------
-- Sequence structure for sys_notice_notice_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."sys_notice_notice_id_seq";
CREATE SEQUENCE "public"."sys_notice_notice_id_seq"
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for sys_operate_log_oper_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."sys_operate_log_oper_id_seq";
CREATE SEQUENCE "public"."sys_operate_log_oper_id_seq"
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Sequence structure for sys_post_post_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."sys_post_post_id_seq";
CREATE SEQUENCE "public"."sys_post_post_id_seq"
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 14
CACHE 1;

-- ----------------------------
-- Sequence structure for sys_role_role_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."sys_role_role_id_seq";
CREATE SEQUENCE "public"."sys_role_role_id_seq"
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 7
CACHE 1;

-- ----------------------------
-- Sequence structure for sys_user_user_id_seq
-- ----------------------------
DROP SEQUENCE IF EXISTS "public"."sys_user_user_id_seq";
CREATE SEQUENCE "public"."sys_user_user_id_seq"
INCREMENT 1
MINVALUE  1
MAXVALUE 9223372036854775807
START 1
CACHE 1;

-- ----------------------------
-- Table structure for QRTZ_BLOB_TRIGGERS
-- ----------------------------
DROP TABLE IF EXISTS "public"."QRTZ_BLOB_TRIGGERS";
CREATE TABLE "public"."QRTZ_BLOB_TRIGGERS" (
                                               "sched_name" varchar(120) COLLATE "pg_catalog"."default" NOT NULL,
                                               "trigger_name" varchar(200) COLLATE "pg_catalog"."default" NOT NULL,
                                               "trigger_group" varchar(200) COLLATE "pg_catalog"."default" NOT NULL,
                                               "blob_data" bytea
)
;

-- ----------------------------
-- Records of QRTZ_BLOB_TRIGGERS
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for QRTZ_CALENDARS
-- ----------------------------
DROP TABLE IF EXISTS "public"."QRTZ_CALENDARS";
CREATE TABLE "public"."QRTZ_CALENDARS" (
                                           "sched_name" varchar(120) COLLATE "pg_catalog"."default" NOT NULL,
                                           "calendar_name" varchar(200) COLLATE "pg_catalog"."default" NOT NULL,
                                           "calendar" bytea NOT NULL
)
;
-----------------------------
-- Records of QRTZ_CALENDARS
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for QRTZ_CRON_TRIGGERS
-- ----------------------------
DROP TABLE IF EXISTS "public"."QRTZ_CRON_TRIGGERS";
CREATE TABLE "public"."QRTZ_CRON_TRIGGERS" (
                                               "sched_name" varchar(120) COLLATE "pg_catalog"."default" NOT NULL,
                                               "trigger_name" varchar(200) COLLATE "pg_catalog"."default" NOT NULL,
                                               "trigger_group" varchar(200) COLLATE "pg_catalog"."default" NOT NULL,
                                               "cron_expression" varchar(200) COLLATE "pg_catalog"."default" NOT NULL,
                                               "time_zone_id" varchar(80) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."QRTZ_CRON_TRIGGERS" OWNER TO "yangge";

-- ----------------------------
-- Records of QRTZ_CRON_TRIGGERS
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for QRTZ_FIRED_TRIGGERS
-- ----------------------------
DROP TABLE IF EXISTS "public"."QRTZ_FIRED_TRIGGERS";
CREATE TABLE "public"."QRTZ_FIRED_TRIGGERS" (
                                                "sched_name" varchar(120) COLLATE "pg_catalog"."default" NOT NULL,
                                                "entry_id" varchar(95) COLLATE "pg_catalog"."default" NOT NULL,
                                                "trigger_name" varchar(200) COLLATE "pg_catalog"."default" NOT NULL,
                                                "trigger_group" varchar(200) COLLATE "pg_catalog"."default" NOT NULL,
                                                "instance_name" varchar(200) COLLATE "pg_catalog"."default" NOT NULL,
                                                "fired_time" int8 NOT NULL,
                                                "sched_time" int8 NOT NULL,
                                                "priority" int4 NOT NULL,
                                                "state" varchar(16) COLLATE "pg_catalog"."default" NOT NULL,
                                                "job_name" varchar(200) COLLATE "pg_catalog"."default",
                                                "job_group" varchar(200) COLLATE "pg_catalog"."default",
                                                "is_nonconcurrent" varchar(1) COLLATE "pg_catalog"."default",
                                                "requests_recovery" varchar(1) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."QRTZ_FIRED_TRIGGERS" OWNER TO "yangge";

-- ----------------------------
-- Records of QRTZ_FIRED_TRIGGERS
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for QRTZ_JOB_DETAILS
-- ----------------------------
DROP TABLE IF EXISTS "public"."QRTZ_JOB_DETAILS";
CREATE TABLE "public"."QRTZ_JOB_DETAILS" (
                                             "sched_name" varchar(120) COLLATE "pg_catalog"."default" NOT NULL,
                                             "job_name" varchar(200) COLLATE "pg_catalog"."default" NOT NULL,
                                             "job_group" varchar(200) COLLATE "pg_catalog"."default" NOT NULL,
                                             "description" varchar(250) COLLATE "pg_catalog"."default",
                                             "job_class_name" varchar(250) COLLATE "pg_catalog"."default" NOT NULL,
                                             "is_durable" varchar(1) COLLATE "pg_catalog"."default" NOT NULL,
                                             "is_nonconcurrent" varchar(1) COLLATE "pg_catalog"."default" NOT NULL,
                                             "is_update_data" varchar(1) COLLATE "pg_catalog"."default" NOT NULL,
                                             "requests_recovery" varchar(1) COLLATE "pg_catalog"."default" NOT NULL,
                                             "job_data" bytea
)
;
ALTER TABLE "public"."QRTZ_JOB_DETAILS" OWNER TO "yangge";

-- ----------------------------
-- Records of QRTZ_JOB_DETAILS
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for QRTZ_LOCKS
-- ----------------------------
DROP TABLE IF EXISTS "public"."QRTZ_LOCKS";
CREATE TABLE "public"."QRTZ_LOCKS" (
                                       "sched_name" varchar(120) COLLATE "pg_catalog"."default" NOT NULL,
                                       "lock_name" varchar(40) COLLATE "pg_catalog"."default" NOT NULL
)
;
ALTER TABLE "public"."QRTZ_LOCKS" OWNER TO "yangge";

-- ----------------------------
-- Records of QRTZ_LOCKS
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for QRTZ_PAUSED_TRIGGER_GRPS
-- ----------------------------
DROP TABLE IF EXISTS "public"."QRTZ_PAUSED_TRIGGER_GRPS";
CREATE TABLE "public"."QRTZ_PAUSED_TRIGGER_GRPS" (
                                                     "sched_name" varchar(120) COLLATE "pg_catalog"."default" NOT NULL,
                                                     "trigger_group" varchar(200) COLLATE "pg_catalog"."default" NOT NULL
)
;
ALTER TABLE "public"."QRTZ_PAUSED_TRIGGER_GRPS" OWNER TO "yangge";

-- ----------------------------
-- Records of QRTZ_PAUSED_TRIGGER_GRPS
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for QRTZ_SCHEDULER_STATE
-- ----------------------------
DROP TABLE IF EXISTS "public"."QRTZ_SCHEDULER_STATE";
CREATE TABLE "public"."QRTZ_SCHEDULER_STATE" (
                                                 "sched_name" varchar(120) COLLATE "pg_catalog"."default" NOT NULL,
                                                 "instance_name" varchar(200) COLLATE "pg_catalog"."default" NOT NULL,
                                                 "last_checkin_time" int8 NOT NULL,
                                                 "checkin_interval" int8 NOT NULL
)
;
ALTER TABLE "public"."QRTZ_SCHEDULER_STATE" OWNER TO "yangge";

-- ----------------------------
-- Records of QRTZ_SCHEDULER_STATE
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for QRTZ_SIMPLE_TRIGGERS
-- ----------------------------
DROP TABLE IF EXISTS "public"."QRTZ_SIMPLE_TRIGGERS";
CREATE TABLE "public"."QRTZ_SIMPLE_TRIGGERS" (
                                                 "sched_name" varchar(120) COLLATE "pg_catalog"."default" NOT NULL,
                                                 "trigger_name" varchar(200) COLLATE "pg_catalog"."default" NOT NULL,
                                                 "trigger_group" varchar(200) COLLATE "pg_catalog"."default" NOT NULL,
                                                 "repeat_count" int8 NOT NULL,
                                                 "repeat_interval" int8 NOT NULL,
                                                 "times_triggered" int8 NOT NULL
)
;
ALTER TABLE "public"."QRTZ_SIMPLE_TRIGGERS" OWNER TO "yangge";

-- ----------------------------
-- Records of QRTZ_SIMPLE_TRIGGERS
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for QRTZ_SIMPROP_TRIGGERS
-- ----------------------------
DROP TABLE IF EXISTS "public"."QRTZ_SIMPROP_TRIGGERS";
CREATE TABLE "public"."QRTZ_SIMPROP_TRIGGERS" (
                                                  "sched_name" varchar(120) COLLATE "pg_catalog"."default" NOT NULL,
                                                  "trigger_name" varchar(200) COLLATE "pg_catalog"."default" NOT NULL,
                                                  "trigger_group" varchar(200) COLLATE "pg_catalog"."default" NOT NULL,
                                                  "str_prop_1" varchar(512) COLLATE "pg_catalog"."default",
                                                  "str_prop_2" varchar(512) COLLATE "pg_catalog"."default",
                                                  "str_prop_3" varchar(512) COLLATE "pg_catalog"."default",
                                                  "int_prop_1" int4,
                                                  "int_prop_2" int4,
                                                  "long_prop_1" int8,
                                                  "long_prop_2" int8,
                                                  "dec_prop_1" numeric(13,4),
                                                  "dec_prop_2" numeric(13,4),
                                                  "bool_prop_1" varchar(1) COLLATE "pg_catalog"."default",
                                                  "bool_prop_2" varchar(1) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."QRTZ_SIMPROP_TRIGGERS" OWNER TO "yangge";

-- ----------------------------
-- Records of QRTZ_SIMPROP_TRIGGERS
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for QRTZ_TRIGGERS
-- ----------------------------
DROP TABLE IF EXISTS "public"."QRTZ_TRIGGERS";
CREATE TABLE "public"."QRTZ_TRIGGERS" (
                                          "sched_name" varchar(120) COLLATE "pg_catalog"."default" NOT NULL,
                                          "trigger_name" varchar(200) COLLATE "pg_catalog"."default" NOT NULL,
                                          "trigger_group" varchar(200) COLLATE "pg_catalog"."default" NOT NULL,
                                          "job_name" varchar(200) COLLATE "pg_catalog"."default" NOT NULL,
                                          "job_group" varchar(200) COLLATE "pg_catalog"."default" NOT NULL,
                                          "description" varchar(250) COLLATE "pg_catalog"."default",
                                          "next_fire_time" int8,
                                          "prev_fire_time" int8,
                                          "priority" int4,
                                          "trigger_state" varchar(16) COLLATE "pg_catalog"."default" NOT NULL,
                                          "trigger_type" varchar(8) COLLATE "pg_catalog"."default" NOT NULL,
                                          "start_time" int8 NOT NULL,
                                          "end_time" int8,
                                          "calendar_name" varchar(200) COLLATE "pg_catalog"."default",
                                          "misfire_instr" int2,
                                          "job_data" bytea
)
;
ALTER TABLE "public"."QRTZ_TRIGGERS" OWNER TO "yangge";

-- ----------------------------
-- Records of QRTZ_TRIGGERS
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for gen_table
-- ----------------------------
DROP TABLE IF EXISTS "public"."gen_table";
CREATE TABLE "public"."gen_table" (
                                      "table_id" int8 NOT NULL GENERATED BY DEFAULT AS IDENTITY (
                                          INCREMENT 1
                                          MINVALUE  1
                                          MAXVALUE 9223372036854775807
                                          START 1
                                          CACHE 1
                                          ),
                                      "table_name" varchar(200) COLLATE "pg_catalog"."default",
                                      "table_comment" varchar(500) COLLATE "pg_catalog"."default",
                                      "sub_table_name" varchar(64) COLLATE "pg_catalog"."default",
                                      "sub_table_fk_name" varchar(64) COLLATE "pg_catalog"."default",
                                      "class_name" varchar(100) COLLATE "pg_catalog"."default",
                                      "tpl_category" varchar(200) COLLATE "pg_catalog"."default",
                                      "package_name" varchar(100) COLLATE "pg_catalog"."default",
                                      "module_name" varchar(30) COLLATE "pg_catalog"."default",
                                      "business_name" varchar(30) COLLATE "pg_catalog"."default",
                                      "function_name" varchar(50) COLLATE "pg_catalog"."default",
                                      "function_author" varchar(50) COLLATE "pg_catalog"."default",
                                      "gen_type" char(1) COLLATE "pg_catalog"."default",
                                      "gen_path" varchar(200) COLLATE "pg_catalog"."default",
                                      "options" varchar(1000) COLLATE "pg_catalog"."default",
                                      "create_by" varchar(64) COLLATE "pg_catalog"."default",
                                      "create_time" timestamp(6),
                                      "update_by" varchar(64) COLLATE "pg_catalog"."default",
                                      "update_time" timestamp(6),
                                      "remark" varchar(500) COLLATE "pg_catalog"."default",
                                      "swagger" int4,
                                      "excel_export" int4,
                                      "vue_version" int4
)
;
ALTER TABLE "public"."gen_table" OWNER TO "yangge";
COMMENT ON COLUMN "public"."gen_table"."table_id" IS '编号';
COMMENT ON COLUMN "public"."gen_table"."table_name" IS '表名称';
COMMENT ON COLUMN "public"."gen_table"."table_comment" IS '表描述';
COMMENT ON COLUMN "public"."gen_table"."sub_table_name" IS '关联子表的表名';
COMMENT ON COLUMN "public"."gen_table"."sub_table_fk_name" IS '子表关联的外键名';
COMMENT ON COLUMN "public"."gen_table"."class_name" IS '实体类名称';
COMMENT ON COLUMN "public"."gen_table"."tpl_category" IS '使用的模板（crud单表操作 tree树表操作）';
COMMENT ON COLUMN "public"."gen_table"."package_name" IS '生成包路径';
COMMENT ON COLUMN "public"."gen_table"."module_name" IS '生成模块名';
COMMENT ON COLUMN "public"."gen_table"."business_name" IS '生成业务名';
COMMENT ON COLUMN "public"."gen_table"."function_name" IS '生成功能名';
COMMENT ON COLUMN "public"."gen_table"."function_author" IS '生成功能作者';
COMMENT ON COLUMN "public"."gen_table"."gen_type" IS '生成代码方式（0zip压缩包 1自定义路径）';
COMMENT ON COLUMN "public"."gen_table"."gen_path" IS '生成路径（不填默认项目路径）';
COMMENT ON COLUMN "public"."gen_table"."options" IS '其它生成选项';
COMMENT ON COLUMN "public"."gen_table"."create_by" IS '创建者';
COMMENT ON COLUMN "public"."gen_table"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."gen_table"."update_by" IS '更新者';
COMMENT ON COLUMN "public"."gen_table"."update_time" IS '更新时间';
COMMENT ON COLUMN "public"."gen_table"."remark" IS '备注';
COMMENT ON COLUMN "public"."gen_table"."swagger" IS '是否生成swagger注释(0:否，1：是)';
COMMENT ON COLUMN "public"."gen_table"."excel_export" IS '是否生成导出excel注释(0:否，1：是)';
COMMENT ON COLUMN "public"."gen_table"."vue_version" IS 'vue版本(2或者3，默认是2)';
COMMENT ON TABLE "public"."gen_table" IS '代码生成业务表';

-- ----------------------------
-- Records of gen_table
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for gen_table_column
-- ----------------------------
DROP TABLE IF EXISTS "public"."gen_table_column";
CREATE TABLE "public"."gen_table_column" (
                                             "column_id" int8 NOT NULL GENERATED BY DEFAULT AS IDENTITY (
                                                 INCREMENT 1
                                                 MINVALUE  1
                                                 MAXVALUE 9223372036854775807
                                                 START 1
                                                 CACHE 1
                                                 ),
                                             "table_id" varchar(64) COLLATE "pg_catalog"."default",
                                             "column_name" varchar(200) COLLATE "pg_catalog"."default",
                                             "column_comment" varchar(500) COLLATE "pg_catalog"."default",
                                             "column_type" varchar(100) COLLATE "pg_catalog"."default",
                                             "go_type" varchar(500) COLLATE "pg_catalog"."default",
                                             "go_field" varchar(200) COLLATE "pg_catalog"."default",
                                             "is_pk" char(1) COLLATE "pg_catalog"."default",
                                             "is_increment" char(1) COLLATE "pg_catalog"."default",
                                             "is_required" char(1) COLLATE "pg_catalog"."default",
                                             "is_insert" char(1) COLLATE "pg_catalog"."default",
                                             "is_edit" char(1) COLLATE "pg_catalog"."default",
                                             "is_list" char(1) COLLATE "pg_catalog"."default",
                                             "is_query" char(1) COLLATE "pg_catalog"."default",
                                             "query_type" varchar(200) COLLATE "pg_catalog"."default",
                                             "html_type" varchar(200) COLLATE "pg_catalog"."default",
                                             "dict_type" varchar(200) COLLATE "pg_catalog"."default",
                                             "sort" int4,
                                             "create_by" varchar(64) COLLATE "pg_catalog"."default",
                                             "create_time" timestamp(6),
                                             "update_by" varchar(64) COLLATE "pg_catalog"."default",
                                             "update_time" timestamp(6)
)
;
ALTER TABLE "public"."gen_table_column" OWNER TO "yangge";
COMMENT ON COLUMN "public"."gen_table_column"."column_id" IS '编号';
COMMENT ON COLUMN "public"."gen_table_column"."table_id" IS '归属表编号';
COMMENT ON COLUMN "public"."gen_table_column"."column_name" IS '列名称';
COMMENT ON COLUMN "public"."gen_table_column"."column_comment" IS '列描述';
COMMENT ON COLUMN "public"."gen_table_column"."column_type" IS '列类型';
COMMENT ON COLUMN "public"."gen_table_column"."go_type" IS 'golang类型';
COMMENT ON COLUMN "public"."gen_table_column"."go_field" IS 'golang字段名';
COMMENT ON COLUMN "public"."gen_table_column"."is_pk" IS '是否主键（1是）';
COMMENT ON COLUMN "public"."gen_table_column"."is_increment" IS '是否自增（1是）';
COMMENT ON COLUMN "public"."gen_table_column"."is_required" IS '是否必填（1是）';
COMMENT ON COLUMN "public"."gen_table_column"."is_insert" IS '是否为插入字段（1是）';
COMMENT ON COLUMN "public"."gen_table_column"."is_edit" IS '是否编辑字段（1是）';
COMMENT ON COLUMN "public"."gen_table_column"."is_list" IS '是否列表字段（1是）';
COMMENT ON COLUMN "public"."gen_table_column"."is_query" IS '是否查询字段（1是）';
COMMENT ON COLUMN "public"."gen_table_column"."query_type" IS '查询方式（等于、不等于、大于、小于、范围）';
COMMENT ON COLUMN "public"."gen_table_column"."html_type" IS '显示类型（文本框、文本域、下拉框、复选框、单选框、日期控件）';
COMMENT ON COLUMN "public"."gen_table_column"."dict_type" IS '字典类型';
COMMENT ON COLUMN "public"."gen_table_column"."sort" IS '排序';
COMMENT ON COLUMN "public"."gen_table_column"."create_by" IS '创建者';
COMMENT ON COLUMN "public"."gen_table_column"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."gen_table_column"."update_by" IS '更新者';
COMMENT ON COLUMN "public"."gen_table_column"."update_time" IS '更新时间';
COMMENT ON TABLE "public"."gen_table_column" IS '代码生成业务表字段';

-- ----------------------------
-- Records of gen_table_column
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_config
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_config";
CREATE TABLE "public"."sys_config" (
                                       "config_id" int8 NOT NULL GENERATED BY DEFAULT AS IDENTITY (
                                           INCREMENT 1
                                           MINVALUE  1
                                           MAXVALUE 9223372036854775807
                                           START 5
                                           CACHE 1
                                           ),
                                       "config_name" varchar(100) COLLATE "pg_catalog"."default",
                                       "config_key" varchar(100) COLLATE "pg_catalog"."default",
                                       "config_value" varchar(500) COLLATE "pg_catalog"."default",
                                       "config_type" char(1) COLLATE "pg_catalog"."default",
                                       "create_by" varchar(64) COLLATE "pg_catalog"."default",
                                       "create_time" timestamp(6),
                                       "update_by" varchar(64) COLLATE "pg_catalog"."default",
                                       "update_time" timestamp(6),
                                       "remark" varchar(500) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."sys_config" OWNER TO "yangge";
COMMENT ON COLUMN "public"."sys_config"."config_id" IS '参数主键';
COMMENT ON COLUMN "public"."sys_config"."config_name" IS '参数名称';
COMMENT ON COLUMN "public"."sys_config"."config_key" IS '参数键名';
COMMENT ON COLUMN "public"."sys_config"."config_value" IS '参数键值';
COMMENT ON COLUMN "public"."sys_config"."config_type" IS '系统内置（Y是 N否）';
COMMENT ON COLUMN "public"."sys_config"."create_by" IS '创建者';
COMMENT ON COLUMN "public"."sys_config"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."sys_config"."update_by" IS '更新者';
COMMENT ON COLUMN "public"."sys_config"."update_time" IS '更新时间';
COMMENT ON COLUMN "public"."sys_config"."remark" IS '备注';
COMMENT ON TABLE "public"."sys_config" IS '参数配置表';

-- ----------------------------
-- Records of sys_config
-- ----------------------------
BEGIN;
INSERT INTO "public"."sys_config" ("config_id", "config_name", "config_key", "config_value", "config_type", "create_by", "create_time", "update_by", "update_time", "remark") VALUES (2, '用户管理-账号初始密码', 'sys.user.initPassword', '123456', 'Y', 'admin', '2022-05-20 02:27:32', 'admin', '2022-05-20 02:27:32', '初始化密码 123456');
INSERT INTO "public"."sys_config" ("config_id", "config_name", "config_key", "config_value", "config_type", "create_by", "create_time", "update_by", "update_time", "remark") VALUES (3, '主框架页-侧边栏主题', 'sys.index.sideTheme', 'theme-dark', 'Y', 'admin', '2022-05-20 02:27:32', 'admin', '2022-05-20 02:27:32', '深色主题theme-dark，浅色主题theme-light');
INSERT INTO "public"."sys_config" ("config_id", "config_name", "config_key", "config_value", "config_type", "create_by", "create_time", "update_by", "update_time", "remark") VALUES (1, '主框架页-默认皮肤样式名称', 'sys.index.skinName', 'skin-blue', 'Y', 'admin', '2022-05-20 02:27:32', 'admin', '2022-05-23 14:37:08.816799', '蓝色 skin-blue、绿色 skin-green、紫色 skin-purple、红色 skin-red、黄色 skin-yellow');
INSERT INTO "public"."sys_config" ("config_id", "config_name", "config_key", "config_value", "config_type", "create_by", "create_time", "update_by", "update_time", "remark") VALUES (8, '33', '333', '3333', 'N', '', '2022-04-17 20:20:31', '', '2022-04-17 20:23:23.648722', '44444');
INSERT INTO "public"."sys_config" ("config_id", "config_name", "config_key", "config_value", "config_type", "create_by", "create_time", "update_by", "update_time", "remark") VALUES (4, '修改', '修改', '13434343', 'N', '大萨达大厦', NULL, '章三', '2022-04-17 20:36:01.310517', '这是备注');
COMMIT;

-- ----------------------------
-- Table structure for sys_dept
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_dept";
CREATE TABLE "public"."sys_dept" (
                                     "id" int8 NOT NULL GENERATED BY DEFAULT AS IDENTITY (
                                         INCREMENT 1
                                         MINVALUE  1
                                         MAXVALUE 9223372036854775807
                                         START 1
                                         CACHE 1
                                         ),
                                     "parent_id" int8,
                                     "ancestors" varchar(50) COLLATE "pg_catalog"."default",
                                     "dept_name" varchar(30) COLLATE "pg_catalog"."default",
                                     "order_num" int4,
                                     "dept_status" char(1) COLLATE "pg_catalog"."default" DEFAULT 0,
                                     "create_by" varchar(64) COLLATE "pg_catalog"."default",
                                     "create_time" timestamp(6),
                                     "update_by" varchar(64) COLLATE "pg_catalog"."default",
                                     "update_time" timestamp(6),
                                     "remarks" varchar(100) COLLATE "pg_catalog"."default",
                                     "del_flag" char(1) COLLATE "pg_catalog"."default" DEFAULT 0,
                                     "leader" varchar(50) COLLATE "pg_catalog"."default",
                                     "phone" varchar(12) COLLATE "pg_catalog"."default",
                                     "email" varchar(100) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."sys_dept" OWNER TO "yangge";
COMMENT ON COLUMN "public"."sys_dept"."id" IS '部门ID';
COMMENT ON COLUMN "public"."sys_dept"."parent_id" IS '父部门ID';
COMMENT ON COLUMN "public"."sys_dept"."ancestors" IS '祖级列表';
COMMENT ON COLUMN "public"."sys_dept"."dept_name" IS '部门名称';
COMMENT ON COLUMN "public"."sys_dept"."order_num" IS '排序';
COMMENT ON COLUMN "public"."sys_dept"."dept_status" IS '部门状态（0正常 1停用）';
COMMENT ON COLUMN "public"."sys_dept"."create_by" IS '创建者';
COMMENT ON COLUMN "public"."sys_dept"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."sys_dept"."update_by" IS '更新者';
COMMENT ON COLUMN "public"."sys_dept"."update_time" IS '更新时间';
COMMENT ON COLUMN "public"."sys_dept"."remarks" IS '备注';
COMMENT ON COLUMN "public"."sys_dept"."del_flag" IS '删除标识（0:正常；1:删除）';
COMMENT ON COLUMN "public"."sys_dept"."leader" IS '负责人';
COMMENT ON COLUMN "public"."sys_dept"."phone" IS '联系电话';
COMMENT ON COLUMN "public"."sys_dept"."email" IS '邮箱';
COMMENT ON TABLE "public"."sys_dept" IS '部门表';

-- ----------------------------
-- Records of sys_dept
-- ----------------------------
BEGIN;
INSERT INTO "public"."sys_dept" ("id", "parent_id", "ancestors", "dept_name", "order_num", "dept_status", "create_by", "create_time", "update_by", "update_time", "remarks", "del_flag", "leader", "phone", "email") VALUES (101, 100, '0,100', '深圳总公司', 1, '0', 'admin', '2021-03-07 17:58:10', '', NULL, NULL, '0', '若依', '15888888888', 'ry@qq.com');
INSERT INTO "public"."sys_dept" ("id", "parent_id", "ancestors", "dept_name", "order_num", "dept_status", "create_by", "create_time", "update_by", "update_time", "remarks", "del_flag", "leader", "phone", "email") VALUES (102, 100, '0,100', '长沙分公司', 2, '0', 'admin', '2021-03-07 17:58:10', 'admin', '2022-08-01 08:08:20', NULL, '0', '若依', '15588888888', 'ry@qq.com');
INSERT INTO "public"."sys_dept" ("id", "parent_id", "ancestors", "dept_name", "order_num", "dept_status", "create_by", "create_time", "update_by", "update_time", "remarks", "del_flag", "leader", "phone", "email") VALUES (103, 101, '0,100,101', '研发部门', 1, '0', 'admin', '2021-03-07 17:58:10', '', NULL, NULL, '0', '若依', '15888888888', 'ry@qq.com');
INSERT INTO "public"."sys_dept" ("id", "parent_id", "ancestors", "dept_name", "order_num", "dept_status", "create_by", "create_time", "update_by", "update_time", "remarks", "del_flag", "leader", "phone", "email") VALUES (104, 101, '0,100,101', '市场部门', 2, '0', 'admin', '2021-03-07 17:58:10', '', NULL, NULL, '0', '若依', '15888888888', 'ry@qq.com');
INSERT INTO "public"."sys_dept" ("id", "parent_id", "ancestors", "dept_name", "order_num", "dept_status", "create_by", "create_time", "update_by", "update_time", "remarks", "del_flag", "leader", "phone", "email") VALUES (105, 101, '0,100,101', '测试部门', 3, '0', 'admin', '2021-03-07 17:58:10', '', NULL, NULL, '0', '若依', '15888888888', 'ry@qq.com');
INSERT INTO "public"."sys_dept" ("id", "parent_id", "ancestors", "dept_name", "order_num", "dept_status", "create_by", "create_time", "update_by", "update_time", "remarks", "del_flag", "leader", "phone", "email") VALUES (106, 101, '0,100,101', '财务部门', 4, '0', 'admin', '2021-03-07 17:58:10', '', NULL, NULL, '0', '若依', '15888888888', 'ry@qq.com');
INSERT INTO "public"."sys_dept" ("id", "parent_id", "ancestors", "dept_name", "order_num", "dept_status", "create_by", "create_time", "update_by", "update_time", "remarks", "del_flag", "leader", "phone", "email") VALUES (107, 101, '0,100,101', '运维部门', 5, '0', 'admin', '2021-03-07 17:58:10', '', NULL, NULL, '0', '若依', '15888888888', 'ry@qq.com');
INSERT INTO "public"."sys_dept" ("id", "parent_id", "ancestors", "dept_name", "order_num", "dept_status", "create_by", "create_time", "update_by", "update_time", "remarks", "del_flag", "leader", "phone", "email") VALUES (108, 102, '0,100,102', '市场部门', 1, '0', 'admin', '2021-03-07 17:58:10', '', NULL, NULL, '0', '若依', '15888888888', 'ry@qq.com');
INSERT INTO "public"."sys_dept" ("id", "parent_id", "ancestors", "dept_name", "order_num", "dept_status", "create_by", "create_time", "update_by", "update_time", "remarks", "del_flag", "leader", "phone", "email") VALUES (109, 102, '0,100,102', '财务部门', 2, '0', 'admin', '2021-03-07 17:58:10', '', NULL, NULL, '0', '若依', '15888888888', 'ry@qq.com');
INSERT INTO "public"."sys_dept" ("id", "parent_id", "ancestors", "dept_name", "order_num", "dept_status", "create_by", "create_time", "update_by", "update_time", "remarks", "del_flag", "leader", "phone", "email") VALUES (111, 100, '0,100', '杭州分公司', 33, '0', 'test', '2022-07-16 04:52:46', 'admin', '2022-07-20 09:52:36', NULL, '0', NULL, NULL, NULL);
INSERT INTO "public"."sys_dept" ("id", "parent_id", "ancestors", "dept_name", "order_num", "dept_status", "create_by", "create_time", "update_by", "update_time", "remarks", "del_flag", "leader", "phone", "email") VALUES (119, 111, '0,100,111', '研发一部', 1, '0', 'admin', '2022-07-20 09:52:50', '', NULL, NULL, '0', NULL, NULL, NULL);
INSERT INTO "public"."sys_dept" ("id", "parent_id", "ancestors", "dept_name", "order_num", "dept_status", "create_by", "create_time", "update_by", "update_time", "remarks", "del_flag", "leader", "phone", "email") VALUES (120, 111, '0,100,111', '研发二部', 2, '0', 'admin', '2022-07-20 09:52:59', '', NULL, NULL, '0', NULL, NULL, NULL);
INSERT INTO "public"."sys_dept" ("id", "parent_id", "ancestors", "dept_name", "order_num", "dept_status", "create_by", "create_time", "update_by", "update_time", "remarks", "del_flag", "leader", "phone", "email") VALUES (121, 111, '0,100,111', '测试部门', 3, '0', 'admin', '2022-07-20 09:53:11', '', NULL, NULL, '0', NULL, NULL, NULL);
INSERT INTO "public"."sys_dept" ("id", "parent_id", "ancestors", "dept_name", "order_num", "dept_status", "create_by", "create_time", "update_by", "update_time", "remarks", "del_flag", "leader", "phone", "email") VALUES (122, 111, '0,100,111', '服务运维', 4, '0', 'admin', '2022-07-20 09:53:29', '', NULL, NULL, '0', NULL, NULL, NULL);
INSERT INTO "public"."sys_dept" ("id", "parent_id", "ancestors", "dept_name", "order_num", "dept_status", "create_by", "create_time", "update_by", "update_time", "remarks", "del_flag", "leader", "phone", "email") VALUES (123, 111, '0,100,111', '售后维护', 6, '1', 'admin', '2022-08-01 03:33:01', 'admin', '2022-08-01 03:33:51', NULL, '0', '大大大啊', '15684848888', '111@sss.com');
INSERT INTO "public"."sys_dept" ("id", "parent_id", "ancestors", "dept_name", "order_num", "dept_status", "create_by", "create_time", "update_by", "update_time", "remarks", "del_flag", "leader", "phone", "email") VALUES (124, 100, '0,100', '大大啊', 200, '1', 'admin', '2022-08-05 08:20:12', 'admin', '2022-08-05 08:21:13', NULL, '2', '奥术大师S', NULL, '111@163.xom');
INSERT INTO "public"."sys_dept" ("id", "parent_id", "ancestors", "dept_name", "order_num", "dept_status", "create_by", "create_time", "update_by", "update_time", "remarks", "del_flag", "leader", "phone", "email") VALUES (125, 100, '0,100', '大撒啊啊', 201, '0', 'admin', '2022-08-05 08:20:28', '', NULL, NULL, '2', NULL, NULL, NULL);
INSERT INTO "public"."sys_dept" ("id", "parent_id", "ancestors", "dept_name", "order_num", "dept_status", "create_by", "create_time", "update_by", "update_time", "remarks", "del_flag", "leader", "phone", "email") VALUES (126, 124, '0,100,124', '大大撒啊', 1111, '1', 'admin', '2022-08-05 08:25:50', 'admin', '2022-08-05 09:27:38', NULL, '2', NULL, NULL, NULL);
INSERT INTO "public"."sys_dept" ("id", "parent_id", "ancestors", "dept_name", "order_num", "dept_status", "create_by", "create_time", "update_by", "update_time", "remarks", "del_flag", "leader", "phone", "email") VALUES (127, 104, '0,100,101,104', '大大啊', 2222, '0', 'admin', '2022-08-05 09:20:05', '', NULL, NULL, '0', '大大啊', NULL, NULL);
INSERT INTO "public"."sys_dept" ("id", "parent_id", "ancestors", "dept_name", "order_num", "dept_status", "create_by", "create_time", "update_by", "update_time", "remarks", "del_flag", "leader", "phone", "email") VALUES (128, 100, '0,100', 'dsadaaa', 11111, '0', 'admin', '2022-08-05 09:44:21', 'admin', '2022-08-05 09:44:26', NULL, '2', 'sdsdsdsasa', NULL, NULL);
INSERT INTO "public"."sys_dept" ("id", "parent_id", "ancestors", "dept_name", "order_num", "dept_status", "create_by", "create_time", "update_by", "update_time", "remarks", "del_flag", "leader", "phone", "email") VALUES (129, 128, '0,100,128', 'dsadaa', 1, '0', 'admin', '2022-08-05 09:44:33', '', NULL, NULL, '2', 'dsadada', NULL, NULL);
INSERT INTO "public"."sys_dept" ("id", "parent_id", "ancestors", "dept_name", "order_num", "dept_status", "create_by", "create_time", "update_by", "update_time", "remarks", "del_flag", "leader", "phone", "email") VALUES (130, 105, '0,100,101,105', '测试试是', 10, '0', 'admin', '2022-08-06 10:01:04', '', NULL, NULL, '0', '大大啊啊啊', NULL, NULL);
INSERT INTO "public"."sys_dept" ("id", "parent_id", "ancestors", "dept_name", "order_num", "dept_status", "create_by", "create_time", "update_by", "update_time", "remarks", "del_flag", "leader", "phone", "email") VALUES (131, 130, '0,100,101,105,130', '大萨达啊大', 1, '1', 'admin', '2022-08-06 10:01:16', 'admin', '2022-08-06 10:01:21', NULL, '0', '大大啊啊', NULL, NULL);
INSERT INTO "public"."sys_dept" ("id", "parent_id", "ancestors", "dept_name", "order_num", "dept_status", "create_by", "create_time", "update_by", "update_time", "remarks", "del_flag", "leader", "phone", "email") VALUES (110, 100, '0,100', '这种消息', 555, '1', 'test', '2022-07-16 04:52:07', 'admin', '2022-04-10 01:36:23.72256', NULL, '0', '打算的撒啊啊', NULL, 'dsadaa@111.com');
INSERT INTO "public"."sys_dept" ("id", "parent_id", "ancestors", "dept_name", "order_num", "dept_status", "create_by", "create_time", "update_by", "update_time", "remarks", "del_flag", "leader", "phone", "email") VALUES (112, 110, '0,100,110', '1515', 1, '0', 'admin', '2022-07-16 06:45:21', '', NULL, NULL, '2', NULL, NULL, NULL);
INSERT INTO "public"."sys_dept" ("id", "parent_id", "ancestors", "dept_name", "order_num", "dept_status", "create_by", "create_time", "update_by", "update_time", "remarks", "del_flag", "leader", "phone", "email") VALUES (113, 112, '0,100,110,112', '大大撒啊', 1, '0', 'admin', '2022-07-16 06:45:30', '', NULL, NULL, '2', NULL, NULL, NULL);
INSERT INTO "public"."sys_dept" ("id", "parent_id", "ancestors", "dept_name", "order_num", "dept_status", "create_by", "create_time", "update_by", "update_time", "remarks", "del_flag", "leader", "phone", "email") VALUES (114, 113, '0,100,110,112,113', '2434343', 1, '0', 'admin', '2022-07-16 06:45:39', '', NULL, NULL, '2', NULL, NULL, NULL);
INSERT INTO "public"."sys_dept" ("id", "parent_id", "ancestors", "dept_name", "order_num", "dept_status", "create_by", "create_time", "update_by", "update_time", "remarks", "del_flag", "leader", "phone", "email") VALUES (115, 114, '0,100,110,112,113,114', '2222', 1, '0', 'admin', '2022-07-16 06:45:47', '', NULL, NULL, '2', NULL, NULL, NULL);
INSERT INTO "public"."sys_dept" ("id", "parent_id", "ancestors", "dept_name", "order_num", "dept_status", "create_by", "create_time", "update_by", "update_time", "remarks", "del_flag", "leader", "phone", "email") VALUES (116, 115, '0,100,110,112,113,114,115', '微微儿', 1, '0', 'admin', '2022-07-16 06:46:07', '', NULL, NULL, '2', NULL, NULL, NULL);
INSERT INTO "public"."sys_dept" ("id", "parent_id", "ancestors", "dept_name", "order_num", "dept_status", "create_by", "create_time", "update_by", "update_time", "remarks", "del_flag", "leader", "phone", "email") VALUES (117, 116, '0,100,110,112,113,114,115,116', '1111', 1, '0', 'admin', '2022-07-16 06:46:16', '', NULL, NULL, '2', NULL, NULL, NULL);
INSERT INTO "public"."sys_dept" ("id", "parent_id", "ancestors", "dept_name", "order_num", "dept_status", "create_by", "create_time", "update_by", "update_time", "remarks", "del_flag", "leader", "phone", "email") VALUES (118, 117, '0,100,110,112,113,114,115,116,117', '23232', 1, '0', 'admin', '2022-07-16 06:46:32', '', NULL, NULL, '2', NULL, NULL, NULL);
INSERT INTO "public"."sys_dept" ("id", "parent_id", "ancestors", "dept_name", "order_num", "dept_status", "create_by", "create_time", "update_by", "update_time", "remarks", "del_flag", "leader", "phone", "email") VALUES (100, 0, '0', '若依科技', 0, '0', 'admin', '2021-03-07 17:58:10', 'admin', '2022-04-10 01:27:19.51494', NULL, '0', '若依', '15888888888', 'ry@qq.com');
COMMIT;



-- ----------------------------
-- Table structure for sys_dict_data
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_dict_data";
CREATE TABLE "public"."sys_dict_data" (
                                          "id" int8 NOT NULL GENERATED BY DEFAULT AS IDENTITY (
                                              INCREMENT 1
                                              MINVALUE  1
                                              MAXVALUE 9223372036854775807
                                              START 29
                                              CACHE 1
                                              ),
                                          "dict_sort" int4,
                                          "dict_label" varchar(100) COLLATE "pg_catalog"."default",
                                          "dict_value" varchar(100) COLLATE "pg_catalog"."default",
                                          "dict_type" varchar(100) COLLATE "pg_catalog"."default",
                                          "css_class" varchar(100) COLLATE "pg_catalog"."default",
                                          "list_class" varchar(100) COLLATE "pg_catalog"."default",
                                          "is_default" char(1) COLLATE "pg_catalog"."default",
                                          "dict_status" char(1) COLLATE "pg_catalog"."default",
                                          "create_by" varchar(64) COLLATE "pg_catalog"."default",
                                          "create_time" timestamp(6),
                                          "update_by" varchar(64) COLLATE "pg_catalog"."default",
                                          "update_time" timestamp(6),
                                          "remarks" varchar(500) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."sys_dict_data" OWNER TO "yangge";
COMMENT ON COLUMN "public"."sys_dict_data"."id" IS '字典编码';
COMMENT ON COLUMN "public"."sys_dict_data"."dict_sort" IS '字典排序';
COMMENT ON COLUMN "public"."sys_dict_data"."dict_label" IS '字典标签';
COMMENT ON COLUMN "public"."sys_dict_data"."dict_value" IS '字典键值';
COMMENT ON COLUMN "public"."sys_dict_data"."dict_type" IS '字典类型';
COMMENT ON COLUMN "public"."sys_dict_data"."css_class" IS '样式属性（其他样式扩展）';
COMMENT ON COLUMN "public"."sys_dict_data"."list_class" IS '表格回显样式';
COMMENT ON COLUMN "public"."sys_dict_data"."is_default" IS '是否默认（Y是 N否）';
COMMENT ON COLUMN "public"."sys_dict_data"."dict_status" IS '状态（0正常 1停用）';
COMMENT ON COLUMN "public"."sys_dict_data"."create_by" IS '创建者';
COMMENT ON COLUMN "public"."sys_dict_data"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."sys_dict_data"."update_by" IS '更新者';
COMMENT ON COLUMN "public"."sys_dict_data"."update_time" IS '更新时间';
COMMENT ON COLUMN "public"."sys_dict_data"."remarks" IS '备注';
COMMENT ON TABLE "public"."sys_dict_data" IS '字典数据表';

-- ----------------------------
-- Records of sys_dict_data
-- ----------------------------
BEGIN;
INSERT INTO "public"."sys_dict_data" ("id", "dict_sort", "dict_label", "dict_value", "dict_type", "css_class", "list_class", "is_default", "dict_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1, 1, '男', '0', 'sys_user_sex', '', '', 'Y', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '性别男');
INSERT INTO "public"."sys_dict_data" ("id", "dict_sort", "dict_label", "dict_value", "dict_type", "css_class", "list_class", "is_default", "dict_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (2, 2, '女', '1', 'sys_user_sex', '', '', 'N', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '性别女');
INSERT INTO "public"."sys_dict_data" ("id", "dict_sort", "dict_label", "dict_value", "dict_type", "css_class", "list_class", "is_default", "dict_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (3, 3, '未知', '2', 'sys_user_sex', '', '', 'N', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '性别未知');
INSERT INTO "public"."sys_dict_data" ("id", "dict_sort", "dict_label", "dict_value", "dict_type", "css_class", "list_class", "is_default", "dict_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (4, 1, '显示', '0', 'sys_show_hide', '', 'primary', 'Y', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '显示菜单');
INSERT INTO "public"."sys_dict_data" ("id", "dict_sort", "dict_label", "dict_value", "dict_type", "css_class", "list_class", "is_default", "dict_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (5, 2, '隐藏', '1', 'sys_show_hide', '', 'danger', 'N', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '隐藏菜单');
INSERT INTO "public"."sys_dict_data" ("id", "dict_sort", "dict_label", "dict_value", "dict_type", "css_class", "list_class", "is_default", "dict_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (6, 1, '正常', '0', 'sys_normal_disable', '', 'primary', 'Y', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '正常状态');
INSERT INTO "public"."sys_dict_data" ("id", "dict_sort", "dict_label", "dict_value", "dict_type", "css_class", "list_class", "is_default", "dict_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (7, 2, '停用', '1', 'sys_normal_disable', '', 'danger', 'N', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '停用状态');
INSERT INTO "public"."sys_dict_data" ("id", "dict_sort", "dict_label", "dict_value", "dict_type", "css_class", "list_class", "is_default", "dict_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (8, 1, '正常', '0', 'sys_job_status', '', 'primary', 'Y', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '正常状态');
INSERT INTO "public"."sys_dict_data" ("id", "dict_sort", "dict_label", "dict_value", "dict_type", "css_class", "list_class", "is_default", "dict_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (9, 2, '暂停', '1', 'sys_job_status', '', 'danger', 'N', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '停用状态');
INSERT INTO "public"."sys_dict_data" ("id", "dict_sort", "dict_label", "dict_value", "dict_type", "css_class", "list_class", "is_default", "dict_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (10, 1, '默认', 'DEFAULT', 'sys_job_group', '', '', 'Y', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '默认分组');
INSERT INTO "public"."sys_dict_data" ("id", "dict_sort", "dict_label", "dict_value", "dict_type", "css_class", "list_class", "is_default", "dict_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (11, 2, '系统', 'SYSTEM', 'sys_job_group', '', '', 'N', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '系统分组');
INSERT INTO "public"."sys_dict_data" ("id", "dict_sort", "dict_label", "dict_value", "dict_type", "css_class", "list_class", "is_default", "dict_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (12, 1, '是', 'Y', 'sys_yes_no', '', 'primary', 'Y', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '系统默认是');
INSERT INTO "public"."sys_dict_data" ("id", "dict_sort", "dict_label", "dict_value", "dict_type", "css_class", "list_class", "is_default", "dict_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (13, 2, '否', 'N', 'sys_yes_no', '', 'danger', 'N', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '系统默认否');
INSERT INTO "public"."sys_dict_data" ("id", "dict_sort", "dict_label", "dict_value", "dict_type", "css_class", "list_class", "is_default", "dict_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (14, 1, '通知', '1', 'sys_notice_type', '', 'warning', 'Y', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '通知');
INSERT INTO "public"."sys_dict_data" ("id", "dict_sort", "dict_label", "dict_value", "dict_type", "css_class", "list_class", "is_default", "dict_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (15, 2, '公告', '2', 'sys_notice_type', '', 'success', 'N', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '公告');
INSERT INTO "public"."sys_dict_data" ("id", "dict_sort", "dict_label", "dict_value", "dict_type", "css_class", "list_class", "is_default", "dict_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (16, 1, '正常', '0', 'sys_notice_status', '', 'primary', 'Y', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '正常状态');
INSERT INTO "public"."sys_dict_data" ("id", "dict_sort", "dict_label", "dict_value", "dict_type", "css_class", "list_class", "is_default", "dict_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (17, 2, '关闭', '1', 'sys_notice_status', '', 'danger', 'N', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '关闭状态');
INSERT INTO "public"."sys_dict_data" ("id", "dict_sort", "dict_label", "dict_value", "dict_type", "css_class", "list_class", "is_default", "dict_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (18, 1, '新增', '1', 'sys_oper_type', '', 'info', 'N', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '新增操作');
INSERT INTO "public"."sys_dict_data" ("id", "dict_sort", "dict_label", "dict_value", "dict_type", "css_class", "list_class", "is_default", "dict_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (19, 2, '修改', '2', 'sys_oper_type', '', 'info', 'N', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '修改操作');
INSERT INTO "public"."sys_dict_data" ("id", "dict_sort", "dict_label", "dict_value", "dict_type", "css_class", "list_class", "is_default", "dict_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (20, 3, '删除', '3', 'sys_oper_type', '', 'danger', 'N', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '删除操作');
INSERT INTO "public"."sys_dict_data" ("id", "dict_sort", "dict_label", "dict_value", "dict_type", "css_class", "list_class", "is_default", "dict_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (21, 4, '授权', '4', 'sys_oper_type', '', 'primary', 'N', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '授权操作');
INSERT INTO "public"."sys_dict_data" ("id", "dict_sort", "dict_label", "dict_value", "dict_type", "css_class", "list_class", "is_default", "dict_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (22, 5, '导出', '5', 'sys_oper_type', '', 'warning', 'N', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '导出操作');
INSERT INTO "public"."sys_dict_data" ("id", "dict_sort", "dict_label", "dict_value", "dict_type", "css_class", "list_class", "is_default", "dict_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (23, 6, '导入', '6', 'sys_oper_type', '', 'warning', 'N', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '导入操作');
INSERT INTO "public"."sys_dict_data" ("id", "dict_sort", "dict_label", "dict_value", "dict_type", "css_class", "list_class", "is_default", "dict_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (24, 7, '强退', '7', 'sys_oper_type', '', 'danger', 'N', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '强退操作');
INSERT INTO "public"."sys_dict_data" ("id", "dict_sort", "dict_label", "dict_value", "dict_type", "css_class", "list_class", "is_default", "dict_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (25, 8, '生成代码', '8', 'sys_oper_type', '', 'warning', 'N', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '生成操作');
INSERT INTO "public"."sys_dict_data" ("id", "dict_sort", "dict_label", "dict_value", "dict_type", "css_class", "list_class", "is_default", "dict_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (26, 9, '清空数据', '9', 'sys_oper_type', '', 'danger', 'N', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '清空操作');
INSERT INTO "public"."sys_dict_data" ("id", "dict_sort", "dict_label", "dict_value", "dict_type", "css_class", "list_class", "is_default", "dict_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (27, 1, '成功', '0', 'sys_common_status', '', 'primary', 'N', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '正常状态');
INSERT INTO "public"."sys_dict_data" ("id", "dict_sort", "dict_label", "dict_value", "dict_type", "css_class", "list_class", "is_default", "dict_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (28, 2, '失败', '1', 'sys_common_status', '', 'danger', 'N', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '停用状态');
INSERT INTO "public"."sys_dict_data" ("id", "dict_sort", "dict_label", "dict_value", "dict_type", "css_class", "list_class", "is_default", "dict_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (33, 1, '33', '333', '5', '', '', ' ', '0', '', '2022-04-17 20:23:55', '', '2022-04-17 20:34:46.077261', '333333');
COMMIT;

-- ----------------------------
-- Table structure for sys_dict_type
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_dict_type";
CREATE TABLE "public"."sys_dict_type" (
                                          "dict_id" int8 NOT NULL GENERATED BY DEFAULT AS IDENTITY (
                                              INCREMENT 1
                                              MINVALUE  1
                                              MAXVALUE 9223372036854775807
                                              START 14
                                              CACHE 1
                                              ),
                                          "dict_name" varchar(100) COLLATE "pg_catalog"."default",
                                          "dict_type" varchar(100) COLLATE "pg_catalog"."default",
                                          "status" char(1) COLLATE "pg_catalog"."default",
                                          "create_by" varchar(64) COLLATE "pg_catalog"."default",
                                          "create_time" timestamp(6),
                                          "update_by" varchar(64) COLLATE "pg_catalog"."default",
                                          "update_time" timestamp(6),
                                          "remark" varchar(500) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."sys_dict_type" OWNER TO "yangge";
COMMENT ON COLUMN "public"."sys_dict_type"."dict_id" IS '字典主键';
COMMENT ON COLUMN "public"."sys_dict_type"."dict_name" IS '字典名称';
COMMENT ON COLUMN "public"."sys_dict_type"."dict_type" IS '字典类型';
COMMENT ON COLUMN "public"."sys_dict_type"."status" IS '状态（0正常 1停用）';
COMMENT ON COLUMN "public"."sys_dict_type"."create_by" IS '创建者';
COMMENT ON COLUMN "public"."sys_dict_type"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."sys_dict_type"."update_by" IS '更新者';
COMMENT ON COLUMN "public"."sys_dict_type"."update_time" IS '更新时间';
COMMENT ON COLUMN "public"."sys_dict_type"."remark" IS '备注';
COMMENT ON TABLE "public"."sys_dict_type" IS '字典类型表';

-- ----------------------------
-- Records of sys_dict_type
-- ----------------------------
BEGIN;
INSERT INTO "public"."sys_dict_type" ("dict_id", "dict_name", "dict_type", "status", "create_by", "create_time", "update_by", "update_time", "remark") VALUES (1, '用户性别', 'sys_user_sex', '0', 'admin', '2022-05-20 02:27:32', 'admin', '2022-05-20 02:27:32', '用户性别列表');
INSERT INTO "public"."sys_dict_type" ("dict_id", "dict_name", "dict_type", "status", "create_by", "create_time", "update_by", "update_time", "remark") VALUES (2, '菜单状态', 'sys_show_hide', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '菜单状态列表');
INSERT INTO "public"."sys_dict_type" ("dict_id", "dict_name", "dict_type", "status", "create_by", "create_time", "update_by", "update_time", "remark") VALUES (3, '系统开关', 'sys_normal_disable', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '系统开关列表');
INSERT INTO "public"."sys_dict_type" ("dict_id", "dict_name", "dict_type", "status", "create_by", "create_time", "update_by", "update_time", "remark") VALUES (4, '任务状态', 'sys_job_status', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '任务状态列表');
INSERT INTO "public"."sys_dict_type" ("dict_id", "dict_name", "dict_type", "status", "create_by", "create_time", "update_by", "update_time", "remark") VALUES (5, '任务分组', 'sys_job_group', '0', 'admin', '2022-05-20 02:27:32', 'admin', '2022-05-20 02:27:32', '任务分组列表');
INSERT INTO "public"."sys_dict_type" ("dict_id", "dict_name", "dict_type", "status", "create_by", "create_time", "update_by", "update_time", "remark") VALUES (6, '系统是否', 'sys_yes_no', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '系统是否列表');
INSERT INTO "public"."sys_dict_type" ("dict_id", "dict_name", "dict_type", "status", "create_by", "create_time", "update_by", "update_time", "remark") VALUES (7, '通知类型', 'sys_notice_type', '0', 'admin', '2022-05-20 02:27:32', 'admin', '2022-05-20 02:27:32', '通知类型列表');
INSERT INTO "public"."sys_dict_type" ("dict_id", "dict_name", "dict_type", "status", "create_by", "create_time", "update_by", "update_time", "remark") VALUES (8, '通知状态', 'sys_notice_status', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '通知状态列表');
INSERT INTO "public"."sys_dict_type" ("dict_id", "dict_name", "dict_type", "status", "create_by", "create_time", "update_by", "update_time", "remark") VALUES (9, '操作类型', 'sys_oper_type', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '操作类型列表');
INSERT INTO "public"."sys_dict_type" ("dict_id", "dict_name", "dict_type", "status", "create_by", "create_time", "update_by", "update_time", "remark") VALUES (10, '系统状态', 'sys_common_status', '0', 'admin', '2022-05-20 02:27:32', 'admin', '2022-05-20 02:27:32', '登录状态列表');
INSERT INTO "public"."sys_dict_type" ("dict_id", "dict_name", "dict_type", "status", "create_by", "create_time", "update_by", "update_time", "remark") VALUES (11, '测试', 'TEST', '0', 'admin', '2022-05-20 02:27:32', 'admin', '2022-05-20 02:27:32', 'SDADADAAA');
COMMIT;

-- ----------------------------
-- Table structure for sys_job
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_job";
CREATE TABLE "public"."sys_job" (
                                    "job_id" int8 NOT NULL GENERATED BY DEFAULT AS IDENTITY (
                                        INCREMENT 1
                                        MINVALUE  1
                                        MAXVALUE 9223372036854775807
                                        START 5
                                        CACHE 1
                                        ),
                                    "job_name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
                                    "job_group" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
                                    "invoke_target" varchar(500) COLLATE "pg_catalog"."default" NOT NULL,
                                    "cron_expression" varchar(255) COLLATE "pg_catalog"."default",
                                    "misfire_policy" varchar(20) COLLATE "pg_catalog"."default",
                                    "concurrent" char(1) COLLATE "pg_catalog"."default",
                                    "status" char(1) COLLATE "pg_catalog"."default",
                                    "create_by" varchar(64) COLLATE "pg_catalog"."default",
                                    "create_time" timestamp(6),
                                    "update_by" varchar(64) COLLATE "pg_catalog"."default",
                                    "update_time" timestamp(6),
                                    "remark" varchar(500) COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."sys_job" OWNER TO "yangge";
COMMENT ON COLUMN "public"."sys_job"."job_id" IS '任务ID';
COMMENT ON COLUMN "public"."sys_job"."job_name" IS '任务名称';
COMMENT ON COLUMN "public"."sys_job"."job_group" IS '任务组名';
COMMENT ON COLUMN "public"."sys_job"."invoke_target" IS '调用目标字符串';
COMMENT ON COLUMN "public"."sys_job"."cron_expression" IS 'cron执行表达式';
COMMENT ON COLUMN "public"."sys_job"."misfire_policy" IS '计划执行错误策略（1立即执行 2执行一次 3放弃执行）';
COMMENT ON COLUMN "public"."sys_job"."concurrent" IS '是否并发执行（0允许 1禁止）';
COMMENT ON COLUMN "public"."sys_job"."status" IS '状态（0正常 1暂停）';
COMMENT ON COLUMN "public"."sys_job"."create_by" IS '创建者';
COMMENT ON COLUMN "public"."sys_job"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."sys_job"."update_by" IS '更新者';
COMMENT ON COLUMN "public"."sys_job"."update_time" IS '更新时间';
COMMENT ON COLUMN "public"."sys_job"."remark" IS '备注信息';
COMMENT ON TABLE "public"."sys_job" IS '定时任务调度表';

-- ----------------------------
-- Records of sys_job
-- ----------------------------
BEGIN;
INSERT INTO "public"."sys_job" ("job_id", "job_name", "job_group", "invoke_target", "cron_expression", "misfire_policy", "concurrent", "status", "create_by", "create_time", "update_by", "update_time", "remark") VALUES (1, '系统默认（无参）', 'DEFAULT', 'ryTask.ryNoParams', '0/10 * * * * ?', '3', '1', '1', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_job" ("job_id", "job_name", "job_group", "invoke_target", "cron_expression", "misfire_policy", "concurrent", "status", "create_by", "create_time", "update_by", "update_time", "remark") VALUES (2, '系统默认（有参）', 'DEFAULT', 'ryTask.ryParams(''ry'')', '0/15 * * * * ?', '3', '1', '1', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_job" ("job_id", "job_name", "job_group", "invoke_target", "cron_expression", "misfire_policy", "concurrent", "status", "create_by", "create_time", "update_by", "update_time", "remark") VALUES (3, '系统默认（多参）', 'DEFAULT', 'ryTask.ryMultipleParams(''ry'', true, 2000L, 316.50D, 100)', '0/20 * * * * ?', '3', '1', '1', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
COMMIT;

-- ----------------------------
-- Table structure for sys_job_log
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_job_log";
CREATE TABLE "public"."sys_job_log" (
                                        "job_log_id" int8 NOT NULL GENERATED BY DEFAULT AS IDENTITY (
                                            INCREMENT 1
                                            MINVALUE  1
                                            MAXVALUE 9223372036854775807
                                            START 1
                                            CACHE 1
                                            ),
                                        "job_name" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
                                        "job_group" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
                                        "invoke_target" varchar(500) COLLATE "pg_catalog"."default" NOT NULL,
                                        "job_message" varchar(500) COLLATE "pg_catalog"."default",
                                        "status" char(1) COLLATE "pg_catalog"."default",
                                        "exception_info" varchar(2000) COLLATE "pg_catalog"."default",
                                        "create_time" timestamp(6)
)
;
ALTER TABLE "public"."sys_job_log" OWNER TO "yangge";
COMMENT ON COLUMN "public"."sys_job_log"."job_log_id" IS '任务日志ID';
COMMENT ON COLUMN "public"."sys_job_log"."job_name" IS '任务名称';
COMMENT ON COLUMN "public"."sys_job_log"."job_group" IS '任务组名';
COMMENT ON COLUMN "public"."sys_job_log"."invoke_target" IS '调用目标字符串';
COMMENT ON COLUMN "public"."sys_job_log"."job_message" IS '日志信息';
COMMENT ON COLUMN "public"."sys_job_log"."status" IS '执行状态（0正常 1失败）';
COMMENT ON COLUMN "public"."sys_job_log"."exception_info" IS '异常信息';
COMMENT ON COLUMN "public"."sys_job_log"."create_time" IS '创建时间';
COMMENT ON TABLE "public"."sys_job_log" IS '定时任务调度日志表';

-- ----------------------------
-- Records of sys_job_log
-- ----------------------------
BEGIN;
COMMIT;

-- ----------------------------
-- Table structure for sys_login_record
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_login_record";
CREATE TABLE "public"."sys_login_record" (
                                             "info_id" int8 NOT NULL GENERATED BY DEFAULT AS IDENTITY (
                                                 INCREMENT 1
                                                 MINVALUE  1
                                                 MAXVALUE 9223372036854775807
                                                 START 1
                                                 CACHE 1
                                                 ),
                                             "user_name" varchar(50) COLLATE "pg_catalog"."default",
                                             "ipaddr" varchar(128) COLLATE "pg_catalog"."default",
                                             "login_location" varchar(255) COLLATE "pg_catalog"."default",
                                             "browser" varchar(50) COLLATE "pg_catalog"."default",
                                             "os" varchar(50) COLLATE "pg_catalog"."default",
                                             "status" char(1) COLLATE "pg_catalog"."default",
                                             "msg" varchar(255) COLLATE "pg_catalog"."default",
                                             "login_time" timestamptz(6),
                                             "token" text COLLATE "pg_catalog"."default"
)
;
ALTER TABLE "public"."sys_login_record" OWNER TO "yangge";
COMMENT ON COLUMN "public"."sys_login_record"."info_id" IS '访问ID';
COMMENT ON COLUMN "public"."sys_login_record"."user_name" IS '用户账号';
COMMENT ON COLUMN "public"."sys_login_record"."ipaddr" IS '登录IP地址';
COMMENT ON COLUMN "public"."sys_login_record"."login_location" IS '登录地点';
COMMENT ON COLUMN "public"."sys_login_record"."browser" IS '浏览器类型';
COMMENT ON COLUMN "public"."sys_login_record"."os" IS '操作系统';
COMMENT ON COLUMN "public"."sys_login_record"."status" IS '登录状态（0成功 1失败）';
COMMENT ON COLUMN "public"."sys_login_record"."msg" IS '提示消息';
COMMENT ON COLUMN "public"."sys_login_record"."login_time" IS '访问时间';
COMMENT ON COLUMN "public"."sys_login_record"."token" IS '登录token';
COMMENT ON TABLE "public"."sys_login_record" IS '系统访问记录';


-- ----------------------------
-- Table structure for sys_menu
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_menu";
CREATE TABLE "public"."sys_menu" (
                                     "id" int8 NOT NULL GENERATED BY DEFAULT AS IDENTITY (
                                         INCREMENT 1
                                         MINVALUE  1
                                         MAXVALUE 9223372036854775807
                                         START 20
                                         CACHE 1
                                         ),
                                     "menu_name" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
                                     "parent_id" int8,
                                     "order_num" int4,
                                     "router_path" varchar(200) COLLATE "pg_catalog"."default",
                                     "component" varchar(255) COLLATE "pg_catalog"."default",
                                     "is_frame" int4,
                                     "is_cache" int4,
                                     "menu_type" char(1) COLLATE "pg_catalog"."default",
                                     "visible" char(1) COLLATE "pg_catalog"."default",
                                     "menu_status" char(1) COLLATE "pg_catalog"."default",
                                     "perms" varchar(100) COLLATE "pg_catalog"."default",
                                     "icon" varchar(100) COLLATE "pg_catalog"."default",
                                     "create_by" varchar(64) COLLATE "pg_catalog"."default",
                                     "create_time" timestamp(6),
                                     "update_by" varchar(64) COLLATE "pg_catalog"."default",
                                     "update_time" timestamp(6),
                                     "remarks" varchar(500) COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "public"."sys_menu"."id" IS '菜单ID';
COMMENT ON COLUMN "public"."sys_menu"."menu_name" IS '菜单名称';
COMMENT ON COLUMN "public"."sys_menu"."parent_id" IS '父菜单ID';
COMMENT ON COLUMN "public"."sys_menu"."order_num" IS '显示顺序';
COMMENT ON COLUMN "public"."sys_menu"."router_path" IS '路由地址';
COMMENT ON COLUMN "public"."sys_menu"."component" IS '组件路径';
COMMENT ON COLUMN "public"."sys_menu"."is_frame" IS '是否为外链（0是 1否）';
COMMENT ON COLUMN "public"."sys_menu"."is_cache" IS '是否缓存（0缓存 1不缓存）';
COMMENT ON COLUMN "public"."sys_menu"."menu_type" IS '菜单类型（M目录 C菜单 F按钮）';
COMMENT ON COLUMN "public"."sys_menu"."visible" IS '菜单状态（0显示 1隐藏）';
COMMENT ON COLUMN "public"."sys_menu"."menu_status" IS '菜单状态（0正常 1停用）';
COMMENT ON COLUMN "public"."sys_menu"."perms" IS '权限标识';
COMMENT ON COLUMN "public"."sys_menu"."icon" IS '菜单图标';
COMMENT ON COLUMN "public"."sys_menu"."create_by" IS '创建者';
COMMENT ON COLUMN "public"."sys_menu"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."sys_menu"."update_by" IS '更新者';
COMMENT ON COLUMN "public"."sys_menu"."update_time" IS '更新时间';
COMMENT ON COLUMN "public"."sys_menu"."remarks" IS '备注';
COMMENT ON TABLE "public"."sys_menu" IS '菜单权限表';

-- ----------------------------
-- Records of sys_menu
-- ----------------------------
BEGIN;
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1, '系统管理', 0, 1, 'system', NULL, 1, 0, 'M', '0', '0', '', 'system', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '系统管理目录');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (2, '系统监控', 0, 2, 'monitor', NULL, 1, 0, 'M', '0', '0', '', 'monitor', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '系统监控目录');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (3, '系统工具', 0, 3, 'tool', NULL, 1, 0, 'M', '0', '0', '', 'tool', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '系统工具目录');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (4, '若依官网', 0, 4, 'http://ruoyi.vip', NULL, 0, 0, 'M', '0', '0', '', 'github', 'admin', '2022-05-20 02:27:32', 'admin', '2022-05-20 02:27:32', '若依官网地址');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (100, '用户管理', 1, 1, 'user', 'system/user/index', 1, 0, 'C', '0', '0', 'system:user:list', 'user', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '用户管理菜单');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (101, '角色管理', 1, 2, 'role', 'system/role/index', 1, 0, 'C', '0', '0', 'system:role:list', 'peoples', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '角色管理菜单');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (102, '菜单管理', 1, 3, 'menu', 'system/menu/index', 1, 0, 'C', '0', '0', 'system:menu:list', 'tree-table', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '菜单管理菜单');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (103, '部门管理', 1, 4, 'dept', 'system/dept/index', 1, 0, 'C', '0', '0', 'system:dept:list', 'table', 'admin', '2022-05-20 02:27:32', 'admin', '2022-05-20 02:27:32', '部门管理菜单');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (105, '字典管理', 1, 6, 'dict', 'system/dict/index', 1, 0, 'C', '0', '0', 'system:dict:list', 'dict', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '字典管理菜单');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (106, '参数设置', 1, 7, 'config', 'system/config/index', 1, 0, 'C', '0', '0', 'system:config:list', 'edit', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '参数设置菜单');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (107, '通知公告', 1, 8, 'notice', 'system/notice/index', 1, 0, 'C', '0', '0', 'system:notice:list', 'message', 'admin', '2022-05-20 02:27:32', 'admin', '2022-05-20 02:27:32', '通知公告菜单');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (108, '日志管理', 1, 9, 'log', '', 1, 0, 'M', '0', '0', '', 'logininfor', 'admin', '2022-05-20 02:27:32', 'admin', '2022-07-28 08:16:18', '日志管理菜单');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (109, '在线用户', 2, 1, 'online', 'monitor/online/index', 1, 0, 'C', '0', '0', 'monitor:online:list', 'online', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '在线用户菜单');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (110, '定时任务', 2, 2, 'job', 'monitor/job/index', 1, 0, 'C', '0', '0', 'monitor:job:list', 'job', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '定时任务菜单');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (112, '服务监控', 2, 4, 'server', 'monitor/server/index', 1, 0, 'C', '0', '0', 'monitor:server:list', 'server', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '服务监控菜单');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (113, '缓存监控', 2, 5, 'cache', 'monitor/cache/index', 1, 0, 'C', '0', '0', 'monitor:cache:list', 'redis', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '缓存监控菜单');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (114, '表单构建', 3, 1, 'build', 'tool/build/index', 1, 0, 'C', '0', '0', 'tool:build:list', 'build', 'admin', '2022-05-20 02:27:32', 'admin', '2022-05-20 02:27:32', '表单构建菜单');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (115, '代码生成', 3, 2, 'gen', 'tool/gen/index', 1, 0, 'C', '0', '0', 'tool:gen:list', 'code', 'admin', '2022-05-20 02:27:32', '', NULL, '代码生成菜单');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (500, '操作日志', 108, 1, 'operlog', 'monitor/operlog/index', 1, 0, 'C', '0', '0', 'monitor:operlog:list', 'form', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '操作日志菜单');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (501, '登录日志', 108, 2, 'logininfor', 'monitor/logininfor/index', 1, 0, 'C', '0', '0', 'monitor:logininfor:list', 'logininfor', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '登录日志菜单');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1001, '用户查询', 100, 1, '', '', 1, 0, 'F', '0', '0', 'system:user:query', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1002, '用户新增', 100, 2, '', '', 1, 0, 'F', '0', '0', 'system:user:add', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1003, '用户修改', 100, 3, '', '', 1, 0, 'F', '0', '0', 'system:user:edit', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1004, '用户删除', 100, 4, '', '', 1, 0, 'F', '0', '0', 'system:user:remove', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1005, '用户导出', 100, 5, '', '', 1, 0, 'F', '0', '0', 'system:user:export', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1006, '用户导入', 100, 6, '', '', 1, 0, 'F', '0', '0', 'system:user:import', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1007, '重置密码', 100, 7, '', '', 1, 0, 'F', '0', '0', 'system:user:resetPwd', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1008, '角色查询', 101, 1, '', '', 1, 0, 'F', '0', '0', 'system:role:query', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1009, '角色新增', 101, 2, '', '', 1, 0, 'F', '0', '0', 'system:role:add', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1010, '角色修改', 101, 3, '', '', 1, 0, 'F', '0', '0', 'system:role:edit', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1011, '角色删除', 101, 4, '', '', 1, 0, 'F', '0', '0', 'system:role:remove', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1012, '角色导出', 101, 5, '', '', 1, 0, 'F', '0', '0', 'system:role:export', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1013, '菜单查询', 102, 1, '', '', 1, 0, 'F', '0', '0', 'system:menu:query', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1014, '菜单新增', 102, 2, '', '', 1, 0, 'F', '0', '0', 'system:menu:add', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1015, '菜单修改', 102, 3, '', '', 1, 0, 'F', '0', '0', 'system:menu:edit', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1016, '菜单删除', 102, 4, '', '', 1, 0, 'F', '0', '0', 'system:menu:remove', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1017, '部门查询', 103, 1, '', '', 1, 0, 'F', '0', '0', 'system:dept:query', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1018, '部门新增', 103, 2, '', '', 1, 0, 'F', '0', '0', 'system:dept:add', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1019, '修改', 103, 3, '', '', 1, 0, 'F', '0', '0', 'system:dept:edit', '#', 'admin', '2022-05-20 02:27:32', 'admin', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1020, '删除', 103, 4, '', '', 1, 0, 'F', '0', '0', 'system:dept:remove', '#', 'admin', '2022-05-20 02:27:32', 'admin', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1021, '岗位查询', 104, 1, '', '', 1, 0, 'F', '0', '0', 'system:post:query', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1022, '岗位新增', 104, 2, '', '', 1, 0, 'F', '0', '0', 'system:post:add', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1023, '岗位修改', 104, 3, '', '', 1, 0, 'F', '0', '0', 'system:post:edit', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1024, '岗位删除', 104, 4, '', '', 1, 0, 'F', '0', '0', 'system:post:remove', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1025, '岗位导出', 104, 5, '', '', 1, 0, 'F', '0', '0', 'system:post:export', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1026, '字典查询', 105, 1, '#', '', 1, 0, 'F', '0', '0', 'system:dict:query', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1027, '字典新增', 105, 2, '#', '', 1, 0, 'F', '0', '0', 'system:dict:add', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1028, '字典修改', 105, 3, '#', '', 1, 0, 'F', '0', '0', 'system:dict:edit', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (104, '岗位管理', 1, 5, 'post', 'system/post/index', 1, 1, 'C', '0', '0', 'system:post:list', 'people', 'admin', '2022-05-20 02:27:32', 'admin', '2022-05-20 02:27:32', '岗位管理菜单');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (111, '数据监控', 2, 3, 'druid', 'monitor/druid/index', 1, 0, 'C', '1', '0', 'monitor:druid:list', 'druid', 'admin', '2022-05-20 02:27:32', '', '2022-04-23 17:53:59.374655', '数据监控菜单');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1029, '字典删除', 105, 4, '#', '', 1, 0, 'F', '0', '0', 'system:dict:remove', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1030, '字典导出', 105, 5, '#', '', 1, 0, 'F', '0', '0', 'system:dict:export', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1031, '参数查询', 106, 1, '#', '', 1, 0, 'F', '0', '0', 'system:config:query', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1032, '参数新增', 106, 2, '#', '', 1, 0, 'F', '0', '0', 'system:config:add', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1033, '参数修改', 106, 3, '#', '', 1, 0, 'F', '0', '0', 'system:config:edit', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1034, '参数删除', 106, 4, '#', '', 1, 0, 'F', '0', '0', 'system:config:remove', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1035, '参数导出', 106, 5, '#', '', 1, 0, 'F', '0', '0', 'system:config:export', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1036, '公告查询', 107, 1, '#', '', 1, 0, 'F', '0', '0', 'system:notice:query', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1037, '公告新增', 107, 2, '#', '', 1, 0, 'F', '0', '0', 'system:notice:add', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1038, '公告修改', 107, 3, '#', '', 1, 0, 'F', '0', '0', 'system:notice:edit', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1039, '公告删除', 107, 4, '#', '', 1, 0, 'F', '0', '0', 'system:notice:remove', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1040, '操作查询', 500, 1, '#', '', 1, 0, 'F', '0', '0', 'monitor:operlog:query', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1041, '操作删除', 500, 2, '#', '', 1, 0, 'F', '0', '0', 'monitor:operlog:remove', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1042, '日志导出', 500, 4, '#', '', 1, 0, 'F', '0', '0', 'monitor:operlog:export', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1043, '登录查询', 501, 1, '#', '', 1, 0, 'F', '0', '0', 'monitor:logininfor:query', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1044, '登录删除', 501, 2, '#', '', 1, 0, 'F', '0', '0', 'monitor:logininfor:remove', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1045, '日志导出', 501, 3, '#', '', 1, 0, 'F', '0', '0', 'monitor:logininfor:export', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1046, '在线查询', 109, 1, '#', '', 1, 0, 'F', '0', '0', 'monitor:online:query', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1047, '批量强退', 109, 2, '#', '', 1, 0, 'F', '0', '0', 'monitor:online:batchLogout', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1048, '单条强退', 109, 3, '#', '', 1, 0, 'F', '0', '0', 'monitor:online:forceLogout', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1049, '任务查询', 110, 1, '#', '', 1, 0, 'F', '0', '0', 'monitor:job:query', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1050, '任务新增', 110, 2, '#', '', 1, 0, 'F', '0', '0', 'monitor:job:add', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1051, '任务修改', 110, 3, '#', '', 1, 0, 'F', '0', '0', 'monitor:job:edit', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1052, '任务删除', 110, 4, '#', '', 1, 0, 'F', '0', '0', 'monitor:job:remove', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1053, '状态修改', 110, 5, '#', '', 1, 0, 'F', '0', '0', 'monitor:job:changeStatus', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1054, '任务导出', 110, 7, '#', '', 1, 0, 'F', '0', '0', 'monitor:job:export', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1055, '生成查询', 115, 1, '#', '', 1, 0, 'F', '0', '0', 'tool:gen:query', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1056, '生成修改', 115, 2, '#', '', 1, 0, 'F', '0', '0', 'tool:gen:edit', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1057, '生成删除', 115, 3, '#', '', 1, 0, 'F', '0', '0', 'tool:gen:remove', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1058, '导入代码', 115, 2, '#', '', 1, 0, 'F', '0', '0', 'tool:gen:import', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1059, '预览代码', 115, 4, '#', '', 1, 0, 'F', '0', '0', 'tool:gen:preview', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1060, '生成代码', 115, 5, '#', '', 1, 0, 'F', '0', '0', 'tool:gen:code', '#', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1061, '缓存列表', 2, 6, 'cacheList', 'monitor/cache/list', 1, 0, 'C', '0', '0', 'monitor:cache:list', 'list', 'admin', '2022-05-20 02:27:32', 'admin', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_menu" ("id", "menu_name", "parent_id", "order_num", "router_path", "component", "is_frame", "is_cache", "menu_type", "visible", "menu_status", "perms", "icon", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (116, '接口文档', 3, 3, 'swagger', 'tool/swagger/index', 1, 0, 'C', '0', '0', 'tool:swagger:list', 'swagger', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '系统接口菜单');
COMMIT;

-- ----------------------------
-- Table structure for sys_notice
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_notice";
CREATE TABLE "public"."sys_notice" (
                                       "id" int8 NOT NULL GENERATED BY DEFAULT AS IDENTITY (
                                           INCREMENT 1
                                           MINVALUE  1
                                           MAXVALUE 9223372036854775807
                                           START 1
                                           CACHE 1
                                           ),
                                       "notice_title" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
                                       "notice_type" char(1) COLLATE "pg_catalog"."default" NOT NULL,
                                       "notice_content" text COLLATE "pg_catalog"."default",
                                       "notice_status" char(1) COLLATE "pg_catalog"."default",
                                       "create_by" varchar(64) COLLATE "pg_catalog"."default",
                                       "create_time" timestamp(6),
                                       "update_by" varchar(64) COLLATE "pg_catalog"."default",
                                       "update_time" timestamp(6),
                                       "remarks" varchar(200) COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "public"."sys_notice"."id" IS '公告ID';
COMMENT ON COLUMN "public"."sys_notice"."notice_title" IS '公告标题';
COMMENT ON COLUMN "public"."sys_notice"."notice_type" IS '公告类型（1通知 2公告）';
COMMENT ON COLUMN "public"."sys_notice"."notice_content" IS '公告内容';
COMMENT ON COLUMN "public"."sys_notice"."notice_status" IS '公告状态（0正常 1关闭）';
COMMENT ON COLUMN "public"."sys_notice"."create_by" IS '创建者';
COMMENT ON COLUMN "public"."sys_notice"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."sys_notice"."update_by" IS '更新者';
COMMENT ON COLUMN "public"."sys_notice"."update_time" IS '更新时间';
COMMENT ON COLUMN "public"."sys_notice"."remarks" IS '备注';
COMMENT ON TABLE "public"."sys_notice" IS '通知公告表';

-- ----------------------------
-- Records of sys_notice
-- ----------------------------
BEGIN;
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1, '温馨提醒：2018-07-01 若依新版本发布啦', '2', '\xe696b0e78988e69cace58685e5aeb9', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '管理员');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (2, '维护通知：2018-07-01 若依系统凌晨维护', '1', '\xe7bbb4e68aa4e58685e5aeb9', '0', 'admin', '2022-05-20 02:27:32', 'admin', '2022-05-20 02:27:32', '管理员');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (3, '办他议的展资信专却光。', '1', '进将事。入持是红。集拉维。布极较认适将。', '1', '闵斌', '2022-03-11 21:31:03.338329', '', NULL, '变处力。品设活周事组几边。济必红北金局大。
海状指思广理个。大放石格命造上。争少全。
般精界车立照被月。学即周日。极今率热况员极。
群作保只地总质儿开。圆切质还验是。任状对。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (4, '其与些少点定设数行。', '2', '改精元术等手许一层。得界知原流容。分青身。快使期化感段式等水管。例较听响各。权常非理民。', '1', '牵奕辰', '2022-03-11 21:32:43.094855', '', NULL, '取处海头即的际内。叫那前于温设。动属次根改长每。
和酸全明权始织北是。低车适今命比式。性保队阶。
内八种表场济理包成建。自断特由。民人斗识立真论中志。
达至直命力至点。写外众界才者们路联结。系组电自半发也记包。
做将低。克争次确观她主并整当。火及最。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (5, '目地造万他行。', '1', '太题相连划。数学史品求明局。件图连阶。', '0', '森娜', '2022-03-11 21:33:04.636901', '', NULL, '通见知。年海王等热十周业。名目次。
图难响。音须价改无象油。集法使火市主你好多部。
属别色将地比作领作。红长间结五。广率展命是需效。
天农安图将京清。花西更置拉半放林。二知存段相真。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (6, '点种总斗层整节论发比。', '2', '种好联列之光色除。太类越通。率白难。心任点对律对想七量个。出情品么原数往都。', '1', '寇凤英', '2022-03-11 21:33:05.446812', '', NULL, '感持数识被走三决压走。什前亲特研。子易上百则级常压第又。
速被际期程己斯事。被表队则比可见界严。约分构半须已程们九。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (7, '除求选。', '2', '看内里产。收心花要清标图。', '1', '无万佳', '2022-03-11 21:33:06.0759', '', NULL, '极小南。的水使件开出。就亲化会步与积。
白合运构运。什生导号市给产参。据打青步派果你点被众。
场作准除连八。线中代便着别型组为身。展类科相声管后气决。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (8, '回验应容立。', '2', '下斯酸严按老及。就书计图。容比安队成力六半。金建层为几步学水速。角影加物适究。', '0', '乌雅俊杰', '2022-03-11 21:33:06.834518', '', NULL, '明已山以想本队代。万也石。议果厂术对较着北。
要该名。采都因好开子自经比快。八边己并音使采整感。
没油做质。县些变处按民处使合例。我完支白天最。
所金低压。照音光识道。建红干相。
什飞织难。入候亲候代面。五段海世存马手京。
金多阶查明它县元系有。提太此。意意子位易放这厂和。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (9, '须见许于指进问识也。', '2', '千有般往存。石体近照置价开强约。道只县。重出月。', '1', '睢婷婷', '2022-03-11 21:33:07.618164', '', NULL, '完儿千十。便亲土调近动日风身之。圆目林精该收。
次研引以精百音队。维只必是根况变。状开天用有局圆门。
变称关性值这成安。意美压量所才存至条因。总越世行了再议当办。
集飞一给办。条千期观选林指。音程从路号生空流条者。
去才者老地候于器放称。状类上各消前报的作越。年思走劳须拉图。
全心就。风以地写分内数计他。会建调表号全厂干速算。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (10, '律速加么想义验他电节。', '1', '发中过到近转离再。话并受热此龙识制。区只须会学度。么广温十产放速。行结米别更阶。受学能具取识油电书。', '0', '守蒙', '2022-03-11 21:33:08.176854', '', NULL, '光可部员发书。步情片民角律太声备。选务间来图名志该。
人群得样写区动。般照之备。值更思备切文持干三。
从记身。员快万业联断。放造济件织参。
五去行火提最之眼。方取手能。产利组。
你酸育。电准例快图后了。引团非红表斯高素流只。
识南育式算世设象严式。极切当过直展去值制导。器最半构。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (11, '作支十七响派持约。', '1', '价个车包参例是确日公。热动市压局青石族内。事儿节她务究发但太给。选常它深。系再增则结流至取达。', '1', '尚涛', '2022-03-11 21:33:08.767304', '', NULL, '华三表通队斗。委别格非。须体况及较。
由多取正县历交管。相几规式。少值感历。
式花织好八调同接他。当且状等列上油路。布活取这且制万。
术经红东农土。农叫号则候作手样门制。极阶委无过生县报法这。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (13, '公向命电力它组。', '2', '但劳建小但格。员铁容使应经现此。', '0', '丁乙萍', '2022-03-11 21:33:10.124574', '', NULL, '着白使元自老亲。应多带下展理思号把感。她间取使如。
用感半子西支电干六。族王斗千里管学。装节程以水铁。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (14, '马做与算。', '2', '红以步物它确性水资产。确精不工圆。主机转体很车须快约。', '1', '碧鹏', '2022-03-11 21:33:10.733767', '', NULL, '开小立。手日知极论种连连打西。义集地省采。
说争间土商时过可。期石至光半压研声。根压连严争带要消。
划们开到她认律。少什通派华起。况写题象。
值统上指重年全切整。省决科所少切任温场。酸统布三精有代群资类。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (15, '备个群看力明时。', '1', '究理阶置那。周商教所提记。行有局别处。西型被性。', '0', '丛国珍', '2022-03-11 21:33:12.484055', '', NULL, '线很八。安千大族南群点当。必同众才内例酸里离。
起育样目器取把王以。论权因使次代月作类。构器难。
东直文算约少白明力计。青分点界。一动起样根。
叫连点斗在样矿角。真都类定组完白。需料效时华年。
持总不被六命。理书展照克矿强好议。九设取委任制于。
其前风。带电问六转。交低张市业育交今反片。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (16, '也当火需县候改却就安。', '1', '叫车万量积能度算斗斗。只代常层你度安维气。', '1', '孛桂英', '2022-03-11 21:33:13.063981', '', NULL, '道类系积。流确划解今转头开给应。色金相非。
队更线形育劳内。根各节二动受就。称听真统做给气风石打。
复具现常再率。明证经究验只专院眼。高布打思具。
作土要点交界头目。大面切。反例南全世完史书确月。
特类制。形代青。选件段年。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (17, '结已造名成深外都状。', '2', '样高究术越活海北向。空天或周角想。', '0', '诸葛呈轩', '2022-03-11 21:33:13.626395', '', NULL, '们农细去发研利员。亲效特下场关相北明部。构而许书派石严进文。
更通直知日技几。级农活变入给是原。经低容需且。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (12, '书意如至最机组张或问。', '2', '列千式先要养。回列程必受将你加。观而持。造步增集身进水。联指公根内周该中厂这。', '1', '衣军', '2022-03-11 21:33:09.528098', '零安琪', '2022-03-11 21:34:30.499514', '面叫什许。争中按节极情。况自知性过长节比者。
适花满增线别着。总入维近。王机年取老即商就率。
影片青好属。带便面求几眼见变原。物称还音变听及想列。
象写多进保全候验。县且酸类。按实建志并步己名回照。
成关大易养最局做火。解资什。百圆务精。
油求象新组状民。式再制算已律种决接。里流始次观。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (18, '引学表示打力人话步叫。', '2', '至大产有。想级关通参事通最难除。须必石。角工真。', '0', '兴阳', '2022-03-11 22:05:36.967325', '', NULL, '眼开马确。教手电特重拉。看近美史。
只分三及行。产定议品内清万。自积采力局所个。
这非交克系局能标们。或原会月制回。口己中近积选。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (19, '理选矿构力。', '2', '子能织界。才口史型。', '1', '盘志国', '2022-03-11 22:05:36.967325', '', NULL, '象报气日行红今见。回名议。东象连。
件标清北。去主铁却北各总组团。线外华改指公。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (20, '教代区器取议信。', '1', '工记有然。选从日头全基求率这体。自文把这。我题调解会所车拉。', '1', '鄞家明', '2022-03-11 22:05:36.967325', '', NULL, '京识历求任。花使特参子七再院。非造期。
从际统角量中形查。制消据儿热。八克长院要物。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (21, '都强改节养受。', '2', '许命级。放位京别而无民。情土情况石。带到真主。美好对色通太。但层以会。', '1', '仆桂兰', '2022-03-11 22:07:37.55757', '', NULL, '得类线近适个前。白实极。百养劳对。
同各书话。该光使度次提品很火。向十家记非化到活。
收认工团具好快设委处。了有置体将阶且积局发。王造列。
那易组治。根育间华放使实样。果被深文节把水地后指。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (22, '报律外离量。', '1', '分局同。过响来交办院通响除。', '1', '单玉梅', '2022-03-11 22:07:37.55757', '', NULL, '被率化听通影选机则便。权装事极。约是置有想部空水。
好南经本。重办管无手报温研本次。为标般际车生流示单。
运中根大。路能局和族压历。式次思。
各部给北现。组龙白论。斯结始。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (23, '给化并思热进片圆但省。', '1', '你济约转律文置过件说。分全般铁门大备该见来。', '1', '蒿国香', '2022-03-11 22:07:37.55757', '', NULL, '里照联。都会周长。式相层安空术。
人使准眼何包争属许工。规带太治员。成太局学第一。
次化才步要导记。清金各争六。实求他极。
改全象究热示南太家话。结确算今将张厂热片然。转能斯段十法正影根单。
次类商采接能率信用。克示你车员书。个程场。
实际直。周世你两。受系学。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (52, '边论品文。', '1', '小之候手权。情极亲公县斯劳青观半。最办着体期之许收示市。', '0', '孔雯静', '2022-03-12 02:17:09.254148', '', NULL, '造油前王口外步别。大中作。各存置具指查酸能才压。
领林术把经况路天。济历地。三到率土做使并。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (35, '打团石型。', '2', '数支以她都定并。开切以布织日西心。', '0', '于国秀', '2022-03-12 00:07:47.397833', '', NULL, '越向道。如论象。制准方千列商区机都该。
出取适老质最入公步下。人分交话见为。取知风。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (36, '例全装快。', '1', '明它之从叫可海使看业。场万八际比心事矿增。', '0', '占娜', '2022-03-12 02:03:13.714424', '', NULL, '圆料之已学周种般斯。问路林市素有传九低。风然分第。
人价响商角完。知体种京转我精铁经信。易阶次形认。
其严红极各老。路行群决当位育。她公油地。
话整常。必关义。达号消律红压。
清活正亲照利。反样再。成高先例相适。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (37, '形方全器属好。', '2', '指相备制着。识产须常说育程热。论来通群后转六光三。民当化。', '0', '栋国良', '2022-03-12 02:03:15.865577', '', NULL, '花目史达起给上属整前。王选温难条老。具因却今单关。
代认几广题认。写往立。月数见打步着。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (38, '成时识许状队。', '2', '听题定多入。放思济影。再完性县增。广结种表回据。', '0', '卜俊凯', '2022-03-12 02:03:20.790601', '', NULL, '院拉展已派几结。布定空厂今按量按。需并干系小空。
月质间一米情。治民深。我主要她积水多做研。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (39, '不风群比将圆。', '2', '小工交场点从六消号。标儿原米可回议历风。', '1', '滕燕', '2022-03-12 02:03:22.185455', '', NULL, '选七你至住。在观别是走。还日器复间公农。
族几两思保合关和。已到技比给七广件老起。加离半都。
断导需决。层本节适别究头酸空话。周至热置定值离色铁。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (40, '走同根志。', '1', '委布所样。京行说海。', '1', '景子豪', '2022-03-12 02:03:23.28924', '', NULL, '使将段。使由存向。增须们节己会合身除。
样导们已着。照角照件明做正保需所。着县联进单容叫。
量立下或代带。青素民主风通积。马越级。
来强九局年。型把半装取。系明更住收区适。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (41, '立金业基矿下局现接越。', '1', '厂表界书备值分后。派较团她听质。年准样际与真育听系。', '0', '娄雨桐', '2022-03-12 02:03:24.335096', '', NULL, '空品周。地状油认。干管精育查织律议。
学界写。流还活时。片适传日作。
她想精及百统进划。二时之次较从养。七放正格好住看置。
九亲况清大使。或厂带委步矿感风复。号空想只但育率声种多。
音周每都克集上。期圆里内生当商整前感。调员处因商话。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (42, '系多同传和基重。', '1', '中建断照低是农实回。至在何表称众果。', '1', '肥智杰', '2022-03-12 02:03:25.3942', '', NULL, '今值来志周光从。九界程再些边状验。集京律较气快领型。
现时支。节权头。成光接。
本交局。部为子九都铁上统没小。领加厂矿以等龙。
值公科。过今电里一今。者用主属好何。
其市精温式。度值上率。整效员。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (43, '置族心加都素力你。', '1', '取部划当。听产设明你族白具大近。住近争后力。得计写市选管准。海按近老省过员目比百。度半规但。', '1', '道诗雨', '2022-03-12 02:03:26.293513', '', NULL, '山重资划统。打立全别认约没。展月育。
六变集太专导。划加元化周。清心易省适车成效光。
半达最到多去。原低断多。始计图养便切受准。
变七动引常米你置。验因适东真前教科小。它号通边华使则电。
圆一规部来矿组委一她。近提生拉交声上厂。传你备由前红土打查。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (44, '车每海。', '1', '然身况适然。极全受可准他离。价斯手。', '0', '饶军', '2022-03-12 02:03:27.031205', '', NULL, '候新酸自马龙心没。目养支法。选建细作号此响状东连。
还断外专分半飞心化叫。使论确体。铁自管解住着数海。
了着心农例。院后铁。比一原到。
并教实展然务热对。级状造转争。人风产回立。
利单识压家构比。识作老统严里位位。部且命看。
别况的切龙。安做题那入。不片队此达术已。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (45, '农什路正型率。', '2', '果根信义专合活组。它世队许式叫九因织又。议金自做便造格小主。通百利团太红土。', '0', '瓮辉', '2022-03-12 02:03:27.740296', '', NULL, '受权白过。头体机果。多土证话度它石意。
白县包文。基基题区原分就六。入议么就定值。
方热使该。色真号候往争。对各回区写。
影程转次真此。但万验。的斗一联属。
太应及。许以制结无思。全区运养积断传市。
非近今期。东五难更流权通率连统。自究天下中七。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (46, '阶识提青。', '2', '王组主斗观。如发事须技查。专知她这来支县分并。给报相金局记何到。', '1', '孔治涛', '2022-03-12 02:10:27.376294', '', NULL, '点件红与展采高必京时。算议原积色。儿外林。
出己基北。满回状必。经步次老支县金示。
数例集千主美省。变务那。规什光。
千清传口转段山原门。易运示清。低动已交务市应。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (47, '的叫心美精叫。', '2', '品表级温议亲名。者意之持领列经。须准的给。何门部打速元据。证己等需市听快。', '0', '施紫林', '2022-03-12 02:11:08.776294', '', NULL, '车不光照。气现从照温术派的。任林风志火斯再。
角道收劳门重直长而。难采元会需计。然来第。
儿利解又年问主金。性间太以式科。美子细安劳。
商历月没无。光团己写部。前据接人。
布石运来列。主着住位先象。划京织者也须土同技积。
领节求克元。领团小织装。节引带观来儿利电。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (48, '可年无现眼要外。', '1', '进算广和任千义。给对干位。书受边车。他精动成。达者复学劳史格改。验和流但长。', '1', '普晨阳', '2022-03-12 02:11:33.411972', '', NULL, '斯们增重实周列前。低来义系果从识。历正细确下好说。
将达命。商机利领来大业千在。面相车东公广运工调次。
达技必名验意部。般南养引选次权离米装。必组位速号区住收京。
此于得。科出文基引具论。受分响而反。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (49, '多正情。', '1', '状造际格决系米叫再。本事次定术学达通以。收表式。信片温本山严影人气。叫龙动起资阶具个。在立般厂如究报。', '1', '宇萍', '2022-03-12 02:14:29.154628', '', NULL, '式么音。事问做条口种。片时亲该联。
眼音们常正关。感合有品已再去期。会格结心音解。
结须调值热半立并认干。程本亲除科把质事。安开界往年度话直。
教写设质他任究。火引自节资支。容成将交等主常亲之。
往连走可目。己阶约日装。位马即史值。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (50, '义门一力书于须展做水。', '2', '完三几上意。查九速观相状亲华拉。法因石我。世价合子信。少产区。', '0', '公羊玉英', '2022-03-12 02:15:02.178007', '', NULL, '养量一劳去运而和。放太八。规连别回天院被它起持。
总权太非百。空素少物存格听矿图温。件论车化些。
任南美力段步六几干。断许温月按花中信。号东美通强达。
情认海展一走济。加率将际格增类了张。例向长积按斗儿。
了做石处省具直认才。人于证数。料今保日层只前但。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (51, '用决无第干即将着除。', '2', '主眼住展花教学多。放油少变型老光最儿使。争细了增们。对何把头选器火和。团本着验。际品快置查办公系质声。', '1', '铁家豪', '2022-03-12 02:16:18.639083', '', NULL, '张眼设。程种铁想存决发。群维示划。
百群候最多观。变二应干走基果程事安。路式组见新对规计今。
光期验必。支不王长当科农。真少起小。
命调整你音体例。压走存用提带且。当十元向声。
名面八。参农斗打代象论快下年。更起号年。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (53, '段于条感完。', '1', '能识五生分。运复机无打。速一几治。半管采本石干建候。火道张面派存非治。京石办从总方。', '0', '黎丽萍', '2022-03-12 02:18:59.702396', '', NULL, '它元状。土加据这设。高物天八意。
叫影七听千强北。集和日元多得干。六元只公高。
住准火无周在。正史命单立。包基包切资求水准。
须济们。广规影反。老上统精群儿主局后。
规关育。劳全次得消立酸价。何得细过能改。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (54, '厂法难周设海公式家性。', '1', '对治速连立满名研出取。何量造都原离示。见想然至收照识。之决精没断复类。', '1', '敏俊凯', '2022-03-12 02:19:14.853753', '', NULL, '求料目写则林消年学。比育率复光由引表称。又引真节技先我须区。
三期口经九构支。办历太路。斯展安。
热目需与角步列场属构。金连通比指严时效南世。号人论十装。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (55, '见际再比情采别低影。', '2', '议数小我向回地华。存规位教温目最理求。支光较想速少持路认。改很九级多级意。整西天五求取众位南九。', '0', '哀沐阳', '2022-03-12 02:19:16.755351', '', NULL, '叫做海高产有中义步响。交年米东白完本切。家记或。
及安同件。列然里区太方物划。些看程。
且例业也任建最老切。物变构青类统声。动济指装资装成。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (56, '西价亲。', '1', '向多机。片间造。青标六与角省果定即。', '0', '乌雅浩晨', '2022-03-12 02:19:18.382169', '', NULL, '文情根照经。将收整七结保大关经自。可少今验好细。
得行人学容调切然那就。人通每些。铁他市。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (57, '建带照教音件花飞。', '2', '快目这称性研。文进志目细角边。选山可团。问信地大所发。位利题建场组民。', '1', '但家豪', '2022-03-12 02:19:20.781377', '', NULL, '高养没也日。从面运标段特设单打。目名步深住物自成委历。
热广细级层容光组对。入们住写化京派其量心。近交生外持林。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (58, '么六上。', '2', '样上管流交清争。开问指识以林米片号。我响统。解没形己效。什重气拉造用现安得看。我属率至保定拉气接。', '0', '励雅婷', '2022-03-12 02:19:22.512541', '', NULL, '车中只制好解教证式导。理本什级发利。种状光水断叫。
员百标青细。常划结思。划高正别。
品之市争张品制花后。压常中条集采月。起权保。
集物将各。还写象算部期。该样因领。
声个件。我素先快。动治标复天们。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (59, '米增比好细美。', '1', '者据外它。被能位特带着。动路油西京能。将单队用土。重支许类小。', '1', '睦若汐', '2022-03-12 02:19:23.444939', '', NULL, '各队山设打关十据性理。飞存在十断例节开。照根带去土步形标提确。
近月回品声。更流群常常地。交一用基油都。
老满又过济照一越。每调局复量起众备。制用点员先相适色度。
两不安家计样效除。济而第己程青一示列。所己规。
新定广往农。温率具意用前以形形。般且直专。
里术到才已。属特划状西交。都儿知入文领况表必。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (61, '又感门资质教样历什。', '2', '义有山听。书相米联团员派管重类。王上准今对离当级。之十整易条一界花已。', '0', '隋雅鑫', '2022-03-12 02:21:00.281183', '', NULL, '定现法。论段至一万交飞快然。须并十斯工以候力。
过般万每单力列干。种和须选话着也求相。省图火。
半接文由。他算成元。再气己置队。
么解有认究电二委因建。完适识状题不易。放不织世方张容各。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (62, '建况究。', '1', '完增主便委王标。些规金间。收号及或铁压果。', '0', '虞国珍', '2022-03-12 02:21:04.095613', '', NULL, '较身圆。千层很变始拉需算。传导处他料组听。
先划收件如据。之温制调数作规。年重形式土已从很。
技持造。示达教值到万。门土形无算想。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (63, '山去与。', '1', '说事太之。起更近般所务称从无。信温华图厂本民。', '1', '林玉梅', '2022-03-12 02:21:05.142963', '', NULL, '交对六。完象和克象细美备十们。统交表。
整边感说花何至。变改子价深。节统直文改学进十队分。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (64, '相称等角听公东六。', '2', '记角那住。儿价级。素史日代飞不土。风约解常权。速验越。', '0', '蹇玉珍', '2022-03-12 02:21:07.599682', '', NULL, '设议用明油影说六情。质原西据地。做基类长养量特即题放。
至感为当流看。矿往委受石发外认真制。式称较而都队的。
取十他民不。金满太。自金治组太。
于之世东选长方。和容被白准明好细。六代且十群然律斗机。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (65, '单门第。', '1', '地相装术划林么近造。对建保品争族精报然百。把内场复所。万识受。', '1', '俎婷婷', '2022-03-12 02:21:09.371451', '', NULL, '周用我分。格达表。派验记合。
比力几形做之影前。及正装性济当价根任面。展公图。
表求三面感什变则。必七活改取不飞际越。水写接的角经人。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (66, '动适高选般问声。', '1', '音先众空并还已。器部她支反压别中如。斗三和才。强义维率养构量南。始等打学改又没然正。际自前很有青力千。', '0', '简思佳', '2022-03-12 02:21:10.437617', '', NULL, '见说车经收斯最二。能支热象以院。设精经先义适引。
京去备志。把么因只长引存南。中治九间将土。
石资张教飞素。议制持复持。维造院体而热铁资。
反示百体西示做养也列。小提才非只以电路取。层识就规开验最都空里。
况派切线五量为许素。任标每无得其。声取人。
图厂设状。如转们很说为工育花包。处建状。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (67, '备员可。', '2', '格很第备养。间达存分称全离被指。联必影导月么。非育结面候能所年。', '0', '甘丽芳', '2022-03-12 02:21:11.292042', '', NULL, '就义交求理北车改边约。办了次派。广水严铁再改风究放。
般己般调七。起区空西。始才青。
少定领等世。始除命具委实除。温斯构二属交己青群建。
应海有风龙持议。酸率叫身响复化。类油查须般白导然题并。
之么土程。反活白造需心质也导。量文你与力养总。
什业目便一员见员。青种论感。与加建种派多千素儿效。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (68, '员还局。', '2', '月传北力活。易口更。', '0', '岑成', '2022-03-12 02:21:12.348768', '', NULL, '热边划圆白无。感信比京大就们离。但如除。
始对目。正集铁你地放不们率身。区无地。
只步律产带流按至月。儿即表条及现是达目查。华决族样族色大名达复。
性必公。历决然。包花候老例布。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (100, '成色见三单次。', '1', '为设书业到山样可。民商法。金直低标象中。', '1', '茆婷方', '2022-03-12 03:17:29.222451', '', NULL, '放任流经点过列难我亲。美领斯那始支山格。者西住。
图各即业查于观周比。持须常还支装那众。真条华领里由才油。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (69, '外复式们物直交决历知。', '2', '最术于能火种系油济热。观没教值性周。', '0', '嬴治涛', '2022-03-12 02:21:13.259803', '', NULL, '认达空具说。八论她却得精。许须便们。
铁系从明农使看。人用为。联式候你民日了拉部来。
文压两队一层改度。经新劳知林儿广此。目持把影。
集活断克。识据水只。油精部设但相较。
风理马查数没公处受。经土只只明积打去明。人史段海如县专。
从算去利为务大时地间。派农下工对酸复。后义米传将物那连然但。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (70, '队件心其。', '2', '农调利。张片行将速青书年。东矿价整种。成才称断上证难受果还。至导月队京石具极。土龙状。', '0', '师梓涵', '2022-03-12 02:53:26.60218', '', NULL, '利成特运达观。斯形说出长还知大会。建等龙这转半及什油三。
目把车半决题电节。机少信利将面。例候半红节。
情方引。来运内间的为立。连好比识消容。
照太经性家说这。应后合火着期给况。公它象根口用。
信织作领。角其法号例局下己技。斯下音整。
按法安间。还义状者压济。进计且又马角直二厂。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (71, '队件心其。', '2', '农调利。张片行将速青书年。东矿价整种。成才称断上证难受果还。至导月队京石具极。土龙状。', '0', '师梓涵', '2022-03-12 02:54:27.546184', '', NULL, '利成特运达观。斯形说出长还知大会。建等龙这转半及什油三。
目把车半决题电节。机少信利将面。例候半红节。
情方引。来运内间的为立。连好比识消容。
照太经性家说这。应后合火着期给况。公它象根口用。
信织作领。角其法号例局下己技。斯下音整。
按法安间。还义状者压济。进计且又马角直二厂。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (72, '队件心其。', '2', '农调利。张片行将速青书年。东矿价整种。成才称断上证难受果还。至导月队京石具极。土龙状。', '0', '师梓涵', '2022-03-12 02:56:46.723974', '', NULL, '利成特运达观。斯形说出长还知大会。建等龙这转半及什油三。
目把车半决题电节。机少信利将面。例候半红节。
情方引。来运内间的为立。连好比识消容。
照太经性家说这。应后合火着期给况。公它象根口用。
信织作领。角其法号例局下己技。斯下音整。
按法安间。还义状者压济。进计且又马角直二厂。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (73, '队件心其。', '2', '农调利。张片行将速青书年。东矿价整种。成才称断上证难受果还。至导月队京石具极。土龙状。', '0', '师梓涵', '2022-03-12 03:00:17.440374', '', NULL, '利成特运达观。斯形说出长还知大会。建等龙这转半及什油三。
目把车半决题电节。机少信利将面。例候半红节。
情方引。来运内间的为立。连好比识消容。
照太经性家说这。应后合火着期给况。公它象根口用。
信织作领。角其法号例局下己技。斯下音整。
按法安间。还义状者压济。进计且又马角直二厂。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (74, '队件心其。', '2', '农调利。张片行将速青书年。东矿价整种。成才称断上证难受果还。至导月队京石具极。土龙状。', '0', '师梓涵', '2022-03-12 03:06:36.039699', '', NULL, '利成特运达观。斯形说出长还知大会。建等龙这转半及什油三。
目把车半决题电节。机少信利将面。例候半红节。
情方引。来运内间的为立。连好比识消容。
照太经性家说这。应后合火着期给况。公它象根口用。
信织作领。角其法号例局下己技。斯下音整。
按法安间。还义状者压济。进计且又马角直二厂。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (75, '队件心其。', '2', '农调利。张片行将速青书年。东矿价整种。成才称断上证难受果还。至导月队京石具极。土龙状。', '0', '师梓涵', '2022-03-12 03:09:11.182325', '', NULL, '利成特运达观。斯形说出长还知大会。建等龙这转半及什油三。
目把车半决题电节。机少信利将面。例候半红节。
情方引。来运内间的为立。连好比识消容。
照太经性家说这。应后合火着期给况。公它象根口用。
信织作领。角其法号例局下己技。斯下音整。
按法安间。还义状者压济。进计且又马角直二厂。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (76, '很带候场边算使属需身。', '1', '率其入对从会。位员很量米单统飞界。证存作。', '0', '抗家明', '2022-03-12 03:09:39.483204', '', NULL, '只维动包素。在连温直才好务。计三才。
律技西。书任干才那你力离有质。个委值器后我办容。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (77, '地起她放前通带后。', '2', '只二机求用装。条年义土把了选他放。', '1', '函治涛', '2022-03-12 03:17:16.798081', '', NULL, '第划验压力进期从。我三增次并着等百原。气候近物史上织音华体。
建都你产。亲算能三打书。老也二更始参些火无。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (78, '飞养包装能根。', '1', '市节除六方。发你斗现织。证高结信器精然就而展。感研各书心。', '1', '富秀兰', '2022-03-12 03:17:17.619606', '', NULL, '离间精。农五较。该花者候用位出解及。
始米大科。文消华米造同积存油。再厂自图亲马。
写话那新目定学片称委。效转极。可明月教理。
样生法近自。志根门员积才家在厂石。们形用非可作和布改。
美度矿。存属增把历子石步这号。么外成式深反整。
火为传世去过至条。报被基热区素度。满育速度然观知得。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (79, '看九和布照老。', '2', '地快多划治或号开。年上再整如标界什业只。养应么位海参记。劳比物原拉参。', '1', '贾丽萍', '2022-03-12 03:17:18.242378', '', NULL, '没十改图总工只着。向入断出青这目下。证第分单。
你公律便面可事离空。下石养斗圆史员准利。金克越结走上。
交儿出位能本地给想员。容音例。知红出却。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (80, '深内青过再指较广生立。', '2', '以马美持内。海快除消。级三状级风飞设安。增看联商。', '1', '秘俊凯', '2022-03-12 03:17:18.922158', '', NULL, '圆往规年示连需。红清资世。同内办细还么六。
当变米。场回千规。必格可也。
青展题。金打定火能以选月。农华事程商气。
队里路。众山成想。阶例员外山查分。
界关织问约两根议。她就消过将便。报提生识斗离广术着。
识你路据权义代复大。间万段三式步青当际。会整们无区适相。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (81, '县发没京这。', '2', '它机查转使往先日收。算干族群听委圆利资。化除张算。指民看志级团气使。准报质众情话华头广深。', '1', '冯国兰', '2022-03-12 03:17:19.488508', '', NULL, '县者红教。维总米是感主。研省型细交。
对存动别老石列对。还证流者个面来明次并。即老机增众张。
看一阶生与展热生。会引始精。才候间中说道导。
南张响术世物装前。素物组分气。方应物增天院。
完元构口。能算矿查选于色手京。接位权。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (82, '去长给照流。', '2', '能离需他会观育。识京光。称整目院支快高现角。情队再。系根六海关。快相受装结其它。', '0', '丁梓浩', '2022-03-12 03:17:20.111117', '', NULL, '内容连。研展其变世这的。象飞人进劳治形专识。
可主基步反体。由水做。成素听温这工。
快分张见联什这。从次空使也质一式达学。料由史质华。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (83, '己信管儿原统。', '2', '许产真。一利适等但酸。', '0', '籍思佳', '2022-03-12 03:17:20.638682', '', NULL, '维口目系。美传果生化如手状。边增有根动速存。
情无车往团组的劳。叫不算子可。全见论。
交北又加都七响素。号我始。物青土拉在马标达今白。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (84, '山按半日根些造场。', '1', '打放向县半身。入体如流例。由件即劳明就众段。不到事候水。素正层他原不。被任组制能解委体起无。', '1', '阳国香', '2022-03-12 03:17:21.225882', '', NULL, '科表表没团山效花做。南识划济化音义风。期今效得示求拉。
观合部书。量号和研志。七后回以水该。
红者类实者划增适安等。名月的阶作。反它无已局此林。
集王文。入本门员己号。下起参济置金。
据教明角意百进去。每内长置存中识更感。件观建。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (85, '离量些引认种格变。', '1', '影二工业整适制导组管。它交易管按反会头。民教别的对处同区确们。', '0', '鞠浩然', '2022-03-12 03:17:21.731171', '', NULL, '数响得么数低究第指。除己求。次级入等离还不。
度节温何七增万备。拉农色济统治。数验样院将适去后个。
金选做技到电存统出。二我导着理量。根消教路。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (86, '世气集机算新。', '2', '采都放。斗近形。别交边科般始求儿化。展当记离力风。做世效。认近八且油民集造。', '0', '顾中海', '2022-03-12 03:17:22.320387', '', NULL, '建保新增。新定别应济分车切按。之农入。
家两素起多听题步传。半周候至。完合将。
海各下拉务备里维形。打进水。种资术矿影便也给。
期好图九石少权学。外志速克阶。三压王张千。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (87, '知力题。', '1', '论最金加清有却场。车状离。眼斗也写回林组总设。等十权通型七通。电铁县织先算。', '1', '仪癸霖', '2022-03-12 03:17:22.749947', '', NULL, '千根素权出并意工约容。料断治备报半种的构。月几任改决则记众。
个用须严但。快质须务飞。组表除集实此美引。
龙手参。过切论必热。代为办地者积级开信。
思月花。广能条则次。北队叫活话周号。
书但布理选。车内观记地内。天后如干必教酸华。
她特去得制。选至海时过质。写新自工。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (88, '重片美被步形铁起分级。', '2', '为铁快组。基可织。', '1', '充鑫', '2022-03-12 03:17:23.247222', '', NULL, '除道般转科都。两需取由相她。才些这济构。
老查建。例省结六日。取部至工外相看上观所。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (89, '往看精系新该便。', '1', '全矿许通也干率。中手格。县到消。和最验。', '1', '丹浩辰', '2022-03-12 03:17:23.711581', '', NULL, '发车导列型。自眼西阶资基员放清。领从原种。
据指时劳标列海时准则。界所证据什了二准完什。圆月万办。
主也过响几千制。空问很石。据取商通而量该规基东。
别实或长例程指写上置。计青连两约是走。术头日文观。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (90, '只传今非会志直联儿第。', '1', '治路头达能格。将形研。易放精边几八几看。', '1', '靖熙瑶', '2022-03-12 03:17:24.232794', '', NULL, '开办她办压状。备离走只重团律情带条。术员众般育度组看该。
联持派。院力类理元着常。般技被过机会海段。
马九民质这物天土则。身拉极专原如型用具。角上采京报专具。
石却积由会包多点长式。写因也候派边许人。过当员八根义近。
体第养部石装周儿家。育由证定回。养保总能无。
定整片验今须住。受了商必而能据示酸权。系极压美。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (91, '名品院。', '2', '然日存机红。事线还工这。断界求回。', '1', '府依诺', '2022-03-12 03:17:24.716189', '', NULL, '手候计须及由成定。风明直京系领但铁东市。家接道任月料。
要权它他。才百保验在段不候义。劳资个相片合反美。
究改得。确新技长族。厂色难内六构持有进。
法运近油以接支五地。之议生这权再作一。人周场往但十更着。
治千已易地。进世把更层系京上厂认。儿当林积立斯。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (92, '总产强想。', '2', '与斗程。光常维天。色正解。片林感前断林所列由效。或斯标反农传解学此。', '0', '阙娟', '2022-03-12 03:17:25.245367', '', NULL, '交周作据干。安集步。体行回最收南程完。
作步但。定走除放和机。一器学特话能情铁商离。
易达别山济已。层本大。参务基开却往米。
报得要海难。见界领。管农律。
气件力资整但历题。度及常。斗对权后置众。
省者包却酸照。内回属教存数众包油。天主则。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (93, '路口单快。', '1', '三工权不究按。候切很体型才次例快。铁从内由儿清教京半品。花厂角除治性金以速器。加广本集光事。路调整石农确反系。', '1', '禄治涛', '2022-03-12 03:17:25.720146', '', NULL, '子作律实法步消开技。快立须后叫志按运向。育还总改酸过度经。
思间增劳两进几件容相。权点积议林提。地局北南据小。
民六类加精至。看些她去象作族济相。传持角当据品数观影酸。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (94, '思况光技品和。', '1', '资道县素龙构调。斯关那各阶。', '1', '古波', '2022-03-12 03:17:26.255373', '', NULL, '改证示到出只铁过。亲好西步特的。结了料层新受车安。
质热需动。认支设价。结解热确究。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (95, '矿小长劳数进。', '1', '过性只因面性。克它素应容给车意。', '0', '秦诚', '2022-03-12 03:17:26.735788', '', NULL, '传就五标实。存断区想组系里。转备结看。
族治准林。家应土设头更。看圆回便县。
料人斯能难那造了记个。住热单公。海门重展在九标第。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (96, '成常连。', '2', '但一打前连利加。生口知动造音品名眼前。队省际百何厂。眼得族压与。', '1', '寻沐阳', '2022-03-12 03:17:27.286367', '', NULL, '严地所该将接各情它。元第权干管元话领层。治后干边。
任合据边值将大使示转。形人正。行太并员究备两识华。
此着论严给员选。养百酸数学验。外八时会面始收程织。
回八育史。路青派照下于列。性却得特太天是。
科法叫世层。华清千易单月原。好业市处要周圆米。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (97, '志安使等目儿。', '2', '构为将阶门什支性。教该铁。同特西约标。东书己器应会争题。', '0', '后丽芳', '2022-03-12 03:17:27.720186', '', NULL, '任全响约进。年用包始将目事题们专。度任次先按中以证究为。
展建步圆造体去。方只基拉。号率况建消没广正求。
最育法月展省周速。认段直正证化步。的在事内值别比叫交即。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (98, '或务今出。', '1', '单这将该将资。太起外华。红际调备再。京群影知青众比就众型。构众所干法京严。新温级铁。', '0', '言梓晨', '2022-03-12 03:17:28.280784', '', NULL, '际类之层领生现题化包。复离调清非始。处例增只半热效响。
太新完子切用多。还也电。须派东两南青还片些带。
可着使证际感完极历意。低还九广米有间子。权队次花。
须气始机。将亲而标。面近石。
究级条些下。具如运验。该式选多并温意。
关西话素思和十。教力权从石。步位整世管严用。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (99, '为委利器规务王回。', '2', '门究放斗示那第飞。队信信美算济成生质度。区传能局眼。选通管小七着主。程知飞两别级。导此上运事美。', '1', '萧宇航', '2022-03-12 03:17:28.725759', '', NULL, '话用几正劳厂断安选细。马完比高每。约精式大际。
非群带表之联回表把很。精等观局教大门新。被取子规。
变次办头也。向圆不从史对它。位强专动看。
月现气。酸场已不管也会。克感清总。
商样义属。便边积派般转论再没满。大它之值通值线划员通。
满图身般据管回照。也人准热。青该素元万。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (101, '清须分派从容声后。', '1', '教队始细群身。条己能计把经结起划成。参加名际斗较龙花标布。效法斯。', '0', '所军', '2022-03-12 03:17:29.749933', '', NULL, '七斗接。题况用好。极当属。
面查组。业厂般领圆。后京口三他原子题着争。
为研时南。相于养特时律根几格发。就世内应便她业该。
动且当化。三法众员利且。思育张儿太准。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (102, '事作论。', '2', '商众五山研识外于人。技结最准处确干分。', '0', '陈癸霖', '2022-03-12 03:17:30.284442', '', NULL, '很决造值。着林象备影组。消温切合。
与复无儿包。内都太。正红段清说识不。
给果外。受值指面义需数果术用。细细般油干业然器。
从应再。据都委必好种回群。计给变写便。
节此因始张很边化并内。性断叫。东组机一志先片亲年。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (103, '火色技强。', '2', '院片传史等片气。米难任她。验光自后等处格。', '1', '包治文', '2022-03-12 03:17:30.681605', '', NULL, '车即满即七众作反年保。号族人济清先内就候。米需什性同会。
们斯边这须石书理米。受系然只。全所林感。
知部何价。回己矿这以张。更产领全证作流我全节。
因声做。律维长复题。温华制识物实思人写。
算合验改争向。快电列车发因电。在活段八近属型龙教。
约斯后识住学料圆响样。的十眼还造我今。想五通。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (104, '种她反事林天越。', '2', '高并备边且期义于少非。你往造快影。西十加了温着备完感。命例律算示。她矿命类群记线。', '0', '车语汐', '2022-03-12 03:17:31.322657', '', NULL, '界们照子组。与始些门设电济相。接十只位人东置电统参。
收先八该何。儿被是无它除过。备动支装产却流么圆。
百因文号当于。点及存为支要十。加消公。
传一局更工。知变历用不确认八达。育去率事步术太阶。
和领过等积和。来次无展儿它类。美展飞圆铁信。
酸九算其同品本型等。上口活该干华两车示片。切市适事划亲。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (105, '行阶主。', '1', '之数解次他。证约六总。', '0', '占奕辰', '2022-03-12 03:17:31.777356', '', NULL, '用受求角。程布利候管科务。果状程极等与价交其导。
发照论经场基。打象运划取子。结月际。
比划干联。世放查需往。治众成到向完。
化作应这京众因对等际。此二合群切气会西。酸电别经。
争说极始。命来会清当除条。矿各点造。
意通知精报于值间。器内规入思天面制关题。书确切常民儿。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (106, '别气称部参这传。', '2', '外子然米全表石是科。思也周。', '0', '靳宇泽', '2022-03-12 03:17:32.349299', '', NULL, '然资决记起组严确八也。去便学适知点精。千天条需。
题水飞计。及收出别酸。权资存至光质查斗安听。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (107, '想志员治作究交斯。', '1', '分口土所。派问己严织被西开。把家无。', '0', '进梓浩', '2022-03-12 03:17:32.923719', '', NULL, '是目置例量许变。过话热始别其能。层行着给引系解派置图。
确任车过技么文省受京。三易是达导需深部些。民说动。
引低八。传较业布使。水数容总记。
日况研美。省何一经各人厂。改立带边些何能常这。
马带子场。资低便属头。基现经片手。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (108, '很备结第家例接想何。', '2', '种具便前。又备术得进近标原并明。十构式任被造结。感精选眼。斯合层数具育此别事。', '0', '睦思佳', '2022-03-12 03:17:33.628349', '', NULL, '完上办向中然。于西资群管山层二热电。边场难再被每际。
状业节斯八。引先团后。量斗如。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (110, '率近今华。', '1', '始花很。物往或许。信争连记局出要农。是且低类已周资压还。', '0', '刘建国', '2022-03-12 03:17:34.650489', '', NULL, '专民他年历集间。提一最八几你变。回联标快本又约立。
日间五提石目级目区。已为决金算。数然二省再才通更成同。
当做马术权决置关位华。响和得山。到克称治打地较值。
北文又将位现达口。知始高。高人被决商王队族层。
些议白美世办种会光又。候角积放记研。行目前由照。
传打压学求之通气二建。王存事场说。阶民离。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (111, '回与没团。', '2', '接如重造在离。还去率即们圆群员。表指百领般。而行形学数劳时价着。', '0', '隆家明', '2022-03-12 03:17:35.167304', '', NULL, '非容立月我大。圆高拉斯法。复义消机事。
气据金两局回内反。也按家调支名小新派正。性外着。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (112, '根米接些达回。', '1', '导史员八。节角果张县装心。直且又强务。事采例计动受。导书县。到收造满段儿般意。', '0', '姬乙萍', '2022-03-12 03:17:35.639702', '', NULL, '制类团流响论因身路风。育格地。系是一业真西确三。
类近色连会则论目始。入着存化增候人活。林和报土。
千压整。道即局。种变回般形。
市特常就元型能周有第。分类手人。风根行道次。
组标备半者。证白克养。领每厂设造精。
主指精特。因解常。放中格上示集信然比。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (113, '效县得持断克只每可。', '1', '打报农前质该加结海。才志养场林细近头品等。及名家过打支间派边。', '1', '井超', '2022-03-12 03:17:36.272832', '', NULL, '消置向多提。部证化还科。样去气。
才看系广面其专需论地。酸华位路。年场史见书表矿还。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (114, '记期思目会。', '1', '再来始队该拉角。白联需格构有日。', '0', '绳民', '2022-03-12 03:17:36.840257', '', NULL, '由六定质件。争油选格严位千。始步效。
究热级程性。之机花青理前七连。变行出。
老济养类能型命区转。进需思他自。音极传。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (115, '效技状走步劳新。', '2', '原带或权可该对新研目。广电义角专明是气利京。光算完种与。', '0', '邴国荣', '2022-03-12 03:17:37.311764', '', NULL, '极团者很铁法总能。但次议制根相规与。还千安今与许往设。
备大过。表入如。争第易色但发队。
合酸型反机格消铁高术。铁复入几状。养格其。
万被表济要铁听确团难。节包提八马史连门品原。入学写。
张火活边解标知持近分。质知同节法科车现。持器响运步信。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (116, '明如部书周所般县斯干。', '2', '保进是小往大志存常。导信多离。形油多支后此论快求。', '0', '佘苡沫', '2022-03-12 03:17:37.796468', '', NULL, '拉量光转想按对千清。打方易期带派再更代。今比住马。
例支到样号向。标关再动二。南高非情个。
于反除明结置技重构。发号放也别当之半。族自无按收便复问她较。
各美由明上安青参。京方院价家们。问保身场以响次长快用。
设万至万铁。识我究属象深看花必们。性置公号就就派新。
品较八志现产样般题。据书办要立日日面又资。清原商叫使究以听。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (117, '得标海你小。', '2', '者把立中情。知组查收维。民消飞入准合么海应。情被领强比些给。集表龙。', '0', '礼茗泽', '2022-03-12 03:17:38.403948', '', NULL, '布人次解将。开整议面量求部克已。二究线位。
经容大专层角对。准率改飞手其马要。组酸五标无青务引属。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (118, '越具个。', '1', '北县海因造南水合。效权许原。东所些。和快全属列族将才究。已极所及除自真样。', '0', '范姜癸霖', '2022-03-12 03:17:38.874942', '', NULL, '要及七。很群南单管派说必该子。技商系每。
采等自进但。手低什然过。大整消。
好车条信。事连象根数矿置。查议必始度变体装。
情头起。段容构周第反机。的安连应日都对。
便活铁技三。样装里题精方。才影后边都性阶治候。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (119, '表加约百图。', '1', '产称还十象。及路济本况电求基。须般重南数。手构北好。组争研件规这华车通。越斯阶权。', '0', '候浩', '2022-03-12 03:17:39.319879', '', NULL, '起干元型影变周低问。带约导划专也程较半可。始规人布更数许级引。
变热几号。深例斯度往与类称科其。且求委。
局团手育名没由。林利构组界长且上达术。活圆济。
求压持期引用别及。阶七道。计意六院个。
非记济常群治。六北几。数真住话认效料看候象。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (120, '本增眼算些我。', '1', '热或行已也连。往动了在正前流生格。调料族义族二品做。', '0', '镜玉英', '2022-03-12 03:17:39.803748', '', NULL, '做展应太积示。例准叫局有下铁精技山。业始易无角。
安业反精。条空何电拉机正调。后少件热段西。
都年内取。例难完局百打。性用却。
太类志极般快住。要单上满治历照红作。成头气同里。
百电行志示。定史会构家干加信片金。只精华间他。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (121, '记只则多西一圆报史情。', '1', '能好几院飞他切美转出。重干商教满五大听。非我名技制。正极素受理速。多整和先。', '1', '蹉娟', '2022-03-12 03:17:40.257746', '', NULL, '般华于传。价解极。强月只走管者比亲油。
便程眼自斯。想被参事适也。同少究带。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (122, '节听边且题火断亲相。', '1', '管具就拉断通手使明张。理来矿群观向况儿器。代反代多建区济点局。即引七办红。长克它把各将大际温。线性第。', '0', '有梓诚', '2022-03-12 03:17:40.782584', '', NULL, '受支构。存每照类场联。育什听达元体。
不边第十。又县周厂列除能联。用实养法实出置细。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (123, '我属布活。', '2', '美几工到地。东总斯现数。于容增。', '0', '纪浩轩', '2022-03-12 03:17:41.338051', '', NULL, '又题个按型候事地量进。论声样并。知何八切内义。
果生自。复料打之。区手受海由先。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (124, '式程素。', '1', '性再观高体业活八。取口规带马。也数心也劳。很段报。目拉别料边即。', '1', '桐国芳', '2022-03-12 03:17:41.689956', '', NULL, '成工上出装选开去。也县少速低些又很教争。或六科越型直管团。
用象什往果热布。出教石更收见起。省场难与。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (125, '完更高查许证问记技。', '2', '管边量调量情。通成支两院查型。规展风变。', '1', '厉宇', '2022-03-12 03:17:42.387084', '', NULL, '须加拉产回。满断作为但。处格明应何低去多百林。
子布及流果半。家我直果信自。华阶带据领则出么应。
经省很十金。日世没。器见眼结维。
打非还。更选任则最。化回约。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (126, '林建也可再包。', '2', '风外层新红。层和以直叫越民情位。三山的接还管文单。响公开。日明须名商取元同步。影候面给务照育观。', '1', '九奕辰', '2022-03-12 03:17:42.892034', '', NULL, '取收安理。局越经到实。维花区真标又专。
光万物华。老打利每。直海除感指等。
三图教眼子关三复按。值调议同万影的。物元龙正这适合。
支题命南前酸组图。达于会拉原向。话始自马很有开。
农应阶会离它。太行马持候红新。立至较们步公据各发。
改内体圆计验农给天定。周质间近。话相常几确现三全将教。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (127, '与需况计样。', '1', '一率听类信。内么其达。', '1', '秦雪', '2022-03-12 03:17:44.562273', '', NULL, '九史料千任应容。五安许深流。具明很王的铁个。
带里声见入。完正在口圆石术位。红集造据效观话。
了保力。三门真量价形增。识用西青条。
济里史适金系断名。成子还众统则算体更。记性日斗高相。
传专书状又。完周王据议。回正构方劳。
较研商片条少断。运算取都北。光十样易高质连内。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (128, '口体构劳果运查。', '1', '受维许写则收满却。算每也具派由。家与只际建务议者。状名切名。', '1', '尤依诺', '2022-03-12 03:17:45.266759', '', NULL, '律候局情响除质。值是将期状红省可。也原经分。
天现下。统空民术关比支全。知它近金值气北备料不。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (129, '起变如物如第发观。', '2', '象断学始你明么。儿则地五由果。王目主用号铁级利上。', '0', '蓬一全', '2022-03-12 03:17:45.89464', '', NULL, '便低置管目种张。那影电改照。你导标没。
增新命任。马复土无可织备算根极。物次走加制。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (130, '据满飞团素北理。', '2', '做展问六之动节直算。出日正准。场年速金色活反到参该。', '0', '芒智杰', '2022-03-12 03:17:46.597622', '', NULL, '问验办段按矿清。精两中但近及己场成话。基你管年常商收养必。
教前西照由。热物整。话例示压听定也团。
市边选。开团变种。治采九七。
无速引给写例外反至率。都分中看形传来根非系。火才五的农。
京思专进合代深分严。名我着正公。会造月厂。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (131, '学阶安样眼图满通效别。', '2', '我较开再门联空月华。照表号起切。基件眼打。采走间统石进。', '1', '锺俊熙', '2022-03-12 03:17:47.259983', '', NULL, '太斗部引格。律西管候命新还装还。口计来面。
做场场计参。后持今权。直图就成细。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (132, '备花发放正眼候员性。', '2', '达十车下准直。克办改。与真务二消。者还其斯中时见除极里。知系展每与飞。在易求形。', '1', '於晨', '2022-03-12 03:17:47.889222', '', NULL, '么那广计适解量色。种列支自带飞需需物造。活队群西后。
米日农。数所民族省。易标热车自学还化老认。
花压想数事广上发取分。志等马。例王音际科何整求记。
量业道没观设长在间。术转却。张度时难达眼求作。
维很去料。后科专。最便对给期。
没调权对住青包。从观些改当只走。步领才上温。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (133, '群中证众的因题。', '2', '养研出件口使次整二育。话教分点意性。', '0', '尧文韬', '2022-03-12 03:17:48.50604', '', NULL, '五白与半展。法体响局下工亲代路管。计由复量最热大。
着加次。见县标经这图验。要成成元思设石本。
间她对角。着点其教后员约又养三。心非业北万火准信元。
量子去规。广铁却家般时。转离张前支八争间见声。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (134, '六片给支新色白与经果。', '1', '门开才收。着建与这。联克行易争。', '0', '慎国兰', '2022-03-12 03:17:49.226966', '', NULL, '走空细术受示办从都。合使南定认被。多半布院声区。
养无能带制各低同带。命难它事调带号人器。思克看起指真细。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (135, '元各正什子。', '2', '观空斗采研济论。者亲研以动称由照细所。我即院证收正别料。加时声。', '1', '乾秀英', '2022-03-12 03:17:50.326983', '', NULL, '系住术行。定两强下观委千阶复。将见系外青。
条单象队产快。满非些济意多应。商认开效识县。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (152, '持正术。', '1', '族世天基。该值切边将。就把及争角。米维性明每持存前金。况门志色原或制报。', '1', '衣政君', '2022-03-12 03:18:29.28855', '', NULL, '通效全相律查则验。老须学千位都。类听此史者。
起传百表文长织。包矿然能除划运。也争活。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (136, '越建重矿。', '2', '音体运强。先到由数属率力。种了达。按真持走成为。', '0', '琦语汐', '2022-03-12 03:17:51.213393', '', NULL, '七值商下完。些观小并问电对不。八消对构更力导存华指。
流结入专报示。组成约例较而法界。素易条确万较际置但。
阶难事满级连。学油叫给者导极办。都五越下装身料当格取。
连转资标如为内受。最色或。究好着法儿调厂收族速。
多二名八具义。劳接满。花没界务团。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (137, '真总周什较部华务上。', '1', '记开解王九海织世。新知油九。常总命先效工往素。规九省。团劳时传分易立。', '1', '励鑫', '2022-03-12 03:17:52.215095', '', NULL, '件百图。日节比并。子律细员青众形支深。
传导即身金调今周达。最们资且信己局斗这。级量之几效基近斗保提。
参空信效它。我素或命起意改流现。体算次但界色造。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (138, '情除下构里表万今此称。', '1', '技话物上眼已少统究。身界导系验本族。空么命。长家术议九花行才委。必商出资治段。', '0', '屈治文', '2022-03-12 03:17:53.041137', '', NULL, '料门使强。科验往。工品地将知越住什会但。
按制通况。系解易儿。运公完料体龙其号争场。
严思情。速所酸。感利光复。
天由县理先之出动。有认个。机率听市。
务构受口同音目区议道。划持华算除。农确几青研决通周才。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (139, '土她民该组常。', '1', '集入情无飞化积场容。拉级存海必步水身。光心接素。', '0', '粟军', '2022-03-12 03:18:15.01562', '', NULL, '说张我我极口。张当白。许意也低什华白性而。
易少即且铁八价小。受效却内任育你专色义。看带体带历般历。
记节她里识代。率已认它学报整。明员还快引去除。
克各适。何型两即市。得用划队风六达。
人她转。铁红同。任思学中不精多方土件。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (140, '调信报始别。', '1', '事非西导县况完并已。其拉题她周活开领从路。众写转省由最万。往立影。多取连个细离按度。目好法口切住。', '1', '官国珍', '2022-03-12 03:18:15.721373', '', NULL, '增七即当但山果真门。关石说还由果有速所。群回查油加消。
二路维适。心复报照无知儿保。交特从信装运当提。
月海际期算基会。技品观。直见委越很精标。
行花六发。系色油米据。收量效半始民备相片。
术量状能无去。书线委信精须。较何第。
置将全该。压转走期委引历示南现。省当天业满的联领它。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (141, '马动清律而。', '2', '查两具列市。里制叫满品争此将利。它适期器以圆入种。大合过间列治始里周周。业飞于节积增原了。积经之。', '0', '刘红', '2022-03-12 03:18:16.414084', '', NULL, '济消交万般事。并没关周。部达重解么图形。
级程百。持查转前非说南。料能际打认属许技线装。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (142, '响文率总声关中查必克。', '2', '你眼和儿。制论然如济装论亲非。调必记用二常引。', '0', '烟鹏', '2022-03-12 03:18:17.458088', '', NULL, '调长总除。农持示题最法率九强。其至期色照等后老工火。
许很须生时当前转质口。极能广太。石头部议心位也。
院圆全。东统点打南京期过才张。象名阶只美。
已领精满。联出划容根。多往说应。
专两感正什九由报和。道个白。难正层得示适件样。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (143, '深后半老家西状物眼书。', '2', '反期专面除元风件学。团加全管验意林度要被。本内问不生展酸进期。又到应。方器无。科关界白片期。', '1', '朱欣怡', '2022-03-12 03:18:18.525871', '', NULL, '得于月电。感交生快京万可历价。厂易种。
土却易长教制适团响。效由细老划自他也出。那以阶市向历林样。
民和五县史上色九产。计生把称转提本史志。式业县非又做马厂。
还结金真。指许离么手山构。青下它管治示力年。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (144, '无特连生酸意。', '1', '劳三品养众十民专。观京般转放活么。', '0', '丁艺涵', '2022-03-12 03:18:20.314713', '', NULL, '石张石除采。上流立类层二了。节之头数县工取。
安切听收和节管育。适带算养些京资用明。边风东。
特结长济适同法严包。增程率数设又完场层。目至米下强造得但当真。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (145, '马因期太。', '1', '我道运下委眼采门。厂整表布适。百议位。', '0', '零玉珍', '2022-03-12 03:18:21.103747', '', NULL, '年能放值几种好压。话素们百。火东历。
通发示美值教。么从实理究或十儿原候。事保专容量际。
大干体己治看。白期间速论老南。增入接现期。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (146, '半引人约。', '2', '种气克求。很给难音众进八。复主由素其容界转局织。正状其名千花动在收部。确反山行实成。化种次西研基能或现。', '0', '森一诺', '2022-03-12 03:18:21.964748', '', NULL, '华对区被问感准向听。边增特。世局认之度动家。
第外可。影方生院六质更变今力。周数海红打京安须决看。
交性身存求效。越半合何。回际历场合从。
把作层近教。点元求写都万质带且清。在半历建真列行。
持正日亲比。场程极。查产派整需决热三。
办思听处计。备报除林。面事众果却样单效由。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (147, '候她把式单史米头完。', '2', '议提口场低做专及思事。想月性安最。规小保正局响入。离管并住。', '0', '狂志明', '2022-03-12 03:18:24.539215', '', NULL, '术可半细片型装关造此。干斯月都。造消线。
者生些层组。志设已。看基最么。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (148, '被维气空制。', '1', '采法积必。么程运高局。写被厂知为。就着斗人表外叫专。由单组完合东。称世节产常。', '1', '莘梓豪', '2022-03-12 03:18:25.860762', '', NULL, '结热矿力表。很海好认活实太反。单照级矿定交过。
北矿知时。局极但群样。少复受装时引。
取前听况身次正近。入装存斯。此通安没价多。
地实些务加元济才。便积时品存节难。关术将压酸生。
导格着。总张步己事。划必两清争运。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (149, '型金车应情经物众得头。', '1', '实何品被务形分压少。太意的算值入保运日运。些着全样写线。值节全。精你长并长。', '0', '沈超栋', '2022-03-12 03:18:26.68397', '', NULL, '即见思意身可小。片技才。铁也要标统族地什西验。
道治西图点系空并成使。人中争或或又六面织。行把火空。
具拉东众斯听内运。看他放民省。认史力同完后历县真律。
被才效群三计层金数。求由而一界千育总图团。许西更转场体利。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (150, '任王则马点火。', '2', '史电快识群东己议。值自水家速员率路段立。', '1', '姚文韬', '2022-03-12 03:18:27.494064', '', NULL, '元增红热情群酸体命。什气场。示生运相极了热立观。
不并是反。并改温行细太解再率文。住群这前话。
后持百近素干断门。持八织知例准头火大。节给整好并。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (171, '思设石产先。', '1', '给识派。提速见一集交照料看。状术龙理参。', '0', '乔沐宸', '2022-03-12 13:24:04.394533', '本诗雨', '2022-03-12 13:27:43.815572', '经声质般向后非动设八。意亲义该头许山。根快系京速。
万区知的二市次厂。整形心处和集响流。半京调层统是从。
青打结约开完参理。劳要关油。常车再族被界局温是。
备感度向维。主划地布各真水。类被有声动类。
回再制入天事学质交。但文走连真研包参法。查去无存级。
观或据百。义验管名间际。音九着生干算相车间也。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (172, '众小山。', '2', '程加以号影育定界。听领多价。证界车提王加性标常。海史切争造低将众。律细建身置克造意商往。', '0', '', NULL, '剧沐辰', '2022-03-12 13:59:53.002882', '图无油海知人并西快基。资列期。据日品命人属。
年而号天即理统车物些。除信且制受各今按。太回么段称林说决个并。
月包面叫过红京。离往重风各员而色。技一济导。
个了越教计总。委发产根群员合无维精。系马响。
系感流。度活条这华说。个几了立西厂斯积还非。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (109, '资可小。', '2', '效段候管转心交果京。具日做常可火后查多接。按意特。半实商其要。非总非增位整度。拉的安难积断资对。', '0', '闵国良', '2022-03-12 03:17:34', '', '2022-04-17 21:54:02.082218', '团别例验斯该。明认非。在通气取道们。
亲型领般队北就经身计。美进属及织提离强在开。听信率之候证叫。
对无果求干况。队与验。满研亲过据干解京。
全离所间。地知则住参。委西权广总业难区成连。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (60, '至见王声。', '2', '改调价选和资定。则众并起治程。', '1', '沙馨羽', '2022-03-12 02:19:24', '', '2022-04-17 21:55:31.559925', '点代比资去么段论度省。观自新应真都特。备没对快复研县称张。
工过器专数。然张观美构。领清象干状北。
持放质时两律品长更三。书实划生可。族于安构家起出半。
断事其命然深么。去选始下满再要此斗正。合格铁眼群。
基该提关权才将清见华。务名示民始和。最工例近手属。
员种则容文值入个。统之效义效研以打。全石专温严水走。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (348, '2222', '1', '<p><br></p>', '0', '', '2022-04-17 22:42:17', '', '2022-04-18 01:03:57.705945', '');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (174, '象总己。_123', '1', '<h1>切物历院。组北众。交众他想近意例然离。</h1>', '1', '望义轩', '2022-03-12 14:03:45', '', '2022-04-18 01:05:20.982347', '打领三处因处专。越分信走易定感题引我。这种状声物更情级必个。
之安放才始说想产。做领义力一南。种能规京什全只。
复严题造花族。年图思。着适火风由位业影形众。
业低车备身。意组集采。消约放。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (173, '知及路则。', '2', '<p>111111</p>', '0', '习子豪', '2022-03-12 14:00:58', '葛玉珍', '2022-04-18 01:00:39.660865', '系细再术也且例度同。就内采。全参王角。
两质号满教当算已常头。容力史海资意。只至几应成地。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (151, '半交程历第表。', '1', '<p>123</p>', '1', '迟熙成', '2022-03-12 03:18:28', '', '2022-04-18 01:00:55.989663', '眼可段农。队类研周线证观些。议技名族究铁写程实。
满相备素。正音到。比式置支对表标空把复。
进新走叫办难向有经整。每群农书斯己知。接龙万化本将。
号每人龙存值东土。边情到例重。何长油连将理起单情形。
接度山。往近按石半因县受以。级维两身交。
科可便究。集性王确才达对。马行却事集事节只并电。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (175, '越万则。_123', '2', '<p>际细织万。七器积人看学机什金准。象增华低非。花准始。</p>', '1', '蒙梓豪', '2022-03-12 14:03:45', '', '2022-04-18 01:07:23.81769', '主名组土基分面结。东布对么可本。和争开再阶。
路争被对。群管但能建持计断省局。者美样通展中市构根作。
与县油门华土气决么。家研取质低现先。可任温被。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (347, '222', '2', '<p>11111111</p>', '0', '', '2022-04-17 21:50:57', '', '2022-04-18 04:15:06.889248', '');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (153, '还第发切段复规极。', '2', '<p><br></p>', '0', '费斌', '2022-03-12 12:39:32', '', '2022-04-18 04:15:26.460897', '十则路。则新内因解。分还土受天在利适第。
现示八太与易消。先道收律。理经论是上直。');
INSERT INTO "public"."sys_notice" ("id", "notice_title", "notice_type", "notice_content", "notice_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (349, '555555', '1', '<p>555555</p><p><img src="data:image/png;base64,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"></p>', '0', '', '2022-04-18 01:30:16', '', '2022-04-18 04:16:51.136386', '');
COMMIT;

-- ----------------------------
-- Table structure for sys_operation_log
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_operation_log";
CREATE TABLE "public"."sys_operation_log" (
                                              "id" int8 NOT NULL GENERATED BY DEFAULT AS IDENTITY (
                                                  INCREMENT 1
                                                  MINVALUE  1
                                                  MAXVALUE 9223372036854775807
                                                  START 1
                                                  CACHE 1
                                                  ),
                                              "module_name" varchar(50) COLLATE "pg_catalog"."default",
                                              "business_type" int4,
                                              "method_name" varchar(100) COLLATE "pg_catalog"."default",
                                              "request_method" varchar(10) COLLATE "pg_catalog"."default",
                                              "operation_type" int4,
                                              "operation_user" varchar(50) COLLATE "pg_catalog"."default",
                                              "dept_name" varchar(50) COLLATE "pg_catalog"."default",
                                              "request_url" varchar(255) COLLATE "pg_catalog"."default",
                                              "request_ip" varchar(128) COLLATE "pg_catalog"."default",
                                              "request_location" varchar(255) COLLATE "pg_catalog"."default",
                                              "request_param" text COLLATE "pg_catalog"."default",
                                              "response_result" text COLLATE "pg_catalog"."default",
                                              "operation_status" int4,
                                              "error_msg" text COLLATE "pg_catalog"."default",
                                              "operation_time" timestamp(6)
)
;
COMMENT ON COLUMN "public"."sys_operation_log"."id" IS '日志主键';
COMMENT ON COLUMN "public"."sys_operation_log"."module_name" IS '模块标题';
COMMENT ON COLUMN "public"."sys_operation_log"."business_type" IS '业务类型（0其它 1新增 2修改 3删除）';
COMMENT ON COLUMN "public"."sys_operation_log"."method_name" IS '方法名称';
COMMENT ON COLUMN "public"."sys_operation_log"."request_method" IS '请求方式';
COMMENT ON COLUMN "public"."sys_operation_log"."operation_type" IS '操作类别（0其它 1后台用户 2手机端用户）';
COMMENT ON COLUMN "public"."sys_operation_log"."operation_user" IS '操作人员';
COMMENT ON COLUMN "public"."sys_operation_log"."dept_name" IS '部门名称';
COMMENT ON COLUMN "public"."sys_operation_log"."request_url" IS '请求URL';
COMMENT ON COLUMN "public"."sys_operation_log"."request_ip" IS '主机地址';
COMMENT ON COLUMN "public"."sys_operation_log"."request_location" IS '操作地点';
COMMENT ON COLUMN "public"."sys_operation_log"."request_param" IS '请求参数';
COMMENT ON COLUMN "public"."sys_operation_log"."response_result" IS '返回参数';
COMMENT ON COLUMN "public"."sys_operation_log"."operation_status" IS '操作状态（0正常 1异常）';
COMMENT ON COLUMN "public"."sys_operation_log"."error_msg" IS '错误消息';
COMMENT ON COLUMN "public"."sys_operation_log"."operation_time" IS '操作时间';
COMMENT ON TABLE "public"."sys_operation_log" IS '操作日志记录';

-- ----------------------------
-- Records of sys_operation_log
-- ----------------------------
BEGIN;
INSERT INTO "public"."sys_operation_log" ("id", "module_name", "business_type", "method_name", "request_method", "operation_type", "operation_user", "dept_name", "request_url", "request_ip", "request_location", "request_param", "response_result", "operation_status", "error_msg", "operation_time") VALUES (1, '测试', 1, '收拾收拾', 'post', 2, '菜哥', '测试', '/system/dept/list', '127.0.0.1', NULL, NULL, NULL, NULL, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_post
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_post";
CREATE TABLE "public"."sys_post" (
                                     "id" int8 NOT NULL GENERATED BY DEFAULT AS IDENTITY (
                                         INCREMENT 1
                                         MINVALUE  1
                                         MAXVALUE 9223372036854775807
                                         START 14
                                         CACHE 1
                                         ),
                                     "post_code" varchar(64) COLLATE "pg_catalog"."default" NOT NULL,
                                     "post_name" varchar(50) COLLATE "pg_catalog"."default" NOT NULL,
                                     "post_sort" int4 NOT NULL,
                                     "post_status" char(1) COLLATE "pg_catalog"."default" NOT NULL,
                                     "create_by" varchar(64) COLLATE "pg_catalog"."default",
                                     "create_time" timestamp(6),
                                     "update_by" varchar(64) COLLATE "pg_catalog"."default",
                                     "update_time" timestamp(6),
                                     "remarks" varchar(500) COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "public"."sys_post"."id" IS '岗位ID';
COMMENT ON COLUMN "public"."sys_post"."post_code" IS '岗位编码';
COMMENT ON COLUMN "public"."sys_post"."post_name" IS '岗位名称';
COMMENT ON COLUMN "public"."sys_post"."post_sort" IS '显示顺序';
COMMENT ON COLUMN "public"."sys_post"."post_status" IS '状态（0正常 1停用）';
COMMENT ON COLUMN "public"."sys_post"."create_by" IS '创建者';
COMMENT ON COLUMN "public"."sys_post"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."sys_post"."update_by" IS '更新者';
COMMENT ON COLUMN "public"."sys_post"."update_time" IS '更新时间';
COMMENT ON COLUMN "public"."sys_post"."remarks" IS '备注';
COMMENT ON TABLE "public"."sys_post" IS '岗位信息表';

-- ----------------------------
-- Records of sys_post
-- ----------------------------
BEGIN;
INSERT INTO "public"."sys_post" ("id", "post_code", "post_name", "post_sort", "post_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (2, 'se', '项目经理', 2, '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_post" ("id", "post_code", "post_name", "post_sort", "post_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (3, 'hr', '人力资源', 3, '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_post" ("id", "post_code", "post_name", "post_sort", "post_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (4, 'user', '普通员工', 4, '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '');
INSERT INTO "public"."sys_post" ("id", "post_code", "post_name", "post_sort", "post_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (5, '434', '大大', 6, '1', 'admin', '2022-05-20 02:27:32', 'admin', '2022-05-20 02:27:32', '184884844848');
INSERT INTO "public"."sys_post" ("id", "post_code", "post_name", "post_sort", "post_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (6, 'cs', '测试相关', 7, '0', 'admin', '2022-05-20 02:27:32', 'admin', '2022-05-20 02:27:32', '测试');
INSERT INTO "public"."sys_post" ("id", "post_code", "post_name", "post_sort", "post_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (7, 'yw', '网络运维', 8, '0', 'admin', '2022-05-20 02:27:32', 'admin', '2022-05-20 02:27:32', 'dsaaa');
INSERT INTO "public"."sys_post" ("id", "post_code", "post_name", "post_sort", "post_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (8, 'cp', '产品经理', 9, '0', 'admin', '2022-05-20 02:27:32', 'admin', '2022-05-20 02:27:32', '的萨达第三大啊啊撒倒萨大');
INSERT INTO "public"."sys_post" ("id", "post_code", "post_name", "post_sort", "post_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (9, 'sj', 'UI与设计', 10, '0', 'admin', '2022-05-20 02:27:32', 'admin', '2022-05-20 02:27:32', '1111dsaa');
INSERT INTO "public"."sys_post" ("id", "post_code", "post_name", "post_sort", "post_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (11, 'kf', '软件开发', 12, '0', 'admin', '2022-05-20 02:27:32', 'admin', '2022-05-20 02:27:32', 'dsadadaa');
INSERT INTO "public"."sys_post" ("id", "post_code", "post_name", "post_sort", "post_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (13, 'dsadadaa', 'dsadada', 14, '1', 'admin', '2022-05-20 02:27:32', 'admin', '2022-05-20 02:27:32', 'dsadadaa');
INSERT INTO "public"."sys_post" ("id", "post_code", "post_name", "post_sort", "post_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (12, 'jgs', '软件架构', 13, '0', 'admin', '2022-05-20 02:27:32', 'admin', '2022-05-20 02:27:32', '1111');
INSERT INTO "public"."sys_post" ("id", "post_code", "post_name", "post_sort", "post_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1, 'ceo', '董事长', 1, '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', 'sssssss');
INSERT INTO "public"."sys_post" ("id", "post_code", "post_name", "post_sort", "post_status", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (16, '2222', '2222', 222, '0', '', '2022-04-17 06:58:43', '', '2022-04-24 17:46:16.790668', '33333');
COMMIT;

-- ----------------------------
-- Table structure for sys_role
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_role";
CREATE TABLE "public"."sys_role" (
                                     "id" int8 NOT NULL GENERATED BY DEFAULT AS IDENTITY (
                                         INCREMENT 1
                                         MINVALUE  1
                                         MAXVALUE 9223372036854775807
                                         START 7
                                         CACHE 1
                                         ),
                                     "role_name" varchar(30) COLLATE "pg_catalog"."default" NOT NULL,
                                     "role_key" varchar(100) COLLATE "pg_catalog"."default" NOT NULL,
                                     "role_sort" int4 NOT NULL,
                                     "data_scope" char(1) COLLATE "pg_catalog"."default",
                                     "menu_check_strictly" char(1) COLLATE "pg_catalog"."default",
                                     "dept_check_strictly" char(1) COLLATE "pg_catalog"."default",
                                     "role_status" char(1) COLLATE "pg_catalog"."default" NOT NULL,
                                     "del_flag" char(1) COLLATE "pg_catalog"."default" DEFAULT 0,
                                     "create_by" varchar(64) COLLATE "pg_catalog"."default",
                                     "create_time" timestamp(6),
                                     "update_by" varchar(64) COLLATE "pg_catalog"."default",
                                     "update_time" timestamp(6),
                                     "remarks" varchar(500) COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "public"."sys_role"."id" IS '角色ID';
COMMENT ON COLUMN "public"."sys_role"."role_name" IS '角色名称';
COMMENT ON COLUMN "public"."sys_role"."role_key" IS '角色权限字符串';
COMMENT ON COLUMN "public"."sys_role"."role_sort" IS '显示顺序';
COMMENT ON COLUMN "public"."sys_role"."data_scope" IS '数据范围（1：全部数据权限 2：自定数据权限 3：本部门数据权限 4：本部门及以下数据权限）';
COMMENT ON COLUMN "public"."sys_role"."menu_check_strictly" IS '菜单树选择项是否关联显示';
COMMENT ON COLUMN "public"."sys_role"."dept_check_strictly" IS '部门树选择项是否关联显示';
COMMENT ON COLUMN "public"."sys_role"."role_status" IS '角色状态（0正常 1停用）';
COMMENT ON COLUMN "public"."sys_role"."del_flag" IS '删除标志（0代表存在 2代表删除）';
COMMENT ON COLUMN "public"."sys_role"."create_by" IS '创建者';
COMMENT ON COLUMN "public"."sys_role"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."sys_role"."update_by" IS '更新者';
COMMENT ON COLUMN "public"."sys_role"."update_time" IS '更新时间';
COMMENT ON COLUMN "public"."sys_role"."remarks" IS '备注';
COMMENT ON TABLE "public"."sys_role" IS '角色信息表';

-- ----------------------------
-- Records of sys_role
-- ----------------------------
BEGIN;
INSERT INTO "public"."sys_role" ("id", "role_name", "role_key", "role_sort", "data_scope", "menu_check_strictly", "dept_check_strictly", "role_status", "del_flag", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1, '超级管理员', 'admin', 1, '1', '1', '1', '0', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '超级管理员');
INSERT INTO "public"."sys_role" ("id", "role_name", "role_key", "role_sort", "data_scope", "menu_check_strictly", "dept_check_strictly", "role_status", "del_flag", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (2, '普通角色', 'common', 2, '2', '1', '1', '0', '0', 'admin', '2022-05-20 02:27:32', 'admin', '2022-05-20 02:27:32', '普通角色');
INSERT INTO "public"."sys_role" ("id", "role_name", "role_key", "role_sort", "data_scope", "menu_check_strictly", "dept_check_strictly", "role_status", "del_flag", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (3, '测试相关', 'test', 3, '2', '1', '1', '0', '0', 'admin', '2022-05-20 02:27:32', 'admin', '2022-05-20 02:27:32', '测试角色');
INSERT INTO "public"."sys_role" ("id", "role_name", "role_key", "role_sort", "data_scope", "menu_check_strictly", "dept_check_strictly", "role_status", "del_flag", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (4, '开发相关', 'kf', 4, '1', '1', '1', '0', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', NULL);
INSERT INTO "public"."sys_role" ("id", "role_name", "role_key", "role_sort", "data_scope", "menu_check_strictly", "dept_check_strictly", "role_status", "del_flag", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (5, '运维相关', 'yw', 5, '1', '1', '1', '0', '0', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_role_dept
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_role_dept";
CREATE TABLE "public"."sys_role_dept" (
                                          "role_id" int8 NOT NULL,
                                          "dept_id" int8 NOT NULL,
                                          "create_time" timestamp(6),
                                          "update_time" timestamp(6)
)
;
COMMENT ON COLUMN "public"."sys_role_dept"."role_id" IS '角色ID';
COMMENT ON COLUMN "public"."sys_role_dept"."dept_id" IS '部门ID';
COMMENT ON COLUMN "public"."sys_role_dept"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."sys_role_dept"."update_time" IS '修改时间';
COMMENT ON TABLE "public"."sys_role_dept" IS '角色和部门关联表';

-- ----------------------------
-- Records of sys_role_dept
-- ----------------------------
BEGIN;
INSERT INTO "public"."sys_role_dept" ("role_id", "dept_id", "create_time", "update_time") VALUES (2, 100, NULL, NULL);
INSERT INTO "public"."sys_role_dept" ("role_id", "dept_id", "create_time", "update_time") VALUES (2, 101, NULL, NULL);
INSERT INTO "public"."sys_role_dept" ("role_id", "dept_id", "create_time", "update_time") VALUES (2, 104, NULL, NULL);
INSERT INTO "public"."sys_role_dept" ("role_id", "dept_id", "create_time", "update_time") VALUES (2, 105, NULL, NULL);
INSERT INTO "public"."sys_role_dept" ("role_id", "dept_id", "create_time", "update_time") VALUES (2, 107, NULL, NULL);
INSERT INTO "public"."sys_role_dept" ("role_id", "dept_id", "create_time", "update_time") VALUES (3, 100, NULL, NULL);
INSERT INTO "public"."sys_role_dept" ("role_id", "dept_id", "create_time", "update_time") VALUES (3, 101, NULL, NULL);
INSERT INTO "public"."sys_role_dept" ("role_id", "dept_id", "create_time", "update_time") VALUES (3, 105, NULL, NULL);
INSERT INTO "public"."sys_role_dept" ("role_id", "dept_id", "create_time", "update_time") VALUES (7, 101, NULL, NULL);
INSERT INTO "public"."sys_role_dept" ("role_id", "dept_id", "create_time", "update_time") VALUES (7, 106, NULL, NULL);
INSERT INTO "public"."sys_role_dept" ("role_id", "dept_id", "create_time", "update_time") VALUES (7, 107, NULL, NULL);
INSERT INTO "public"."sys_role_dept" ("role_id", "dept_id", "create_time", "update_time") VALUES (9, 100, NULL, NULL);
INSERT INTO "public"."sys_role_dept" ("role_id", "dept_id", "create_time", "update_time") VALUES (9, 101, NULL, NULL);
INSERT INTO "public"."sys_role_dept" ("role_id", "dept_id", "create_time", "update_time") VALUES (9, 103, NULL, NULL);
INSERT INTO "public"."sys_role_dept" ("role_id", "dept_id", "create_time", "update_time") VALUES (9, 104, NULL, NULL);
INSERT INTO "public"."sys_role_dept" ("role_id", "dept_id", "create_time", "update_time") VALUES (9, 127, NULL, NULL);
INSERT INTO "public"."sys_role_dept" ("role_id", "dept_id", "create_time", "update_time") VALUES (9, 105, NULL, NULL);
INSERT INTO "public"."sys_role_dept" ("role_id", "dept_id", "create_time", "update_time") VALUES (9, 130, NULL, NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_role_menu
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_role_menu";
CREATE TABLE "public"."sys_role_menu" (
                                          "role_id" int8 NOT NULL,
                                          "menu_id" int8 NOT NULL,
                                          "create_time" timestamp(6),
                                          "update_time" timestamp(6)
)
;
COMMENT ON COLUMN "public"."sys_role_menu"."role_id" IS '角色ID';
COMMENT ON COLUMN "public"."sys_role_menu"."menu_id" IS '菜单ID';
COMMENT ON COLUMN "public"."sys_role_menu"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."sys_role_menu"."update_time" IS '修改时间';
COMMENT ON TABLE "public"."sys_role_menu" IS '角色和菜单关联表';

-- ----------------------------
-- Records of sys_role_menu
-- ----------------------------
BEGIN;
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 1, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 100, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 101, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 102, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 103, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 104, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 105, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 106, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 107, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 108, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 500, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 501, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 1001, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 1002, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 1008, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 1013, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 1017, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 1018, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 1021, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 1022, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 1023, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 1024, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 1025, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 1026, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 1027, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 1028, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 1029, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 1030, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 1031, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 1032, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 1033, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 1034, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 1035, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 1036, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 1037, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 1038, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 1039, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 1040, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 1041, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 1042, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 1043, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 1044, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (2, 1045, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (3, 3, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (3, 4, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (3, 114, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (3, 115, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (3, 116, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (3, 1055, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (3, 1056, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (3, 1057, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (3, 1058, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (3, 1059, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (3, 1060, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (4, 2, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (4, 3, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (4, 109, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (4, 110, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (4, 111, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (4, 112, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (4, 113, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (4, 114, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (4, 115, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (4, 116, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (4, 1046, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (4, 1047, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (4, 1048, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (4, 1049, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (4, 1050, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (4, 1051, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (4, 1052, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (4, 1053, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (4, 1054, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (4, 1055, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (4, 1056, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (4, 1057, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (4, 1058, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (4, 1059, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (4, 1060, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (5, 2, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (5, 3, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (5, 110, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (5, 111, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (5, 112, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (5, 113, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (5, 116, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (5, 1049, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (6, 1, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (6, 4, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (6, 107, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (6, 108, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (6, 500, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (6, 501, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (6, 1036, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (6, 1037, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (6, 1040, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (6, 1042, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (6, 1043, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (6, 1045, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (7, 2, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (7, 3, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (7, 111, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (7, 1061, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (7, 116, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (7, 4, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (11, 4, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (10, 4, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (21, 108, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (21, 500, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (21, 1040, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (21, 1041, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (21, 1042, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (21, 501, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (21, 1043, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (21, 1044, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (21, 1045, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (21, 114, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (21, 116, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (21, 1, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (21, 3, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (22, 1031, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (22, 1035, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (22, 1036, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (22, 1037, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (22, 112, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (22, 1061, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (22, 4, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (22, 1, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (22, 106, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (22, 107, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (22, 2, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (23, 4, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (12, 3, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (12, 114, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (12, 115, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (12, 1055, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (12, 1056, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (12, 1057, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (12, 1058, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (12, 1059, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (12, 1060, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (12, 116, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (12, 4, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (24, 200, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (25, 200, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (26, 4, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (26, 200, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (28, 4, NULL, NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (29, 114, '2022-04-20 00:12:05.256892', NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (29, 116, '2022-04-20 00:12:05.256915', NULL);
INSERT INTO "public"."sys_role_menu" ("role_id", "menu_id", "create_time", "update_time") VALUES (29, 3, '2022-04-20 00:12:05.256933', NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_user
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_user";
CREATE TABLE "public"."sys_user" (
                                     "id" int8 NOT NULL GENERATED BY DEFAULT AS IDENTITY (
                                         INCREMENT 1
                                         MINVALUE  1
                                         MAXVALUE 9223372036854775807
                                         START 1
                                         CACHE 1
                                         ),
                                     "dept_id" int8,
                                     "username" varchar(30) COLLATE "pg_catalog"."default" NOT NULL,
                                     "nickname" varchar(30) COLLATE "pg_catalog"."default" NOT NULL,
                                     "user_type" varchar(2) COLLATE "pg_catalog"."default",
                                     "email" varchar(50) COLLATE "pg_catalog"."default",
                                     "phone_no" varchar(11) COLLATE "pg_catalog"."default",
                                     "sex" char(1) COLLATE "pg_catalog"."default",
                                     "avatar" varchar(100) COLLATE "pg_catalog"."default",
                                     "user_pwd" varchar(100) COLLATE "pg_catalog"."default",
                                     "user_status" char(1) COLLATE "pg_catalog"."default",
                                     "del_flag" char(1) COLLATE "pg_catalog"."default" DEFAULT 0,
                                     "login_ip" varchar(128) COLLATE "pg_catalog"."default",
                                     "login_date" timestamp(6),
                                     "create_by" varchar(64) COLLATE "pg_catalog"."default",
                                     "create_time" timestamp(6),
                                     "update_by" varchar(64) COLLATE "pg_catalog"."default",
                                     "update_time" timestamp(6),
                                     "remarks" varchar(500) COLLATE "pg_catalog"."default"
)
;
COMMENT ON COLUMN "public"."sys_user"."id" IS '用户ID';
COMMENT ON COLUMN "public"."sys_user"."dept_id" IS '部门ID';
COMMENT ON COLUMN "public"."sys_user"."username" IS '用户账号';
COMMENT ON COLUMN "public"."sys_user"."nickname" IS '用户昵称';
COMMENT ON COLUMN "public"."sys_user"."user_type" IS '用户类型（00系统用户）';
COMMENT ON COLUMN "public"."sys_user"."email" IS '用户邮箱';
COMMENT ON COLUMN "public"."sys_user"."phone_no" IS '手机号码';
COMMENT ON COLUMN "public"."sys_user"."sex" IS '用户性别（0男 1女 2未知）';
COMMENT ON COLUMN "public"."sys_user"."avatar" IS '头像地址';
COMMENT ON COLUMN "public"."sys_user"."user_pwd" IS '密码';
COMMENT ON COLUMN "public"."sys_user"."user_status" IS '帐号状态（0正常 1停用）';
COMMENT ON COLUMN "public"."sys_user"."del_flag" IS '删除标志（0代表存在 2代表删除）';
COMMENT ON COLUMN "public"."sys_user"."login_ip" IS '最后登录IP';
COMMENT ON COLUMN "public"."sys_user"."login_date" IS '最后登录时间';
COMMENT ON COLUMN "public"."sys_user"."create_by" IS '创建者';
COMMENT ON COLUMN "public"."sys_user"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."sys_user"."update_by" IS '更新者';
COMMENT ON COLUMN "public"."sys_user"."update_time" IS '更新时间';
COMMENT ON COLUMN "public"."sys_user"."remarks" IS '备注';
COMMENT ON TABLE "public"."sys_user" IS '用户信息表';

-- ----------------------------
-- Records of sys_user
-- ----------------------------
BEGIN;
INSERT INTO "public"."sys_user" ("id", "dept_id", "username", "nickname", "user_type", "email", "phone_no", "sex", "avatar", "user_pwd", "user_status", "del_flag", "login_ip", "login_date", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (1, 103, 'admin', '若依', '00', 'ry@163.com', '15888888888', '1', '', '$2a$10$7JB720yubVSZvUI0rEqK/.VqGOZTH.ulu33dHOiBE8ByOhJIrdAu2', '0', '0', '127.0.0.1', '2022-05-20 02:27:32', 'admin', '2022-05-20 02:27:32', '', '2022-05-20 02:27:32', '管理员');
INSERT INTO "public"."sys_user" ("id", "dept_id", "username", "nickname", "user_type", "email", "phone_no", "sex", "avatar", "user_pwd", "user_status", "del_flag", "login_ip", "login_date", "create_by", "create_time", "update_by", "update_time", "remarks") VALUES (3, 108, 'test', 'test', '00', 'dsadada@168.com', '15555555555', '0', '/profile/avatar/2022/07/16/blob_20220716042515A002.jpeg', '$2a$10$sIcNKQqPJ5dmsO9TXKfHGOgU6.thgYIhatp.BkHd2suX/wKeRFHWy', '0', '0', '', '2022-05-20 02:27:32', 'admin', '2022-05-20 02:27:32', 'admin', '2022-05-20 02:27:32', 'test');
COMMIT;

-- ----------------------------
-- Table structure for sys_user_post
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_user_post";
CREATE TABLE "public"."sys_user_post" (
                                          "user_id" int8 NOT NULL,
                                          "post_id" int8 NOT NULL,
                                          "create_time" timestamp(6),
                                          "update_time" timestamp(6)
)
;
COMMENT ON COLUMN "public"."sys_user_post"."user_id" IS '用户ID';
COMMENT ON COLUMN "public"."sys_user_post"."post_id" IS '岗位ID';
COMMENT ON COLUMN "public"."sys_user_post"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."sys_user_post"."update_time" IS '修改时间';
COMMENT ON TABLE "public"."sys_user_post" IS '用户与岗位关联表';

-- ----------------------------
-- Records of sys_user_post
-- ----------------------------
BEGIN;
INSERT INTO "public"."sys_user_post" ("user_id", "post_id", "create_time", "update_time") VALUES (1, 1, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_user_post" ("user_id", "post_id", "create_time", "update_time") VALUES (3, 4, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_user_post" ("user_id", "post_id", "create_time", "update_time") VALUES (378533, 4, '2022-04-16 08:36:06.986748', NULL);
INSERT INTO "public"."sys_user_post" ("user_id", "post_id", "create_time", "update_time") VALUES (378533, 6, '2022-04-16 08:36:06.986766', NULL);
INSERT INTO "public"."sys_user_post" ("user_id", "post_id", "create_time", "update_time") VALUES (378533, 7, '2022-04-16 08:36:06.986785', NULL);
INSERT INTO "public"."sys_user_post" ("user_id", "post_id", "create_time", "update_time") VALUES (5867, 4, '2022-04-16 08:43:25.279734', NULL);
INSERT INTO "public"."sys_user_post" ("user_id", "post_id", "create_time", "update_time") VALUES (5867, 3, '2022-04-16 08:43:25.279755', NULL);
INSERT INTO "public"."sys_user_post" ("user_id", "post_id", "create_time", "update_time") VALUES (5867, 6, '2022-04-16 08:43:25.279772', NULL);
INSERT INTO "public"."sys_user_post" ("user_id", "post_id", "create_time", "update_time") VALUES (6, 2, '2022-04-16 10:39:45.670809', NULL);
INSERT INTO "public"."sys_user_post" ("user_id", "post_id", "create_time", "update_time") VALUES (2, 2, '2022-04-16 23:10:05.098007', NULL);
INSERT INTO "public"."sys_user_post" ("user_id", "post_id", "create_time", "update_time") VALUES (5, 4, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_user_post" ("user_id", "post_id", "create_time", "update_time") VALUES (5, 8, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_user_post" ("user_id", "post_id", "create_time", "update_time") VALUES (378534, 2, '2022-04-17 00:49:52.38638', NULL);
INSERT INTO "public"."sys_user_post" ("user_id", "post_id", "create_time", "update_time") VALUES (378535, 3, '2022-04-17 00:59:19.760777', NULL);
INSERT INTO "public"."sys_user_post" ("user_id", "post_id", "create_time", "update_time") VALUES (83, 3, '2022-04-18 05:26:44.349007', NULL);
INSERT INTO "public"."sys_user_post" ("user_id", "post_id", "create_time", "update_time") VALUES (378536, 4, '2022-04-20 00:27:40.805988', NULL);
INSERT INTO "public"."sys_user_post" ("user_id", "post_id", "create_time", "update_time") VALUES (378536, 6, '2022-04-20 00:27:40.806023', NULL);
INSERT INTO "public"."sys_user_post" ("user_id", "post_id", "create_time", "update_time") VALUES (378536, 7, '2022-04-20 00:27:40.806053', NULL);
INSERT INTO "public"."sys_user_post" ("user_id", "post_id", "create_time", "update_time") VALUES (378536, 8, '2022-04-20 00:27:40.806081', NULL);
COMMIT;

-- ----------------------------
-- Table structure for sys_user_role
-- ----------------------------
DROP TABLE IF EXISTS "public"."sys_user_role";
CREATE TABLE "public"."sys_user_role" (
                                          "user_id" int8 NOT NULL,
                                          "role_id" int8 NOT NULL,
                                          "create_time" timestamp(6),
                                          "update_time" timestamp(6)
)
;
COMMENT ON COLUMN "public"."sys_user_role"."user_id" IS '用户ID';
COMMENT ON COLUMN "public"."sys_user_role"."role_id" IS '角色ID';
COMMENT ON COLUMN "public"."sys_user_role"."create_time" IS '创建时间';
COMMENT ON COLUMN "public"."sys_user_role"."update_time" IS '修改时间';
COMMENT ON TABLE "public"."sys_user_role" IS '用户和角色关联表';

-- ----------------------------
-- Records of sys_user_role
-- ----------------------------
BEGIN;
INSERT INTO "public"."sys_user_role" ("user_id", "role_id", "create_time", "update_time") VALUES (1, 1, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_user_role" ("user_id", "role_id", "create_time", "update_time") VALUES (3, 2, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_user_role" ("user_id", "role_id", "create_time", "update_time") VALUES (5, 2, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_user_role" ("user_id", "role_id", "create_time", "update_time") VALUES (5, 6, '2022-05-20 02:27:32', '2022-05-20 02:27:32');
INSERT INTO "public"."sys_user_role" ("user_id", "role_id", "create_time", "update_time") VALUES (378533, 2, '2022-04-16 08:36:06.981668', NULL);
INSERT INTO "public"."sys_user_role" ("user_id", "role_id", "create_time", "update_time") VALUES (5867, 5, '2022-04-16 08:43:25.276444', NULL);
INSERT INTO "public"."sys_user_role" ("user_id", "role_id", "create_time", "update_time") VALUES (6, 5, '2022-04-16 10:39:45.663514', NULL);
INSERT INTO "public"."sys_user_role" ("user_id", "role_id", "create_time", "update_time") VALUES (2, 2, '2022-04-16 23:10:05.089887', NULL);
INSERT INTO "public"."sys_user_role" ("user_id", "role_id", "create_time", "update_time") VALUES (378534, 5, '2022-04-17 00:49:52.383978', NULL);
INSERT INTO "public"."sys_user_role" ("user_id", "role_id", "create_time", "update_time") VALUES (378535, 5, '2022-04-17 00:59:19.750933', NULL);
INSERT INTO "public"."sys_user_role" ("user_id", "role_id", "create_time", "update_time") VALUES (5752, 2, '2022-04-17 20:48:33.626154', NULL);
INSERT INTO "public"."sys_user_role" ("user_id", "role_id", "create_time", "update_time") VALUES (83, 5, '2022-04-18 05:26:44.33554', NULL);
INSERT INTO "public"."sys_user_role" ("user_id", "role_id", "create_time", "update_time") VALUES (378536, 2, '2022-04-20 00:27:37.850605', NULL);
INSERT INTO "public"."sys_user_role" ("user_id", "role_id", "create_time", "update_time") VALUES (378536, 5, '2022-04-20 00:27:37.850644', NULL);
INSERT INTO "public"."sys_user_role" ("user_id", "role_id", "create_time", "update_time") VALUES (378536, 7, '2022-04-20 00:27:37.850677', NULL);
INSERT INTO "public"."sys_user_role" ("user_id", "role_id", "create_time", "update_time") VALUES (378536, 6, '2022-04-20 00:27:37.850719', NULL);
COMMIT;

-- ----------------------------
-- Function structure for mask_password
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."mask_password"("pwd" text);
CREATE OR REPLACE FUNCTION "public"."mask_password"("pwd" text)
              RETURNS "pg_catalog"."text" AS $BODY$
              DECLARE
              prefix_len INT := 5;
suffix_len INT := 5;
BEGIN
    RETURN
    CASE
      WHEN LENGTH(pwd) < prefix_len + suffix_len THEN pwd  -- 短于10位不处理
      ELSE LEFT(pwd, prefix_len) ||
           REPEAT('*', GREATEST(LENGTH(pwd) - prefix_len - suffix_len, 4)) ||
           RIGHT(pwd, suffix_len)
    END;
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
COMMENT ON FUNCTION "public"."mask_password"("pwd" text) IS '密码脱敏';

-- ----------------------------
-- Function structure for mask_phone
-- ----------------------------
DROP FUNCTION IF EXISTS "public"."mask_phone"("phone" text);
CREATE OR REPLACE FUNCTION "public"."mask_phone"("phone" text)
              RETURNS "pg_catalog"."text" AS $BODY$
              DECLARE
              prefix_len INT := 3;
suffix_len INT := 4;
BEGIN
    RETURN
    CASE
      WHEN LENGTH(phone) < prefix_len + suffix_len THEN phone  -- 短于7位不处理
      ELSE LEFT(phone, prefix_len) ||
           REPEAT('*', GREATEST(LENGTH(phone) - prefix_len - suffix_len, 4)) ||
           RIGHT(phone, suffix_len)
    END;
END;
$BODY$
  LANGUAGE plpgsql VOLATILE
  COST 100;
COMMENT ON FUNCTION "public"."mask_phone"("phone" text) IS '手机号脱敏';

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."gen_table_column_column_id_seq"
OWNED BY "public"."gen_table_column"."column_id";
SELECT setval('"public"."gen_table_column_column_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."gen_table_table_id_seq"
OWNED BY "public"."gen_table"."table_id";
SELECT setval('"public"."gen_table_table_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."sys_config_config_id_seq"
OWNED BY "public"."sys_config"."config_id";
SELECT setval('"public"."sys_config_config_id_seq"', 8, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."sys_dept_dept_id_seq"
OWNED BY "public"."sys_dept"."id";
SELECT setval('"public"."sys_dept_dept_id_seq"', 15993, true);


-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."sys_dict_data_dict_code_seq"
OWNED BY "public"."sys_dict_data"."id";
SELECT setval('"public"."sys_dict_data_dict_code_seq"', 33, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."sys_dict_type_dict_id_seq"
OWNED BY "public"."sys_dict_type"."dict_id";
SELECT setval('"public"."sys_dict_type_dict_id_seq"', 30, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."sys_job_job_id_seq"
OWNED BY "public"."sys_job"."job_id";
SELECT setval('"public"."sys_job_job_id_seq"', 34, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."sys_job_log_job_log_id_seq"
OWNED BY "public"."sys_job_log"."job_log_id";
SELECT setval('"public"."sys_job_log_job_log_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."sys_login_record_info_id_seq"
OWNED BY "public"."sys_login_record"."info_id";
SELECT setval('"public"."sys_login_record_info_id_seq"', 306, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."sys_menu_menu_id_seq"
OWNED BY "public"."sys_menu"."id";
SELECT setval('"public"."sys_menu_menu_id_seq"', 43, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."sys_notice_notice_id_seq"
OWNED BY "public"."sys_notice"."id";
SELECT setval('"public"."sys_notice_notice_id_seq"', 349, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."sys_operate_log_oper_id_seq"
OWNED BY "public"."sys_operation_log"."id";
SELECT setval('"public"."sys_operate_log_oper_id_seq"', 1, false);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."sys_post_post_id_seq"
OWNED BY "public"."sys_post"."id";
SELECT setval('"public"."sys_post_post_id_seq"', 16, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."sys_role_role_id_seq"
OWNED BY "public"."sys_role"."id";
SELECT setval('"public"."sys_role_role_id_seq"', 29, true);

-- ----------------------------
-- Alter sequences owned by
-- ----------------------------
ALTER SEQUENCE "public"."sys_user_user_id_seq"
OWNED BY "public"."sys_user"."id";
SELECT setval('"public"."sys_user_user_id_seq"', 378536, true);

-- ----------------------------
-- Primary Key structure for table QRTZ_BLOB_TRIGGERS
-- ----------------------------
ALTER TABLE "public"."QRTZ_BLOB_TRIGGERS" ADD CONSTRAINT "QRTZ_BLOB_TRIGGERS_pkey" PRIMARY KEY ("sched_name", "trigger_name", "trigger_group");

-- ----------------------------
-- Primary Key structure for table QRTZ_CALENDARS
-- ----------------------------
ALTER TABLE "public"."QRTZ_CALENDARS" ADD CONSTRAINT "QRTZ_CALENDARS_pkey" PRIMARY KEY ("sched_name", "calendar_name");

-- ----------------------------
-- Primary Key structure for table QRTZ_CRON_TRIGGERS
-- ----------------------------
ALTER TABLE "public"."QRTZ_CRON_TRIGGERS" ADD CONSTRAINT "QRTZ_CRON_TRIGGERS_pkey" PRIMARY KEY ("sched_name", "trigger_name", "trigger_group");

-- ----------------------------
-- Primary Key structure for table QRTZ_FIRED_TRIGGERS
-- ----------------------------
ALTER TABLE "public"."QRTZ_FIRED_TRIGGERS" ADD CONSTRAINT "QRTZ_FIRED_TRIGGERS_pkey" PRIMARY KEY ("sched_name", "entry_id");

-- ----------------------------
-- Primary Key structure for table QRTZ_JOB_DETAILS
-- ----------------------------
ALTER TABLE "public"."QRTZ_JOB_DETAILS" ADD CONSTRAINT "QRTZ_JOB_DETAILS_pkey" PRIMARY KEY ("sched_name", "job_name", "job_group");

-- ----------------------------
-- Primary Key structure for table QRTZ_LOCKS
-- ----------------------------
ALTER TABLE "public"."QRTZ_LOCKS" ADD CONSTRAINT "QRTZ_LOCKS_pkey" PRIMARY KEY ("sched_name", "lock_name");

-- ----------------------------
-- Primary Key structure for table QRTZ_PAUSED_TRIGGER_GRPS
-- ----------------------------
ALTER TABLE "public"."QRTZ_PAUSED_TRIGGER_GRPS" ADD CONSTRAINT "QRTZ_PAUSED_TRIGGER_GRPS_pkey" PRIMARY KEY ("sched_name", "trigger_group");

-- ----------------------------
-- Primary Key structure for table QRTZ_SCHEDULER_STATE
-- ----------------------------
ALTER TABLE "public"."QRTZ_SCHEDULER_STATE" ADD CONSTRAINT "QRTZ_SCHEDULER_STATE_pkey" PRIMARY KEY ("sched_name", "instance_name");

-- ----------------------------
-- Primary Key structure for table QRTZ_SIMPLE_TRIGGERS
-- ----------------------------
ALTER TABLE "public"."QRTZ_SIMPLE_TRIGGERS" ADD CONSTRAINT "QRTZ_SIMPLE_TRIGGERS_pkey" PRIMARY KEY ("sched_name", "trigger_name", "trigger_group");

-- ----------------------------
-- Primary Key structure for table QRTZ_SIMPROP_TRIGGERS
-- ----------------------------
ALTER TABLE "public"."QRTZ_SIMPROP_TRIGGERS" ADD CONSTRAINT "QRTZ_SIMPROP_TRIGGERS_pkey" PRIMARY KEY ("sched_name", "trigger_name", "trigger_group");

-- ----------------------------
-- Indexes structure for table QRTZ_TRIGGERS
-- ----------------------------
CREATE INDEX "sched_name" ON "public"."QRTZ_TRIGGERS" USING btree (
    "sched_name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
    "job_name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST,
    "job_group" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Primary Key structure for table QRTZ_TRIGGERS
-- ----------------------------
ALTER TABLE "public"."QRTZ_TRIGGERS" ADD CONSTRAINT "QRTZ_TRIGGERS_pkey" PRIMARY KEY ("sched_name", "trigger_name", "trigger_group");

-- ----------------------------
-- Auto increment value for gen_table
-- ----------------------------
SELECT setval('"public"."gen_table_table_id_seq"', 1, false);

-- ----------------------------
-- Indexes structure for table gen_table
-- ----------------------------
CREATE UNIQUE INDEX "table_name_unique_index" ON "public"."gen_table" USING btree (
    "table_name" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );
COMMENT ON INDEX "public"."table_name_unique_index" IS '表名唯一索引';

-- ----------------------------
-- Primary Key structure for table gen_table
-- ----------------------------
ALTER TABLE "public"."gen_table" ADD CONSTRAINT "gen_table_pkey" PRIMARY KEY ("table_id");

-- ----------------------------
-- Auto increment value for gen_table_column
-- ----------------------------
SELECT setval('"public"."gen_table_column_column_id_seq"', 1, false);

-- ----------------------------
-- Primary Key structure for table gen_table_column
-- ----------------------------
ALTER TABLE "public"."gen_table_column" ADD CONSTRAINT "gen_table_column_pkey" PRIMARY KEY ("column_id");

-- ----------------------------
-- Auto increment value for sys_config
-- ----------------------------
SELECT setval('"public"."sys_config_config_id_seq"', 8, true);

-- ----------------------------
-- Primary Key structure for table sys_config
-- ----------------------------
ALTER TABLE "public"."sys_config" ADD CONSTRAINT "sys_config_pkey" PRIMARY KEY ("config_id");

-- ----------------------------
-- Auto increment value for sys_dept
-- ----------------------------
SELECT setval('"public"."sys_dept_dept_id_seq"', 15993, true);

-- ----------------------------
-- Primary Key structure for table sys_dept
-- ----------------------------
ALTER TABLE "public"."sys_dept" ADD CONSTRAINT "sys_dept_pkey" PRIMARY KEY ("id");


-- ----------------------------
-- Auto increment value for sys_dict_data
-- ----------------------------
SELECT setval('"public"."sys_dict_data_dict_code_seq"', 33, true);

-- ----------------------------
-- Primary Key structure for table sys_dict_data
-- ----------------------------
ALTER TABLE "public"."sys_dict_data" ADD CONSTRAINT "sys_dict_data_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Auto increment value for sys_dict_type
-- ----------------------------
SELECT setval('"public"."sys_dict_type_dict_id_seq"', 30, true);

-- ----------------------------
-- Indexes structure for table sys_dict_type
-- ----------------------------
CREATE UNIQUE INDEX "dict_type" ON "public"."sys_dict_type" USING btree (
    "dict_type" COLLATE "pg_catalog"."default" "pg_catalog"."text_ops" ASC NULLS LAST
    );

-- ----------------------------
-- Primary Key structure for table sys_dict_type
-- ----------------------------
ALTER TABLE "public"."sys_dict_type" ADD CONSTRAINT "sys_dict_type_pkey" PRIMARY KEY ("dict_id");

-- ----------------------------
-- Auto increment value for sys_job
-- ----------------------------
SELECT setval('"public"."sys_job_job_id_seq"', 34, true);

-- ----------------------------
-- Primary Key structure for table sys_job
-- ----------------------------
ALTER TABLE "public"."sys_job" ADD CONSTRAINT "sys_job_pkey" PRIMARY KEY ("job_id", "job_name", "job_group");

-- ----------------------------
-- Auto increment value for sys_job_log
-- ----------------------------
SELECT setval('"public"."sys_job_log_job_log_id_seq"', 1, false);

-- ----------------------------
-- Primary Key structure for table sys_job_log
-- ----------------------------
ALTER TABLE "public"."sys_job_log" ADD CONSTRAINT "sys_job_log_pkey" PRIMARY KEY ("job_log_id");

-- ----------------------------
-- Auto increment value for sys_login_record
-- ----------------------------
SELECT setval('"public"."sys_login_record_info_id_seq"', 306, true);

-- ----------------------------
-- Primary Key structure for table sys_login_record
-- ----------------------------
ALTER TABLE "public"."sys_login_record" ADD CONSTRAINT "sys_login_record_pkey" PRIMARY KEY ("info_id");

-- ----------------------------
-- Auto increment value for sys_menu
-- ----------------------------
SELECT setval('"public"."sys_menu_menu_id_seq"', 43, true);

-- ----------------------------
-- Primary Key structure for table sys_menu
-- ----------------------------
ALTER TABLE "public"."sys_menu" ADD CONSTRAINT "sys_menu_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Auto increment value for sys_notice
-- ----------------------------
SELECT setval('"public"."sys_notice_notice_id_seq"', 349, true);

-- ----------------------------
-- Primary Key structure for table sys_notice
-- ----------------------------
ALTER TABLE "public"."sys_notice" ADD CONSTRAINT "sys_notice_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Auto increment value for sys_operation_log
-- ----------------------------
SELECT setval('"public"."sys_operate_log_oper_id_seq"', 1, false);

-- ----------------------------
-- Primary Key structure for table sys_operation_log
-- ----------------------------
ALTER TABLE "public"."sys_operation_log" ADD CONSTRAINT "sys_oper_log_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Auto increment value for sys_post
-- ----------------------------
SELECT setval('"public"."sys_post_post_id_seq"', 16, true);

-- ----------------------------
-- Primary Key structure for table sys_post
-- ----------------------------
ALTER TABLE "public"."sys_post" ADD CONSTRAINT "sys_post_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Auto increment value for sys_role
-- ----------------------------
SELECT setval('"public"."sys_role_role_id_seq"', 29, true);

-- ----------------------------
-- Primary Key structure for table sys_role
-- ----------------------------
ALTER TABLE "public"."sys_role" ADD CONSTRAINT "sys_role_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table sys_role_dept
-- ----------------------------
ALTER TABLE "public"."sys_role_dept" ADD CONSTRAINT "sys_role_dept_pkey" PRIMARY KEY ("role_id", "dept_id");

-- ----------------------------
-- Primary Key structure for table sys_role_menu
-- ----------------------------
ALTER TABLE "public"."sys_role_menu" ADD CONSTRAINT "sys_role_menu_pkey" PRIMARY KEY ("role_id", "menu_id");

-- ----------------------------
-- Auto increment value for sys_user
-- ----------------------------
SELECT setval('"public"."sys_user_user_id_seq"', 378536, true);

-- ----------------------------
-- Primary Key structure for table sys_user
-- ----------------------------
ALTER TABLE "public"."sys_user" ADD CONSTRAINT "sys_user_pkey" PRIMARY KEY ("id");

-- ----------------------------
-- Primary Key structure for table sys_user_post
-- ----------------------------
ALTER TABLE "public"."sys_user_post" ADD CONSTRAINT "sys_user_post_pkey" PRIMARY KEY ("user_id", "post_id");

-- ----------------------------
-- Primary Key structure for table sys_user_role
-- ----------------------------
ALTER TABLE "public"."sys_user_role" ADD CONSTRAINT "sys_user_role_pkey" PRIMARY KEY ("user_id", "role_id");

-- ----------------------------
-- Foreign Keys structure for table QRTZ_BLOB_TRIGGERS
-- ----------------------------
ALTER TABLE "public"."QRTZ_BLOB_TRIGGERS" ADD CONSTRAINT "qrtz_blob_triggers_ibfk_1" FOREIGN KEY ("sched_name", "trigger_name", "trigger_group") REFERENCES "public"."QRTZ_TRIGGERS" ("sched_name", "trigger_name", "trigger_group") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table QRTZ_CRON_TRIGGERS
-- ----------------------------
ALTER TABLE "public"."QRTZ_CRON_TRIGGERS" ADD CONSTRAINT "qrtz_cron_triggers_ibfk_1" FOREIGN KEY ("sched_name", "trigger_name", "trigger_group") REFERENCES "public"."QRTZ_TRIGGERS" ("sched_name", "trigger_name", "trigger_group") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table QRTZ_SIMPLE_TRIGGERS
-- ----------------------------
ALTER TABLE "public"."QRTZ_SIMPLE_TRIGGERS" ADD CONSTRAINT "qrtz_simple_triggers_ibfk_1" FOREIGN KEY ("sched_name", "trigger_name", "trigger_group") REFERENCES "public"."QRTZ_TRIGGERS" ("sched_name", "trigger_name", "trigger_group") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table QRTZ_SIMPROP_TRIGGERS
-- ----------------------------
ALTER TABLE "public"."QRTZ_SIMPROP_TRIGGERS" ADD CONSTRAINT "qrtz_simprop_triggers_ibfk_1" FOREIGN KEY ("sched_name", "trigger_name", "trigger_group") REFERENCES "public"."QRTZ_TRIGGERS" ("sched_name", "trigger_name", "trigger_group") ON DELETE NO ACTION ON UPDATE NO ACTION;

-- ----------------------------
-- Foreign Keys structure for table QRTZ_TRIGGERS
-- ----------------------------
ALTER TABLE "public"."QRTZ_TRIGGERS" ADD CONSTRAINT "qrtz_triggers_ibfk_1" FOREIGN KEY ("sched_name", "job_name", "job_group") REFERENCES "public"."QRTZ_JOB_DETAILS" ("sched_name", "job_name", "job_group") ON DELETE NO ACTION ON UPDATE NO ACTION;
