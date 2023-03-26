from copy import deepcopy

data = open('input.txt', 'r').read().split('\n\n')

starting_stacks = data[0].split('\n')[::-1][1:]
procedures = data[1].split('\n')

crate_poisitions = [*range(1, len(starting_stacks[0]), 4)]
raw_stacks = [[line[x] for x in crate_poisitions] for line in starting_stacks]
# transpose the matrix and remove all empty values
stacks = [[x for x in list(line) if x != ' '] for line in zip(*raw_stacks)]
stacksB = deepcopy(stacks)

for procedure in procedures:
    [_, move, _, fromm, _, to] = procedure.split(' ')

    cache_list = []
    for _ in range(0, int(move)):
        cache_list.append(stacksB[int(fromm)-1].pop())
        stacks[int(to)-1].append(stacks[int(fromm)-1].pop())
    stacksB[int(to)-1] += cache_list[::-1]


# get all top crates and join to a string
def printTopOfStack(stacks):
    print(''.join([stack.pop() for stack in stacks]))

printTopOfStack(stacks)
printTopOfStack(stacksB)
