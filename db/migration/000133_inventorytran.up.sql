---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Inventory_Tran (
---------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Trn_Head_ID bigint NOT NULL,
  Series bigint NOT NULL,
  Inventory_Detail_ID bigint NOT NULL,
  Repository_ID bigint NOT NULL,
  Quantity numeric(16,6) NOT NULL,
  Unit_Price  numeric(16,6) NOT NULL,
  Discount numeric(10,6) NOT NULL,
  Tax_Amt numeric(16,6) NOT NULL,
  Net_Trn_Amt numeric(16,6) NOT NULL, 
  Other_Info jsonb NULL,
  
  CONSTRAINT Inventory_Tran_pkey PRIMARY KEY (Trn_Head_ID, Series),
  CONSTRAINT fkInventoryTranRepo FOREIGN KEY (Repository_ID) REFERENCES Inventory_Repository(ID),
  CONSTRAINT fkInventoryTranAccqny FOREIGN KEY (Inventory_Detail_ID) REFERENCES Inventory_Detail(ID),
  CONSTRAINT fkInventoryTranTrn_Head FOREIGN KEY (Trn_Head_ID) REFERENCES Trn_Head(ID))
 ;
CREATE UNIQUE INDEX IF NOT EXISTS idxInventory_Tran_UUID ON public.Inventory_Tran(UUID);

DROP TRIGGER IF EXISTS trgInventory_TranIns on Inventory_Tran;
---------------------------------------------------------------------------
CREATE TRIGGER trgInventory_TranIns
---------------------------------------------------------------------------
    BEFORE INSERT ON Inventory_Tran
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgInventory_Tranupd on Inventory_Tran;
---------------------------------------------------------------------------
CREATE TRIGGER trgInventory_Tranupd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Inventory_Tran
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgInventory_Tran_del on Inventory_Tran;
---------------------------------------------------------------------------
CREATE TRIGGER trgInventory_Tran_del
---------------------------------------------------------------------------
    AFTER DELETE ON Inventory_Tran
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();



  INSERT into Inventory_Tran(
      UUID, Trn_Head_ID, Series, Inventory_Detail_ID, Quantity, Unit_Price , Discount, Tax_Amt, Net_Trn_Amt)
  
  SELECT 
      Acc.UUID, h.ID Trn_Head_ID, Acc.Series, q.ID Inventory_ID, 
      Acc.Quantity, Acc.Unit_Price , Acc.Discount, Acc.Tax_Amt, Acc.Net_Trn_Amt
  FROM (Values      
      ('cd5a1c3e-08b5-4979-8948-678038536dc9'::UUID, 'a7fa02d3-5c91-4683-8429-a88d96030c8a'::UUID, 1, 'c3476afe-bd50-49e6-8de3-074555a8e1bd'::UUID,10, 100, 0, .12, 100 )
      )   
    Acc(
      UUID, Trn_HeadUUID, Series, Account_InventoryUUID, Quantity, Unit_Price , Discount, Tax_Amt, Net_Trn_Amt
      )
  INNER JOIN Trn_Head h on h.UUID = Acc.Trn_HeadUUID
  INNER JOIN Account_Inventory q on q.UUID = Acc.Account_InventoryUUID

  ON CONFLICT(UUID) DO UPDATE SET
    Trn_Head_ID = excluded.Trn_Head_ID,
    Series = excluded.Series,
    Inventory_Detail_ID = excluded.Inventory_Detail_ID,
    Quantity = excluded.Quantity,
    Unit_Price  = excluded.Unit_Price ,
    Discount = excluded.Discount,
    Tax_Amt = excluded.Tax_Amt,
    Net_Trn_Amt = excluded.Net_Trn_Amt
    ;
