import math

paths1 = []
paths2 = []
with open('input.txt', 'r') as file:
    paths1 = file.readline().strip().split(',')
    paths2 = file.readline().strip().split(',')

def nearest_cross(paths1, paths2):
    # populate first path
    full_path1 = set((0, 0))
    x = 0
    y = 0
    for path in paths1:
        direction = path[0]
        amount = int(path[1:])
        for i in range(0, amount):
            if direction == 'R':
                x += 1
            elif direction == 'L':
                x -= 1
            elif direction == 'U':
                y += 1
            elif direction == 'D':
                y -= 1
            full_path1.add((x, y))
    
    # check for nodes 
    nearest_distance = math.inf

    x = 0
    y = 0
    for path in paths2:
        direction = path[0]
        amount = int(path[1:])
        for i in range(0, amount):
            if direction == 'R':
                x += 1
            elif direction == 'L':
                x -= 1
            elif direction == 'U':
                y += 1
            elif direction == 'D':
                y -= 1

            if (x, y) in full_path1:
                distance = abs(x) + abs(y)
                if distance < nearest_distance:
                    nearest_distance = distance
    return nearest_distance

assert nearest_cross(['R8','U5','L5','D3'],['U7','R6','D4','L4']) == 6
assert nearest_cross(['R75','D30','R83','U83','L12','D49','R71','U7','L72'], ['U62','R66','U55','R34','D71','R55','D58','R83']) == 159
assert nearest_cross(['R98','U47','R26','D63','R33','U87','L62','D20','R33','U53','R51'], ['U98','R91','D20','R16','D67','R40','U7','R15','U6','R7']) == 135

print(nearest_cross(paths1, paths2))
