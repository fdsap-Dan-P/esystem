---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Syllabus_Subject (
---------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Syllabus_Id bigint NOT NULL,
  Subject_Id bigint NOT NULL,
  Units varchar(50) NOT NULL,
  Type_Id bigint NOT NULL,
  Status_Id bigint NOT NULL,
  Remarks Varchar(200) NOT NULL DEFAULT '',
  Other_Info jsonb NULL,
 
  CONSTRAINT Syllabus_Subject_pkey PRIMARY KEY (ID),
  CONSTRAINT fk_Syllabus_Subject_Syllabus FOREIGN KEY (Syllabus_Id) REFERENCES Syllabus(ID),
  CONSTRAINT fk_Syllabus_Subject_Subject FOREIGN KEY (Subject_ID) REFERENCES Subject(ID),
  CONSTRAINT fk_Syllabus_Subject_Type FOREIGN KEY (Type_Id) REFERENCES Reference(ID),
  CONSTRAINT fk_Syllabus_Subject_Status FOREIGN KEY (Status_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxSyllabus_Subject_UUID ON public.Syllabus_Subject(UUID);
CREATE UNIQUE INDEX IF NOT EXISTS idxSyllabus_Subject_Subj ON public.Syllabus_Subject(Syllabus_Id, Subject_Id);

DROP TRIGGER IF EXISTS trgSyllabus_Subject_Ins on Syllabus_Subject;
---------------------------------------------------------------------------
CREATE TRIGGER trgSyllabus_Subject_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Syllabus_Subject
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgSyllabus_Subject_upd on Syllabus_Subject;
---------------------------------------------------------------------------
CREATE TRIGGER trgSyllabus_Subject_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Syllabus_Subject
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgSyllabus_Subject_del on Syllabus_Subject;
---------------------------------------------------------------------------
CREATE TRIGGER trgSyllabus_Subject_del
---------------------------------------------------------------------------
    AFTER DELETE ON Syllabus_Subject
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
    
