CREATE TABLE Users(
 id int NOT NULL AUTO_INCREMENT,
 name   varchar(255),
 occupation   varchar(255),
 email   varchar(255),
 password_hash   varchar(255),
 avatar_file_name   varchar(255),
 Role   varchar(255),
 token   varchar,
 created_at   datetime,
 updated_at   datetime,
 PRIMARY KEY (id)
 );

CREATE TABLE Campaigns(
 id int NOT NULL AUTO_INCREMENT,
 users_id   int NOT NULL,
 name   varchar(255),
 short_description   varchar(255),
 goal_amount   int,
 current_amount  int,
 desctiption  text,
 perks  text,
 slug  varchar,
 backer_count  int,
 created_at  datetime,
 updated_at  datetime,
 PRIMARY KEY (id)
 );

CREATE TABLE Campaigns_Images()
 id int NOT NULL AUTO_INCREMENT,
 Campaigns_id int NOT NULL,
 file_name varchar,
 is_primary boolean,
 created_at datetime,
 updated_at datetime,
 PRIMARY_KEY (id)
 );

CREATE TABLE Transactions (
 id int NOT NULL AUTO_INCREMENT,
 Campaigns_id int NOT NULL,
 users_id int,
 amount int,
 status varchar,
 code varchar,
 created_at datetime,
 updated_at datetime,
 ;