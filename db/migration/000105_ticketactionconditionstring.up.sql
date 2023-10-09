----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Ticket_Action_Condition_String (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Item_Code VarChar(50) NOT NULL,
  Ticket_Type_Status_Id bigint NOT NULL,
  Item_ID bigint NOT NULL,
  Condition_ID bigint NOT NULL,
  Value text NOT NULL,
  
  CONSTRAINT Ticket_Action_Condition_String_pkey PRIMARY KEY (UUID),
  CONSTRAINT Ticket_Action_Condition_String_Ticket FOREIGN KEY (Ticket_Type_Status_Id) REFERENCES Ticket_Type_Status(ID),
  CONSTRAINT Ticket_Action_Condition_String_Item FOREIGN KEY (Item_ID) REFERENCES Reference(ID),
  CONSTRAINT Ticket_Action_Condition_String_Condi FOREIGN KEY (Condition_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxTicket_Action_Condition_String_Unique ON public.Ticket_Action_Condition_String(Ticket_Type_Status_Id, Item_ID);
CREATE UNIQUE INDEX IF NOT EXISTS idxTicket_Action_Condition_String_Code ON public.Ticket_Action_Condition_String(Ticket_Type_Status_Id, lower(Item_Code));

DROP TRIGGER IF EXISTS trgTicket_Action_Condition_String_Ins on Ticket_Action_Condition_String;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicket_Action_Condition_String_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Ticket_Action_Condition_String
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericItemInsert();
 
DROP TRIGGER IF EXISTS trgTicket_Action_Condition_String_upd on Ticket_Action_Condition_String;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicket_Action_Condition_String_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Ticket_Action_Condition_String
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericItemUpdate();

DROP TRIGGER IF EXISTS trgTicket_Action_Condition_String_del on Ticket_Action_Condition_String;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicket_Action_Condition_String_del
---------------------------------------------------------------------------
    AFTER DELETE ON Ticket_Action_Condition_String
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
 