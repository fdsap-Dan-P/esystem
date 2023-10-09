----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Charge_Type (
----------------------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Charge_Type varchar(100) NOT NULL,
  UnRealized_ID bigint NOT NULL,
  Realized_ID bigint NOT NULL,
  Other_Info jsonb NULL,
  CONSTRAINT Charge_Type_pkey PRIMARY KEY (ID),
  CONSTRAINT fk_Charge_Type_unrel FOREIGN KEY (Unrealized_ID) REFERENCES Chartof_Account(ID),
  CONSTRAINT fk_Charge_Type_rel FOREIGN KEY (Realized_ID) REFERENCES Chartof_Account(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxCharge_Type_UUID ON public.Charge_Type(UUID);
CREATE UNIQUE INDEX IF NOT EXISTS idxCharge_Type_Name ON public.Charge_Type(LOWER(Charge_Type));

DROP TRIGGER IF EXISTS trgCharge_Type_Ins on Charge_Type;
---------------------------------------------------------------------------
CREATE TRIGGER trgCharge_Type_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Charge_Type
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgCharge_Type_upd on Charge_Type;
---------------------------------------------------------------------------
CREATE TRIGGER trgCharge_Type_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Charge_Type
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgCharge_Type_del on Charge_Type;
---------------------------------------------------------------------------
CREATE TRIGGER trgCharge_Type_del
---------------------------------------------------------------------------
    AFTER DELETE ON Charge_Type
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();


----------------------------------------------------------------------------------------
CREATE OR REPLACE VIEW public.vwCharge_Type
----------------------------------------------------------------------------------------
AS SELECT 
    typ.ID, mr.UUID,
    typ.Charge_Type,
    COAR.ID Unrealized_ID, COAR.UUID UnRealizedUUID, COAR.Title UnRealized, 
    COAU.ID Realized_ID,   COAU.UUID RealizedUUID, COAU.Title Realized,   

    mr.Mod_Ctr,
    typ.Other_Info,
    mr.Created,
    mr.Updated 
 
    FROM Charge_Type typ
    LEFT JOIN Main_Record mr ON typ.UUID = mr.UUID
    LEFT JOIN Chartof_Account COAR ON COAR.ID = typ.Unrealized_ID
    LEFT JOIN Chartof_Account COAU ON COAU.ID = typ.Realized_ID;

----------------------------------------------------------------------------------------
    INSERT INTO Charge_Type(
       Charge_Type, Unrealized_ID, Realized_ID)    
    SELECT 
      a.Charge_Type, COAR.ID Unrealized_ID, COAU.ID Realized_ID
    FROM (Values
      ('Service Fee','Service Charges/Fees', 'Service Charges/Fees')
      )   
      a(Charge_Type, UnRealized, Realized)
    LEFT JOIN Chartof_Account COAR on COAR.Title = a.UnRealized
    LEFT JOIN Chartof_Account COAU on COAU.Title = a.Realized

    ON CONFLICT (LOWER(Charge_Type)) DO UPDATE SET 
      Charge_Type = excluded.Charge_Type,
      Unrealized_ID = excluded.Unrealized_ID,
      Realized_ID = excluded.Realized_ID
    ;
    



