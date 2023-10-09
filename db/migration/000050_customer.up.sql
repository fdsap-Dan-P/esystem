----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Customer (
----------------------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  IIID bigint NOT NULL,
  Central_Office_ID bigint NOT NULL,
  CID bigint NOT NULL,
  Customer_Alt_ID varchar(100) NULL,
  Debit_Limit float8 NOT NULL,
  Credit_Limit float8 NOT NULL,
  Date_Entry Date NULL,
  Date_Recognized Date NULL,
  Date_Resigned Date NULL,
  Resigned bool NULL,
  Reason_Resigned varchar(100) NULL,
  Last_Activity_Date Date NULL,
  dosri bool NOT NULL,
  Classification_ID bigint NULL,
  Sub_Classification_ID bigint NULL,
  Customer_Group_ID bigint NULL,
  Office_ID bigint NOT NULL,
  Restriction_ID bigint NULL,
  Risk_Class_ID bigint NULL,
  Industry_ID bigint NULL,
  Status_Code integer NOT NULL,
  Refferedby_ID bigint NULL,
  Remarks varchar(200) NULL,
  Primary_Acc varchar(21) NULL,
  Other_Info jsonb NULL,
  
  CONSTRAINT Customer_pkey PRIMARY KEY (ID),
  CONSTRAINT idxCustomer_Alt_ID UNIQUE (Customer_Alt_ID),
  CONSTRAINT idxCustomerCID UNIQUE (Central_Office_ID, CID),
  CONSTRAINT fkCustomerCentralOffice FOREIGN KEY (Central_Office_ID) REFERENCES Office(ID),
  CONSTRAINT fkCustomerClassification FOREIGN KEY (Classification_ID) REFERENCES Reference(ID),
  CONSTRAINT fkCustomerCustomer_Group FOREIGN KEY (Customer_Group_ID) REFERENCES Customer_Group(ID),
  CONSTRAINT fkCustomer_IDentity_Info FOREIGN KEY (IIID) REFERENCES Identity_Info(ID),
  CONSTRAINT fkCustomerOffice FOREIGN KEY (Office_ID ) REFERENCES Office(ID),
  CONSTRAINT fkCustomerRestrictionid FOREIGN KEY (Restriction_ID) REFERENCES Reference(ID),
  CONSTRAINT fkCustomerRisk_Class FOREIGN KEY (Risk_Class_ID) REFERENCES Reference(ID),
  CONSTRAINT fkCustomerRisk_Reffered FOREIGN KEY (Refferedby_ID) REFERENCES Identity_Info(ID),
--  CONSTRAINT fkCustomerStatus FOREIGN KEY (Status_ID) REFERENCES Reference(ID),
  CONSTRAINT fkCustomerSub_Classification FOREIGN KEY (Sub_Classification_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxCustomer_UUID ON public.Customer(UUID);

DROP TRIGGER IF EXISTS trgCustomerIns on Customer;
---------------------------------------------------------------------------
CREATE TRIGGER trgCustomerIns
---------------------------------------------------------------------------
    BEFORE INSERT ON Customer
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgCustomerupd on Customer;
---------------------------------------------------------------------------
CREATE TRIGGER trgCustomerupd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Customer
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgCustomer_del on Customer;
---------------------------------------------------------------------------
CREATE TRIGGER trgCustomer_del
---------------------------------------------------------------------------
    AFTER DELETE ON Customer
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();


----------------------------------------------------------------------------------------
CREATE OR REPLACE VIEW public.vwCustomer
----------------------------------------------------------------------------------------
AS SELECT
    cust.ID, mr.UUID,
    cust.IIID,
    p.Title,    
    p.Last_Name,
    p.First_Name,
    p.Middle_Name,
    p.Mother_Maiden_Name,
    p.Birthday,
    p.Sex,
    
    p.Gender_ID, p.Gender,    
    
    p.Civil_Status_ID , p.Civil_Status,
        
    p.Birth_Place_ID, p.Birth_Place, p.Full_Birth_Place, p.Birth_PlaceURL, 
    
    p.Contact_ID, 
    p.Contact_Last_Name, p.Contact_First_Name, p.Contact_Middle_Name,
    
    p.Alternate_ID, p.Phone, p.Email, p.Identity_Map_ID, p.Simple_Name, 

    p.iiMod_Ctr,
    p.iiOther_Info,
    p.iiCreated,
    p.iiUpdated,
    p.Current_Address_ID,
    p.Current_Detail,
    p.Current_URL,    
    p.Location,
    p.Full_Location,
 
    p.Marriage_Date,
    p.Known_Language,
    p.isAdopted,
  
    p.Source_Income_ID, p.Source_Income,
    p.Disability_ID, p.Disability,
    p.Occupation_ID, p.Occupation,
    p.Sector_ID, p.Sector,
    p.Industry_ID, p.industry,
    p.Religion_ID, p.Religion,
    p.Nationality_ID, p.Nationality,         
    
    cust.dosri,
    cust.Customer_Alt_ID,

    co.ID Central_Office_ID, co.Office_Name CentralOffice_Name,
    cust.CID,
 
    cust.Credit_Limit,
    cust.Date_Entry,
    cust.Debit_Limit,
    cust.Last_Activity_Date,

    
    cls.ID Classification_ID, cls.UUID ClassificationUUID, cls.Title Classification,
    grp.ID Customer_Group_ID, grp.UUID Customer_GroupUUID, grp.Group_Name Customer_Group_Name,
    
    o.ID Office_ID , o.UUID OfficeUUID, o.Office_Name,

    rest.ID Restriction_ID, rest.UUID RestrictionUUID, rest.Title Restriction,
    Risk.ID Risk_Class_ID, Risk.UUID Risk_ClassUUID, Risk.Title Risk_Class,
    cust.Status_Code, stat.ID Status_ID, stat.UUID StatusUUID, stat.Title Status,
    subcls.ID Sub_Classification_ID,  subcls.UUID Sub_ClassificationUUID, subcls.Title Sub_Classification,
    
    mr.Mod_Ctr,
    cust.Other_Info,
    mr.Created,
    mr.Updated 
    
   FROM Customer cust
   INNER JOIN Main_Record mr on mr.UUID = cust.UUID  
   LEFT JOIN vwperson p ON cust.IIID = p.IIID
   LEFT JOIN Reference cls ON cls.ID = cust.Classification_ID
   LEFT JOIN Reference rest ON rest.ID = cust.Restriction_ID
   LEFT JOIN Reference Risk ON Risk.ID = cust.Risk_Class_ID
   LEFT JOIN Reference stat ON stat.Code = cust.Status_Code and lower(stat.Ref_Type) = 'customercode'
   LEFT JOIN Reference subcls ON subcls.ID = cust.Sub_Classification_ID

   LEFT JOIN Customer_Group grp ON cls.ID = cust.Customer_Group_ID

   LEFT JOIN Office o  ON o.ID = cust.Office_ID 
   LEFT JOIN Office co ON co.ID = cust.Central_Office_ID;
 
 
  INSERT INTO Identity_Info 
   (isPerson, Last_Name, Alternate_ID)
  SELECT FALSE isPerson, 'System' Last_Name, 'System'    
  
  ON CONFLICT (Alternate_ID) 
  DO UPDATE SET 
    Last_Name = EXCLUDED.Last_Name;
  
  INSERT INTO Customer (
    IIID, Central_Office_ID, CID, Customer_Alt_ID, Debit_Limit, Credit_Limit, 
    dosri, Office_ID 
    ) 
  SELECT 
   ii.ID IIID, o.ID Central_Office_ID, 0 CID, 'System' Customer_Alt_ID, 
   cast(-1 as double precision) Debit_Limit, cast(-1 as double precision) Credit_Limit, 
   FALSE dosri, o.ID Office_ID 
  FROM Identity_Info ii, vwOffice o
  WHERE ii.Alternate_ID = 'System' and o.OfficeType = 'System' and o.Office_Name = 'System'
  
  
  ON CONFLICT(Central_Office_ID, CID) 
  DO UPDATE SET
    IIID = excluded.IIID, 
    Central_Office_ID = excluded.Central_Office_ID, 
    CID = excluded.CID, 
    Customer_Alt_ID = excluded.Customer_Alt_ID, 
    Debit_Limit = excluded.Debit_Limit, 
    Credit_Limit = excluded.Credit_Limit, 
    dosri = excluded.dosri, 
    Office_ID  = excluded.Office_ID 
  ; 

/* 
 ----------------------------------------------------------------------------------------
DO $$
----------------------------------------------------------------------------------------
DECLARE 
  Titles TEXT DEFAULT '';
  rec   RECORD;
  cur CURSOR FOR 
  
   SELECT 
     FALSE isPerson, a.Last_Name, cast(a.Birthday as Date), 
     ct.ID Contact_ID, a.Alternate_ID, Address_Detail, Address_URL
   FROM
     (ValueS
      ('United Church of Christ in the Philippines','05/25/1948','1067','10001','West Triangle, Quezon City. Metro Manila','https://www.google.com/Maps/@14.6462099,121.0360189,3a,75y,239.26h,86.91t/data=!3m6!1e1!3m4!1s1PFeC-7yipSjCZOmoFP8dg!2e0!7i13312!8i6656'),
      ('South Luzon Jurisdictional Area',null,'1054','10002','Los Baños, Laguna','https://www.google.com/Maps/@14.1672649,121.2405257,3a,75y,339.28h,80.44t/data=!3m6!1e1!3m4!1sJ708bLw0rcP9IneTpkuLBw!2e0!7i13312!8i6656'),
      ('Batangas Associate Conference',null,'1072','10003','Lipa City','https://www.google.com/Maps/@13.9410034,121.158245,0a,82.2y,144.39h,96.49t/data=!3m4!1e1!3m2!1sHt4XitK9nN2MapFCRuXyRA!2e0?Source=apiv3'),
      ('BPI - Los Banos',null,null,'10004','Batong Malake, Los Baños City, Laguna','0'),
      ('BPI - New Mia',null,null,'10005','City of Parañaque, Metro Manila','0'),
      ('EBF',null,'1068','10006','Quezon City, 1104 Metro Manila','0'),
      ('IDPIP-ST',null,'1064','10007','Batong Malake, Los Baños, Laguna','0'),
      ('Kowloon Union Church, Hongkong',null,'1062','10008',' Kwun Chung, Hong Kong','0'),
      ('LCSMC',null,'1006','10009','Malate, Manila, 1004 Metro Manila','https://www.google.com/Maps/@14.574546,120.9873325,3a,60y,59.63h,87.11t/data=!3m6!1e1!3m4!1sEF33csD5NRunEWqAk1XYzA!2e0!7i13312!8i6656'),
      ('North Bicol Conference',null,'1061','10010','San Jose, San Jose City, OcCIDental Mindoro','https://www.google.com/Maps/@12.3533273,121.071089,3a,75y,219.41h,85.77t/data=!3m6!1e1!3m4!1ssEd4huRj8H-ebteOT4PA!2e0!7i13312!8i6656'),
      ('OcCIDental Mindoro Assoc. Conference',null,'1011','10011','Brgy. Bethel, Victoria, Oriental Mindoro','https://www.google.com/Maps/@13.1463288,121.251938,3a,82.2y,7.23h,69.42t/data=!3m6!1e1!3m4!1sBH-x3JBXbt5qzAIkH0zHBw!2e0!7i13312!8i6656'),
      ('UCCP Palawan Associate Conference',null,'1059','10012','Magara, Roxas, Palawan','https://www.google.com/Maps/@10.3036729,119.2645837,3a,66.6y,130.41h,80.27t/data=!3m6!1e1!3m4!1sHmT732qANcSVTpjxqLT5Q!2e0!7i13312!8i6656'),
      ('PAC - UCCP Magara',null,'1071','10013','Brgy. San Miguel, Puerto Princesa, Palawan','https://www.google.com/Maps/@9.7524795,118.7661807,0a,82.2y,311.09h,82.77t/data=!3m4!1e1!3m2!1s8DQjMvxYQj8yTpP4ojYwrg!2e0?Source=apiv3'),
      ('PAC District V',null,'1059','10014','Calamba City, Laguna','0'),
      ('Pastive Arce',null,'1069','10015','Lusacan, Tiaong, Quezon','https://www.google.com/Maps/@13.9628722,121.3211301,0a,82.2y,210.33h,86.12t/data=!3m4!1e1!3m2!1s17lxe63AuXOTXXs0dkiOFQ!2e0?Source=apiv3'),
      ('South Bicol Conference',null,'1024','10016','San Pablo City, Laguna','https://www.google.com/Maps/@14.0652933,121.3235006,0a,82.2y,332.31h,90.61t/data=!3m4!1e1!3m2!1sC1ho4OvyA-SuWDfB37QfIA!2e0?Source=apiv3'),
      ('Southern Tagalog Conference',null,'1063','10017','Amadeo, Cavite','https://www.google.com/Maps/@14.1345142,120.9399264,0a,82.2y,297.89h,84.46t/data=!3m4!1e1!3m2!1sa-0IfeiihWO5E4i0OR6Hqg!2e0?Source=apiv3'),
      ('STC - FEC',null,'1063','10018','Los Baños, Laguna','https://www.google.com/Maps/@14.1767695,121.2198035,0a,82.2y,353.06h,93.66t/data=!3m4!1e1!3m2!1sRqObxTroK69XRlJF9y19Eg!2e0?Source=apiv3'),
      ('UCCP Amadeo',null,'1004','10019','Poblacion, Batangas, 4200 Batangas','https://www.google.com/Maps/@13.7587154,121.0599069,0a,66.6y,264.17h,80.72t/data=!3m4!1e1!3m2!1sUB0-rftb9zUMCFpBPIOyvg!2e0?Source=apiv3'),
      ('UCCP Bambang - OMC',null,'1065','10020','Tagum, Davao del Norte','https://www.google.com/Maps/@7.4473733,125.8003235,0a,82.2y,322.55h,100.02t/data=!3m4!1e1!3m2!1sJ7HaDlfZPvp2HEf9Zw9ZRQ!2e0?Source=apiv3'),
      ('UCCP Batangas',null,'1073','10021','Buli, Pinamalayan, Oriental Mindoro','https://www.google.com/Maps/@13.0667505,121.4980613,0a,66.6y,256.51h,85.89t/data=!3m4!1e1!3m2!1sfuil1lS-B7AdY3n59sBiUw!2e0?Source=apiv3'),
      ('UCCP Bethel',null,'1060','10022','0','0'),
      ('UCCP Buli - OMC',null,'1011','10023','Los Baños, Laguna','https://www.google.com/Maps/@14.1672649,121.2405257,3a,75y,339.28h,80.44t/data=!3m6!1e1!3m4!1sJ708bLw0rcP9IneTpkuLBw!2e0!7i13312!8i6656'),
      ('UCCP Campung Alay',null,null,'10024','Culasisi, Sablayan, Oriental Mindoro','https://www.google.com/Maps/@12.8581433,120.8926738,0a,66.6y,309.76h,80.75t/data=!3m4!1e1!3m2!1ssGoc88Hb2O3lGjMA1Zp1pg!2e0?Source=apiv3'),
      ('UCCP Church Among the Palms',null,'1070','10025','Rizal, Palawan','https://www.google.com/Maps/@7.4478307,125.8003397,3a,75y,188h,69.41t/data=!3m6!1e1!3m4!1sVilb9Hxm2IL9TTQNNFkV0w!2e0!7i13312!8i6656?hl=fil'),
      ('UCCP Culasisi',null,'1011','10026','Labasan, Bongabong, Oriental Mindoro ','https://www.google.com/Maps/@12.7782126,121.469986,0a,82.2y,294.36h,70.11t/data=!3m4!1e1!3m2!1szaakmdVAlIfaRW5yptpDuA!2e0?Source=apiv3'),
      ('UCCP Danum-danum',null,'1059','10027','Manila, Metro manila','https://www.google.com/Maps/@14.5747202,120.9872387,0a,82.2y,15.21h,97.14t/data=!3m4!1e1!3m2!1s1EGvfpcYNuaa9xpdIee5IA!2e0?Source=apiv3'),
      ('UCCP Labasan  - OMC',null,'1011','10028','Lipa, Batangas','https://www.google.com/Maps/@13.9410034,121.158245,0a,82.2y,144.39h,96.49t/data=!3m4!1e1!3m2!1sHt4XitK9nN2MapFCRuXyRA!2e0?Source=apiv3'),
      ('UCCP LCSMC',null,'1006','10029','Rizal, Palawan','0'),
      ('UCCP Lipa Evangelical church',null,'1072','10030','West Triangle, Quezon City. Metro Manila','https://www.google.com/Maps/@14.6462099,121.0360189,3a,75y,239.26h,86.91t/data=!3m6!1e1!3m4!1s1PFeC-7yipSjCZOmoFP8dg!2e0!7i13312!8i6656'),
      ('UCCP Malapandig',null,'1059','10031','Los Baños, Laguna','https://www.google.com/Maps/@14.1672649,121.2405257,3a,75y,339.28h,80.44t/data=!3m6!1e1!3m4!1sJ708bLw0rcP9IneTpkuLBw!2e0!7i13312!8i6656'),
      ('UCCP National Office',null,'1067','10032','Peñafrancia, Naga City','https://www.google.com/Maps/@13.6247431,123.2039188,3a,60y,54.07h,81.84t/data=!3m6!1e1!3m4!1slloRygB3buiK4oQerIC1Ug!2e0!7i13312!8i6656'),
      ('UCCP NESTCON',null,'1057','10033','San Pedro, Laguna','https://www.google.com/Maps/@14.3441138,121.0556519,3a,82.2y,170.06h,91.44t/data=!3m6!1e1!3m4!1su0pUg2GyJYwAIEPlUKDxqw!2e0!7i13312!8i6656'),
      ('UCCP North Bicol Conference',null,'1012','10034','Pagsanjan, Laguna','https://www.google.com/Maps/@14.2749298,121.4531951,3a,66.6y,90t/data=!3m6!1e1!3m4!1sHt36GMDn3ac9Hygd5Z1ow!2e0!7i13312!8i6656'),
      ('UCCP PACITA',null,'1003','10035','Panitian, Quezon, Palawan','https://www.google.com/Maps/@9.2381267,118.0270671,0a,82.2y,78h,70.77t/data=!3m4!1e1!3m2!1sXreK2SHQS9EcVYFiEX79mQ!2e0?Source=apiv3'),
      ('UCCP Pagsanjan',null,'1066','10036','Pinamalayan','https://www.google.com/Maps/@13.0667505,121.4980613,0a,66.6y,256.51h,85.89t/data=!3m4!1e1!3m2!1sfuil1lS-B7AdY3n59sBiUw!2e0?Source=apiv3'),
      ('UCCP Panitian',null,'1059','10037','Salcedo, Bansud, Oriental Mindoro','0'),
      ('UCCP Pinamalayan - OMC',null,'1011','10038','San Buenaventura, Luisiana, Laguna','https://www.google.com/Maps/@14.1861381,121.5105019,0a,82.2y,156.69h,89.75t/data=!3m4!1e1!3m2!1sZIxeIy7ju0xT9XKLL53aFA!2e0?Source=apiv3'),
      ('UCCP Salcedo - OMC',null,'1011','10039','Poblacion, San Vicente, Palawan','https://www.google.com/Maps/@10.5184238,119.3514341,0a,66.6y,151.27h,76.06t/data=!3m4!1e1!3m2!1symBgiTC7U0sHMST4i6NoQw!2e0?Source=apiv3'),
      ('UCCP San Buenaventura',null,'1057','10040','Pangil, Laguna','https://www.google.com/Maps/@14.408538,121.4741708,0a,82.2y,42.67h,101.17t/data=!3m4!1e1!3m2!1srdLSFqhtmNB-nAOhZ89vyQ!2e0?Source=apiv3'),
      ('UCCP San Vicente',null,'1059','10041','Sucol, Calamba, Laguna','https://www.google.com/Maps/@14.1802064,121.2000268,3a,75y,280.64h,88.45t,359.32r/data=!3m6!1e1!3m4!1ssFqa7JrEw-cJ2kxLB9tBnQ!2e0!7i13312!8i6656'),
      ('UCCP SKP NESTCON',null,'1058','10042','Taytay, Palawan','0'),
      ('UCCP Sucol',null,'1074','10043','0','0'),
      ('UCCP Taytay',null,'1059','10044','Dasmariñas City','0'),
      ('UCCP Tubon',null,null,'10045','Victoria, Laguna','0'),
      ('UCCP UCC',null,'1004','10046','Victoria, oriental','0'),
      ('UCCP Victoria - NESTCON',null,'1057','10047','0','0'),
      ('UCCP Victoria - OMC',null,'1011','10048','0','0'),
      ('Uniting Church of Sweden',null,'1056','10049','0','0'),
      ('Oriental Mindoro Conference',null,null,'1075','0','0'),
      ('PAC thru Perla Francisco',null,null,'1076','0','0'),
      ('UCCP NBC',null,null,'1077','0','0'),
      ('UCCP OMAC',null,null,'1078','0','0'),
      ('CCTV Kwiknet',null,null,'1082','0','0'),
      ('Ecumenical Bishops Forum, Inc',null,null,'1084','0','0'),
      ('El Guapo',null,null,'1085','0','0'),
      ('Gabriela',null,null,'1086','0','0'),
      ('Holiday Planners Travel and Tours',null,null,'1087','0','0'),
      ('Jesus Gili Thru Marlyn Anicete',null,null,'1088','0','0'),
      ('Kairos Philippines',null,null,'1089','0','0'),
      ('LBC thru MBA',null,null,'1090','0','0'),
      ('Northwest Mindanao Jurisdiction',null,null,'1093','0','0'),
      ('UCCP NWMJ',null,null,'1098','0','0'),
      ('Union Theological Seminary',null,null,'1099','0','0')
    ) a (Last_Name, Birthday, Contactaltid, Alternate_ID, Address_Detail, Address_URL)
  LEFT JOIN Identity_Info ii on ii.Alternate_ID =  a.Alternate_ID
  LEFT JOIN Identity_Info ct on ct.Alternate_ID =  a.Contactaltid;

    i integer;
    Parent_ID bigint;
  BEGIN
     -- Open the cursor
    Open cur;   
    LOOP
      FETCH cur INTO rec;
      -- exit when no more row to fetch
      EXIT WHEN NOT FOUND;
            
      INSERT INTO Identity_Info 
       (isPerson, Last_Name, Birthday, Contact_ID, Alternate_ID)
      SELECT rec.isPerson, rec.Last_Name, rec.Birthday, rec.Contact_ID, rec.Alternate_ID      
      
      ON CONFLICT (Alternate_ID) 
      DO UPDATE SET 
        Last_Name = EXCLUDED.Last_Name, 
        Birthday = EXCLUDED.Birthday, 
        Contact_ID = EXCLUDED.Contact_ID, 
        Alternate_ID = EXCLUDED.Alternate_ID
      ;    
      
    END LOOP;  
   -- Close the cursora
    CLOSE cur;
  END; $$ ;  

  INSERT INTO Customer (
   IIID, dosri, Customer_Alt_ID, Central_Office_ID, 
   CID, Debit_Limit, Credit_Limit, 
   Date_Entry, Last_Activity_Date, Classification_ID, Customer_Group_ID, 
   Office_ID , Restriction_ID, Risk_Class_ID, Status_ID , 
   Sub_Classification_ID, Other_Info 
   ) 
 SELECT 
   ii.ID IIID, a.dosri, a.Customer_Alt_ID, Central.ID Central_Office_ID, 
   a.CID, cast(a.Debit_Limit as double precision), cast(a.Credit_Limit as double precision), 
   cast(a.Date_Entry as Date), cast(a.Last_Activity_Date as Date), cls.ID Classification_ID, 
   grp.ID Customer_Group_ID, 
   Office.ID Office_ID , rst.ID Restriction_ID, rsk.ID Risk_Class_ID, stat.ID Status_ID , 
   sc.ID Sub_Classification_ID, cast(a.Other_Info  as jsonb)
   FROM (Values
      ('100',TRUE,'1','1',1,'Active','Member',null,'Not Restricted','Risk 1','G10001',0,0,'2019/01/01','2019/01/01',FALSE,'1',null),
      ('10001',FALSE,'10002','10002',10001,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10001',null),
      ('10002',FALSE,'10002','10002',10002,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10002',null),
      ('10003',FALSE,'10002','10002',10003,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10003',null),
      ('10004',FALSE,'10002','10002',10004,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10004',null),
      ('10005',FALSE,'10002','10002',10005,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10005',null),
      ('10006',FALSE,'10002','10002',10006,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10006',null),
      ('10007',FALSE,'10002','10002',10007,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10007',null),
      ('10008',FALSE,'10002','10002',10008,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10008',null),
      ('10009',FALSE,'10002','10002',10009,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10009',null),
      ('10010',FALSE,'10002','10002',10010,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10010',null),
      ('10011',FALSE,'10002','10002',10011,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10011',null),
      ('10012',FALSE,'10002','10002',10013,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10013',null),
      ('10013',FALSE,'10002','10002',10014,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10014',null),
      ('10014',FALSE,'10002','10002',10015,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10015',null),
      ('10015',FALSE,'10002','10002',10017,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10017',null),
      ('10016',FALSE,'10002','10002',10018,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10018',null),
      ('10017',FALSE,'10002','10002',10019,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10019',null),
      ('10018',FALSE,'10002','10002',10020,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10020',null),
      ('10019',FALSE,'10002','10002',10021,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10021',null),
      ('10020',FALSE,'10002','10002',10022,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10022',null),
      ('10021',FALSE,'10002','10002',10023,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10023',null),
      ('10022',FALSE,'10002','10002',10024,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10024',null),
      ('10023',FALSE,'10002','10002',10025,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10025',null),
      ('10024',FALSE,'10002','10002',10026,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10026',null),
      ('10025',FALSE,'10002','10002',10027,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10027',null),
      ('10026',FALSE,'10002','10002',10028,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10028',null),
      ('10027',FALSE,'10002','10002',10029,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10029',null),
      ('10028',FALSE,'10002','10002',10030,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10030',null),
      ('10029',FALSE,'10002','10002',10031,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10031',null),
      ('10030',FALSE,'10002','10002',10032,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10032',null),
      ('10031',FALSE,'10002','10002',10033,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10033',null),
      ('10032',FALSE,'10002','10002',10034,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10034',null),
      ('10033',FALSE,'10002','10002',10035,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10035',null),
      ('10034',FALSE,'10002','10002',10036,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10036',null),
      ('10035',FALSE,'10002','10002',10037,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10037',null),
      ('10036',FALSE,'10002','10002',10038,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10038',null),
      ('10037',FALSE,'10002','10002',10039,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10039',null),
      ('10038',FALSE,'10002','10002',10040,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10040',null),
      ('10039',FALSE,'10002','10002',10041,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10041',null),
      ('10040',FALSE,'10002','10002',10042,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10042',null),
      ('10041',FALSE,'10002','10002',10043,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10043',null),
      ('10042',FALSE,'10002','10002',10044,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10044',null),
      ('10043',FALSE,'10002','10002',10045,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10045',null),
      ('10044',FALSE,'10002','10002',10046,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10046',null),
      ('10045',FALSE,'10002','10002',10047,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10047',null),
      ('10046',FALSE,'10002','10002',10048,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10048',null),
      ('10047',FALSE,'10002','10002',10049,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10049',null),
      ('10048',FALSE,'10002','10002',10050,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10050',null),
      ('10049',FALSE,'10002','10002',10051,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10051',null),
      ('100',TRUE,'10002','10002',10052,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10052',null),
      ('101',TRUE,'10002','10002',10053,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10053',null),
      ('1001',TRUE,'10002','10002',10054,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10054',null),
      ('1002',TRUE,'10002','10002',10055,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10055',null),
      ('1003',TRUE,'10002','10002',10056,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10056',null),
      ('1004',TRUE,'10002','10002',10057,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10057',null),
      ('1005',TRUE,'10002','10002',10058,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10058',null),
      ('1006',TRUE,'10002','10002',10059,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10059',null),
      ('1007',TRUE,'10002','10002',10060,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10060',null),
      ('1008',TRUE,'10002','10002',10061,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10061',null),
      ('1009',TRUE,'10002','10002',10062,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10062',null),
      ('1010',TRUE,'10002','10002',10063,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10063',null),
      ('1011',TRUE,'10002','10002',10064,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10064',null),
      ('1012',TRUE,'10002','10002',10065,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10065',null),
      ('1013',TRUE,'10002','10002',10066,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10066',null),
      ('1014',TRUE,'10002','10002',10067,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10067',null),
      ('1015',TRUE,'10002','10002',10068,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10068',null),
      ('1016',TRUE,'10002','10002',10069,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10069',null),
      ('1017',TRUE,'10002','10002',10070,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10070',null),
      ('1018',TRUE,'10002','10002',10071,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10071',null),
      ('1019',TRUE,'10002','10002',10072,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10072',null),
      ('1020',TRUE,'10002','10002',10073,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10073',null),
      ('1021',TRUE,'10002','10002',10074,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10074',null),
      ('1022',TRUE,'10002','10002',10075,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10075',null),
      ('1023',TRUE,'10002','10002',10076,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10076',null),
      ('1024',TRUE,'10002','10002',10077,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10077',null),
      ('1025',TRUE,'10002','10002',10078,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10078',null),
      ('1026',TRUE,'10002','10002',10079,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10079',null),
      ('1027',TRUE,'10002','10002',10080,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10080',null),
      ('1028',TRUE,'10002','10002',10081,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10081',null),
      ('1029',TRUE,'10002','10002',10082,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10082',null),
      ('1030',TRUE,'10002','10002',10083,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10083',null),
      ('1031',TRUE,'10002','10002',10084,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10084',null),
      ('1032',TRUE,'10002','10002',10085,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10085',null),
      ('1033',TRUE,'10002','10002',10086,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10086',null),
      ('1034',TRUE,'10002','10002',10087,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10087',null),
      ('1035',TRUE,'10002','10002',10088,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10088',null),
      ('1036',TRUE,'10002','10002',10089,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10089',null),
      ('1037',TRUE,'10002','10002',10090,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10090',null),
      ('1038',TRUE,'10002','10002',10091,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10091',null),
      ('1039',TRUE,'10002','10002',10092,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10092',null),
      ('1040',TRUE,'10002','10002',10093,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10093',null),
      ('1041',TRUE,'10002','10002',10094,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10094',null),
      ('1042',TRUE,'10002','10002',10095,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10095',null),
      ('1043',TRUE,'10002','10002',10096,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10096',null),
      ('1044',TRUE,'10002','10002',10097,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10097',null),
      ('1045',TRUE,'10002','10002',10098,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10098',null),
      ('1046',TRUE,'10002','10002',10099,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10099',null),
      ('1047',TRUE,'10002','10002',10100,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10100',null),
      ('1048',TRUE,'10002','10002',10101,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10101',null),
      ('1049',TRUE,'10002','10002',10102,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10102',null),
      ('1050',TRUE,'10002','10002',10103,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10103',null),
      ('1051',TRUE,'10002','10002',10104,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10104',null),
      ('1052',TRUE,'10002','10002',10105,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10105',null),
      ('1053',TRUE,'10002','10002',10106,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10106',null),
      ('1054',TRUE,'10002','10002',10107,'Active','Member',null,null,null,null,0,0,'2019/01/01','2019/01/01',FALSE,'10107',null)
      )   
    a(Alternate_ID, isPerson,
      CentralOfficealt, Officealt, CID, Status, Classification, 
      Sub_Classification, Restriction, Risk_Class, Customer_GroupaltID, 
      Debit_Limit, Credit_Limit, Date_Entry, Last_Activity_Date, dosri, Customer_Alt_ID, Other_Info)     
          
  LEFT JOIN Identity_Info ii      on ii.Alternate_ID = a.Alternate_ID
  LEFT JOIN Customer cus          on cus.Customer_Alt_ID = a.Alternate_ID
  
  LEFT JOIN vwReference stat      on lower(stat.Title)    = lower(a.Status)         and stat.Ref_Type = 'CustomerStatus'
  LEFT JOIN vwReference cls       on lower(cls.Title)     = lower(a.Classification) and cls.Ref_Type = 'CustomerClass'
  LEFT JOIN vwReference sc        on lower(sc.Title)      = lower(a.Sub_Classification)    and sc.Ref_Type = 'CustomerSubClass'  
  LEFT JOIN vwReference rst       on lower(rst.Title)     = lower(a.Restriction)    and rst.Ref_Type = 'CustomerRestriction'
  LEFT JOIN vwReference rsk       on lower(rsk.Title)     = lower(a.Risk_Class)    and rst.Ref_Type = 'Risk_Class'
  LEFT JOIN Office  on Office.Alternate_ID  = a.CentralOfficealt
  LEFT JOIN Office Central on Central.Alternate_ID = a.Officealt
  LEFT JOIN Customer_Group grp on grp.ID = cus.Customer_Group_ID  
  ON CONFLICT(Central_Office_ID, CID) 
  DO UPDATE SET
    IIID = excluded.IIID, 
    dosri = excluded.dosri, 
    Customer_Alt_ID = excluded.Customer_Alt_ID, 
    Central_Office_ID = excluded.Central_Office_ID, 
    CID = excluded.CID, 
    Debit_Limit = excluded.Debit_Limit, 
    Credit_Limit = excluded.Credit_Limit, 
    Date_Entry = excluded.Date_Entry, 
    Last_Activity_Date = excluded.Last_Activity_Date, 
    Classification_ID = excluded.Classification_ID, 
    Customer_Group_ID = excluded.Customer_Group_ID, 
    Office_ID  = excluded.Office_ID , 
    Restriction_ID = excluded.Restriction_ID, 
    Risk_Class_ID = excluded.Risk_Class_ID, 
    Status_ID  = excluded.Status_ID , 
    Sub_Classification_ID = excluded.Sub_Classification_ID, 
    Other_Info  = excluded.Other_Info 
  ; 

*/
