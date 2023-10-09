---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Ticket (
---------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Central_Office_ID bigint NOT NULL,
  Ticket_Type_ID bigint NOT NULL,
  Ticket_Date  Date NOT NULL,
  Postedby_ID bigint NOT NULL,
  Status_ID bigint NOT NULL,
  Remarks varchar(100) NULL,
  Other_Info jsonb NULL,
  
  CONSTRAINT Ticket_pkey PRIMARY KEY (ID),
  CONSTRAINT fkTicket_Office FOREIGN KEY (Central_Office_ID) REFERENCES Office(ID),
  CONSTRAINT fkTicket_Type FOREIGN KEY (Ticket_Type_ID) REFERENCES Reference(ID),
  CONSTRAINT fkTicketUserName FOREIGN KEY (Postedby_ID) REFERENCES Users(ID),
  CONSTRAINT fkTicketStatus FOREIGN KEY (Status_ID ) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxTicket_UUID ON public.Ticket(UUID);

DROP TRIGGER IF EXISTS trgTicketIns on Ticket;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicketIns
---------------------------------------------------------------------------
    BEFORE INSERT ON Ticket
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgTicketupd on Ticket;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicketupd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Ticket
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgTicket_del on Ticket;
---------------------------------------------------------------------------
CREATE TRIGGER trgTicket_del
---------------------------------------------------------------------------
    AFTER DELETE ON Ticket
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();


  INSERT INTO Ticket(
    UUID, Central_Office_ID, Ticket_Date, Ticket_Type_ID, Postedby_ID, Status_ID , Remarks)
  SELECT 
    cast(a.UUID as UUID), o.ID, cast(a.Ticket_Date  as Date), typ.ID Ticket_Type_ID, 
    ul.ID Postedby_ID, sta.ID Status_ID , a.Remarks
    
   FROM (Values
      ('da970ce4-dc2f-44af-b1a8-49a987148922','01-01-2020','Disburse','erick1421@gmail.com', 'Completed', '')
      )   
    a(UUID, Ticket_Date, Ticket_Type, Postedby, Status, Remarks)  

  LEFT JOIN vwReference typ     on lower(typ.Title) = lower(a.Ticket_Type) and lower(typ.Ref_Type) = 'tickettype'
  LEFT JOIN vwReference sta     on lower(sta.Title) = lower(a.Status) and lower(sta.Ref_Type) = 'ticketstatus'
  LEFT JOIN Users ul        on lower(Login_Name) = lower(Postedby) 
  LEFT JOIN Office o  on lower(o.Alternate_ID) = lower('1')  
  ON CONFLICT(UUID)
  DO UPDATE SET
    Ticket_Date  = excluded.Ticket_Date,
    Ticket_Type_ID = excluded.Ticket_Type_ID,
    Postedby_ID = excluded.Postedby_ID,
    Status_ID  = excluded.Status_ID,
    Remarks = excluded.Remarks
  ;  

select * from vwReference typ     where  lower(typ.Ref_Type) = 'tickettype'
