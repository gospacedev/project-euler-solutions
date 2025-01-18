def factorial(n):
	if n == 0:
		return 1
	return n * factorial(n-1)

def main():
	k = 20
	n = 40

	ans = factorial(n) / (factorial(n-k) * factorial(k))

	print(ans)


if __name__ == "__main__":
	main()