
minimum = 372304
maximum = 847060

def password_count(minimum, maximum):
    valid_codes = 0
    for code_int in range(minimum, maximum +1):
        code = str(code_int)
        if len(code) == 6:
            double_found = False
            change = True
            double = False
            ascending = True
            last_digit = int(code[0])

            for digit_str in code[1:]:
                digit = int(digit_str)
                if digit < last_digit:
                    ascending = False
                if digit != last_digit:
                    change = True
                    if double == True:
                        double_found = True
                    double = False
                else:
                    if change:
                        double = True
                    else:
                        double = False
                    change = False

                last_digit = digit

            if double:
                double_found = True

            if double_found and ascending:
                valid_codes += 1
                #print(code)
    #print(valid_codes)
    return(valid_codes)

assert password_count(128392, 643281) == 1390

print(password_count(minimum, maximum))
