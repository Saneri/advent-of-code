data = [line.strip() for line in open('input.txt', 'r')]

def convertCharToPriority(char):
    if char.islower():
        return ord(char) - 96
    return ord(char) - 38
 
assert(convertCharToPriority('a') == 1)
assert(convertCharToPriority('z') == 26)
assert(convertCharToPriority('A') == 27)

prioritySum = 0
for rucksack in data:
    half = int(len(rucksack)/2)
    first = rucksack[:half]
    second = rucksack[half:]
    commonChar = list(set(first)&set(second))[0]
    prioritySum += convertCharToPriority(commonChar)

threeElfSum = 0
myit = iter(data)
for _ in range(int(len(data)/3)):
    commonChar = list(set(next(myit))&set(next(myit))&set(next(myit)))[0];
    threeElfSum += convertCharToPriority(commonChar)

print(prioritySum)
print(threeElfSum)
