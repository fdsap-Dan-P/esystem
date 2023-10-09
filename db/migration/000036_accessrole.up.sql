---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Access_Role (
---------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Access_Name varchar(50) NOT NULL,
  Description varchar(500) NULL,
  Other_Info jsonb NULL,

  CONSTRAINT Access_Role_pkey PRIMARY KEY (ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccess_Role_UUID ON public.Access_Role(UUID);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccess_Role_Name ON public.Access_Role(lower(Access_Name));

DROP TRIGGER IF EXISTS trgAccess_Role_Ins on Access_Role;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccess_Role_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Access_Role
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgAccess_Role_upd on Access_Role;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccess_Role_upd
---------------------------------------------------------------------------
  BEFORE UPDATE ON Access_Role
  FOR EACH ROW
  WHEN (OLD.* IS DISTINCT FROM NEW.*)
  EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgAccess_Role_del on Access_Role;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccess_Role_del
---------------------------------------------------------------------------
    AFTER DELETE ON Access_Role
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();


  INSERT INTO Access_Role(Access_Name, Description)
  SELECT 
    Access_Name, null Description 
  FROM (Values
      ('Admin'),
      ('Bookkeeper'),
      ('Cashier'),
      ('Area Manager'),
      ('Unit Manager'),
      ('IT Officer'),
      ('Savings and Loans Teller'),
      ('SysAdmin')
      )   
    a(Access_Name)
  ON CONFLICT(lower(Access_Name))
  DO UPDATE SET Description = EXCLUDED.Description;
