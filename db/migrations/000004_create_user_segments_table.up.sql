CREATE TABLE users_segments
(
    user_id     BIGINT NOT NULL,
    segment_ids BIGINT[] NOT NULL,
    PRIMARY KEY (user_id, segment_ids),
    FOREIGN KEY (user_id) REFERENCES users(id),
    FOREIGN KEY (segment_ids) REFERENCES segments(id)
)