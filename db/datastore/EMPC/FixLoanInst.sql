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

BEGIN TRAN
  EXEC FixLoanInst '0000-4048-0003447'

  select * from loaninst where acc = '0000-4048-0003447'

  select prin+intR, * from trnmaster where acc =  '0000-4048-0003447'
  order by trndate,trn
ROLLBACK

