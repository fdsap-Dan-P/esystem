---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Ticket_Item_Assigned (
---------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Ticket_Item_ID bigint NOT NULL,
  User_ID bigint NOT NULL,
  Assigned_By_ID bigint NOT NULL,
  Assigned_Date Date NOT NULL,
  Remarks Varchar(200) NOT NULL DEFAULT '',
  Status_ID bigint NOT NULL,
  Other_Info jsonb NULL,
 
  CONSTRAINT Ticket_Item_Assigned_pkey PRIMARY KEY (UUID),
  CONSTRAINT fk_Ticket_Item_Assigned_Item FOREIGN KEY (Ticket_Item_ID) REFERENCES Ticket_Item(ID),
  CONSTRAINT fk_Ticket_Item_Assigned_Assignedby FOREIGN KEY (Assigned_By_ID) REFERENCES Users(ID),
  CONSTRAINT fk_Ticket_Item_Assigned_User FOREIGN KEY (User_ID) REFERENCES Users(ID),
  CONSTRAINT fk_Ticket_Item_Assigned_Status FOREIGN KEY (Status_ID) REFERENCES Reference(ID)
);
CREATE INDEX IF NOT EXISTS idxTicket_Item_Assigned_Unq ON public.Ticket_Item_Assigned(Ticket_Item_ID, User_ID);

DROP TRIGGER IF EXISTS trgTicket_Item_Assigned_Ins on Ticket_Item_Assigned;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicket_Item_Assigned_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Ticket_Item_Assigned
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgTicket_Item_Assigned_upd on Ticket_Item_Assigned;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicket_Item_Assigned_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Ticket_Item_Assigned
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgTicket_Item_Assigned_del on Ticket_Item_Assigned;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicket_Item_Assigned_del
---------------------------------------------------------------------------
    AFTER DELETE ON Ticket_Item_Assigned
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
