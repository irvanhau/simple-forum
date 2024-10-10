CREATE TABLE IF NOT EXISTS posts (
    id BIGINT AUTO_INCREMENT PRIMARY KEY,
    user_id INT NOT NULL,
    post_title VARCHAR(250) NOT NULL,
    post_content LONGTEXT NOT NULL,
    post_hashtags LONGTEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT  CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by LONGTEXT NOT NULL,
    updated_by LONGTEXT NOT NULL
);