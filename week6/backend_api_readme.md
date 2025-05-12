# 🧠 AI 题库系统后端接口文档

本项目基于 Go + Gin + SQLite 实现，支持调用大模型生成题目，并进行题库管理。

## 🗂 接口汇总

| 路由地址                      | 方法   | 描述                   |
|-------------------------------|--------|------------------------|
| `/api/questions/generate`     | POST   | 调用大模型生成题目（不保存） |
| `/api/questions/save`         | POST   | 保存生成的题目到数据库      |
| `/api/questions/list`         | GET    | 获取题库列表（支持筛选分页） |
| `/api/questions/delete/:id`   | DELETE | 删除题目               |       
| `/api/questions/modify/:id`   | POST   | 修改题目          |  
| `/api/questions/query/:id `   | GET    | 获取题目 |
| `/api/questions/type `        | GET    | 获取题型 |
---

## 1. 生成题目（不保存）

- **URL**：`POST /api/questions/generate`
- **参数**（JSON）：
```json
{
  "model": "tongyi",
  "language": "go",
  "type": 1,
  "keyword": "Gin 路由",
  "difficulty": "中等",
  "count": 3
}
```
- **返回**：
```json
{
  "code": 0,
  "msg": "",
  "data": [ { 题目结构 }, ... ]
}
```

---

## 2. 保存题目

- **URL**：`POST /api/questions/save`
- **参数**（JSON，结构与生成题目返回一致）：
```json
[
  {
    "aiStartTime": "...",
    "aiEndTime": "...",
    "aiCostTime": 123,
    "aiReq": { ... },
    "aiRes": { ... }
  },
  ...
]
```
- **返回**：
```json
{ "code": 0, "msg": "保存成功", "data": null }
```

---

## 3. 获取题目列表

- **URL**：`GET /api/questions/list?page=1&size=10&type=1`
- **返回**：
```json
{
  "code": 0,
  "msg": "",
  "data": {
    "total": 100,
    "list": [ { 题目结构 }, ... ]
  }
}
```

---

## 4. 删除题目

- **URL**：`DELETE /api/questions/delete/:id`
- **返回**：
```json
{ "code": 0, "msg": "删除成功" }
```

---

## 5. 批量删除题目

- **URL**：`POST /api/questions/batch-delete`
- **参数**：
```json
{ "ids": [1, 2, 3] }
```
- **返回**：
```json
{ "code": 0, "msg": "删除成功" }
```

---

## 📦 项目目录结构建议

```
server/
├── config/          // 配置文件读取
├── db/              // SQLite 初始化
├── handler/         // 接口逻辑
├── router/          // 路由注册
├── schema/          // 结构体定义
├── service/         // LLM调用逻辑
├── data/            // 生成JSON文件存储
├── main.go          // 程序入口
```
