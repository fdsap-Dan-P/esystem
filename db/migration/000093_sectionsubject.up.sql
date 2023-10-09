---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Section_Subject (
---------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  School_Section_ID bigint NOT NULL,
  Subject_ID bigint NOT NULL,
  Teacher_ID bigint NULL,
  Type_ID bigint NOT NULL,
  Status_ID bigint NOT NULL,
  Schedule_Code VarChar(100) NULL,
  Schedule_Json jsonb NULL,
  Remarks Varchar(200) NOT NULL DEFAULT '',
  Other_Info jsonb NULL,
 
  CONSTRAINT Section_Subject_pkey PRIMARY KEY (ID),
  CONSTRAINT fk_Section_Subject_Section FOREIGN KEY (School_Section_ID) REFERENCES School_Section(ID),
  CONSTRAINT fk_Section_Subject_Subject FOREIGN KEY (Subject_ID) REFERENCES Subject(ID),
  CONSTRAINT fk_Section_Subject_Teacher FOREIGN KEY (Teacher_ID) REFERENCES Employee(ID),
  CONSTRAINT fk_Section_Subject_Type FOREIGN KEY (Type_ID) REFERENCES Reference(ID),
  CONSTRAINT fk_Section_Subject_Status FOREIGN KEY (Status_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxSection_Subject_UUID ON public.Section_Subject(UUID);
CREATE UNIQUE INDEX IF NOT EXISTS idxSection_Subject_Subject ON public.Section_Subject(School_Section_ID, Subject_ID);

DROP TRIGGER IF EXISTS trgSection_Subject_Ins on Section_Subject;
---------------------------------------------------------------------------
CREATE TRIGGER trgSection_Subject_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Section_Subject
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgSection_Subject_upd on Section_Subject;
---------------------------------------------------------------------------
CREATE TRIGGER trgSection_Subject_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Section_Subject
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgSection_Subject_del on Section_Subject;
---------------------------------------------------------------------------
CREATE TRIGGER trgSection_Subject_del
---------------------------------------------------------------------------
    AFTER DELETE ON Section_Subject
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
  