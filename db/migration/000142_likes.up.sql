---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Likes (
---------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  User_ID bigint NOT NULL,
  Mood integer NOT NULL,  
  Date_Liked timestamp NOT NULL DEFAULT CURRENT_TIMESTAMP,
  CONSTRAINT Likes_pkey PRIMARY KEY (UUID, User_ID),
  CONSTRAINT fkLikes_User FOREIGN KEY (User_ID) REFERENCES Users(ID),
  CONSTRAINT fkLikes_UUID FOREIGN KEY (UUID) REFERENCES Main_Record(UUID)
);

CREATE UNIQUE INDEX IF NOT EXISTS idxLikes_UUID ON public.Likes(UUID);

DROP TRIGGER IF EXISTS trgLikes_Ins on Likes;
---------------------------------------------------------------------------
CREATE TRIGGER trgLikes_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Likes
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
  
DROP TRIGGER IF EXISTS trgLikes_upd on Likes;
---------------------------------------------------------------------------
CREATE TRIGGER trgLikes_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Likes
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();
    
DROP TRIGGER IF EXISTS trgLikes_del on Likes;
---------------------------------------------------------------------------
CREATE TRIGGER trgLikes_del
---------------------------------------------------------------------------
    AFTER DELETE ON Likes
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();