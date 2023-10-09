---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Journal_Detail (
---------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Trn_Head_ID bigint NOT NULL,
  Series int2 NULL,
  Office_ID bigint NOT NULL,
  COA_ID bigint NOT NULL,
  Account_Type_ID bigint NULL,
  Currency varchar(3) NULL,
  Partition_ID bigint NULL,
  Trn_Amt numeric(16,6) NOT NULL DEFAULT 0,
  Other_Info jsonb NULL,
  
  CONSTRAINT Journal_Detail_pkey PRIMARY KEY (Trn_Head_ID, Series),
  CONSTRAINT idxJournal_Detail_ID  UNIQUE (Trn_Head_ID, Series),  
  CONSTRAINT fkJournal_Detail_Office FOREIGN KEY (Office_ID ) REFERENCES Office(ID),
  CONSTRAINT fk_Journal_Detail_Account_Type FOREIGN KEY (Account_Type_ID) REFERENCES Account_Type(ID),
  CONSTRAINT fk_Journal_Detail_Partition FOREIGN KEY (Partition_ID ) REFERENCES Reference(ID),
  CONSTRAINT fkJournal_Detail_COA FOREIGN KEY (coa_id) REFERENCES Chartof_Account(ID),
  CONSTRAINT fkJournal_Detail_Trn_Head FOREIGN KEY (Trn_Head_ID) REFERENCES Trn_Head(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxJournal_Detail_UUID ON public.Journal_Detail(UUID);
CREATE INDEX IF NOT EXISTS Journal_Detailglid ON public.Journal_Detail USING btree (Office_ID , coa_id);

DROP TRIGGER IF EXISTS trgJournal_DetailIns on Journal_Detail;
---------------------------------------------------------------------------
CREATE TRIGGER trgJournal_DetailIns
---------------------------------------------------------------------------
    BEFORE INSERT ON Journal_Detail
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();

DROP TRIGGER IF EXISTS trgJournal_Detailupd on Journal_Detail;
---------------------------------------------------------------------------
CREATE TRIGGER trgJournal_Detailupd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Journal_Detail
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgJournal_Detail_del on Journal_Detail;
---------------------------------------------------------------------------
CREATE TRIGGER trgJournal_Detail_del
---------------------------------------------------------------------------
    AFTER DELETE ON Journal_Detail
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();



  INSERT into Journal_Detail(
      Trn_Head_ID, Series, Office_ID , coa_id, Currency, Account_Type_ID, Partition_ID , Trn_Amt)
  
  SELECT 
      h.ID Trn_Head_ID, Series, o.ID Office_ID , ca.ID coa_id, 
      acc.Currency, y.ID Account_Type_ID, par.ID Partition_ID , Trn_Amt
  FROM (Values      
      ('2af90d74-3bee-48c5-8935-443edafb8f5a'::UUID, 1, '10019', 'Cash on Hand', 'Sikap 1', 'PHP', 'FundSource', 'GSB', 100 )
      )   
    Acc(
      Trn_HeadUUID, Series, OfficeAltID, COA, Account_Type, Currency, Partition_Type, Partition_Title, Trn_Amt
      )
  INNER JOIN Trn_Head h on h.UUID = Acc.Trn_HeadUUID  
  LEFT JOIN Office o on o.Alternate_ID = Acc.Officealtid
  LEFT JOIN Chartof_Account ca on ca.Title = Acc.COA
  LEFT JOIN Account_Type y on y.Account_Type = Acc.Account_Type
  LEFT JOIN vwReference par on par.Ref_Type = Acc.Partition_Type and par.Title = Acc.Partition_Title
  
  ON CONFLICT(Trn_Head_ID, Series) DO UPDATE SET
    Trn_Head_ID = excluded.Trn_Head_ID,
    Series = excluded.Series,
    Office_ID  = excluded.Office_ID ,
    coa_id = excluded.coa_id,
    Currency = excluded.Currency,
    Account_Type_ID = excluded.Account_Type_ID,
    Partition_ID  = excluded.Partition_ID ,
    Trn_Amt = excluded.Trn_Amt
    ;
