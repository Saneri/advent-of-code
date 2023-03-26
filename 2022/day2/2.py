data = [line.strip().split(' ') for line in open('input.txt', 'r')]

result_score = {
    "X": 0,
    "Y": 3,
    "Z": 6
}

moves = {
    "A": {
        "X": 3,
        "Y": 1,
        "Z": 2
    },
    "B": {
        "X": 1,
        "Y": 2,
        "Z": 3
    },
    "C": {
        "X": 2,
        "Y": 3,
        "Z": 1
    }
}

score = 0
for round in data:
    score += result_score[round[1]]
    score += moves[round[0]][round[1]]

print(score)
