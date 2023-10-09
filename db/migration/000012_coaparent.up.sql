---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.COA_Parent (
---------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Acc varchar(21) NOT NULL,
  COA_Seq bigint NULL,
  Title varchar(200) NOT NULL,
  Parent_ID bigint NULL,
  Other_Info jsonb NULL,
  
  CONSTRAINT COA_Parent_pkey PRIMARY KEY (ID),
  CONSTRAINT idxCOA_Parent_Title UNIQUE (Title),
  CONSTRAINT idxCOA_Parent_Acc UNIQUE (Acc),
  CONSTRAINT fk_COA_Parent_Parent FOREIGN KEY (Parent_ID) REFERENCES COA_Parent(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxCOA_Parent_ID ON public.COA_Parent(UUID);

CREATE UNIQUE INDEX IF NOT EXISTS idx_COA_Parent ON public.COA_Parent 
  USING btree ( COALESCE(Parent_ID,0), lower((Title)::text) ) ;

DROP TRIGGER IF EXISTS trgCOA_Parent_Ins on COA_Parent;
---------------------------------------------------------------------------
CREATE TRIGGER trgCOA_Parent_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON COA_Parent
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgCOA_Parent_upd on COA_Parent;
---------------------------------------------------------------------------
CREATE TRIGGER trgCOA_Parent_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON COA_Parent
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

  DROP TRIGGER IF EXISTS trgCOA_Parent_del on COA_Parent;
---------------------------------------------------------------------------
CREATE TRIGGER trgCOA_Parent_del
---------------------------------------------------------------------------
    AFTER DELETE ON COA_Parent
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();


DO $$
DECLARE 
  titles TEXT DEFAULT '';
  rec   RECORD;
  cur CURSOR FOR 

 with RECURSIVE coa as
 (SELECT acc, title, parent FROM (Values
      ('100','Asset', null),
      ('101','Cash And Cash Equivalents', 'Asset'),
      ('101-0100','Cash', 'Cash And Cash Equivalents'),
      ('101-0200','Cash Equivalents', 'Cash And Cash Equivalents'),
      ('101-0300','Short-Term Investments', 'Cash And Cash Equivalents'),
      ('101-0400','Other Cash And Cash Equivalents', 'Cash And Cash Equivalents'),
      ('102','Receivables', 'Asset'),
      ('102-0100','Inter-Office Account', 'Receivables'),
      ('102-0200','Trade Receivables', 'Receivables'),
      ('102-0300','Contract Assets', 'Receivables'),
      ('102-0400','Other Receivables', 'Receivables'),
      ('102-0500','Adjustments on Receivables', 'Receivables'),
      ('103','Inventories', 'Asset'),
      ('103-0100','Raw Materials and Production Supplies', 'Inventories'),
      ('103-0200','Merchandise Inventory', 'Inventories'),
      ('103-0300','Food and Beverage', 'Inventories'),
      ('103-0400','Agricultural Produce', 'Inventories'),
      ('103-0500','Work In Progress', 'Inventories'),
      ('103-0600','Materials and Supplies To Be Consumed In Production Process Or Rendering Services', 'Inventories'),
      ('103-0700','Finished Goods', 'Inventories'),
      ('103-0800','Packaging And Storage Materials', 'Inventories'),
      ('103-0900','Spare Parts', 'Inventories'),
      ('103-1000','Fuel', 'Inventories'),
      ('103-1100','Property Intended For Sale In Ordinary Course Of Business', 'Inventories'),
      ('103-1200','Inventories In Transit', 'Inventories'),
      ('103-1300','Other Inventories', 'Inventories'),
      ('103-1400','Inventories Pledged As Security For Liabilities', 'Inventories'),
      ('103-1500','Inventories At Fair Value Less Costs To Sell', 'Inventories'),
      ('103-1600','Acquisition of Inventories In Progress', 'Inventories'),
      ('103-1700','Additional Inventory Items', 'Inventories'),
      ('104','Accrued And Other Assets', 'Asset'),
      ('104-0100','Prepayments And Other Current Assets', 'Accrued And Other Assets'),
      ('104-0200','Accrued Interest Receivable', 'Accrued And Other Assets'),
      ('104-0300','Tax Related Receivables', 'Accrued And Other Assets'),
      ('104-0400','Service Providers', 'Accrued And Other Assets'),
      ('104-0500','Construction Contract Asset', 'Accrued And Other Assets'),
      ('104-0600','Set Up Costs', 'Accrued And Other Assets'),
      ('104-0700','Restricted Assets', 'Accrued And Other Assets'),
      ('104-0800','Current Investments Not Classified As Cash Equivalents', 'Accrued And Other Assets'),
      ('104-0900','Additional, Other And Miscellaneous Assets', 'Accrued And Other Assets'),
      ('105','Biological Assets', 'Asset'),
      ('105-0100','Biological Assets At Cost', 'Biological Assets'),
      ('105-0200','Biological Assets At Fair Value', 'Biological Assets'),
      ('106','Financial Assets', 'Asset'),
      ('106-0100','Group Companies (Intercompany Investments)', 'Financial Assets'),
      ('106-0200','Investments And Financial Instruments', 'Financial Assets'),
      ('106-0201','Trade Securities', 'Investments And Financial Instruments'),
      ('106-0202','Available-for-sale Financial Assets', 'Investments And Financial Instruments'),
      ('106-0203','Held-to-maturity investments', 'Investments And Financial Instruments'),
      ('106-0204','Loans and Receivables', 'Investments And Financial Instruments'),
      ('106-0205','Investment in Bonds', 'Investments And Financial Instruments'),
      ('106-0206','Equity Investment', 'Investments And Financial Instruments'),
      ('106-0300','Derivative Financial Assets', 'Financial Assets'),
      ('106-0400','Other Financial Assets', 'Financial Assets'),
      ('106-0500','Allowance For Credit Losses (Aggregate)', 'Financial Assets'),
      ('106-0600','Financial Assets Classified By Designation', 'Financial Assets'),
      ('107','Intangible Assets', 'Asset'),
      ('107-0100','Intellectual Property', 'Intangible Assets'),
      ('107-0200','Computer Software', 'Intangible Assets'),
      ('107-0300','Trade And Distribution Assets', 'Intangible Assets'),
      ('107-0400','Contracts And Rights', 'Intangible Assets'),
      ('107-0500','Right To Use Assets (Classified By Type)', 'Intangible Assets'),
      ('107-0600','Other Intangible Assets', 'Intangible Assets'),
      ('107-0700','Acquisition of Intangibles In Progress', 'Intangible Assets'),
      ('107-9000','Goodwill', 'Asset'),
      ('108','Investment Property', 'Asset'),
      ('108-0100','Investment Property Under Construction Or Development', 'Investment Property'),
      ('108-0200','Investment Property Completed', 'Investment Property'),
      ('109','Property, Plant And Equipment', 'Asset'),
      ('109-0100','Land And Land Improvements', 'Property, Plant And Equipment'),
      ('109-0101','Land', 'Land And Land Improvements'),
      ('109-0102','Land Improvements', 'Land And Land Improvements'),
      ('109-0200','Leasehold Rights and Improvements', 'Property, Plant And Equipment'),
      ('109-0300','Buildings, Structures And Improvements', 'Property, Plant And Equipment'),
      ('109-0301','Buildings', 'Property, Plant And Equipment'),
      ('109-0302','Structures', 'Property, Plant And Equipment'),
      ('109-0303','Improvements', 'Property, Plant And Equipment'),
      ('109-0400','Machinery And Equipment', 'Property, Plant And Equipment'),
      ('109-0500','Vehicles', 'Property, Plant And Equipment'),
      ('109-0600','Furniture And Fixtures', 'Property, Plant And Equipment'),
      ('109-0700','Exploration And Evaluation Assets', 'Property, Plant And Equipment'),
      ('109-0800','Other Property, Plant And Equipment', 'Property, Plant And Equipment'),
      ('109-0900','Construction In Progress', 'Property, Plant And Equipment'),
      ('109-2000','Retirement Asset (Liability)', 'Property, Plant And Equipment'),
      ('200','Liabilities', null),
      ('201','Financial Liabilities', 'Liabilities'),
      ('201-0010','Bank Deposits', 'Financial Liabilities'),
      ('201-0100','Notes Payable', 'Financial Liabilities'),
      ('201-0200','Loans Payable', 'Financial Liabilities'),
      ('201-0300','Bonds (Debentures)', 'Financial Liabilities'),
      ('201-0400','Other Debts And Borrowings', 'Financial Liabilities'),
      ('201-0500','Lease Obligations', 'Financial Liabilities'),
      ('201-0600','Derivative Financial Liabilities', 'Financial Liabilities'),
      ('202','Provisions', 'Liabilities'),
      ('202-0100','Employee Benefits Provisions', 'Provisions'),
      ('202-0200','Customer Related Provisions', 'Provisions'),
      ('202-0300','Warranties', 'Provisions'),
      ('202-0400','Refunds', 'Provisions'),
      ('202-0500','Decommissioning, Restoration And Rehabilitation', 'Provisions'),
      ('202-0600','Restructuring', 'Provisions'),
      ('202-0700','Onerous Contracts', 'Provisions'),
      ('202-0800','Ligation And Regulatory Provisions', 'Provisions'),
      ('202-0900','Business Combinations', 'Provisions'),
      ('202-1000','Liabilities Included In Disposal Groups', 'Provisions'),
      ('202-1100','Other Provisions', 'Provisions'),
      ('203','Trade And Other Payables', 'Liabilities'),
      ('203-0100','Trade Payables', 'Trade And Other Payables'),
      ('203-0200','Contract Liabilities', 'Trade And Other Payables'),
      ('203-0300','Related Party Payables', 'Trade And Other Payables'),
      ('203-0400','Retention Payables', 'Trade And Other Payables'),
      ('203-0500','Adjustments on Trade and Other Payables', 'Trade And Other Payables'),
      ('203-0600','Dividend Payables', 'Trade And Other Payables'),
      ('203-0700','Interest Payable', 'Trade And Other Payables'),
      ('203-0800','Advances', 'Trade And Other Payables'),
      ('203-0900','Construction Contract Liability', 'Trade And Other Payables'),
      ('203-1000','Other Payables', 'Trade And Other Payables'),
      ('204','Accrued, Deferred And Other Liabilities', 'Liabilities'),
      ('204-0100','Accrued Expenses', 'Accrued, Deferred And Other Liabilities'),
      ('204-0101','Accrued Salaries Expenses', 'Accrued Expenses'),
      ('204-0102','Accrued Interest Expenses', 'Accrued Expenses'),
      ('204-0109','Other Accrued Expenses', 'Accrued Expenses'),
      ('204-0200','Deferred Income (Unearned Revenue) ', 'Accrued, Deferred And Other Liabilities'),
      ('204-0300','Accrued Taxes (Other Than Payroll)', 'Accrued, Deferred And Other Liabilities'),
      ('204-0400','Other (Non-Financial) Liabilities', 'Accrued, Deferred And Other Liabilities'),
      ('205','Tax Liabilities', 'Liabilities'),
      ('205-0100','Current Tax Liabilities', 'Tax Liabilities'),
      ('205-0200','Deferred Tax Liabilities', 'Tax Liabilities'),
      ('206','Other And Miscellaneous Liabilities', 'Liabilities'),
      ('206-0100','Finance Leases', 'Other And Miscellaneous Liabilities'),
      ('206-0200','Deposits', 'Other And Miscellaneous Liabilities'),
      ('206-0400','Government Grant Obligations', 'Other And Miscellaneous Liabilities'),
      ('206-0500','Liabilities Due To Central Banks', 'Other And Miscellaneous Liabilities'),
      ('206-0600','Subordinated Liabilities', 'Other And Miscellaneous Liabilities'),
      ('206-0700','Other Liabilities', 'Other And Miscellaneous Liabilities'),
      ('300','Equity', null),
      ('301','Issued Capital', 'Equity'),
      ('301-0100','Ordinary Shares', 'Issued Capital'),
      ('301-0200','Preferred Shares', 'Issued Capital'),
      ('301-0300','Par Value Per Share', 'Issued Capital'),
      ('301-0400','Share Premium', 'Issued Capital'),
      ('301-0500','Additional Paid In Capital', 'Issued Capital'),
      ('302','Retained Earnings', 'Equity'),
      ('302-0100','Current Year''s Retained Profit (Loss)', 'Retained Earnings'),
      ('302-0200','Appropriated', 'Retained Earnings'),
      ('302-0300','Unappropriated', 'Retained Earnings'),
      ('302-0400','In Suspense', 'Retained Earnings'),
      ('303','Other Reserves (Accumulated Other Comprehensive Income)', 'Equity'),
      ('303-0100','Revaluation Surplus', 'Other Reserves (Accumulated Other Comprehensive Income)'),
      ('303-0200','Reserve Of Exchange Differences On Translation', 'Other Reserves (Accumulated Other Comprehensive Income)'),
      ('303-0300','Reserve Of Cash Flow Hedges', 'Other Reserves (Accumulated Other Comprehensive Income)'),
      ('303-0400','Reserve Of Gains And Losses On Hedging Instruments That Hedge Investments In Equity Instruments', 'Other Reserves (Accumulated Other Comprehensive Income)'),
      ('303-0500','Reserve Of Change In Value Of Time Value Of Options', 'Other Reserves (Accumulated Other Comprehensive Income)'),
      ('303-0600','Reserve Of Change In Value Of Forward Elements Of Forward Contracts', 'Other Reserves (Accumulated Other Comprehensive Income)'),
      ('303-0700','Reserve Of Change In Value Of Foreign Currency Basis Spreads', 'Other Reserves (Accumulated Other Comprehensive Income)'),
      ('303-0800','Reserve Of Gains And Losses On Financial Assets Measured At Fair Value Through Other Comprehensive Income', 'Other Reserves (Accumulated Other Comprehensive Income)'),
      ('303-0900','Reserve Of Gains And Losses On Remeasuring Available-For-Sale Financial Assets', 'Other Reserves (Accumulated Other Comprehensive Income)'),
      ('303-1000','Reserve Of Share-Based Payments', 'Other Reserves (Accumulated Other Comprehensive Income)'),
      ('303-1100','Reserve Of Remeasurements Of Defined Benefit Plans', 'Other Reserves (Accumulated Other Comprehensive Income)'),
      ('303-1200','Amount Recognised In Other Comprehensive Income And Accumulated In Equity Relating To Non-Current Assets Or Disposal Groups Held For Sale', 'Other Reserves (Accumulated Other Comprehensive Income)'),
      ('303-1300','Reserve Of Gains And Losses From Investments In Equity Instruments', 'Other Reserves (Accumulated Other Comprehensive Income)'),
      ('303-1400','Reserve Of Change In Fair Value Of Financial Liability Attributable To Change In Credit Risk Of Liability', 'Other Reserves (Accumulated Other Comprehensive Income)'),
      ('303-1500','Reserve For Catastrophe', 'Other Reserves (Accumulated Other Comprehensive Income)'),
      ('303-1600','Reserve For Equalisation', 'Other Reserves (Accumulated Other Comprehensive Income)'),
      ('303-1700','Reserve Of Discretionary Participation Features', 'Other Reserves (Accumulated Other Comprehensive Income)'),
      ('303-1800','Reserve Of Equity Component Of Convertible Instruments', 'Other Reserves (Accumulated Other Comprehensive Income)'),
      ('303-1900','Capital Redemption Reserve', 'Other Reserves (Accumulated Other Comprehensive Income)'),
      ('303-2000','Merger Reserve', 'Other Reserves (Accumulated Other Comprehensive Income)'),
      ('303-2100','Statutory Reserve', 'Other Reserves (Accumulated Other Comprehensive Income)'),
      ('304','Other Equity', 'Equity'),
      ('304-0100','Other Equity Interest', 'Other Equity'),
      ('304-0200','Capital Reserve', 'Other Equity'),
      ('304-0300','Receivable For Shares (Share Subscribed But Unissued)', 'Other Equity'),
      ('304-0400','Treasury Shares', 'Other Equity'),
      ('304-0500','Controlling Interest', 'Other Equity'),
      ('304-0600','Non-Controlling Interest', 'Other Equity'),
      ('400','Revenue', null),
      ('401','Sales of Goods', 'Revenue'),
      ('401-0100','Products', 'Sales of Goods'),
      ('401-0200','Sales of Merchandise', 'Sales of Goods'),
      ('401-0300','Adjustments on Sales', 'Sales of Goods'),
      ('401-0400','Specifically Itemized Goods Revenues', 'Sales of Goods'),
      ('402','Revenue from Services', 'Revenue'),
      ('402-0100','Specifically Itemized Service Revenues', 'Revenue from Services'),
      ('403','Revenue From Contracts With Customers', 'Revenue'),
      ('404','Revenue From Construction Contracts', 'Revenue'),
      ('405','Franchise Fee Income', 'Revenue'),
      ('406','Barter Sales', 'Revenue'),
      ('406-0100','Barter Sales of Goods', 'Barter Sales'),
      ('406-0200','Barter Sales Services', 'Barter Sales'),
      ('406-0300','Barter Sales Construction Contracts', 'Barter Sales'),
      ('406-0400','Barter Sales of Royalties', 'Barter Sales'),
      ('406-0500','Barter Sales Interest', 'Barter Sales'),
      ('406-0600','Barter Sales Dividends', 'Barter Sales'),
      ('406-0700','Other Exchange Revenue', 'Barter Sales'),
      ('407','Financial Income', 'Revenue'),
      ('407-0100','Finance Income', 'Other Income'),
      ('407-0200','Interest Income', 'Other Income'),
      ('407-0300','Dividends Income', 'Other Income'),
      ('408','By Nature Non-Revenue Income', 'Revenue'),
      ('408-0100','Changes In Inventories', 'By Nature Non-Revenue Income'),
      ('408-0200','Work Performed By Entity And Capitalised', 'By Nature Non-Revenue Income'),
      ('409','Non-Operating Income  (Peripheral Activities)', 'Revenue'),
      ('409-9000','Other Income', 'Non-Operating Income  (Peripheral Activities)'),
      ('409-9001','Royalties Income', 'Other Income'),
      ('409-9002','Licensees', 'Other Income'),
      ('409-9003','Rental Income', 'Other Income'),
      ('409-9004','Contractual Fines And Penalties', 'Other Income'),
      ('409-9005','Income From Government Grants', 'Other Income'),
      ('409-9006','Property Service Charge Income', 'Other Income'),
      ('409-9007','Income From Reimbursements Under Insurance Policies', 'Other Income'),
      ('410','Income from Donation', 'Non-Operating Income  (Peripheral Activities)'),
      ('440','Other Comprehensive Income Reclassification Adjustments', 'Revenue'),
      ('440-0100','Group Companies', 'Other Comprehensive Income Reclassification Adjustments'),
      ('480-0101','Equity Method Investments', 'Group Companies'),
      ('480-0102','Subsidiaries, Jointly Controlled Entities And Associates', 'Group Companies'),
      ('480-0200','Exchange Differences On Translation', 'Other Comprehensive Income Reclassification Adjustments'),
      ('480-0300','Available-For-Sale Financial Assets Gain (Loss)', 'Other Comprehensive Income Reclassification Adjustments'),
      ('480-0400','Cash Flow Hedges', 'Other Comprehensive Income Reclassification Adjustments'),
      ('480-0500','Hedges Of Net Investment In Foreign Operations', 'Other Comprehensive Income Reclassification Adjustments'),
      ('480-0600','Change In Value Of Time Value Of Options', 'Other Comprehensive Income Reclassification Adjustments'),
      ('480-0700','Change In Value Of Forward Elements Of Forward Contracts', 'Other Comprehensive Income Reclassification Adjustments'),
      ('480-0800','Change In Value Of Foreign Currency Basis Spreads', 'Other Comprehensive Income Reclassification Adjustments'),
      ('480-0900','Financial Assets Measured At Fair Value Through Other Comprehensive Income', 'Other Comprehensive Income Reclassification Adjustments'),
      ('500','Expenses', null),
      ('501','Expenses Classified By Nature', 'Expenses'),
      ('501-0100','Material And Merchandise', 'Expenses Classified By Nature'),
      ('501-0200','Salaries and Benefits', 'Expenses Classified By Nature'),
      ('501-0201','Salaries and Wages', 'Salaries and Benefits'),
      ('501-0202','Employee Benefits', 'Salaries and Benefits'),
      ('501-020201','Staff Benefits', 'Staff Benefits'),
      ('501-020202','Officers Benefits', 'Staff Benefits'),
      ('501-020203','SSS - Bank''s Share ', 'Staff Benefits'),
      ('501-020204','Pag-ibig - Bank''s Share ', 'Staff Benefits'),
      ('501-020205','Contributions to Retirement/Provident Fund ', 'Staff Benefits'),
      ('501-020206','Medical Assistance', 'Staff Benefits'),
      ('501-0300','Services', 'Expenses Classified By Nature'),
      ('501-030001','IT Expenses', 'Services'),
      ('501-030002','Cost of IT Services', 'Services'),
      ('501-030003','Outsourcing Fee', 'Services'),
      ('501-0400','Depreciation, Amortization And Depletion', 'Expenses Classified By Nature'),
      ('501-040001','Depreciation', 'Rent, Depreciation, Amortization And Depletion'),
      ('501-040002','Amortization', 'Rent, Depreciation, Amortization And Depletion'),
      ('501-040003','Depletion', 'Rent, Depreciation, Amortization And Depletion'),
      ('501-0500','Repairs and Maintenance', 'Expenses Classified By Nature'),
      ('501-0600','Power, Light and Water', 'Expenses Classified By Nature'),
      ('501-0700','Communication and Postage', 'Expenses Classified By Nature'),
      ('501-0800','Telephone', 'Expenses Classified By Nature'),
      ('501-0900','Gasoline and Travel', 'Expenses Classified By Nature'),
      ('501-090001','Travel', 'Gasoline and Travel'),
      ('501-090002','Car Maintenance', 'Gasoline and Travel'),
      ('501-090003','Driver''s Honorarium', 'Gasoline and Travel'),
      ('501-090004','Fuel and Lubricants', 'Gasoline and Travel'),
      ('501-090005','Travelling Expense', 'Gasoline and Travel'),
      ('501-1000','Benevolence', 'Expenses Classified By Nature'),
      ('501-1100','Office Supplies', 'Expenses Classified By Nature'),
      ('501-1200','Representation', 'Expenses Classified By Nature'),
      ('501-1300','Interest Expense', 'Expenses Classified By Nature'),
      ('501-1400','Management and Other Professional Fees', 'Expenses Classified By Nature'),
      ('501-1500','Supervision and Examination', 'Expenses Classified By Nature'),
      ('501-1600','Seminars and Meetings', 'Expenses Classified By Nature'),
      ('501-1700','Program Monitoring Evaluation Expenses', 'Expenses Classified By Nature'),
      ('501-1800','Staff Training and Development', 'Expenses Classified By Nature'),
      ('501-1900','Banking Fees', 'Expenses Classified By Nature'),
      ('501-2000','Insurance', 'Expenses Classified By Nature'),
      ('501-2100','Rent', 'Expenses Classified By Nature'),
      ('501-210001','Office Rental', 'Rent'),
      ('501-210002','Staff House Rental', 'Rent'),
      ('501-210003','Other Rental', 'Rent'),
      ('501-2200','Stationery and Supplies Used', 'Expenses Classified By Nature'),
      ('501-2300','Service Fee', 'Expenses Classified By Nature'),
      ('501-2400','Representation and Entertainment', 'Expenses Classified By Nature'),
      ('501-2500','Donations and Charitable Contributions', 'Expenses Classified By Nature'),
      ('501-2600','Provision for impairment losses', 'Expenses Classified By Nature'),
      ('501-2700','Increase (Decrease) In Inventories Of Finished Goods And Work In Progress', 'Expenses Classified By Nature'),
      ('501-2800','Other Work Performed By Entity And Capitalized', 'Expenses Classified By Nature'),
      ('501-2900','Miscellaneous Expenses', 'Expenses Classified By Nature'),
      ('502','Expenses Classified By Function', 'Expenses'),
      ('502-0100','Cost Of Sales', 'Expenses Classified By Function'),
      ('502-0200','Selling, General And Administrative ', 'Expenses Classified By Function'),
      ('502-0300','Accounts Receivable, Credit Loss (Reversal)', 'Expenses Classified By Function'),
      ('503','Other (Non-Operating) Income And Expenses', 'Expenses'),
      ('504','Other Revenue And Expenses', 'Expenses'),
      ('504-0100','Other Revenue', 'Other Revenue And Expenses'),
      ('504-0200','Other Expenses', 'Other Revenue And Expenses'),
      ('505','Gains And Losses', 'Expenses'),
      ('505-0100','Foreign Currency Transaction Gain (Loss)', 'Gains And Losses'),
      ('505-0200','Gain (Loss) On Investments', 'Gains And Losses'),
      ('505-0300','Gain (Loss) On Derivatives', 'Gains And Losses'),
      ('505-0400','Gain (Loss) On Disposal Of Assets', 'Gains And Losses'),
      ('505-0500','Debt Related Gain (Loss)', 'Gains And Losses'),
      ('505-0600','Impairment Loss', 'Gains And Losses'),
      ('505-0700','Other Gains And (Losses)', 'Gains And Losses'),
      ('506','Taxes (Other Than Income And Payroll) And Fees', 'Expenses'),
      ('506-0100','Real Estate Taxes And Insurance', 'Taxes (Other Than Income And Payroll) And Fees'),
      ('506-0200','Highway (Road) Taxes And Tolls', 'Taxes (Other Than Income And Payroll) And Fees'),
      ('506-0300','Direct Tax And License Fees', 'Taxes (Other Than Income And Payroll) And Fees'),
      ('506-0400','Taxes and licenses', 'Taxes (Other Than Income And Payroll) And Fees'),
      ('506-0500','Documentary Stamps Used', 'Taxes (Other Than Income And Payroll) And Fees'),
      ('506-0600','Excise And Sales Taxes', 'Taxes (Other Than Income And Payroll) And Fees'),
      ('506-0700','Customs Fees And Duties (Not Classified As Sales Or Excise)', 'Taxes (Other Than Income And Payroll) And Fees'),
      ('506-0800','Non-Deductible VAT (GST)', 'Taxes (Other Than Income And Payroll) And Fees'),
      ('506-0900','General Insurance Expense', 'Taxes (Other Than Income And Payroll) And Fees'),
      ('506-1000','Administrative Fees (Revenue Stamps)', 'Taxes (Other Than Income And Payroll) And Fees'),
      ('506-1100','Penalties and Other Charges', 'Taxes (Other Than Income And Payroll) And Fees'),
      ('506-1200','Miscellaneous Taxes', 'Taxes (Other Than Income And Payroll) And Fees'),
      ('506-1300','Other Taxes And Fees', 'Taxes (Other Than Income And Payroll) And Fees'),
      ('700','Intercompany And Related Party Accounts', null),
      ('701','Intercompany And Related Party Assets', 'Intercompany And Related Party Accounts'),
      ('701-0100','Intercompany Balances in Asset', 'Intercompany And Related Party Assets'),
      ('701-0200','Related Party Balances in Asset', 'Intercompany And Related Party Assets'),
      ('701-0300','Intercompany Investments', 'Intercompany And Related Party Assets'),
      ('702','Intercompany And Related Party Liabilities', 'Intercompany And Related Party Accounts'),
      ('702-0100','Intercompany Balances in Liability', 'Intercompany And Related Party Liabilities'),
      ('702-0101','Related Party Balances in Liability', 'Intercompany And Related Party Liabilities'),
      ('703','Intercompany And Related Party Income And Expense', 'Intercompany And Related Party Accounts'),
      ('703-0100','Intercompany And Related Party Income', 'Intercompany And Related Party Income And Expense'),
      ('703-0200','Intercompany And Related Party Expenses', 'Intercompany And Related Party Income And Expense'),
      ('703-0300','Income (Loss) From Equity Method Investments', 'Intercompany And Related Party Income And Expense'),
      ('800','Temporary Acount', 'Temporary Acount')
      )   
  a(acc, title, parent)),

  
  a as (
    SELECT 
      coa.acc, coa.title, coa.parent, cast('' as text) ParentPath, 0 lvl
    FROM coa
    WHERE parent is null
    UNION ALL 
    SELECT 
      g.acc, g.title, g.parent, cast(a.ParentPath || g.parent || '>' as text) ParentPath, lvl+1 lvl
    FROM coa g
    INNER JOIN a on a.title = g.parent
    WHERE g.parent is not null)

    SELECT acc, title, parent FROM a order by lvl, acc
  ;
  
  i integer;
  _id integer;
BEGIN
   -- Open the cursor
   OPEN cur;
   
   LOOP

      FETCH cur INTO rec;
    -- exit when no more row to fetch
      EXIT WHEN NOT FOUND;
 
      IF rec.parent is not null THEN      
        SELECT id into _id FROM COA_Parent p WHERE p.Title = rec.parent;
        IF _id is null THEN 
          RAISE EXCEPTION 'Nonexistent parent --> %' , rec.parent USING HINT = 'Please check your parent';        
        END IF;
      ELSE
        _id = null;
      END IF;
    
      INSERT INTO COA_Parent(
        Acc,   Title,   Parent_ID, COA_Seq) 
      SELECT rec.Acc, rec.Title, _id, 0
      ON CONFLICT(Acc) DO UPDATE SET 
        Title = excluded.Title,
        Parent_ID = excluded.Parent_ID;
      
   END LOOP;
  
   -- Close the cursor
   CLOSE cur;

END; $$ ;
