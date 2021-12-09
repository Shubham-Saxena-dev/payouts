CREATE DATABASE mydb;

GRANT ALL PRIVILEGES ON mydb.* TO
'myuser'@'%' IDENTIFIED BY 'mysql';
GRANT ALL PRIVILEGES ON mydb.* TO
'myuser'@'localhost' IDENTIFIED BY 'mysql';

USE mydb;

DROP TABLE IF EXISTS Payout;

CREATE TABLE `Payout`
(
    `SellerReference` int,
    `Currency`        varchar(20),
    `Amount`          varchar(255)
);

--Insert data into Payout Table
INSERT INTO Payout
VALUES ("1", "400", "USD")
INSERT INTO Payout
VALUES ("2", "1700", "EUR")
INSERT INTO Payout
VALUES ("3", "700", "GBP")
INSERT INTO Payout
VALUES ("4", "1100", "EUR")


