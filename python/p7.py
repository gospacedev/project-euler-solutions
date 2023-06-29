def is_prime(number):
    if number < 2:
        return False

    for i in range(2, int(number ** 0.5) + 1):
        if number % i == 0:
            return False
        
    return True
        
def main():
    primeCount = 0
    i = 0
    while True:
        i += 1
        if is_prime(i):
            primeCount += 1
        if primeCount == 10001:
            print(i)
            break

if __name__ == "__main__":
    main()