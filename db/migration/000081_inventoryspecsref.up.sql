----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Inventory_Specs_Ref (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Inventory_Item_ID bigint NOT NULL,
  Specs_Code VarChar(50) NOT NULL,
  Specs_ID bigint NOT NULL,
  Ref_ID bigint NOT NULL,
  
  CONSTRAINT Inventory_Specs_Ref_pkey PRIMARY KEY (UUID),
  CONSTRAINT Inventory_Specs_Ref_ID FOREIGN KEY (Inventory_Item_ID) REFERENCES Inventory_Item(ID),
  CONSTRAINT Inventory_Specs_Ref_Item FOREIGN KEY (Specs_ID) REFERENCES Reference(ID),
  CONSTRAINT Inventory_Specs_Ref_Ref FOREIGN KEY (Ref_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxInventory_Specs_Ref_Unique ON public.Inventory_Specs_Ref(Inventory_Item_ID, Specs_ID);
CREATE UNIQUE INDEX IF NOT EXISTS idxInventory_Specs_Ref_Code ON public.Inventory_Specs_Ref(Inventory_Item_ID, lower(Specs_Code));

DROP TRIGGER IF EXISTS trgInventory_Specs_Ref_Ins on Inventory_Specs_Ref;
---------------------------------------------------------------------------
CREATE TRIGGER trgInventory_Specs_Ref_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Inventory_Specs_Ref
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericSpecsInsert();
 
DROP TRIGGER IF EXISTS trgInventory_Specs_Ref_upd on Inventory_Specs_Ref;
---------------------------------------------------------------------------
CREATE TRIGGER trgInventory_Specs_Ref_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Inventory_Specs_Ref
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericSpecsUpdate();

DROP TRIGGER IF EXISTS trgInventory_Specs_Ref_del on Inventory_Specs_Ref;
---------------------------------------------------------------------------
CREATE TRIGGER trgInventory_Specs_Ref_del
---------------------------------------------------------------------------
    AFTER DELETE ON Inventory_Specs_Ref
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
 