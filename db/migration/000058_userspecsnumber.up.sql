----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Users_Specs_Number (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Users_ID bigint NOT NULL,
  Specs_Code varchar(50) NOT NULL,
  Specs_ID bigint NOT NULL,
  Value numeric(16,6) NOT NULL,
  Value2 numeric(16,6) NOT NULL,
  Measure_ID bigint NULL,

  CONSTRAINT Users_Specs_Number_pkey PRIMARY KEY (UUID),
  CONSTRAINT Users_Specs_Number_ID FOREIGN KEY (Users_ID) REFERENCES Users(ID),
  CONSTRAINT fk_Users_Specs_Unitmeasure FOREIGN KEY (Measure_ID) REFERENCES Reference(ID),
  CONSTRAINT Users_Specs_Number_Item FOREIGN KEY (Specs_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxUsers_Specs_Number_Unique ON public.Users_Specs_Number(Users_ID, Specs_ID);
CREATE UNIQUE INDEX IF NOT EXISTS idxUsers_Specs_Number_Code ON public.Users_Specs_Number(Users_ID, lower(Specs_Code));

DROP TRIGGER IF EXISTS trgUsers_Specs_Number_Ins on Users_Specs_Number;
---------------------------------------------------------------------------
CREATE TRIGGER trgUsers_Specs_Number_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Users_Specs_Number
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericSpecsInsert();
 
DROP TRIGGER IF EXISTS trgUsers_Specs_Number_upd on Users_Specs_Number;
---------------------------------------------------------------------------
CREATE TRIGGER trgUsers_Specs_Number_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Users_Specs_Number
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericSpecsUpdate();

DROP TRIGGER IF EXISTS trgUsers_Specs_Number_del on Users_Specs_Number;
---------------------------------------------------------------------------
CREATE TRIGGER trgUsers_Specs_Number_del
---------------------------------------------------------------------------
    AFTER DELETE ON Users_Specs_Number
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
 
