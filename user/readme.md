# User Service

## Feature Schema

Feature                       | Method  |Endpoint
--------                      |-------  |--------
Register ( Admin & User )     | Post    |{base_url}/api/user/register 
login                         | Post    |{base_url}/api/user/auth/login
Update Profile                | Patch   |{base_url}/api/user/{user_id}
Show Profile                  | Get     |{base_url}/api/user/{user_id}

<br>

## User Table

- tbl_user

Field          | Type
------         |------
id             | int-pk
name           | string
email          | string-unique
no_hp          | string
address        | text
is_admin       | boolean
avatar         | string-nullable
created_at     | timestamp
updated_at     | timestamp-nullable
deleted_at     | timestamp-nullable


- tbl_user_auth

Field          | Type
------         |------
id             | int-pk
user_id        | int-fk
email          | string-unique
password       | string-hash
is_active      | boolean
last_login     | timestamp-nullable
created_at     | timestamp
updated_at     | timestamp-nullable
deleted_at     | timestamp-nullable