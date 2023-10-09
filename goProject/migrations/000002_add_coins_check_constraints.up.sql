ALTER TABLE coins ADD CONSTRAINT coins_quantity_check CHECK (quantity >= 0);
ALTER TABLE coins ADD CONSTRAINT coins_year_check CHECK (year BETWEEN 1600 AND date_part('year', now()));