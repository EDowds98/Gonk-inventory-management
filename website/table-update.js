"use strict";

function displayESPData() {
    fetch("/SendToJS", {
        body: JSON.stringify(data)
    }).then((response) => {
        response.text().then(function (data) {
            let result = JSON.parse(data);
            console.log(result)
        });
    }).catch((error) => {
        console.log(error)
    });
}

displayESPData();