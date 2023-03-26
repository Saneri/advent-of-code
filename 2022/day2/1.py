data = [line.strip().split(' ') for line in open('input.txt', 'r')]

base_score = {
    "X": 1,
    "Y": 2,
    "Z": 3
}

result_score = {
    "win": 6,
    "draw": 3,
    "lose": 0
}

winning_move = {
    "A": "Y",
    "B": "Z",
    "C": "X"
}

draw_move = {
    "A": "X",
    "B": "Y",
    "C": "Z"
}

score = 0
for round in data:
    score += base_score[round[1]]
    if winning_move[round[0]] is round[1]:
        score += result_score['win']
    elif draw_move[round[0]] is round[1]:
        score += result_score["draw"]

print(score)
