---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Trn_Head (
---------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Trn_Serial VarChar(40) NOT NULL,
  Ticket_ID bigint NOT NULL,
  Trn_Date Date NOT NULL,
  Type_ID bigint NOT NULL,
  Particular varchar(300) NULL,
  Office_ID bigint NOT NULL,
  User_ID bigint NOT NULL,
  Transacting_IIID bigint NULL,
  ORNo varchar(30) NULL,
  isFinal bool NULL,
  isManual bool NULL,
  Alternate_Trn varchar(40) NULL,
  Reference varchar(50) NOT NULL,
  Other_Info jsonb NULL,
  
  CONSTRAINT Trn_Head_pkey PRIMARY KEY (ID),
  CONSTRAINT Trn_HeadSerial UNIQUE (Trn_Serial),
  CONSTRAINT Trn_HeadaltTrn UNIQUE (Alternate_Trn),
  CONSTRAINT Trn_Header UNIQUE (ORNo),
  CONSTRAINT fkTrn_Head_IDentity FOREIGN KEY (Transacting_IIID) REFERENCES Identity_Info(ID),
  CONSTRAINT fkTrn_HeadOffice FOREIGN KEY (Office_ID ) REFERENCES Office(ID),
  CONSTRAINT fkTrn_HeadTrn_Type FOREIGN KEY (Type_ID) REFERENCES Reference(ID),
  CONSTRAINT fkTrn_HeadUserName FOREIGN KEY (User_ID) REFERENCES Users(ID),
  CONSTRAINT fkTrn_HeadTicket FOREIGN KEY (Ticket_ID) REFERENCES Ticket(ID)
); 
CREATE UNIQUE INDEX IF NOT EXISTS idxTrn_Head_UUID ON public.Trn_Head(UUID);

DROP TRIGGER IF EXISTS trgTrn_HeadIns on Trn_Head;
---------------------------------------------------------------------------
CREATE TRIGGER trgTrn_HeadIns
---------------------------------------------------------------------------
    BEFORE INSERT ON Trn_Head
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgTrn_Headupd on Trn_Head;
---------------------------------------------------------------------------
CREATE TRIGGER trgTrn_Headupd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Trn_Head
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgTrn_Head_del on Trn_Head;
---------------------------------------------------------------------------
CREATE TRIGGER trgTrn_Head_del
---------------------------------------------------------------------------
    AFTER DELETE ON Trn_Head
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();

------------------------------------------------------ 
CREATE OR REPLACE VIEW vwTrn_Head AS
------------------------------------------------------ 
  SELECT
    th.ID, mr.UUID, th.Trn_Serial, th.Ticket_ID, th.Trn_Date, th.Type_ID, th.Particular,
    th.Office_ID , th.User_ID, th.Transacting_IIID, th.ORNo, th.isFinal, 
    th.isManual, th.Alternate_Trn, th.Reference, 

    mr.Mod_Ctr,
    th.Other_Info,
    mr.Created,
    mr.Updated 
    
  FROM Trn_Head th
  INNER JOIN Main_Record mr on th.UUID = mr.UUID
  INNER JOIN Ticket on Ticket.ID = th.Ticket_ID 
 ;

  INSERT INTO Trn_Head(
    UUID, Trn_Serial, Ticket_ID, Trn_Date, Type_ID, Office_ID , User_ID, 
    Transacting_IIID, ORNo, isFinal, isManual, Alternate_Trn, Reference, Particular)
  SELECT 
    a.UUID, a.TrnSerial, T.ID Ticket_ID, cast(Trn_Date as Date), 
    typ.ID Type_ID, o.ID Office_ID , ul.ID User_ID, 
    ii.ID Transacting_IIID, a.ORNo, a.isFinal, a.isManual, a.Alternate_Trn, a.Reference, a.Particular
    
   FROM (Values
      ('2af90d74-3bee-48c5-8935-443edafb8f5a'::UUID, 'Test-01', 'da970ce4-dc2f-44af-b1a8-49a987148922'::UUID,'01-01-2020', 'Payment', '10019', 'erick1421@gmail.com', '100', 20010, true, true, null, 'ref', 'Particular'),
      ('26dfab18-f80b-46cf-9c54-be79d4fc5d23'::UUID, 'Test-02','da970ce4-dc2f-44af-b1a8-49a987148922'::UUID,'01-01-2020', 'Payment', '10019', 'erick1421@gmail.com', '100', 20011, true, true, null, 'ref', 'Particular'),
      ('3793422c-eb9f-49f0-9ec6-e5cf80caac25'::UUID, 'Test-03', 'da970ce4-dc2f-44af-b1a8-49a987148922'::UUID,'01-01-2020', 'Payment', '10019', 'erick1421@gmail.com', '100', 20012, true, true, null, 'ref', 'Particular')
      )   
    a(UUID, TrnSerial, TicketUUID, Trn_Date, Trn_Type, Officealtid, Login_Name,  
      Transactingaltid, ORNo, isFinal, isManual, Alternate_Trn, Reference, Particular)  

  LEFT JOIN vwReference typ  on lower(typ.Title) = lower(a.Trn_Type) and lower(typ.Ref_Type) = 'trnheadtype'
  LEFT JOIN Ticket T on T.UUID = a.TicketUUID
  LEFT JOIN Users ul     on lower(ul.Login_Name) = lower(a.Login_Name) 
  LEFT JOIN Identity_Info ii on ii.Alternate_ID = a.Transactingaltid
  LEFT JOIN Office o         on lower(o.Alternate_ID) = lower(a.OfficeAltID) 
  
  ON CONFLICT(Trn_Serial)
  DO UPDATE SET
    Trn_Serial = excluded.Trn_Serial,
    Trn_Date = excluded.Trn_Date,
    Type_ID = excluded.Type_ID,
    Particular = excluded.Particular,
    Office_ID  = excluded.Office_ID ,
    User_ID = excluded.User_ID,
    Transacting_IIID = excluded.Transacting_IIID,
    ORNo = excluded.ORNo,
    isFinal = excluded.isFinal,
    isManual = excluded.isManual,
    Alternate_Trn = excluded.Alternate_Trn,
    Reference = excluded.Reference
  ;   
