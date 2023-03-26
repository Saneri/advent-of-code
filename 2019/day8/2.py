import math

with open('input.txt', 'r') as file:
    image_data = file.readline()

#image_data = '0222112222120000'

WIDE = 25
TALL = 6

image = [[2 for _ in range(WIDE)] for _ in range(TALL)]

line_amount = int(len(image_data) / WIDE / TALL)

i = 0
layer = 1
for layer in range(0, line_amount):
    for x in range(0, TALL):
        y = 0
        for index in range(i, i+WIDE):
            pixel = image_data[index]
            #print("{} : [{},{}] - {}".format(pixel, x, y, image[x][y]))
            if (pixel == '0' or pixel == '1') and image[x][y] == 2:
                image[x][y] = int(pixel)
            y += 1
            
        i += WIDE

message = ''
for line in image:
    for pixel in line:
        if pixel == 1:
            message += '#'
        elif pixel == 0:
            message += '.'
    message += '\n'

print(message)
