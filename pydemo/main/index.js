let pricesElement = document.getElementById('prices');
let sellsElement = document.getElementById('sells');
let ownsElement = document.getElementById('owns');

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
        let node = document.createElement('div');
        node.classList.add('card-long');
        node.innerHTML = own.name + ': is owned by: ' + own.owner;
        ownsElement.appendChild(node);
      }
    });

}

function reload() {
  loadPrices();
  loadSells();
  loadOwns();
}

//setInterval(reload, 2000);
