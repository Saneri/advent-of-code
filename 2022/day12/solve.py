data = [[char for char in line.strip()] for line in open('input.txt', 'r')]

def get_adjecent_squares(position, map):
    def check(pos_a, pos_b, map, arr, alph):
        # check if out of bounds
        if not 0 <= pos_a < len(map) or not 0 <= pos_b < len(map[0]):
            return
        alphabet = map[pos_a][pos_b]
        if alphabet == 'E':
            alphabet = 'z'
        are_adjecent_alphabets = ord(alphabet)-ord(alph) <= 1
        if are_adjecent_alphabets:
            arr.append((pos_a,pos_b))

    current_alphabet = map[position[0]][position[1]]
    availble_positions = []

    check(position[0]+1, position[1], map, availble_positions, current_alphabet)
    check(position[0], position[1]+1, map, availble_positions, current_alphabet)
    check(position[0]-1, position[1], map, availble_positions, current_alphabet)
    check(position[0], position[1]-1, map, availble_positions, current_alphabet)
    return availble_positions

def dijkstras(start, map):
    # initiate distance to infinity
    dist = [[float('inf') for _ in range(len(map[0]))] for _ in range(len(map))]
    dist[start[0]][start[1]] = 0
    checked_graphs = set()
    discovered_graphs = [start]

    try:
        while (True):
            current_graph = discovered_graphs.pop(0)
            for graph in get_adjecent_squares(current_graph, map):
                if not graph in checked_graphs:
                    if dist[graph[0]][graph[1]] == float('inf'):
                        dist[graph[0]][graph[1]] = dist[current_graph[0]][current_graph[1]]+1
                    discovered_graphs.append(graph)
                    checked_graphs.add(graph)
                
                if map[graph[0]][graph[1]] == 'E':
                    return dist[graph[0]][graph[1]]
    # no path found
    except:
        return float('inf')

def starting_locations(map):
    starts = []
    for j in range(len(map)):
        for i in range(len(map[0])):
            if data[j][i] == 'a':
                starts.append((j,i))
            elif data[j][i] == 'S':
                data[j][i] = 'a'
                starts.append((j, i))
    return starts

current_locations = starting_locations(data)
distances = [dijkstras(start, data) for start in current_locations]
print(min(distances))
