
DROP TABLE IF EXISTS Abgeordnete.User;

CREATE TABLE User(
	id int NOT NULL AUTO_INCREMENT,
	name varchar(255),
	email varchar(255),
	PRIMARY KEY (id)
	);
	
INSERT INTO Abgeordnete.User VALUES (0,"Fabian","fabian@hallo.de");
select * from Abgeordnete.User;


DROP TABLE IF EXISTS Abgeordnete.Abstimmung;
DROP TABLE IF EXISTS Abgeordnete.Gesetz;
DROP TABLE IF EXISTS Abgeordnete.Abgeordnete;

CREATE TABLE Abgeordnete.Abgeordnete(
	id int NOT NULL AUTO_INCREMENT,
	name varchar(255),
	partei varchar(255),
	PRIMARY KEY (id)
);

CREATE TABLE Abgeordnete.Gesetz(
	id int NOT NULL AUTO_INCREMENT,
	date varchar(255),
	titel varchar(255),
	PRIMARY KEY (id)
);




CREATE TABLE Abgeordnete.Abstimmung(
	id int NOT NULL AUTO_INCREMENT,
	gesetzId int,
	abgeorneteId int,
	vote varchar(255),
	PRIMARY KEY (id),
	FOREIGN KEY (gesetzId) REFERENCES Gesetz(id),
	FOREIGN KEY (abgeorneteId) REFERENCES Abgeordnete(id)
	UNIQUE KEY (abgeorneteId,gesetzId)
);

select * From Abstimmung where gesetzId = 11;


#INSERT INTO Abgeordnete.User (name,email) VALUES ("Fabian","fabian@hallo.de");

SELECT * from Gesetz;
SELECT * from Abgeordnete;





