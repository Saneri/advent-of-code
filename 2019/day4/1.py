
minimum = 372304
maximum = 847060

def password_count(minimum, maximum):
    valid_codes = 0
    for code_int in range(minimum, maximum +1):
        code = str(code_int)
        if len(code) == 6:
            double_found = False
            ascending = True
            last_digit = int(code[0])

            for digit_str in code[1:]:
                digit = int(digit_str)
                if digit < last_digit:
                    ascending = False
                if digit == last_digit:
                    double_found = True
                last_digit = digit

            if double_found and ascending:
                valid_codes += 1
    print(valid_codes)
    return(valid_codes)

assert password_count(331, 335) == 0
assert password_count(128392, 643281) == 2050

print(password_count(minimum, maximum))
