insert into customer_specs_number(customer_id, specs_code, specs_id, value, value2, measure_id)
  SELECT c.ID customer_id, specs.short_name specs_code, specs.id specs_id, 1000 "value",0 value2, measure.id measure_id
  FROM Customer c, Reference specs, Reference measure
  WHERE c.id = 4798 and specs.ref_type = 'CustomerSpecs' and specs.title = 'ShareofStocks'
  and measure.ref_type = 'Currency' and measure.Title = 'Philippine peso'

INSERT INTO IDs(
  iiid, series, id_number, registration_date, validity_date, type_id )
SELECT cus.iiid, COALESCE(id.series,0) + 1, '143' id_number, null registration_date, null validity_date, typ.id type_id
FROM Customer cus
INNER JOIN reference typ on lower(typ.ref_type) = 'idtype' and lower(typ.title) = 'inai-iiid'
LEFT JOIN 
  (SELECT iiid, Max(Series)  series
   FROM IDs
   GROUP BY iiid) id on id.iiid = cus.iiid
WHERE cus.id = 4798
ON CONFLICT (iiid, type_id, lower(trim(id_number))) DO UPDATE SET
  registration_date = EXCLUDED.registration_date, 
  validity_date = EXCLUDED.validity_date



INSERT INTO IDs(
  iiid, series, id_number, registration_date, validity_date, type_id )
SELECT cus.iiid, COALESCE(id.series,0) + 1, '1435254' id_number, null registration_date, null validity_date, typ.id type_id
FROM Customer cus
INNER JOIN reference typ on lower(typ.ref_type) = 'idtype' and lower(typ.title) = 'inai-iiid'
LEFT JOIN 
  (SELECT iiid, Max(Series)  series
   FROM IDs
   GROUP BY iiid) id on id.iiid = cus.iiid
WHERE cus.id = 22185
ON CONFLICT (iiid, type_id, lower(trim(id_number))) DO UPDATE SET
  registration_date = EXCLUDED.registration_date, 
  validity_date = EXCLUDED.validity_date


-- INSERT INTO Contact(
--   iiid, series, Contact, Type_Id)
-- SELECT cus.iiid, COALESCE(id.series,0) + 1 Series, '09988513220' Contact, typ.id type_id
-- FROM Customer cus
-- INNER JOIN reference typ on lower(typ.ref_type) = 'contacttype' and lower(typ.title) = 'cellphone'
-- LEFT JOIN 
--   (SELECT iiid, Max(Series)  series
--    FROM Contact
--    GROUP BY iiid) id on id.iiid = cus.iiid
-- WHERE cus.id = 4798
-- ON CONFLICT (iiid, type_id, lower(trim(Contact))) DO NOTHING 


