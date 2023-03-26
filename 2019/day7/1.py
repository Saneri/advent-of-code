from intcode import process_code

from itertools import permutations

phase_settings = [0,1,2,3,4]

def amplifier_row_output(amp_seq, code):
    amp_input = 0
    for amp in amp_seq:
        amp_input, _, _ = process_code(code, [amp, amp_input], True)

    return amp_input

def get_biggest_output(amp_seq, code):
    max_output = 0
    for seq in permutations(amp_seq, len(amp_seq)):
        output = amplifier_row_output(seq, code)
        if output > max_output:
            max_output = output
    return max_output
    


code = []
with open('input.txt', 'r') as file:
    firstline = file.readline().split(',')
    for code_str in firstline:
        code.append(int(code_str))


print(get_biggest_output(phase_settings, code))
assert get_biggest_output(phase_settings, [3,15,3,16,1002,16,10,16,1,16,15,15,4,15,99,0,0]) == 43210
assert get_biggest_output(phase_settings, [3,23,3,24,1002,24,10,24,1002,23,-1,23, 101,5,23,23,1,24,23,23,4,23,99,0,0]) == 54321
assert get_biggest_output(phase_settings, [3,31,3,32,1002,32,10,32,1001,31,-2,31,1007,31,0,33, 1002,33,7,33,1,33,31,31,1,32,31,31,4,31,99,0,0,0]) == 65210
