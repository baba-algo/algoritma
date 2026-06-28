const stack = [];

function push(value) {
  stack.push(value);
}

function pop() {
  if (stack.length === 0) return undefined;
  return stack.pop();
}

function peek() {
  if (stack.length === 0) return undefined;
  return stack.at(-1);
}
