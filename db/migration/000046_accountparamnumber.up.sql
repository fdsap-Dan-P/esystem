----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Account_Param_Number (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Param_ID bigint NOT NULL,
  Item_Code varchar(50) NOT NULL,
  Item_ID bigint NOT NULL,
  Value numeric(16,6) NOT NULL,
  Value2 numeric(16,6) NOT NULL,
  Measure_ID bigint NULL,

  CONSTRAINT Account_Param_Number_pkey PRIMARY KEY (UUID),
  CONSTRAINT Account_Param_Number_ID FOREIGN KEY (Param_ID) REFERENCES Account_Param(ID),
  CONSTRAINT Account_Param_Number_Item FOREIGN KEY (Item_ID) REFERENCES Reference(ID),
  CONSTRAINT fk_Account_Param_Unitmeasure FOREIGN KEY (Measure_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccount_Param_Number_Unique ON public.Account_Param_Number(Param_ID, Item_ID);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccount_Param_Number_Code ON public.Account_Param_Number(Param_ID, lower(Item_Code));

DROP TRIGGER IF EXISTS trgAccount_Param_Number_Ins on Account_Param_Number;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Param_Number_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Account_Param_Number
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericItemInsert();
 
DROP TRIGGER IF EXISTS trgAccount_Param_Number_upd on Account_Param_Number;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Param_Number_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Account_Param_Number
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericItemUpdate();

DROP TRIGGER IF EXISTS trgAccount_Param_Number_del on Account_Param_Number;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Param_Number_del
---------------------------------------------------------------------------
    AFTER DELETE ON Account_Param_Number
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
 
