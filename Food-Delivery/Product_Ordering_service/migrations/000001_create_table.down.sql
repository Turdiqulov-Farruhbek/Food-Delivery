-- Bog'langan jadvallarni o'chirish
DROP TABLE IF EXISTS OrderRecommendations CASCADE;
DROP TABLE IF EXISTS OrderItems CASCADE;
DROP TABLE IF EXISTS CartItems CASCADE;

-- Asosiy jadvallarni o'chirish
DROP TABLE IF EXISTS P_Orders CASCADE;
DROP TABLE IF EXISTS Carts CASCADE;
DROP TABLE IF EXISTS Products CASCADE;
DROP TABLE IF EXISTS schema_migrations CASCADE;c



-- Enum turini o'chirish
DROP TYPE IF EXISTS Card_Type CASCADE;
