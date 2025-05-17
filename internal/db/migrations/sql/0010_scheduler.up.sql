CREATE TABLE schedulers (
    scheduler_id VARCHAR(255) NOT NULL PRIMARY KEY,
    room_id VARCHAR(255) NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    class_id VARCHAR(255),
    start_time DATETIME NOT NULL,
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,
    
    CONSTRAINT fk_scheduler_classroom FOREIGN KEY (class_id) REFERENCES classrooms(class_id)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    
    CONSTRAINT fk_scheduler_room FOREIGN KEY (room_id) REFERENCES rooms(room_id)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    
    CONSTRAINT fk_scheduler_user FOREIGN KEY (user_id) REFERENCES users(user_id)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    
    INDEX idx_room_id (room_id),
    INDEX idx_user_id (user_id),
    INDEX idx_class_id (class_id)
);