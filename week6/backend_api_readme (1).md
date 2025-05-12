
# 生成题库系统后端接口文档

本项目是一个使用 Go 语言 + Gin 框架构建的后端服务，支持 AI 出题、题目管理和存储。前端页面功能包括题目生成、展示、筛选、编辑、删除等操作。

---

## 📌 接口汇总

### 1. 生成 AI 题目（调用大模型）
- **URL**：`POST /api/questions/generate`
- **说明**：调用大模型生成题目，支持选择题、编程题等。
- **请求参数**（JSON）：
```json
{
  "model": "tongyi",
  "language": "go",
  "type": 1,
  "keyword": "gin",
  "difficulty": "简单",
  "count": 3
}
```
- **返回参数**：
```json
{
  "code": 0,
  "msg": "",
  "data": [ {...}, {...} ]
}
```

---

### 2. 保存 AI 题目
- **URL**：`POST /api/questions/save`
- **说明**：将生成的题目保存至 SQLite。
- **请求参数**（JSON）：
```json
{
  "questions": [ {...}, {...} ]
}
```
- **返回**：
```json
{
  "code": 0,
  "msg": "保存成功",
  "data": null
}
```

---

### 3. 获取题目列表（分页）
- **URL**：`GET /api/questions`
- **说明**：获取题库中所有题目，支持分页。
- **参数**（Query）：
  - `page`：页码，默认 1
  - `size`：每页数量，默认 10

---

### 4. 删除题目
- **URL**：`DELETE /api/questions/:id`
- **说明**：根据题目 ID 删除题目。

---

### 5. 批量删除题目
- **URL**：`POST /api/questions/batch-delete`
- **说明**：批量删除题目。
- **请求参数**：
```json
{
  "ids": [1, 2, 3]
}
```

---

### ✅ 6. 获取指定类型题目列表
- **URL**：`GET /api/questions/type/:type`
- **说明**：按题型分类获取题目，如：单选题、多选题、编程题。
- **示例请求**：`/api/questions/type/单选题`

---

### ✅ 7. 搜索题目标题
- **URL**：`GET /api/questions/search`
- **说明**：根据题目标题关键词进行模糊查询。
- **Query 参数**：
  - `keyword`：题目关键字
- **示例请求**：`/api/questions/search?keyword=输出`

---

## 🗃️ 数据库字段设计（questions 表）

| 字段名           | 类型    | 说明             |
|------------------|---------|------------------|
| id               | int     | 主键             |
| model            | text    | 使用的模型       |
| language         | text    | 编程语言         |
| type             | int     | 题目类型（1/2/3）|
| keyword          | text    | 提示词           |
| question         | text    | 题干             |
| options          | text    | JSON 格式选项    |
| answer           | text    | JSON 格式答案    |
| explanation      | text    | 解析             |
| ai_start_time    | text    | 开始时间         |
| ai_end_time      | text    | 结束时间         |
| ai_cost_time     | int     | 耗时（毫秒）     |

---

## 📌 说明
- 支持按类型筛选（单选、多选、编程）
- 编程题没有 options 和 answer 字段，页面需适配
- 支持搜索题目名模糊匹配

---
