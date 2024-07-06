CREATE TYPE Courier_status AS ENUM ('available', 'unavailable', 'on_delivery') 
CREATE TYPE User_role AS ENUM ('user', 'admin') 

CREATE TABLE Users (
    user_id UUID PRIMARY KEY, -- Foydalanuvchi noyob ID'si
    email VARCHAR(100) UNIQUE NOT NULL, -- Foydalanuvchi email manzili
    password_hash VARCHAR(255) NOT NULL, -- Parolni hash qilingan qiymati
    full_name VARCHAR(100) NOT NULL, -- Foydalanuvchi to'liq ismi
    role User_role DEFAULT 'user', -- Foydalanuvchi rol (masalan, 'user', 'admin')
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP, -- Foydalanuvchi yaratilgan vaqti
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);


CREATE TABLE Couriers (
    courier_id UUID PRIMARY KEY, -- Kuryer noyob ID'si
    email VARCHAR(100) UNIQUE NOT NULL, -- Kuryer email manzili
    password_hash VARCHAR(255) NOT NULL, -- Parolni hash qilingan qiymati
    full_name VARCHAR(100) NOT NULL, -- Kuryer to'liq ismi
    status Courier_status DEFAULT 'available', -- Kuryer holati ("mavjud", "mavjud emas", "etkazib berishda")
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP, -- Kuryer yaratilgan vaqti
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);


CREATE TABLE UserSessions (
    session_id UUID PRIMARY KEY,
    user_id UUID NOT NULL REFERENCES Users(user_id) ON DELETE CASCADE,
    token VARCHAR(500) UNIQUE NOT NULL, -- JWT token or session ID
    created_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    expires_at TIMESTAMP, -- Token expiration time
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT CURRENT_TIMESTAMP,
    deleted_at BIGINT DEFAULT 0
);
