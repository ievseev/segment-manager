CREATE TABLE users_segments
(
    id         SERIAL PRIMARY KEY,
    user_id    BIGINT NOT NULL,
    segment_id BIGINT NOT NULL
)