"use strict";

let ESPObject;
// dummy data because no testbench
const dummyData = {
    Module1: [false, false, false, false, false, false, false, false],
    Module2: [false, false, false, false, false, false, false, false],
    Module3: [false, false, false, false, false, false, false, false],
    Module4: [false, false, false, false, false, false, false, false],
    Module5: [false, false, false, false, false, false, false, false],
    Module6: [false, false, false, false, false, false, false, false],
    Module7: [false, false, false, false, false, false, false, false],
    Module8: [false, false, false, false, false, false, false, false]
}
function displayESPData() {
    fetch('/SendToJS')
    .then(response => {
        console.log('Response: ', response)
        console.log(response.json())
    })
    .then(data => {
        try {
            ESPObject = JSON.parse(data)
        } catch(error) {
            ESPObject = dummyData;
        }
        console.log("parsed json: ", ESPObject)
    });
  
}

displayESPData();

