---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Post (
---------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  User_ID bigint NOT NULL,
  Caption varChar(100) NOT NULL,
  Message_Body text NOT NULL,
  URL text NOT NULL,
  Image_URI text NOT NULL,
  Thumbnail_URI text NOT NULL,
  Keywords text[] NOT NULL,
  Mood integer NOT NULL,
  Mood_Emoji varchar(100) NOT NULL,
  Other_Info jsonb NULL,

  CONSTRAINT fkpublic_User FOREIGN KEY (User_ID) REFERENCES Users(ID), 
  CONSTRAINT Post_pkey PRIMARY KEY (UUID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxSocial_Media_Credential_UUID ON public.Social_Media_Credential(UUID);

 
-- CREATE UNIQUE INDEX IF NOT EXISTS idxPost_UUID ON public.Post(UUID);

DROP TRIGGER IF EXISTS trgPostIns on Post;
---------------------------------------------------------------------------
CREATE TRIGGER trgPostIns
---------------------------------------------------------------------------
    BEFORE INSERT ON Post
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgPostupd on Post;
---------------------------------------------------------------------------
CREATE TRIGGER trgPostupd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Post
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgPost_del on Post;
---------------------------------------------------------------------------
CREATE TRIGGER trgPost_del
---------------------------------------------------------------------------
    AFTER DELETE ON Post
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
---------------------------------------------------------------------------
