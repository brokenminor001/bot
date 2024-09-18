function getValue (id) {
    text = document.getElementById("name").value;
    textdva = document.getElementById("date").value;
    texttre = document.getElementById("time").value;
    textfour = document.getElementById("type").value //value of the text input
    let user = {
        PostRaspisanie: textfour,
        PostRaspisanieData: textdva,
        PostRaspisanieVremya: texttre,
        PostRaspisanieName: text
      };
      
fetch('http://127.0.0.1:8081/postraspisanie', {
    method: 'POST',
    headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/x-www-form-urlencoded'
    },
    body: JSON.stringify(user)
    
    
})


   .then(response => response.json())
   .then(response => console.log(JSON.stringify(response)))
}
function DispOne() {
  
        
fetch("http://127.0.0.1:8082/getlast")
  .then((response) => response.text())
  .then((json) => document.getElementById("DispLast").innerHTML = json);
     
fetch("http://127.0.0.1:8082/getbeforelast")
  .then((response) => response.text())
  .then((json) => document.getElementById("DispBeforeLast").innerHTML = json);
fetch("http://127.0.0.1:8082/getbeforebeforelast")
  .then((response) => response.text())
  .then((json) => document.getElementById("DispBeforeBeforeLast").innerHTML = json);
     
  }
  