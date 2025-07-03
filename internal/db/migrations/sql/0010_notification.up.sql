CREATE TABLE notifications (
    notification_id VARCHAR(255) NOT NULL PRIMARY KEY,
    user_id VARCHAR(255) NOT NULL,
    class_id VARCHAR(255),
    scheduler_id VARCHAR(255),
    description TEXT,
    type ENUM('success', 'warning', 'info'),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    
    CONSTRAINT fk_notification_user FOREIGN KEY (user_id) REFERENCES users(user_id)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    
    CONSTRAINT fk_notification_class FOREIGN KEY (class_id) REFERENCES classrooms(class_id)
        ON UPDATE CASCADE
        ON DELETE CASCADE,
    
    
    INDEX idx_user_id (user_id),
    INDEX idx_class_id (class_id)
);
