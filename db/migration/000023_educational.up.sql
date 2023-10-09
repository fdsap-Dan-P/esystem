----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Educational (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  IIID bigint NOT NULL,
  Series int2 NOT NULL,
  Level_ID bigint NULL,
  Course_Type_ID bigint NOT NULL,
  Course_ID bigint NULL,
  Course VarChar(200) NOT NULL,
  School VarChar(200) NOT NULL,
  Address_Detail varchar(150) NULL,
  Address_URL varchar(200) NULL,
  Geography_ID bigint NULL,  
  Start_Date date NULL,
  End_Date date NULL,
  Period_Date VarChar(100),
  Completed Bool NULL,
  Other_Info jsonb NULL,

  CONSTRAINT Educational_pkey PRIMARY KEY (IIID, Series),
  CONSTRAINT fkEducational_Identity FOREIGN KEY (IIID) REFERENCES Identity_Info(ID),
  CONSTRAINT fkEducational_Level FOREIGN KEY (Level_ID) REFERENCES Reference(ID),
  CONSTRAINT fkEducational_Geo FOREIGN KEY (Geography_ID) REFERENCES Geography(ID),
  CONSTRAINT fkEducational_Course_Type FOREIGN KEY (Course_Type_ID) REFERENCES Reference(ID),
  CONSTRAINT fkEducational_Course FOREIGN KEY (Course_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxEducational_UUID ON public.Educational(UUID);

DROP TRIGGER IF EXISTS trgEducationalIns on Educational;
---------------------------------------------------------------------------
CREATE TRIGGER trgEducationalIns
---------------------------------------------------------------------------
    BEFORE INSERT ON Educational
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgEducationalupd on Educational;
---------------------------------------------------------------------------
CREATE TRIGGER trgEducationalupd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Educational
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgEducational_del on Educational;
---------------------------------------------------------------------------
CREATE TRIGGER trgEducational_del
---------------------------------------------------------------------------
    AFTER DELETE ON Educational
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();

/*
INSERT into Educational(
      IIID, Series, Level_ID, Course_Type_ID, Course_ID, Course, School, 
      Address_Detail, Address_URL, Geography_ID, 
      Start_Date, End_Date, Period_Date, Completed) 
  
  SELECT 
      i.ID IIID, Series, lvl.ID Level_ID, typ.ID Course_Type_ID, crs.ID Course_ID, COALESCE(Course,Course_Type) Course, School, 
      Address_Detail, a.URL, g.ID Geography_ID, 
      cast(Start_Date as Date) Start_Date, Cast(End_Date as Date) End_Date, Period_Date, Completed::bool Completed
  FROM (Values
      ('100','Grade 6',1,'Pre-College','Brgy. Sta. Maria Elementary School','Brgy. Sta Maria, San Pablo City, Laguna','https://www.google.com/maps/@14.0274821,121.3114398,3a,75y,53h,83.94t/data=!3m6!1e1!3m4!1sVXYSO32Q6BrY7nnVsN3jFA!2e0!7i13312!8i6656','Brgy. santa Maria, San Pablo City, Laguna',NULL,'06/01/1985','03/31/1991','1985-1991',1),
      ('100','High School Graduate',2,'Pre-College','San Pablo City National High School','Marasigan St, Brgy. VI-D, San Pablo City, San Pablo City, 4000 Laguna','https://www.google.com/maps/@14.0761487,121.321002,3a,75y,345.19h,86.12t/data=!3m6!1e1!3m4!1sOzaCK361kliGr0MwJqsaOg!2e0!7i13312!8i6656','brgy VI-D, San Pablo City, San Pablo City',NULL,'06/01/1991','03/31/1995','1991-1995',1),
      ('100','College Graduate',4,'Professions and Applied Sciences','San Pablo Colleges','Brgy III-A, Hermanos Belen Street, San Pablo City, 4000 Laguna',NULL,'brgy III-A pablo','Accountancy','06/01/2001','03/31/2007','2001-2007',1)
      )   
    a(
      Alternate_ID, Level, Series, Course_Type, School, 
      Address_Detail, URL, Geography, 
      Course, Start_Date, End_Date, Period_Date, Completed
      )   
  LEFT JOIN Identity_Info i on i.Alternate_ID = a.Alternate_ID 
  LEFT JOIN vwReference lvl on lower(lvl.Ref_Type) = 'educationallevel' and lower(lvl.Title)  = lower(a.Level)
  LEFT JOIN vwReference typ on lower(typ.Ref_Type) = 'coursetype' and lower(typ.Title)  = lower(a.Course_Type)
  LEFT JOIN vwReference crs on lower(crs.Ref_Type) = 'courses' and lower(crs.Title)  = lower(a.Course)
  LEFT JOIN LATERAL SearchLocation(a.Geography,1) g ON true

  ON CONFLICT(IIID, Series) DO UPDATE SET
    Level_ID = excluded.Level_ID, 
    Course_Type_ID = excluded.Course_Type_ID, 
    Course_ID = excluded.Course_ID,
    School = excluded.School, 
    Address_Detail = excluded.Address_Detail, 
    Address_URL = excluded.Address_URL, 
    Geography_ID = excluded.Geography_ID, 
    Start_Date = excluded.Start_Date, 
    Completed = excluded.Completed, 
    Period_Date = excluded.Period_Date
    ;
*/