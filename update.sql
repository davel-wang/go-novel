-- --------------------------------------------------------
-- 主机:                           127.0.0.1
-- 服务器版本:                        5.5.53 - MySQL Community Server (GPL)
-- 服务器OS:                        Win32
-- HeidiSQL 版本:                  10.2.0.5599
-- --------------------------------------------------------

/*!40101 SET @OLD_CHARACTER_SET_CLIENT=@@CHARACTER_SET_CLIENT */;
/*!40101 SET NAMES utf8 */;
/*!50503 SET NAMES utf8mb4 */;
/*!40014 SET @OLD_FOREIGN_KEY_CHECKS=@@FOREIGN_KEY_CHECKS, FOREIGN_KEY_CHECKS=0 */;
/*!40101 SET @OLD_SQL_MODE=@@SQL_MODE, SQL_MODE='NO_AUTO_VALUE_ON_ZERO' */;

-- Dumping structure for table novel.authors
CREATE TABLE IF NOT EXISTS `authors` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `name` varchar(50) NOT NULL DEFAULT '' COMMENT '作者信息',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8;

-- Dumping data for table novel.authors: ~1 rows (大约)
DELETE FROM `authors`;
/*!40000 ALTER TABLE `authors` DISABLE KEYS */;
INSERT INTO `authors` (`id`, `name`) VALUES
	(1, '巫门老九');
/*!40000 ALTER TABLE `authors` ENABLE KEYS */;

-- Dumping structure for table novel.books
CREATE TABLE IF NOT EXISTS `books` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `author_id` int(11) NOT NULL DEFAULT '0' COMMENT '作者',
  `name` varchar(100) NOT NULL DEFAULT '' COMMENT '书名',
  `from` varchar(500) NOT NULL DEFAULT '' COMMENT '来源 一行一个',
  `views` int(11) NOT NULL DEFAULT '0' COMMENT '总访问次数',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB AUTO_INCREMENT=27 DEFAULT CHARSET=utf8;

-- Dumping data for table novel.books: ~25 rows (大约)
DELETE FROM `books`;
/*!40000 ALTER TABLE `books` DISABLE KEYS */;
INSERT INTO `books` (`id`, `author_id`, `name`, `from`, `views`) VALUES
	(2, 1, '黄泉杂货铺', 'https://m.biqushu.com/chapters_108751/', 0),
	(3, 1, '黄泉杂货铺', 'https://m.biqushu.com/chapters_108751/', 0),
	(4, 1, '黄泉杂货铺', 'https://m.biqushu.com/chapters_108751/', 0),
	(5, 1, '黄泉杂货铺', 'https://m.biqushu.com/chapters_108751/', 0),
	(6, 1, '黄泉杂货铺', 'https://m.biqushu.com/chapters_108751/', 0),
	(7, 1, '黄泉杂货铺', 'https://m.biqushu.com/chapters_108751/', 0),
	(8, 1, '黄泉杂货铺', 'https://m.biqushu.com/chapters_108751/', 0),
	(9, 1, '黄泉杂货铺', 'https://m.biqushu.com/chapters_108751/', 0),
	(10, 1, '黄泉杂货铺', 'https://m.biqushu.com/chapters_108751/', 0),
	(11, 1, '黄泉杂货铺', 'https://m.biqushu.com/chapters_108751/', 0),
	(12, 1, '黄泉杂货铺', 'https://m.biqushu.com/chapters_108751/', 0),
	(13, 1, '黄泉杂货铺', 'https://m.biqushu.com/chapters_108751/', 0),
	(14, 1, '黄泉杂货铺', 'https://m.biqushu.com/chapters_108751/', 0),
	(15, 1, '黄泉杂货铺', 'https://m.biqushu.com/chapters_108751/', 0),
	(16, 1, '黄泉杂货铺', 'https://m.biqushu.com/chapters_108751/', 0),
	(17, 1, '黄泉杂货铺', 'https://m.biqushu.com/chapters_108751/', 0),
	(18, 1, '黄泉杂货铺', 'https://m.biqushu.com/chapters_108751/', 0),
	(19, 1, '黄泉杂货铺', 'https://m.biqushu.com/chapters_108751/', 0),
	(20, 1, '黄泉杂货铺', 'https://m.biqushu.com/chapters_108751/', 0),
	(21, 1, '黄泉杂货铺', 'https://m.biqushu.com/chapters_108751/', 0),
	(22, 1, '黄泉杂货铺', 'https://m.biqushu.com/chapters_108751/', 0),
	(23, 1, '黄泉杂货铺', 'https://m.biqushu.com/chapters_108751/', 0),
	(24, 1, '黄泉杂货铺', 'https://m.biqushu.com/chapters_108751/', 0),
	(25, 1, '黄泉杂货铺', 'https://m.biqushu.com/chapters_108751/', 0),
	(26, 1, '黄泉杂货铺', 'https://m.biqushu.com/chapters_108751/', 0);
/*!40000 ALTER TABLE `books` ENABLE KEYS */;

-- Dumping structure for table novel.books_list
CREATE TABLE IF NOT EXISTS `books_list` (
  `id` int(11) NOT NULL AUTO_INCREMENT,
  `author_id` int(11) NOT NULL DEFAULT '0' COMMENT '作者',
  `book_id` int(11) NOT NULL DEFAULT '0' COMMENT '书',
  `title` varchar(500) NOT NULL DEFAULT '' COMMENT '标题',
  PRIMARY KEY (`id`)
) ENGINE=InnoDB DEFAULT CHARSET=utf8;

-- Dumping data for table novel.books_list: ~0 rows (大约)
DELETE FROM `books_list`;
/*!40000 ALTER TABLE `books_list` DISABLE KEYS */;
/*!40000 ALTER TABLE `books_list` ENABLE KEYS */;

/*!40101 SET SQL_MODE=IFNULL(@OLD_SQL_MODE, '') */;
/*!40014 SET FOREIGN_KEY_CHECKS=IF(@OLD_FOREIGN_KEY_CHECKS IS NULL, 1, @OLD_FOREIGN_KEY_CHECKS) */;
/*!40101 SET CHARACTER_SET_CLIENT=@OLD_CHARACTER_SET_CLIENT */;
