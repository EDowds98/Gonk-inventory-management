"use strict";

let ESPObject = {};
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

function getESPData() {
    fetch('/SendToJS')
    .then(response => {
        console.log('Response: ', response);
        console.log(response.json());
        ESPObject = response.json();
    }/*)
    .then(data => {
        try { 
            ESPObject = data;
        } catch(error) {
            ESPObject = dummyData;
            console.log("couldn't parse JSON, using dummy data");
        }
        console.log("parsed json: ", ESPObject);
    }*/);
  
}

function updateTable() {
    console.log("ESPObject: ", ESPObject)
    let table = document.getElementById('mainTable');
    let arrays = Object.values(ESPObject);

    for(let i = 0; i < arrays.length; i++) {
        let innerArrayLength = arrays[i].length;
        for(let j = 0; j < innerArrayLength; j++) {
            table.rows[i+1].cells[j+1].innerHTML = arrays[i][j];
        }
    }
}

function main() {
    setInterval(()=> {
        getESPData();
        updateTable();
    }, 5000)
  }
  
main();
