----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Office_Account_Type (
----------------------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Office_Account_Type varchar(100) NOT NULL,
  COA_ID bigint NOT NULL,
  Other_Info jsonb NULL,

  CONSTRAINT Office_Account_Type_pkey PRIMARY KEY (ID),
  CONSTRAINT idxOffice_Account_Type_Unq UNIQUE (Office_Account_Type),
  CONSTRAINT fk_Office_Account_Type_COA FOREIGN KEY (coa_id) REFERENCES Chartof_Account(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxOffice_Account_Type_UUID ON public.Office_Account_Type(UUID);

DROP TRIGGER IF EXISTS trgOffice_Account_Type_Ins on Office_Account_Type;
---------------------------------------------------------------------------
CREATE TRIGGER trgOffice_Account_Type_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Office_Account_Type
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgOffice_Account_Type_upd on Office_Account_Type;
---------------------------------------------------------------------------
CREATE TRIGGER trgOffice_Account_Type_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Office_Account_Type
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgOffice_Account_Type_del on Office_Account_Type;
---------------------------------------------------------------------------
CREATE TRIGGER trgOffice_Account_Type_del
---------------------------------------------------------------------------
    AFTER DELETE ON Office_Account_Type
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();


----------------------------------------------------------------------------------------
CREATE OR REPLACE VIEW public.vwOffice_Account_Type
----------------------------------------------------------------------------------------
AS SELECT 
    Acctyp.ID, mr.UUID,
    Acctyp.Office_Account_Type,
    COA.ID coa_id, COA.UUID COAUUID, COA.Title COA, COA.Parent_ID,

    mr.Mod_Ctr,
    Acctyp.Other_Info,
    mr.Created,
    mr.Updated 
 
    FROM Office_Account_Type Acctyp
    JOIN Main_Record mr ON Acctyp.UUID = mr.UUID
    JOIN Chartof_Account COA ON COA.ID = Acctyp.coa_id;

 ----------------------------------------------------------------------------------------
   INSERT INTO Office_Account_Type(
       Office_Account_Type,   coa_id)    
    SELECT 
      a.Office_Account_Type, COA.ID coa_id
    FROM (Values
      ('Cash','Cash on Hand'),
      ('Bank Charge','Banking Fees'),
      ('Unrealized Interest','Interest Receivable'),
      ('Realized Interest','Interest Income'),
      ('Service Fee Income','Service Income')
      )   
      a(Office_Account_Type, COA)
    LEFT JOIN Chartof_Account COA on COA.Title = a.COA
    ON CONFLICT (Office_Account_Type) DO UPDATE SET 
      Office_Account_Type = excluded.Office_Account_Type,
      coa_id = excluded.coa_id
    ;
    