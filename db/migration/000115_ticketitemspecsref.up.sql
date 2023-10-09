----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Ticket_Item_Specs_Ref (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Ticket_Item_ID bigint NOT NULL,
  Specs_Code VarChar(50) NOT NULL,
  Specs_ID bigint NOT NULL,
  Ref_ID bigint NOT NULL,
  
  CONSTRAINT Ticket_Item_Specs_Ref_pkey PRIMARY KEY (UUID),
  CONSTRAINT Ticket_Item_Specs_Ref_ID FOREIGN KEY (Ticket_Item_ID) REFERENCES Ticket_Item(ID),
  CONSTRAINT Ticket_Item_Specs_Ref_Item FOREIGN KEY (Specs_ID) REFERENCES Reference(ID),
  CONSTRAINT Ticket_Item_Specs_Ref_Ref FOREIGN KEY (Ref_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxTicket_Item_Specs_Ref_Unique ON public.Ticket_Item_Specs_Ref(Ticket_Item_ID, Specs_ID);
CREATE UNIQUE INDEX IF NOT EXISTS idxTicket_Item_Specs_Ref_Code ON public.Ticket_Item_Specs_Ref(Ticket_Item_ID, lower(Specs_Code));

DROP TRIGGER IF EXISTS trgTicket_Item_Specs_Ref_Ins on Ticket_Item_Specs_Ref;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicket_Item_Specs_Ref_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Ticket_Item_Specs_Ref
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericSpecsInsert();
 
DROP TRIGGER IF EXISTS trgTicket_Item_Specs_Ref_upd on Ticket_Item_Specs_Ref;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicket_Item_Specs_Ref_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Ticket_Item_Specs_Ref
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericSpecsUpdate();

DROP TRIGGER IF EXISTS trgTicket_Item_Specs_Ref_del on Ticket_Item_Specs_Ref;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicket_Item_Specs_Ref_del
---------------------------------------------------------------------------
    AFTER DELETE ON Ticket_Item_Specs_Ref
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
 