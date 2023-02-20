let infoElement = document.getElementById('info');
let addressElement = document.getElementById('address');
let nameElement = document.getElementById('name');

function addChild(element, cont, cClass) {
  let node = document.createElement('div');
  if (cClass !== '') node.classList.add(cClass);
  node.innerHTML = cont;
  element.appendChild(node);
}

function register() {
  if (infoElement.innerHTML !== '') return;
  fetch('/register/' + addressElement.value +'/' + nameElement.value)
    .then((response) => response.text())
    .then((ica) => {
      addChild(infoElement, 'Name: ' + nameElement.value, 'info');
      addChild(infoElement, 'Address: ' + addressElement.value, 'info');
      addChild(infoElement, 'ICA address: ' + ica, 'info');
    });
}

//setInterval(reload, 2000);
