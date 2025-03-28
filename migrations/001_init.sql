CREATE DATABASE IF NOT EXISTS companyapi;

USE companyapi;


CREATE TABLE IF NOT EXISTS companies (
    id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    name STRING NOT NULL,
    country STRING NOT NULL,
    email STRING
);

-- Create owners table
CREATE TABLE IF NOT EXISTS owners (
    id UUID NOT NULL PRIMARY KEY DEFAULT gen_random_uuid(),
    company_id UUID REFERENCES companies(id) ON DELETE CASCADE, -- Foreign key constraint
    name STRING NOT NULL,
    ssn STRING NOT NULL,
    UNIQUE(ssn)
);

CREATE INDEX IF NOT EXISTS index_company_owner ON owners (company_id);
