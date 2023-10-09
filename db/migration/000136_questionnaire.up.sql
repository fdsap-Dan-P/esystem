---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Questionaire (
---------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Code varchar(30) NOT NULL,
  Version smallint NOT NULL,
  Title varchar(50) NOT NULL,
  Type_ID bigint NOT NULL,
  Subject_ID bigint NULL,
  Date_Revised date NOT NULL,
  Office_ID bigint NULL,
  Author_ID bigint NULL,
  Status_ID bigint NOT NULL,
  Point_Equivalent jsonb NULL,
  Remarks varchar(200) NULL,
  Other_Info jsonb NULL,
 
  CONSTRAINT Questionaire_pkey PRIMARY KEY (ID),
  CONSTRAINT fk_Questionaire_Type FOREIGN KEY (Type_ID) REFERENCES Reference(ID),
  CONSTRAINT fk_Questionaire_Subject FOREIGN KEY (Subject_ID) REFERENCES Reference(ID),
  CONSTRAINT fk_Questionaire_Office FOREIGN KEY (Office_ID) REFERENCES Office(ID),
  CONSTRAINT fk_Questionaire_Author FOREIGN KEY (Author_ID) REFERENCES Identity_Info(ID),
  CONSTRAINT fk_Questionaire_Status FOREIGN KEY (Status_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxQuestionaire_UUID ON public.Questionaire(UUID);
CREATE UNIQUE INDEX IF NOT EXISTS idxQuestionaire_CODE ON public.Questionaire(Code,Version);

DROP TRIGGER IF EXISTS trgQuestionaire_Ins on Questionaire;
---------------------------------------------------------------------------
CREATE TRIGGER trgQuestionaire_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Questionaire
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgQuestionaire_upd on Questionaire;
---------------------------------------------------------------------------
CREATE TRIGGER trgQuestionaire_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Questionaire
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgQuestionaire_del on Questionaire;
---------------------------------------------------------------------------
CREATE TRIGGER trgQuestionaire_del
---------------------------------------------------------------------------
   AFTER DELETE ON Questionaire
   FOR EACH ROW 
   EXECUTE FUNCTION trgGenericDelete();
  
   INSERT into Questionaire(
     UUID, Code, Version, Title, Type_ID, Subject_ID, Date_Revised, Office_ID, 
     Author_ID, Status_ID , Remarks
    ) 
 
  SELECT
     Acc.UUID, Acc.Code, Version, Acc.Title, typ.ID Type_ID, Subj.ID Subject_ID, Date_Revised, o.ID Office_ID, 
     auth.ID Author_ID, stat.ID Status_ID, Remarks
  FROM (Values
      ('50fd99e8-f430-4ff3-983d-e4f61390def0'::UUID, 'Q1', 1, 'English Question', 'Quiz', 'English', '01/01/2020'::Date, '10019', '1054','Active', 'Remarks')
      )   
  Acc(
     UUID, Code, Version, Title, Question_Type, Subject, Date_Revised, Office_altid, 
     AltIIID, Status, Remarks
    )
  LEFT JOIN vwReference typ on lower(typ.Title) = lower(Acc.Question_Type) and lower(typ.Ref_Type) = lower('QuestionaireType')
  LEFT JOIN vwReference Subj on lower(Subj.Title) = lower(Acc.Subject) and lower(Subj.Ref_Type) = lower('Subject')
  LEFT JOIN Identity_Info auth on auth.alternate_id = Acc.AltIIID 
  LEFT JOIN vwReference stat on lower(stat.Title) = lower(Acc.Status) and lower(stat.Ref_Type) = lower('QuestionStatus')
  LEFT JOIN Office o on o.Alternate_ID = Acc.Office_altid
  
  ON CONFLICT(Code, Version) DO UPDATE SET
    Code  = excluded.Code,
    Version = excluded.Version,
    Title = excluded.Title,
    Type_ID  = excluded.Type_ID,
    Subject_ID = excluded.Subject_ID,
    Date_Revised = excluded.Date_Revised,
    Office_ID = excluded.Office_ID,
    Author_ID = excluded.Author_ID,
    Status_ID  = excluded.Status_ID,
    Remarks = excluded.Remarks
  ;
