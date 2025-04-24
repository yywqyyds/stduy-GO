# AI 出题生成服务

本项目是一个基于 Gin + Go 的后端服务，支持通过调用阿里云百炼（通义/DeepSeek）等大模型接口，生成编程类选择题，并保存到本地 JSON 文件。

---

## ✨ 功能亮点

- ✅ 支持题型：单选 / 多选
- ✅ 支持语言：Go / JavaScript / Java / Python / C++
- ✅ 模型支持：通义千问、DeepSeek
- ✅ 数据保存：按日期存储到 `data/YYYY-MM-DD.json` 文件中
- ✅ 返回结构统一，便于前端解析
- ✅ 失败时不写入文件，返回错误信息

---

## 📦 请求说明

接口地址：

```
POST /api/questions/create
```

请求体参数（JSON）：

| 字段      | 类型   | 说明              |
|-----------|--------|-------------------|
| model     | string | 模型名称（tongyi / deepseek）|
| language  | string | 编程语言（go、javascript、java、python、c++）|
| type      | int    | 题型（1 单选，2 多选）|
| keyword   | string | 关键词（如：路由）|

---

## 📤 响应说明

成功示例：

```json
{
  "code": 0,
  "msg": "",
  "data": {
    "aiRes":{
      "question": "...",
      "options": {
        "A": "...",
        "B": "..."
      },
      "answer": ["A"],
      "explanation": "...",
    },
    "aiStartTime": "",
    "aiEndTime": "",
    "aiCostTime": 3,
    "aiReq": {
      "model": "tongyi",
      "language": "go",
      "type": "single",
      "keyword": "路由"
    }
  }
}
```

---

## 🗂️ 本地文件结构

所有生成的题目将保存在 `data/` 目录中，每日一个文件：

```
data/
├── 2025-04-23.json
├── 2025-04-24.json
```

每行一个 JSON 对象，便于后续导入数据库或展示。

---

## 🧱 模块说明

| 模块       | 说明                         |
|------------|------------------------------|
| `handler/` | 接口请求处理逻辑             |
| `service/` | 模型调用逻辑（调用 DashScope）|
| `schema/`  | 数据结构定义（Question、AiRequest）|
| `config/`  | 配置读取，读取 `.env` 中的 API_KEY 等 |

---

## ⚙️ 环境依赖

- Go 1.20+
- Gin Web Framework
- openai-go SDK （用于兼容阿里百炼）

---

## 🚀 快速运行

```bash
go mod tidy
go run main.go
```

访问地址：

```
http://localhost:8080/api/questions/create
```

建议使用 Postman 或 curl 测试接口。

---

## 🧠 作者提示

- 提示词中明确要求返回 **纯 JSON**，禁止 Markdown。
- 出现解析错误一般是模型返回格式不规范。

---

更新时间：2025-04-24
