# book_crawler

豆瓣图书TOP250榜单数据抓取

本项目仅供学习参考

---

### 1. 依赖项安装
```
go get github.com/jinzhu/gorm
go get github.com/PuerkitoBio/goquery
```

### 2. 创建数据库并配置`model`中的数据库连接参数
``` sql
CREATE DATABASE book_crawler DEFAULT CHARACTER SET utf8 DEFAULT COLLATE utf8_general_ci;
```

### 3. 运行项目
``` bash
go run main.go
```