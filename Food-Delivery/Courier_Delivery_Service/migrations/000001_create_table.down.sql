-- Enum turlarini o'chirish
DROP TYPE IF EXISTS courier_status, pay_status, CourOrd_status;

-- Jadvallarni o'chirish
DROP TABLE IF EXISTS Couriers CASCADE;
DROP TABLE IF EXISTS Orders CASCADE;
DROP TABLE IF EXISTS CourierOrders CASCADE;
DROP TABLE IF EXISTS schema_migrations CASCADE;

