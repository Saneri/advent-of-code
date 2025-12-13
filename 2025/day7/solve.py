with open('input.txt', 'r') as file:
    lines = file.readlines()

width = len(lines[0])
height = len(lines)
cache = {}
def flow_down(beam_col: int, row: int) -> int:
    if row >= height:
        return 1
    
    cached = cache.get((beam_col, row))
    if cached:
        return cached
    
    sum = 0
    if lines[row][beam_col] == '^':
        if beam_col > 0:
            sum += flow_down(beam_col - 1, row + 1)
        if beam_col < width:
            sum += flow_down(beam_col + 1, row + 1)
    else:
        sum = flow_down(beam_col, row + 1)
    cache[(beam_col, row)] = sum
    return sum

start = lines[0].index('S')
print(flow_down(start, 1))
