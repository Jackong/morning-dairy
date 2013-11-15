CREATE TABLE dairy (
  id INT UNSIGNED NOT NULL AUTO_INCREMENT,
  date DATE DEFAULT 0,
  time TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  boards VARCHAR(100) NOT NULL,
  PRIMARY KEY(id),
  UNIQUE KEY `idx_date`(date)
) ENGINE = InnoDB charset = utf8;