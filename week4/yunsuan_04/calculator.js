function add(a,b){
  return a+b
}

// 减法函数
function subtract(a, b) {
  return a - b;
}

// 乘法函数
function multiply(a, b) {
  return a * b;
}

// 除法函数
function divide(a, b) {
  if (b === 0) {
      throw new Error("除数不能为零");
  }
  return a / b;
}

//导出模块
module.exports = {add, subtract, multiply,divide}