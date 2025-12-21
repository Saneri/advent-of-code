with open('input.txt', 'r') as file:
    lines = [l.strip().split(':') for l in file.readlines()]

portal_map = {}
for line in lines:
    start = line[0].strip()
    ends = line[1].strip().split()
    portal_map[start] = ends

cache = {}

def travel_portals(in_portal: str, visited: tuple[bool, bool]) -> int:
    if (in_portal, visited) in cache:
        return cache[(in_portal, visited)]
    match in_portal:
        case 'out':
            return 1 if visited == (True, True) else 0
        case 'dac':
            visited = (True, visited[1])
        case 'fft':
            visited = (visited[0], True)

    out_portals = portal_map[in_portal]
    total = 0
    for out_portal in out_portals:
        total += travel_portals(out_portal, visited)
    cache[(in_portal, visited)] = total
    return total

print(travel_portals('svr', (False, False)))