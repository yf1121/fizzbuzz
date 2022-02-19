for i in range(40):
    # １から40まで
    i += 1
    s = ""
    if i % 3 == 0 or "3" in str(i):
        s += "アホ"
    if i % 5 == 0:
        s += "ナルシスト"
    s += str(i)
    print(s)
