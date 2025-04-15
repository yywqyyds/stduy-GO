// index.js
const calculator = require('./calculator'); // 导入 calculator 模块

// 获取命令行输入的参数
const args = process.argv.slice(2); // 提取命令行参数，去掉前两个默认参数

// 检查参数是否存在
if (args.length < 3) {
    console.log("错误: 请提供操作符和两个数字参数");
    console.log("用法示例: node index.js add 2 3");
    process.exit(1); // 退出程序
}

const operator = args[0]; // 操作符（如 add, subtract, multiply, divide）
const num1 = parseFloat(args[1]); // 第一个数字
const num2 = parseFloat(args[2]); // 第二个数字

// 根据操作符执行相应的计算
let result;
switch (operator) {
    case 'add':
        result = calculator.add(num1, num2);
        break;
    case 'subtract':
        result = calculator.subtract(num1, num2);
        break;
    case 'multiply':
        result = calculator.multiply(num1, num2);
        break;
    case 'divide':
        try {
            result = calculator.divide(num1, num2);
        } catch (error) {
            console.log("错误:", error.message);
            process.exit(1); // 如果除数为 0，退出程序
        }
        break;
    default:
        console.log("错误: 无效的操作符，请使用 add, subtract, multiply 或 divide");
        process.exit(1); // 退出程序
}

console.log("结果:", result); // 输出计算结果
