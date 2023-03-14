let pricesTable = document.getElementById("prices");
let sellsTable = document.getElementById("sells");
let ownsTable = document.getElementById("owns");
let oracleResultsTable = document.getElementById("oracle-result");

function addChild(element, cont, cClass) {
  let node = document.createElement("div");
  if (cClass !== "") node.classList.add(cClass);
  node.innerHTML = cont;
  element.appendChild(node);
}

function loadPrices() {
  fetch("/prices")
    .then((response) => response.json())
    .then((data) => {
      pricesTable.innerHTML = "";
      for (const [key, value] of Object.entries(data)) {
        let row = pricesTable.insertRow(-1);
        let itemCell = row.insertCell(0);
        itemCell.classList.add('w-25');
        let rangeCell = row.insertCell(1);

        itemCell.innerHTML = key;
        rangeCell.innerHTML = value[0] + " - " + value[1];
      }
    });
}

function loadSells() {
  fetch("/sells")
    .then((response) => response.json())
    .then((data) => {
      console.log(data);
      sellsTable.innerHTML = "";
      for (let sell of data) {
        let row = sellsTable.insertRow(-1);
        let itemCell = row.insertCell(0);
        itemCell.classList.add('w-25');
        let priceCell = row.insertCell(1);

        itemCell.innerHTML = sell.name;
        priceCell.innerHTML = sell.price;
      }
    });
}

function loadOwns() {
  fetch("/owns")
    .then((response) => response.json())
    .then((data) => {
      ownsTable.innerHTML = "";
      for (let own of data) {
        let row = ownsTable.insertRow(-1);
        let itemCell = row.insertCell(0);
        itemCell.classList.add('w-25');
        let ownerCell = row.insertCell(1);

        itemCell.innerHTML = own.name;
        fetch("/alias/" + own.owner)
          .then((response) => response.text())
          .then((data) => {
            ownerCell.innerHTML = data;
          });
      }
    });
}

function loadOracleResults() {
  fetch("/oracle-result")
    .then((response) => response.json())
    .then((data) => {
      console.log(data);
      oracleResultsTable.innerHTML = "";
      // reconstruct list & sort data by timestamp
      var results = []
      for (let sell of data) {
        // sell.timestamp = new Date(sell.timestamp);
        let result = sell.result.split("::::")[0];
        let timeStr = sell.result.split("::::")[1] || ""; //handle undefined and wrong format string
        let timestamp = parseInt(timeStr.replace("\n", "").replace("::", ""));
        results.push({
          request: sell.request,
          result: result,
          timestamp: timestamp
        });
      }
      results.sort((a,b) => b.timestamp - a.timestamp);
      for (let result of results) {
        let row = oracleResultsTable.insertRow(-1);
        let itemCell = row.insertCell(0);
        itemCell.classList.add('w-25');
        itemCell.innerHTML = result.request.slice(0,24) + "..." + result.request.slice(-8);
        let resultCell = row.insertCell(1);
        resultCell.innerHTML = result.result;
        let timestampCell = row.insertCell(2);
        let dateObject = new Date(result.timestamp*1000);
        timestampCell.innerHTML = dateObject.toLocaleString();
      }
    });
}

function reload() {
  loadPrices();
  loadSells();
  loadOwns();
  loadOracleResults();
}

//setInterval(reload, 2000);
