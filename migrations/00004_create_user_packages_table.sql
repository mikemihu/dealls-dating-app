-- +goose Up
-- +goose StatementBegin
CREATE TABLE "user_packages"
(
    id         UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    user_id    UUID                                   NOT NULL
        CONSTRAINT user_packages_user_id_fk REFERENCES "users" ON DELETE CASCADE,
    package_id UUID                                   NOT NULL
        CONSTRAINT user_packages_package_id_fk REFERENCES "packages" ON DELETE CASCADE,
    expired_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);

CREATE INDEX idx_user_packages_user_id ON user_packages (user_id);

-- seeder
INSERT INTO user_packages (user_id, package_id, expired_at)
VALUES ('18aba2b3-ec74-447d-a9ab-83b8cd6d2b24', '4847798b-f27d-4f67-8f82-65b8ef3bbbe2', NOW() + INTERVAL '1 month'),
       ('c68964b7-a97e-4d33-b180-c3a6d0f52c09', '4847798b-f27d-4f67-8f82-65b8ef3bbbe2', NOW() + INTERVAL '1 month'),
       ('f5cabe66-c5f8-43de-aeee-b3e527ba2843', '120cc6a0-69ce-465c-8a50-c42be122134e', NOW() + INTERVAL '1 month'),
       ('dd37d63c-ac20-41ec-90a8-73ad6d15321f', '4847798b-f27d-4f67-8f82-65b8ef3bbbe2', NOW() + INTERVAL '1 month')
;
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "user_packages" CASCADE;
-- +goose StatementEnd
