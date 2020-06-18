CREATE TABLE Inventory
(
    idINV    INTEGER NOT NULL,
    name     VARCHAR(30),
    price    DECIMAL NOT NULL DEFAULT 100,
    PRIMARY KEY (idINV)
);
INSERT INTO Inventory (idINV, name, price)
VALUES (10,'trousers', 0);

INSERT INTO Inventory (idINV, name, price)
VALUES (11,'t-shirt',0);
INSERT INTO Inventory (idINV, name, price)
VALUES ( 12,'sandals',0);
INSERT INTO Inventory (idINV, name, price)
VALUES ( 13,'helmet',1);
INSERT INTO Inventory (idINV, name, price)
VALUES ( 14,'gloves ',1);
INSERT INTO Inventory (idINV, name, price)
VALUES ( 15,'jacket',1);
INSERT INTO Inventory (idINV, name, price)
VALUES ( 16,'boot',1);

INSERT INTO Inventory (idINV, name, price)
VALUES ( 20,'armor',0);

INSERT INTO Inventory (idINV, name, price)
VALUES (21 ,'bow', 0);

INSERT INTO Inventory (idINV, name, price)
VALUES (22 ,'crossbow ',1);
INSERT INTO Inventory (idINV, name, price)
VALUES ( 23,'rainbow',0);
INSERT INTO Inventory (idINV, name, price)
VALUES (24 ,'blade',1);

INSERT INTO Inventory (idINV, name, price)
VALUES (25 ,'sword',2);

INSERT INTO Inventory (idINV, name, price)
VALUES ( 26,'two-handed sword',3);
INSERT INTO Inventory (idINV, name, price)
VALUES (27 ,'ax',4);
INSERT INTO Inventory (idINV, name, price)
VALUES (28 ,'spear',5);
INSERT INTO Inventory (idINV, name, price)
VALUES ( 30,'magic shield',0);

INSERT INTO Inventory (idINV, name, price)
VALUES ( 31,'wooden shield',0);
INSERT INTO Inventory (idINV, name, price)
VALUES ( 32,'iron shield',1);
INSERT INTO Inventory (idINV, name, price)
VALUES (33 ,'plastic shield',3);
INSERT INTO Inventory (idINV, name, price)
VALUES (40 ,'gold',1);
INSERT INTO Inventory (idINV, name, price)
VALUES ( 41,'crystals',3);

INSERT INTO Inventory (idINV, name, price)
VALUES ( 50,'food',2);
