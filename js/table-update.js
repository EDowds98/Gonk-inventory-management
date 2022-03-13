"use strict";

 //global variables
let ESPObject;
const data;

function displayESPData() {
    fetch('/SendToJS')
    .then(response => {
        console.log('Response: ', response)
        console.log(response.json())
    })
    .then(data => {
        ESPObject = JSON.parse(data)
        console.log("parsed json: ", ESPObject)
    });
  
}

displayESPData();

