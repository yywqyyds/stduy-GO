const axios = require("axios");
const url = "http://localhost:8081/api/questions/create";

async function getData(requestBody) {
  axios
    .post(url, requestBody, {
      headers: {
        "Content-Type": "application/json",
      },
      proxy:{
        protocol: 'http',
        host: '127.0.0.1',
        port: 8899,
      }
    })
    .then((response) => {
      console.log("响应数据:", response.data);
    })
    .catch((error) => {
      console.error("请求出错:", error.code);
    });
}

async function main() {
  const requestBody1 = {
    model: "deepseek", // 模型名称，不传就是 tongyi
    language: "javascript",
    type: 2, // 多选题
    keyword: "react基础知识",
  };
  await getData(requestBody1);

  new Promise((resolve) => setTimeout(resolve, 1000)); // 等待1秒

  const requestBody2 = {
    language: "go",
    // type: 1, //单选题
    keyword: "go语言gin框架",
  };
  await getData(requestBody2);
}

main();
