-- Migration Down Script

-- Drop the triggers that update 'updated_at' before removing tables
DROP TRIGGER IF EXISTS update_orders_updated_at ON orders;
DROP TRIGGER IF EXISTS update_products_updated_at ON products;
DROP TRIGGER IF EXISTS update_order_items_updated_at ON order_items;
DROP TRIGGER IF EXISTS update_kitchens_updated_at ON kitchens;
DROP TRIGGER IF EXISTS update_locations_updated_at ON locations;
DROP TRIGGER IF EXISTS update_notifications_updated_at ON notifications;

-- Drop the function used by the triggers
DROP FUNCTION IF EXISTS update_updated_at_column();

-- Drop the tables in reverse order of creation to handle dependencies
DROP TABLE IF EXISTS order_items;
DROP TABLE IF EXISTS products;
DROP TABLE IF EXISTS orders;
DROP TABLE IF EXISTS kitchens;
DROP TABLE IF EXISTS locations;
DROP TABLE IF EXISTS notifications;


-- (Optional) Drop the UUID extension if it was specifically created for this migration
DROP EXTENSION IF EXISTS "uuid-ossp";
