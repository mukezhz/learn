function add(x) {
  return x + 10;
}
function subtract(x) {
  return x - 5;
}
function multiply(x) {
  return x * 2;
}

// Without pipeline operator
let val1 = add(subtract(add(multiply(10))));
console.log(val1);

// Using pipeline operator
let val2 = add(subtract(add(multiply(10))));
console.log(val2);
function oldStuff() {
  const numbers = [1, 2, 3, 4, 5];
  const result = numbers.map(n => n * 2).filter(n => n > 4).reduce((acc, n) => acc + n, 0);
  console.log(result); // 14
}
function newStuff() {
  const numbers = [1, 2, 3, 4, 5];
  const result = numbers.map(n => n * 2).filter(n => n > 4).reduce((acc, n) => acc + n, 0);
  console.log({
    result
  });
}
oldStuff();
newStuff();
