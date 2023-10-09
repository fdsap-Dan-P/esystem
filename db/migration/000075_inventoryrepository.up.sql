
----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Inventory_Repository (
----------------------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Central_Office_ID bigint NOT NULL,
  Repository_Code varchar(50) NOT NULL,
  Repository varchar(100) NOT NULL,
  Office_Id bigint NOT NULL,
  Custodian_Id bigint NULL,
  Geography_Id bigint NULL,
  Location_Description varchar(200) NULL,
  Remarks varchar(1000) NULL,
  Other_Info jsonb NULL,
  
  CONSTRAINT Inventory_Repository_pkey PRIMARY KEY (ID),
  CONSTRAINT Inventory_Repository_Cen_Off FOREIGN KEY (Central_Office_Id) REFERENCES Office(ID),
  CONSTRAINT Inventory_Repository_Off FOREIGN KEY (Office_Id) REFERENCES Office(ID),
  CONSTRAINT Inventory_Repository_Custodian FOREIGN KEY (Custodian_Id) REFERENCES Identity_Info(ID),
  CONSTRAINT Inventory_Repository_Geo FOREIGN KEY (Geography_Id) REFERENCES Geography(ID)
);

CREATE UNIQUE INDEX IF NOT EXISTS idxInventory_Repository_UUID ON public.Inventory_Repository(UUID);
CREATE UNIQUE INDEX IF NOT EXISTS idxInventory_Repository_Code ON public.Inventory_Repository(Central_Office_Id,lower(Repository_Code));

DROP TRIGGER IF EXISTS trgInventory_Repository_Ins on Inventory_Repository;
---------------------------------------------------------------------------
CREATE TRIGGER trgInventory_Repository_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Inventory_Repository
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgInventory_Repository_upd on Inventory_Repository;
---------------------------------------------------------------------------
CREATE TRIGGER trgInventory_Repository_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Inventory_Repository
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgInventory_Repository_del on Inventory_Repository;
---------------------------------------------------------------------------
CREATE TRIGGER trgInventory_Repository_del
---------------------------------------------------------------------------
    AFTER DELETE ON Inventory_Repository
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();

  
   