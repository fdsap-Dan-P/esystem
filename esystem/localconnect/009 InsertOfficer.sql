INSERT INTO Officer (
  Office_Id, officer_iiid, employee_id, is_Head, position, period_start, period_end, status_id)
SELECT 
  ofc.id Office_Id, emp.iiid officer_iiid, emp.id employee_id, true is_Head, sf.position, 
  sf.begDate::date period_start, sf.endDate::date period_end, stat.id status_id 
FROM
   (VALUES 
('49eb057e-7b97-45a4-bf1b-dd04ecbd97a1','201511-08638','Unit Manager','2020/01/01','2099/01/01'),
('4ddbe818-1cb7-4019-b9cd-b5e7d6d3011c','200601-00565','Unit Manager','2020/01/01','2099/01/01'),
('db505498-64f0-4537-b43f-b7291e98789d','201511-08579','Unit Manager','2020/01/01','2099/01/01'),
('07cdf596-b452-431d-9a40-93f316b18793','200804-01645','Unit Manager','2020/01/01','2099/01/01'),
('f402eecc-b383-4326-9a82-12df8b6ac380','200804-01644','Unit Manager','2020/01/01','2099/01/01'),
('8ac1e0f2-d088-457e-845c-53ed268019e5','200712-01365','Unit Manager','2020/01/01','2099/01/01'),
('7795eecc-916e-4729-b09d-1b6d9d30aa03','201407-06086','Unit Manager','2020/01/01','2099/01/01')
)
sf (OfficeUIID, StaffNo, Position, begDate, endDate)
INNER JOIN Office ofc on sf.OfficeUIID::UUID = ofc.uuid
INNER JOIN Employee emp on sf.StaffNo = emp.employee_no 
INNER JOIN Reference stat on lower(stat.title) = 'active' and lower(ref_type) = 'officerstatus'

with a as 
(SELECT *
FROM Officer
)
update office set officer_iiid = a.officer_iiid
from a 
where a.office_id = office.id
