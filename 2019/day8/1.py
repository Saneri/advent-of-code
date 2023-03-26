import math

with open('input.txt', 'r') as file:
    image_data = file.readline()

#image_data = '123456789012'

WIDE = 25
TALL = 6

line_amount = int(len(image_data) / WIDE / TALL)

i = 0
layer = 1
fewest_zeros_layer = None
amount_of_zeros = math.inf

for _ in range(0, line_amount):
    current_zeros = 0
    for _ in range(0, TALL):
        current_zeros += image_data[i:i+WIDE].count('0')
        i += WIDE
    if current_zeros < amount_of_zeros:
        amount_of_zeros = current_zeros
        fewest_zeros_layer = i - WIDE * TALL
    layer += 1

i = fewest_zeros_layer
line = ''
for _ in range(0, TALL):
    line += image_data[i:i+WIDE]
    i += WIDE

print(line)
print(line.count('1') * line.count('2'))

