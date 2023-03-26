data = [[int(x) for x in x.split('\n')] for x in open("input.txt", "r").read().split('\n\n')]
ordered_sums = sorted([sum(x) for x in data])

print(ordered_sums[-1])
print(sum(ordered_sums[-3:]))
