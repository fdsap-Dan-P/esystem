---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Follower (
---------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  User_ID bigint NOT NULL,
  Follower_ID bigint NOT NULL,
  Date_Followed timestamp DEFAULT CURRENT_TIMESTAMP NOT NULL,
  is_Follower Boolean NOT NULL,
  CONSTRAINT Follower_pkey PRIMARY KEY (Follower_ID, User_ID),
  CONSTRAINT fkFollower_User FOREIGN KEY (User_ID) REFERENCES Users(ID),
  CONSTRAINT fkFollower_Follower FOREIGN KEY (Follower_ID) REFERENCES Users(ID)
);

CREATE UNIQUE INDEX IF NOT EXISTS idxFollower_user ON public.Follower(User_ID);

DROP TRIGGER IF EXISTS trgFollower_Ins on Follower;
---------------------------------------------------------------------------
CREATE TRIGGER trgFollower_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Follower
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
  
DROP TRIGGER IF EXISTS trgFollower_upd on Follower;
---------------------------------------------------------------------------
CREATE TRIGGER trgFollower_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Follower
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();
    
DROP TRIGGER IF EXISTS trgFollower_del on Follower;
---------------------------------------------------------------------------
CREATE TRIGGER trgFollower_del
---------------------------------------------------------------------------
    AFTER DELETE ON Follower
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();