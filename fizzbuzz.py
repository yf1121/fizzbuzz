for i in range(100):
    # １から40まで
    i += 1
    s = ""
    if i % 3 == 0:
        s += "Fizz"
    if i % 5 == 0:
        s += "Buzz"
    if s == "":
        s += str(i)
    print(s)
