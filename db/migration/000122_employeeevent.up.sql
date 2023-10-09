----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.Employee_Event (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  Employee_ID bigint NOT NULL,
  Ticket_ID bigint NOT NULL,
  Event_Type_ID bigint NOT NULL,
  Office_ID bigint NOT NULL,
  Position_ID bigint NOT NULL,
  Basic_Pay numeric NOT NULL,
  Status_ID bigint NOT NULL,
  Job_Grade int2 NOT NULL,
  Job_Step int2 NOT NULL,
  Level_ID bigint NULL,
  Employee_Type_ID bigint NOT NULL,
  Remarks VarChar(1000) NULL,
  Other_Info jsonb NULL,

  CONSTRAINT Employee_Event_pkey PRIMARY KEY (UUID),
  CONSTRAINT fk_Employee_Event_Emp FOREIGN KEY (Employee_ID) REFERENCES Employee(ID),
  CONSTRAINT fk_Employee_Event_Ticket FOREIGN KEY (Ticket_ID) REFERENCES Ticket(ID),
  CONSTRAINT fk_Employee_Event_Level FOREIGN KEY (Level_ID) REFERENCES Reference(ID),
  CONSTRAINT fk_Employee_Event_Office FOREIGN KEY (Office_ID ) REFERENCES Office(ID),
  CONSTRAINT fk_Employee_Event_Position FOREIGN KEY (Position_ID) REFERENCES Reference(ID),
  CONSTRAINT fk_Employee_Event_Status FOREIGN KEY (Status_ID ) REFERENCES Reference(ID),
  CONSTRAINT fk_Employee_Event_Event_Type FOREIGN KEY (Event_Type_ID) REFERENCES Reference(ID),
  CONSTRAINT fk_Employee_Event_EmpType FOREIGN KEY (Employee_Type_ID) REFERENCES Reference(ID)
);

DROP TRIGGER IF EXISTS trgEmployee_Event_Ins on Employee_Event;
---------------------------------------------------------------------------
CREATE TRIGGER trgEmployee_Event_Ins
---------------------------------------------------------------------------
    BEFORE INSERT ON Employee_Event
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();

DROP TRIGGER IF EXISTS trgEmployee_Event_upd on Employee_Event;
---------------------------------------------------------------------------
CREATE TRIGGER trgEmployee_Event_upd
---------------------------------------------------------------------------
    BEFORE UPDATE ON Employee_Event
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgEmployee_Event_del on Employee_Event;
---------------------------------------------------------------------------
CREATE TRIGGER trgEmployee_Event_del
---------------------------------------------------------------------------
    AFTER DELETE ON Employee_Event
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();

  
----------------------------------------------------------------------------------------
CREATE OR REPLACE VIEW public.vwEmployee_Event
----------------------------------------------------------------------------------------
AS SELECT 
    evt.UUID, evt.Employee_ID,
 
    ofc.ID Office_ID , ofc.Code OfficeCode, 
    ofc.Short_Name OfficeShort_Name, ofc.Office_Name Office_Name,
 
    pos.ID  Position_ID,     pos.UUID  PositionUUID,     pos.Title Position_desc,     
    sta.ID  Status_ID ,       sta.UUID  StatusUUID,       sta.Title Status,    
    typ.ID  Employee_Type_ID, typ.UUID  Employee_TypeUUID, typ.Title Employee_Type,    
    vtyp.ID Event_Type_ID,    vtyp.UUID Event_TypeUUID,    vtyp.Title Event_Type,    
    lvl.ID  Level_ID,        lvl.UUID  LevelUUID,        lvl.Title LevelDesc,    

    mr.Mod_Ctr,
    evt.Other_Info,
    mr.Created,
    mr.Updated 
    
   FROM Employee_Event evt 
   INNER JOIN Main_Record mr  ON mr.UUID = evt.UUID
   LEFT JOIN  Reference pos  ON pos.ID  = evt.Position_ID   
   LEFT JOIN  Reference sta  ON sta.ID  = evt.Status_ID    
   LEFT JOIN  Reference typ  ON typ.ID  = evt.Employee_Type_ID   
   LEFT JOIN  Reference vtyp ON vtyp.ID = evt.Event_Type_ID   
   LEFT JOIN  Reference lvl  ON lvl.ID  = evt.Level_ID   
   LEFT JOIN  Office ofc     ON ofc.ID  = evt.Office_ID  
   ;

  INSERT INTO Employee_Event(
      UUID, Employee_ID, Ticket_ID, Event_Type_ID, Office_ID , Position_ID, Basic_Pay, Status_ID , 
      Job_Grade, Job_Step, Level_ID, Employee_Type_ID, Remarks)
  
  SELECT 
      a.UUID, Emp.ID Employee_ID, t.ID Ticket_ID, v.ID Event_Type_ID, Ofc.ID Office_ID , Title.ID Position_ID, a.Basic_Pay, Stat.ID Status_ID , 
      a.Job_Grade, a.Job_Step, Lvl.ID Level_ID, Typ.ID Employee_Type_ID, a.Remarks
     
  FROM (Values
      ('e757ae5d-d0c3-46f0-bc03-6efc92a98803'::UUID,'97-0114','da970ce4-dc2f-44af-b1a8-49a987148922'::UUID,'Regularization','10002','IT Officer',10000,'Regular',4,3,'Officer','Employed','Primary Income')
      )   
    a(UUID, Employee_No, TicketUUID, Event_Type, OfficeAltID, PositionDesc, Basic_Pay, Status, 
      Job_Grade, Job_Step, EmpLevel, Employee_Type, Remarks
      )  
  LEFT JOIN Employee Emp on emp.Employee_No  = a.Employee_No
  LEFT JOIN Ticket t on t.UUID  = a.TicketUUID
  LEFT JOIN Office Ofc  on lower(Ofc.Alternate_ID) = lower(a.OfficeAltID)  
  LEFT JOIN vwReference v on lower(v.Title) = lower(a.Event_Type)     and lower(v.Ref_Type)  = 'employeeevent'
  LEFT JOIN vwReference Title on lower(Title.Title) = lower(a.PositionDesc)     and lower(Title.Ref_Type)  = 'position'
  LEFT JOIN vwReference Lvl   on lower(Lvl.Title)   = lower(a.EmpLevel)  and lower(Lvl.Ref_Type)  = 'employeelevel'
  LEFT JOIN vwReference Stat  on lower(Stat.Title)  = lower(a.Status)    and lower(Stat.Ref_Type) = 'employeestatus'
  LEFT JOIN vwReference Typ   on lower(Typ.Title)   = lower(a.Employee_Type)   and lower(Typ.Ref_Type) = 'employeetype'
  
  ON CONFLICT(UUID) DO UPDATE SET
    Employee_ID = EXCLUDED.Employee_ID,
    Ticket_ID = EXCLUDED.Ticket_ID, 
    Event_Type_ID = EXCLUDED.Event_Type_ID, 
    Office_ID  = EXCLUDED.Office_ID , 
    Position_ID = EXCLUDED.Position_ID,
    Basic_Pay = EXCLUDED.Basic_Pay, 
    Status_ID  = EXCLUDED.Status_ID , 
    Job_Grade = EXCLUDED.Job_Grade, 
    Job_Step = EXCLUDED.Job_Step, 
    Level_ID = EXCLUDED.Level_ID, 
    Employee_Type_ID = EXCLUDED.Employee_Type_ID, 
    Remarks = EXCLUDED.Remarks
   ;
