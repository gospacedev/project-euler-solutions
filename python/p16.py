def add_digits(n):
  sum_of_digits = 0
  carry = 0
  while n > 0:
    digit = n % 10 + carry
    sum_of_digits += digit % 10
    carry = digit // 10
    n //= 10
  return sum_of_digits

def main():
   print(add_digits(pow(2, 1000)))

if __name__ == "__main__":
    main()
