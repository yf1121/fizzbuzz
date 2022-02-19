for(var i = 1; i < 101; i++){
  var s = "";
  if(!(i % 3)){
    s += "Fizz";
  }
  if(!(i % 5)){
    s += "Buzz";
  }
  if(s == ""){
    s = i.toString()
  }
  console.log("%s", s);
}
