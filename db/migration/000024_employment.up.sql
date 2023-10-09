----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Employment (
----------------------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  IIID bigint NOT NULL,
  Series int2 NOT NULL,
  Company VarChar(200) NOT NULL,
  Title VarChar(200) NOT NULL,
  Address_Detail varchar(150) NULL,
  Address_URL varchar(200) NULL,
  Geography_ID bigint NULL,  
  Start_Date date NULL,
  End_Date date NULL,
  Period_Date VarChar(100),
  Remarks Text,
  Other_Info jsonb NULL,

  CONSTRAINT Employment_pkey PRIMARY KEY (ID),
  CONSTRAINT fkEmployment_Identity FOREIGN KEY (IIID) REFERENCES Identity_Info(ID),
  CONSTRAINT fkEmployment_Geo FOREIGN KEY (Geography_ID) REFERENCES Geography(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxEmployment_UUID ON public.Employment(UUID);
CREATE UNIQUE INDEX IF NOT EXISTS idxEmployment_Unique ON public.Employment(IIID, Series);

DROP TRIGGER IF EXISTS trgEmploymentIns on Employment;
---------------------------------------------------------------------------
CREATE TRIGGER trgEmploymentIns
---------------------------------------------------------------------------
    BEFORE INSERT ON Employment
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgEmploymentupd on Employment;
---------------------------------------------------------------------------
CREATE TRIGGER trgEmploymentupd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Employment
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgEmployment_del on Employment;
---------------------------------------------------------------------------
CREATE TRIGGER trgEmployment_del
---------------------------------------------------------------------------
    AFTER DELETE ON Employment
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();

/*
INSERT into Employment(
      IIID, Series, Company, Title,
      Address_Detail, Address_URL, Geography_ID, 
      Start_Date, End_Date, Period_Date, Remarks) 
  
  SELECT 
      i.ID IIID, Series, Company, Title, 
      Address_Detail, a.URL, g.ID GographyID, 
      cast(Start_Date as Date) Start_Date, Cast(End_Date as Date) End_Date, Period_Date, Remarks
  FROM (Values
      ('100',1,'VYP MSC Computer Technology Center','Instructor','4th Floor Las Swerte Building, Rizal Ave.',NULL,'Brgy. santa Maria, San Pablo City, Laguna','06/01/1985','03/31/1991','1997','Computer Instructor in SPCNHS HighSchool Student')
      )   
    a(
      Alternate_ID, Series, Company, Title, 
      Address_Detail, URL, Geography, 
      Start_Date, End_Date, Period_Date, Remarks
      )   
  LEFT JOIN Identity_Info i on i.Alternate_ID = a.Alternate_ID 
  LEFT JOIN LATERAL SearchLocation(a.Geography,1) g ON true

  ON CONFLICT(IIID, Series) DO UPDATE SET
    Company = excluded.Company, 
    Title = excluded.Title, 
    Address_Detail = excluded.Address_Detail, 
    Address_URL = excluded.Address_URL, 
    Geography_ID = excluded.Geography_ID, 
    Start_Date = excluded.Start_Date, 
    Remarks = excluded.Remarks, 
    Period_Date = excluded.Period_Date
    ;
*/