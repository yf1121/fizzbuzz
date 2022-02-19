for(var i = 1; i < 41; i++){
  var s = "";
  if(!(i % 3) || String(i).match(/3/)){
    s += "アホ";
  }
  if(!(i % 5)){
    s += "ナルシスト";
  }
  s += i
  console.log("%s", s);
}
