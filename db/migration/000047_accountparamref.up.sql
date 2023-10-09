----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Account_Param_Ref (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Param_ID bigint NOT NULL,
  Item_Code varchar(50) NOT NULL,
  Item_ID bigint NOT NULL,
  Ref_ID bigint NOT NULL,
  
  CONSTRAINT Account_Param_Ref_pkey PRIMARY KEY (UUID),
  CONSTRAINT Account_Param_Ref_ID FOREIGN KEY (Param_ID) REFERENCES Account_Param(ID),
  CONSTRAINT Account_Param_Ref_Item FOREIGN KEY (Item_ID) REFERENCES Reference(ID),
  CONSTRAINT Account_Param_Ref_Ref FOREIGN KEY (Ref_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccount_Param_Ref_Unique ON public.Account_Param_Ref(Param_ID, Item_ID);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccount_Param_Ref_Code ON public.Account_Param_Ref(Param_ID, lower(Item_Code));

DROP TRIGGER IF EXISTS trgAccount_Param_Ref_Ins on Account_Param_Ref;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Param_Ref_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Account_Param_Ref
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericItemInsert();
 
DROP TRIGGER IF EXISTS trgAccount_Param_Ref_upd on Account_Param_Ref;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Param_Ref_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Account_Param_Ref
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericItemUpdate();

DROP TRIGGER IF EXISTS trgAccount_Param_Ref_del on Account_Param_Ref;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Param_Ref_del
---------------------------------------------------------------------------
    AFTER DELETE ON Account_Param_Ref
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
 