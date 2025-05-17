CREATE TABLE emotions (
    id INT AUTO_INCREMENT PRIMARY KEY,
    room_id VARCHAR(255) NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    emotion ENUM('Happy', 'Sad', 'Neutral', 'Fear', 'Surprise'),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,

    INDEX idx_room_id (room_id),
    INDEX idx_user_id (user_id),
    INDEX idx_created_at (created_at),

    CONSTRAINT fk_emotion_room FOREIGN KEY (room_id)
        REFERENCES rooms(room_id)
        ON UPDATE CASCADE
        ON DELETE CASCADE,

    CONSTRAINT fk_emotion_user FOREIGN KEY (user_id)
        REFERENCES users(user_id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);
