
IF NOT EXISTS (SELECT Name FROM sysObjects where Name = 'trnDates')
BEGIN
  CREATE TABLE trnDates
    (TrnDate SmallDateTime NOT NULL,
	 lnTrn Numeric(18,0)   NOT NULL,
	 saTrn Numeric(18,0)   NOT NULL,
	 CONSTRAINT pk_trnDates PRIMARY KEY CLUSTERED (TrnDate))  
END
GO

MERGE trnDates T 
USING (SELECT trnDate, Max(trn) trn FROM trnMaster GROUP BY trnDate) S on S.trnDate = T.trnDate
WHEN NOT MATCHED THEN
  INSERT (trnDate, lnTrn, saTrn) VALUES(S.trnDate, S.trn, 0)
WHEN MATCHED THEN
  UPDATE SET lnTrn = S.trn;

MERGE trnDates T 
USING (SELECT trnDate, Max(trn) trn FROM saTrnMaster GROUP BY trnDate) S on S.trnDate = T.trnDate
WHEN NOT MATCHED THEN
  INSERT (trnDate, lnTrn, saTrn) VALUES(S.trnDate, 0, S.trn)
WHEN MATCHED THEN
  UPDATE SET saTrn = S.trn;

GO

ALTER TRIGGER [dbo].[Insert_trnMaster] ON [dbo].[TRNMASTER]
  INSTEAD OF INSERT
AS 
BEGIN
  MERGE trnDates T 
  USING (SELECT trnDate, Count(*) trn FROM Inserted GROUP BY trnDate) S on S.trnDate = T.trnDate
  WHEN NOT MATCHED THEN
    INSERT (trnDate, lnTrn, saTrn) VALUES(S.trnDate, S.trn, 0)
  WHEN MATCHED THEN
    UPDATE SET lnTrn = lnTrn + S.trn;

  INSERT trnMaster 
  SELECT 
    ACC, i.trnDate, 
	ROW_NUMBER() OVER (PARTITION BY i.trnDate ORDER BY Acc) + d.lnTrn, OrNo, 
    Prin, IntR, Oth, Penalty, TrnAmt, WaivedInt, TrnType, 
    UserName, TermID, Balance, 
    RefNo, TrnDesc, TrnMnem_CD, Particulars, 
    [Time], Cancel 
  FROM Inserted i 
  INNER JOIN trnDates d on d.trnDate = i.trnDate

END
GO

ALTER PROCEDURE [dbo].[usp_Updlntran](
  --@Acc         as Char(17) =  '0000-4065-0023547',  
  --@Acc         as Char(17)    =  '0000-4048-0000021',  
 -- @Acc         as Char(17)    =  '0101-4240-0003757',  
  @Acc         as Char(17),
  @TrnType     as SmallInt    = 3001,      
  @mTrnAmt     as Numeric(14,2),  
  @OrNo        as Numeric     = 0,    
  @PostBy      as VarChar(15) = 'sa',     
  @TermID      as VarChar(15) = '',      
  @Particulars as VarChar(100)= '',       
  @RefNo       as VarChar(16) = '',      
  @lnStatus    as Int         = 0,      
  @TrnDate     as DateTime    = 0)
AS
  DECLARE    
    @CID       as Int,       
    @dNum      as Int,   
    @TrnAmt    as Numeric(14,2),      
	-- Current Balance
    @BalPrin   as Numeric(14,2), 
    @BalInt    as Numeric(14,2),      
    @BalOth    as Numeric(14,2),
    @PaidWaived as Numeric(14,2),
	@Bal        as Numeric(14,2),
	
	-- Balance should be
    @ShouldBalPrin   as Numeric(14,2),       
    @ShouldBalInt    as Numeric(14,2),    
    @ShouldBalOth    as Numeric(14,2),
	
	-- Payment Distribution
    @Prin      as Numeric(14,2),    
    @Int       as Numeric(14,2),    
    @Oth       as Numeric(14,2),
    @WaivedInt as Numeric(14,2) = 0,
	@Penalty as Numeric(14,2) = 0, -- Not in Used
    @TrnDesc    as VarChar(100),	
	
	-- Info from lnMaster
    @Frequency   as Int, 
	@AcctType    as Numeric,

    @RefDate as DateTime,       

	-- Info from ebSysDate
	@SysDate as DateTime,
    @WaivableInt as Bit,
    @RunState    as Int,

	@MnemCode   as Int

  SELECT @MnemCode = isNull(t,43)
  FROM 
   (VALUES 
    (3001, 43),                           
    (3097, 56),                           
    (3098, 54),          
    (3099, 48),        
    (3899, 0 ),       
    (3201, 43),       
    (3202, 43)) m(t,m)
  WHERE t = @TrnType

  SELECT 
    @CID = CID,
    @Frequency = Frequency,
	@AcctType  = AcctType
  FROM lnMaster 
  WHERE Acc = @Acc

  SELECT 
    @SysDate = ebSysDate, 
    @RunState = RunState, 
    @WaivableInt = CASE WHEN @AcctType in (315) THEN 0 ELSE IsNull(WaivableInt,1) END 
  FROM orgParms  

  IF @RunState <> 0       
  BEGIN       
    RAISERROR ('Cannot Transaction. System is not open yet...', 16, 1)        
    RETURN       
  END    

  -- Force trnDate to SystemDate
  SET @TrnDate = @SysDate 

  SET @RefDate = dbo.RefDueDate(@Frequency,@SysDate,0)       

  SELECT 
    @BalPrin = m.Principal - isNull(t.Prin,0),
    @BalInt  = m.Interest - isNull(t.IntR,0),
    @BalOth  = m.Others - Discounted - isNull(t.Oth,0),
    @PaidWaived = isNull(WaivedInt,0),
	@ShouldBalPrin = isNull(i.Prin,0),
	@ShouldBalInt = isNull(i.IntR,0),
	@ShouldBalOth = isNull(i.Oth,0)
  FROM lnMaster m
  ,(SELECT 
      Sum(Prin) Prin, Sum(IntR) IntR, Sum(WaivedInt) WaitedInt, Sum(Oth) Oth
    FROM trnMaster
    WHERE Acc = @Acc and trnType in (3001,3097,3098,3099,3899,3201,3202) 
   ) t
  ,(SELECT 
      Sum(Prin) Prin, Sum(IntR) IntR, Sum(Oth) Oth
    FROM LoanInst i
    WHERE Acc = @Acc and DueDate > @RefDate
      ) i
  WHERE m.Acc = @Acc
  
  SET @Bal = @BalPrin+@BalInt+@BalOth

  -- If Payment is enough to fullpaid the loan apply all waivedint
  IF @Bal-@ShouldBalInt-@ShouldBalOth-@mTrnAmt <= 1 and @WaivableInt = 1
      SET @WaivedInt = @ShouldBalInt+@ShouldBalOth

  -- Do this if there is already waived interest applied in the past which makes the balance less than 25
  IF @BalPrin+@BalInt+@BalOth-@PaidWaived-@mTrnAmt < 25 and @ShouldBalInt+@ShouldBalOth < @PaidWaived
     SET  @WaivedInt = @PaidWaived

  IF  @WaivedInt > @ShouldBalInt + @ShouldBalOth
    SET @WaivedInt = @ShouldBalInt + @ShouldBalOth

  SET @Oth = @BalOth - @ShouldBalOth
  SET @Int = @BalInt - @ShouldBalInt

  SET @TrnAmt = @mTrnAmt

  IF @TrnAmt > 0
  BEGIN
	IF @Oth > @TrnAmt
      SET @Oth = @TrnAmt  
	SET @TrnAmt = @TrnAmt - @Oth 

	IF @Int > @TrnAmt
	  SET @Int = @TrnAmt
	SET @TrnAmt = @TrnAmt - @Int
  
	SET @Prin = @TrnAmt
	IF @BalPrin < @TrnAmt
	  SET @Prin = @BalPrin
	SET @TrnAmt = @TrnAmt - @Prin 

  END ELSE
  BEGIN
	IF @Oth < @TrnAmt 
      SET @Oth = @TrnAmt  
	SET @TrnAmt = @TrnAmt - @Oth 

	IF @Int < @TrnAmt
	  SET @Int = @TrnAmt
	SET @TrnAmt = @TrnAmt - @Int
  
	SET @Prin = @TrnAmt
	IF @BalPrin < @TrnAmt
	  SET @Prin = @BalPrin
	SET @TrnAmt = @TrnAmt - @Prin 
  END

  IF @TrnAmt < 0 
  BEGIN
     SET @Prin = @Prin + @TrnAmt
     SET @TrnAmt = 0    
  END

  SET @WaivedInt = @WaivedInt - @PaidWaived 

  SET @Bal = @Bal - @Prin - @Int - @Oth - @WaivedInt 

  SELECT 
    @RefDate RefDate, @Prin Prin, @Int Intr, @Oth Oth, @TrnAmt TrnAmt, @mTrnAmt mTrnAmt, @WaivedInt WaivedInt, @PaidWaived PaidWaived, 
    @ShouldBalPrin ShouldBalPrin, @ShouldBalInt ShouldBalInt, @BalPrin BalPrin, @BalInt BalInt, @Bal Bal

  -- Exit if there is no transaction
  IF abs(@Prin)+abs(@Int)+abs(@Oth)+abs(@WaivedInt) = 0 RETURN

  SELECT @TrnDesc = TrnDesc from trnTypes where TrnType = @TrnType

  DECLARE @TrnTbl TABLE (Trn Numeric(18,0))

  INSERT trnMaster (
    ACC, trnDate, TrnType, OrNo, TrnAmt, Prin,        
    IntR, Oth, Penalty,       
    WaivedInt, Balance, UserName, TermID,        
    RefNo, TrnDesc, TrnMnem_CD, Particulars,        
    [Time], Cancel)                 
  OUTPUT Inserted.trn INTO @TrnTbl(trn)
  VALUES(
    @Acc, @TrnDate, @TrnType, @OrNo, @TrnAmt, @Prin, @Int, @Oth, @Penalty,       
    @WaivedInt, @Bal, @PostBy, @TermID,        
    @RefNo, @TrnDesc, @MnemCode, @Particulars,       
    GetDate(), 0)
	
  UPDATE LoanInst SET 
    InstPD   = CASE WHEN EndBal+EndInt+EndOth >= @Bal THEN Prin+IntR+Oth 
	                WHEN EndBal+EndInt+EndOth+Prin+IntR+Oth < @Bal THEN 0
					ELSE EndBal+EndInt+EndOth+Prin+IntR+Oth - @Bal END,
    InstFlag = CASE WHEN EndBal+EndInt+EndOth >= @Bal THEN 9 ELSE 0 END       
  WHERE Acc = @Acc     
  DECLARE @trn Numeric(18,0)
  SELECT @trn = trn FROM @TrnTbl

  SELECT @dNum = isNull(Max(dNum),0) FROM LoanInst WHERE @Acc = Acc and InstFlag = 9

  UPDATE lnMaster SET 
    Prin       = Principal-@BalPrin+@Prin,
    IntR       = Interest-@BalInt+@Int+Discounted,       
    Oth        = Others-@BalOth+@Oth,
    WaivedInt  = @PaidWaived+@WaivedInt,         
    doLastTrn  = @TrnDate,       
    LastTrn    = IsNull(@Trn,LastTrn),       
    WeeksPaid  = @dNum,     
    Status     = CASE WHEN @Bal = 0 THEN CASE WHEN @TrnType = 3899 THEN 98 ELSE 99  END        
                 ELSE CASE WHEN Status in (30,91) THEN Status ELSE 30 END END                        
  WHERE Acc    = @Acc  


  DECLARE @UpSaf as Bit       
  SET @UpSaf = 0       

-- Update Tellers Cash       
  IF @TrnType in (3001,3097,3098,3099)         
  BEGIN       
    UPDATE SAF SET Cash_On_Hand = Cash_On_Hand + @TrnAmt        
    WHERE TlrName = @PostBy       
    SET @UpSaf = 1       
  END -- End Update Tellers Cash 

  IF @TrnAmt > 0       
  BEGIN       
    DECLARE @sAcc as VarChar(17)       
	--(addded by kent)excess payment for closed account will be posted on AP Account
    SELECT @sAcc = Acc from saMaster       
    WHERE CID = @CID and AcctType = 40 and Status in (10,20,90)  

    IF @sAcc is Null       
      SET @sAcc = 'P'+rTrim(Convert(VarChar(10),@CID)) + '-40'       
      IF @TrnType  = 3899 -- From Renewals       
      BEGIN       
        SET @Particulars = 'Excess ' + @Acc       
 --print 'Insert savings 1'   
        EXEC usp_UpdSaTran @sAcc, 3, 903, @TrnAmt, 0, @ORNo, @PostBy, @TermID, @Particulars, 0, 'A', @UpSaf,1         
        END 
      ELSE BEGIN        
        IF not @TrnType  in (3201,3202)       
        BEGIN       
--print 'Insert savings 2'   
          EXEC usp_UpdSaTran @sAcc, 3, 7012, @TrnAmt, 0, @ORNo, @PostBy, @TermID, @Particulars, 0, 'A', @UpSaf,1       
        END ELSE       
--print 'Insert savings 3'   
          EXEC usp_UpdSaTran @sAcc, 231, 15, @TrnAmt, 0, @ORNo, @PostBy, @TermID, @Particulars, 0, 'A', @UpSaf,1       
        END       
      END      

  select * from satrnmaster where trndate = 

  SELECT * from loaninst where acc = @Acc

  SELECT * from trnmaster where acc = @acc and trndate = @trndate
  SELECT * FROM lnMaster WHERE Acc = @Acc

  IF @Bal = 0       
    UPDATE lnMaster SET Status = 99 WHERE Acc = @Acc and Status not in (99, 98)       
 
--  IF @Bal > 0 and @Particulars = 'Full Payment' and @TrnType = 3099  
--    exec fixloantran @ACC,0,0 
--- Fixed special Loan   
  IF @AcctType = 330 exec usp_FixSpecialLoan @Acc
GO

BEGIN TRAN
 EXEC usp_Updlntran
  @Acc =  '0000-4048-0000021',  
  @TrnType    = 3001,      
  @mTrnAmt     =  19449.49,  
  @OrNo          = 0,    
  @PostBy      = 'sa',     
  @TermID       = '',      
  @Particulars = '',       
  @RefNo       = '',      
  @lnStatus      = 0,      
  @TrnDate   = '2021-3-17'


ROLLBACK