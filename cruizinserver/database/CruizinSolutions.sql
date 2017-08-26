BEGIN TRANSACTION;
CREATE TABLE IF NOT EXISTS `vehicles` (
	`vid`	INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	`year`	TEXT,
	`make`	TEXT,
	`model`	TEXT
);
CREATE TABLE IF NOT EXISTS `tires` (
	`itemnum`	INTEGER,
	`brand`	TEXT,
	`model`	TEXT,
	`size`	TEXT,
	`servicedesc`	TEXT,
	`warranty`	TEXT,
	`condition`	TEXT,
	`price`	TEXT,
	`qty`	INTEGER,
	PRIMARY KEY(`itemnum`),
	FOREIGN KEY(`itemnum`) REFERENCES `items`(`itemnum`)
);
CREATE TABLE IF NOT EXISTS `services` (
	`itemnum`	INTEGER,
	`description`	TEXT,
	`price`	TEXT,
	PRIMARY KEY(`itemnum`),
	FOREIGN KEY(`itemnum`) REFERENCES `items`(`itemnum`)
);
CREATE TABLE IF NOT EXISTS `rims` (
	`itemnum`	INTEGER,
	`brand`	TEXT,
	`model`	TEXT,
	`size`	TEXT,
	`boltpattern`	TEXT,
	`finish`	TEXT,
	`condition`	TEXT,
	`price`	TEXT,
	`qty`	INTEGER,
	PRIMARY KEY(`itemnum`),
	FOREIGN KEY(`itemnum`) REFERENCES `items`(`itemnum`)
);
CREATE TABLE IF NOT EXISTS `parts` (
	`itemnum`	INTEGER,
	`description`	TEXT,
	`condition`	TEXT,
	`price`	TEXT,
	PRIMARY KEY(`itemnum`),
	FOREIGN KEY(`itemnum`) REFERENCES `items`(`itemnum`)
);
CREATE TABLE IF NOT EXISTS `packages` (
	`itemnum`	INTEGER,
	`description`	TEXT,
	`price`	TEXT,
	PRIMARY KEY(`itemnum`),
	FOREIGN KEY(`itemnum`) REFERENCES `items`(`itemnum`)
);
CREATE TABLE IF NOT EXISTS `orders` (
	`ordernum`	INTEGER PRIMARY KEY AUTOINCREMENT,
	`date`	TEXT,
	`cid`	INTEGER,
	`vid`	INTEGER,
	`odometer`	INTEGER,
	`comments`	TEXT,
	`subtotal`	TEXT,
	`tax`	TEXT,
	`total`	TEXT,
	`payment`	TEXT,
	FOREIGN KEY(`vid`) REFERENCES `vehicles`(`vid`),
	FOREIGN KEY(`cid`) REFERENCES `customers`(`cid`)
);
CREATE TABLE IF NOT EXISTS `items` (
	`itemnum`	INTEGER PRIMARY KEY AUTOINCREMENT,
	`description`	TEXT
);
CREATE TABLE IF NOT EXISTS `itemorders` (
	`id`	INTEGER PRIMARY KEY AUTOINCREMENT,
	`ordernum`	INTEGER,
	`itemnum`	INTEGER,
	`qty`	INTEGER,
	`amount`	TEXT,
	`price`	TEXT,
	FOREIGN KEY(`itemnum`) REFERENCES `items`(`itemnum`),
	FOREIGN KEY(`ordernum`) REFERENCES `orders`(`ordernum`)
);
CREATE TABLE IF NOT EXISTS `customers` (
	`cid`	INTEGER PRIMARY KEY AUTOINCREMENT,
	`name`	TEXT,
	`address`	TEXT,
	`city`	TEXT,
	`state`	TEXT,
	`zipcode`	TEXT,
	`phone`	TEXT,
	`email`	TEXT
);
COMMIT;
