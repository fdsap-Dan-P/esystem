---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Document_User (
---------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Document_ID bigint NOT NULL,
  User_ID bigint NOT NULL,
  Access_Code char(1) NULL,
  Other_Info jsonb NULL,
  
  CONSTRAINT document_User_pkey PRIMARY KEY (UUID),
  CONSTRAINT fk_document_User_User FOREIGN KEY (User_ID) REFERENCES Users(ID)
);

DROP TRIGGER IF EXISTS trgdocument_User_Ins on Document_User;
---------------------------------------------------------------------------
CREATE TRIGGER trgdocument_User_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Document_User
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgdocument_User_upd on Document_User;
---------------------------------------------------------------------------
CREATE TRIGGER trgdocument_User_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Document_User
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();
---------------------------------------------------------------------------

DROP TRIGGER IF EXISTS trgDocument_User_del on Document_User;
---------------------------------------------------------------------------
CREATE TRIGGER trgDocument_User_del
---------------------------------------------------------------------------
    AFTER DELETE ON Document_User
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
