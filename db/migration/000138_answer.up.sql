---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Answer (
---------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Subject_Event_ID bigint NOT NULL,
  Question_ID bigint NOT NULL,
  Answers jsonb NULL,
  Points numeric(8,5) NOT NULL DEFAULT 0,
  Remarks varchar(200) NULL,
  Other_Info jsonb NULL,
 
  CONSTRAINT Answer_pkey PRIMARY KEY (ID),
  CONSTRAINT fk_Answer_Event FOREIGN KEY (Subject_Event_ID) REFERENCES Subject_Event(ID),
  CONSTRAINT fk_Answer_Question FOREIGN KEY (Question_ID) REFERENCES Question(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxAnswer_UUID ON public.Answer(UUID);

DROP TRIGGER IF EXISTS trgAnswer_Ins on Answer;
---------------------------------------------------------------------------
CREATE TRIGGER trgAnswer_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Answer
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgAnswer_upd on Answer;
---------------------------------------------------------------------------
CREATE TRIGGER trgAnswer_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Answer
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgAnswer_del on Answer;
---------------------------------------------------------------------------
CREATE TRIGGER trgAnswer_del
---------------------------------------------------------------------------
    AFTER DELETE ON Answer
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
  
