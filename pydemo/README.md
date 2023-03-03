# Web UI for the demo

## Overview
After setting up the chains, relayers, and oracles, you can start the UI.
- In this `pydemo` folder, start the server by running `./server.py`
- There are 2 sites, status site and account site.

## Status site
- Open `localhost:5555/index.html`. Click `Reload` to get the latest status of the chains

  Example: After running `make docker-unitest` then clicking `Reload`:
  
 ![webUI](/images/WebUI.png)

## Account site
- Open `localhost:5555/account.html`

  Type the following and click `Register` to have the alias for the accounts. (`Address` is `$WALLET_1`)
  
  ![register](/images/account.png)
  
  + This will register an interchain account for the address if it does not exist (`Chain` has to be `test-1` or `test-3`)
  
- Click `Reload` on status site, it will update account name to Alice:

  ![Updated-UI](/images/Updated-UI.png)
  
- Next, `$WALLET_1` can sell a domain name like this:

  ![sell](/images/sell.png)
  
  + Reload status site to see pending sell updated

- Next, `$WALLET_2` can buy the domain:

  ![buy](/images/buy.png)
  
  + Reload status site to see update!
