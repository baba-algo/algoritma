stack = []

def push(value):
    stack.append(value)

def pop():
    if len(stack) == 0:
        return None
    return stack.pop()

def peek():
    if len(stack) == 0:
        return None
    return stack[-1]
