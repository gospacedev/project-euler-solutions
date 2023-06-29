def find_pythagorean_triplets(n):
  for a in range(1, n//3):
    for b in range(a + 1, n//2):
      c = n - a - b
      if a*a + b*b == c*c:
          return a * b * c


def main():
  print(find_pythagorean_triplets(1000))


if __name__ == "__main__":
    main()