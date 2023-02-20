let demoContent = document.getElementById('progress');
let transferButton = document.getElementById('butt');

transferButton.addEventListener('click', function() {
  let xhr = new XMLHttpRequest();
  let amount = document.getElementById('amount').value;
  if (document.getElementById('tm_to_fab').checked) {
    xhr.open("GET", "/transfer/tm-to-fab/" + amount);
  }
  else {
    xhr.open("GET", "/transfer/fab-to-tm/" + amount);
  }
  xhr.send();
});

function reloadProgress() {
  let xhr = new XMLHttpRequest();
  xhr.onload = function() {
    if (xhr.status === 200) {
      demoContent.innerHTML = xhr.responseText;
      demoContent.scrollTop = demoContent.scrollHeight;
    }
  }
  xhr.open("GET", "/plain/test-tx", true);
  xhr.send();
}

let ibc0Bal = document.getElementById('ibc0-bal');
let ibc1Bal = document.getElementById('ibc1-bal');

//setInterval(reloadProgress, 1000);
function reloadBalance() {
  let xhr0 = new XMLHttpRequest();
  xhr0.onload = function() {
    if (xhr0.status === 200) {
      ibc0Bal.innerHTML = xhr0.responseText;
    }
  }
  xhr0.open("GET", "/query-balance/ibc0", true);
  xhr0.send();

  let xhr1 = new XMLHttpRequest();
  xhr1.onload = function() {
    if (xhr1.status === 200) {
      ibc1Bal.innerHTML = xhr1.responseText;
    }
  }
  xhr1.open("GET", "/query-balance/ibc1", true);
  xhr1.send();
}

setInterval(reloadBalance, 1000);
