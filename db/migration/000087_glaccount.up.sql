---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.GL_Account (
---------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Office_ID bigint NOT NULL,
  COA_ID bigint NULL,
  Balance numeric(16,6) NOT NULL DEFAULT 0,
  Pending_Trn_Amt numeric(16,6) NOT NULL DEFAULT 0,
  Account_Type_ID bigint NULL,
  Currency varchar(3) NULL,
  Partition_ID bigint NULL,
  Remark varchar(500) NULL,
  Other_Info jsonb NULL,

  CONSTRAINT GL_Account_pkey PRIMARY KEY (ID),
  CONSTRAINT idxGL_Account_ID UNIQUE (Office_ID , coa_id),
  CONSTRAINT fk_GL_Account_COA FOREIGN KEY (coa_id) REFERENCES Chartof_Account(ID),
  CONSTRAINT fk_GL_Account_Account_Type FOREIGN KEY (Account_Type_ID) REFERENCES Account_Type(ID),
  CONSTRAINT fk_GL_Account_Partition FOREIGN KEY (Partition_ID ) REFERENCES Reference(ID),
  CONSTRAINT fk_GL_Account_Office FOREIGN KEY (Office_ID) REFERENCES Office(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxGL_Account_UUID ON public.GL_Account(UUID);
DROP TRIGGER IF EXISTS trgGL_Account_Ins on GL_Account;
---------------------------------------------------------------------------
CREATE TRIGGER trgGL_Account_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON GL_Account
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgGL_Account_upd on GL_Account;
---------------------------------------------------------------------------
CREATE TRIGGER trgGL_Account_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON GL_Account
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgGL_Account_del on GL_Account;
---------------------------------------------------------------------------
CREATE TRIGGER trgGL_Account_del
---------------------------------------------------------------------------
    AFTER DELETE ON GL_Account
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();


  INSERT into GL_Account(
    Office_ID , coa_id, Balance, Pending_Trn_Amt, 
    Currency, Account_Type_ID, Partition_ID , Remark
    ) 
 
  SELECT
    o.ID Office_ID , ca.ID coa_id, Acc.Balance, Acc.Pending_Trn_Amt, 
    acc.Currency, y.ID Account_Type_ID, par.ID Partition_ID , Acc.Remark
  FROM (Values
      ('10019', 'Cash on Hand', 0, 0, '', 'PHP', 'FundSource', 'GSB', 'Remarks')
      )   
  Acc(
    Office_altid, COA, Balance, Pending_Trn_Amt, 
    Account_Type, Currency, Partition_Type, Partition_Title, Remark
    )
  LEFT JOIN Office o on o.Alternate_ID = Acc.Office_altid
  LEFT JOIN Chartof_Account ca on ca.Title = Acc.COA
  LEFT JOIN Account_Type y on y.Account_Type = Acc.Account_Type
  LEFT JOIN vwReference par on par.Ref_Type = Acc.Partition_Type and par.Title = Acc.Partition_Title
  ON CONFLICT(Office_ID , coa_id) DO UPDATE SET
    Office_ID  = excluded.Office_ID ,
    coa_id = excluded.coa_id,
    Balance = excluded.Balance,
    Pending_Trn_Amt = excluded.Pending_Trn_Amt,
    Currency = excluded.Currency,
    Account_Type_ID = excluded.Account_Type_ID,
    Partition_ID  = excluded.Partition_ID ,
    Remark = excluded.Remark
  ;