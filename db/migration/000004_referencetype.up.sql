-- Add Reference Type Table
---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Reference_Type (
---------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Code        varchar(15) NULL,
  Title       varchar(35) NOT NULL,
  Description varchar(1000) NULL,  
  Other_Info jsonb NULL,
  CONSTRAINT Reference_Type_pkey PRIMARY KEY (ID)
);

CREATE UNIQUE INDEX IF NOT EXISTS idxReference_Type_Title ON public.Reference_Type(LOWER(Title));
CREATE UNIQUE INDEX IF NOT EXISTS idxReference_Type_UUID ON public.Reference_Type(UUID);

-- Add Reference Trigger
-- DROP TRIGGER Reference_UPDATE ON Reference_Type;
DROP TRIGGER IF EXISTS trgReference_Type_Ins on Reference_Type;
---------------------------------------------------------------------------
CREATE TRIGGER trgReference_Type_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Reference_Type
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
  
DROP TRIGGER IF EXISTS trgReference_Type_upd on Reference_Type;
---------------------------------------------------------------------------
CREATE TRIGGER trgReference_Type_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Reference_Type
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();
    
DROP TRIGGER IF EXISTS trgReference_Type_del on Reference_Type;
---------------------------------------------------------------------------
CREATE TRIGGER trgReference_Type_del
---------------------------------------------------------------------------
    AFTER DELETE ON Reference_Type
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();

-- Populate Reference_Type  

INSERT INTO Reference_Type(Code, Title, Description)
ValueS('','NA','NA')
ON CONFLICT(LOWER(Title))
DO NOTHING;

---------------------------------------------------------------------------
INSERT INTO Reference_Type(Code, Title, Description)
---------------------------------------------------------------------------
  SELECT '' Code, a.Title, a.Title Description 
  FROM (Values
      ('Parameter'),
      ('ParameterDetail'),
      ('SubParameter'),
      ('DepositoryBank'),
      ('Application'),
      ('Account Attribute'),
      ('FundSource'),
      ('AccessObjectType'),
      ('Title'),
      ('TypeOfProvider'),
      ('ValidID'),
      ('RefAddress'),
      ('CustomerRestriction'),
      ('ReferenceStatus'),
      ('CustomerStatus'),
      ('CustomerRiskClass'),
      ('CustomerEventType'),
      ('AccountStatus'),
      ('Gender'),
      ('LoanClass'),
      ('SavingsClass'),
      ('Currency'),
      ('CivilStatus'),
      ('lnBeneChangeLevel'),
      ('Lending_Type'),
      ('lnBorrType'),
      ('LNECOActivity'),
      ('ln_Category'),
      ('LNGroup'),
      ('LoanTypes'),
      ('Ownership'),
      ('RelationshipType'),
      ('AccountTypeGroup'),
      ('AccountClass'),
      ('CustomerGroupType'),
      ('trnItemList'),
      ('SSSParameters'),
      ('SSSPaymentChannel'),
      ('SSSPaymentType'),
      ('TrnHeadType'),
      ('ProviderType'),
      ('IDGroupType'),
      ('IDType'),
      ('SSSPayorType'),
      ('Position'),
      ('ActionList'),
      ('ActorGroup'),
      ('LocationType'),
      ('AddressType'),
      ('EducationalLevel'),
      ('ContactType'),
      ('OfficeType'),
      ('TrnItem'),
      ('TicketType'),
      ('TicketStatus'),
      ('OfficeAccountType'),
      ('OfficeAccountTypeGroup'),
      ('Disabilities'),
      ('Occupation'),
      ('Religion'),
      ('Nationality'),
      ('UserStatus'),
      ('SourceofIncome'),
      ('Industry'),
      ('EmployeeType'),
      ('EmployeeStatus'),
      ('EmployeeLevel'),
      ('CustomerClass'),
      ('CustomerSubClass'),
      ('BeneficiaryType'),
      ('ConversionType'),
      ('MeasureType'),
      ('UnitMeasure'),
      ('AccessConfig'),
      ('AccessObject'),
      ('ScheduleType'),
      ('DonationClass'),
      ('BrandName'),
      ('GenericName'),
      ('TrnType'),
      ('OfficeAccountStatus'),
      ('CourseType'),
      ('CourseStatus'),
      ('SubjectStatus'),
      ('SchoolSectionStatus'),
      ('SubjectType'),
      ('SchoolTerm'),
      ('SchoolSemister'),
      ('ChurchFundItem'),
      ('EmployeeEvent'),
      ('EmployeeEventStatus'),
      ('Courses'),
      ('Subject'),
      ('DocumentType'),
      ('TrnHeadAction'),
      ('InventorySpecs'),
      ('Milestone'),
      ('QuestionaireType'),
      ('QuestionType'),
      ('QuestionStatus'),
      ('SubjectEvent'),
      ('SchoolGroup'),
      ('ActionListStatus'),
      ('LogicalCondition'),
      ('BSP Category'),
      ('MiDAS Category'),
      ('CustomerGeneralCategory'),
      ('CustomerSpecificSubCategory'),
      ('CustomerSpecificCategory'),
      ('CustomerSpecs'),
      ('OfficerStatus'),
      ('StoreInventoryItem')
   ) a(Title)
 ON CONFLICT (LOWER(Title)) 
 -- DO NOTHING
 DO UPDATE 
   SET Code = EXCLUDED.Code, Title = EXCLUDED.Title, Description = EXCLUDED.Description;

