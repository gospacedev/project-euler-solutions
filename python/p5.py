import math

def lcm(a, b):
	return a * b // math.gcd(a, b)

def smallest_multiple(numbers):
	lcm_of_numbers = numbers[0]

	for number in numbers[1:]:
		lcm_of_numbers = lcm(lcm_of_numbers, number)
	return lcm_of_numbers


if __name__ == "__main__":
	numbers = range(1, 21)
	smallest_multiple = smallest_multiple(numbers)
	print(smallest_multiple)