---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Customer_Event (
---------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Trn_Head_ID bigint NOT NULL,
  Customer_ID bigint NOT NULL,
  Type_ID bigint NOT NULL,
  Remarks varchar(1000) NOT NULL,
  Other_Info jsonb NULL,

  CONSTRAINT Customer_Event_pkey PRIMARY KEY (UUID),
  CONSTRAINT fkCustomer_EventCustomer FOREIGN KEY (Customer_ID) REFERENCES Customer(ID),
  CONSTRAINT fkCustomer_TranTrn_Head FOREIGN KEY (Trn_Head_ID) REFERENCES Trn_Head(ID),
  CONSTRAINT fkCustomer_Event_Type FOREIGN KEY (Type_ID) REFERENCES Reference(ID)
);

DROP TRIGGER IF EXISTS trgCustomer_EventIns on Customer_Event;
---------------------------------------------------------------------------
CREATE TRIGGER trgCustomer_EventIns
---------------------------------------------------------------------------
    BEFORE INSERT ON Customer_Event
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgCustomer_Eventupd on Customer_Event;
---------------------------------------------------------------------------
CREATE TRIGGER trgCustomer_Eventupd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Customer_Event
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgCustomer_Event_del on Customer_Event;
---------------------------------------------------------------------------
CREATE TRIGGER trgCustomer_Event_del
---------------------------------------------------------------------------
    AFTER DELETE ON Customer_Event
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();

/*
---------------------------------------------------------------------------
  INSERT into Customer_Event(
     UUID, Trn_Head_ID, Customer_ID, Type_ID, Remarks
    ) 
 
  SELECT
    Acc.UUID, h.ID Trn_Head_ID, cust.ID Customer_ID, 
    e.ID Type_ID, Acc.Remarks
  FROM (Values
      ('aa5fa555-83e8-4c58-9b54-b0f801342a12'::UUID, '2af90d74-3bee-48c5-8935-443edafb8f5a'::UUID, '10017', 'Recognize', 'new Member')
      )   
  Acc(
    UUID, Trn_HeadUUID, Customer_Alt_ID, Event_Type, Remarks
    )
    
  INNER JOIN Trn_Head h on h.UUID = Acc.Trn_HeadUUID
  LEFT JOIN Customer cust on cust.Customer_Alt_ID = Acc.Customer_Alt_ID
  LEFT JOIN vwReference e on e.Title = Acc.Event_Type and lower(e.Ref_Type) = 'customereventtype'
  
  ON CONFLICT(UUID) DO UPDATE SET
    Trn_Head_ID = excluded.Trn_Head_ID,
    Customer_ID = excluded.Customer_ID,
    Type_ID = excluded.Type_ID,
    Remarks = excluded.Remarks
  ;

*/