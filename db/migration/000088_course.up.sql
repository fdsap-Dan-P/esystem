---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Course (
---------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  School_ID bigint NOT NULL,
  Course_Title varchar(50) NOT NULL,
  Course_Ref_ID bigint NOT NULL,
  Status_ID bigint NOT NULL,
  Remarks Varchar(200) NOT NULL DEFAULT '',
  Other_Info jsonb NULL,
 
  CONSTRAINT Course_pkey PRIMARY KEY (ID),
  CONSTRAINT fk_Course_Office FOREIGN KEY (School_ID) REFERENCES Office(ID),
  CONSTRAINT fk_Course_Course FOREIGN KEY (Course_Ref_ID) REFERENCES Reference(ID),
  CONSTRAINT fk_Course_Status FOREIGN KEY (Status_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxCourse_UUID ON public.Course(UUID);

DROP TRIGGER IF EXISTS trgCourse_Ins on Course;
---------------------------------------------------------------------------
CREATE TRIGGER trgCourse_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Course
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgCourse_upd on Course;
---------------------------------------------------------------------------
CREATE TRIGGER trgCourse_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Course
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgCourse_del on Course;
---------------------------------------------------------------------------
CREATE TRIGGER trgCourse_del
---------------------------------------------------------------------------
    AFTER DELETE ON Course
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
  
  INSERT into Course(
    UUID, School_ID, Course_Title, Course_Ref_ID, Status_ID, Remarks
    )
  SELECT
   a.UUID::UUID, o.ID Office_ID, a.Title, cors.ID Course_ID, stat.ID Status_ID, Remarks
  FROM (Values
      ('8845d46a-c910-4467-a40b-37008c78a288', '10019', 'Elementary', 'Elementary', 'Current', 'Elementary')
      )   
  a(UUID, OfficeAlt, Title, Course, Status, Remarks)
  LEFT JOIN Office o  on lower(o.Alternate_ID) = lower(a.OfficeAlt) 
  LEFT JOIN vwReference cors on lower(cors.Title) = lower(a.Course) and lower(cors.Ref_Type) = lower('Courses')
  LEFT JOIN vwReference stat on lower(stat.Title) = lower(a.Status) and lower(stat.Ref_Type) = lower('CourseStatus')
  
  ON CONFLICT(UUID) DO UPDATE SET
    School_ID = excluded.School_ID, 
    Course_Title = excluded.Course_Title,
    Course_Ref_ID = excluded.Course_Ref_ID, 
    Status_ID = excluded.Status_ID, 
    Remarks = excluded.Remarks
  ;
