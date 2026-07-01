queue = []

def enqueue(value):
    queue.append(value)

def dequeue():
    if len(queue) == 0:
        return None
    return queue.pop(0)
