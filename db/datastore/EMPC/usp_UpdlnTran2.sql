
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

ALTER TRIGGER [dbo].[Insert_satrnMaster] ON [dbo].[saTrnMaster]
  INSTEAD OF INSERT
AS 
BEGIN
  MERGE trnDates T 
  USING (SELECT trnDate, Count(*) trn FROM Inserted GROUP BY trnDate) S on S.trnDate = T.trnDate
  WHEN NOT MATCHED THEN
    INSERT (trnDate, lnTrn, saTrn) VALUES(S.trnDate, 0, S.trn)
  WHEN MATCHED THEN
    UPDATE SET saTrn = saTrn + S.trn;

  INSERT satrnMaster (
    Acc, TrnDate, Trn, 
    Orno, [Time], Particulars, trnAmt,
    Balance, AvlBal, AccrdInt, trnType, trnMnem_cd, UserName,
    TermID, pbkPosted, PendApprove, RefDate, RefNo)
  SELECT 
    Acc, i.TrnDate,
	ROW_NUMBER() OVER (PARTITION BY i.trnDate ORDER BY Acc) + d.saTrn Trn, 
    Orno, [Time], Particulars, trnAmt,
    Balance, AvlBal, AccrdInt, trnType, trnMnem_cd, UserName,
    TermID, pbkPosted, PendApprove, RefDate, RefNo
  FROM Inserted i 
  INNER JOIN trnDates d on d.trnDate = i.trnDate

END
GO

ALTER PROCEDURE [dbo].[usp_Updsatran](
  @Acc         AS Char(17),         
  @TrnType     AS SmallINT    = 1,        
  @MnemCode    AS SmallINT    = 1,        
  @TrnAmt      AS NUMERIC(14,2) = 0,        
  @ChkAmt      AS NUMERIC(14,2) = 0,        
  @ORno        AS NUMERIC     = 0,        
  @PostBY      AS VARCHAR(15) = 'sa',        
  @TermID      AS VARCHAR(15) = '',        
  @Particulars AS VARCHAR(100)= '',        
  @PbkPosted   AS SmallINT    = 0,        
  @PENDApprove AS VARCHAR(1)  = 'A',        
  @UpdateSAF   AS Bit         = 1,        
  @UpdateCust  AS Bit         = 0,        
  @RefNo       AS VARCHAR(16) = 0,        
  @ReverSETrn  AS VARCHAR(16) = 0,        
  @CardTerminal AS VARCHAR(16) = '')        
AS         
SET NOCOUNT ON      

BEGIN        
  DECLARE 
    @CID              AS INT,        
    @AcctType         AS INT,        
    @TrnDate          AS DATETIME,         
    @Trn              AS INT,        
    @RefDate          AS DATETIME,        
    @Bal              AS NUMERIC(14,2),
    @MembershipFee    AS NUMERIC(8,2) = 0,  
	@WithExisting     AS Bit = 1,

	-- Info from ebSysDate
	@SysDate as DateTime,
    @RunState    as Int

  IF @TrnAmt = 0 RETURN

  SELECT 
    @SysDate = ebSysDate, 
    @RunState = RunState
  FROM orgParms  

  SET @TrnDate = @SysDate

  IF @RunState <> 0       
  BEGIN       
    RAISERROR ('Cannot Transaction. System is not open yet...', 16, 1)        
    RETURN       
  END         

  IF LEFT(@Acc,1) = 'P'        
  BEGIN        
    SET @CID = SubString(@Acc,2,CHARINDEX('-',@Acc)-2)        
    SET @AcctType = SubString(@Acc,CHARINDEX('-',@Acc)+1,2)        

	IF NOT EXISTS (SELECT Acc FROM saMaster WHERE Status in (10,20,90) and AcctType = 10) 
	  SET @WithExisting = 0     

    EXEC usp_NewSavAcct @CID, @AcctType, 1        
      
    SELECT @Acc = Acc      
    FROM SaMaster     
    WHERE CID = @CID
      AND AcctType = @AcctType AND Status in (10,20,90)	
  END    

  SELECT @Acc = Acc, @Bal = Balance      
  FROM SaMaster     
  WHERE Acc = @Acc

  SET @Bal = @Bal + CASE WHEN @TrnType % 2 = 1 THEN @TrnAmt ELSE -@TrnAmt END
  
  INSERT saTrnMaster (
    Acc, TrnDate, Trn, Orno, Time, Particulars, trnAmt,          
    Balance, AvlBal, AccrdINT, trnType, trnMnem_cd, UserName, TermID,          
    pbkPosted, PENDApprove, RefDate, RefNo)          
  VALUES(
    @Acc, @TrnDate, @Trn, @ORno, GetDate(), @Particulars, @TrnAmt,           
    @Bal, @Bal, 0, @TrnType, @MnemCode, @PostBY, @TermID,           
    @PbkPosted, @PENDApprove, NULL, @RefNo)   

  IF @WithExisting = 0 
  BEGIN
	SET @MembershipFee = (SELECT Membershipfee FROM Settings) 
	IF @TrnAmt < @MembershipFee
	   SET @MembershipFee = @TrnAmt  
	  --SET @TrnAmt = @TrnAmt - @MembershipFee

	SET @Bal = @Bal - @MembershipFee

    INSERT saTrnMaster (
      Acc, TrnDate, Trn, Orno, Time, Particulars, trnAmt,          
      Balance, AvlBal, AccrdINT, trnType, trnMnem_cd, UserName, TermID,          
      pbkPosted, PENDApprove, RefDate, RefNo)          
    VALUES(
	  @Acc, @TrnDate, @Trn, @ORno, GetDate(), @Particulars, @MembershipFee,           
      @Bal, @Bal, 0, 4, 0, @PostBY, @TermID,           
      @PbkPosted, @PENDApprove, NULL, @RefNo)

    UPDATE saMaster         
    SET -- unCldBal       = @Bal - @AvlBal,        
      Balance        = @Bal,        
      AvlBal         = @Bal,        
      PBBal          = @Bal,        
      LastTrnType    = @TrnType,        
      LastTrnAmount  = @TrnAmt,        
      DoLAStTrn      = @TrnDate,        
      Status         = CASE WHEN @Bal = 0 AND @TrnType in (506,2506,1008) THEN 99 ELSE Status END        
    WHERE Acc = @Acc
  END
END
GO

ALTER PROCEDURE FixLoanInst(@Acc as VarChar(17))  
AS  
  
  DECLARE 
    @BalPrin as Numeric(14,2),
    @BalInt  as Numeric(14,2),
    @BalOth  as Numeric(14,2),
    @CarVal  as Numeric(14,2),
    @Amort as Numeric(14,2),  
    @Prin  as Numeric(14,2),  
    @Int   as Numeric(14,2),  
    @Oth   as Numeric(14,2),  
    @pDate as DateTime,  
    @DueDate as DateTime,  
    @CutDate as DateTime = '2021-05-01',  
    @SysDate as DateTime,
    @CumInt as Numeric(14,2) = 0,  
    @Int2 as Numeric(14,2),
    @sAcc as VarChar(17),  

    @Days as Int,
    @dNum as Int,  
    @Term as Int,
    @i as Int,
    @ReCompute as Bit = 0,
    @IntRate Float

  DECLARE @LN TABLE (      
    [dNum] [SmallInt] NOT NULL,
    [DueDate] [SmallDateTime] NOT NULL,
    [Prin] numeric (16, 2) NULL,
    [IntR] numeric (16, 2) NULL,
    [Oth]  numeric (16, 2) NULL
  )  

  SELECT @SysDate = ebSysDate FROM OrgParms

  SELECT 
    @BalPrin = Principal,  
    @BalInt = Interest-Discounted,  
    @BalOth = Others,
    @pDate   = DisbDate  
  FROM lnMaster WHERE Acc = @Acc  

  SELECT @Term = Max(dNum) FROM LoanInst WHERE Acc = @Acc
  SET @dNum = 1
  SET @CarVal = @BalPrin

  WHILE @dNum <= @Term
  BEGIN
    SELECT
      @DueDate = DueDate,
      @Prin = isNull(Prin,0),
      @Int  = isNull(IntR,0),
      @Oth  = isNull(Oth,0)
    FROM LoanInst
    WHERE Acc = @Acc and dNum = @dNum

    IF @DueDate >= @CutDate and @SysDate < '2021-09-01'
    BEGIN      
      SET @IntRate = .08   
      SET @ReCompute = 1   
    END ELSE SET @IntRate = 0

    SET @Amort = @Prin + @Int
	SET @CarVal = @BalPrin

    IF @ReCompute = 1
    BEGIN
      SET @Days = DateDiff(dd, @pDate, @DueDate)            
      IF @DueDate > @CutDate and @pDate < @CutDate -- Change of Interest from 10% to 8%      
      BEGIN
        SET @Days = DateDiff(dd, @pDate, @CutDate)-1
        SET @Int2 = Round(.14 * @CarVal / 365 * @Days,2)
        SET @Days = DateDiff(dd, @CutDate, @DueDate) + 1
		select @DueDate, @Days, @CarVal, @Int2, Round(.08 * @CarVal / 365 * @Days,2)
        SET @Int2 = @Int2 + Round(.08 * @CarVal / 365 * @Days,2)
      END ELSE SET @Int2 = Round(@IntRate * @CarVal / 365 * @Days,2)     
      SET @CumInt = @CumInt + @Int2

      IF @CumInt <= @Amort
      BEGIN
        SET @Int = @CumInt
        SET @Prin = @Amort - @Int
      END ELSE
      BEGIN
        SET @Prin = 0
        SET @Int = @Amort
      END  
      SET @CumInt = @CumInt - @Int
    END ELSE
    BEGIN
      IF @Int > @BalInt
        SET @Int = @BalInt
    END

    SET @BalInt = @BalInt - @Int
    IF @BalPrin < @Prin OR @dNum = @Term
      SET @Prin = @BalPrin
	SET @BalPrin = @BalPrin -@Prin

    IF @BalOth < @Oth OR @dNum = @Term
      SET @Oth = @BalOth
    SET @BalOth = @BalOth - @Oth

    INSERT @LN(dNum, DueDate, Prin, IntR, Oth)
    VALUES(@dNum, @DueDate, @Prin, @Int, @Oth)

    SET @pDate = @DueDate  
    SET @dNum = @dNum + 1
  END

  UPDATE lnMaster SET
    Interest = (SELECT Sum(IntR) FROM @LN)
  WHERE @Acc = Acc

  UPDATE LoanInst SET
    DueDate = l.DueDate,
    Prin = l.Prin,
    Intr = l.Intr,
    Oth = l.Oth,
    EndBal = l.EndBal,
    EndInt = l.EndInt,
    EndOth = l.EndOth
  FROM 
   (SELECT 
      l.dNum, l.DueDate, l.Prin, l.IntR, l.Oth,
      isNull(Sum(i.Prin),0) EndBal,
      isNull(Sum(i.IntR),0) EndInt,
      isNull(Sum(i.Oth),0)  EndOth
    FROM @LN l
    LEFT JOIN @LN i ON l.dNum < i.dNum
	GROUP BY l.dNum, l.DueDate, l.Prin, l.IntR, l.Oth
    ) l
  WHERE LoanInst.Acc = @Acc and LoanInst.dNum = l.dNum
GO

ALTER PROCEDURE FixLoanTran(
  @Acc as VarChar(17), @nAmt as Numeric(14,2),
  @bAdd as bit = 0, @TrnDate as DateTime = 0)
AS

DECLARE
  @TrnType     as SmallInt    = 3001,      
  @mTrnAmt     as Numeric(14,2),  
  @OrNo        as Numeric     = 0,    
  @PostBy      as VarChar(15) = 'sa',     
  @TermID      as VarChar(15) = '',      
  @Particulars as VarChar(100)= '',       
  @RefNo       as VarChar(16) = '',      
  @lnStatus    as Int         = 0

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

  IF @bAdd = 0 and @nAmt <> 0       
  BEGIN       
    RAISERROR ('Unable to force account balance other then zero', 16, 1)        
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
  
  -- Check if Add Balance or Force to adjust balance
  If @bAdd = 1
  BEGIN
    SET @mTrnAmt = @nAmt
  END ELSE
    SET @mTrnAmt = @Bal - @nAmt
 
  IF @Bal = 0 and @bAdd = 0 -- Force to Close Account
  BEGIN
      SET @ShouldBalPrin = 0
      SET @ShouldBalOth = @BalOth
      SET @ShouldBalInt = @BalInt
      SET @Prin = @BalPrin
      SET @Int = -@BalPrin
      SET @Oth = 0
      SET @WaivableInt = @BalPrin+@BalInt+@BalOth
  END ELSE
  BEGIN -- Add = 1 (Normal Transaction)
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

    IF @TrnAmt >= 0
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
  END -- for Add = 1

  SET @WaivedInt = @WaivedInt - @PaidWaived 

  SET @Bal = @Bal - @Prin - @Int - @Oth - @WaivedInt 

  SELECT 
    @TrnDate TrnDate, @RefDate RefDate, @Prin Prin, @Int Intr, @Oth Oth, @TrnAmt TrnAmt, @mTrnAmt mTrnAmt, @WaivedInt WaivedInt, @PaidWaived PaidWaived, 
    @ShouldBalPrin ShouldBalPrin, @ShouldBalInt ShouldBalInt, @BalPrin BalPrin, @BalInt BalInt, @Bal Bal

  -- Exit if there is no transaction
  IF abs(@Prin)+abs(@Int)+abs(@Oth)+abs(@WaivedInt) = 0 RETURN

  SELECT @TrnDesc = TrnDesc from trnTypes where TrnType = @TrnType

  DECLARE @TrnTbl TABLE (Trn Numeric(18,0))

  SET @TrnType = CASE WHEN @TrnAmt > 0 THEN 3001 ELSE 3098 END

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
    @BalPrin = m.Principal - m.Prin, -- isNull(t.Prin,0),
    @BalInt  = m.Interest -m.IntR, -- isNull(t.IntR,0),
    @BalOth  = m.Others - Discounted - m.Oth, -- isNull(t.Oth,0),
    @PaidWaived = isNull(WaivedInt,0),
	@ShouldBalPrin = isNull(i.Prin,0),
	@ShouldBalInt = isNull(i.IntR,0),
	@ShouldBalOth = isNull(i.Oth,0)
  FROM lnMaster m
  --,(SELECT 
  --    Sum(Prin) Prin, Sum(IntR) IntR, Sum(WaivedInt) WaitedInt, Sum(Oth) Oth
  --  FROM trnMaster
  --  WHERE Acc = @Acc and trnType in (3001,3097,3098,3099,3899,3201,3202) 
  -- ) t
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
    @TrnDate TrnDate, @RefDate RefDate, @Prin Prin, @Int Intr, @Oth Oth, @TrnAmt TrnAmt, @mTrnAmt mTrnAmt, @WaivedInt WaivedInt, @PaidWaived PaidWaived, 
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

	  --select * from lnmaster where status = 30 and cid in (select cid from samaster where accttype = 40 and balance > 0)

/*
  select m.Balance,* from satrnmaster t
  inner join samaster m on m.Acc = t.Acc 
  where trndate = @TrnDate and m.CID = @CID ORDER BY Trn Desc
*/
--  IF @Bal > 0 and @Particulars = 'Full Payment' and @TrnType = 3099  
--    exec fixloantran @ACC,0,0 
--- Fixed special Loan   
  IF @AcctType = 330 exec usp_FixSpecialLoan @Acc

  SELECT Sum(trnAmt)-@mTrnAmt
  FROM
 (SELECT sum(prin+intr+oth) trnamt from trnmaster where orno = @orno
  UNION ALL
  SELECT sum(case when trntype%2=1 then trnamt else -trnamt end) trnamt from satrnmaster where orno = @orno) a

  SELECT * from loaninst where acc = @Acc

  SELECT * from trnmaster where acc = @acc and trndate = @trndate
  SELECT * FROM lnMaster WHERE Acc = @Acc

GO

UPDATE lnMaster SET
	Prin = isNull(t.Prin,0),
    IntR = isNull(t.IntR,0),
    Oth = isNull(t.Oth,0),
    WaivedInt = isNull(t.WaivedInt,0)
FROM
   (SELECT 
      Acc, Sum(Prin) Prin, Sum(IntR) IntR, Sum(WaivedInt) WaivedInt, Sum(Oth) Oth
    FROM trnMaster
    WHERE trnType in (3001,3097,3098,3099,3899,3201,3202,3901) 
	GROUP BY Acc
   ) t 
 WHERE lnMaster.Acc = t.Acc 


BEGIN TRAN
 EXEC usp_Updlntran
  --@Acc =  '0000-4048-0000021',  
  --@Acc = '0000-4050-0030196',
  --@Acc = '0000-4048-0003447', 
  @Acc = '0101-4440-0000110',
  @TrnType    = 3001,      
  @mTrnAmt     =  13352.66,  
  @OrNo          = 3000000,    
  @PostBy      = 'sa',     
  @TermID       = '',      
  @Particulars = '',       
  @RefNo       = '',      
  @lnStatus      = 0,      
  @TrnDate   = '2021-3-17'

ROLLBACK

--select * from lnmaster where cid = 343 and status = 30


/*
SELECT * 
FROM
 (SELECT 
    m.Acc, 
	m.Principal - isNull(t.Prin,0) BalPrin,
	m.Interest - isNull(t.IntR,0) - isNull(t.WaivedInt,0) BalInt,
	m.Others - isNull(t.Oth,0) BalOth,
	isNull(t.WaivedInt,0) WaivedInt,
	m.Prin - isNull(t.Prin,0) lnPrin,
    m.IntR - isNull(t.IntR,0) lnInt,
    m.Oth - isNull(t.Oth,0)   lnOth,
    m.WaivedInt - isNull(t.WaivedInt,0) lnWaived,
	isNull(i.Prin,0) - isNull(t.Prin,0) iPrin,
	isNull(i.IntR,0) - isNull(t.IntR,0) iInt,
	isNull(i.Oth,0)  - isNull(t.Oth,0)  iOth
  FROM lnMaster m
  LEFT JOIN
   (SELECT 
      Acc, Sum(Prin) Prin, Sum(IntR) IntR, Sum(WaivedInt) WaivedInt, Sum(Oth) Oth
    FROM trnMaster
    WHERE trnType in (3001,3097,3098,3099,3899,3201,3202,3901) 
	GROUP BY Acc
   ) t on t.Acc = m.Acc
  LEFT JOIN 
   (SELECT 
      Acc,
	  Sum(CASE WHEN InstFlag = 9 THEN Prin ELSE CASE WHEN InstPD > IntR+Oth THEN InstPD-IntR-Oth ELSE 0 END END) Prin, 
	  Sum(CASE WHEN InstFlag = 9 THEN IntR ELSE CASE WHEN InstPD > Oth     THEN (CASE WHEN InstPD-Oth > IntR THEN IntR ELSE InstPD-Oth END) ELSE 0 END END) IntR, 
	  Sum(CASE WHEN InstFlag = 9 THEN Oth  ELSE CASE WHEN InstPD > Oth THEN Oth ELSE InstPD END END) Oth
    FROM LoanInst i
	GROUP BY Acc
      ) i on i.Acc = m.Acc
  WHERE m.Status in (30,99,98)) d
WHERE abs(lnPrin)+abs(lnInt)+abs(lnOth)+abs(iPrin) <> 0 
order by lnprin
*/
