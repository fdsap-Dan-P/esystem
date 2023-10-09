---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Ticket_Type_Item (
---------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Ticket_Type_ID bigint NOT NULL,
  Item_ID bigint NOT NULL,
  Item Varchar(200) NOT NULL,
  Status_ID bigint NOT NULL,
  Remarks Varchar(200) NOT NULL DEFAULT '',
  Other_Info jsonb NULL,
 
  CONSTRAINT Ticket_Type_Item_pkey PRIMARY KEY (ID),
  CONSTRAINT fk_Ticket_Type_Item_Type FOREIGN KEY (Ticket_Type_ID) REFERENCES Reference(ID),
  CONSTRAINT fk_Ticket_Type_Item_Item FOREIGN KEY (Item_ID) REFERENCES Reference(ID),
  CONSTRAINT fk_Ticket_Type_Item_Status FOREIGN KEY (Status_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxTicket_Type_Item_UUID ON public.Ticket_Type_Item(UUID);
CREATE UNIQUE INDEX IF NOT EXISTS idxTicket_Type_Item_Unq ON public.Ticket_Type_Item(Ticket_Type_ID, Item_ID);

DROP TRIGGER IF EXISTS trgTicket_Type_Item_Ins on Ticket_Type_Item;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicket_Type_Item_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Ticket_Type_Item
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgTicket_Type_Item_upd on Ticket_Type_Item;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicket_Type_Item_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Ticket_Type_Item
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgTicket_Type_Item_del on Ticket_Type_Item;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicket_Type_Item_del
---------------------------------------------------------------------------
    AFTER DELETE ON Ticket_Type_Item
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
