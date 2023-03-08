let pricesElement = document.getElementById("prices");
let sellsElement = document.getElementById("sells");
let ownsElement = document.getElementById("owns");

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
      pricesElement.innerHTML = "";
      for (const [key, value] of Object.entries(data)) {
        pricesElement.innerHTML =
          "<tr>" + "<td class='w-100'>" + key + " : " + value[0] + " - ";
        value[1] + "</td>" + "</tr>";
      }
    });
}

function loadSells() {
  fetch("/sells")
    .then((response) => response.json())
    .then((data) => {
      sellsElement.innerHTML = "";
      for (let sell of data) {
        "<tr>" +
          "<td class='w-100'>" +
          sell.name +
          " : on sale for : " +
          sell.price +
          "</td>" +
          "</tr>";
      }
    });
}

function loadOwns() {
  fetch("/owns")
    .then((response) => response.json())
    .then((data) => {
      ownsElement.innerHTML = "";
      for (let own of data) {
        let subjectNode = document.createElement("tr");

        let aliasNode1 = document.createElement("td");
        aliasNode1.classList.add("w50");
        aliasNode1.innerHTML = own.name + " : is owned by : ";
        subjectNode.appendChild(aliasNode1);

        let aliasNode2 = document.createElement("td");
        aliasNode2.classList.add("w50");
        aliasNode2.classList.add("brown");
        fetch("/alias/" + own.owner)
          .then((response) => response.text())
          .then((data) => {
            aliasNode2.innerHTML = data;
          });
        subjectNode.appendChild(aliasNode2);

        ownsElement.appendChild(subjectNode);
      }
    });
}

function reload() {
  loadPrices();
  loadSells();
  loadOwns();
}

//setInterval(reload, 2000);
