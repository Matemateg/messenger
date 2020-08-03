CREATE TABLE `users` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `username` varchar(255) NOT NULL DEFAULT '',
    `created_at` datetime NOT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_user` (`username`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE `chats` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `name` varchar(255) NOT NULL DEFAULT '',
    `created_at` datetime NOT NULL,
    `last_message_time` datetime DEFAULT NULL,
    PRIMARY KEY (`id`),
    UNIQUE KEY `uniq_name` (`name`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE `chats_has_users` (
    `chat_id` int(11) unsigned NOT NULL,
    `user_id` int(11) unsigned NOT NULL,
    PRIMARY KEY (`chat_id`,`user_id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;


CREATE TABLE `messages` (
    `id` int(11) unsigned NOT NULL AUTO_INCREMENT,
    `chat_id` int(11) NOT NULL,
    `author_id` int(11) NOT NULL,
    `text` varchar(255) NOT NULL DEFAULT '',
    `created_at` datetime NOT NULL,
    PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;
