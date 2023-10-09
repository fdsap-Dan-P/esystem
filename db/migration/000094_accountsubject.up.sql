
---------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Account_Subject (
---------------------------------------------------------------------------
  ID BIGSERIAL NOT NULL,
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Account_ID bigint NOT NULL,
  Section_Subject_ID bigint NOT NULL,
  Subject_ID bigint NOT NULL,
  Ratings_1st_Qtr Numeric(8,5) NOT NULL DEFAULT 0,
  Ratings_2nd_Qtr Numeric(8,5) NOT NULL DEFAULT 0,
  Ratings_3rd_Qtr Numeric(8,5) NOT NULL DEFAULT 0,
  Ratings_4th_Qtr Numeric(8,5) NOT NULL DEFAULT 0,
  Ratings_Final  Numeric(8,5) NOT NULL DEFAULT 0,
  Attendance_Ctr Int NOT NULL DEFAULT 0,
  Absent_Ctr Int NOT NULL DEFAULT 0,
  Late_Ctr Int NOT NULL DEFAULT 0,
  Status_ID bigint NOT NULL,
  Remarks Varchar(200) NOT NULL DEFAULT '',
  Other_Info jsonb NULL,
 
  CONSTRAINT Account_Subject_pkey PRIMARY KEY (ID),
  CONSTRAINT fk_Account_Subject_Acc FOREIGN KEY (Account_ID) REFERENCES Account(ID),
  CONSTRAINT fk_Account_Subject_SubjSection FOREIGN KEY (Section_Subject_ID) REFERENCES Section_Subject(ID),
  CONSTRAINT fk_Account_Subject_Subject FOREIGN KEY (Subject_ID) REFERENCES Subject(ID),
  CONSTRAINT fk_Account_Subject_Status FOREIGN KEY (Status_ID) REFERENCES Reference(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccount_Subject_UUID ON public.Account_Subject(UUID);
CREATE UNIQUE INDEX IF NOT EXISTS idxAccount_Subject_Subject ON public.Account_Subject(Account_ID, Section_Subject_ID);

DROP TRIGGER IF EXISTS trgAccount_Subject_Ins on Account_Subject;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Subject_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Account_Subject
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
 
DROP TRIGGER IF EXISTS trgAccount_Subject_upd on Account_Subject;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Subject_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Account_Subject
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgAccount_Subject_del on Account_Subject;
---------------------------------------------------------------------------
CREATE TRIGGER trgAccount_Subject_del
---------------------------------------------------------------------------
    AFTER DELETE ON Account_Subject
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();
  
