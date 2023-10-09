
CREATE TABLE IF NOT EXISTS staging.Employee(
    UUID uuid NULL,
    EmpNo VarChar(50) NULL,
    FName VarChar(100) NULL,
    LName VarChar(100) NULL,
    MName VarChar(100) NULL,
    Suffix VarChar(15) NULL,
    Position VarChar(100) NULL,
    OfficeId BigInt NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS idxEmployee ON staging.Employee(EmpNo);
CREATE UNIQUE INDEX IF NOT EXISTS idxEmployeeuuid ON staging.Employee(uuid);



elect * from staging.office


select * from office where uuid in (
'5e09397a-d1b9-46f8-bae9-7ba46e4febbc',
'4b942ed0-af9b-4e7f-bbc8-6567800cb6f2',
'4132fa43-e990-48e8-b0fb-0d2db2397ad1',
'bf3e966f-f70b-4b56-8c4a-21a049289266',
'f62a3f99-7d8a-4d3c-88a0-c12721b23ba8',
'b02bf534-395f-479e-8243-77142aca8be4')


select version()

select * from office where uuid = '2ffa8a47-fe54-4687-b842-7f16467e3f72'


INSERT INTO staging.AOMap(EmpNo, BrCode, AOID)
SELECT * FROM 
(VALUES
('200601-00565','E3',39),
('200712-01365','E3',52),
('200804-01644','E3',66),
('201309-05147','E3',10),
('201407-06085','E3',23),
('201408-06263','E3',28),
('201408-06267','E3',63),
('201410-06461','E3',35), 
('201506-07540','E3',44),
('201510-08476','E3',20),
('201511-08577','E3',7),
('201511-08579','E3',30),
('201511-08638','E3',18),
('201608-10959','E3',1),
('201612-12469','E3',36),
('201706-13191','E3',15),
('201707-14095','E3',22),
('201707-14531','E3',42),
('201710-16800','E3',17),
('201801-20708','E3',24),
('201802-17763','E3',29),
('201804-18338','E3',50),
('201805-20101','E3',51),
('201807-20704','E3',55),
('201807-21831','E3',67),
('201808-21881','E3',26),
('201811-24500','E3',16),
('201903-26328','E3',57),
('201905-28117','E3',14),
('201905-28120','E3',4),
('201906-29039','E3',11),
('202102-35883','E3',2),
('202208-40126','E3',31),
('202208-40353','E3',60),
('202210-41718','E3',68),
('201801-17287','E3',65),
('202203-38617','E3',9),
('202208-40736','E3',54),
('202208-40798','E3',45),
('202209-41252','E3',46),
('201807-20705','E3',3),
('201707-14532','E3',5),
('201702-13011','E3',6),
('201410-06461','E3',8),
('201811-24569','E3',21)) ao(EmpNo, BrCode, AOID)
