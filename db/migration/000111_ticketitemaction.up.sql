---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Ticket_Item_Action (
---------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Ticket_Item_ID bigint NOT NULL,
  Trn_Head_ID bigint NOT NULL,
  User_ID bigint NOT NULL,
  Action_ID bigint NOT NULL,
  Action_Date date NOT NULL,
  Remarks text NULL DEFAULT '',
  Other_Info jsonb NULL,
 
  CONSTRAINT Ticket_Item_Action_pkey PRIMARY KEY (ID),
  CONSTRAINT fk_Ticket_Item_Action_Ticket FOREIGN KEY (Ticket_Item_ID) REFERENCES Ticket(ID),
  CONSTRAINT fk_Ticket_Item_Action_Trn FOREIGN KEY (Trn_Head_ID) REFERENCES Trn_Head(ID),
  CONSTRAINT fk_Ticket_Item_Action_User FOREIGN KEY (User_ID) REFERENCES Users(ID),
  CONSTRAINT fk_Ticket_Item_Action_Action FOREIGN KEY (Action_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxTicket_Item_Action_UUID ON public.Ticket_Item_Action(UUID);

DROP TRIGGER IF EXISTS trgTicket_Item_Action_Ins on Ticket_Item_Action;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicket_Item_Action_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Ticket_Item_Action
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgTicket_Item_Action_upd on Ticket_Item_Action;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicket_Item_Action_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Ticket_Item_Action
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgTicket_Item_Action_del on Ticket_Item_Action;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicket_Item_Action_del
---------------------------------------------------------------------------
    AFTER DELETE ON Ticket_Item_Action
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
