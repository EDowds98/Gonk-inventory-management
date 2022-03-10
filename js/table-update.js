"use strict";

function displayESPData() {
    fetch('/SendToJS')
    .then(response => {
        console.log('Response: ', response)
        response.json()
    })
    .then(data => console.log(JSON.stringify(data)));
  
}

displayESPData();

