one = two = six = ten = and_word = 3
three = seven = eight = forty = fifty = sixty = 5
four = five = nine = 4
eleven = twelve = twenty = thirty = eighty = ninety = 6
thirteen = fourteen = eighteen = nineteen = thousand = 8
fifteen = sixteen = seventy = hundred = 7
seventeen = 9

number_letter_count = 0

first_digits = [one, two, three, four, five, six, seven, eight, nine]

second_teen_digits = [ten, eleven, twelve, thirteen,
                      fourteen, fifteen, sixteen, seventeen, eighteen, nineteen]

second_digits = [twenty, thirty, forty, fifty, sixty, seventy, eighty, ninety]

# 1 - 19
number_letter_count += sum(first_digits) + sum(second_teen_digits) 

# 20 - 99
for i in second_digits:
    number_letter_count += i
    for j in first_digits:
        number_letter_count += i + j

# 100 - 999
for hundreds in first_digits:
    number_letter_count += hundreds + hundred

    for t in first_digits:
        number_letter_count += hundreds + hundred + and_word + t

    for k in second_teen_digits:
        number_letter_count += hundreds + hundred + and_word + k

    for i in second_digits:
        number_letter_count += hundreds + hundred + and_word + i

        for j in first_digits:
            number_letter_count += hundreds + hundred + and_word + i + j


number_letter_count += one + thousand

print(number_letter_count)
