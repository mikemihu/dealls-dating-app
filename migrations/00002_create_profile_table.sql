-- +goose Up
-- +goose StatementBegin
CREATE TYPE gender AS ENUM ('male', 'female');

CREATE TABLE "profiles"
(
    id         UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    user_id    UUID UNIQUE                            NOT NULL
        CONSTRAINT profiles_user_id_fk REFERENCES "users" ON DELETE CASCADE,
    name       TEXT                                   NOT NULL,
    gender     gender                                 NOT NULL,
    bio        TEXT                                   NOT NULL,
    updated_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL
);

-- seeder
INSERT INTO profiles (user_id, name, gender, bio)
VALUES ('18aba2b3-ec74-447d-a9ab-83b8cd6d2b24', 'Andre', 'male', ''),
       ('c68964b7-a97e-4d33-b180-c3a6d0f52c09', 'Budi', 'male', ''),
       ('87d847d5-a7c5-46ea-9f28-5e3400de4fe8', 'Cynthia', 'female', ''),
       ('5be97d24-d8bb-433a-ade6-5e334dd70d20', 'Debby', 'female', ''),
       ('a177469d-e980-41d2-8e05-2c877d4e1abd', 'Elon', 'male', ''),
       ('f5cabe66-c5f8-43de-aeee-b3e527ba2843', 'Fina', 'female', ''),
       ('6b95e7dc-6484-48d1-be44-412d7eae7b08', 'George', 'male', ''),
       ('a6b3231d-f030-40bc-8e95-5a9f0831eaf6', 'Hannah', 'female', ''),
       ('0679b780-bdcb-4fb0-992a-7ffe96a532b1', 'Isabel', 'female', ''),
       ('773f1ee7-57ae-48c7-babf-082d62cb63dd', 'Jason', 'male', ''),
       ('dd37d63c-ac20-41ec-90a8-73ad6d15321f', 'Kevin', 'male', ''),
       ('52ab4454-fb3c-4189-b9f4-677cd2929b4c', 'Leonard', 'male', ''),
       ('3ef82d2a-1d01-48b4-b7f9-68c0b72fb8c7', 'Mike', 'male', '');
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "profiles" CASCADE;
DROP TYPE IF EXISTS "gender";
-- +goose StatementEnd
