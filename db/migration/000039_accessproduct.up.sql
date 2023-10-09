---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Access_Product (
---------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Role_ID bigint NOT NULL,
  Product_ID bigint NOT NULL,
  Allow boolean DEFAULT TRUE,
  Other_Info jsonb NULL,

  CONSTRAINT Access_Product_pkey PRIMARY KEY (Role_ID, Product_ID),
  CONSTRAINT fk_Access_Product_role   FOREIGN KEY (Role_ID)   REFERENCES Access_Role(ID));

CREATE UNIQUE INDEX IF NOT EXISTS idxAccess_Product_UUID ON public.Access_Product(UUID);

--------------------------------------------
-- CREATE Reference View
--------------------------------------------
CREATE OR REPLACE VIEW public.vwAccess_Product
AS SELECT 
    mr.UUID,
    r.ID as Role_ID, r.UUID as Access_RoleUUID,
    r.Access_Name,

    c.ID Product_ID, c.Code ProductCode, c.UUID ProductUUID, c.Product_Name, 
   
    rf.Allow,

    mr.Mod_Ctr,
    rf.Other_Info,
    mr.Created,
    mr.Updated 
   FROM Access_Product rf
   INNER JOIN Main_Record mr on mr.UUID = rf.UUID
   LEFT JOIN Product c ON rf.Product_ID = c.ID
   LEFT JOIN Access_Role r ON r.ID = rf.Role_ID;

DROP TRIGGER IF EXISTS trgAccess_Product_Ins on Access_Product;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccess_Product_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Access_Product
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgAccess_Product_upd on Access_Product;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccess_Product_upd
---------------------------------------------------------------------------
  BEFORE UPDATE ON Access_Product
  FOR EACH ROW
  WHEN (OLD.* IS DISTINCT FROM NEW.*)
  EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgAccess_Product_del on Access_Product;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccess_Product_del
---------------------------------------------------------------------------
    AFTER DELETE ON Access_Product
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();


  INSERT INTO Access_Product(Role_ID, Product_ID, Allow)
  SELECT 
    r.ID Role_ID, c.ID Product_ID, a.Allow
  FROM (Values
      ('Admin', 'Loan', True)
      )   
    a(Access_Name, Product_Name, Allow)
  INNER JOIN Access_Role r on r.Access_Name = a.Access_Name
  INNER JOIN Product    c on c.Product_Name = a.Product_Name

  ON CONFLICT(Role_ID, Product_ID)
  DO UPDATE SET 
    Allow = EXCLUDED.Allow
  ;
