def main():
    sum_of_squares, square_of_sums = 0

    for i in range(1, 100):
        sum_of_squares += pow(i, 2)
        square_of_sums += i

    square_of_sums = pow(square_of_sums, 2)

    ans = square_of_sums - sum(sum_of_squares)

    print(ans)

if __name__ == "__main__":
    main()