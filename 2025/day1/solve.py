with open('input.txt', 'r') as file:
    input: list[str] = [l.strip() for l in file.readlines()]

dial: int = 50
total: int = 0

for l in input:
    dir, clicks = l[0], int(l[1:])
    oldDial = dial
    match dir:
        case 'R':
            dial += clicks
            laps = dial // 100 - oldDial // 100
            total += abs(laps)
        case 'L':
            dial -= clicks
            laps = dial // (-100) - oldDial // (-100)
            total += abs(laps)
        case _:
            raise Exception("unknown direction", dir)

print(total)