-----------------------------------------------------------------------------------     
ALTER PROCEDURE [dbo].usp_FixSpecialLoan (@Acc VarChar(17))      
-----------------------------------------------------------------------------------  
as      
  DECLARE @LN TABLE (      
    [DNUM] [smallint] Identity(1,1) NOT NULL,      
    [ACC] [varchar](22) NOT NULL,      
    [DUEDATE] [smalldatetime] NOT NULL,      
    [INSTFLAG] [smallint] NOT NULL,      
    [PRIN] [money] NOT NULL,      
    [INTR] [money] NOT NULL,      
    [Oth] [numeric](16, 2) NULL,      
    [PENALTY] [money] NOT NULL,      
    [ENDBAL] [money] NOT NULL,      
    [ENDINT] [money] NOT NULL,      
    [EndOth] [numeric](16, 2) NULL,      
    [INSTPD] [money] NOT NULL,      
    [PenPD] [money] NOT NULL,      
    [CarVal] [money] NULL,      
    [UpInt] [money] NULL,      
    [ServFee] [money] NULL      
  )      
      
  INSERT @LN      
  SELECT Distinct      
    @Acc acc, DueDate, 0 InstFlag, TrnAmt Prin, 0 IntR, 0 Oth,      
    0 Penalty, 0 EndBal, 0 EndInt, 0 EndOth, 0 InstPD,       
    0 PenPD, 0 CarVal, 0 UpInt, 0 ServFee     
  FROM        
   (SELECT a, CASE WHEN a=1 THEN TrnDate ELSE DueDate END DueDate,   
      SUM(CASE WHEN a=1 THEN Prin+IntR ELSE 0 END) TrnAmt  
    FROM trnMaster t   
    INNER JOIN 
      (SELECT Acc, MAX(DueDate) DueDate FROM LoanInst GROUP BY Acc) i on i.Acc = t.Acc  
    INNER JOIN 
      (SELECT 1 a UNION SELECT 2) a on 0=0  
    WHERE trnType in (3001,3097,3098,3099,3899,3202,3201,3901) and Prin+IntR+Oth <> 0  
      and t.Acc = @Acc  
    GROUP BY CASE WHEN a=1 THEN TrnDate ELSE DueDate END, a) d  
  WHERE (trnamt <> 0 or a = 2)  
  
  DECLARE       
    @Principal Money,      
    @Interest Money,      
    @TrnAmt Money,      
    @IntRate Float,      
    @PrinBal Money,      
    @IntBal Money,      
    @OrgDue Money,      
    @NetDue Money,      
    @DueDate DateTime,      
    @pDate DateTime,      
    @dNum Int,      
    @Days Int,      
    @Int Money,      
    @Prin Money,      
    @Int2 Money,      
    @Prin2 Money,      
    @CarVal Money,      
    @CumInt Money,      
    @Net Money,      
    @NewDue Money,    
    @CutDate DateTime,  
    @Maturity DateTime,  
    @SysDate DateTime,  
    @NextDueDate DateTime,  
    @MaxdNum Int,  
    @ColinYear Money,  
    @ColinYearInt Money,  
    @i Int  
      
  SELECT @Maturity = MAX(DueDate), @MaxdNum = Max(DNUM) FROM LoanInst Where Acc = @Acc
  SELECT @SysDate = ebSysDate from OrgParms         
    
  SELECT @pDate =   
     CAST(CAST(Year(@SysDate)-1 as VarChar(4)) + '/' +    
     CAST(Month(@Maturity) as VarChar(2)) + '/' +  
     CAST(Day(@Maturity) as VarChar(2)) as DateTime)  
    
  SET @pDate = CASE WHEN @pDate < @SysDate THEN @pDate ELSE DATEADD(yy,1,@pDate) END  
    
  SET @dNum = @MaxdNum   
  
  WHILE Year(@pDate) < YEAR(@Maturity)   
  BEGIN  
    SET @pDate = DATEADD(yy,1,@pDate)  
    --SELECT @pDate, @Maturity   
      
    IF @dNum <> @MaxdNum  
    BEGIN  
      INSERT @LN  
      SELECT Distinct      
        @Acc, @pDate, 0 InstFlag, 0 Prin, 0 IntR, 0 Oth,      
        0 Penalty, 0 EndBal, 0 EndInt, 0 EndOth, 0 InstPD,       
        0 PenPD, 0 CarVal, 0 UpInt, 0 ServFee     
      END ELSE     
      UPDATE @LN SET DueDate = @pDate WHERE DNUM = @dNum  
        
      SET @dNum = @dNum + 1        
      END      
  
   -- SELECT @Maturity = MAX(DueDate), @MaxdNum = Max(DNUM) FROM @LN  
      
    SET @CumInt = 0      
    SET @OrgDue = 0      
    SET @NetDue = 0      
    SET @Interest = 0                
    SET @CutDate = '2015-01-01'    
    SELECT       
      @pDate = DisbDate, @Principal = Principal       
    FROM lnMaster       
    WHERE @Acc = Acc       
  
    DECLARE ln CURSOR KEYSET      
    FOR          
      SELECT dNum, DueDate, Prin, IntR, UpInt, PenPD      
      FROM @LN   
      WHERE Prin+IntR+Oth < 0     
      ORDER BY dNum      
        
    OPEN ln      
    FETCH NEXT FROM ln INTO @dNum, @DueDate, @Prin, @Int, @TrnAmt, @OrgDue      
    SET @PrinBal = @Principal      
        
    WHILE (@@fetch_status <> -1) -- and @dNum < 3      
    BEGIN      
      IF (@@fetch_status <> -2)       
      BEGIN  
        SET @NewDue = -(@Prin+@Int)  
        SET @i = @dNum - 1  
        WHILE @NewDue > 0 and @i > 1  
        BEGIN         
          SELECT @Prin = Prin+Intr FROM @LN WHERE DNUM = @i  
          IF @Prin > @NewDue  
          BEGIN  
            SET @Int = @Prin - @NewDue  
          END ELSE SET @Int = @Prin  
          SET @NewDue = @NewDue-@Int  
          SET @Prin = @Prin - @Int  
            
        -- select @Int rhick, @i  
          UPDATE @LN   
          SET PRIN = @Prin, INTR = 0   
          WHERE DNUM = @i  
            
        -- select * from @LN WHERE DNUM = @i  
          SET @i = @i-1  
          END                
        
        UPDATE @LN   
          SET PRIN = @NewDue, INTR = 0   
          WHERE @dNum = dNum  
        END      
        FETCH NEXT FROM ln INTO @dNum, @DueDate, @Prin, @Int, @TrnAmt, @OrgDue      
      END                   
    -- UPDATE @LN SET EndInt = Sum(IntR)      
    CLOSE ln      
    DEALLOCATE ln      

  --------      
  DECLARE ln CURSOR KEYSET      
  FOR          
    SELECT dNum, DueDate, Prin, IntR, UpInt, PenPD      
    FROM @LN      
    ORDER BY dNum      
        
  SET @dNum =0      
  OPEN ln      
  FETCH NEXT FROM ln INTO @dNum, @DueDate, @Prin, @Int, @TrnAmt, @OrgDue      
  SET @PrinBal = @Principal      
        
  SET @CarVal = @Principal      
  SET @IntBal = 0      
  SET @Net = 0      
  WHILE (@@fetch_status <> -1) -- and @dNum < 3      
  BEGIN      
    IF (@@fetch_status <> -2)       
    BEGIN                   
      IF @DueDate >= @CutDate    
      BEGIN      
        SET @IntRate = .08      
      END ELSE SET @IntRate = .1      
      SET @Days = DateDiff(dd, @pDate, @DueDate)      
      
      IF @DueDate > @CutDate and @pDate < @CutDate -- Change of Interest from 10% to 8%      
      BEGIN      
        SET @Days = DateDiff(dd, @pDate, @CutDate)-1    
        SET @Int2 = Round(.1 * @CarVal / 365 * @Days,2)      
        SET @Days = DateDiff(dd, @CutDate, @DueDate) + 1    
        SET @Int2 = @Int2 + Round(.08 * @CarVal / 365 * @Days,2)      
      END ELSE SET @Int2 = Round(@IntRate * @CarVal / 365 * @Days,2)      
  
  select @Days, @Int2, @CutDate,  @pDate,  @DueDate, @CarVal, @prin

      SET @CumInt = @CumInt + @Int2      
        
  --    if @dnum = 38      
  --    select 'p',@CumInt, @trnAmt Trnamt, @PrinBal PrinBal, @dNum dnum, @Prin Prin, @Int Int, @Int2 Int2, @Days Days, @DueDate DueDate, @pDate PrevDate      
        
      SET @NewDue = @Prin + @Int    
        
      IF @NewDue = 0   
      BEGIN  
        SELECT @ColinYear = IsNull(Sum(Prin),0),   
                @ColinYearInt = IsNull(Sum(IntR),0) FROM @LN   
        WHERE DUEDATE Between DateAdd(yy,-1,@DueDate)+1 and @DueDate  
        SELECT @NewDue = Prin+IntR FROM LoanInst where DueDate = @DueDate and Acc = @Acc
		IF isNull(@NewDue,0) = 0
           SET @NewDue = dbo.PMT(.08,@MaxdNum-@dNum+1 , @PrinBal+@ColinYear)-@ColinYear-@ColinYearInt  
        SET @NewDue = round(CASE WHEN @NewDue < 0 THEN 0 ELSE @NewDue END,2)
      END  
        select @carval
      SET @Int  = CASE WHEN @NewDue > @CumInt  THEN @CumInt  ELSE @NewDue END      
      SET @NewDue = @NewDue - @Int      
      SET @Prin = CASE WHEN @NewDue > @PrinBal THEN @PrinBal ELSE @NewDue END      
      SET @NewDue = @NewDue - @Prin         
      SET @PrinBal = @PrinBal - @Prin      
            select @prin,@int
      SET @CarVal = @CarVal - @Prin      
      SET @NetDue = @NetDue + @OrgDue - @Prin - @Int      
        
  --select 'int', @CumInt, @Int      
        --select @IntRate      
      SET @Interest = @Interest + @Int2      
        
      UPDATE @LN       
      SET       
        Prin = @Prin,      
        IntR = @Int,      
        EndBal = @PrinBal,      
        CarVal = @CarVal,      
        UpInt = @Int2      
        --PenPD = @NetDue,      
        --Penalty = @NetDue,      
        --EndOth = @TrnAmt      
      WHERE dNum = @dNum      
        
      SET @CumInt = @CumInt - @Int      
      SET @pDate = @DueDate  
        
    -- select @dNum d, @CumInt, @Interest      
    END      
  FETCH NEXT FROM ln INTO @dNum, @DueDate, @Prin, @Int, @TrnAmt, @OrgDue      
  END      
        
  UPDATE @LN SET Prin = Prin+EndBal, EndBal = 0, IntR = @Int + @CumInt WHERE dNum = @dNum       
        
  -- UPDATE @LN SET EndInt = Sum(IntR)      
  CLOSE ln      
  DEALLOCATE ln      
          
  UPDATE @ln       
  SET EndInt =  i.IntBal      
  FROM @ln i2      
  LEFT JOIN      
   (SELECT i1.dNum, Sum(IsNull(i2.IntR,0)) IntBal       
  FROM @LN i1      
  LEFT JOIN @LN i2 on i1.dNum < i2.dNum       
  GROUP BY i1.dNum) i on i.dNum = i2.dNum      
    
  DELETE LoanInst WHERE Acc = @Acc       
  INSERT LoanInst       
    SELECT * FROM @LN      
  
  UPDATE LNMASTER Set INTEREST = @Interest WHERE Acc = @Acc      
  EXEC FixLoanTran @Acc, 0, 1      
  GO
