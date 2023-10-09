---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Office_Account_Tran (
---------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Trn_Head_ID bigint NOT NULL,
  Series bigint NOT NULL,
  Office_Account_ID bigint NOT NULL,
  Trn_Amt  numeric(16,6) NOT NULL DEFAULT 0,
  Other_Info jsonb NULL,
  
  CONSTRAINT Office_Account_Tran_pkey PRIMARY KEY (Trn_Head_ID, Series),
 
  CONSTRAINT fkOffice_Account_TranTrn_Head FOREIGN KEY (Trn_Head_ID) REFERENCES Trn_Head(ID),
  CONSTRAINT fkOffice_Account_TranOffice_Account FOREIGN KEY (Office_Account_ID) REFERENCES Office_Account(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxOffice_Account_Tran_UUID ON public.Office_Account_Tran(UUID);

DROP TRIGGER IF EXISTS trgOffice_Account_TranIns on Office_Account_Tran;
---------------------------------------------------------------------------
CREATE TRIGGER trgOffice_Account_TranIns
---------------------------------------------------------------------------
    BEFORE INSERT ON Office_Account_Tran
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgOffice_Account_Tranupd on Office_Account_Tran;
---------------------------------------------------------------------------
CREATE TRIGGER trgOffice_Account_Tranupd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Office_Account_Tran
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgOffice_Account_Tran_del on Office_Account_Tran;
---------------------------------------------------------------------------
CREATE TRIGGER trgOffice_Account_Tran_del
---------------------------------------------------------------------------
    AFTER DELETE ON Office_Account_Tran
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();


   INSERT into Office_Account_Tran(
    Trn_Head_ID, Series, Office_Account_ID, Trn_Amt
    ) 
    SELECT
      h.ID Trn_Head_ID, Acc.Series, oa.ID Office_Account_ID, Trn_Amt
    FROM (Values
        ('2af90d74-3bee-48c5-8935-443edafb8f5a'::UUID, 1, '10019', 'Cash', 'PHP', 'FundSource', 'GSB', 100)
        )   
    Acc(
      Trn_HeadUUID, Series, OfficeAltID, Office_Account_Type, Currency, PartitionType, PartitionTitle, Trn_Amt
      )
      
    INNER JOIN Trn_Head h on h.UUID = Acc.Trn_HeadUUID  
    LEFT JOIN Office o on o.Alternate_ID = Acc.OfficeAltID
    LEFT JOIN Office_Account_Type y on y.Office_Account_Type = Acc.Office_Account_Type
    LEFT JOIN vwReference par on par.Ref_Type = Acc.PartitionType and par.Title = Acc.PartitionTitle

    LEFT JOIN Office_Account oa on oa.Office_ID  = o.ID and oa.Type_ID = y.ID 
      and oa.Currency = Acc.Currency 
      and COALESCE(oa.Partition_ID ,0) = 
          COALESCE(par.ID,0)
  
  ON CONFLICT(Trn_Head_ID, Series) DO UPDATE SET
    Trn_Head_ID = excluded.Trn_Head_ID,
    Series = excluded.Series,
    Office_Account_ID = excluded.Office_Account_ID,
    Trn_Amt = excluded.Trn_Amt;
