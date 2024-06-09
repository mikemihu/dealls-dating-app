-- +goose Up
-- +goose StatementBegin
CREATE TABLE "likes"
(
    id             UUID PRIMARY KEY         DEFAULT uuid_generate_v4(),
    user_id        UUID                                   NOT NULL
        CONSTRAINT likes_user_id_fk REFERENCES "users" ON DELETE CASCADE,
    target_user_id UUID                                   NOT NULL
        CONSTRAINT likes_target_user_id_fk REFERENCES "users" ON DELETE CASCADE,
    updated_at     TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,
    created_at     TIMESTAMP WITH TIME ZONE DEFAULT NOW() NOT NULL,

    UNIQUE (user_id, target_user_id)
);

CREATE INDEX idx_likes_user_id_target_user_id ON likes (user_id, target_user_id);
-- +goose StatementEnd

-- +goose Down
-- +goose StatementBegin
DROP TABLE IF EXISTS "likes" CASCADE;
-- +goose StatementEnd
