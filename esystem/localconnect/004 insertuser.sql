CREATE TABLE IF NOT EXISTS staging.UserMap(
    EmpNo VarChar(200) NOT NULL, 
    BrCode VarChar(2) NOT NULL, 
    LoginName VarChar(200) NULL
);
CREATE UNIQUE INDEX IF NOT EXISTS idxUserMap ON staging.UserMap(BrCode,LoginName);


INSERT INTO Staging.UserMap(EmpNo, BrCode, LoginName)
SELECT EmpNo, BrCode, LoginName
FROM (Values
('200212-00238','E3','magonos'),
('200302-00247','E3','mblasco'),
('200307-00352','E3','apanopio'),
('200601-00565','E3','ebagamasbad'),
('200706-01080','E3','lmacasaet'),
('200712-01365','E3','mpiano'),
('200804-01604','E3','ebaleciado'),
('200804-01644','E3','cabangan'),
('200906-02420','E3','mbejeno'),
('201008-03172','E3','apaloyo'),
('201008-16117','E3','jasamson'),
('201304-04913','E3','itneseo'),
('201405-05803','E3','itfarciosa'),
('201407-06046','E3','itjrejuso'),
('201408-06182','E3','itmjuan'),
('201410-06461','E3','jvera'),
('201503-06933','E3','itrnening'),
('201511-08579','E3','rbasilan'),
('201511-08638','E3','aballeta'),
('201512-08710','E3','avega'),
('201512-08760','E3','mrabin'),
('201512-08808','E3','dgbarbosa'),
('201607-11525','E3','itasalem'),
('201703-12911','E3','itgcdomingo'),
('201707-15115','E3','itrbernal'),
('201709-15302','E3','itlpasiliao'),
('201709-16046','E3','itncamigla'),
('201801-20707','E3','cmguab'),
('201806-20794','E3','itmantero'),
('201905-28117','E3','rtaduran'),
('202205-39209','E3','itjnitollano'),
('a81fa2bc-89a9-4d69-8a2b-bbe8052ece74','E3','itcaquino'),
('af97ae82-4d19-4b36-8a15-3b0c3efbc303','E3','ebeliciado'),
('9e72d324-a00d-4f7b-8686-7490331c51f6','E3','ibustarga'),
('10b91256-9e33-4b44-bbce-cb363a9817b5','E3','itmconseja'),
('f4bac085-b614-47e6-b0e6-dbc23a9f67ec','E3','mjdeguzman'),
('16288c72-b5a7-4f9d-bd8a-868122c4802c','E3','rdelarama'),
('f4bac085-b614-47e6-b0e6-dbc23a9f67ec','E3','mguzman'),
('553c9709-76e1-49e7-b86e-8e3c8cac9771','E3','itjestipona'),
('e4b7674f-a4a5-4fc6-b654-b388c3270c06','E3','cguab'),
('75e56aed-07d0-4210-b2cc-1434243026e6','E3','jhade'),
('6dbca81c-679a-45e5-b967-1bc9348aaf3f','E3','ngorit'),
('56bf26f1-6005-4098-827b-d01b17a9a5f1','E3','ritngo'),
('64505c8e-1194-4e97-838c-5098721d08ea','E3','amabini'),
('ea815506-e72c-4f94-97c5-71d89e49fdfa','E3','mmangampo'),
('f108877a-812a-45f5-880f-d5b820af636e','E3','lmartin'),
('e229557c-28db-40db-9b07-ca314048c8e8','E3','kmendez'),
('b2261675-5629-4a8b-8f5a-438bd597cd4f','E3','emijares'),
('b2261675-5629-4a8b-8f5a-438bd597cd4f','E3','esmijares'),
('4ce5ce29-20ae-4586-b9f7-1602aefeb31d','E3','apelaez'),
('80a5e964-3b65-41ba-a94e-a601c62a7a98','E3','aprena'),
('5da4e0c3-67a7-49da-bb99-0dc80f6f11c7','E3','araro'),
('e4f8274a-2f02-4056-a770-b6d85d71ecb4','E3','srit'),
('94cc5c82-135c-457c-bd39-8c0eaee9144d','E3','SA'),
('c6b28c0b-4105-45b6-8124-5299c2af9b06','E3','jsamson'),
('5e142189-5b54-4365-a3f6-d2f99f1e5a05','E3','msanone'),
('defca4a7-6cc3-4818-9e0b-ddfeb8131d2e','E3','aserafines'),
('87549473-93bc-4328-a3e2-0846b4c2bbe1','E3','jsoria'),
('6842747a-c349-4ad4-9e09-f882fa0ead15','E3','itjrtabunan'),
('38166795-338b-4acd-a327-dcabc79d1b2b','E3','aamor'),
('d2766420-4074-46c4-96b5-d8f1cd6e25e8','E3','abuenaflor'),
('277c3534-095f-4fb8-9ccc-dfd1c67ff530','E3','adesengano'),
('134c755e-93a2-43dc-a84c-59ecc1ed04d8','E3','asalve'),
('2921cb24-6fc9-4bf3-ae7f-91f4a9a5303b','E3','BINDEFENSO'),
('92124aab-8d2d-4d45-9ae3-8376d29b1781','E3','cador'),
('bedaec1f-188f-4ac6-8d50-36647d431183','E3','DPARZA'),
('0218616b-a317-4d1b-ab6b-fcb9b311bee7','E3','egay'),
('3e3309c2-0b68-46e9-9ac1-4cffb6d0a860','E3','esalcedo'),
('c26bd8e5-7b67-4e2f-af05-5f36bfca0807','E3','gabanes'),
('b7e05952-59f2-4914-96c4-78f8e7dd4616','E3','gbalane'),
('27b4ed8e-3ce7-47e6-ba89-e7167456d3ce','E3','jalarcon'),
('5c763df9-dda9-4bc7-8641-a2e16e5a83d7','E3','jbueno'),
('cb9b0753-19dc-4918-82f9-47b0bbbdd863','E3','jmorales'),
('2225c70e-b302-415e-9702-b5767cf476c2','E3','jmrepaso'),
('aa5de502-8b5c-4d82-97a5-af6b43b2391d','E3','joy'),
('1619d15d-783b-40ce-b590-b5529f77cdbc','E3','jroyo'),
('4a048b63-7883-4ae9-9d5d-24d385a2274c','E3','jtindugan'),
('31e9bdba-0ea4-433d-bec6-8f040c5ffc1c','E3','kcobilla'),
('91730ed5-f38e-4ce6-bedd-40c912fced13','E3','knunez'),
('18f47688-e463-44ef-9768-3b407a44d8b2','E3','ldimaculangan'),
('dcdd7f3c-87ec-4296-ba0f-27659249ae5e','E3','lgasga'),
('37dff958-2740-40b6-a1c4-a44f36c378d7','E3','lmadrona'),
('9b45ba1f-7059-4a3b-aaab-4e972ce9a500','E3','maynova'),
('230fce1f-010d-4d01-8df9-6d571db2e7ee','E3','mimperial'),
('4cb552df-04f5-42d4-b520-21a08bba1bb2','E3','mjparcia'),
('23fac2a7-fb48-4b76-b78d-0c3f47a687b3','E3','mnova'),
('fc13fbff-8c6d-49a6-a6a6-920b28eba4e2','E3','msamonte'),
('a15fe9d5-5eb3-4b96-8d2d-f7e54d859712','E3','nnoble'),
('657a568a-6b70-46f1-9c5d-a4c60be77530','E3','phomnes'),
('6732f477-3139-46f7-a015-662bdf7e7147','E3','pmarquez'),
('4ffda815-fed5-44d7-aa59-910c76ccf079','E3','rarcayera'),
('1f0c237d-4943-4e97-974f-8ae6b2d9bd0f','E3','rgcapisonda'),
('b287c273-73d9-414e-bfea-9907cdc5d82d','E3','rjudane'),
('2badd313-3a4f-447d-a507-b945ce79e362','E3','sbnacario'),
('0092d23b-7be3-42e5-8a0f-c72864839717','E3','sflores'),
('320952ee-6a2b-49c1-815e-174683c148d9','E3','sguerrero'),
('35198ed6-3ba9-4ed3-9428-040be8dedafe','E3','shiela'),
('fa96da3d-885f-4680-a553-2c7ea885815f','E3','snacario'),
('73e318af-d642-4ee8-a10e-adaf1a485db6','E3','spalano'),
('2604fc96-352a-4543-aa54-85e0a7f6fee3','E3','spinon'),
('79f3eb98-ae2f-42c8-bc56-5ee1687a3408','E3','svergara'),
('e5061056-6f46-4135-aa6b-27fda6da8590','E3','vdeguzman'),
('abe775ee-147e-4388-8d6a-249e505b4dea','E3','virgilio'),
('8178cc4d-214c-483b-bbc4-9c32dfd8f7eb','E3','VVIOLANDA'),
('12649fef-0732-4719-8f8f-470c609a22ae','E3','wbosch'),
('7cf096f8-f01d-4b66-842d-f83d7067be40','E3','salano'),
('cb2bedde-f2b5-460c-83dc-f2d0883d7cc6','E3','cvgeqgmniz'),
('7d480a2c-3c23-43b5-a171-74c1ca4032d8','E3','ejacyxorxb'),
('f0d87013-55e9-46ce-a90f-26845a82cbfa','E3','kcgellgqll'),
('227f497e-c36e-4bbc-afde-23d4bd4a9665','E3','vgepubozfk'),
) e(EmpNo, BrCode, LoginName)
ON CONFLICT (BrCode, LoginName) DO UPDATE SET
  EmpNo = Excluded.EmpNo;


INSERT INTO Users(
  iiid, login_name, display_name, access_role_id, status_code, 
  date_given, date_expired, date_locked, password_changed_at, hashed_password, 
  attempt, isloggedin
)
SELECT 
    ii.Id iiid, u.BrCode || '-' || u.userid login_name, FullNameTFMLS(ii.Title, ii.Last_Name, ii.First_Name, ii.Middle_Name, ii.Suffix_Name), 
    r.Id access_role_id, stat.code status_code, 
    u.dategiven date_given, u.dateexpired date_expired, u.datelocked date_locked, NULL password_changed_at, E'aa\\000aa'::bytea hashed_password, 
    u.attempt attempt, (u.isloggedin = 1) isloggedin
  
FROM Staging.userslist u 
LEFT JOIN Staging.UserMap e on e.BrCode = u.BrCode and e.LoginName = u.UserId
LEFT JOIN identity_info ii on ii.alternate_id = e.empno 
LEFT JOIN Access_Role r on lower(r.access_name) = lower(u.position)
LEFT JOIN reference stat on lower(stat.ref_type) = 'userstatus' and lower(stat.title) = 'active'
ON CONFLICT (lower(login_name)) DO UPDATE SET
iiid =  EXCLUDED.iiid,
display_name =  EXCLUDED.display_name,
access_role_id =  EXCLUDED.access_role_id,
status_code =  EXCLUDED.status_code,
date_given =  EXCLUDED.date_given,
date_expired =  EXCLUDED.date_expired,
date_locked =  EXCLUDED.date_locked,
password_changed_at =  EXCLUDED.password_changed_at,
hashed_password =  EXCLUDED.hashed_password,
attempt =  EXCLUDED.attempt,
isloggedin =  EXCLUDED.isloggedin;



INSERT INTO Users(
  iiid, login_name, display_name, access_role_id, status_code, 
  date_given, date_expired, date_locked, password_changed_at, hashed_password, 
  attempt, isloggedin
)
SELECT 
    ii.Id iiid, e.BrCode || '-' || e.loginname, FullNameTFMLS(ii.Title, ii.Last_Name, ii.First_Name, ii.Middle_Name, ii.Suffix_Name), 
    r.Id access_role_id, stat.code status_code, 
    '1900/01/01' date_given, '1900/01/01' date_expired, '1900/01/01' date_locked, NULL password_changed_at, E'aa\\000aa'::bytea hashed_password, 
    0 attempt, false isloggedin
  
FROM Staging.UserMap e 
LEFT JOIN identity_info ii on ii.alternate_id = e.empno 
LEFT JOIN Access_Role r on lower(r.access_name) = lower('none')
LEFT JOIN reference stat on lower(stat.ref_type) = 'userstatus' and lower(stat.title) = 'deleted'
ON CONFLICT (lower(login_name)) DO UPDATE SET
iiid =  EXCLUDED.iiid,
display_name =  EXCLUDED.display_name,
access_role_id =  EXCLUDED.access_role_id,
status_code =  EXCLUDED.status_code,
date_given =  EXCLUDED.date_given,
date_expired =  EXCLUDED.date_expired,
date_locked =  EXCLUDED.date_locked,
password_changed_at =  EXCLUDED.password_changed_at,
hashed_password =  EXCLUDED.hashed_password,
attempt =  EXCLUDED.attempt,
isloggedin =  EXCLUDED.isloggedin;
