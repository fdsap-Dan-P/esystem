----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Customer_Group (
----------------------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Central_Office_ID bigint NOT NULL,
  Code varchar(15) NOT NULL,
  Type_ID bigint NOT NULL,
  Group_Name varchar(200) NOT NULL,
  Short_Name varchar(30),
  Date_Stablished Date NULL,
  Meeting_Day int2 NULL,
  Office_ID bigint NOT NULL,
  Officer_ID bigint NULL,
  Parent_ID bigint NULL,
  Alternate_ID VarChar(30) NULL,
  Address_Detail varchar(150) NULL,
  Address_URL varchar(200) NULL,
  Geography_ID bigint NULL,  
  Other_Info jsonb NULL,
  
  CONSTRAINT Customer_Group_pkey PRIMARY KEY (ID),
  CONSTRAINT idxCustomer_Group_CID UNIQUE (Type_ID, Central_Office_ID, Code),
  CONSTRAINT idxCustomer_Group_Alt UNIQUE (Alternate_ID),

  CONSTRAINT fk_Customer_Group_CentralOffice FOREIGN KEY (Central_Office_ID) REFERENCES Office(ID),
  CONSTRAINT fk_Customer_Group_Customer_GroupType FOREIGN KEY (Type_ID) REFERENCES Reference(ID),
  CONSTRAINT fk_Customer_Group_Office FOREIGN KEY (Office_ID ) REFERENCES Office(ID),
  CONSTRAINT fk_Customer_Group_Officer FOREIGN KEY (Officer_ID) REFERENCES Employee(ID),
  CONSTRAINT fk_Customer_Group_Parent FOREIGN KEY (Parent_ID) REFERENCES Customer_Group(ID),
  CONSTRAINT fk_Customer_Group_Geography FOREIGN KEY (Geography_ID) REFERENCES Geography(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxCustomer_Group_UUID ON public.Customer_Group(UUID);

-- UPDATE Customer_Group Trigger
---------------------------------------------------------------------------
CREATE or REPLACE FUNCTION Customer_Group_UPDATE() returns trigger AS 
---------------------------------------------------------------------------
$$ BEGIN
   NEW.Mod_Ctr := Nextval('Customer_Group_Mod_Ctr_seq');
   NEW.Updated := CURRENT_TIMESTAMP;
   
   RETURN NEW;
END $$ Language plpgsql;

DROP TRIGGER IF EXISTS trgCustomer_Group_Ins on Customer_Group;
---------------------------------------------------------------------------
CREATE TRIGGER trgCustomer_Group_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Customer_Group
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgCustomer_Group_upd on Customer_Group;
---------------------------------------------------------------------------
CREATE TRIGGER trgCustomer_Group_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Customer_Group
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgCustomer_Group_del on Customer_Group;
---------------------------------------------------------------------------
CREATE TRIGGER trgCustomer_Group_del
---------------------------------------------------------------------------
    AFTER DELETE ON Customer_Group
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();


----------------------------------------------------------------------------------------
CREATE OR REPLACE VIEW public.vwCustomer_Group
----------------------------------------------------------------------------------------
AS SELECT 
    grp.ID, mr.UUID,

    co.ID Central_Office_ID, co.Office_Name CentralName,

    grp.Code,
    grp.Group_Name,
    grp.Date_Stablished,
    grp.Alternate_ID,
    
    gType.ID Type_ID, gType.UUID TypeUUID, gType.Title AS GroupType,
    
    o.ID Office_ID , o.Office_Name,
    
    ii.ID OfficerIIID, ii.UUID OfficerUUID, emp.ID Officer_ID, emp.Employee_No,
    ii.Last_Name OfficerLast_Name, ii.First_Name OfficerFirst_Name, ii.BirthDay OfficerBirthday,

    grp.Address_Detail, grp.Address_URL,

    vgeo.ID Geography_ID, vgeo.UUID GeographyUUID, vgeo.Location, vgeo.Full_Location,

    mr.Mod_Ctr, 
    grp.Other_Info,
    mr.Created,
    mr.Updated 
   FROM Customer_Group grp
   INNER JOIN Main_Record mr on mr.UUID = grp.UUID
   JOIN Reference gType ON gType.ID = grp.Type_ID
   JOIN Office o  ON o.ID = grp.Office_ID 
   JOIN Office co ON co.ID = grp.Office_ID 
   JOIN Employee emp on emp.ID = grp.Officer_ID
   JOIN Identity_Info ii ON ii.ID = emp.IIID
   
   LEFT JOIN Customer_Group p ON p.ID = o.Parent_ID
   LEFT JOIN vwGeography vgeo ON vgeo.ID = o.Geography_ID;

-------------------------------------------------------------------------------------
INSERT INTO Customer_Group

  (Central_Office_ID, Code, Type_ID, Short_Name, 
   Group_Name, Date_Stablished, Meeting_Day,
   Office_ID , Officer_ID, Parent_ID, Alternate_ID, 
   Address_Detail, Address_URL, Geography_ID, Other_Info)

   SELECT 
     Cen.ID Central_Office_ID, a.Code, typ.ID Type_ID, a.Short_Name, 
     a.Group_Name, cast(a.Date_Stablished as Date) Date_Stablished, a.Meeting_Day, 
     Ofc.ID Office_ID , Ofr.ID Officer_ID, p.ID Parent_ID, a.Alternate_ID, 
     a.Address_Detail, a.Address_URL, f.ID Geography_ID, NULL
  FROM (Values
      ('10001','001','Center','Center 1','Center 1','01/01/2020',1,'10002','97-0114',NULL,'10000','Brgy. Soledad, San Pablo City',NULL,'Soledad San Pablo City')
      )   
    a(CentralAlt, Code, GroupType, Short_Name, Group_Name, Date_Stablished, Meeting_Day, 
      OfficeAlt, Employee_No, ParentAlt, Alternate_ID, Address_Detail, Address_URL, Geography)  
      
  INNER JOIN Office Cen on Cen.Alternate_ID = a.CentralAlt  
  INNER JOIN Office Ofc on Ofc.Alternate_ID = a.OfficeAlt  
  INNER JOIN vwReference typ on lower(typ.Title) = lower(a.GroupType) 
    and lower(typ.Ref_Type) = 'customergrouptype' 

  LEFT JOIN Employee Ofr   on Ofr.Employee_No = a.Employee_No
  LEFT JOIN Customer_Group p   on p.Alternate_ID  = a.ParentAlt
  LEFT JOIN Office on Office.Alternate_ID = a.OfficeAlt
  LEFT JOIN LATERAL searchLocation(a.Geography,1) f ON true  
  ON CONFLICT (Alternate_ID) DO 
  UPDATE SET
     Central_Office_ID = excluded.Central_Office_ID, 
     Code = excluded.Code, 
     Type_ID = excluded.Type_ID, 
     Short_Name = excluded.Short_Name, 
     Group_Name = excluded.Group_Name, 
     Date_Stablished = excluded.Date_Stablished, 
     Office_ID  = excluded.Office_ID , 
     Officer_ID = excluded.Officer_ID, 
     Parent_ID = excluded.Parent_ID, 
     Alternate_ID = excluded.Alternate_ID, 
     Address_Detail = excluded.Address_Detail, 
     Address_URL = excluded.Address_URL, 
     Geography_ID = excluded.Geography_ID, 
     Other_Info = excluded.Other_Info;    
