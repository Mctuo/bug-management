use bug_management

create table `person_info`(
    `id` int AUTO_INCREMENT,
    `name` varchar(10) ,
    `account` bigint(12)  not null,
    `mail` varchar(20),
    `job` varchar(30) ,
    `note` varchar(30),
    `avatar` varchar(100),
    primary key (`id`),
	index  (`account`)
)engine=innodb,charset=utf8mb4;


create table `project`(
    `id` int AUTO_INCREMENT comment '项目编号',
    `project_name` varchar(30) comment '项目名称',
    `account` bigint(12)  not null comment '项目创建人',
    `project_people` int comment '项目人数',
    `task_total` int comment '项目总任务数',
    `task_unfinished` int comment '尚未完成的任务数',
    `task_finished` int comment '已经完成的任务数',
    `creater` int comment '创建人Id',
    primary key (`id`),
	index  (`account`)
)engine=innodb,charset=utf8mb4;

create table `project_people`(
     `id` int AUTO_INCREMENT comment '项目编号',
     `account` int comment '小组成员账号',
     index  (`id`)
)engine=innodb,charset=utf8mb4;


drop table test_case;

create table `test_case`(
    `id` int AUTO_INCREMENT comment '测试用例Id',
    `projectId` int  comment '项目Id',
    `title` varchar(30) comment '测试用例标题',
    `module_path` varchar(30) comment '被测试对象模块路径',
    `assigned` bigint(12) comment '指派给测试人员',
    `priority` int comment '优先级',
    `type_method` varchar(20) comment '类型方法',
    `type_plan` varchar(20) comment '测试计划',
    `creator` bigint(12) comment '创建人员',
     primary key (`id`),
     index (`creator`)
)engine=innodb,charset=utf8mb4;


drop table test_result;

create table `test_result`(
     `id` int AUTO_INCREMENT comment '测试用例Id',
     `case_id` int unique comment 'case_id',
     `projectId` int  comment '项目Id',
     `status` varchar(10) comment '状态',
     `assigned` bigint(12) comment '指派给开发人员',
     `test_env` varchar(30) comment '测试运行环境',
     `test_step` varchar(200) comment '步骤',
     primary key (`id`),
     index(`projectId`)
)engine=innodb,charset=utf8mb4;


drop table bug_info;

create table `bug_info`(
     `id` int AUTO_INCREMENT comment 'bugId',
     `bug_title` varchar(30) comment 'bug标题',
     `case_id` int unique comment 'case_id',
     `projectId` int  comment '项目Id',
     `module_path` varchar(30) comment '被测试对象模块路径',
     `assigned` bigint(12) comment '指派给开发人员',
     `severity` int comment '严重程度',
     `priority` int comment '优先级',
     `type` varchar(20) comment '类型',
     `find_way` varchar(100) comment '如何发现',
     `test_env` varchar(30) comment '测试运行环境',
     primary key (`id`),
     index(`projectId`)
)engine=innodb,charset=utf8mb4;



drop table bug_solution;

create table `bug_solution`(
    `id` int AUTO_INCREMENT comment 'bugId',
    `case_id` int unique comment 'case_id',
    `projectId` int comment '项目Id',
    `solver` bigint(12) comment '解决者',
    `solve_time` bigint(12) comment '解决日期',
    `solution` varchar(200) comment '解决方案',
    primary key (`id`),
    index(`projectId`)
)engine=innodb,charset=utf8mb4;

