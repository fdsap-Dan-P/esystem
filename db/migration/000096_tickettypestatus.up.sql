
---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Ticket_Type_Status (
---------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Product_Ticket_Type_Id bigint NOT NULL,
  Status_ID bigint NOT NULL,
  Ticket_Type_Action_Array bigint[] NOT NULL,
  Other_Info jsonb NULL,
 
  CONSTRAINT Ticket_Type_Status_pkey PRIMARY KEY (ID),
  CONSTRAINT fk_ProductTicket_Office FOREIGN KEY (Product_Ticket_Type_Id) REFERENCES Product_Ticket_Type(ID),
  CONSTRAINT fk_Ticket_Type_Status FOREIGN KEY (Status_ID) REFERENCES Reference(ID)
);

CREATE UNIQUE INDEX IF NOT EXISTS idxTicket_Type_Status_UUID ON public.Ticket_Type_Status(UUID);
CREATE UNIQUE INDEX IF NOT EXISTS idxTicket_Type_Status_Unq ON public.Ticket_Type_Status(Product_Ticket_Type_Id, Status_ID);

DROP TRIGGER IF EXISTS trgTicket_Type_Status_Ins on Ticket_Type_Status;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicket_Type_Status_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Ticket_Type_Status
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgTicket_Type_Status_upd on Ticket_Type_Status;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicket_Type_Status_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Ticket_Type_Status
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgTicket_Type_Status_del on Ticket_Type_Status;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicket_Type_Status_del
---------------------------------------------------------------------------
    AFTER DELETE ON Ticket_Type_Status
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
