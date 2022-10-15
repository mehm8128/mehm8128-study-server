DROP DATABASE IF EXISTS mehm8128_study;
CREATE DATABASE mehm8128_study;
USE mehm8128_study;

CREATE TABLE IF NOT EXISTS `users` (
  `id` char(36) NOT NULL UNIQUE,
  `name` varchar(20) NOT NULL,
  `hashed_pass` varchar(200) NOT NULL UNIQUE,
  `description` varchar(140) DEFAULT '',
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NUll,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `files` (
  `id` char(36) NOT NULL UNIQUE,
  `file_name` varchar(36) NOT NULL,
  `created_by` char(36) NOT NULL,
  `created_at` datetime NOT NULL,
  `file` blob,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `goals` (
  `id` char(36) NOT NULL UNIQUE,
  `title` varchar(20) NOT NULL,
  `comment` varchar(140) DEFAULT '',
  `goal_date` char(10) NOT NULL,
  `is_completed` boolean DEFAULT False,
  `favorite_num` decimal(40) DEFAULT 0,
  `created_by` char(36) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NUll,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`created_by`) REFERENCES users(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `records` (
  `id` char(36) NOT NULL UNIQUE,
  `title` varchar(20) NOT NULL,
  `page` decimal(3) DEFAULT 0,
  `time` decimal(3) DEFAULT 0,
  `comment` varchar(140) DEFAULT '',
  `favorite_num` decimal(40) DEFAULT 0,
  `file_id` char(36),
  `created_by` char(36) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NUll,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`created_by`) REFERENCES users(`id`),
  FOREIGN KEY (`file_id`) REFERENCES files(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `record_favorites` (
  `id` char(36) NOT NULL UNIQUE,
  `record_id` char(36) NOT NULL,
  `created_by` char(36) NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`record_id`) REFERENCES records(`id`),
  FOREIGN KEY (`created_by`) REFERENCES users(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `goal_favorites` (
  `id` char(36) NOT NULL UNIQUE,
  `goal_id` char(36) NOT NULL,
  `created_by` char(36) NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`goal_id`) REFERENCES goals(`id`),
  FOREIGN KEY (`created_by`) REFERENCES users(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `memorizes` (
  `id` char(36) NOT NULL UNIQUE,
  `name` varchar(36) NOT NULL,
  `created_at` datetime NOT NULL,
  `updated_at` datetime NOT NULL,
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;

CREATE TABLE IF NOT EXISTS `words` (
  `id` char(36) NOT NULL UNIQUE,
  `memorize_id` char(36) NOT NULL,
  `word` varchar(36) NOT NULL,
  `word_jp` varchar(36) NOT NULL,
  `created_at` datetime NOT NULL,
  PRIMARY KEY (`id`),
  FOREIGN KEY (`memorize_id`) REFERENCES memorizes(`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_general_ci;