def is_prime(number):
    if number < 2:
        return False

    for i in range(2, int(number ** 0.5) + 1):
        if number % i == 0:
            return False
        
    return True

def main():
    ans = 0
    for i in range(1, 2000000):
        if is_prime(i):
           ans += i

    print(ans)

if __name__ == "__main__":
    main()