sp_helptext usp_UpdlnTranText
---------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------
CREATE PROCEDURE [dbo].[usp_Updlntran](
  @Acc         as Char(17),        
  @TrnType     as SmallInt    = 3001,      
  @mTrnAmt     as Money       = 0,  
  @OrNo        as Numeric     = 0,    
  @PostBy      as VarChar(15) = 'sa',     
  @TermID      as VarChar(15) = '',      
  @Particulars as VarChar(100)= '',       
  @RefNo       as VarChar(16) = '',      
  @lnStatus    as Int         = 0,      
  @TrnDate     as DateTime    = 0)
AS        
  DECLARE 
    @CID        as INT,       
    @dNum       as Int,       
    @RefDate as DateTime,       
    @SysDate as DateTime,      
    @Trn        as Int,       
    @Frequency  as Int,       
    @TrnAmt     as Numeric(14,2),         
    @iPrinEnd   as Numeric(14,2), @iIntEnd    as Numeric(14,2), @iOthEnd    as Numeric(14,2),       
    @iPrinAmort as Numeric(14,2), @iIntAmort  as Numeric(14,2), @iOthAmort  as Numeric(14,2),       
    @Bal        as Numeric(14,2), @PrinBal    as Numeric(14,2), @IntBal     as Numeric(14,2), @OthBal     as Numeric(14,2),       
    @PrinPaid   as Numeric(14,2), @IntPaid    as Numeric(14,2), @OthPaid    as Numeric(14,2),       
    @Prin       as Numeric(14,2), @IntR       as Numeric(14,2),       
    @Oth        as Numeric(14,2), @Penalty    as Numeric(14,2),       
    @PrinRel    as Numeric(14,2), @IntRel     as Numeric(14,2), @OthRel     as Numeric(14,2),       
    @nShort     as Numeric(14,2), @TotalPaid  as  Numeric(14,2),  @CanWaived  as Numeric(14,2),    
    @PaidWaived as Numeric(14,2), @TrnDesc    as VarChar(100),       
    @MnemCode   as Int,   @UpdLn      as Bit, @RunState as Int,       
    @Excess     as Numeric(14,2),       
    @SavTrnAmt  as Numeric(14,2), @WaivableInt as Bit,   
    @AcctType Numeric   

  SET @TrnAmt = Cast(@mTrnAmt as Numeric(14,2))       
  SET @SavTrnAmt = @TrnAmt       

-- 3001 Loan Collection/Repayment   43       
-- 3097 Cancel Close                56       
-- 3098 Cancel Payment              54       
-- 3099 Repayment To Close          48       
-- 3899 Loan Renewal                0       
-- 3201 L/N Journal Credit          43       
-- 3202 L/N Journal Debit           43       
  SET @MnemCode = 
    CASE @TrnType        
      WHEN 3001 THEN 43                           
      WHEN 3097 THEN 56                           
      WHEN 3098 THEN 54          
      WHEN 3099 THEN 48        
      WHEN 3899 THEN 0        
      WHEN 3201 THEN 43       
      WHEN 3202 THEN 43 ELSE 43       
      END       

  DECLARE @LoanInst TABLE (       
    DNUM smallint,       
    ACC varchar(22),       
    DUEDATE datetime,       
    PRIN Numeric(14,2),       
    ENDBAL Numeric(14,2),       
    INTR Numeric(14,2),       
    ENDINT Numeric(14,2),       
    INSTAMT Numeric(14,2),       
    INSTPD Numeric(14,2),       
    PENALTY Numeric(14,2),       
    PenPD Numeric(14,2),       
    CarVal Numeric(14,2),       
    UpInt Numeric(14,2),       
    ServFee Numeric(14,2),       
    Oth Numeric(14,2),       
    EndOth Numeric(14,2))       

-- Get the system Date       

  SELECT @SysDate = ebSysDate, @RunState = RunState, @WaivableInt = IsNull(WaivableInt,1) FROM orgParms       

  --overwirte added BY ANthony Caya
  IF @AcctType = 315
  BEGIN
	  Set @WaivableInt = 0
  END
  
  IF @TrnDate = 0       
    SELECT @TrnDate = @SysDate       

  IF @RunState <> 0       
  BEGIN       
    RAISERROR ('Cannot Transaction. System is not open yet...', 16, 1)        
    RETURN       
  END       

  INSERT @LoanInst        
  SELECT 
    DNUM, ACC, DUEDATE, PRIN, ENDBAL, INTR,         
    ENDINT, Prin+IntR+Oth, INSTPD, PENALTY, PenPD,       
    IsNull(CarVal,EndBal), IsNull(UpInt,0), IsNull(ServFee,0), IsNull(Oth,0), IsNull(EndOth,0)       
  FROM LoanInst Where @Acc = Acc       

-- Get the MnemCode       
-- SET @MnemCode = CASE WHEN @TrnType in (3097,3099,3899) THEN 48 ELSE 43 END       
-- Get the Waived Interest or the remaining interest and Other information from lnmaster       

  SELECT 
    @PrinRel   = Principal,       
    @CID       = CID,        
    @Frequency = Frequency,
--Modified June 6, 2017 PLP loan due was deducted to contribbution account
  --@Bal       = Principal + Interest + Others -Oth - Discounted,      
    @Bal       = Principal + Interest + Others ,   
    @AcctType  = AcctType   
  FROM lnMaster m       
  WHERE @Acc = m.Acc        
  
  print 'BAL@'
	print @Bal

  SELECT 
    @IntRel    = Sum(IntR),       
    @OthRel    = Sum(IsNull(Oth,0))       
   FROM @LoanInst m
   WHERE @Acc = m.Acc        
  
  SET @RefDate = dbo.RefDueDate(@Frequency,@SysDate,0)       
  
  SELECT @CanWaived = Sum(IntR)       
  FROM @LoanInst        
  WHERE DueDate-DatePart(dw,DueDate)+DatePart(dw,@RefDate) > =@RefDate        
    AND ACC = @Acc       
  
  SET @CanWaived = IsNull(@CanWaived,0)       
  
  IF @WaivableInt = 0 SET @CanWaived = 0       
-- Get the Transaction       
  SELECT 
    @TotalPaid  = isnull(Sum(Prin+IntR+Oth+Penalty),0),       
    @PaidWaived = isNull(Sum(WaivedInt),0),           
    @PrinPaid   = isnull(Sum(Prin),0),            
    @IntPaid    = isnull(Sum(IntR),0),       
    @OthPaid    = isnull(Sum(Oth),0)       
  FROM trnMaster         
  WHERE Acc = @Acc        
    and trnType in (3001,3097,3098,3099,3899,3201,3202)       
           
-- se       
--    SELECT isnull(Sum(lnPriCr-lnPriDr),0),isnull(Sum(lnIntCr-lnIntDr),0)       
--        FROM trnMaster         
--        WHERE Acc = @Acc        
--              and trnType in (3001,3097,3098,3099,3899,3201,3202)       

  SET @PrinPaid  = isNull(@PrinPaid,0)
  SET @IntPaid   = isNull(@IntPaid,0)       
  SET @OthPaid   = isNull(@OthPaid,0)       
  SET @TotalPaid = IsNull(@TotalPaid,0)       
  
  print @Bal
	print @TotalPaid
	print @TrnAmt

  SET @Bal = @Bal - @TotalPaid - @TrnAmt       
 
  Print '1'
	print @Bal

  IF @Bal - @CanWaived > 1 -- If the Remaining Balance is below P 1 it should Fully Paid and apply Waived Interest       
   SET @CanWaived = 0       
  
  IF @Bal - @PaidWaived  < 25 AND @CanWaived < @PaidWaived       
   SET @CanWaived = @PaidWaived       
  
  SET @Bal = @Bal - @CanWaived       

  IF @Bal < 0        
   SET @Bal = 0       

/*       
  IF @Bal < 0        
  BEGIN       
    SET @Excess = -@Bal       
    SET @Bal = 0        
    END        
  ELSE        
    SET @Excess = 0       
*/       

--------------------------------------------------       
-- Innialized Variables       
--------------------------------------------------       
  SET @dNum    = 1       
  SET @Prin    = 0       
  SET @IntR    = 0       
  SET @Oth     = 0      
  SET @Penalty = 0       
  SET @PrinBal = 0       
  SET @IntBal  = 0       

-- Get the Corresponding dNum of Balance from LoanInst       
  SELECT 
    @dNum       = dNum,       
    @iPrinEnd   = EndBal,       
    @iIntEnd    = EndInt,       
    @iOthEnd    = EndOth,       
    @iPrinAmort = Prin,       
    @iIntAmort  = IntR,       
    @iOthAmort  = Oth       
  FROM @LoanInst         
  WHERE Acc = @Acc       
    and @Bal Between EndBal+EndInt+EndOth 
    and EndBal+EndInt+EndOth+Prin+IntR+Oth-.0001       
  
  IF @iPrinAmort is null AND @Bal > 0 -- If Balance is more than 0 and Amortization Not Found       
  BEGIN -- Used the Information of dnum       
    SELECT @dNum       = dNum,       
          @iPrinEnd   = EndBal,       
          @iIntEnd    = EndInt,       
          @iOthEnd    = IsNull(EndOth,0),             
          @iPrinAmort = Prin,              
          @iIntAmort  = IntR,       
          @iOthAmort  = IsNull(Oth,0)       
     FROM @LoanInst         
     WHERE  Acc = @Acc and dNum = 1                    
     ORDER BY dnum                                                                                                                                                                                                                                                   
    END ELSE IF @iPrinAmort is null AND @Bal < 0 -- If Balance is Less than 0 and Amortization Not Found       
    BEGIN       
      SELECT @dNum = Max(dNum)       
      FROM @LoanInst         
      WHERE  Acc = @Acc    

      SET @iPrinEnd   = @Bal       
      SET @iIntEnd    = 0       
      SET @iOthEnd    = 0       
      SET @iPrinAmort = 0       
      SET @iIntAmort  = 0       
      SET @iOthAmort  = 0       
      END             
 
-- Update @Bal (There is a problem if @Bal is greater than {Principal+Interest}) "Need to solve"       
--   SET @Bal = @iPrinEnd + @iIntEnd + @iPrinAmort + @iIntAmort          
-- select @PrinRel prinrel,@IntRel,@PrinPaid,@IntPaid,@Bal       
--print 'test5'   
--print @prinrel   
--print @intrel   
--print @othrel   
--print @prinpaid   
--print @intpaid   
--print @canwaived   
--print @othpaid   
--print @bal   
--print 'test amount' + Convert(Varchar(10),@TrnAmt)   

  SET @TrnAmt  =  (@PrinRel+@IntRel+@OthRel-@PrinPaid-@IntPaid-@CanWaived-@OthPaid) - @Bal       
   --print 'test amount' + Convert(Varchar(10),@TrnAmt)   
------------------------------------***************************************************************       
-- select trntype,trndate,trn,trnamt,balance,* from trnmaster where acc = '0111-4043-0024027' order by trndate,trn       
-- select @trnamt trnamt , @Bal bal, @PrinRel prinrel, @IntRel intrel, @OthRel OthRel,        
--    @PrinPaid PrinPaid, @IntPaid IntPaid, @CanWaived CanWaived,@OthPaid OthPaid       
------------------------------------***************************************************************       
-- Difference between the Balance in LoanInst and the Balance Should Be after the transaction       
-- to get the payment for the corresponding dNum in LoanInst       

  SET @nShort  = @Bal - (@iPrinEnd + @iIntEnd + @iOthEnd)         
  SET @PrinBal = @iPrinEnd       
  SET @IntBal  = @iIntEnd       
  SET @OthBal  = @iOthEnd       

  print '@BAL'
  print @Bal

-- Checking Variable Values (for debugging only)       
----------------------------------------------------------------------       
--   SELECT @TrnAmt TrnAmt, @nShort short, @Acc Acc, @Bal Mustbe, @PrinRel Release,        
--          @IntRel Interest, @PrinPaid Payment, @iPrinEnd schedrinbal,       
--          @PrinRel+@IntRel-@PrinPaid-@IntPaid PrevBal       
----------------------------------------------------------------------       
  IF @nShort < @iPrinAmort -- Payment is not enough to pay the corresponding Principal Due       
  BEGIN       
    SET @PrinBal = @PrinBal + @nShort       
    SET @nShort  = 0       
    END
  ELSE     
  BEGIN -- There is a remaining Amount to pay part or full interest       
    SET @PrinBal = @PrinBal + @iPrinAmort       
    SET @nShort  = @nShort  - @iPrinAmort       
  END       

  SET @IntBal = @IntBal + @nShort       
-- Compute for the PrinCr and IntCr needed to be posted in trnmaster       
  print @PrinRel
  print @PrinBal
  print @PrinPaid
  SET @Prin    = (@PrinRel-@PrinBal) - @PrinPaid       
  print '----'
  print @IntRel
  print @IntBal
  print @IntPaid
  print @CanWaived

  SET @IntR    = (@IntRel -@IntBal)  - @IntPaid - @CanWaived       

  SET @Oth     = (@OthRel -@OthBal)  - @OthPaid       
  print @Prin
  print @IntR
    
  --   select 'to post', @trnamt, ABS(@PrinCr),ABS(@PrinDr),ABS(@IntCr),ABS(@IntDr)       
  
  SET @UpdLN = 0       
  IF ABS(@Prin)+ABS(@IntR)+ABS(@Oth)+ABS(@Penalty) <> 0       
  BEGIN         
    SET @UpdLN = 1       
    SELECT @TrnDesc = TrnDesc from trnTypes where TrnType = @TrnType       
-- EXEC @Trn = GetNewSerial 3       
-- SELECT @Trn = Max(Trn) from trnMaster Where TrnDate = @TrnDate       
-- SET @Trn = 0 --isnull(@Trn,0) + 1       
    SELECT @trn = Max(Trn)+1 --jhed        
    FROM  TrnMaster        
    WHERE TrnDate = @TrnDate       
    If isnumeric(@trn)=0     
      BEGIN     
        set @trn = 1     
      end     

  -- Correct the Waived Interest Value       
  -- CanWavied means both (Can be Waived and Cancel Waived)       
  SET @CanWaived = @CanWaived - @PaidWaived       
-- select 0,0,0,@CanWaived,@PaidWaived,       
--                 CAST(@TrnDate as Int),NULL,1,0,'A',       
--                 NULL,@TrnDate,@Trn,@Acc,0,3,@CID,@TrnType,NULL,NULL,       
--                 @PostBy,9,@TermID,Abs(@TrnAmt),@Bal,@RefNo,1477,@TrnDesc,30,       
--                 @PrinDr,@PrinCr,@IntDr,@IntCr,0,0,       
--                 0,0,@MnemCode,@Particulars,@Bal,       
--     0,0,0,0,0,0,GetDate(),0,@OrNo       
--sp_help trnmaster       
       -------------------------------------- 
 --For Loan Bal In Renewal 
 -------------------------------------- 
 IF @TrnType = 3899 and @Bal > 0 
 BEGIN 
  SET @CanWaived = @Bal  
  SET @Bal = 0 
  END 

  INSERT trnMaster (
    ACC, trnDate, TRN, TrnType, OrNo, TrnAmt, Prin,        
    IntR, Oth, Penalty,       
    WaivedInt, Balance, UserName, TermID,        
    RefNo, TrnDesc, TrnMnem_CD, Particulars,        
    [Time], Cancel)       
          
  VALUES(
    @Acc, @TrnDate, @Trn, @TrnType, @OrNo, @TrnAmt, @Prin,          
    @IntR, @Oth, @Penalty,       
    @CanWaived, @Bal, @PostBy, @TermID,        
    @RefNo, @TrnDesc, @MnemCode, @Particulars,       
    GetDate(), 0)       
          
--      IF @SysDate = @TrnDate       
--      UPDATE tblSerial               
--         SET ln_TrnSerial =        
--              isnull((SELECT Max(Trn) + 1        
--                         FROM trnmaster        
--                         WHERE trnDate = @TrnDate),0)       

  END       

--print @TrnAmt   
--print 'test3'    
-- Update LoanInst position current payment        

  UPDATE LoanInst SET 
    InstPD   = Prin+IntR+Oth - IsNull((@Bal - (@iPrinEnd + @iIntEnd + @iOthEnd)),0),       
    InstFlag = CASE WHEN (@Bal - (@iPrinEnd + @iIntEnd + @iOthEnd)) = 0 THEN 9 ELSE 0 END       
  WHERE Acc = @Acc and dNum = @dNum       
  
  UPDATE LoanInst SET 
    InstPD   = Prin+IntR+Oth,       
    InstFlag = 9       
  WHERE Acc = @Acc and dNum < @dNum       
   
  UPDATE LoanInst        
  SET InstPD   = 0,       
       InstFlag = 0           
  WHERE Acc = @Acc and dNum > @dNum       
-- Get Total Collection       
  SELECT 
    @Prin       = Sum(Prin),       
    @IntR       = Sum(IntR),       
    @PaidWaived = Sum(WaivedInt),       
    @Oth        = Sum(Oth),           
    @Penalty    = Sum(Penalty)       
  FROM trnMaster       
  WHERE acc = @Acc        
      and trnType in (3001,3097,3098,3099,3899,3201,3202)         
      
/*       
  DECLARE @ChrAmt as Money       
  SELECT @ChrAmt = Sum(ChrAmnt)       
   FROM lnchrgdata       
   WHERE Acc = @Acc
  SET @ChrAmt = IsNull(@ChrAmt,0)       
  SET @Oth = @Oth + @ChrAmt       
*/       
    
-- select @PrinCr,@PrinDr, @IntCr, @IntDr, @TrnDate, @TrnType, @Trn, @Bal, @TrnType, @OthCr, @OthDr       
-- Add to Excess if Transaction for Loans is not posted       
--   IF @UpdLN = 0       
--      SET @Excess = @SavTrnAmt - (@PrinCr - @PrinDr + @IntCr - @IntDr + @OthCr - @OthDr)       
--    IF @UpdLN = 1       
--    BEGIN       

   UPDATE lnMaster       
      SET Prin       = isnull(@Prin,0),       
          IntR       = isnull(@IntR + Discounted,0),             
          Oth        = isnull(@Oth,0),          
          WaivedInt  = isnull(@PaidWaived,0),          
          doLastTrn  = @TrnDate,       
--             LastTrnType= @TrnType,       
          LastTrn    = IsNull(@Trn,LastTrn),       
          WeeksPaid  = @dNum - CASE WHEN (@Bal - (@iPrinEnd + @iIntEnd + @iOthEnd)) = 0 THEN 0 ELSE 1 END,       
          Status     = CASE WHEN @Bal = 0 THEN CASE WHEN @TrnType = 3899 THEN 98 ELSE 99  END        
                            ELSE CASE WHEN Status in (30,91) THEN Status ELSE 30 END END                        
      WHERE Acc      = @Acc       
--    END       

  DECLARE @UpSaf as Bit       
  SET @UpSaf = 0       

-- Update Tellers Cash       

  IF @TrnType in (3001,3097,3098,3099)         
  BEGIN       
    UPDATE SAF SET Cash_On_Hand = Cash_On_Hand + @TrnAmt        
    WHERE TlrName = @PostBy       
    SET @UpSaf = 1       
    END        

    SET @Excess = @SavTrnAmt-@TrnAmt       
  IF Round(@SavTrnAmt,2) <>  Round(@TrnAmt,2)       

  BEGIN       
    DECLARE @sAcc as VarChar(17)       
    SELECT @sAcc = Acc from saMaster       
    WHERE CID = @CID and AcctType = 40 and Status in (10,20,90)     --(addded by kent)excess payment for closed account will be posted on AP Account

    IF @sAcc is Null       
      SET @sAcc = 'P'+rTrim(Convert(VarChar(10),@CID)) + '-40'       
      IF @TrnType  = 3899 -- From Renewals       
      BEGIN       
        SET @Particulars = 'Excess ' + @Acc       
 --print 'Insert savings 1'   
        EXEC usp_UpdSaTran @sAcc, 3, 903, @Excess, 0, @ORNo, @PostBy, @TermID, @Particulars, 0, 'A', @UpSaf,1         
        END 
      ELSE BEGIN        
-- Add to Excess if Transaction for Loans is not posted       
--select @PrinCr , @PrinDr , @IntCr , @IntDr , @OthCr , @OthDr       
--      SET @Excess = @SavTrnAmt - @TrnAmt --(@PrinCr - @PrinDr + @IntCr - @IntDr + @OthCr - @OthDr)       
        IF not @TrnType  in (3201,3202)       
        BEGIN       
--print 'Insert savings 2'   
          EXEC usp_UpdSaTran @sAcc, 3, 7012, @Excess, 0, @ORNo, @PostBy, @TermID, @Particulars, 0, 'A', @UpSaf,1       
          END 
        ELSE       
--print 'Insert savings 3'   
          EXEC usp_UpdSaTran @sAcc, 231, 15, @Excess, 0, @ORNo, @PostBy, @TermID, @Particulars, 0, 'A', @UpSaf,1       
        END       
      END       

-- select * from trntypes where apptype = 0 order by 1       
-- Remarks to make process faster       

/*       

-- Fix TrnMaster       
--DEClare @acc  as varchar(17), @sAcc as varchar(17), @TrnAmt as Numeric(14,2), @Bal as Numeric(14,2)       
--set @sAcc = ''       
  DECLARE lnTrn CURSOR       
  KEYSET       
  FOR        
  SELECT 
    Acc, Case When lnpricr-lnpridr+lnintcr-lnintdr+CumAccrdInt-       
      ACCRDINT+lnOthcr-lnOthdr = 0 then TrnAmt        
            Else lnpricr-lnpridr+lnintcr-lnintdr+CumAccrdInt-ACCRDINT+lnOthcr-lnOthdr end       
    FROM trnmaster       
    WHERE Acc = @Acc AND trnType in (3001,3097,3098,3099,3899,3201,3202)        
  and TrnDate >= @TrnDate and Trn >= @Trn       
    ORDER BY Acc,trnDate,trn       
  -- Remarks and remove the Previous Tran in Checking       
  --   SET @Bal = @PrinRel + @IntRel + @OthRel       
  OPEN lnTrn       
  FETCH NEXT FROM lnTrn INTO @Acc, @TrnAmt       
  WHILE (@@fetch_status <> -1)       
  BEGIN       
  IF (@@fetch_status <> -2)       
  BEGIN       
  --        IF @Acc <> @sAcc        
  --        BEGIN       
  --           SELECT @Bal = Principal + Interest - Discounted from lnMaster Where acc = @Acc      
  --           Set @sAcc = @Acc       
  --        END       
  SET @Bal = @Bal - @TrnAmt       
          UPDATE trnMaster       
                SET Balance = @Bal, AvlBal = @Bal       
    WHERE CURRENT OF lnTrn       
      END       
    FETCH NEXT FROM lnTrn INTO @Acc, @TrnAmt       
  END       
  CLOSE lnTrn       
  DEALLOCATE lnTrn       
*/       

  IF @Bal = 0       
    UPDATE lnMaster SET Status = 99 WHERE Acc = @Acc and Status not in (99 ,98)       
 
  IF @Bal > 0 and @Particulars = 'Full Payment' and @TrnType = 3099  
    exec fixloantran @ACC,0,0 

  --- Fixed special Loan   
  IF @AcctType = 330 exec usp_FixSpecialLoan @Acc
