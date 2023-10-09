
---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Product_Ticket_Type (
---------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Central_Office_ID bigint NOT NULL,
  Product_ID bigint NOT NULL,
  Ticket_Type_ID bigint NOT NULL,
  Status_ID bigint NOT NULL,
  Other_Info jsonb NULL,
 
  CONSTRAINT Product_Ticket_Type_pkey PRIMARY KEY (ID),
  CONSTRAINT fk_ProductTicket_Office FOREIGN KEY (Central_Office_ID) REFERENCES Office(ID),
  CONSTRAINT fk_Product_Ticket_Type FOREIGN KEY (Ticket_Type_ID) REFERENCES Reference(ID),
  CONSTRAINT fk_ProductTicket_Product FOREIGN KEY (Product_ID) REFERENCES Product(ID),
  CONSTRAINT fk_ProductTicket_Status FOREIGN KEY (Status_ID) REFERENCES Reference(ID)
);

CREATE UNIQUE INDEX IF NOT EXISTS idxProduct_Ticket_Type_UUID ON public.Product_Ticket_Type(UUID);
CREATE UNIQUE INDEX IF NOT EXISTS idxProduct_Ticket_Type_Unq ON public.Product_Ticket_Type(Central_Office_ID, Product_ID, Ticket_Type_ID);

DROP TRIGGER IF EXISTS trgProduct_Ticket_Type_Ins on Product_Ticket_Type;
---------------------------------------------------------------------------
CREATE TRIGGER trgProduct_Ticket_Type_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Product_Ticket_Type
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgProduct_Ticket_Type_upd on Product_Ticket_Type;
---------------------------------------------------------------------------
CREATE TRIGGER trgProduct_Ticket_Type_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Product_Ticket_Type
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgProduct_Ticket_Type_del on Product_Ticket_Type;
---------------------------------------------------------------------------
CREATE TRIGGER trgProduct_Ticket_Type_del
---------------------------------------------------------------------------
    AFTER DELETE ON Product_Ticket_Type
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
