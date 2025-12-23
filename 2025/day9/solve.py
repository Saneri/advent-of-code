from itertools import combinations

with open('input.txt', 'r') as file:
    lines = [tuple(map(int, l.strip().split(','))) for l in file.readlines()]

def is_inside(p1: tuple[int, int], p2: tuple[int, int], edges) -> bool:
    left = p1[0] if p1[0] < p2[0] else p2[0]
    right = p1[0] if p1[0] > p2[0] else p2[0]
    top = p1[1] if p1[1] < p2[1] else p2[1]
    bottom = p1[1] if p1[1] > p2[1] else p2[1]

    if top == bottom or left == right:
        return True

    # check if area is inside polygon at all by ray tracing from middle to top
    # find all the edges that are horizontal, are above the top line, and cross the middle x
    # if the number of these edges is odd, then we are outside the polygon
    rays_crossed = 0
    mid_x = (left + right) / 2
    for edge in edges:
        a = edge[0]
        b = edge[1]
        if a[1] == b[1] and a[1] < top: # horizontal edge above
            edge_left = a[0] if a[0] < b[0] else b[0]
            edge_right = a[0] if a[0] > b[0] else b[0]
            if edge_left <= mid_x <= edge_right:
                rays_crossed += 1
    if rays_crossed % 2 == 1:
        return False

    for edge in edges:
        a = edge[0]
        b = edge[1]
        # check edge is not p1 and p2
        if (a == p1 and b == p2) or (a == p2 and b == p1):
            continue

        # no points should be inside rectangle
        if right > a[0] > left and top < a[1] < bottom:
            return False

        # vertical edge
        if a[0] == b[0] and right > a[0] > left:
            edge_top = a[1] if a[1] < b[1] else b[1]
            edge_bottom = a[1] if a[1] > b[1] else b[1]
            if edge_top < top < edge_bottom <= bottom:
                return False
            if edge_bottom > bottom > edge_top >= top:
                return False
            if edge_top <= top and edge_bottom >= bottom:
                return False
            
        # horizontal edge
        if a[1] == b[1] and top < a[1] < bottom:
            edge_left = a[0] if a[0] < b[0] else b[0]
            edge_right = a[0] if a[0] > b[0] else b[0]
            if edge_left < left < edge_right <= right:
                return False
            if edge_right > right > edge_left >= left:
                return False
            if edge_left <= left and edge_right >= right:
                return False

    return True

edges = list(zip(lines, lines[1:] + lines[:1]))
max = 0
for combo in combinations(lines, 2):
    if not is_inside(combo[0], combo[1], edges): 
        continue
    area = (abs(combo[0][0] - combo[1][0]) + 1) * (abs(combo[0][1] - combo[1][1]) + 1)
    if area > max:
        max = area
print(max)
