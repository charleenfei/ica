let infoElement = document.getElementById('info');
let addressElement = document.getElementById('address');
let nameElement = document.getElementById('name');
let chainElement = document.getElementById('chain');
let rCom = document.getElementById('r-com');
let rItem = document.getElementById('r-item');
let rPrice = document.getElementById('r-price');

let icaTxt = document.getElementById('ica-txt');
let balTxt = document.getElementById('bal-txt');
let spinner = document.getElementById('spinner');

function register() {
  addressElement.readOnly = true;
  nameElement.readOnly = true;
  chainElement.readOnly = true;

  fetch('/register/' + addressElement.value +'/' + nameElement.value + '/' + chainElement.value)
    .then((response) => response.text())
    .then((ica) => {
      icaTxt.innerHTML = ica;
      loadBalance();
    });
}

function loadBalance() {
  spinner.classList.remove('visually-hidden');
  ica = icaTxt.innerHTML;
  fetch('/balance/' + ica)
    .then((response) => response.text())
    .then((balance) => {
      balTxt.innerHTML = balance;
    });
  spinner.classList.add('visually-hidden');
}

function run() {
  fetch('/run/' + rCom.value + '/' + addressElement.value + '/' + chainElement.value + '/' + rItem.value + '/' + rPrice.value);
}

//setInterval(reload, 2000);
