-- INSERT Identity Info
INSERT INTO identity_info(
    identity_map_id, isperson, alternate_id, title, last_name,
    first_name, middle_name, mother_maiden_name, suffix_name, professional_suffixes,
    birthday, sex, gender_id, civil_status_id, birth_place_id,
    contact_id, phone, email
  )

SELECT
  null identity_map_id, true isperson, c.BrCode || '-' || c.CID alternate_id, title.title title, COALESCE(Lname,'') last_name,
   COALESCE(Fname,'') first_name,  COALESCE(Mname,'') middle_name, 
  CASE WHEN MaidenFName is null THEN '' ELSE  MaidenFName || ' ' END  ||
  CASE WHEN MaidenLName is null THEN '' ELSE  MaidenLName || ' ' END ||
  CASE WHEN MaidenMName is null THEN '' ELSE  MaidenMName  END mother_maiden_name, null suffix_name, null professional_suffixes,
  c.birthdate birthday, CASE WHEN c.sex = 'F' THEN false ELSE true END sex, null gender_id, stat.ID civil_status_id, null birth_place_id,
  null contact_id, null phone, null email
FROM staging.Customer c
LEFT JOIN
  (VALUES
    (1,'Company Name'),
    (2,'Mr.'),
    (3,'Mrs.'),
    (4,'Ms.'),
    (5,'Atty.'),
    (6,'Engr.'),
    (7,'Hon.'),
    (8,'Dr.')
   ) title(code, title) on title.code = c.Title
   
LEFT JOIN
  (VALUES
    (1,'Single'),
    (2,'Married'),
    (3,'Spinster'),
    (4,'Widow'),
    (0,'No Status'),
    (5,'Live-in'),
    (6,'Separated')
   ) civil(code, civil) on civil.code = c.CivilStatus
LEFT JOIN vwReference stat on stat.ref_type  = 'CivilStatus' and lower(stat.Title) = lower(civil.civil)
ON CONFLICT(UUID) DO UPDATE SET
  identity_map_id = EXCLUDED.identity_map_id,
  isperson = EXCLUDED.isperson,
  alternate_id = EXCLUDED.alternate_id,
  title = EXCLUDED.title,
  last_name = EXCLUDED.last_name,
  first_name = EXCLUDED.first_name,
  middle_name = EXCLUDED.middle_name,
  mother_maiden_name = EXCLUDED.mother_maiden_name,
  suffix_name = EXCLUDED.suffix_name,
  professional_suffixes = EXCLUDED.professional_suffixes,
  birthday = EXCLUDED.birthday,
  sex = EXCLUDED.sex,
  gender_id = EXCLUDED.gender_id,
  civil_status_id = EXCLUDED.civil_status_id,
  birth_place_id = EXCLUDED.birth_place_id,
  contact_id = EXCLUDED.contact_id,
  phone = EXCLUDED.phone,
  email = EXCLUDED.email,
  simple_name = EXCLUDED.simple_name,
  vec_simple_name = EXCLUDED.vec_simple_name,
  vec_full_simple_name = EXCLUDED.vec_full_simple_name,
  other_info = EXCLUDED.other_info;  
  

-- Insert Customer
INSERT INTO Customer(
  IIID, central_office_id, cid, customer_alt_id, debit_limit, credit_limit, 
  date_entry, last_activity_date, dosri, classification_id, sub_classification_id, 
  customer_group_id, office_id, restriction_id, risk_class_id, industry_id, 
  status_code 
)

SELECT 
    ii.ID IIID, co.ID central_office_id, ii.ID cid, customer_alt_id, 150000 debit_limit, 150000 credit_limit, 
    c.date_entry, null last_activity_date, false dosri, cls.ID classification_id, 
    subcls.ID sub_classification_id, cg.ID customer_group_id, o.ID office_id, null restriction_id, null risk_class_id, null industry_id,
    stat.Code status_Code
FROM
 (SELECT 
    c.cid, c.BrCode, c.BrCode || '-' || c.CID customer_alt_id,  
    c.dateentry date_entry, CASE WHEN Classification = 0 THEN 'Member' ELSE 'Non-Member' END classification, 
    CASE WHEN COALESCE(SubClassification,0) = 0 THEN  1564 ELSE SubClassification END sub_classification_id,
    CenterCode, CASE Status WHEN 0 THEN 'Active' WHEN 1 THEN CASE WHEN i.cid is null THEN 'Active' ELSE 'InActive' END 
    WHEN 2 THEN 'Resigned' WHEN 4 THEN 'Balik CARD' ELSE '' END Status
  FROM staging.Customer c
  LEFT JOIN staging.inactivecid i on c.brcode = i.brcode and c.cid = i.cid
) c 
LEFT JOIN identity_info ii on ii.alternate_id = c.BrCode || '-' || c.CID 
LEFT JOIN vwOffice co on co.Code = 'CI'
LEFT JOIN vwReference cls on cls.ref_type = 'CustomerClass' and c.classification = cls.Title 
LEFT JOIN vwReference subcls on subcls.ref_type = 'CustomerSubClass' and c.sub_classification_id = subcls.code 
LEFT JOIN customer_group cg  on cg.alternate_id = c.BrCode || '-' || c.CenterCode
LEFT JOIN vwoffice o  on o.officetype  = 'Area' and o.Code = c.BrCode
LEFT JOIN vwReference stat on stat.ref_type = 'CustomerStatus' and c.Status = stat.Title;
 