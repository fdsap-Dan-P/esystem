----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Ticket_Item_Specs_String (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Ticket_Item_ID bigint NOT NULL,
  Specs_Code VarChar(50) NOT NULL,
  Specs_ID bigint NOT NULL,
  Value text NOT NULL,
  
  CONSTRAINT Ticket_Item_Specs_String_pkey PRIMARY KEY (UUID),
  CONSTRAINT Ticket_Item_Specs_String_ID FOREIGN KEY (Ticket_Item_ID) REFERENCES Ticket_Item(ID),
  CONSTRAINT Ticket_Item_Specs_String_Item FOREIGN KEY (Specs_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxTicket_Item_Specs_String_Unique ON public.Ticket_Item_Specs_String(Ticket_Item_ID, Specs_ID);
CREATE UNIQUE INDEX IF NOT EXISTS idxTicket_Item_Specs_String_Code ON public.Ticket_Item_Specs_String(Ticket_Item_ID, lower(Specs_Code));

DROP TRIGGER IF EXISTS trgTicket_Item_Specs_String_Ins on Ticket_Item_Specs_String;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicket_Item_Specs_String_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Ticket_Item_Specs_String
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericSpecsInsert();
 
DROP TRIGGER IF EXISTS trgTicket_Item_Specs_String_upd on Ticket_Item_Specs_String;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicket_Item_Specs_String_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Ticket_Item_Specs_String
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericSpecsUpdate();

DROP TRIGGER IF EXISTS trgTicket_Item_Specs_String_del on Ticket_Item_Specs_String;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicket_Item_Specs_String_del
---------------------------------------------------------------------------
    AFTER DELETE ON Ticket_Item_Specs_String
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
 