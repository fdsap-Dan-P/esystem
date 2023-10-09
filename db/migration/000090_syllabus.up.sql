---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Syllabus (
---------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Course_ID bigint NOT NULL,
  Version varchar(20) NOT NULL,
  Course_Year smallint NOT NULL,
  Semister_ID bigint NOT NULL,
  Status_ID bigint NOT NULL,
  Date_Implement Date NOT NULL,
  Remarks Varchar(200) NOT NULL DEFAULT '',
  Other_Info jsonb NULL,
 
  CONSTRAINT Syllabus_pkey PRIMARY KEY (ID),
  CONSTRAINT fk_Syllabus_Course FOREIGN KEY (Course_ID) REFERENCES Course(ID),
  CONSTRAINT fk_Syllabus_Semister FOREIGN KEY (Semister_ID) REFERENCES Reference(ID),
  CONSTRAINT fk_Syllabus_Status FOREIGN KEY (Status_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxSyllabus_UUID ON public.Syllabus(UUID);
CREATE INDEX IF NOT EXISTS idxSyllabus_Cors ON public.Syllabus(Course_ID);

DROP TRIGGER IF EXISTS trgSyllabus_Ins on Syllabus;
---------------------------------------------------------------------------
CREATE TRIGGER trgSyllabus_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Syllabus
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgSyllabus_upd on Syllabus;
---------------------------------------------------------------------------
CREATE TRIGGER trgSyllabus_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Syllabus
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgSyllabus_del on Syllabus;
---------------------------------------------------------------------------
CREATE TRIGGER trgSyllabus_del
---------------------------------------------------------------------------
    AFTER DELETE ON Syllabus
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
    
  INSERT into Syllabus(
    UUID, Course_ID, Version, Course_Year, Semister_ID, Status_ID, 
    Date_Implement, Remarks
    )
  SELECT
   a.UUID, cors.ID Course_ID, Version, Course_Year, sem.ID Semister_ID,
   stat.ID Status_ID, Date_Implement, a.Remarks
  FROM (Values
      ('034ae2f0-ff93-405d-a7cf-0983750a6ed9'::UUID, '8845d46a-c910-4467-a40b-37008c78a288'::UUID, '2001', 1, 'Full Year', 'Current', '06/05/2001'::Date, 'Elementary English')
      )   
    a(UUID, Course, Version, Course_Year, Semister, Status, Date_Implement, Remarks)
    
  LEFT JOIN Course cors on a.Course = cors.UUID
  LEFT JOIN vwReference sem on lower(sem.Title) = lower(a.Semister) and lower(sem.Ref_Type) = lower('SchoolSemister')
  LEFT JOIN vwReference stat on lower(stat.Title) = lower(a.Status) and lower(stat.Ref_Type) = lower('SubjectStatus')
  
  ON CONFLICT(UUID) DO UPDATE SET
    Course_ID = excluded.Course_ID,
    Version = excluded.Version, 
    Course_Year = excluded.Course_Year, 
    Semister_ID = excluded.Semister_ID, 
    Status_ID = excluded.Status_ID, 
    Date_Implement = excluded.Date_Implement, 
    Remarks = excluded.Remarks
  ;
