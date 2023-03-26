with open('input.txt', 'r') as file:
    data = file.read()

MARKER_LEN = 14

processed_chars = MARKER_LEN
queue = list(data[:MARKER_LEN])

while len(set(queue)) < MARKER_LEN:
    queue.pop(0)
    queue.append(data[processed_chars])
    processed_chars += 1

print(processed_chars)
