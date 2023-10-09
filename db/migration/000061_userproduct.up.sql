----------------------------------------------------------------------------------------
CREATE TABLE IF NOT EXISTS public.User_Product (
----------------------------------------------------------------------------------------
  UUID uuid NOT NULL DEFAULT uuid_generate_v4(),
  User_ID bigint NOT NULL,
  Product_ID bigint NOT NULL,
  Allow boolean NOT NULL,
  Other_Info jsonb NULL,

  CONSTRAINT User_Product_pkey PRIMARY KEY (User_ID, Product_ID),  
  CONSTRAINT fkUserAccType FOREIGN KEY (Product_ID) REFERENCES Product(ID),
  CONSTRAINT fkUserUser FOREIGN KEY (User_ID) REFERENCES Users(ID)
);
CREATE UNIQUE INDEX IF NOT EXISTS idxUser_Product_UUID ON public.User_Product(UUID);

DROP TRIGGER IF EXISTS trgUser_ProductIns on User_Product;
---------------------------------------------------------------------------
CREATE TRIGGER trgUser_ProductIns
---------------------------------------------------------------------------
    BEFORE INSERT ON User_Product
    FOR EACH ROW
    EXECUTE PROCEDURE trgGenericInsert();
  
DROP TRIGGER IF EXISTS trgUser_Productupd on User_Product;
---------------------------------------------------------------------------
CREATE TRIGGER trgUser_Productupd
---------------------------------------------------------------------------
    BEFORE UPDATE ON User_Product
    FOR EACH ROW
    WHEN (OLD.* IS DISTINCT FROM NEW.*)
    EXECUTE PROCEDURE trgGenericUpdate();

DROP TRIGGER IF EXISTS trgUser_Product_del on User_Product;
---------------------------------------------------------------------------
CREATE TRIGGER trgUser_Product_del
---------------------------------------------------------------------------
    AFTER DELETE ON User_Product
    FOR EACH ROW 
    EXECUTE FUNCTION trgGenericDelete();

    
----------------------------------------------------------------------------------------
CREATE OR REPLACE VIEW public.vwUser_Product
----------------------------------------------------------------------------------------
  AS SELECT 
    mr.UUID,
    u.ID User_ID, u.Login_Name,

    y.ID Product_ID, y.UUID ProductUUID, y.Product_Name,
        
    uat.Allow,
   
    mr.Mod_Ctr,
    uat.Other_Info,
    mr.Created,
    mr.Updated 
   FROM User_Product uat
   INNER JOIN Main_Record mr on mr.UUID = uat.UUID
   LEFT JOIN Product y ON y.ID = uat.Product_ID
   LEFT JOIN Users u ON u.ID = uat.User_ID
    ;

  INSERT INTO User_Product
    (User_ID, Product_ID, Allow)
  SELECT 
     u.ID User_ID, y.ID Product_ID, a.Allow
  FROM
   (Values
    ('erick1421@gmail.com', 'Loan', true))
     a(Login_Name, Product, Allow)  
  INNER JOIN Product y on y.Product_Name = a.Product
  INNER JOIN Users    u on u.Login_Name = a.Login_Name
  
  ON CONFLICT(User_ID, Product_ID) DO UPDATE SET
    Allow = EXCLUDED.Allow,
    Other_Info = EXCLUDED.Other_Info
  ;
