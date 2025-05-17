CREATE TABLE classrooms (
    class_id VARCHAR(255) PRIMARY KEY UNIQUE NOT NULL,
    class_name VARCHAR(255) NOT NULL,
    user_created VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    url TEXT NOT NULL,
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP,

    CONSTRAINT fk_classroom_teacher FOREIGN KEY (user_created)
        REFERENCES teachers(user_id)
        ON UPDATE CASCADE
        ON DELETE CASCADE
);