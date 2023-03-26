data = [[int(num) for num in line.strip()] for line in open('input.txt', 'r')]

def check(num, arr):
    for number in arr:
        if number >= num:
            return False
    return True
        
def is_visible(trees, x, y):
    tree = trees[y][x]
    horizontal = trees[y]
    vertical = [row[x] for row in trees]
    left = check(tree, horizontal[:x])
    right = check(tree, horizontal[x+1:])
    up = check(tree, vertical[:y])
    down = check(tree, vertical[y+1:])

    return left or right or up or down

def count(num, arr):
    count = 0
    for number in arr:
        count += 1
        if number >= num:
            return count
    return count

def scenic_score(trees, x, y):
    tree = trees[y][x]
    horizontal = trees[y]
    vertical = [row[x] for row in trees]
    left = count(tree, horizontal[:x][::-1])
    right = count(tree, horizontal[x+1:])
    up = count(tree, vertical[:y][::-1])
    down = count(tree, vertical[y+1:])
    return left * right * up * down
    
visible_trees = 0
max_scenic = 0
for j in range(0, len(data)):
    for i in range(0, len(data[0])):
        visible_trees += is_visible(data, i, j)
        scenic = scenic_score(data, i, j)
        if (scenic > max_scenic):
            max_scenic = scenic
    
print(visible_trees)
print(max_scenic)
