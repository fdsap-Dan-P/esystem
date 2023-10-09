---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.School_Section (
---------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Syllabus_ID bigint NOT NULL,
  School_ID bigint NOT NULL,
  Course_ID bigint NOT NULL,
  Start_Date Date NULL NULL,
  End_Date Date NULL NULL,
  Adviser_ID bigint NULL,
  Status_ID bigint NOT NULL,
  Section_Name VarChar(100) NULL,
  Remarks Varchar(200) NOT NULL DEFAULT '',
  Other_Info jsonb NULL,
 
  CONSTRAINT School_Section_pkey PRIMARY KEY (ID),
  CONSTRAINT fk_School_Section_Syllabus FOREIGN KEY (Syllabus_ID) REFERENCES Syllabus(ID),
  CONSTRAINT fk_School_Section_School FOREIGN KEY (School_ID) REFERENCES Office(ID),
  CONSTRAINT fk_School_Section_Course FOREIGN KEY (Course_ID) REFERENCES Reference(ID),
  CONSTRAINT fk_School_Section_Adviser FOREIGN KEY (Adviser_ID) REFERENCES Employee(ID),
  CONSTRAINT fk_School_Section_Status FOREIGN KEY (Status_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxSchool_Section_UUID ON public.School_Section(UUID);
CREATE INDEX IF NOT EXISTS idxSchool_Section_Section ON public.School_Section(Lower(Section_Name));

DROP TRIGGER IF EXISTS trgSchool_Section_Ins on School_Section;
---------------------------------------------------------------------------
CREATE TRIGGER trgSchool_Section_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON School_Section
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgSchool_Section_upd on School_Section;
---------------------------------------------------------------------------
CREATE TRIGGER trgSchool_Section_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON School_Section
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgSchool_Section_del on School_Section;
---------------------------------------------------------------------------
CREATE TRIGGER trgSchool_Section_del
---------------------------------------------------------------------------
    AFTER DELETE ON School_Section
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
  
  INSERT into School_Section(
    UUID, Syllabus_ID, School_ID, Course_ID, Start_Date, End_Date, Adviser_ID, 
    Status_ID, Section_Name, Remarks
    )
  SELECT
   a.UUID, sy.ID Syllabus_ID, cors.School_ID, cors.ID Course_ID, a.Start_Date::date, a.End_Date::date, adviser.ID Adviser_ID, 
    stat.ID Status_ID, Section_Name, a.Remarks
  FROM (Values
      ('3bdd00c6-dffb-4b7b-9b02-bf358205a690'::UUID, '034ae2f0-ff93-405d-a7cf-0983750a6ed9'::UUID, '06/05/2022', '04/05/2023', '97-0114', 'Current', 'Grade 1 Section A', 'Grade 1 Section A')
      )   
  a(UUID, Syllabus, Start_Date, End_Date, Employee_No, Status, Section_Name, Remarks)
    
  LEFT JOIN Syllabus sy  on sy.UUID = a.Syllabus
  LEFT JOIN Course cors on cors.ID = sy.Course_ID
  LEFT JOIN Employee adviser on lower(adviser.Employee_No) = lower(a.Employee_No) 
  LEFT JOIN vwReference stat on lower(stat.Title) = lower(a.Status) and lower(stat.Ref_Type) = lower('SchoolSectionStatus')
  
  ON CONFLICT(UUID) DO UPDATE SET
    Syllabus_ID = excluded.Syllabus_ID, 
    School_ID = excluded.School_ID, 
    Course_ID = excluded.Course_ID,
    Start_Date = excluded.Start_Date, 
    End_Date = excluded.End_Date, 
    Adviser_ID = excluded.Adviser_ID, 
    Status_ID = excluded.Status_ID, 
    Section_Name = excluded.Section_Name, 
    Remarks = excluded.Remarks
  ;
