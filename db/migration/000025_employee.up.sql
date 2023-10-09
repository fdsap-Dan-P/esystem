----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Employee (
----------------------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  IIID bigint NOT NULL,
  Central_Office_ID bigint NOT NULL,
  Employee_No varchar(36) NOT NULL,
--   Employee_Name varchar(200) NOT NULL,
  Basic_Pay float8 NOT NULL,
  Date_Hired Date NULL,
  Date_Regular Date NULL,
  Job_Grade int2 NOT NULL,
  Job_Step int2 NOT NULL,
  Level_ID bigint NULL,
  Office_ID bigint NOT NULL,
  Position_ID bigint NOT NULL,
  Status_Code bigint NOT NULL,
  Superior_ID bigint NULL,
  Type_ID integer NOT NULL,
  Other_Info jsonb NULL,

  CONSTRAINT Employee_pkey PRIMARY KEY (ID),
  CONSTRAINT fk_Employee_CentralOffice FOREIGN KEY (Central_Office_ID) REFERENCES Office(ID),
  CONSTRAINT fk_Employee_IDentity_Info FOREIGN KEY (IIID) REFERENCES Identity_Info(ID),
  CONSTRAINT fk_Employee_Level FOREIGN KEY (Level_ID) REFERENCES Reference(ID),
  CONSTRAINT fk_Employee_Office FOREIGN KEY (Office_ID ) REFERENCES Office(ID),
  CONSTRAINT fk_Employee_Position FOREIGN KEY (Position_ID) REFERENCES Reference(ID),
  CONSTRAINT fk_Employee_superior FOREIGN KEY (Superior_ID) REFERENCES Employee(ID),
  CONSTRAINT fk_Employee_Type FOREIGN KEY (Type_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxEmployee_UUID ON public.Employee(UUID);
CREATE UNIQUE INDEX IF NOT EXISTS idxEmployee_No ON public.Employee(Central_Office_ID, lower(Employee_No));
CREATE INDEX IF NOT EXISTS idxEmployee_Status ON public.Employee(Status_Code);

DROP TRIGGER IF EXISTS trgEmployee_Ins on Employee;
---------------------------------------------------------------------------
CREATE TRIGGER trgEmployee_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Employee
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();

DROP TRIGGER IF EXISTS trgEmployee_upd on Employee;
---------------------------------------------------------------------------
CREATE TRIGGER trgEmployee_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Employee
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgEmployee_del on Employee;
---------------------------------------------------------------------------
CREATE TRIGGER trgEmployee_del
---------------------------------------------------------------------------
    AFTER DELETE ON Employee
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();

  
----------------------------------------------------------------------------------------
CREATE OR REPLACE VIEW public.vwEmployee
----------------------------------------------------------------------------------------
AS SELECT 
    emp.ID, mr.UUID,
    p.IIID,
 
    p.Title, p.Last_Name, p.First_Name, p.Middle_Name, p.Mother_Maiden_Name,
    p.Birthday, p.Sex,
    
    p.Gender_ID, p.Gender,    
    
    p.Civil_Status_ID , p.Civil_Status,
        
    p.Birth_Place_ID, p.Birth_Place, p.Full_Birth_Place, p.Birth_PlaceURL, 
    
    p.Contact_ID,
    p.Contact_Last_Name, p.Contact_First_Name,p.Contact_Middle_Name,
    
    p.Alternate_ID, p.Identity_Map_ID, p.Simple_Name, 

    p.iiMod_Ctr,
    p.iiOther_Info,
    p.iiCreated,
    p.iiUpdated,

    p.Current_Address_ID,
    p.Current_Detail,
    p.Current_URL,    
    p.Location, p.Full_Location,
 
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
   
    cOf.ID Central_Office_ID, cOf.Code CentralOfficeCode, 
    cOf.Short_Name CentralOfficeShort_Name, cOf.Office_Name CentralOffice_Name,
    
    
    ofc.ID Office_ID , ofc.Code OfficeCode, 
    ofc.Short_Name OfficeShort_Name, ofc.Office_Name Office_Name,
 
    emp.Employee_No,
 
    pos.ID Position_ID, pos.UUID PositionUUID, pos.Title PositionDesc,    
    
    emp.Superior_ID,
    
    sta.ID Status_ID , sta.UUID StatusUUID, sta.Title Status,    
    typ.ID Type_ID, typ.UUID Employee_TypeUUID, typ.Title Employee_Type,    

    lvl.ID Level_ID, lvl.UUID LevelUUID, lvl.Title LevelDesc,    

    emp.Date_Hired,
    emp.Date_Regular,
    emp.Job_Grade,
    emp.Job_Step,
    emp.Basic_Pay,
    
    mr.Mod_Ctr,
    emp.Other_Info,
    mr.Created,
    mr.Updated 
    
   FROM Employee emp
   INNER JOIN Main_Record mr on mr.UUID = emp.UUID
   LEFT JOIN vwperson p ON p.IIID = emp.IIID
   LEFT JOIN Reference pos ON pos.ID = emp.Position_ID   
   LEFT JOIN Reference sta ON sta.ID = emp.Status_Code and lower(sta.Ref_Type) = 'employeestatus'
   LEFT JOIN Reference typ ON typ.ID = emp.Type_ID 
   LEFT JOIN Reference lvl ON lvl.ID = emp.Level_ID   
   LEFT JOIN Office cOf ON cOf.ID = emp.Central_Office_ID   
   LEFT JOIN Office ofc ON ofc.ID = emp.Office_ID  
   --LEFT JOIN Employee sv on sv.ID = emp.Superior_ID
   ;

  INSERT INTO Employee(
      IIID, Central_Office_ID, Employee_No, Date_Hired, Basic_Pay, Date_Regular,
      Job_Grade, Job_Step, Level_ID, Office_ID , Position_ID, Status_Code , 
      Superior_ID, Type_ID)
  
  SELECT 
      ii.ID IIID, Cen.ID Central_Office_ID, a.Employee_No, 
      CAST(a.Date_Hired as Date) Date_Hired, a.Basic_Pay, CAST(a.Date_Regular as Date) Date_Regular,
      a.Job_Grade, a.Job_Step, Lvl.ID EmpLevelID, Ofc.ID Office_ID , 
      Title.ID Position_ID, Stat.Code Status_Code, 
      visor.ID Superior_ID, typ.ID EmpType_ID
      
  FROM (Values
      ('100','10001','97-0114',10000,'11/03/1997','04/01/1998',3,3,'Officer','10002','IT Officer','Regular',NULL,'Employed')
      )   
    a(Alternate_ID, CentralAlt, Employee_No, Basic_Pay, Date_Hired, Date_Regular,
      Job_Grade, Job_Step, EmpLevel, OfficeAlt, Title, Status, SuperiorNo, EmpType
      )  
  LEFT JOIN Identity_Info ii   on ii.Alternate_ID   = a.Alternate_ID
  LEFT JOIN Office Cen  on lower(Cen.Alternate_ID) = lower(a.CentralAlt)  
  LEFT JOIN Office Ofc  on lower(Ofc.Alternate_ID) = lower(a.OfficeAlt)  
  LEFT JOIN vwReference Title on lower(Title.Title) = lower(a.Title)     and lower(Title.Ref_Type)  = 'position'
  LEFT JOIN vwReference Lvl   on lower(Lvl.Title)   = lower(a.EmpLevel)  and lower(Lvl.Ref_Type)  = 'employeelevel'
  LEFT JOIN vwReference Stat  on lower(Stat.Title)  = lower(a.Status)    and lower(Stat.Ref_Type) = 'employeestatus'
  LEFT JOIN vwReference Typ   on lower(Typ.Title)   = lower(a.EmpType)   and lower(Typ.Ref_Type) = 'employeetype'
  LEFT JOIN Employee visor    on lower(visor.Employee_No) = lower(a.SuperiorNo) 
  ON CONFLICT(Central_Office_ID, lower(Employee_No)) DO UPDATE SET
    IIID = EXCLUDED.IIID, 
    Central_Office_ID = EXCLUDED.Central_Office_ID,
    Employee_No = EXCLUDED.Employee_No, 
    Date_Hired = EXCLUDED.Date_Hired, 
    Basic_Pay = EXCLUDED.Basic_Pay, 
    Date_Regular = EXCLUDED.Date_Regular,
    Job_Grade = EXCLUDED.Job_Grade, 
    Job_Step = EXCLUDED.Job_Step, 
    Level_ID = EXCLUDED.Level_ID, 
    Office_ID  = EXCLUDED.Office_ID , 
    Position_ID = EXCLUDED.Position_ID, 
    Status_Code  = EXCLUDED.Status_Code , 
    Superior_ID = EXCLUDED.Superior_ID, 
    Type_ID = EXCLUDED.Type_ID
   ;
 


