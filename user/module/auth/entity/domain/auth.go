package domain

type AuthDomain struct{
     id string
     user_id string
     email string
     is_active bool
     last_login int64
     created_at int64
     updated_at int64
     deleted_at int64
}