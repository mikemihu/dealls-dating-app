-- +goose Up
-- +goose StatementBegin
CREATE TABLE "packages"
(
    id                 UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    name               TEXT                                   NOT NULL,
    is_unlimited_swipe BOOLEAN                                NOT NULL,
    is_verified        BOOLEAN                                NOT NULL,
    price              NUMERIC(10)                            NOT NULL,
    updated_at         TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    created_at         TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);

-- seeder
INSERT INTO packages (id, name, is_unlimited_swipe, is_verified, price)
VALUES ('120cc6a0-69ce-465c-8a50-c42be122134e', 'Unlimited Swipe', true, false, 99000),
       ('4847798b-f27d-4f67-8f82-65b8ef3bbbe2', 'Verified Label', false, true, 99000);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "packages" CASCADE;
-- +goose StatementEnd
