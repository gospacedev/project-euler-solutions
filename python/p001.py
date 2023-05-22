def main():
    place = []

    for x in range(1000):
        if x % 3 == 0 or x % 5 == 0:
            place.append(x)

    print(sum(place))

if __name__ == "__main__":
    main()