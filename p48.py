def main():
    n = []
    for x in range(1, 1001):
        n.append(pow(x, x))

    print(sum(n))

if __name__ == "__main__":
    main()