let infoElement = document.getElementById('info');
let addressElement = document.getElementById('address');
let nameElement = document.getElementById('name');
let chainElement = document.getElementById('chain');
let rCom = document.getElementById('r-com');
let rAddr = document.getElementById('r-addr');
let rChain = document.getElementById('r-chain');
let rItem = document.getElementById('r-item');
let rPrice = document.getElementById('r-price');

function addChild(element, cont, cClass) {
  let node = document.createElement('div');
  if (cClass !== '') node.classList.add(cClass);
  node.innerHTML = cont;
  element.appendChild(node);
}

function register() {
  fetch('/register/' + addressElement.value +'/' + nameElement.value + '/' + chainElement.value)
    .then((response) => response.text())
    .then((ica) => {
      infoElement.innerHTML = '<tr><td>Name: ' + nameElement.value + "</td></tr>"  
      + '<tr><td>Address: ' + addressElement.value + "</td></tr>"
      + '<tr><td>Chain: ' + chainElement.value + "</td></tr>"
      + '<tr><td>ICA address: ' + ica + "</td></tr>"
    });
}

function run() {
  fetch('/run/' + rCom.value + '/' + rAddr.value + '/' + rChain.value + '/' + rItem.value + '/' + rPrice.value);
}

//setInterval(reload, 2000);
