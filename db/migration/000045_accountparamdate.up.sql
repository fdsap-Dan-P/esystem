----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Account_Param_Date (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Param_ID bigint NOT NULL,
  Item_Code varchar(50) NOT NULL,
  Item_ID bigint NOT NULL,
  Value Date NOT NULL,
  Value2 Date NOT NULL,
  
  CONSTRAINT Account_Param_Date_pkey PRIMARY KEY (UUID),
  CONSTRAINT Account_Param_ID FOREIGN KEY (Param_ID) REFERENCES Account_Param(ID),
  CONSTRAINT Account_Param_Date_Item FOREIGN KEY (Item_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccount_Param_Date_Unique ON public.Account_Param_Date(Param_ID, Item_ID);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccount_Param_Date_Code ON public.Account_Param_Date(Param_ID, lower(Item_Code));

DROP TRIGGER IF EXISTS trgAccount_Param_Date_Ins on Account_Param_Date;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Param_Date_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Account_Param_Date
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericItemInsert();
 
DROP TRIGGER IF EXISTS trgAccount_Param_Date_upd on Account_Param_Date;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Param_Date_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Account_Param_Date
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericItemUpdate();

DROP TRIGGER IF EXISTS trgAccount_Param_Date_del on Account_Param_Date;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Param_Date_del
---------------------------------------------------------------------------
    AFTER DELETE ON Account_Param_Date
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
 
  