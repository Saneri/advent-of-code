with open('input.txt', 'r') as file:
    lines = [l.strip().split() for l in file.readlines()]

total = 0
for line in lines:
    goal = tuple(map(int, line[-1][1:-1].split(',')))
    start = [0] * len(goal)
    buttons = [tuple(map(int, l[1:-1].split(','))) for l in line[1:-1]]
    visited = set()

    def push_button(states: set[list[int]], count: int) -> int:
        next_states = set()
        for state in states:
            if state == goal:
                return count
            
            if state in visited:
                continue
            visited.add(state)
            
            for button in buttons:
                new_state = list(state)
                for action in button:
                    new_state[action] += 1
                if all(new_state[i] <= goal[i] for i in range(len(goal))):
                    next_states.add(tuple(new_state))
        return push_button(next_states, count + 1)
    
    total += push_button(set([tuple(start)]), 0)
print(total)

# this code was too inefficient for part 2 so I solved it with linear equations else where