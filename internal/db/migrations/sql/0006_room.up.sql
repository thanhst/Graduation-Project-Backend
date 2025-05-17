CREATE TABLE rooms (
    room_id VARCHAR(255) NOT NULL PRIMARY KEY,
    class_id VARCHAR(255),
    state ENUM('opening','closed'),
    host VARCHAR(255) NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    ended_at TIMESTAMP NULL,

    CONSTRAINT fk_room_class FOREIGN KEY (class_id) REFERENCES classrooms(class_id)
        ON UPDATE CASCADE
        ON DELETE CASCADE,

    CONSTRAINT fk_room_user FOREIGN KEY (host) REFERENCES users(user_id)
        ON UPDATE CASCADE
        ON DELETE CASCADE,

    INDEX idx_class_id (class_id)
);
