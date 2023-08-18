def collatzSequenceLength(num):
    length = 1

    while num > 1:
        if num % 2 == 0:
            num = num / 2
        else:
            num = 3*num + 1

        length+=1
        

    return length

def main():
    currentLength = 0
    longestLength = 0
    ans = 0

    for i in range(1, 1000000):
        currentLength = collatzSequenceLength(i)

        if currentLength > longestLength:
            longestLength = currentLength
            ans = i

    print(ans)
    
if __name__ == "__main__":
	main()