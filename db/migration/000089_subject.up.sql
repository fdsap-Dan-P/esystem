---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Subject (
---------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Subject_Title varchar(50) NOT NULL,
  Subject_Ref_ID bigint NOT NULL,
  Type_ID bigint NOT NULL,
  Remarks Varchar(200) NOT NULL DEFAULT '',
  Other_Info jsonb NULL,
 
  CONSTRAINT Subject_pkey PRIMARY KEY (ID),
  CONSTRAINT fk_Subject_Subject FOREIGN KEY (Subject_Ref_ID) REFERENCES Reference(ID),
  CONSTRAINT fk_Subject_Type FOREIGN KEY (Type_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxSubject_UUID ON public.Subject(UUID);

DROP TRIGGER IF EXISTS trgSubject_Ins on Subject;
---------------------------------------------------------------------------
CREATE TRIGGER trgSubject_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Subject
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgSubject_upd on Subject;
---------------------------------------------------------------------------
CREATE TRIGGER trgSubject_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Subject
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgSubject_del on Subject;
---------------------------------------------------------------------------
CREATE TRIGGER trgSubject_del
---------------------------------------------------------------------------
    AFTER DELETE ON Subject
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
  
  INSERT into Subject(
    UUID, Subject_Title, Subject_Ref_ID, Type_ID, Remarks)
  SELECT
    a.UUID::UUID, a.Subject, subj.ID Subject_Ref_ID, typ.ID Type_ID, a.Remarks
  FROM (Values
      ('945b7370-d414-4a6c-9533-3e2d1337cf5f', 'English', 'English', 'Regular', 'Elementary English')
      )   
  a(UUID, Title, Subject, SubjectType, Remarks)
    
  LEFT JOIN vwReference subj on lower(subj.Title) = lower(a.Subject) and lower(subj.Ref_Type) = lower('Subject')
  LEFT JOIN vwReference typ on lower(typ.Title) = lower(a.SubjectType) and lower(typ.Ref_Type) = lower('SubjectType')
  
  ON CONFLICT(UUID) DO UPDATE SET
    Subject_Title = excluded.Subject_Title, 
    Subject_Ref_ID = excluded.Subject_Ref_ID,
    Type_ID = excluded.Type_ID, 
    Remarks = excluded.Remarks
  ;
