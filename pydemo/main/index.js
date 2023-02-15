let pricesElement = document.getElementById('prices');
let sellsElement = document.getElementById('sells');
let ownsElement = document.getElementById('owns');

function addChild(element, cont, cClass) {
  let node = document.createElement('div');
  if (cClass !== '') node.classList.add(cClass);
  node.innerHTML = cont;
  element.appendChild(node);
}

function loadPrices() {
  fetch('/prices')
    .then((response) => response.json())
    .then((data) => { 
      pricesElement.innerHTML = '';
      for (const [key, value] of Object.entries(data)) {
        let node = document.createElement('div');
        node.classList.add('card');
        node.innerHTML = key + ': ' + value[0] + ' - ' + value[1];
        pricesElement.appendChild(node);
      }
    });
}

function loadSells() {
  fetch('/sells')
    .then((response) => response.json())
    .then((data) => {
      sellsElement.innerHTML = '';
      for (let sell of data) {
        let node = document.createElement('div');
        node.classList.add('card');
        node.innerHTML = sell.name + ': on sale for: ' + sell.price;
        sellsElement.appendChild(node);
      }
    });
}

function loadOwns() {
  fetch('/owns')
    .then((response) => response.json())
    .then((data) => {
      owns.innerHTML = '';
      for (let own of data) {
        let subjectNode = document.createElement('span');
        subjectNode.innerHTML = own.name + ': isowned by: '
        ownsElement.appendChild(subjectNode);

        let aliasNode = document.createElement('span');
        aliasNode.classList.add('brown');
        fetch('/alias/' + own.owner)
          .then((response) => response.text())
          .then((data) => {
            aliasNode.innerHTML = data;
          });

        ownsElement.appendChild(aliasNode);
      }
    });

}

function reload() {
  loadPrices();
  loadSells();
  loadOwns();
}

//setInterval(reload, 2000);
