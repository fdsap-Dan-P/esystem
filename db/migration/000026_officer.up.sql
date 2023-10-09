----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Officer (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Office_ID bigint NOT NULL,
  Officer_IIID bigint NOT NULL,
  Employee_ID bigint NULL,
  Is_Head bool NULL,
  Position varchar(50) NULL,
  Period_Start Date NOT NULL,
  Period_End Date NULL,
  Status_ID bigint NOT NULL,
  Other_Info jsonb NULL,

  CONSTRAINT Officer_pkey PRIMARY KEY (UUID),
  CONSTRAINT fk_Officer_Office FOREIGN KEY (Office_ID) REFERENCES Office(ID),
  CONSTRAINT fk_Officer_Officer FOREIGN KEY (Employee_ID) REFERENCES Employee(ID),
  CONSTRAINT fk_Officer_IIID FOREIGN KEY (Officer_IIID) REFERENCES Identity_Info(ID),
  CONSTRAINT fk_Officer_Status FOREIGN KEY (Status_ID ) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxOfficer_ofc ON public.Officer(Office_ID, Is_Head);
ALTER TABLE Officer DROP CONSTRAINT IF EXISTS officer_ishead;
ALTER TABLE Officer ADD CONSTRAINT officer_ishead CHECK (Is_Head IN (true, null) );

DROP TRIGGER IF EXISTS trgOfficer_Ins on Officer;
---------------------------------------------------------------------------
CREATE TRIGGER trgOfficer_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Officer
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();

DROP TRIGGER IF EXISTS trgOfficer_upd on Officer;
---------------------------------------------------------------------------
CREATE TRIGGER trgOfficer_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Officer
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgOfficer_del on Officer;
---------------------------------------------------------------------------
CREATE TRIGGER trgOfficer_del
---------------------------------------------------------------------------
    AFTER DELETE ON Officer
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
