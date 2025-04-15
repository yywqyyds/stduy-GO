## 📂 项目结构

```
├── main.go                  # 主程序
├── json                    # 示例 JSON 文件
├── words.db               # 输出的 SQLite 数据库
├── go.mod                 # 依赖文件
├── go.sum                 # 校验依赖的完整性、安全性
```

---

程序运行后将：
   - 读取 JSON 文件
   - 将单词、翻译、短语分别插入三张表中
   - 生成 `word.db` 数据库
   - 打印总耗时（单位：秒）

---

## 🧱 数据库表结构

### `words`
| 字段   | 类型     | 描述   |
|--------|----------|--------|
| id     | INTEGER  | 主键   |
| word   | TEXT     | 单词   |

### `translations`
| 字段       | 类型     | 描述     |
|------------|----------|----------|
| id         | INTEGER  | 主键     |
| word_id    | INTEGER  | 外键     |
| type       | TEXT     | 词性     |
| translation | TEXT    | 翻译     |

### `phrases`
| 字段               | 类型     | 描述         |
|--------------------|----------|--------------|
| id                 | INTEGER  | 主键         |
| word_id            | INTEGER  | 外键         |
| phrase             | TEXT     | 短语         |
| phrase_translation | TEXT     | 短语含义     |

---

