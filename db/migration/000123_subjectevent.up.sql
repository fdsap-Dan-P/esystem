---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Subject_Event (
---------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Type_ID bigint NOT NULL,
  Ticket_Item_ID bigint NOT NULL,
  IIID bigint NOT NULL,
  Section_Subject_ID bigint NOT NULL,
  Event_Date Date NOT NULL, 
  Grading_Period SmallInt NOT NULL DEFAULT 0, -- (0-N/A, 1,2,3,4 Qtr, 10-Finals)
  Item_Count Numeric(8,5) NOT NULL DEFAULT 0,
  Status_ID bigint NOT NULL,
  Remarks varchar(200) NULL,
  Other_Info jsonb NULL,
 
  CONSTRAINT Subject_Event_pkey PRIMARY KEY (ID),
  CONSTRAINT fk_Subject_Event_iiid FOREIGN KEY (IIID) REFERENCES Identity_Info(ID),
  CONSTRAINT fk_Subject_Event_Type FOREIGN KEY (Type_ID) REFERENCES Reference(ID),
  CONSTRAINT fk_Subject_Event_Ticket FOREIGN KEY (Ticket_Item_ID) REFERENCES Ticket_Item(ID),
  CONSTRAINT fk_Subject_Event_Subject FOREIGN KEY (Section_Subject_ID) REFERENCES Section_Subject(ID),
  CONSTRAINT fk_Subject_Event_Status FOREIGN KEY (Status_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxSubject_Event_UUID ON public.Subject_Event(UUID);

DROP TRIGGER IF EXISTS trgSubject_Event_Ins on Subject_Event;
---------------------------------------------------------------------------
CREATE TRIGGER trgSubject_Event_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Subject_Event
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgSubject_Event_upd on Subject_Event;
---------------------------------------------------------------------------
CREATE TRIGGER trgSubject_Event_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Subject_Event
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgSubject_Event_del on Subject_Event;
---------------------------------------------------------------------------
CREATE TRIGGER trgSubject_Event_del
---------------------------------------------------------------------------
    AFTER DELETE ON Subject_Event
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
  
