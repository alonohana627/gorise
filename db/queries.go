package db

// Queries for initialization of the DB
const isTableExistsQuery string = `
SELECT EXISTS (
	SELECT FROM information_schema.tables 
	WHERE table_name = $1
)`

const createContactTableQuery string = `
CREATE TABLE contact (
    id SERIAL PRIMARY KEY,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    name VARCHAR(255) NOT NULL,
    last_name VARCHAR(255) NOT NULL,
    phone_number VARCHAR(255) NOT NULL UNIQUE,
    address VARCHAR(255),
    UNIQUE (name, last_name)
);
`

// Queries for CRUD and search operations
// Create
const InsertContactQuery string = `INSERT INTO contact (name, last_name, phone_number, address) VALUES ($1, $2, $3, $4);`

// Read
const GetAllContactsQuery string = `SELECT * FROM contact`
const GetContactByNameQuery string = `SELECT name, last_name, phone_number, address FROM contact LIMIT 10 OFFSET $1`
const GetSpecificContactByNameQuery string = `SELECT name, last_name, phone_number, address FROM contact WHERE name=$1`

// Update
const UpdateContactPhoneByNameQuery = `UPDATE contact SET phone_number=$3 WHERE name=$1 AND last_name=$2`
const UpdateContactAddressByNameQuery = `UPDATE contact SET address=$3 WHERE name=$1 AND last_name=$2`

const UpdateContactNameByPhoneQuery = `UPDATE contact SET name=$2 WHERE phone_number=$1`
const UpdateContactLastNameByPhoneQuery = `UPDATE contact SET last_name=$2 WHERE phone_number=$1`
const UpdateContactAddressByPhoneQuery = `UPDATE contact SET address=$2 WHERE phone_number=$1`

// Delete
const DeleteContactByPhone = `DELETE FROM contact WHERE phone_number=$1`

// search
const SearchContactByNameQuery string = `SELECT name, last_name, phone_number, address FROM contact WHERE name LIKE $1 OR phone_number LIKE $2`
