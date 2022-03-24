CREATE TABLE addresses (
	id integer AUTO_INCREMENT PRIMARY KEY,
	country varchar(50) NOT NULL,
	city varchar(50) NOT NULL,
	street varchar(50) NOT NULL,
	house varchar(50) NOT NULL,
	apartments varchar(50) 
);

CREATE TABLE persons (
	id integer AUTO_INCREMENT PRIMARY KEY,
	person_name varchar(50) NOT NULL,
	phone varchar(50) NOT NULL,
	addressID integer NOT NULL
);
