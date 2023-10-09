---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Account_Param (
---------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Account_Type_ID bigint NOT NULL,
  Date_Implemented Date NOT NULL,
  Other_Info jsonb NULL,
  
  CONSTRAINT Account_Param_pkey PRIMARY KEY (ID),
  CONSTRAINT idxAccount_ParamID UNIQUE (Account_Type_ID, Date_Implemented),
  CONSTRAINT idxAccount_Param_UUID UNIQUE (UUID),
  CONSTRAINT fk_Account_Param_Account_Type FOREIGN KEY (Account_Type_ID) REFERENCES Account_Type(ID)
);


DROP TRIGGER IF EXISTS trgAccount_Param_Ins on Account_Param;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Param_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Account_Param
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgAccount_Param_upd on Account_Param;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Param_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Account_Param
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgAccount_Param_del on Account_Param;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Param_del
---------------------------------------------------------------------------
    AFTER DELETE ON Account_Param
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();

/*
--------------------------------------------
-- CREATE Account_Param View
--------------------------------------------
CREATE OR REPLACE VIEW public.vwAccount_Param
AS SELECT 
    mr.UUID,
        
    Date_Implemented,
    rf.Value_Int, rf.Value_Decimal, rf.Value_Date, rf.Value_String,
    
    
    par.ID Item_ID, par.UUID Param_ItemUUID, par.Title param_Item,
    y.ID Type_ID, y.UUID Account_TypeUUID, y.Account_Type,
    
    mr.Mod_Ctr,
    rf.Other_Info,
    mr.Created,
    mr.Updated 
   FROM Account_Param rf
     INNER JOIN Main_Record mr on mr.UUID = rf.UUID
     LEFT JOIN Reference par ON rf.Item_ID = par.ID
     LEFT JOIN Account_Type y ON y.ID = rf.Type_ID;

--------------------------------------------
    INSERT into Account_Param(
      Type_ID, Item_ID, Date_Implemented, 
      Value_Int, Value_Decimal, Value_Date, Value_String)     
    SELECT y.ID Type_ID, par.ID Item_ID, cast(a.Date_Implemented as Date), 
      a.Value_Int, cast(a.Value_Decimal as decimal), cast(a.Value_Date as Date), a.Value_String
    FROM (Values
        ('Sikap 1', 'Interest', '2010-10-01',
         30, null, null, null))
        a(Account_Type, param_Item, Date_Implemented, 
          Value_Int, Value_Decimal, Value_Date, Value_String)
    LEFT JOIN vwReference par on par.Ref_Type = 'Parameter' and par.Title = a.param_Item
    INNER JOIN Account_Type y on y.Account_Type = a.Account_Type
    LEFT JOIN Account_Param p on par.ID = p.Item_ID and y.ID = p.Type_ID 
    
    ON CONFLICT(Type_ID, Item_ID, Date_Implemented) DO UPDATE SET
      Date_Implemented  = EXCLUDED.Date_Implemented, 
      Value_Int = EXCLUDED.Value_Int, 
      Value_Decimal = EXCLUDED.Value_Decimal, 
      Value_Date = EXCLUDED.Value_Date, 
      Value_String = EXCLUDED.Value_String 
    ;
*/