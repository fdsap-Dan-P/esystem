----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Users_Specs_Ref (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Users_ID bigint NOT NULL,
  Specs_Code varchar(50) NOT NULL,
  Specs_ID bigint NOT NULL,
  Ref_Id bigint NOT NULL,
  
  CONSTRAINT Users_Specs_Ref_pkey PRIMARY KEY (UUID),
  CONSTRAINT Users_Specs_Ref_ID FOREIGN KEY (Users_ID) REFERENCES Users(ID),
  CONSTRAINT Users_Specs_Ref_Item FOREIGN KEY (Specs_ID) REFERENCES Reference(ID),
  CONSTRAINT Users_Specs_Ref_Ref FOREIGN KEY (Ref_Id) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxUsers_Specs_Ref_Unique ON public.Users_Specs_Ref(Users_ID, Specs_ID);
CREATE UNIQUE INDEX IF NOT EXISTS idxUsers_Specs_Ref_Code ON public.Users_Specs_Ref(Users_ID, lower(Specs_Code));
  
DROP TRIGGER IF EXISTS trgUsers_Specs_Ref_Ins on Users_Specs_Ref;
---------------------------------------------------------------------------
CREATE TRIGGER trgUsers_Specs_Ref_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Users_Specs_Ref
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericSpecsInsert();
 
DROP TRIGGER IF EXISTS trgUsers_Specs_Ref_upd on Users_Specs_Ref;
---------------------------------------------------------------------------
CREATE TRIGGER trgUsers_Specs_Ref_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Users_Specs_Ref
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericSpecsUpdate();

DROP TRIGGER IF EXISTS trgUsers_Specs_Ref_del on Users_Specs_Ref;
---------------------------------------------------------------------------
CREATE TRIGGER trgUsers_Specs_Ref_del
---------------------------------------------------------------------------
    AFTER DELETE ON Users_Specs_Ref
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
 