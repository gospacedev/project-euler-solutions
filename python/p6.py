def main():
    sum_of_squares = []
    lst = []

    for x in range(101):

        lst.append(x)

        sum_of_squares.append(pow(x, 2))

    square_of_sums = pow(sum(lst), 2)

    ans = square_of_sums - sum(sum_of_squares)

    print(ans)

if __name__ == "__main__":
    main()