---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Comment (
---------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Record_UUID uuid NOT NULL,
  User_ID BigInt NOT NULL,
  Comment Text NOT NULL,
  Other_Info jsonb NULL,
  
  CONSTRAINT Comment_pkey PRIMARY KEY (UUID),
  CONSTRAINT fkComment_User FOREIGN KEY (User_ID) REFERENCES Users(ID),
  CONSTRAINT fkComment_Rec FOREIGN KEY (Record_UUID) REFERENCES Main_Record(UUID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxComment_UUID ON public.Comment(UUID);


DROP TRIGGER IF EXISTS trgCommentIns on Comment;
---------------------------------------------------------------------------
CREATE TRIGGER trgCommentIns
---------------------------------------------------------------------------
    BEFORE INSERT ON Comment
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgCommentupd on Comment;
---------------------------------------------------------------------------
CREATE TRIGGER trgCommentupd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Comment
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgComment_del on Comment;
---------------------------------------------------------------------------
CREATE TRIGGER trgComment_del
---------------------------------------------------------------------------
    AFTER DELETE ON Comment
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();


---------------------------------------------------------------------------

INSERT INTO Comment(
    UUID, Record_UUID, User_ID, Comment)
  SELECT 
    a.UUID, Record_UUID, ul.ID, Comment
    
   FROM (Values
      ('6b46011d-ad54-4135-a1c5-cf10a0563743'::UUID, 'da970ce4-dc2f-44af-b1a8-49a987148922'::UUID, 'erick1421@gmail.com', 'Comments from Testing')
      )   
    a(UUID, Record_UUID, Login_Name, Comment)  
  LEFT JOIN Users ul     on lower(ul.Login_Name) = lower(a.Login_Name) 

  ON CONFLICT(UUID)
  DO UPDATE SET
    Record_UUID = excluded.Record_UUID,
    Comment = excluded.Comment
  ;  
