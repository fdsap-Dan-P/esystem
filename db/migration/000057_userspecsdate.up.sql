----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Users_Specs_Date (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Users_ID bigint NOT NULL,
  Specs_Code varchar(50) NOT NULL,
  Specs_ID bigint NOT NULL,
  Value Date NOT NULL,
  Value2 Date NOT NULL,
  
  CONSTRAINT Users_Specs_Date_pkey PRIMARY KEY (UUID),
  CONSTRAINT Users_Specs_Date_ID FOREIGN KEY (Users_ID) REFERENCES Users(ID),
  CONSTRAINT Users_Specs_Date_Item FOREIGN KEY (Specs_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxUsers_Specs_Date_Unique ON public.Users_Specs_Date(Users_ID, Specs_ID);
CREATE UNIQUE INDEX IF NOT EXISTS idxUsers_Specs_Date_Code ON public.Users_Specs_Date(Users_ID, lower(Specs_Code));

DROP TRIGGER IF EXISTS trgUsers_Specs_Date_Ins on Users_Specs_Date;
---------------------------------------------------------------------------
CREATE TRIGGER trgUsers_Specs_Date_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Users_Specs_Date
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericSpecsInsert();
 
DROP TRIGGER IF EXISTS trgUsers_Specs_Date_upd on Users_Specs_Date;
---------------------------------------------------------------------------
CREATE TRIGGER trgUsers_Specs_Date_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Users_Specs_Date
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericSpecsUpdate();

DROP TRIGGER IF EXISTS trgUsers_Specs_Date_del on Users_Specs_Date;
---------------------------------------------------------------------------
CREATE TRIGGER trgUsers_Specs_Date_del
---------------------------------------------------------------------------
    AFTER DELETE ON Users_Specs_Date
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
 
  