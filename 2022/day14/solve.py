data = [[[int(coord) for coord in pair.split(',')] for pair in line.strip().split(' -> ')] for line in open('input.txt', 'r')]

def initiate_rocks(rock_corners):
    def update_lowest(check, low):
        if check > low:
            return check
        return low

    lowest_point = 0
    rocks = set()
    for corner in rock_corners:
        for start, end in zip(corner, corner[1:]):
            lowest_point = update_lowest(start[1], lowest_point)
            if end[0] < start[0]:
                start, end = end, start
            for x in range(start[0], end[0]+1):
                rocks.add((start[1], x))
            if end[1] < start[1]:
                start, end = end, start
            for x in range(start[1], end[1]+1):
                rocks.add((x, start[0]))
                lowest_point = update_lowest(x, lowest_point)
    return rocks, lowest_point

def move_one_down(point, obstacles, lowest):
    one_down = point[0]+1
    one_left = point[1]-1
    one_right = point[1]+1
    if one_down == lowest:
        return point, True
    if not (one_down, point[1]) in obstacles:
        return (one_down, point[1]), False
    elif not (one_down, one_left) in obstacles:
        return (one_down, one_left), False
    elif not (one_down, one_right) in obstacles:
        return (one_down, one_right), False
    return point, True

def drop_sand(source, obstacles, lowest):
    sand_position = source
    while True:
        sand_position, ended = move_one_down(sand_position, obstacles, lowest)
        if ended:
            return sand_position
        if sand_position[0] > lowest:
            return None

rocks, lowest_point = initiate_rocks(data)
sand = set()
sand_source = (0, 500)
while True:
    dropped_sand = drop_sand(sand_source, sand.union(rocks), lowest_point+2)
    sand.add(dropped_sand)
    if dropped_sand == sand_source:
        break
    if dropped_sand == None:
        raise Exception('sand fell out!')

print(len(sand))