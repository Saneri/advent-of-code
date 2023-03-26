data = [line.strip().split(',') for line in open('input.txt', 'r')]

def checkOverlap(first, second, partB = False):
    def check(a, b):
        return (a[0] >= b[0] and a[0 if partB else 1] <= b[1])
    return check(first, second) or check(second, first)

assert(not checkOverlap([2,4],[6,8]))
assert(checkOverlap([6,6],[4,6]))
assert(checkOverlap([4,6],[6,6]))
assert(checkOverlap([1,1],[1,1]))
assert(checkOverlap([1,5],[2,6], True))
assert(checkOverlap([2,6],[1,5], True))

overlapTimes = 0
overlapAtAll = 0
for line in data:
    first = [int(e) for e in line[0].split('-')]
    second = [int(e) for e in line[1].split('-')]
    if checkOverlap(first, second):
        overlapTimes += 1
    if checkOverlap(first, second, True):
        overlapAtAll += 1

print(overlapTimes)
print(overlapAtAll)