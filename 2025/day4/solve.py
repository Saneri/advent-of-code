with open('input.txt', 'r') as file:
    lines = [line.strip() for line in file.readlines()]

def to_coordinates(input_lines: list[str]) -> dict[tuple[int, int], str]:
    grid = {}
    for row_index, row in enumerate(input_lines):
        for col_index, char in enumerate(row):
            grid[(row_index, col_index)] = char
    return grid

def can_remove(map: dict[tuple[int, int], str], pos: tuple[int, int]) -> bool:
    if map[pos] != '@':
        return False
    
    rolls = 0
    positions = [
        (pos[0] - 1, pos[1]),
        (pos[0] + 1, pos[1]),
        (pos[0], pos[1] - 1),
        (pos[0], pos[1] + 1),
        (pos[0] - 1, pos[1] - 1),
        (pos[0] + 1, pos[1] + 1),
        (pos[0] - 1, pos[1] + 1),
        (pos[0] + 1, pos[1] - 1),
    ]
    for pos in positions:
        if map.get(pos) == '@':
            rolls += 1
    return rolls < 4

map = to_coordinates(lines)
count = 0
no_removals = False
while not no_removals:
    removed = []
    for i in map:
        if can_remove(map, i):
            count += 1
            removed.append(i)
            
    for rem in removed:
        map[rem] = '.'
    if not removed:
        no_removals = True

print(count)