use bug_management

create table `person_info`(
    `id` int AUTO_INCREMENT,
    `name` varchar(10) ,
    `account` bigint(12) unique not null,
    `mail` varchar(20),
    `job` varchar(30) ,
    `note` varchar(30),
    `avatar` varchar(100),
    primary key (`id`),
	index  (`account`)
)engine=innodb,charset=utf8mb4;


create table `project_info`(
    `id` int AUTO_INCREMENT,
    `note` varchar(30),
    `role` varchar (30),
    `member` bigint(12) not null,
    primary key (`id`),
	index  (`account`)
)engine=innodb,charset=utf8mb4;