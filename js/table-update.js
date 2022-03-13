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
        .then(function (response) {
            return response.json();
        }).then(function (json) {
            ESPObject = json;
        });

}

function updateTable() {
    console.log("ESPObject: ", ESPObject)
    let table = document.getElementById('mainTable');
    let arrays = Object.values(ESPObject);

    // TESTING ONLY
  
        if (arrays.includes(null)) {
            arrays = Object.values(dummyData)
        }


    for (let i = 0; i < arrays.length; i++) {
        let innerArrayLength = arrays[i].length;
        for (let j = 0; j < innerArrayLength; j++) {
            if(arrays[i][j]) {
                table.rows[i + 1].cells[j + 1].innerHTML = "full! 😄";
            } else {
                table.rows[i + 1].cells[j + 1].innerHTML = "empty! 😞";
            }
        }
    }
}

function main() {
    setInterval(() => {
        getESPData();
        updateTable();
    }, 10000)
}

main();
