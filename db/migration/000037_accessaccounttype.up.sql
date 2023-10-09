---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Access_Account_Type (
---------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Role_ID bigint NOT NULL,
  Type_ID bigint NOT NULL,
  Allow boolean DEFAULT TRUE,
  Other_Info jsonb NULL,

  CONSTRAINT Access_Account_Type_pkey PRIMARY KEY (Role_ID, Type_ID),
  CONSTRAINT fk_Access_Account_Type_role   FOREIGN KEY (Role_ID)   REFERENCES Access_Role(ID),
  CONSTRAINT fk_Access_Account_Type_Acctype   FOREIGN KEY (Type_ID)   REFERENCES Account_Type(ID)
  );

CREATE UNIQUE INDEX IF NOT EXISTS idxAccess_Account_Type_UUID ON public.Access_Account_Type(UUID);

--------------------------------------------
-- CREATE Reference View
--------------------------------------------
CREATE OR REPLACE VIEW public.vwAccess_Account_Type
AS SELECT 
    mr.UUID,
    r.ID as Role_ID, r.UUID as Access_RoleUUID,
    r.Access_Name,

    c.ID Type_ID, c.Code AccountCode, c.UUID AccountUUID, c.Account_Type, c.Product_ID, c.Group_ID, 

    rf.Allow,

    mr.Mod_Ctr,
    rf.Other_Info,
    mr.Created,
    mr.Updated 
   FROM Access_Account_Type rf
   INNER JOIN Main_Record mr on mr.UUID = rf.UUID
   LEFT JOIN Account_Type c ON rf.Type_ID = c.ID
   LEFT JOIN Access_Role r ON r.ID = rf.Role_ID;

DROP TRIGGER IF EXISTS trgAccess_Account_Type_Ins on Access_Account_Type;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccess_Account_Type_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Access_Account_Type
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgAccess_Account_Type_upd on Access_Account_Type;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccess_Account_Type_upd
---------------------------------------------------------------------------
  BEFORE UPDATE ON Access_Account_Type
  FOR EACH ROW
  WHEN (OLD.* IS DISTINCT FROM NEW.*)
  EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgAccess_Account_Type_del on Access_Account_Type;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccess_Account_Type_del
---------------------------------------------------------------------------
    AFTER DELETE ON Access_Account_Type
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();


  INSERT INTO Access_Account_Type(Role_ID, Type_ID, Allow)
  SELECT 
    r.ID Role_ID, c.ID Type_ID, a.Allow
  FROM (Values
      ('Admin', 'Sikap 1', True)
      )   
    a(Access_Name, Account_Type, Allow)
  INNER JOIN Access_Role r on r.Access_Name = a.Access_Name
  INNER JOIN Account_Type c on c.Account_Type = a.Account_Type

  ON CONFLICT(Role_ID, Type_ID)
  DO UPDATE SET 
    Allow = EXCLUDED.Allow
  ;
