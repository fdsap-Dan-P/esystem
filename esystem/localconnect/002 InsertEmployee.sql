CREATE TABLE IF NOT EXISTS staging.Employee(
    UUID uuid NOT NULL,  
    EmpNo VarChar(200) NOT NULL, 
    Fname VarChar(200) NULL,
    Lname VarChar(200) NULL, 
    MName VarChar(200) NULL, 
    Suffix VarChar(200) NULL, 
    Position VarChar(200) NULL, 
    OfficeUUID uuid NULL
);

CREATE UNIQUE INDEX IF NOT EXISTS idxEmployeeUuid ON staging.Employee(UUID);
CREATE UNIQUE INDEX IF NOT EXISTS idxEmployeeEmpNo ON staging.Employee(EmpNo);

INSERT INTO staging.Employee ( 
  UUID, EmpNo, Fname,  Lname, MName, Suffix, 
  Position, OfficeUUID)
SELECT 
  UUID::UUID, EmpNo, Fname,  Lname, MName, Suffix, 
  Position, OfficeUUID::UUID
FROM  
 (VALUES  
('a81fa2bc-89a9-4d69-8a2b-bbe8052ece74','a81fa2bc-89a9-4d69-8a2b-bbe8052ece74','AQUINO','CARL NEALJUN','P','','IT Officer','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('af97ae82-4d19-4b36-8a15-3b0c3efbc303','af97ae82-4d19-4b36-8a15-3b0c3efbc303','Baleciado222','Eric','B','','Unit Manager','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('9e72d324-a00d-4f7b-8686-7490331c51f6','9e72d324-a00d-4f7b-8686-7490331c51f6','Bustarga','Irene','H','','Unit Manager','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('10b91256-9e33-4b44-bbce-cb363a9817b5','10b91256-9e33-4b44-bbce-cb363a9817b5','CONSEJA','MARCO ANTONIO','L','','IT Officer','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('f4bac085-b614-47e6-b0e6-dbc23a9f67ec','f4bac085-b614-47e6-b0e6-dbc23a9f67ec','De Guzman','Madonna Joy','O','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('16288c72-b5a7-4f9d-bd8a-868122c4802c','16288c72-b5a7-4f9d-bd8a-868122c4802c','De La Rama ','Rosalin','','','Unit Manager','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('553c9709-76e1-49e7-b86e-8e3c8cac9771','553c9709-76e1-49e7-b86e-8e3c8cac9771','ESTIPONA','JERICK','M','','IT Officer','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('e4b7674f-a4a5-4fc6-b654-b388c3270c06','e4b7674f-a4a5-4fc6-b654-b388c3270c06','Guab ','Christine Mae ','D','','Bookkeeper','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('75e56aed-07d0-4210-b2cc-1434243026e6','75e56aed-07d0-4210-b2cc-1434243026e6','Hade','Joan','B','','Unit Manager','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('6dbca81c-679a-45e5-b967-1bc9348aaf3f','6dbca81c-679a-45e5-b967-1bc9348aaf3f','It Officer','It Officer','I','','IT Officer','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('56bf26f1-6005-4098-827b-d01b17a9a5f1','56bf26f1-6005-4098-827b-d01b17a9a5f1','Itliaison','Cardinc','C','','IT Officer','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('64505c8e-1194-4e97-838c-5098721d08ea','64505c8e-1194-4e97-838c-5098721d08ea','Mabini','Adelfa','G','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('ea815506-e72c-4f94-97c5-71d89e49fdfa','ea815506-e72c-4f94-97c5-71d89e49fdfa','Mangampo','Myla','T','','Unit Manager','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('f108877a-812a-45f5-880f-d5b820af636e','f108877a-812a-45f5-880f-d5b820af636e','Martin','Lilibeth','B','','Bookkeeper','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('e229557c-28db-40db-9b07-ca314048c8e8','e229557c-28db-40db-9b07-ca314048c8e8','Mendez','Kristine','A','','Unit Manager','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('b2261675-5629-4a8b-8f5a-438bd597cd4f','b2261675-5629-4a8b-8f5a-438bd597cd4f','Mijares','Elvie','S','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('4ce5ce29-20ae-4586-b9f7-1602aefeb31d','4ce5ce29-20ae-4586-b9f7-1602aefeb31d','Pelaez','Anelyn ','B','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('80a5e964-3b65-41ba-a94e-a601c62a7a98','80a5e964-3b65-41ba-a94e-a601c62a7a98','Prena','Almira','M','','Area Manager','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('5da4e0c3-67a7-49da-bb99-0dc80f6f11c7','5da4e0c3-67a7-49da-bb99-0dc80f6f11c7','Raro','Aires','V','','Unit Manager','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('e4f8274a-2f02-4056-a770-b6d85d71ecb4','e4f8274a-2f02-4056-a770-b6d85d71ecb4','rit','sme','','','Sysadmin','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('94cc5c82-135c-457c-bd39-8c0eaee9144d','94cc5c82-135c-457c-bd39-8c0eaee9144d','Sa','Sa','S','','Sysadmin','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('c6b28c0b-4105-45b6-8124-5299c2af9b06','c6b28c0b-4105-45b6-8124-5299c2af9b06','Samson','John Alfred','','','Unit Manager','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('5e142189-5b54-4365-a3f6-d2f99f1e5a05','5e142189-5b54-4365-a3f6-d2f99f1e5a05','Sanone','Mary Joy  ','V','','Unit Manager','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('defca4a7-6cc3-4818-9e0b-ddfeb8131d2e','defca4a7-6cc3-4818-9e0b-ddfeb8131d2e','Serafines','Alma','A','','Area Manager','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('87549473-93bc-4328-a3e2-0846b4c2bbe1','87549473-93bc-4328-a3e2-0846b4c2bbe1','Soria','Joy','M','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('6842747a-c349-4ad4-9e09-f882fa0ead15','6842747a-c349-4ad4-9e09-f882fa0ead15','TABUNAN','JAY R','U','','IT Officer','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('38166795-338b-4acd-a327-dcabc79d1b2b','38166795-338b-4acd-a327-dcabc79d1b2b','','aamor','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('d2766420-4074-46c4-96b5-d8f1cd6e25e8','d2766420-4074-46c4-96b5-d8f1cd6e25e8','','abuenaflor','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('277c3534-095f-4fb8-9ccc-dfd1c67ff530','277c3534-095f-4fb8-9ccc-dfd1c67ff530','','adesengano','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('134c755e-93a2-43dc-a84c-59ecc1ed04d8','134c755e-93a2-43dc-a84c-59ecc1ed04d8','','asalve','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('2921cb24-6fc9-4bf3-ae7f-91f4a9a5303b','2921cb24-6fc9-4bf3-ae7f-91f4a9a5303b','','BINDEFENSO','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('92124aab-8d2d-4d45-9ae3-8376d29b1781','92124aab-8d2d-4d45-9ae3-8376d29b1781','','cador','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('0ebb499c-b30f-4419-af0f-46f6c8a6365c','0ebb499c-b30f-4419-af0f-46f6c8a6365c','','cmguab','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('bc16625a-ff77-4a2c-b7a8-401aa1692057','bc16625a-ff77-4a2c-b7a8-401aa1692057','','dgbarbosa','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('bedaec1f-188f-4ac6-8d50-36647d431183','bedaec1f-188f-4ac6-8d50-36647d431183','','DPARZA','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('0218616b-a317-4d1b-ab6b-fcb9b311bee7','0218616b-a317-4d1b-ab6b-fcb9b311bee7','','egay','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('3e3309c2-0b68-46e9-9ac1-4cffb6d0a860','3e3309c2-0b68-46e9-9ac1-4cffb6d0a860','','esalcedo','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('c26bd8e5-7b67-4e2f-af05-5f36bfca0807','c26bd8e5-7b67-4e2f-af05-5f36bfca0807','','gabanes','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('b7e05952-59f2-4914-96c4-78f8e7dd4616','b7e05952-59f2-4914-96c4-78f8e7dd4616','','gbalane','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('27b4ed8e-3ce7-47e6-ba89-e7167456d3ce','27b4ed8e-3ce7-47e6-ba89-e7167456d3ce','','jalarcon','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('5c763df9-dda9-4bc7-8641-a2e16e5a83d7','5c763df9-dda9-4bc7-8641-a2e16e5a83d7','','jbueno','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('cb9b0753-19dc-4918-82f9-47b0bbbdd863','cb9b0753-19dc-4918-82f9-47b0bbbdd863','','jmorales','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('2225c70e-b302-415e-9702-b5767cf476c2','2225c70e-b302-415e-9702-b5767cf476c2','','jmrepaso','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('aa5de502-8b5c-4d82-97a5-af6b43b2391d','aa5de502-8b5c-4d82-97a5-af6b43b2391d','','joy','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('1619d15d-783b-40ce-b590-b5529f77cdbc','1619d15d-783b-40ce-b590-b5529f77cdbc','','jroyo','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('4a048b63-7883-4ae9-9d5d-24d385a2274c','4a048b63-7883-4ae9-9d5d-24d385a2274c','','jtindugan','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('31e9bdba-0ea4-433d-bec6-8f040c5ffc1c','31e9bdba-0ea4-433d-bec6-8f040c5ffc1c','','kcobilla','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('91730ed5-f38e-4ce6-bedd-40c912fced13','91730ed5-f38e-4ce6-bedd-40c912fced13','','knunez','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('18f47688-e463-44ef-9768-3b407a44d8b2','18f47688-e463-44ef-9768-3b407a44d8b2','','ldimaculangan','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('dcdd7f3c-87ec-4296-ba0f-27659249ae5e','dcdd7f3c-87ec-4296-ba0f-27659249ae5e','','lgasga','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('37dff958-2740-40b6-a1c4-a44f36c378d7','37dff958-2740-40b6-a1c4-a44f36c378d7','','lmadrona','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('9b45ba1f-7059-4a3b-aaab-4e972ce9a500','9b45ba1f-7059-4a3b-aaab-4e972ce9a500','','maynova','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('bba2f4a5-bffe-47ee-ba59-57536d3bb42e','bba2f4a5-bffe-47ee-ba59-57536d3bb42e','','MGUZMAN','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('230fce1f-010d-4d01-8df9-6d571db2e7ee','230fce1f-010d-4d01-8df9-6d571db2e7ee','','mimperial','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('4cb552df-04f5-42d4-b520-21a08bba1bb2','4cb552df-04f5-42d4-b520-21a08bba1bb2','','mjparcia','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('23fac2a7-fb48-4b76-b78d-0c3f47a687b3','23fac2a7-fb48-4b76-b78d-0c3f47a687b3','','mnova','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('fc13fbff-8c6d-49a6-a6a6-920b28eba4e2','fc13fbff-8c6d-49a6-a6a6-920b28eba4e2','','msamonte','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('a15fe9d5-5eb3-4b96-8d2d-f7e54d859712','a15fe9d5-5eb3-4b96-8d2d-f7e54d859712','','nnoble','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('657a568a-6b70-46f1-9c5d-a4c60be77530','657a568a-6b70-46f1-9c5d-a4c60be77530','','phomnes','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('6732f477-3139-46f7-a015-662bdf7e7147','6732f477-3139-46f7-a015-662bdf7e7147','','pmarquez','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('4ffda815-fed5-44d7-aa59-910c76ccf079','4ffda815-fed5-44d7-aa59-910c76ccf079','','rarcayera','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('1f0c237d-4943-4e97-974f-8ae6b2d9bd0f','1f0c237d-4943-4e97-974f-8ae6b2d9bd0f','','rgcapisonda','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('b287c273-73d9-414e-bfea-9907cdc5d82d','b287c273-73d9-414e-bfea-9907cdc5d82d','','rjudane','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('2badd313-3a4f-447d-a507-b945ce79e362','2badd313-3a4f-447d-a507-b945ce79e362','','sbnacario','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('0092d23b-7be3-42e5-8a0f-c72864839717','0092d23b-7be3-42e5-8a0f-c72864839717','','sflores','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('320952ee-6a2b-49c1-815e-174683c148d9','320952ee-6a2b-49c1-815e-174683c148d9','','sguerrero','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('35198ed6-3ba9-4ed3-9428-040be8dedafe','35198ed6-3ba9-4ed3-9428-040be8dedafe','','shiela','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('fa96da3d-885f-4680-a553-2c7ea885815f','fa96da3d-885f-4680-a553-2c7ea885815f','','snacario','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('73e318af-d642-4ee8-a10e-adaf1a485db6','73e318af-d642-4ee8-a10e-adaf1a485db6','','spalano','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('2604fc96-352a-4543-aa54-85e0a7f6fee3','2604fc96-352a-4543-aa54-85e0a7f6fee3','','spinon','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('79f3eb98-ae2f-42c8-bc56-5ee1687a3408','79f3eb98-ae2f-42c8-bc56-5ee1687a3408','','svergara','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('e5061056-6f46-4135-aa6b-27fda6da8590','e5061056-6f46-4135-aa6b-27fda6da8590','','vdeguzman','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('abe775ee-147e-4388-8d6a-249e505b4dea','abe775ee-147e-4388-8d6a-249e505b4dea','','virgilio','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('8178cc4d-214c-483b-bbc4-9c32dfd8f7eb','8178cc4d-214c-483b-bbc4-9c32dfd8f7eb','','VVIOLANDA','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('12649fef-0732-4719-8f8f-470c609a22ae','12649fef-0732-4719-8f8f-470c609a22ae','','wbosch','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('7cf096f8-f01d-4b66-842d-f83d7067be40','7cf096f8-f01d-4b66-842d-f83d7067be40','','salano','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('126596d6-2d13-47db-9bf8-74be28a76668','126596d6-2d13-47db-9bf8-74be28a76668','','rcapisonda','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('cb2bedde-f2b5-460c-83dc-f2d0883d7cc6','cb2bedde-f2b5-460c-83dc-f2d0883d7cc6','','cvgeqgmniz','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('7d480a2c-3c23-43b5-a171-74c1ca4032d8','7d480a2c-3c23-43b5-a171-74c1ca4032d8','','ejacyxorxb','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('f0d87013-55e9-46ce-a90f-26845a82cbfa','f0d87013-55e9-46ce-a90f-26845a82cbfa','','kcgellgqll','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda'),
('227f497e-c36e-4bbc-afde-23d4bd4a9665','227f497e-c36e-4bbc-afde-23d4bd4a9665','','vgepubozfk','','','Savings and Loans Teller','cfbb516e-1b7c-4f23-a7b4-a5fc4234ccda')
)
a (UUID, EmpNo, Fname,  Lname, MName, Suffix, 
  Position, OfficeUUID);
      
-- INSERT Identity Info
INSERT INTO identity_info(
    identity_map_id, isperson, alternate_id, title, last_name,
    first_name, middle_name, mother_maiden_name, suffix_name, professional_suffixes,
    birthday, sex, gender_id, civil_status_id, birth_place_id,
    contact_id, phone, email
  )
SELECT
  null identity_map_id, true isperson, EmpNo alternate_id, null title, Lname last_name,
  Fname first_name, MName middle_name, null mother_maiden_name, Suffix suffix_name, null professional_suffixes,
  null birthday, null sex, null gender_id, null civil_status_id, null birth_place_id,
  null contact_id, null phone, null email
FROM staging.Employee   
ON CONFLICT(alternate_id) DO UPDATE SET
  identity_map_id = EXCLUDED.identity_map_id,
  isperson = EXCLUDED.isperson,
--  alternate_id = EXCLUDED.alternate_id,
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
  

  update staging.employee 
  set "position" = 'Savings Teller'
  where "position"  = 'Savings and Loans Teller';


INSERT INTO Employee(
  UUID, iiid, central_office_id, employee_no, basic_pay, date_hired, 
  date_regular, job_grade, job_step, level_id, office_id, 
  position_id, status_code, superior_id, type_id
)
SELECT   
  e.UUID, ii.id, co.ID central_office_id, e.EmpNo employee_no, 0 basic_pay, null date_hired, 
  null date_regular, 0 job_grade, 0 job_step, null level_id, o.ID office_id, 
  p.ID position_id,  CASE WHEN Position(' oic' in lower(e.position)) > 0 THEN sta.code ELSE sta2.code END status_code, 
  NULL superior_id, typ.ID type_id
FROM   
  staging.Employee e
LEFT JOIN identity_info ii on ii.alternate_id = e.empno 
LEFT JOIN vwOffice co on co.Code = 'CI'
LEFT JOIN Office o on o.UUID = e.OfficeUUID
LEFT JOIN vwreference p on lower(p.ref_type) = 'position' and lower(p.Title) = trim(replace(replace(lower(e.position),' oic', ''),' -',''))
LEFT JOIN vwreference typ on lower(typ.ref_type) = 'employeetype' and lower(typ.Title) = 'employed'
LEFT JOIN vwreference sta on lower(sta.ref_type) = 'employeestatus' and sta.parent_id  = typ.id and lower(sta.Title) = 'oic' 
LEFT JOIN vwreference sta2 on lower(sta2.ref_type) = 'employeestatus' and lower(sta2.title) = 'active'
ON CONFLICT (UUID) DO UPDATE SET
  iiid =  EXCLUDED.iiid,
  central_office_id =  EXCLUDED.central_office_id,
  employee_no =  EXCLUDED.employee_no,
  basic_pay =  EXCLUDED.basic_pay,
  date_hired =  EXCLUDED.date_hired,
  date_regular =  EXCLUDED.date_regular,
  job_grade =  EXCLUDED.job_grade,
  job_step =  EXCLUDED.job_step,
  level_id =  EXCLUDED.level_id,
  office_id =  EXCLUDED.office_id,
  position_id =  EXCLUDED.position_id,
  status_code =  EXCLUDED.status_code,
  superior_id =  EXCLUDED.superior_id,
  type_id =  EXCLUDED.type_id,
  other_info =  EXCLUDED.other_info;


INSERT INTO staging.AOMap(EmpNo, BrCode, AOID)
select * from 
(select a.empno, a.BrCode, a.AOID from 
(values 
('200601-00565','E3',39),
('200712-01365','E3',52),
('200804-01644','E3',66),
('201309-05147','E3',10),
('201407-06085','E3',23),
('201408-06263','E3',28),
('201408-06267','E3',63),
('201410-06461','E3',35),
('201506-07540','E3',44),
('201510-08476','E3',20),
('201511-08577','E3',7),
('201511-08579','E3',30),
('201511-08638','E3',18),
('201608-10959','E3',1),
('201612-12469','E3',36),
('201706-13191','E3',15),
('201707-14095','E3',22),
('201707-14531','E3',42),
('201710-16800','E3',17),
('201801-20708','E3',24),
('201802-17763','E3',29),
('201804-18338','E3',50),
('201805-20101','E3',51),
('201807-20704','E3',55),
('201807-21831','E3',67),
('201808-21881','E3',26),
('201811-24500','E3',16),
('201903-26328','E3',57),
('201905-28117','E3',14),
('201905-28120','E3',4),
('201906-29039','E3',11),
('202102-35883','E3',2),
('202208-40126','E3',31),
('202208-40353','E3',60),
('202210-41718','E3',68),
('201801-17287','E3',65),
('202203-38617','E3',9),
('202208-40736','E3',54),
('202208-40798','E3',45),
('202209-41252','E3',46),
('201807-20705','E3',3),
('201707-14532','E3',5),
('201702-13011','E3',6),
('201410-06461','E3',8),
('201811-24569','E3',21),
('ed5bb1cf-79ac-48d0-8f29-49aab3d3c94a','E3',12),
('097e7677-4de9-44b9-be74-cb81127d5734','E3',13),
('201602-09155','E3',19),
('05494285-4614-40cd-a38d-e7c50c0266c6','E3',25),
('a4392c1e-e334-43b5-99ea-e6b02854dd34','E3',27),
('c7e43b5f-3c53-4bba-9112-7872e7e83206','E3',32),
('f690dcca-2360-4295-b117-5b084671f623','E3',33),
('6ba6091b-71a8-4b66-8875-e9ef1849bf33','E3',34),
('4c326a3a-20da-406e-a6f4-9a58dcc7bc06','E3',37),
('21762780-aa25-4d1c-b829-323d08f8f148','E3',38),
('c910b629-2163-4e3a-b936-6d19b13c7206','E3',40),
('11012eb8-4986-4dd9-bf07-0dede61aa8d6','E3',41),
('b2cae61c-a26d-4f75-9b36-fdd67c859c1d','E3',43),
('09d5665f-a698-4a38-acfd-6fd2c34016a5','E3',47),
('3a2ee21f-ca21-44cd-ad25-cb1c9134501a','E3',48),
('4ef0409f-f8bd-4177-8519-feebfadeb0e6','E3',49),
('1acc3f22-21c7-4cbc-9bb0-73e6b87e0ceb','E3',53),
('46bc30d9-f0d0-4be8-9e40-61f3a06fed65','E3',56),
('53b666a9-3dae-422d-af8e-224dfec46c82','E3',58),
('7a9e1b43-4679-4fc3-b1c7-ada7b90e815d','E3',59),
('7f94970e-40b5-4f7d-806d-2bbaa8a43eeb','E3',61),
('977ed959-9280-4ac0-8fdf-2a2bda6b02ec','E3',62),
('c3bf163a-2a6e-4717-8231-244400f5c3ca','E3',64))
a(EmpNo, BrCode, AOID)
left join staging.aomap m on a.empno = m.empno 
where m.empno is null)a

on conflict(BrCode, AOID) do update 
set 
  EmpNo = excluded.EmpNo;
  