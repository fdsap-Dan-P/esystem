---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Ticket_Type_Action (
---------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Product_Ticket_Type_Id bigint NOT NULL,
  Action_ID bigint NOT NULL,
  ActionDesc Varchar(200) NOT NULL,
  Action_Link_ID bigint NULL,
  Other_Info jsonb NULL,
 
  CONSTRAINT Ticket_Type_Action_pkey PRIMARY KEY (ID),
  CONSTRAINT fk_Ticket_Type_Action_Ticket FOREIGN KEY (Product_Ticket_Type_Id) REFERENCES Product_Ticket_Type(ID),
  CONSTRAINT fk_Ticket_Type_Action_Action FOREIGN KEY (Action_ID) REFERENCES Reference(ID),
  CONSTRAINT fk_Ticket_Type_Action_Status FOREIGN KEY (Action_Link_ID) REFERENCES Action_Link(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxTicket_Type_Action_UUID ON public.Ticket_Type_Action(UUID);
CREATE UNIQUE INDEX IF NOT EXISTS idxTicket_Type_Action_Subject ON public.Ticket_Type_Action(Product_Ticket_Type_Id, Action_ID);

DROP TRIGGER IF EXISTS trgTicket_Type_Action_Ins on Ticket_Type_Action;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicket_Type_Action_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Ticket_Type_Action
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgTicket_Type_Action_upd on Ticket_Type_Action;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicket_Type_Action_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Ticket_Type_Action
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgTicket_Type_Action_del on Ticket_Type_Action;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicket_Type_Action_del
---------------------------------------------------------------------------
    AFTER DELETE ON Ticket_Type_Action
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
    