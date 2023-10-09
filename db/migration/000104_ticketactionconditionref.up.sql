----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public. Ticket_Action_Condition_Ref (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Ticket_Type_Status_Id bigint NOT NULL,
  Item_Code VarChar(50) NOT NULL,
  Item_ID bigint NOT NULL,
  Condition_ID bigint NOT NULL,
  Ref_ID bigint NOT NULL,
  
  CONSTRAINT Ticket_Action_Condition_Ref_pkey PRIMARY KEY (UUID),
  CONSTRAINT Ticket_Action_Condition_Ref_Ticket FOREIGN KEY (Ticket_Type_Status_Id) REFERENCES Ticket_Type_Status(ID),
  CONSTRAINT Ticket_Action_Condition_Ref_Item FOREIGN KEY (Item_ID) REFERENCES Reference(ID),
  CONSTRAINT Ticket_Action_Condition_Ref_Condi FOREIGN KEY (Condition_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxTicket_Action_Condition_Ref_Unique ON public. Ticket_Action_Condition_Ref(Ticket_Type_Status_Id, Item_ID);
CREATE UNIQUE INDEX IF NOT EXISTS idxTicket_Action_Condition_Ref_Code ON public.Ticket_Action_Condition_Ref(Ticket_Type_Status_Id, lower(Item_Code));

DROP TRIGGER IF EXISTS trgTicket_Action_Condition_Ref_Ins on  Ticket_Action_Condition_Ref;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicket_Action_Condition_Ref_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON  Ticket_Action_Condition_Ref
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericItemInsert();
 
DROP TRIGGER IF EXISTS trgTicket_Action_Condition_Ref_upd on Ticket_Action_Condition_Ref;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicket_Action_Condition_Ref_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON  Ticket_Action_Condition_Ref
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericItemUpdate();

DROP TRIGGER IF EXISTS trgTicket_Action_Condition_Ref_del on  Ticket_Action_Condition_Ref;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicket_Action_Condition_Ref_del
---------------------------------------------------------------------------
    AFTER DELETE ON  Ticket_Action_Condition_Ref
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
 