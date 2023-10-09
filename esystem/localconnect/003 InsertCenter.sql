
 -- Insert Unit that are not map
INSERT INTO Office(
  code, short_name, office_name, date_stablished, type_id, parent_id, alternate_id, 
  address_detail, address_url, geography_id, cid_sequence)
SELECT
  u.BrCode || '-' || u.unitcode code, u.BrCode || '-' || u.unitcode short_name, u.unit office_name, null date_stablished, 
  typ.ID Type_id, bo.ID parent_id, u.BrCode || '-' || u.unitcode alternate_id, 
  u.unitaddress address_detail, null address_url, null geography_id, 0 cid_sequence
FROM staging.unit u
LEFT JOIN vwoffice o on o.code = u.brcode || '-' || u.unitcode 
INNER JOIN vwReference typ on typ.Title = 'Unit' and typ.ref_type = 'OfficeType'
LEFT JOIN vwOffice bo on bo.officetype  = 'Area' and lower(bo.code) = lower(u.brcode)
WHERE o.id is null
ON CONFLICT (COALESCE(Parent_ID,0), lower(Code)) DO UPDATE SET
  code = EXCLUDED.code,
  short_name = EXCLUDED.short_name,
  office_name = EXCLUDED.office_name,
  date_stablished = EXCLUDED.date_stablished,
  type_id = EXCLUDED.type_id,
  parent_id = EXCLUDED.parent_id,
  alternate_id = EXCLUDED.alternate_id,
  address_detail = EXCLUDED.address_detail,
  address_url = EXCLUDED.address_url,
  geography_id = EXCLUDED.geography_id,
  cid_sequence = EXCLUDED.cid_sequence;
    
-- INSERT Center to Customer_Group
INSERT INTO Customer_Group(
  central_office_id, code, type_id, group_name, short_name, 
  date_stablished, meeting_day, office_id, officer_id, parent_id, 
  alternate_id, address_detail, address_url, geography_id
)
SELECT 
  co.id central_office_id, c.code, typ.ID type_id, c.centername group_name, null short_name, 
  c.dateestablished date_stablished, c.meetingday meeting_day, o.ID office_id, e.ID officer_id, null parent_id, 
  c.code alternate_id, c.centeraddress address_detail, null address_url, null geography_id
FROM 
 (SELECT BrCode, BrCode||'-'||unit unitcode, BrCode||'-'||centerCode code, centername, centeraddress, meetingday, unit, dateestablished, case when aoid = -1 then 0 else aoid end aoid
  FROM staging.center) c 
LEFT JOIN vwOffice co on co.Code = 'CI'
LEFT JOIN vwReference typ on typ.ref_type  = 'CustomerGroupType' and typ.Title = 'Center'
LEFT JOIN vwoffice br on lower(br.Code) = lower(c.BrCode) and br.officetype = 'Area'
LEFT JOIN vwoffice o on COALESCE(o.parent_id,0) = br.ID and lower(c.unitcode) = lower(o.Code) and o.officetype = 'Unit'
LEFT JOIN staging.aomap a on a.brcode = c.brcode and  a.aoid = c.aoid
LEFT JOIN employee e on e.employee_no = a.empno 
ON CONFLICT (Type_ID, Central_Office_ID, Code) DO UPDATE SET
  group_name = EXCLUDED.group_name,
  short_name = EXCLUDED.short_name,
  date_stablished = EXCLUDED.date_stablished,
  meeting_day = EXCLUDED.meeting_day,
  office_id = EXCLUDED.office_id,
  officer_id = EXCLUDED.officer_id,
  parent_id = EXCLUDED.parent_id,
  alternate_id = EXCLUDED.alternate_id,
  address_detail = EXCLUDED.address_detail,
  address_url = EXCLUDED.address_url,
  geography_id = EXCLUDED.geography_id,
  other_info = EXCLUDED.other_info;



