---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Account_Event (
---------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Trn_Head_ID bigint NOT NULL,
  Account_ID bigint NOT NULL,
  Type_ID bigint NOT NULL,
  Remarks varchar(1000) NOT NULL,
  Other_Info jsonb NULL,

  CONSTRAINT Account_Event_pkey PRIMARY KEY (UUID),
  CONSTRAINT fkAccount_EventAccount FOREIGN KEY (Account_ID) REFERENCES Account(ID),
  CONSTRAINT fkAccount_TranTrn_Head FOREIGN KEY (Trn_Head_ID) REFERENCES Trn_Head(ID),
  CONSTRAINT fkAccount_Event_Type FOREIGN KEY (Type_ID) REFERENCES Reference(ID)
);

DROP TRIGGER IF EXISTS trgAccount_EventIns on Account_Event;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_EventIns
---------------------------------------------------------------------------
    BEFORE INSERT ON Account_Event
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgAccount_Eventupd on Account_Event;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Eventupd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Account_Event
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgAccount_Event_del on Account_Event;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Event_del
---------------------------------------------------------------------------
    AFTER DELETE ON Account_Event
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
