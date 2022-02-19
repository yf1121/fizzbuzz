1.upto(40) {|i|
  if i % 3 == 0 || i.to_s.include?("3")
    print("アホ")
  end
  if i % 5 == 0
    print("ナルシスト")
  end
  print(i.to_s + "\n")
}
