"use strict";

function displayESPData() {
    fetch('/SendToJS')
    .then(response => response.json())
    .then(data => console.log(data));
  
}

displayESPData();

