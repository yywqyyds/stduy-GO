
1. 路由设计

| 接口              | 方法 | 功能说明               |
|-------------------|------|-----------------------|
| `/upload`         | POST | 多文件上传             |
| `/download/:uuid` | GET  | 下载指定文件           |
| `/preview/:uuid`  | GET  | 在线预览文件           |
| `/files`          | GET  | 分页获取文件列表       |
| `/stats`          | GET  | 获取文件统计信息       |

---

2. 数据表结构设计（SQLite）

```sql
CREATE TABLE IF NOT EXISTS files (
  id INTEGER PRIMARY KEY AUTOINCREMENT,
  uuid TEXT NOT NULL UNIQUE,
  filename TEXT NOT NULL,
  path TEXT NOT NULL UNIQUE,
  size INTEGER NOT NULL,
  mime_type TEXT,
  file_type TEXT,
  uploaded_by TEXT,
  created_at DATETIME DEFAULT CURRENT_TIMESTAMP
);
```

示例返回数据（上传成功）

```json
{
  "code": 0,
  "msg": "",
  "data": [
    {
      "uuid": "8e2d9f50-c271-4c33-a38c-d0edec75e921",
      "filename": "hello.jpg",
      "size": 102400
    }
  ]
}
```
