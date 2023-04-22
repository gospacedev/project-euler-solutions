sum_of_squares = []
lst = []

for x in range(101):

    lst.append(x)

    sum_of_squares.append(pow(x, 2))

square_of_sums = pow(sum(lst), 2)

print(square_of_sums - sum(sum_of_squares))