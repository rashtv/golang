CREATE TABLE IF NOT EXISTS coins (
    id bigserial PRIMARY KEY,
    created_at timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    title text NOT NULL,
    description text NOT NULL,
    country text NOT NULL,
    status text NOT NULL,
    quantity integer NOT NULL,
    material text NOT NULL,
    auction_value integer NOT NULL
);