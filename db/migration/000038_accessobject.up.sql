---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Access_Object (
---------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Role_ID bigint NOT NULL,
  Object_ID bigint NOT NULL,
  Allow boolean DEFAULT TRUE,
  Max_Value numeric(16,6) NOT NULL DEFAULT 0,
  Other_Info jsonb NULL,

  CONSTRAINT Access_Object_pkey PRIMARY KEY (Role_ID, Object_ID),
  CONSTRAINT fk_Access_Object_role   FOREIGN KEY (Role_ID)   REFERENCES Access_Role(ID),
  CONSTRAINT fk_Access_Object_Object FOREIGN KEY (Object_ID) REFERENCES Reference(ID),
  CONSTRAINT idxAccess_Object UNIQUE (Role_ID, Object_ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccess_Object_UUID ON public.Access_Object(UUID);

DROP TRIGGER IF EXISTS trgAccess_Object_Ins on Access_Object;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccess_Object_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Access_Object
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgAccess_Object_upd on Access_Object;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccess_Object_upd
---------------------------------------------------------------------------
  BEFORE UPDATE ON Access_Object
  FOR EACH ROW
  WHEN (OLD.* IS DISTINCT FROM NEW.*)
  EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgAccess_Object_del on Access_Object;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccess_Object_del
---------------------------------------------------------------------------
    AFTER DELETE ON Access_Object
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();

--------------------------------------------
-- CREATE Reference View
--------------------------------------------
CREATE OR REPLACE VIEW public.vwAccess_Object
AS SELECT 
    mr.UUID,
    r.ID as Role_ID,
    r.Access_Name,

    c.ID as Object_ID, c.UUID as Access_ObjectUUID,
    c.Title as Access_Object,    
    
    rf.Allow, rf.Max_Value,

    mr.Mod_Ctr,
    rf.Other_Info,
    mr.Created,
    mr.Updated 
   FROM Access_Object rf
     INNER JOIN Main_Record mr on mr.UUID = rf.UUID
     LEFT JOIN Reference c ON rf.Object_ID = c.ID
     LEFT JOIN Access_Role r ON r.ID = rf.Role_ID;

-- DROP TRIGGER Access_Object_UPDATE ON Access_Object;
DROP TRIGGER IF EXISTS trgAccess_Object on Access_Object;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccess_Object
---------------------------------------------------------------------------
  BEFORE UPDATE ON Access_Object
  FOR EACH ROW
  WHEN (OLD.* IS DISTINCT FROM NEW.*)
  EXECUTE PROCEDURE trgGenericUpdate();

  INSERT INTO Access_Object(Role_ID, Object_ID, Allow, Max_Value)
  SELECT 
    r.ID Role_ID, c.ID Object_ID, a.Allow, a.Max_Value
  FROM (Values
      ('Admin', 'File', True, 0),
      ('Bookkeeper', 'File', True, 0),
      ('Cashier', 'File', True, 0),
      ('Pastor', 'File', True, 0),
      ('Member', 'File', True, 0)
      )   
    a(Access_Name, Access_Object, Allow, Max_Value), vwReference c, Access_Role r
   WHERE c.Ref_Type = 'Access_Object' and c.Title = a.Access_Object
     and r.Access_Name = a.Access_Name
   ORDER BY 1
  ON CONFLICT(Role_ID, Object_ID)
  DO UPDATE SET 
    Allow = EXCLUDED.Allow, 
    Max_Value = EXCLUDED.Max_Value;
