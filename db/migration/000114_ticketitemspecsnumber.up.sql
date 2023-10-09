----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Ticket_Item_Specs_Number (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Ticket_Item_ID bigint NOT NULL,
  Specs_Code VarChar(50) NOT NULL,
  Specs_ID bigint NOT NULL,
  Value numeric(16,6) NOT NULL,
  Value2 numeric(16,6) NOT NULL,
  Measure_ID bigint NULL,

  CONSTRAINT Ticket_Item_Specs_Number_pkey PRIMARY KEY (UUID),
  CONSTRAINT Ticket_Item_Specs_Number_ID FOREIGN KEY (Ticket_Item_ID) REFERENCES Ticket_Item(ID),
  CONSTRAINT fk_Ticket_Item_Specs_Unitmeasure FOREIGN KEY (Measure_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxTicket_Item_Specs_Number_Unique ON public.Ticket_Item_Specs_Number(Ticket_Item_ID, Specs_ID);
CREATE UNIQUE INDEX IF NOT EXISTS idxTicket_Item_Specs_Number_Code ON public.Ticket_Item_Specs_Number(Ticket_Item_ID, lower(Specs_Code));

DROP TRIGGER IF EXISTS trgTicket_Item_Specs_Number_Ins on Ticket_Item_Specs_Number;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicket_Item_Specs_Number_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Ticket_Item_Specs_Number
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericSpecsInsert();
 
DROP TRIGGER IF EXISTS trgTicket_Item_Specs_Number_upd on Ticket_Item_Specs_Number;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicket_Item_Specs_Number_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Ticket_Item_Specs_Number
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericSpecsUpdate();

DROP TRIGGER IF EXISTS trgTicket_Item_Specs_Number_del on Ticket_Item_Specs_Number;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicket_Item_Specs_Number_del
---------------------------------------------------------------------------
    AFTER DELETE ON Ticket_Item_Specs_Number
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
 
