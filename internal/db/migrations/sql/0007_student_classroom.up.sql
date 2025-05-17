CREATE TABLE student_classes (
    class_id VARCHAR(255) NOT NULL,
    user_id VARCHAR(255) NOT NULL,
    state ENUM('joined','waiting'),
    created_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    joined_at TIMESTAMP DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    PRIMARY KEY (class_id, user_id),
    
    CONSTRAINT fk_studentclass_user FOREIGN KEY (user_id) REFERENCES users(user_id)
        ON UPDATE CASCADE
        ON DELETE CASCADE,

    CONSTRAINT fk_studentclass_class FOREIGN KEY (class_id) REFERENCES classrooms(class_id)
        ON UPDATE CASCADE
        ON DELETE CASCADE,

    INDEX idx_class_id (class_id),
    INDEX idx_user_id (user_id)
);
