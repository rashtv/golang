CREATE TABLE IF NOT EXISTS coins (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title text NOT NULL DEFAULT 'Coin Title',
    description text NOT NULL DEFAULT 'Coin Description',
    country text NOT NULL DEFAULT 'The World',
    status text NOT NULL DEFAULT 'Coin Status',
    quantity integer NOT NULL DEFAULT 1,
    material text NOT NULL DEFAULT 'Coin Material',
    auction_value integer NOT NULL DEFAULT 1,
    version integer NOT NULL DEFAULT 1
);