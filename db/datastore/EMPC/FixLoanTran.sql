alter PROCEDURE FixLoanTran(
  @Acc as VarChar(17), @nAmt as Numeric(14,2),
  @bAdd as bit = 0, @TrnDate as DateTime = 0)
AS
  DECLARE --@Acc        as VARCHAR(17),
    @CID       as INT,
    @dNum      as Int,
    @iPrinEnd   as Numeric(14,2),
    @iIntEnd    as Numeric(14,2),
    @iPrinAmort as Numeric(14,2),
    @iIntAmort  as Numeric(14,2),
    @iOthEnd    as Numeric(14,2), 
    @Bal        as Numeric(14,2), 
    @PrinBal    as Numeric(14,2),
    @IntBal     as Numeric(14,2),
    @OthBal     as Numeric(14,2),
    @PrinPaid    as Numeric(14,2),
    @IntPaid     as Numeric(14,2),
    @OthPaid     as Numeric(14,2),
    @TrnAmt     as Numeric(14,2),  
    @PrinCr     as Numeric(14,2),
    @PrinDr     as Numeric(14,2),
    @IntCr      as Numeric(14,2),
    @IntDr      as Numeric(14,2),
    @OthCr      as Numeric(14,2),
    @OthDr      as Numeric(14,2),
    @Trn        as Int,
    @PrinRel    as Numeric(14,2),
    @IntRel     as Numeric(14,2),
    @nShort     as Numeric(14,2),
    @TrnType    as int,
    @WaivedInt  as Numeric(14,2),
    @WaivedInt2  as Numeric(14,2),
    @OthRel     as Numeric(14,2),
    @iOthAmort  as Numeric(14,2)

  IF @TrnDate = 0
    SELECT @TrnDate = ebSysDate
    from orgparms

--IF (SELECT Count(Acc) from TrnMaster 
--          Where Acc = @Acc and TrnDate > @TrnDate) = 0
  BEGIN
  --SET @Acc  = '0213-4042-0000021'
  --SET @Bal  = 6833.2+204+1000
  DECLARE @TotalPaid as Numeric(14,2)
  DELETE trnmaster 
  WHERE acc = @Acc 
  AND Particulars = 'To Correct Balance'
  AND trndate = @TrnDate

  If @bAdd = 1
  BEGIN
    SELECT 
      @TotalPaid = IsNull(Sum(Prin+IntR+Oth+WaivedInt),0),
      @WaivedInt2 = IsNull(Sum(WaivedInt),0)
    FROM trnMaster    
    WHERE acc = @Acc 
      and trnType in (3001,3097,3098,3099,3899,3202,3201) 

    SELECT 
      @Bal = Principal + interest - Discounted + OTHERS  
    FROM lnmaster
    WHERE acc = @Acc
    Set @Bal = @Bal - @TotalPaid + @nAmt
    END ELSE SET @Bal = @nAmt

  EXEC FixLoanInst @Acc

  SET @WaivedInt2 = IsNull(@WaivedInt2,0) 
  SET @dNum    = 1
  SET @PrinCr  = 0
  SET @PrinDr  = 0
  SET @IntCr   = 0
  SET @IntDr   = 0
  SET @PrinBal = 0
  SET @IntBal  = 0
  SET @OthBal  = 0
  SET @OthDr = 0
  SET @OthCr = 0

  SELECT 
    @dNum       = dNum,
    @iPrinEnd   = EndBal,
    @iIntEnd    = EndInt,
    @iPrinAmort = Prin,
    @iIntAmort  = IntR,
    @iOthAmort = oth,
    @iOthEnd = endoth
  FROM LoanInst  
  WHERE  
    Acc = @Acc
    and @Bal Between EndBal+EndInt+EndOth and EndBal+EndInt+EndOth+Prin+IntR+Oth-.0001
  ORDER BY dnum

  If @iPrinAmort is null AND @Bal < 0
  BEGIN
    SELECT 
      @dNum       = dNum,
      @iPrinEnd   = 0,
      @iIntEnd    = 0,
      @iPrinAmort = @Bal,
      @iIntAmort  = 0
    FROM LoanInst  
    WHERE  Acc = @Acc
      and dNum = (Select max(dNum) from loaninst where acc = @Acc)
    ORDER BY dnum
    Set @Bal = @iPrinEnd + @iIntEnd + @iPrinAmort + @iIntAmort+@iOthEnd +@iOthAmort
    END  

  If @iPrinAmort is null
  BEGIN
    SELECT 
      @dNum       = dNum,
      @iPrinEnd   = EndBal,
      @iIntEnd    = EndInt,
      @iPrinAmort = Prin,
      @iIntAmort  = IntR,
      @iOthEnd = EndOth,
      @iOthAmort = oth
    FROM LoanInst  
    WHERE Acc = @Acc
      and dNum = 1
    ORDER BY dnum
    Set @Bal = @iPrinEnd + @iIntEnd + @iPrinAmort + @iIntAmort + @iOthEnd +@iOthAmort
    END      

  SELECT @CID = CID, @PrinRel = Principal, @IntRel = Interest - Discounted, @OthRel = OTHERS
    FROM lnMaster
    WHERE Acc = @Acc

  SELECT @PrinPaid = Sum(Prin), @IntPaid = Sum(IntR+WaivedInt) ,@OthPaid = Sum(OTH)
  FROM trnMaster  
  WHERE acc = @Acc 
  and trnType in (3001,3097,3098,3099,3899,3202,3201)  

  SELECT Sum(Prin), Sum(IntR+WaivedInt) 
  FROM trnMaster  
  WHERE acc = @Acc 
    and trnType in (3001,3097,3098,3099,3899,3202,3201)  

  Select @IntPaid
  Set @PrinPaid = isNull(@PrinPaid,0)
  Set @IntPaid  = isNull(@IntPaid,0)

  SET @TrnAmt  =  (isNull(@PrinRel,0)+isNull(@IntRel,0)-isNull(@PrinPaid,0)-isNull(@IntPaid,0)+isNull(@OthRel,0)-isNull(@OthPaid,0)) - isNull(@Bal,0)
  SET @nShort  = @Bal - (@iPrinEnd + @iIntEnd+@iOthEnd) 
  SET @PrinBal = @iPrinEnd
  SET @IntBal  = @iIntEnd

----------------------------------------------------------------------
--   SELECT @TrnAmt TrnAmt, @nShort short, @Acc Acc, @Bal Mustbe, @PrinRel Release, 
--          @IntRel Interest, @PrinPaid Payment, @iPrinEnd schedrinbal,
--          @PrinRel+@IntRel-@PrinPaid-@IntPaid PrevBal
----------------------------------------------------------------------

  IF @nShort < @iPrinAmort
  BEGIN
    SET @PrinBal = @PrinBal + @nShort
    SET @nShort  = 0
    END ELSE 
  BEGIN
      SET @PrinBal = @PrinBal + @iPrinAmort
      SET @nShort  = @nShort
  - @iPrinAmort  -isnull(@iOthAmort,0)
  END

  SET @IntBal = @IntBal + @nShort
  SET @PrinCr  = (@PrinRel-@PrinBal) - @PrinPaid
  SET @IntCr   = (@IntRel -@IntBal)  - @IntPaid
  SET @OthCr   = (@OthRel -@OthBal)  - @OthPaid

--SELECT @PrinBal, @IntBal, @PrinBal+@IntBal
  IF @PrinCr < 0 
  BEGIN 
    SET @PrinDr = -@PrinCr 
    SET @PrinCr = 0 
  END

  IF @IntCr < 0
  BEGIN 
    SET @IntDr = -@IntCr 
    SET @IntCr = 0 
  END

  IF @OthCr < 0 
  BEGIN 
    SET @OthDr = -@OthCr 
    SET @OthCr = 0 
  END

--select ABS(@PrinCr),ABS(@PrinDr),ABS(@IntCr),ABS(@IntDr)
  SET @WaivedInt = 0

  IF @bAdd = 0
  BEGIN
    SET @WaivedInt = @PrinCr-@PrinDr+@IntCr-@IntDr
    SET @IntCr = -@PrinCr
    SET @IntDr = -@PrinDr
    END   

--select 2,ABS(@PrinCr),ABS(@PrinDr),ABS(@IntCr),ABS(@IntDr)
  IF ABS(@WaivedInt)+ABS(@PrinCr)+ABS(@PrinDr)+ABS(@IntCr)+ABS(@IntDr) <> 0
--IF @PrinCr+@PrinDr+@IntCr+@IntDr = 0
  BEGIN
    IF @TrnAmt > 0
    BEGIN
      SET @TrnType = 3001
      END ELSE 
      SET @TrnType = 3098

    SELECT @Trn = Max(Trn) from trnMaster Where TrnDate = @TrnDate
    SET @Trn = isnull(@Trn,0) + 1
    INSERT trnMaster (
      ACC, trnDate, TRN, TrnType, OrNo, TrnAmt, Prin,
      IntR, Oth, Penalty,
      WaivedInt, Balance, UserName, TermID, 
      RefNo, TrnDesc, TrnMnem_CD, Particulars, 
      [Time], Cancel)
    VALUES(
      @Acc, @TrnDate, @Trn, @TrnType, 0, @TrnAmt, @PrinCr-@PrinDr, 
      @IntCr-@IntDr, IsNull(@OthCr-@OthDr,0), 0,
      @WaivedInt, @Bal, 'sa', 'Server', 
      '0', '', 43, 'To Correct Balance', GetDate(), 0)
    END

  UPDATE LoanInst SET 
    InstPD   = Prin+IntR+Oth - (@Bal - (@iPrinEnd + @iIntEnd+@iOthEnd)),
    InstFlag = CASE WHEN (@Bal - (@iPrinEnd + @iIntEnd + @iOthEnd)) = 0 THEN 9 ELSE 0 END
  WHERE Acc = @Acc and dNum = @dNum

  UPDATE LoanInst 
  SET InstPD   = Prin+IntR+Oth,
      InstFlag = 9
  WHERE Acc = @Acc and dNum < @dNum

  UPDATE LoanInst
  SET InstPD   = 0,
      InstFlag = 0
  WHERE Acc = @Acc and dNum > @dNum

-- Get Total Collection
  SELECT 
    @PrinCr = IsNull(Sum(Prin),0),
    @PrinDr = 0, 
    @IntCr  = IsNull(Sum(IntR),0),
    @WaivedInt = IsNull(Sum(WaivedInt),0),
    @IntDr  = 0,
    @OthDr = 0,
    @OthCr = IsNull(Sum(OTH),0)
  FROM trnMaster
  WHERE acc = @Acc 
         --and trnType in (3001,3097,3098,3099,3899)  
    and trnType in (3001,3097,3098,3099,3899,3202,3201)  

  print '----'
  print isnull(@IntCr ,0)
  UPDATE lnMaster SET 
    Prin      = isnull(@PrinCr-@PrinDr,0),
    IntR      = isnull(@IntCr -@IntDr + Discounted,0),Status = CASE WHEN round(@Bal,2) = 0 THEN 99 ELSE 30 END,
    WaivedInt = @WaivedInt,
    Oth = isnull(@OthCr-@OthDr,0)
  Where Acc   = @Acc

-- Fix TrnMaster
--DEClare @acc  as varchar(17), @sAcc as varchar(17), @TrnAmt as Numeric(14,2), @Bal as Numeric(14,2)
--set @sAcc = ''

  DECLARE lnTrn CURSOR
  KEYSET FOR 
    SELECT Acc, Prin+IntR+WaivedInt+OTH
    FROM trnmaster 
    WHERE Acc = @Acc AND trnType in (3001,3097,3098,3099,3899,3201,3202) 
    ORDER BY Acc,trnDate,trn

--SET @Bal = @PrinRel + @IntRel  --Updated 05-13-2020 by Anthony Caya
  SET @Bal = @PrinRel + @IntRel +@OthRel
  OPEN lnTrn
  FETCH NEXT FROM lnTrn INTO @Acc, @TrnAmt
  WHILE (@@fetch_status <> -1)
  BEGIN
    IF ( @@fetch_status <> -2)
    BEGIN
      -- select 'b',@TrnAmt,@Bal
      --        IF @Acc <> @sAcc 
      --        BEGIN
      --           SELECT @Bal = Principal + Interest - Discounted from lnMaster Where acc = @Acc
      --           Set @sAcc = @Acc
      --        END
      SET @Bal = @Bal - @TrnAmt
      --select 'ta', @TrnAmt,@Bal 
      UPDATE trnMaster
      SET Balance = @Bal
      WHERE CURRENT OF lnTrn
      END
        FETCH NEXT FROM lnTrn INTO @Acc, @TrnAmt
    END
  CLOSE lnTrn
  DEALLOCATE lnTrn

  --Declare @acc as varchar(17)
  --Set @acc = '0213-4042-0000126'

  SELECT 
    'New' Source, 
    Prin PrinPaid, 
    IntR IntPaid, WaivedInt, Balance
  FROM trnMaster
  WHERE Acc = @Acc and TrnDate = @TrnDate and Trn = @Trn

  UNION All
  SELECT 
    'Master' Source, Prin PrinPaid, IntR IntPaid, WaivedInt,
    Principal+Interest+Others - (Prin+IntR+WaivedInt+Oth) Balance
  FROM  lnMaster 
  WHERE acc = @Acc

  UNION All
  SELECT 
    Source, Sum(PrinPaid) PrinPaid, Sum(IntPaid) IntPaid, Sum(WaivedInt) WaivedInt, Sum(Balance) Balance
  FROM
   (SELECT 
      'TRAN' Source, Sum(Prin) PrinPaid, 
      Sum(IntR) IntPaid, Sum(WaivedInt) WaivedInt, 0 Balance
    FROM trnmaster 
    WHERE 
      acc = @Acc           
      and trnType in (3001,3097,3098,3099,3899,3202,3201)   

  UNION All
  SELECT 
    'TRAN' Source, 
    0 PrinPaid, 0 IntPaid, 0 WaivedInt, Balance
  FROM trnmaster 
  WHERE acc = @Acc           
    and cast(TrnDate as BigInt)*10000000+Trn in 
    (SELECT Max(cast(TrnDate as BigInt)*10000000+Trn) 
     FROM trnMaster 
     WHERE Acc = @Acc                    
        and trnType in (3001,3097,3098,3099,3899))
    ) as d
  GROUP BY source

  UNION All
  SELECT 
    Source, Sum(PrinPaid) PrinPaid, Sum(IntPaid) IntPaid, 0 WaivedInt, Sum(Balance)  Balance
  FROM
   (SELECT 'Sched' Source,
      Sum(CASE WHEN (InstPD-IntR-oth) > 0 THEN 
            CASE WHEN Prin > InstPD-IntR THEN InstPD-IntR ELSE PRIN END 
          ELSE 0 END) PrinPaid, 
      Sum(CASE WHEN InstPD <> 0 THEN 
            CASE WHEN Intr > InstPD THEN InstPD ELSE IntR END 
          ELSE 0 END) IntPaid, 
      CASE WHEN InstFlag  = 0 THEN Sum(Prin+Intr-InstPD) ELSE 0 END Balance 
    FROM LoanInst
    WHERE acc = @Acc 
    GROUP BY InstFlag, InstPD) as d
    GROUP BY source
  END
