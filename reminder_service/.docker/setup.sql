CREATE TABLE IF NOT EXISTS reminders (
    id VARCHAR(255) PRIMARY KEY,
    pet_id VARCHAR(255),
    type VARCHAR(255),
    drug_name VARCHAR(255),
    drug_usage VARCHAR(255),
    frequency VARCHAR(255),
    created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
    updated_at DATETIME DEFAULT CURRENT_TIMESTAMP ON UPDATE CURRENT_TIMESTAMP
);

CREATE TABLE IF NOT EXISTS period_reminder_info (
    reminder_id VARCHAR(255) PRIMARY KEY,
    morning DATETIME,
    noon DATETIME,
    evening DATETIME,
    before_bed DATETIME,
    FOREIGN KEY (reminder_id) REFERENCES reminders(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS hour_reminder_info (
    reminder_id VARCHAR(255) PRIMARY KEY,
    first_usage DATETIME,
    every int,
    FOREIGN KEY (reminder_id) REFERENCES reminders(id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS notifications (
    id INT AUTO_INCREMENT PRIMARY KEY,
    reminder_id VARCHAR(255),
    user_id VARCHAR(255),
    time DATETIME,
    status VARCHAR(255),
    FOREIGN KEY (reminder_id) REFERENCES reminders(id) ON DELETE CASCADE
);