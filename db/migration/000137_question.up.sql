---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Question (
---------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Questionaire_ID bigint NOT NULL,
  Series smallint NOT NULL,
  Type_ID bigint NOT NULL,
  Question_Item Text NOT NULL,
  Choices jsonb NULL,
  Answer_Type VarChar(2) NULL, --(S:String, N:Number, D:Date, B:Boolean)
  Parent_ID bigint NULL,
  Status_ID bigint NOT NULL,
  Remarks varchar(200) NULL,
  Other_Info jsonb NULL,
 
  CONSTRAINT Question_pkey PRIMARY KEY (ID),
  CONSTRAINT fk_Question_Quest FOREIGN KEY (Questionaire_ID) REFERENCES Questionaire(ID),
  CONSTRAINT fk_Question_Type FOREIGN KEY (Type_ID) REFERENCES Reference(ID),
  CONSTRAINT fk_Question_Parent FOREIGN KEY (Parent_ID) REFERENCES Question(ID),
  CONSTRAINT fk_Question_Status FOREIGN KEY (Status_ID ) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxQuestion_UUID ON public.Question(UUID);
CREATE UNIQUE INDEX IF NOT EXISTS idxQuestion_Unq on Question(Questionaire_ID, Series);

DROP TRIGGER IF EXISTS trgQuestion_Ins on Question;
---------------------------------------------------------------------------
CREATE TRIGGER trgQuestion_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Question
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgQuestion_upd on Question;
---------------------------------------------------------------------------
CREATE TRIGGER trgQuestion_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Question
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgQuestion_del on Question;
---------------------------------------------------------------------------
CREATE TRIGGER trgQuestion_del
---------------------------------------------------------------------------
    AFTER DELETE ON Question
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();

   INSERT into Question(
    UUID, Questionaire_ID, Series, Type_ID, Question_Item, 
    Choices, Answer_Type, Parent_ID, Status_ID, Remarks
    ) 
 
  SELECT
    Acc.UUID, q.ID Questionaire_ID, Acc.Series, typ.ID Type_ID, Acc.Question_Item, 
    Acc.Choices, Acc.Answer_Type, par.ID Parent_ID, stat.ID Status_ID, Acc.Remarks
  FROM (Values
      ('a4bc09b4-db99-4afb-be43-22b4938b1582'::UUID, '50fd99e8-f430-4ff3-983d-e4f61390def0'::UUID, 1, 'Multiple Choice', 'Capital of the Philippines'::text, '[{"Option":"Manila","Score":1},{"Option":"Cebu", "Score":0}]'::jsonb, 'S', NULL::UUID, 'Active', 'Question 1')
      )   
  Acc(
    UUID, QuestUUID, Series, QuestionType, Question_Item, 
    Choices, Answer_Type, ParentUUID, Status, Remarks
    )
    
  LEFT JOIN Questionaire q on q.UUID = Acc.QuestUUID
  LEFT JOIN vwReference typ on typ.Title = Acc.QuestionType and typ.Ref_Type = 'QuestionType'
  LEFT JOIN Question par on par.UUID = Acc.ParentUUID
  LEFT JOIN vwReference stat on stat.Title = Acc.Status and stat.Ref_Type = 'QuestionStatus'
  
  ON CONFLICT(UUID) DO UPDATE SET
    Questionaire_ID  = excluded.Questionaire_ID ,
    Series = excluded.Series,
    Type_ID = excluded.Type_ID,
    Question_Item = excluded.Question_Item,
    Choices = excluded.Choices,
    Answer_Type = excluded.Answer_Type,
    Parent_ID = excluded.Parent_ID,
    Status_ID  = excluded.Status_ID,
    Remarks = excluded.Remarks
  ;

   INSERT into Question(
    UUID, Questionaire_ID, Series, Type_ID, Question_Item, 
    Choices, Answer_Type, Parent_ID, Status_ID, Remarks
    ) 
 
  SELECT
    Acc.UUID, q.ID Questionaire_ID, Acc.Series, typ.ID Type_ID, Acc.Question_Item, 
    Acc.Choices, Acc.Answer_Type, par.ID Parent_ID, stat.ID Status_ID, Acc.Remarks
  FROM (Values
      ('89d84d39-e56f-47a1-8426-f3b5812116dd'::UUID, '50fd99e8-f430-4ff3-983d-e4f61390def0'::UUID, 2, 'Multiple Choice', 'Capital of the Laguna'::text, '[{"Option":"San Pablo City","Score":0},{"Option":"Sta. Cruz", "Score":1}]'::jsonb, 'S', 'a4bc09b4-db99-4afb-be43-22b4938b1582'::UUID, 'Active', 'Question 2')
      )   
  Acc(
    UUID, QuestUUID, Series, QuestionType, Question_Item, 
    Choices, Answer_Type, ParentUUID, Status, Remarks
    )
    
  LEFT JOIN Questionaire q on q.UUID = Acc.QuestUUID
  LEFT JOIN vwReference typ on typ.Title = Acc.QuestionType and typ.Ref_Type = 'QuestionType'
  LEFT JOIN Question par on par.UUID = Acc.ParentUUID
  LEFT JOIN vwReference stat on stat.Title = Acc.Status and stat.Ref_Type = 'QuestionStatus'
  
  ON CONFLICT(UUID) DO UPDATE SET
    Questionaire_ID  = excluded.Questionaire_ID ,
    Series = excluded.Series,
    Type_ID = excluded.Type_ID,
    Question_Item = excluded.Question_Item,
    Choices = excluded.Choices,
    Answer_Type = excluded.Answer_Type,
    Parent_ID = excluded.Parent_ID,
    Status_ID  = excluded.Status_ID,
    Remarks = excluded.Remarks
  ;

