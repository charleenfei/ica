let pricesTable = document.getElementById("prices");
let sellsTable = document.getElementById("sells");
let ownsTable = document.getElementById("owns");

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

function reload() {
  loadPrices();
  loadSells();
  loadOwns();
}

//setInterval(reload, 2000);
