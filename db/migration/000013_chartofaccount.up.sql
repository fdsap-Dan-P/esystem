----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Chartof_Account (
----------------------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Acc varchar(21) NOT NULL,
  Active bool NOT NULL,
  Contra_Account bool NOT NULL,
  Normal_Balance bool NOT NULL,
  Title varchar(200) NOT NULL,
  Parent_ID bigint NOT NULL,
  Short_Name varchar(100) NOT NULL,
  Other_Info jsonb NULL,
  
  CONSTRAINT Chartof_Account_pkey PRIMARY KEY (ID),
  CONSTRAINT idxCOAAcc UNIQUE (Acc),
  CONSTRAINT idxCOATitle UNIQUE (Title),
  CONSTRAINT fk_COA_Parent FOREIGN KEY (Parent_ID) REFERENCES COA_Parent(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxChartof_Account_UUID ON public.Chartof_Account(UUID);

DROP TRIGGER IF EXISTS trgChartof_Account_Ins on Chartof_Account;
---------------------------------------------------------------------------
CREATE TRIGGER trgChartof_Account_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Chartof_Account
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgChartof_Account_upd on Chartof_Account;
---------------------------------------------------------------------------
CREATE TRIGGER trgChartof_Account_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Chartof_Account
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

  DROP TRIGGER IF EXISTS trgChartof_Account_del on Chartof_Account;
---------------------------------------------------------------------------
CREATE TRIGGER trgChartof_Account_del
---------------------------------------------------------------------------
    AFTER DELETE ON Chartof_Account
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();


----------------------------------------------------------------------------------------
CREATE OR REPLACE VIEW public.vwChartof_Account
----------------------------------------------------------------------------------------
AS SELECT 
    COA.ID, mr.UUID,
    COA.Acc,
    COA.Title,
    COA.Short_Name,
    
    p.ID Parent_ID, p.UUID ParentUUID, p.Title AS Parent,
    COA.Contra_Account,
    COA.Normal_Balance,
    COA.Active,
    
    mr.Mod_Ctr,
    COA.Other_Info,
    mr.Created,
    mr.Updated 
   FROM Chartof_Account COA
   INNER JOIN Main_Record mr on mr.UUID = COA.UUID
   JOIN COA_Parent p ON p.ID = COA.Parent_ID;

DO $$
DECLARE 
  Titles TEXT DEFAULT '';
  rec RECORD;
  cur CURSOR FOR 
  
  SELECT 
    a.Acc, a.Title, a.Short_Name, p.ID Parent_ID, 
    a.Contra_Account, a.Normal_Balance 
  FROM (Values
      ('101-0100-0001','Cash on Hand', 'Cash on Hand','Cash',FALSE,TRUE),
      ('101-0100-0002','Cash in Box', 'Cash in Box','Cash',FALSE,TRUE),
      ('101-0100-0003','Cash In Vault', 'Cash In Vault','Cash',FALSE,TRUE),
      ('101-0100-0004','Petty Cash Fund', 'Petty Cash Fund','Cash',FALSE,TRUE),
      ('101-0200-0001','COCI - Checks and Other Cash Items', 'Checks and Other Cash Items','Cash Equivalents',FALSE,TRUE),
      ('101-0200-0002','COCI - Returned Checks and Other Cash Items', 'Returned Checks and Other Cash Items','Cash Equivalents',FALSE,TRUE),
      ('101-0200-0003','COCI - Bangko Sentral ng Pilipinas', 'Bangko Sentral ng Pilipinas','Cash Equivalents',FALSE,TRUE),
      ('101-0200-0004','COCI - National Government', 'National Government','Cash Equivalents',FALSE,TRUE),
      ('101-0200-0005','COCI - Due from Other Banks', 'Due from Other Banks','Cash Equivalents',FALSE,TRUE),
      ('101-0200-0006','COCI - Deposits with Closed Banks / Banks in Liquidation', 'Deposits with Closed Banks / Banks in Liquidation','Cash Equivalents',FALSE,TRUE),
      ('101-0200-0007','COCI - Allowance for Probable Losses - Deposit w/ Closed Banks', 'Allowance for Probable Losses - Deposit w/ Closed Banks','Cash Equivalents',TRUE,FALSE),
      ('101-0200-0008','COCI - Miscellaneous', 'Miscellaneous','Cash Equivalents',FALSE,TRUE),
      ('102-0100-0001','DFHO/B/A ', 'DFHO/B/A ','Inter-Office Account',FALSE,TRUE),
      ('102-0400-0001','Accounts Receivable', 'Accounts Receivable','Other Receivables',FALSE,TRUE),
      ('102-0400-0002','Deficiency Judgment Receivable', 'Deficiency Judgment Receivable','Other Receivables',FALSE,TRUE),
      ('102-0400-0003','Shortages', 'Shortages','Other Receivables',FALSE,TRUE),
      ('102-0400-0004','Sales Contract Receivable', 'Sales Contract Receivable','Other Receivables',FALSE,TRUE),
      ('102-0500-0001','Allowance for Probable Losses - Accounts Receivable', 'Allowance for Probable Losses - Accounts Receivable','Adjustments on Receivables',TRUE,FALSE),
      ('104-0100-0001','Prepaid Expenses', 'Prepaid Expenses','Prepayments And Other Current Assets',FALSE,TRUE),
      ('104-0100-0002','Deferred Charges (including Goodwill)', 'Deferred Charges (including Goodwill)','Prepayments And Other Current Assets',FALSE,TRUE),
      ('104-0200-0001','Accrued Admin Fee Receivables', 'Accrued Admin Fee Receivables','Accrued Interest Receivable',FALSE,TRUE),
      ('104-0200-0002','Interest Receivable', 'Interest Receivable','Accrued Interest Receivable',FALSE,TRUE),
      ('104-0200-0003','Accrued Interest Receivable', 'Accrued Interest Receivable','Accrued Interest Receivable',FALSE,TRUE),
      ('104-0200-0004','Allowance for Uncollected Interest on Loans', 'Allowance for Uncollected Interest on Loans','Accrued Interest Receivable',FALSE,FALSE),
      ('104-0300-0001','Deferred Tax', 'Deferred Tax','Tax Related Receivables',FALSE,TRUE),
      ('104-0300-0002','Input Tax', 'Input Tax','Tax Related Receivables',FALSE,TRUE),
      ('104-0900-0001','Miscellaneous Assets', 'Miscellaneous Assets','Additional, Other And Miscellaneous Assets',FALSE,TRUE),
      ('104-0900-0002','Postage Stamps', 'Postage Stamps','Additional, Other And Miscellaneous Assets',FALSE,TRUE),
      ('104-0900-0003','Documentary Stamps', 'Documentary Stamps','Additional, Other And Miscellaneous Assets',FALSE,TRUE),
      ('106-0203-0001','Held-to-maturity investments', 'Held-to-maturity investments','Held-to-maturity investments',FALSE,TRUE),
      ('106-0204-0001','Microfinance Loan', 'Microfinance Loan','Loans and Receivables',FALSE,TRUE),
      ('106-0204-0002','SME Loan', 'SME Loan','Loans and Receivables',FALSE,TRUE),
      ('106-0204-0003','Restructured Loan', 'Restructured Loan','Loans and Receivables',FALSE,TRUE),
      ('106-0204-0004','Other Loan', 'Other Loan','Loans and Receivables',FALSE,TRUE),
      ('106-0204-0005','Allowance for Probable Losses - ROPA', 'Allowance for Probable Losses - ROPA','Loans and Receivables',TRUE,FALSE),
      ('106-0204-0006','General Loan Loss Provision', 'General Loan Loss Provision','Loans and Receivables',TRUE,FALSE),
      ('106-0204-0007','Allowance for Probable Losses', 'Allowance for Probable Losses','Loans and Receivables',TRUE,FALSE),
      ('106-0205-0001','IBODI - Government', 'Government','Investment in Bonds',FALSE,TRUE),
      ('106-0205-0002','IBODI - Private', 'Private','Investment in Bonds',FALSE,TRUE),
      ('106-0205-0003','Allowance for Probable Losses - IBODI', 'Allowance for Probable Losses - IBODI','Investment in Bonds',TRUE,FALSE),
      ('106-0205-0004','Accumulated Bond Discount - Private', 'Accumulated Bond Discount - Private','Investment in Bonds',TRUE,FALSE),
      ('106-0205-0005','Accumulated Premium Amortization - Private', 'Accumulated Premium Amortization - Private','Investment in Bonds',TRUE,FALSE),
      ('106-0205-0006','Accumulated Premium Amortization - IBODI Govt', 'Accumulated Premium Amortization - IBODI Govt','Investment in Bonds',TRUE,FALSE),
      ('106-0205-0007','Accumulated Bond Discount - IBODI Govt', 'Accumulated Bond Discount - IBODI Govt','Investment in Bonds',TRUE,FALSE),
      ('106-0206-0001','Equity Investments in Financial Allied Undertakings', 'Equity Investments in Financial Allied Undertakings','Equity Investment',FALSE,TRUE),
      ('106-0206-0002','Equity Investments in Non-Financial Allied Undertakings', 'Equity Investments in Non-Financial Allied Undertakings','Equity Investment',FALSE,TRUE),
      ('106-0206-0003','Allowance for Probable Losses - Equity Investments', 'Allowance for Probable Losses - Equity Investments','Equity Investment',TRUE,FALSE),
      ('106-0400-0001','Other Investments', 'Other Investments','Other Financial Assets',FALSE,TRUE),
      ('106-0400-0002','Allowance for Probable Losses - Others', 'Allowance for Probable Losses - Others','Other Financial Assets',TRUE,FALSE),
      ('109-0101-0001','Land', 'Land','Land',FALSE,TRUE),
      ('109-0301-0001','Appraisal Increment (With MB Approval)', 'Appraisal Increment (With MB Approval)','Buildings',FALSE,TRUE),
      ('109-0301-0002','Building', 'Building','Buildings',FALSE,TRUE),
      ('109-0301-0003','Accumulated Depreciation - Building', 'Accumulated Depreciation - Building','Buildings',TRUE,FALSE),
      ('109-0301-0004','Accumulated Depreciation - ROPA Building', 'Accumulated Depreciation - ROPA Building','Buildings',TRUE,FALSE),
      ('109-0301-0005','Transportation Equipment', 'Transportation Equipment','Buildings',FALSE,TRUE),
      ('109-0301-0006','Accumulated Depreciation - Bank Premises - Appraisal Increment', 'Accumulated Depreciation - Bank Premises - Appraisal Increment','Buildings',TRUE,FALSE),
      ('109-0200-0001','Leasehold Rights and Improvements', 'Leasehold Rights and Improvements','Leasehold Rights and Improvements',FALSE,TRUE),
      ('109-0200-0002','Amortization - Leasehold Rights and Improvement', 'Amortization - Leasehold Rights and Improvement','Leasehold Rights and Improvements',FALSE,FALSE),
      ('109-0400-0001','Information Technology Equipment', 'Information Technology Equipment','Machinery And Equipment',FALSE,TRUE),
      ('109-0400-0002','Accumulated Depreciation - Information Technology Equipment', 'Accumulated Depreciation - Information Technology Equipment','Machinery And Equipment',TRUE,FALSE),
      ('109-0500-0001','Accumulated Depreciation - Transportation Equipment', 'Accumulated Depreciation - Transportation Equipment','Vehicles',TRUE,FALSE),
      ('109-0600-0001','Accumulated Depreciation - Furnitures & Fixtures', 'Accumulated Depreciation - Furnitures & Fixtures','Furniture And Fixtures',TRUE,FALSE),
      ('109-0600-0002','Furniture, Fixtures and Equipment', 'Furniture, Fixtures and Equipment','Furniture And Fixtures',FALSE,TRUE),
      ('109-0800-0001','ROPOA', 'ROPOA','Other Property, Plant And Equipment',FALSE,TRUE),
      ('109-0800-0002','ROPA - Real and Other Properties Owned or Acquired', 'ROPA - Real and Other Properties Owned or Acquired','Other Property, Plant And Equipment',FALSE,TRUE),
      ('109-0900-0001','Building Under Construction', 'Building Under Construction','Construction In Progress',FALSE,TRUE),
      ('201-0001','Bills Payable ', 'Bills Payable ','Financial Liabilities',FALSE,FALSE),
      ('201-0002','Savings Deposit', 'Savings Deposit','Financial Liabilities',FALSE,FALSE),
      ('202-0100-0001','Retirement liability', 'Retirement liability','Employee Benefits Provisions',FALSE,FALSE),
      ('204-0102-0001','Accrued Interest Payable', 'Accrued Interest Payable','Accrued Interest Expenses',FALSE,FALSE),
      ('204-0109-0001','AIP - Accrued Interest, Taxes, Fringe Benefits & Other Expense Payable', 'Other Accrued Expenses','Other Accrued Expenses',FALSE,FALSE),
      ('205-0100-0001','Accrued Income Tax Payable', 'Accrued Income Tax Payable','Current Tax Liabilities',FALSE,FALSE),
      ('204-0109-0001','Accrued Expense', 'Accrued Expense','Other Accrued Expenses',FALSE,FALSE),
      ('203-0600-0001','Dividends Payable', 'Dividends Payable','Dividend Payables',FALSE,FALSE),
      ('204-0200-0001','Other Deferred Credits', 'Other Deferred Credits','Deferred Income (Unearned Revenue) ',FALSE,FALSE),
      ('203-1000-0001','Accounts Payable', 'Accounts Payable','Other Payables',FALSE,FALSE),
      ('203-1000-0002','Overages', 'Overages','Other Payables',FALSE,FALSE),
      ('203-1000-0003','Pag-ibig Loans Payable', 'Pag-ibig Loans Payable','Other Payables',FALSE,FALSE),
      ('203-1000-0004','Pag-ibig Fund Premium', 'Pag-ibig Fund Premium','Other Payables',FALSE,FALSE),
      ('203-1000-0005','Philhealth Premium Payable', 'Philhealth Premium Payable','Other Payables',FALSE,FALSE),
      ('203-1000-0006','AP - Documentary Stamps', 'AP - Documentary Stamps','Other Payables',FALSE,FALSE),
      ('203-1000-0007','Other Credits - Unclaimed Balances', 'Other Credits - Unclaimed Balances','Other Payables',FALSE,FALSE),
      ('203-1000-0008','Other Credits - Dormant', 'Other Credits - Dormant','Other Payables',FALSE,FALSE),
      ('203-1000-0009','AP - Other', 'AP - Other','Other Payables',FALSE,FALSE),
      ('206-0700-0001','Witholding Tax Payable', 'Witholding Tax Payable','Other Liabilities',FALSE,FALSE),
      ('206-0700-0002','Tempo Account - Credit', 'Tempo Account - Credit','Other Liabilities',FALSE,FALSE),
      ('206-0700-0003','Tempo Account - Debit', 'Tempo Account - Debit','Other Liabilities',FALSE,FALSE),
      ('206-0700-0004','Deposit for Stock Subscription', 'Deposit for Stock Subscription','Other Liabilities',FALSE,FALSE),
      ('206-0700-0005','Miscellaneous Liabilities', 'Miscellaneous Liabilities','Other Liabilities',FALSE,FALSE),
      ('301-0100-0001','Capital Stock - Common - Paid-in - Government', 'Common - Paid-in - Government','Ordinary Shares',FALSE,FALSE),
      ('301-0100-0002','Capital Stock - Common - Subscribed', 'Common - Subscribed','Ordinary Shares',FALSE,FALSE),
      ('301-0100-0003','Capital Stock - Common - Paid-in - Private', 'Common - Paid-in - Private','Ordinary Shares',FALSE,FALSE),
      ('301-0100-0004','Capital Stock - Common - Subscriptions Receivable', 'Common - Subscriptions Receivable','Ordinary Shares',FALSE,FALSE),
      ('301-0200-0001','Capital Stock - Preferred - Paid-in - Government', 'Preferred - Paid-in - Government','Preferred Shares',FALSE,FALSE),
      ('301-0200-0002','Capital Stock - Preferred - Subscribed', 'Preferred - Subscribed','Preferred Shares',FALSE,FALSE),
      ('301-0200-0003','Capital Stock - Preferred - Paid-in - Private', 'Preferred - Paid-in - Private','Preferred Shares',FALSE,FALSE),
      ('301-0200-0004','Capital Stock - Preferred - Subscriptions Receivable', 'Preferred - Subscriptions Receivable','Preferred Shares',FALSE,FALSE),
      ('301-0100-0001','Paid-in Surplus', 'Paid-in Surplus','Ordinary Shares',FALSE,FALSE),
      ('303-1500-0001','Reserve for Contingencies', 'Reserve for Contingencies','Reserve For Catastrophe',FALSE,FALSE),
      ('303-1500-0002','Other Surplus Reserves', 'Other Surplus Reserves','Reserve For Catastrophe',FALSE,FALSE),
      ('303-1500-0003','Reserve for Self - Insurance', 'Reserve for Self - Insurance','Reserve For Catastrophe',FALSE,FALSE),
      ('304-0100-0001','Outward Bills for Collection', 'Outward Bills for Collection','Other Equity Interest',FALSE,FALSE),
      ('304-0100-0002','Other Contingent Accounts', 'Other Contingent Accounts','Other Equity Interest',FALSE,FALSE),
      ('304-0100-0003','Appraisal Increment Reserve', 'Appraisal Increment Reserve','Other Equity Interest',FALSE,FALSE),
      ('304-0100-0004','Inward Bills for Collection', 'Inward Bills for Collection','Other Equity Interest',FALSE,FALSE),
      ('304-0100-0005','Late Deposits/Payments Received', 'Late Deposits/Payments Received','Other Equity Interest',FALSE,FALSE),
      ('304-0100-0006','Other Comprehensive Income', 'Other Comprehensive Income','Other Equity Interest',FALSE,FALSE),
      ('304-0100-0007','Deficiency Claims Receivable', 'Deficiency Claims Receivable','Other Equity Interest',FALSE,FALSE),
      ('304-0100-0008','Contingent Contra Account', 'Contingent Contra Account','Other Equity Interest',FALSE,FALSE),
      ('304-0100-0009','Items Held as Collateral', 'Items Held as Collateral','Other Equity Interest',FALSE,FALSE),
      ('407-0200-0001','Interest Income', 'Interest Income','Interest Income',FALSE,FALSE),
      ('409-9003-0001','Rental - Bank Premises & Equipment', 'Rental - Bank Premises & Equipment','Rental Income',FALSE,FALSE),
      ('402-0001','Service Charges/Fees', 'Service Charges/Fees - Others','Revenue from Services',FALSE,FALSE),
      ('402-0002','Service Income', 'Service Income','Revenue from Services',FALSE,FALSE),
      ('504-0100-0001','Commission from Microinsurance', 'Commission from Microinsurance','Other Revenue',FALSE,FALSE),
      ('504-0100-0002','Imputed Income', 'Imputed Income','Other Revenue',FALSE,FALSE),
      ('504-0100-0003','Miscellaneous Income', 'Miscellaneous Income','Other Revenue',FALSE,FALSE),
      ('504-0100-0004','Bank Commissions', 'Bank Commissions','Other Revenue',FALSE,FALSE),
      ('504-0100-0005','Sale of Sunday School Materials', 'Sale of Sunday School Materials','Other Revenue',FALSE,FALSE),
      ('504-0100-0006','Sale of UCCP Planner', 'Sale of UCCP Planner','Other Revenue',FALSE,FALSE),
      ('504-0100-0007','Sale of Certificates', 'Sale of Certificates','Other Revenue',FALSE,FALSE),
      ('410-0001','Income from Donation', 'Income from Donation','Income from Donation',FALSE,TRUE),
      ('501-0200-0001','Salaries and Benefits', 'Salaries and Benefits','Salaries and Benefits',FALSE,TRUE),
      ('501-090001-0001','Travel', 'Travel','Travel',FALSE,TRUE),
      ('501-1100-0001','Office Supplies', 'Office Supplies','Office Supplies',FALSE,TRUE),
      ('501-1200-0001','Representation', 'Representation','Representation',FALSE,TRUE),
      ('501-210002-0001','Staff House Rental', 'Staff House Rental','Staff House Rental',FALSE,TRUE),
      ('501-0500-0001','Office Equipment/Repair and Maintenance', 'Repairs and Maintenance','Repairs and Maintenance',FALSE,TRUE),
      ('501-0600-0001','Electricity/Water/Internet', 'Power, Light and Water','Power, Light and Water',FALSE,TRUE),
      ('501-0700-0001','Communication and Postage', 'Communication and Postage','Communication and Postage',FALSE,TRUE),
      ('501-0800-0001','Telephone', 'Telephone','Telephone',FALSE,TRUE),
      ('501-090002-0001','Car Maintenance', 'Car Maintenance','Car Maintenance',FALSE,TRUE),
      ('501-090003-0001','Driver''s Honorarium', 'Driver''s Honorarium','Driver''s Honorarium',FALSE,TRUE),
      ('501-1000-0001','Benevolence', 'Benevolence','Benevolence',FALSE,FALSE),
      ('501-020206-0001','Medical Assistance', 'Medical Assistance','Medical Assistance',FALSE,FALSE),
      ('501-1900-0001','Banking Fees', 'Banking Fees','Banking Fees',FALSE,FALSE)
      )   
    a(Acc, Title, Short_Name, COA_Parent, Contra_Account, Normal_Balance)
  INNER JOIN COA_Parent p on lower(p.Title) = lower(a.COA_Parent)
  ;
  
  i integer;
  _ID integer;
BEGIN
   -- Open the cursor
   Open cur;
   
   LOOP

      FETCH cur INTO rec;
    -- exit when no more row to fetch
      EXIT WHEN NOT FOUND;
 
      INSERT INTO Chartof_Account(
         Acc,   Title,   Short_Name, Parent_ID,   
         Contra_Account,   Normal_Balance,   Active) 
      SELECT
        rec.Acc, rec.Title, rec.Short_Name, rec.Parent_ID, 
        rec.Contra_Account, rec.Normal_Balance, TRUE
      ON CONFLICT(Acc) DO UPDATE SET
         Title = excluded.Title,
         Short_Name = excluded.Short_Name,
         Parent_ID = excluded.Parent_ID,   
         Contra_Account = excluded.Contra_Account,
         Normal_Balance = excluded.Normal_Balance,
         Active = excluded.Active ;

   END LOOP;
  
   -- Close the cursor
   CLOSE cur;

END; $$ ;
