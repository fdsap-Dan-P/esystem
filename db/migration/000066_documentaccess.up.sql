---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Document_Access (
---------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Document_ID bigint NOT NULL,
  Role_ID bigint NOT NULL,
  Access_Code char(1) NOT NULL,
  Other_Info jsonb NULL,
  
  CONSTRAINT document_Access_pkey PRIMARY KEY (UUID),
  CONSTRAINT fk_document_Access_Documents FOREIGN KEY (Document_ID) REFERENCES Documents(ID)
);

DROP TRIGGER IF EXISTS trgdocument_Access_Ins on Document_Access;
---------------------------------------------------------------------------
CREATE TRIGGER trgdocument_Access_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Document_Access
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgdocument_Access_upd on Document_Access;
---------------------------------------------------------------------------
CREATE TRIGGER trgdocument_Access_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Document_Access
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();
---------------------------------------------------------------------------
