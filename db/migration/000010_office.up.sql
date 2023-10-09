----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Office (
----------------------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Code varchar(15) NOT NULL,
  Short_Name varchar(30) NOT NULL,
  Office_Name varchar(200) NOT NULL,
  Date_Stablished Date NULL,
  Type_ID bigint NOT NULL,
  Parent_ID bigint NULL,
  Officer_IIID bigint NULL,
  Alternate_ID VarChar(30) NULL,
  Address_Detail varchar(150) NULL,
  Address_URL varchar(200) NULL,
  Geography_ID bigint NULL,  
  CID_Sequence bigint NULL,
  Other_Info jsonb NULL,

  CONSTRAINT Office_pkey PRIMARY KEY (ID),
  CONSTRAINT ndxOfficealt UNIQUE (Alternate_ID),
  CONSTRAINT fkOfficeOfficeType FOREIGN KEY (Type_ID) REFERENCES Reference(ID),
  CONSTRAINT fkOfficeParent FOREIGN KEY (Parent_ID) REFERENCES Office(ID),
--  CONSTRAINT fkOfficeOfficer FOREIGN KEY (Officer_IIID) REFERENCES Identity_info(ID),
  CONSTRAINT fkofficceGeography FOREIGN KEY (Geography_ID) REFERENCES Geography(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxOffice_UUID ON public.Office(UUID);
CREATE UNIQUE INDEX IF NOT EXISTS idxOfficeCode ON public.Office USING btree (COALESCE(Parent_ID,0), lower(Code));
CREATE UNIQUE INDEX IF NOT EXISTS idxOfficeShort_Name ON public.Office USING btree (COALESCE(Parent_ID,0), lower(Short_Name));

DROP TRIGGER IF EXISTS trgOfficeIns on Office;
---------------------------------------------------------------------------
CREATE TRIGGER trgOfficeIns
---------------------------------------------------------------------------
    BEFORE INSERT ON Office
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();

---------------------------------------------------------------------------
DROP TRIGGER IF EXISTS trgOfficeupd on Office;
CREATE TRIGGER trgOfficeupd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Office
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgOffice_del on Office;
---------------------------------------------------------------------------
CREATE TRIGGER trgOffice_del
---------------------------------------------------------------------------
    AFTER DELETE ON Office
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();

----------------------------------------------------------------------------------------
CREATE OR REPLACE VIEW public.vwOffice
----------------------------------------------------------------------------------------
AS SELECT 
    o.ID, mr.UUID,
    o.Code,
    o.Short_Name,
    o.Office_Name,
    o.Date_Stablished,
 
    oType.ID Type_ID, oType.UUID OfficeTypeUUID, oType.Title OfficeType,
 
    p.ID Parent_ID, p.UUID ParentUUID, p.Code ParentCode,
    p.Short_Name ParentShort_Name, p.Office_Name ParentName,
    
    o.Alternate_ID,
    o.Address_Detail,
    o.Address_URL,
  
    vgeo.ID Geography_ID, vgeo.UUID GeographyUUID,
    vgeo.Location,
    vgeo.Full_Location,
    
    mr.Mod_Ctr,
    o.Other_Info,
    mr.Created,
    mr.Updated
   FROM Office o
   INNER JOIN Main_Record mr on mr.UUID = o.UUID
   INNER JOIN Reference oType ON oType.ID = o.Type_ID
   LEFT JOIN Office p ON p.ID = o.Parent_ID
   LEFT JOIN vwGeography vgeo ON vgeo.ID = o.Geography_ID;
   
----------------------------------------------------------------------------------------
  INSERT INTO Office(Code, Short_Name, Office_Name, Type_ID)
  SELECT 
    '0' Code, 'System' Short_Name, 'System' Office_Name, r.ID Type_ID
  FROM vwReference r 
  WHERE r.Title = 'System' and r.Ref_Type = 'OfficeType'
  ON CONFLICT (Alternate_ID)
  DO UPDATE SET
    Code = excluded.Code, 
    Short_Name = excluded.Short_Name, 
    Office_Name = excluded.Office_Name, 
    Type_ID = excluded.Type_ID;   
      
  
---------------------------------------------------------------
DO $$
---------------------------------------------------------------
DECLARE 
  Titles TEXT DEFAULT '';
  rec   RECORD;
  cur CURSOR FOR 

  SELECT a.Office_Name, cast(a.Date_Stablished as Date) Date_Stablished,  
     a.Office_Address, a.Address_Detail,  a.Address_URL, offTyp.ID Type_ID,   
     a.Alternate_ID, 
     a.Code,   a.Short_Name, Parentalt
  FROM 
    (Values
      ('Generic',null,'Philippines','Philippines',null,null,'1','Institution','null',null,null,'9999', 'Generic'),
      ('United Church of Christ in the Philippines','1948/05/25','West Triangle, Quezon City. Metro Manila','877 EDSA, West Triangle, Quezon City. Metro Manila','https://www.google.com/Maps/@14.6462099,121.0360189,3a,75y,239.26h,86.91t/data=!3m6!1e1!3m4!1s1PFeC-7yipSjCZOmoFP8dg!2e0!7i13312!8i6656','1067','10001','Institution','null',null,null,'0000', 'UCCP'),
      ('South Luzon Jurisdictional Area',null,'Los Baños, Laguna','#9 Kanluran Road, College, Los Banos, Laguna','https://www.google.com/Maps/@14.1672649,121.2405257,3a,75y,339.28h,80.44t/data=!3m6!1e1!3m4!1sJ708bLw0rcP9IneTpkuLBw!2e0!7i13312!8i6656','1054','10002','Institution','null','10001','1054','0000', 'SLJ'),
      ('Batangas Associate Conference',null,'Lipa City','28 B. Morada Avenue, Lipa, Batangas','https://www.google.com/Maps/@13.9410034,121.158245,0a,82.2y,144.39h,96.49t/data=!3m4!1e1!3m2!1sHt4XitK9nN2MapFCRuXyRA!2e0?Source=apiv3','1072','10003','Institution',null,'10001',null,'0001', 'BAC'),
      ('BPI - Los Banos',null,'Batong Malake, Los Baños City, Laguna','9332, Lopez Avenue, Barangay Batong Malake, Los Baños City, Laguna',null,null,'10004','Institution',null,null,null,'0002', 'BPI - LB'),
      ('BPI - New Mia',null,'City of Parañaque, Metro Manila','Ipt Bldg, Arrival Lobby, Naia Terminal 1, City of Parañaque, Metro Manila',null,null,'10005','Institution',null,null,null,'0003', 'BPI - New Mia'),
      ('EBF',null,'Quezon City, 1104 Metro Manila','879 Epifanio de los Santos Ave, Quezon City, 1104 Metro Manila',null,'1068','10006','Institution',null,null,null,'0004', 'EBF'),
      ('IDPIP-ST',null,'Batong Malake, Los Baños, Laguna','10051 MT. Data, Umali Subd., Batong Malake, Los Baños, Laguna',null,'1064','10007','Institution',null,'10002',null,'0005', 'IDPIP-ST'),
      ('Kowloon Union Church, Hongkong',null,' Kwun Chung, Hong Kong','4 Jordan Rd, Kwun Chung, Hong Kong',null,'1062','10008','Institution',null,'10002',null,'0006', 'KUC, Hongkong'),
      ('LCSMC',null,'Malate, Manila, 1004 Metro Manila','1661 Vasquez St, Malate, Manila, 1004 Metro Manila','https://www.google.com/Maps/@14.574546,120.9873325,3a,60y,59.63h,87.11t/data=!3m6!1e1!3m4!1sEF33csD5NRunEWqAk1XYzA!2e0!7i13312!8i6656','1006','10009','Institution',null,'10002',null,'0007', 'LCSMC'),
      ('North Bicol Conference',null,'San Jose, San Jose City, OcCIDental Mindoro','Aurora Quezon St., San Jose, San Jose City, OcCIDental Mindoro','https://www.google.com/Maps/@12.3533273,121.071089,3a,75y,219.41h,85.77t/data=!3m6!1e1!3m4!1ssEd4huRj8H-ebteOT4PA!2e0!7i13312!8i6656','1061','10010','Institution',null,'10002',null,'0009', 'Occ. Mindoro AC'),
      ('OcCIDental Mindoro Assoc. Conference',null,'Brgy. Bethel, Victoria, Oriental Mindoro','Bethel, Brgy. Bethel, Victoria, Oriental Mindoro','https://www.google.com/Maps/@13.1463288,121.251938,3a,82.2y,7.23h,69.42t/data=!3m6!1e1!3m4!1sBH-x3JBXbt5qzAIkH0zHBw!2e0!7i13312!8i6656','1011','10011','Institution',null,'10002',null,'0010', 'Occ. Mindoro Conf.'),
      ('UCCP Palawan Associate Conference',null,'Magara, Roxas, Palawan','Magara, Roxas, Palawan','https://www.google.com/Maps/@10.3036729,119.2645837,3a,66.6y,130.41h,80.27t/data=!3m6!1e1!3m4!1sHmT732qANcSVTpjxqLT5Q!2e0!7i13312!8i6656','1059','10012','Institution',null,'10002',null,'0011', 'Palawan AC'),
      ('PAC - UCCP Magara',null,'Brgy. San Miguel, Puerto Princesa, Palawan','Wescom Rd, Brgy. San Miguel, Puerto Princesa, Palawan','https://www.google.com/Maps/@9.7524795,118.7661807,0a,82.2y,311.09h,82.77t/data=!3m4!1e1!3m2!1s8DQjMvxYQj8yTpP4ojYwrg!2e0?Source=apiv3','1071','10013','Institution',null,'10002',null,'0012', 'PAC UCCP Magara'),
      ('PAC District V',null,'Calamba City, Laguna','Calamba City, Laguna',null,'1059','10014','Institution',null,'10002',null,'0013', 'PAC District V'),
      ('Pastive Arce',null,'Lusacan, Tiaong, Quezon','STC Cornerstone, Lusacan, Tiaong, Quezon','https://www.google.com/Maps/@13.9628722,121.3211301,0a,82.2y,210.33h,86.12t/data=!3m4!1e1!3m2!1s17lxe63AuXOTXXs0dkiOFQ!2e0?Source=apiv3','1069','10015','Institution',null,'10002',null,'0015', 'Pastive Arce'),
      ('South Bicol Conference',null,'San Pablo City, Laguna','Peñalosa St. San Pablo City, Laguna','https://www.google.com/Maps/@14.0652933,121.3235006,0a,82.2y,332.31h,90.61t/data=!3m4!1e1!3m2!1sC1ho4OvyA-SuWDfB37QfIA!2e0?Source=apiv3','1024','10016','Institution',null,'10002',null,'0016', 'South Bicol Conf.'),
      ('Southern Tagalog Conference',null,'Amadeo, Cavite','Talon Rd, Amadeo, Cavite','https://www.google.com/Maps/@14.1345142,120.9399264,0a,82.2y,297.89h,84.46t/data=!3m4!1e1!3m2!1sa-0IfeiihWO5E4i0OR6Hqg!2e0?Source=apiv3','1063','10017','Institution',null,'10002',null,'0017', 'Southern Tagalog Conf.'),
      ('STC - FEC',null,'Los Baños, Laguna','UCCP Bambang, Bambang Highway, Los Baños, Laguna','https://www.google.com/Maps/@14.1767695,121.2198035,0a,82.2y,353.06h,93.66t/data=!3m4!1e1!3m2!1sRqObxTroK69XRlJF9y19Eg!2e0?Source=apiv3','1063','10018','Institution',null,'10002',null,'0018', 'STC - FEC'),
      ('UCCP Amadeo',null,'Poblacion, Batangas, 4200 Batangas','17 C.Tirona, Poblacion, Batangas, 4200 Batangas','https://www.google.com/Maps/@13.7587154,121.0599069,0a,66.6y,264.17h,80.72t/data=!3m4!1e1!3m2!1sUB0-rftb9zUMCFpBPIOyvg!2e0?Source=apiv3','1004','10019','Institution',null,'10002',null,'0019', 'UCCP Amadeo'),
      ('UCCP Bambang - OMC',null,'Tagum, Davao del Norte','315, Mabini St, Tagum, Davao del Norte','https://www.google.com/Maps/@7.4473733,125.8003235,0a,82.2y,322.55h,100.02t/data=!3m4!1e1!3m2!1sJ7HaDlfZPvp2HEf9Zw9ZRQ!2e0?Source=apiv3','1065','10020','Institution',null,'10002',null,'0020', 'UCCP Bambang - OMC'),
      ('UCCP Batangas',null,'Buli, Pinamalayan, Oriental Mindoro','Buli, Pinamalayan, Oriental Mindoro','https://www.google.com/Maps/@13.0667505,121.4980613,0a,66.6y,256.51h,85.89t/data=!3m4!1e1!3m2!1sfuil1lS-B7AdY3n59sBiUw!2e0?Source=apiv3','1073','10021','Institution',null,'10002',null,'0021', 'UCCP Batangas'),
      ('UCCP Bethel',null,null,'unknown',null,'1060','10022','Institution',null,'10002',null,'0022', 'UCCP Bethel'),
      ('UCCP Buli - OMC',null,'Los Baños, Laguna','#9 Kanluran Road, College, Los Banos, Laguna','https://www.google.com/Maps/@14.1672649,121.2405257,3a,75y,339.28h,80.44t/data=!3m6!1e1!3m4!1sJ708bLw0rcP9IneTpkuLBw!2e0!7i13312!8i6656','1011','10023','Institution',null,'10002',null,'0023', 'UCCP Buli - OMC'),
      ('UCCP Campung Alay',null,'Culasisi, Sablayan, Oriental Mindoro','Culasisi, Sablayan, OcCIDental Mindoro','https://www.google.com/Maps/@12.8581433,120.8926738,0a,66.6y,309.76h,80.75t/data=!3m4!1e1!3m2!1ssGoc88Hb2O3lGjMA1Zp1pg!2e0?Source=apiv3',null,'10024','Institution',null,'10002',null,'0024', 'UCCP Campung Alay'),
      ('UCCP Church Among the Palms',null,'Rizal, Palawan','Danum-danum, Rizal, Palawan','https://www.google.com/Maps/@7.4478307,125.8003397,3a,75y,188h,69.41t/data=!3m6!1e1!3m4!1sVilb9Hxm2IL9TTQNNFkV0w!2e0!7i13312!8i6656?hl=fil','1070','10025','Institution',null,'10002',null,'0025', 'UCCP Church Among the Palms'),
      ('UCCP Culasisi',null,'Labasan, Bongabong, Oriental Mindoro ','Strong Republic Nautical Hwy, Labasan, Bongabong, Oriental Mindoro ','https://www.google.com/Maps/@12.7782126,121.469986,0a,82.2y,294.36h,70.11t/data=!3m4!1e1!3m2!1szaakmdVAlIfaRW5yptpDuA!2e0?Source=apiv3','1011','10026','Institution',null,'10002',null,'0026', 'UCCP Culasisi'),
      ('UCCP Danum-danum',null,'Manila, Metro manila','1661 Vasquez St, Malate, Manila, 1004 Metro Manila','https://www.google.com/Maps/@14.5747202,120.9872387,0a,82.2y,15.21h,97.14t/data=!3m4!1e1!3m2!1s1EGvfpcYNuaa9xpdIee5IA!2e0?Source=apiv3','1059','10027','Institution',null,'10002',null,'0027', 'UCCP Danum-danum'),
      ('UCCP Labasan  - OMC',null,'Lipa, Batangas','28 B. Morada Avenue, Lipa, Batangas','https://www.google.com/Maps/@13.9410034,121.158245,0a,82.2y,144.39h,96.49t/data=!3m4!1e1!3m2!1sHt4XitK9nN2MapFCRuXyRA!2e0?Source=apiv3','1011','10028','Institution',null,'10002',null,'0028', 'UCCP Labasan  - OMC'),
      ('UCCP LCSMC',null,'Rizal, Palawan','Malapandig, Rizal, Palawan',null,'1006','10029','Institution',null,'10002',null,'0029', 'UCCP LCSMC'),
      ('UCCP Lipa Evangelical church',null,'West Triangle, Quezon City. Metro Manila','877 EDSA, West Triangle, Quezon City. Metro Manila','https://www.google.com/Maps/@14.6462099,121.0360189,3a,75y,239.26h,86.91t/data=!3m6!1e1!3m4!1s1PFeC-7yipSjCZOmoFP8dg!2e0!7i13312!8i6656','1072','10030','Institution',null,'10002',null,'0030', 'UCCP Lipa Evangelical church'),
      ('UCCP Malapandig',null,'Los Baños, Laguna','#9 Kanluran Road, College, Los Banos, Laguna','https://www.google.com/Maps/@14.1672649,121.2405257,3a,75y,339.28h,80.44t/data=!3m6!1e1!3m4!1sJ708bLw0rcP9IneTpkuLBw!2e0!7i13312!8i6656','1059','10031','Institution',null,'10002',null,'0031', 'UCCP Malapandig'),
      ('UCCP National Office',null,'Peñafrancia, Naga City','1- Peñafrancia St., Peñafrancia, Naga City','https://www.google.com/Maps/@13.6247431,123.2039188,3a,60y,54.07h,81.84t/data=!3m6!1e1!3m4!1slloRygB3buiK4oQerIC1Ug!2e0!7i13312!8i6656','1067','10032','Institution',null,'10002',null,'0032', 'UCCP National Office'),
      ('UCCP NESTCON',null,'San Pedro, Laguna','Almasiga St, Pacita, San Pedro, Laguna','https://www.google.com/Maps/@14.3441138,121.0556519,3a,82.2y,170.06h,91.44t/data=!3m6!1e1!3m4!1su0pUg2GyJYwAIEPlUKDxqw!2e0!7i13312!8i6656','1057','10033','Institution',null,'10002',null,'0033', 'UCCP NESTCON'),
      ('UCCP North Bicol Conference',null,'Pagsanjan, Laguna','Pagsanjan, Laguna','https://www.google.com/Maps/@14.2749298,121.4531951,3a,66.6y,90t/data=!3m6!1e1!3m4!1sHt36GMDn3ac9Hygd5Z1ow!2e0!7i13312!8i6656','1012','10034','Institution',null,'10002',null,'0034', 'UCCP North Bicol Conference'),
      ('UCCP PACITA',null,'Panitian, Quezon, Palawan','Brgy. Panitian, Quezon, Palawan','https://www.google.com/Maps/@9.2381267,118.0270671,0a,82.2y,78h,70.77t/data=!3m4!1e1!3m2!1sXreK2SHQS9EcVYFiEX79mQ!2e0?Source=apiv3','1003','10035','Institution',null,'10002',null,'0035', 'UCCP PACITA'),
      ('UCCP Pagsanjan',null,'Pinamalayan','Brgy. Pinamalayan, Oriental Mindoro','https://www.google.com/Maps/@13.0667505,121.4980613,0a,66.6y,256.51h,85.89t/data=!3m4!1e1!3m2!1sfuil1lS-B7AdY3n59sBiUw!2e0?Source=apiv3','1066','10036','Institution',null,'10002',null,'0036', 'UCCP Pagsanjan'),
      ('UCCP Panitian',null,'Salcedo, Bansud, Oriental Mindoro','Brgy. Salcedo, Bansud, Oriental Mindoro',null,'1059','10037','Institution',null,'10002',null,'0037', 'UCCP Panitian'),
      ('UCCP Pinamalayan - OMC',null,'San Buenaventura, Luisiana, Laguna','Brgy. San Buenaventura, Luisiana, Laguna','https://www.google.com/Maps/@14.1861381,121.5105019,0a,82.2y,156.69h,89.75t/data=!3m4!1e1!3m2!1sZIxeIy7ju0xT9XKLL53aFA!2e0?Source=apiv3','1011','10038','Institution',null,'10002',null,'0038', 'UCCP Pinamalayan - OMC'),
      ('UCCP Salcedo - OMC',null,'Poblacion, San Vicente, Palawan','Brgy. Poblacion, San Vicente, Palawan','https://www.google.com/Maps/@10.5184238,119.3514341,0a,66.6y,151.27h,76.06t/data=!3m4!1e1!3m2!1symBgiTC7U0sHMST4i6NoQw!2e0?Source=apiv3','1011','10039','Institution',null,'10002',null,'0039', 'UCCP Salcedo - OMC'),
      ('UCCP San Buenaventura',null,'Pangil, Laguna','Brgy. Piit, Pangil, Laguna','https://www.google.com/Maps/@14.408538,121.4741708,0a,82.2y,42.67h,101.17t/data=!3m4!1e1!3m2!1srdLSFqhtmNB-nAOhZ89vyQ!2e0?Source=apiv3','1057','10040','Institution',null,'10002',null,'0040', 'UCCP San Buenaventura'),
      ('UCCP San Vicente',null,'Sucol, Calamba, Laguna','Brgy. Sucol, Calamba, Laguna','https://www.google.com/Maps/@14.1802064,121.2000268,3a,75y,280.64h,88.45t,359.32r/data=!3m6!1e1!3m4!1ssFqa7JrEw-cJ2kxLB9tBnQ!2e0!7i13312!8i6656','1059','10041','Institution',null,'10002',null,'0041', 'UCCP San Vicente'),
      ('UCCP SKP NESTCON',null,'Taytay, Palawan','UCCP Nalbot, Taytay, Palawan',null,'1058','10042','Institution',null,'10002',null,'0042', 'UCCP SKP NESTCON'),
      ('UCCP Sucol',null,null,'unknown',null,'1074','10043','Institution',null,'10002',null,'0043', 'UCCP Sucol'),
      ('UCCP Taytay',null,'Dasmariñas City','PCU-UTS, Pala-pala, Dasmariñas, Cavite',null,'1059','10044','Institution',null,'10002',null,'0044', 'UCCP Taytay'),
      ('UCCP Tubon',null,'Victoria, Laguna','Victoria, Laguna',null,null,'10045','Institution',null,'10002',null,'0045', 'UCCP Tubon'),
      ('UCCP UCC',null,'Victoria, oriental','Victoria, Oriental Mindoro',null,'1004','10046','Institution',null,'10002',null,'0046', 'UCCP UCC'),
      ('UCCP Victoria - NESTCON',null,null,'Gustavslundsvägen 18, 167 51 Bromma, Sweden',null,'1057','10047','Institution',null,'10002',null,'0047', 'UCCP Victoria - NESTCON'),
      ('UCCP Victoria - OMC',null,null,null,null,'1011','10048','Institution',null,'10002',null,'0048', 'UCCP Victoria - OMC'),
      ('Uniting Church of Sweden',null,null,null,null,'1056','10049','Institution',null,'10002',null,'0049', 'Uniting Church of Sweden')
      )   
    a(Office_Name,  Date_Stablished,  Office_Address, Address_Detail, Address_URL,
      ContactaltID,   Alternate_ID, OfficeType,  ParentType, Parentalt,  Officeraltid,   Code,   Short_Name)      
  INNER JOIN vwReference offTyp on lower(offTyp.Title) = lower(a.OfficeType) 
    and offTyp.Ref_Type = 'OfficeType'  
  ORDER BY a.Alternate_ID
  ;

  i integer;
  Parent_ID bigint;
  Office_Address_ID bigint;
  AddressType_ID bigint;

BEGIN
   -- Open the cursor
   Open cur;
   
   LOOP

      FETCH cur INTO rec;
    -- exit when no more row to fetch
      EXIT WHEN NOT FOUND;
 
      SELECT a.ID into Office_Address_ID 
      FROM searchLocation(rec.Office_Address,1) a;
    
      SELECT r.ID into AddressType_ID 
      FROM Reference r
      inner join Reference_Type y on r.Type_ID = y.ID
      WHERE r.Title = 'Office Address' and y.Title = 'AddressType';    
    
      SELECT ID into Parent_ID
      FROM Office ii
      WHERE ii.Alternate_ID = rec.ParentAlt;
          
      INSERT INTO Office(
        Code, Short_Name, Office_Name, Date_Stablished, 
        Type_ID, Parent_ID,
        Alternate_ID, Address_Detail, Address_URL, Geography_ID)
      SELECT 
        rec.Code, rec.Short_Name, rec.Office_Name, rec.Date_Stablished, 
        AddressType_ID Type_ID, Parent_ID Parent_ID,
        rec.Alternate_ID, rec.Address_Detail, rec.Address_URL, Office_Address_ID Geography_ID
      ON CONFLICT (Alternate_ID)
      DO UPDATE SET
        Code = excluded.Code, Short_Name = excluded.Short_Name, 
        Office_Name = excluded.Office_Name, Date_Stablished = excluded.Date_Stablished, 
        Type_ID = excluded.Type_ID, 
        Parent_ID = excluded.Parent_ID, Alternate_ID = excluded.Alternate_ID, 
        Address_Detail = excluded.Address_Detail, Address_URL = excluded.Address_URL, 
        Geography_ID = excluded.Geography_ID ;   
        
   END LOOP;
  
   -- Close the cursor
   CLOSE cur;

END; $$ ;
