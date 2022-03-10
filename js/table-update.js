"use strict";

function displayESPData() {
    fetch('/SendToJS')
    .then(response => {
        console.log('Response: ', response)
        console.log(response.json())
    })
    .then(data => console.log(JSON.stringify(data)));
  
}

displayESPData();

