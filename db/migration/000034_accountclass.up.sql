----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Account_Class (
----------------------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Product_ID bigint NOT NULL,
  Group_ID bigint NOT NULL,
  Class_ID bigint NOT NULL,
  Cur_ID bigint NOT NULL,
  NonCur_ID bigint NULL,
  BS_Acc_ID bigint NULL,  
  IS_Acc_ID bigint NULL,
  Other_Info jsonb NULL,
  CONSTRAINT Account_Class_pkey PRIMARY KEY (ID),
  CONSTRAINT idxAccount_Class_Unq UNIQUE (Product_ID, Group_ID, Class_ID),
  CONSTRAINT fk_Account_Class_TypeGroup FOREIGN KEY (Group_ID) REFERENCES Reference(ID),
  CONSTRAINT fk_Account_Class_Account_Class FOREIGN KEY (Class_ID) REFERENCES Reference(ID),
  CONSTRAINT fk_Account_Class_cur FOREIGN KEY (cur_id) REFERENCES Chartof_Account(ID),
  CONSTRAINT fk_Account_Class_noncur FOREIGN KEY (Noncur_id) REFERENCES Chartof_Account(ID),
  CONSTRAINT fk_Account_Class_BS FOREIGN KEY (BS_Acc_ID) REFERENCES Chartof_Account(ID),
  CONSTRAINT fk_Account_Class_IS FOREIGN KEY (IS_Acc_ID) REFERENCES Chartof_Account(ID),
  CONSTRAINT fk_Account_Class_Product FOREIGN KEY (Product_ID) REFERENCES Product(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccount_Class_UUID ON public.Account_Class(UUID);
-- drop table Account_Class cascade
DROP TRIGGER IF EXISTS trgAccount_Class_Ins on Account_Class;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Class_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Account_Class
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgAccount_Class_upd on Account_Class;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Class_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Account_Class
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgAccount_Class_del on Account_Class;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Class_del
---------------------------------------------------------------------------
    AFTER DELETE ON Account_Class
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();


----------------------------------------------------------------------------------------
CREATE OR REPLACE VIEW public.vwAccount_Class
----------------------------------------------------------------------------------------
AS SELECT 
    Acc.ID, mr.UUID,
    
    md.ID Product_ID, md.Product_Name,
    
    Accgrp.ID Group_ID, Accgrp.UUID GroupUUID, Accgrp.Title AccountGroup,

    COAC.ID cur_id, COAC.UUID CurUUID, COAC.Acc CurrentAcc, COAC.Title CurrentTitle,
    COAN.ID Noncur_id, COAN.UUID NonCurUUID, COAN.Acc Non_CurrentAcc, COAN.Title Non_CurrentTitle,

    b.ID BS_Acc_ID, b.UUID BS_AccUUID, b.Acc BS_Acc, b.Title BSTitle, 
    i.ID IS_Acc_ID, i.UUID IS_AccUUID, i.Acc IS_Acc, b.Title ISTitle, 
    
    AccClass.UUID Class_ID, AccClass.Title Account_Class,
    
    mr.Mod_Ctr,
    Acc.Other_Info,
    mr.Created,
    mr.Updated 
    
   FROM Account_Class Acc
   INNER JOIN Main_Record mr on mr.UUID = Acc.UUID   
   JOIN Product md on Acc.Product_ID = md.ID
   JOIN Reference Accgrp ON Accgrp.ID = Acc.Group_ID
   JOIN Reference AccClass ON AccClass.ID = Acc.Class_ID
   JOIN Chartof_Account COAC ON COAC.ID = Acc.cur_id
   LEFT JOIN Chartof_Account COAN ON COAN.ID = Acc.Noncur_id
   LEFT JOIN Chartof_Account b ON b.ID = Acc.BS_Acc_ID
   LEFT JOIN Chartof_Account i ON i.ID = Acc.IS_Acc_ID
 ;

  INSERT INTO 
    Account_Class(Product_ID, Group_ID, Class_ID, cur_id, Noncur_id, IS_Acc_ID, BS_Acc_ID) 
  SELECT 
    md.ID Product_ID, Accgrp.ID Group_ID,  AccClass.ID Class_ID, 
    COAC.ID cur_id, COAC.ID Noncur_id, i.ID IS_Acc_ID, b.ID BS_Acc_ID
  FROM (Values
  
      ('Loan','Microfinance','Current','LoanClass','Other Loan',NULL,'Accrued Interest Receivable','Interest Income'),
      ('Loan','Microfinance','Past Due','LoanClass','Other Loan',NULL,'Accrued Interest Receivable','Interest Income'),
      ('Savings','Microfinance','Active','SavingsClass','Savings Deposit', NULL, NULL, NULL),
      ('Savings','Microfinance','Dormant','SavingsClass','Savings Deposit', NULL, NULL, NULL),
      ('Collecting Facility','MicroInsurance','Current','AccountClass','Accounts Payable', NULL, NULL, NULL),
      ('Donation','Other','Restricted','DonationClass','Income from Donation',NULL,NULL,NULL),
      ('Donation','Other','UnRestricted','DonationClass','Income from Donation',NULL,NULL,NULL),
      ('Receivable','Other','Current','LoanClass','Accounts Receivable',NULL,'Accrued Interest Receivable','Interest Income')      ) 
    Acc(Product, Account_TypeGroup, Account_Class, AccClass, Cur, NonCur, BS_Acc, IS_Acc)
  LEFT JOIN Product md on Acc.Product = md.Product_Name
  LEFT JOIN vwReference Accgrp 
    on lower(Accgrp.Title) = lower(Acc.Account_TypeGroup)
    and lower(Accgrp.Ref_Type) = lower('AccountTypeGroup')
 
  LEFT JOIN Reference AccClass 
    on lower(AccClass.Title) = lower(Acc.Account_Class)
    and lower(AccClass.Ref_Type) = lower(Acc.AccClass)
  
  LEFT JOIN Chartof_Account COAC on lower(COAC.Title) = lower(Acc.Cur)
  LEFT JOIN Chartof_Account COAN on lower(COAN.Title) = lower(Acc.NonCur)
  LEFT JOIN Chartof_Account b on lower(b.Title) = lower(Acc.BS_Acc)
  LEFT JOIN Chartof_Account i on lower(i.Title) = lower(Acc.IS_Acc)

  ON CONFLICT (Product_ID, Group_ID, Class_ID)
  DO UPDATE SET 
    cur_id = excluded.cur_id,
    Noncur_id = excluded.Noncur_id,
    BS_Acc_ID = excluded.BS_Acc_ID,
    IS_Acc_ID = excluded.IS_Acc_ID
  ;