function getValue (id) {
    text = document.getElementById("commands").value; //value of the text input
    textdva = document.getElementById("commandsdva").value;
fetch('http://127.0.0.1/post', {
    method: 'POST',
    headers: {
        'Accept': 'application/json',
        'Content-Type': 'application/x-www-form-urlencoded'
    },
    body: JSON.stringify({ "name": text  })
})
   .then(response => response.json())
   .then(response => console.log(JSON.stringify(response)))
}