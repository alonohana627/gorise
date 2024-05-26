# gorise - Basic Phonebook API w/Postgresql as a DB

## How to Run

One method is through Docker. One can use the docker-compose in the repo as an example:
`docker-compose up -d`

The docker-compose file's template pulls automatically the Postgresql's image and uses it as a DB. The DB is
configurable through the docker-compose.

In order to run locally, go to `/db/client.go` and edit the parameters to be taken out of constant strings or other OS
environment variables.

## API

The API has **6 endpoints** for CRUD operations:
<!--------------------------------------------------------->

### GET /api/v1/read-contacts

Retrieves contacts in pages, with 10 contacts per page. By using the offset query parameter, you can specify the
starting point for pagination to begin from a later position.

#### CURL Example:

`curl --location 'localhost:8080/api/v1/read-contacts?offset=[NUMBER]'`

#### Responses:

- 200 OK: Successfully retrieved contacts list.
- 405 Method Not Allowed: Request method is not GET.
- 500 Internal Server Error: Error fetching contacts or marshaling response.

#### Response example 200:

```json
{
  "name": string,
  "lastName": string,
  "phone_number": string,
  "address": string(nullable)
}
```

<!--------------------------------------------------------->

### POST /api/v1/create-contact

Handles the addition of a new contact to the phonebook.

#### CURL example:

```bash
curl --location 'localhost:8080/api/v1/create-contact'
--header 'Content-Type: application/json'
--data '{
    "name": "Alonusz",
    "lastName": "Ohansus",
    "phone_number": "5554-77",
    "address": null
}'
```

#### Request Body

```json
{
  "name": string,
  "lastName": string,
  "phone_number": string,
  "address": string(nullable)
}
```

#### Responses:

- 200 OK: Contact added successfully or already exists.
- 400 Bad Request: Invalid request body.
- 405 Method Not Allowed: Request method is not POST.
- 500 Internal Server Error: Database insertion error.

<!--------------------------------------------------------->

### PATCH /api/v1/update-contact-by-name

Handles the editing of a contact's details based on the name.

#### CURL Example:

```bash
 curl --location --request PATCH 'localhost:8080/api/v1/update-contact-by-name'
--header 'Content-Type: application/json'
--data '{
    "phone_number": "123-456-7893",
    "name": "Alon",
    "lastName": "Ohana"
}'
```

#### Request Body:

```json
{
  "name": string,
  "last_name": string,
  "phone_number": string(nullable),
  "address": string(nullable)
}
```

#### Responses:

- 200 OK: Contact updated successfully.
- 400 Bad Request: Invalid request body.
- 405 Method Not Allowed: Request method is not PATCH.
- 500 Internal Server Error: Database update error.

<!--------------------------------------------------------->

### PATCH /api/v1/update-contact-by-phone

Handles the editing of a contact's details based on the phone number.

#### CURL Example:

```bash
 curl --location --request PATCH 'localhost:8080/api/v1/update-contact-by-phone'
--header 'Content-Type: application/json'
--data '{
    "phone_number": "052-840-8722",
    "name": "Alon",
    "lastName": "Ohana"
}'
```

#### Request Body:

```json
{
  "phone_number": string,
  "name": string(nullable),
  "lastName": string(nullable),
  "address": string(nullable)
}
```

### DELETE /api/v1/delete-contact

Handles the deletion of a contact from the phonebook.

#### CURL Example:

```bash
 curl --location --request DELETE 'localhost:8080/api/v1/delete-contact'
--header 'Content-Type: application/json'
--data '{
    "phone_number":"052-840-8722"
}'
```

#### Request Body:

```json
{
  "phone_number": string,
  "name": string(nullable),
  "lastName": string(nullable),
  "address": string(nullable)
}
```

#### Responses:

- 200 OK: Contact deleted successfully or not found.
- 400 Bad Request: Invalid request body or deletion error.
- 405 Method Not Allowed: Request method is not DELETE.
- 500 Internal Error.

### POST /api/v1/search-contacts

Handles searching for contacts in the phonebook based on search criteria. Can handle partial names/numbers.

#### CURL Example:

```bash
 curl --location 'localhost:8080/api/v1/search-contacts'
--header 'Content-Type: application/json'
--data '{
    "phone_number": "123-"
}'
```

#### Request Body:
```json
{
  "phone_number": string(nullable),
  "name": string(nullable),
  "lastName": string(nullable),
}
```

#### Responses:
- 200 OK: Successfully retrieved list of matching contacts.
- 400 Bad Request: Invalid request body.
- 405 Method Not Allowed: Request method is not POST.
- 500 Internal Server Error: Error searching contacts or marshaling response.

#### Example for 200 Response:

```json
[
    {
        "name": "Alonusz",
        "lastName": "Ohansus",
        "phone_number": "052-840-8722",
        "address": "HaDror 3 Gedera"
    },
    {
        "name": "Eliahu",
        "lastName": "Anavim",
        "phone_number": "053-844-8722"
    }
]
```

##