----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Account_Type (
----------------------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Central_Office_ID bigint NOT NULL,
  Code bigint NOT NULL,
  Account_Type varchar(255) NOT NULL,
  Product_ID bigint NOT NULL,
  Group_ID bigint NULL,
  -- IIID bigint NULL,
  Normal_Balance bool NOT NULL,
  Isgl bool NOT NULL,
  Active bool NOT NULL,
  Filter_Type smallint NOT NULL DEFAULT 0, 
  -- 0 no Filter, 1 Whitelist, 2 Blacklist
  Other_Info jsonb NULL,
  
  CONSTRAINT Account_Type_pkey PRIMARY KEY (ID),
  CONSTRAINT idxAccount_Type_Code UNIQUE (Product_ID, Central_Office_ID, Code),
  CONSTRAINT idxAccount_Type_Name UNIQUE (Central_Office_ID, Account_Type),
  CONSTRAINT fk_Account_Type_Office FOREIGN KEY (Central_Office_ID) REFERENCES Office(ID),
  CONSTRAINT fk_Account_Type_Product FOREIGN KEY (Product_ID) REFERENCES Product(ID),
  CONSTRAINT fk_Account_Type_Group FOREIGN KEY (Group_ID) REFERENCES Reference(ID)
  -- CONSTRAINT fk_Account_Type_IIID FOREIGN KEY (IIID) REFERENCES Identity_Info(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccount_Type_UUID ON public.Account_Type(UUID);

DROP TRIGGER IF EXISTS trgAccount_Type_Ins on Account_Type;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Type_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Account_Type
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgAccount_Type_upd on Account_Type;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Type_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Account_Type
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgAccount_Type_del on Account_Type;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Type_del
---------------------------------------------------------------------------
    AFTER DELETE ON Account_Type
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();


----------------------------------------------------------------------------------------
CREATE OR REPLACE VIEW public.vwAccount_Type
----------------------------------------------------------------------------------------
AS SELECT 
    Acc.ID, mr.UUID,
    Acc.Code,
    Acc.Account_Type,
    
    md.ID Product_ID, md.Product_Name,
    
    grp.ID Account_Type_Group_ID, grp.UUID Account_Type_GroupUUID, 
    grp.Title Account_Type_Group,

    Acc.Normal_Balance,
    Acc.isgl,
    Acc.active,
    
    mr.Mod_Ctr,
    Acc.Other_Info,
    mr.Created,
    mr.Updated 
   FROM Account_Type Acc
   INNER JOIN Main_Record mr on mr.UUID = Acc.UUID
   JOIN Product md ON Acc.Product_ID = md.ID
   JOIN Reference grp ON grp.ID = Acc.Group_ID;
   
   INSERT INTO Account_Type(
     Central_Office_ID, Code, Account_Type, Product_ID, Group_ID, 
     Normal_Balance, isgl, active, Other_Info) 
   SELECT 
     o.ID, a.Code, a.Account_Type, m.ID Product_ID, grp.ID Group_ID, 
     a.Normal_Balance, a.isgl, a.active, NULL
   FROM (Values
      (0,'General','Accounting','Main',NULL,FALSE,FALSE,TRUE),
      (30,'Time Deposit','Bank Deposit','Main',NULL,FALSE,FALSE,TRUE),
      (20,'Savings Deposit','Bank Deposit','Main',NULL,FALSE,FALSE,TRUE),
      (10,'Demand Deposit','Bank Deposit','Main',NULL,FALSE,FALSE,TRUE),
      (100,'Golden Life 100','Collecting Facility','MicroInsurance',NULL,FALSE,FALSE,TRUE),
      (50,'Golden Life 50','Collecting Facility','MicroInsurance',NULL,FALSE,FALSE,TRUE),
      (20,'Mutual Fund','Collecting Facility','MicroInsurance',NULL,FALSE,FALSE,TRUE),
      (2,'Katuparan','Collecting Facility','MicroInsurance',NULL,FALSE,FALSE,TRUE),
      (1,'Loan Redemption Fund','Collecting Facility','MicroInsurance',NULL,FALSE,FALSE,TRUE),
      (800,'Donation','Donation','Regular',NULL,FALSE,FALSE,TRUE),
      (60,'Transportation Equipment','Fixed Asset','Main',NULL,TRUE,FALSE,TRUE),
      (40,'Leasehold Rights and Improvements','Fixed Asset','Main',NULL,TRUE,FALSE,TRUE),
      (10,'Land','Fixed Asset','Main',NULL,TRUE,FALSE,TRUE),
      (50,'Information Technology Equipment','Fixed Asset','Main',NULL,TRUE,FALSE,TRUE),
      (70,'Furniture, Fixtures and Equipment','Fixed Asset','Main',NULL,TRUE,FALSE,TRUE),
      (80,'Building Under Construction','Fixed Asset','Main',NULL,TRUE,FALSE,TRUE),
      (20,'Building','Fixed Asset','Main',NULL,TRUE,FALSE,TRUE),
      (30,'Appraisal Increment (With MB Approval)','Fixed Asset','Main',NULL,TRUE,FALSE,TRUE),

      (301,'Sikap 1','Loan','Microfinance',NULL,True,False,TRUE),
      (302,'Sikap 2','Loan','Microfinance',NULL,True,False,TRUE),
      (303,'Sikap 3','Loan','Microfinance',NULL,True,False,FALSE),
      (304,'Sikap 4','Loan','Microfinance',NULL,True,False,FALSE),
      (305,'Reserve 305','Loan','Microfinance',NULL,True,False,FALSE),
      (306,'Sipag Flex','Loan','Microfinance',NULL,True,False,FALSE),
      (307,'Sipag Term Loan','Loan','Microfinance',NULL,True,False,FALSE),
      (308,'Bundle Insurance Loan','Loan','Microfinance',NULL,True,False,FALSE),
      (309,'Sipag 4','Loan','Microfinance',NULL,True,False,FALSE),
      (310,'MA-Bundled Insurance Loan 2','Loan','Microfinance',NULL,True,False,FALSE),
      (311,'MF - Sikap (GLIP)','Loan','Microfinance',NULL,True,False,TRUE),
      (312,'Reserve 312','Loan','Microfinance',NULL,True,False,FALSE),
      (313,'Reserve 313','Loan','Microfinance',NULL,True,False,FALSE),
      (316,'Agri Loan','Loan','Microfinance',NULL,True,False,TRUE),
      (317,'Small Business Loan','Loan','Microfinance',NULL,True,False,TRUE),
      (318,'HOUSING REPAIRS/IMPROVEMENTS','Loan','Microfinance',NULL,True,False,TRUE),
      (319,'Reserve 319','Loan','Microfinance',NULL,True,False,FALSE),
      (320,'Reserve 320','Loan','Microfinance',NULL,True,False,FALSE),
      (321,'PHILHEALTH','Loan','Microfinance',NULL,True,False,TRUE),
      (322,'Health Loan - Akap','Loan','Microfinance',NULL,True,False,FALSE),
      (323,'SSS Premium Loan','Loan','Microfinance',NULL,True,False,TRUE),
      (324,'Reserve 324','Loan','Microfinance',NULL,True,False,FALSE),
      (325,'Reserve 325','Loan','Microfinance',NULL,True,False,FALSE),
      (326,'Tindahan Ni Inay Loan','Loan','Microfinance',NULL,True,False,FALSE),
      (327,'CBS','Loan','Microfinance',NULL,True,False,FALSE),
      (328,'Reserve 328','Loan','Microfinance',NULL,True,False,FALSE),
      (329,'Reserve 329','Loan','Microfinance',NULL,True,False,FALSE),
      (330,'Reserve 330','Loan','Microfinance',NULL,True,False,FALSE),
      (331,'Emergency Loan','Loan','Microfinance',NULL,True,False,FALSE),
      (332,'MFST- Calamity Loan','Loan','Microfinance',NULL,True,False,TRUE),
      (333,'Seosanal Loan','Loan','Microfinance',NULL,True,False,FALSE),
      (334,'Educational Loan','Loan','Microfinance',NULL,True,False,FALSE),
      (335,'Reserve 335','Loan','Microfinance',NULL,True,False,FALSE),
      (336,'Solar Loan','Loan','Microfinance',NULL,True,False,FALSE),
      (337,'Reserve 337','Loan','Microfinance',NULL,True,False,FALSE),
      (338,'Reserve 338','Loan','Microfinance',NULL,True,False,FALSE),
      (339,'Reserve 339','Loan','Microfinance',NULL,True,False,FALSE),
      (340,'Reserve 340','Loan','Microfinance',NULL,True,False,FALSE),
      (341,'Individual Loan','Loan','Microfinance',NULL,True,False,FALSE),
      (342,'Salary Loan','Loan','Microfinance',NULL,True,False,FALSE),
      (344,'EDUCATIONAL LOAN - ELEMENTARY','Loan','Microfinance',NULL,True,False,TRUE),
      (345,'Reserve 345','Loan','Microfinance',NULL,True,False,FALSE),
      (346,'Reserve 346','Loan','Microfinance',NULL,True,False,FALSE),
      (347,'Reserve 347','Loan','Microfinance',NULL,True,False,FALSE),
      (348,'Reserve 348','Loan','Microfinance',NULL,True,False,FALSE),
      (349,'Reserve 349','Loan','Microfinance',NULL,True,False,FALSE),
      (350,'Reserve 350','Loan','Microfinance',NULL,True,False,FALSE),
      (351,'SME Working Capital Loan','Loan','Microfinance',NULL,True,False,FALSE),
      (352,'SME Investment Loan','Loan','Microfinance',NULL,True,False,FALSE),
      (353,'Reserve 353','Loan','Microfinance',NULL,True,False,FALSE),
      (354,'Reserve 354','Loan','Microfinance',NULL,True,False,FALSE),
      (355,'Reserve 355','Loan','Microfinance',NULL,True,False,FALSE),
      (356,'Reserve 356','Loan','Microfinance',NULL,True,False,FALSE),
      (357,'Reserve 357','Loan','Microfinance',NULL,True,False,FALSE),
      (358,'Reserve 358','Loan','Microfinance',NULL,True,False,FALSE),
      (359,'Reserve 359','Loan','Microfinance',NULL,True,False,FALSE),
      (360,'Reserve 360','Loan','Microfinance',NULL,True,False,FALSE),
      (361,'Reserve 361','Loan','Microfinance',NULL,True,False,FALSE),
      (362,'Reserve 362','Loan','Microfinance',NULL,True,False,FALSE),
      (363,'Reserve 363','Loan','Microfinance',NULL,True,False,FALSE),
      (364,'Reserve 364','Loan','Microfinance',NULL,True,False,FALSE),
      (365,'Reserve 365','Loan','Microfinance',NULL,True,False,FALSE),
      (366,'Reserve 366','Loan','Microfinance',NULL,True,False,FALSE),
      (367,'Reserve 367','Loan','Microfinance',NULL,True,False,FALSE),
      (368,'Reserve 368','Loan','Microfinance',NULL,True,False,FALSE),
      (369,'Reserve 369','Loan','Microfinance',NULL,True,False,FALSE),
      (370,'Reserve 370','Loan','Microfinance',NULL,True,False,FALSE),
      (371,'Reserve 371','Loan','Microfinance',NULL,True,False,FALSE),
      (372,'Reserve 372','Loan','Microfinance',NULL,True,False,FALSE),
      (373,'Reserve 373','Loan','Microfinance',NULL,True,False,FALSE),
      (374,'Reserve 374','Loan','Microfinance',NULL,True,False,FALSE),
      (375,'Reserve 375','Loan','Microfinance',NULL,True,False,FALSE),
      (376,'Reserve 376','Loan','Microfinance',NULL,True,False,FALSE),
      (377,'Reserve 377','Loan','Microfinance',NULL,True,False,FALSE),
      (378,'Reserve 378','Loan','Microfinance',NULL,True,False,FALSE),
      (379,'Reserve 379','Loan','Microfinance',NULL,True,False,FALSE),
      (380,'Reserve 380','Loan','Microfinance',NULL,True,False,FALSE),
      (381,'Reserve 381','Loan','Microfinance',NULL,True,False,FALSE),
      (382,'Reserve 382','Loan','Microfinance',NULL,True,False,FALSE),
      (383,'Reserve 383','Loan','Microfinance',NULL,True,False,FALSE),
      (384,'Reserve 384','Loan','Microfinance',NULL,True,False,FALSE),
      (385,'Reserve 385','Loan','Microfinance',NULL,True,False,FALSE),
      (386,'Reserve 386','Loan','Microfinance',NULL,True,False,FALSE),
      (387,'Reserve 387','Loan','Microfinance',NULL,True,False,FALSE),
      (388,'Reserve 388','Loan','Microfinance',NULL,True,False,FALSE),
      (389,'Reserve 389','Loan','Microfinance',NULL,True,False,FALSE),
      (390,'Reserve 390','Loan','Microfinance',NULL,True,False,FALSE),
      (391,'Reserve 391','Loan','Microfinance',NULL,True,False,FALSE),
      (392,'Reserve 392','Loan','Microfinance',NULL,True,False,FALSE),
      (393,'Reserve 393','Loan','Microfinance',NULL,True,False,FALSE),
      (394,'Reserve 394','Loan','Microfinance',NULL,True,False,FALSE),
      (395,'Reserve 395','Loan','Microfinance',NULL,True,False,FALSE),
      (396,'Reserve 396','Loan','Microfinance',NULL,True,False,FALSE),
      (397,'Reserve 397','Loan','Microfinance',NULL,True,False,FALSE),
      (398,'Unsecured Restructured Loans','Loan','Microfinance',NULL,True,False,FALSE),
      (399,'Secured Restructured Loans','Loan','Microfinance',NULL,True,False,FALSE),
      (410,'DSHP - Sikap 1','Loan','Microfinance',NULL,True,False,FALSE),
      (411,'DSHP - Sikap 2','Loan','Microfinance',NULL,True,False,FALSE),
      (413,'DSHP - Solar Loan','Loan','Microfinance',NULL,True,False,FALSE),
      (414,'CARD Care','Loan','Microfinance',NULL,True,False,FALSE),
      (415,'Paid Plan','Loan','Microfinance',NULL,True,False,FALSE),
      (416,'DSHP - Agri','Loan','Microfinance',NULL,True,False,FALSE),
      (417,'SAGIP PLAN','Loan','Microfinance',NULL,True,False,FALSE),
      (418,'Educational Loan - HIGH SCHOOL','Loan','Microfinance',NULL,True,False,TRUE),
      (419,'EDUCATIONAL LOAN - COLLEGE','Loan','Microfinance',NULL,True,False,TRUE),
      (420,'IPL-CAMIA (BUNDLED)','Loan','Microfinance',NULL,True,False,TRUE),
      (421,'DSHP Insurance Premium Loan','Loan','Microfinance',NULL,True,False,FALSE),
      (422,'DSHP Philhealth Loan','Loan','Microfinance',NULL,True,False,FALSE),
      (423,'DSHP SSS Premium Loan','Loan','Microfinance',NULL,True,False,FALSE),
      (424,'Kabuklod Plan','Loan','Microfinance',NULL,True,False,FALSE),
      (425,'DSHP - Kabuklod Plan','Loan','Microfinance',NULL,True,False,FALSE),
      (443,'RPA-Educational Loan - HS & College','Loan','Microfinance',NULL,True,False,FALSE),
      (444,'DSHP Elementary','Loan','Microfinance',NULL,True,False,FALSE),
      (445,'DSHP High School','Loan','Microfinance',NULL,True,False,FALSE),
      (446,'DSHP College','Loan','Microfinance',NULL,True,False,FALSE),
      (447,'DSHP Housing and repair','Loan','Microfinance',NULL,True,False,FALSE),
      (448,'Latrine Loan','Loan','Microfinance',NULL,True,False,TRUE),
      (451,'RPA - REGULAR AGRI LOAN','Loan','Microfinance',NULL,True,False,TRUE),
      (461,'IPL-MBA (FSP)','Loan','Microfinance',NULL,True,False,TRUE),
      (462,'Gadget Loan','Loan','Microfinance',NULL,True,False,TRUE),
      (463,'MF - Sikap','Loan','Microfinance',NULL,True,False,TRUE),
      (464,'Sikap - Restructured','Loan','Microfinance',NULL,True,False,TRUE),
      (465,'OL - Restructured','Loan','Microfinance',NULL,True,False,TRUE),
      (474,'Special Liquidity Fund','Loan','Microfinance',NULL,True,False,TRUE),
      (475,'IPL - Rimansi (Dakila)','Loan','Microfinance',NULL,True,False,TRUE),
      (476,'OL - Padala Now, Pay Later','Loan','Microfinance',NULL,True,False,TRUE),
      (477,'IPL - Konek2Protek','Loan','Microfinance',NULL,True,False,TRUE),

      (20,'Borrowed Funds','Loans Payable','Main',NULL,FALSE,FALSE,TRUE),
      (10,'Bills Payable','Loans Payable','Main',NULL,FALSE,FALSE,TRUE),
      (120,'EMPC Pension','Payable','Main',NULL,FALSE,FALSE,TRUE),
      (110,'EMPC Loan','Payable','Main',NULL,FALSE,FALSE,TRUE),
      (100,'Employees Contribution','Payable','Main',NULL,FALSE,FALSE,TRUE),
      (20,'Overages','Payable','Main',NULL,FALSE,FALSE,TRUE),
      (10,'Accounts Payable','Payable','Main',NULL,FALSE,FALSE,TRUE),
      (20,'Shortages','Receivable','Main',NULL,TRUE,FALSE,TRUE),
      (10,'Accounts Receivable','Receivable','Main',NULL,TRUE,FALSE,TRUE),
      (700,'Cash Advance','Receivable','Regular',NULL,FALSE,FALSE,TRUE),
      (240,'Special Time Deposit','Savings','Regular',NULL,FALSE,FALSE,TRUE),
      (80,'Tagumpay Account (Member)','Savings','Regular',NULL,FALSE,FALSE,TRUE),
      (70,'Tagumpay Account (Non-Member)','Savings','Regular',NULL,FALSE,FALSE,TRUE),
      (50,'Kayang-Kaya Account','Savings','Regular',NULL,FALSE,FALSE,TRUE),
      (40,'Tagumpay Savings','Savings','Regular',NULL,FALSE,FALSE,TRUE),
      (30,'Maagap Savings Account','Savings','Regular',NULL,FALSE,FALSE,TRUE),
      (60,'Pledge Account','Savings','Microfinance',NULL,FALSE,FALSE,TRUE),
      (120,'Interest Bearing Current','Savings','Current Account',NULL,FALSE,FALSE,TRUE),
      (710,'Regular Current Account','Savings','Current Account',NULL,FALSE,FALSE,TRUE),
      (720,'Elementary','School','Regular',NULL,FALSE,FALSE,TRUE),
      (730,'High School','School','Regular',NULL,FALSE,FALSE,TRUE),
      (740,'College','School','Regular',NULL,FALSE,FALSE,TRUE),
      (750,'Seminars','School','Regular',NULL,FALSE,FALSE,TRUE)
      )   
    a(Code, Account_Type, Product, grp, Alternate_ID, Normal_Balance, isGL, Active) 

      
  INNER JOIN Product m on m.Product_Name = a.Product
  INNER JOIN vwReference grp on lower(grp.Title) = lower(a.grp) 
    and lower(grp.Ref_Type) = 'accounttypegroup'
  LEFT JOIN Office o  on lower(o.Alternate_ID) = lower('1')  
  
  ON CONFLICT (Product_ID, Central_Office_ID, Code) DO UPDATE SET
    Central_Office_ID = EXCLUDED.Central_Office_ID,
    Product_ID = EXCLUDED.Product_ID,
    Group_ID = EXCLUDED.Group_ID,
    Code = EXCLUDED.Code,
    active = EXCLUDED.active,   
    Normal_Balance = EXCLUDED.Normal_Balance,
    isgl = EXCLUDED.isgl,
    Account_Type = EXCLUDED.Account_Type,
    Other_Info = EXCLUDED.Other_Info
  ;
