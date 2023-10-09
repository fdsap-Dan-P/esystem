----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Ticket_Account_Type_Group (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Ticket_Type_ID bigint NOT NULL,
  Account_Type_Group_ID bigint NOT NULL,
  Status_ID bigint NOT NULL,
  
  CONSTRAINT Ticket_Account_Type_Group_pkey PRIMARY KEY (UUID),
  CONSTRAINT Ticket_Account_Type_Group_Type FOREIGN KEY (Ticket_Type_ID) REFERENCES Reference(ID),
  CONSTRAINT Ticket_Account_Type_Group_AccGrp FOREIGN KEY (Account_Type_Group_ID) REFERENCES Account_Type_Group(ID),
  CONSTRAINT Ticket_Account_Type_Group_Status FOREIGN KEY (Status_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxTicket_Account_Type_Group_Unique ON public.Ticket_Account_Type_Group(Ticket_Type_ID, Account_Type_Group_ID);


DROP TRIGGER IF EXISTS trgTicket_Account_Type_Group_Ins on Ticket_Account_Type_Group;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicket_Account_Type_Group_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Ticket_Account_Type_Group
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgTicket_Account_Type_Group_upd on Ticket_Account_Type_Group;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicket_Account_Type_Group_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Ticket_Account_Type_Group
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgTicket_Account_Type_Group_del on Ticket_Account_Type_Group;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicket_Account_Type_Group_del
---------------------------------------------------------------------------
    AFTER DELETE ON Ticket_Account_Type_Group
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
 