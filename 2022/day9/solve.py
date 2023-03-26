data = [line.strip().split(' ') for line in open('input.txt', 'r')]

def move_head(head, direction):
    match direction:
        case 'U':
            head[1] += 1
        case 'D':
            head[1] -= 1
        case 'L':
            head[0] -= 1
        case 'R':
            head[0] += 1
    
def is_touching(head, tail):
    return abs(head[0] - tail[0]) < 2 and abs(head[1] - tail[1]) < 2

def move_tail(tail, head):
    if head[0] > tail[0]:
        tail[0] += 1
    if head[0] < tail[0]:
        tail[0] -= 1
    if head[1] > tail[1]:
        tail[1] += 1
    if head[1] < tail[1]:
        tail[1] -= 1
    
knots = [[0,0] for _ in range(10)]
visited_spots = set()
for line in data:
    direction = line[0]
    amount = int(line[1])
    for _ in range(amount):
        move_head(knots[0], direction)
        for i in range(1, len(knots)):
            if not is_touching(knots[i], knots[i-1]):
                move_tail(knots[i], knots[i-1])
        visited_spots.add(tuple(knots[9]))

print(len(visited_spots))
