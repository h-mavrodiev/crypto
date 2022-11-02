import './App.css';
import BootstrapTable from 'react-bootstrap-table-next';
import React, { useState } from "react";


const gateData = [ {AskPrice:1000, AskAmount: 2, BidPrice: 900, BidAmount: 2}];
const stexData = [ {askPrice:1000, askAmount: 2, bidPrice: 900, bidAmount: 2}];

const columns = [{
  dataField: 'AskPrice',
  text: 'Ask Price'
}, {
  dataField: 'AskAmount',
  text: 'Ask Amount'
}, {
  dataField: 'BidPrice',
  text: 'Bid Price'
}, {
  dataField: 'BidAmount',
  text: 'Bid Amount'
}];

const GateCaptionElement = () => <h3 className='Gate-caption-element'>Gate</h3>;

const StexCaptionElement = () => <h3 className='Stex-caption-element'>Stex</h3>;

const WalletHeaderElement = () => <h3 className='Wallet-header-element'>Wallet Info</h3>;

const WalletCaptionElement = () => <h3 className='Wallet-caption-element'>Wallet</h3>;

const PlatformHeaderElement = () => <h3 className='Platform-header-element'>Platform Info</h3>;

function App() {

  return (
    <div className="App">
      <div className='Wallet-header'>
        <WalletHeaderElement/>
        <body>
          <BootstrapTable bootstrap4 keyField="id" data={ [{gateCurrentAmount: 200, stexCurrentAmount: 100}] } caption={<WalletCaptionElement/>} columns={ [{dataField:"gateCurrentAmount", text:"Gate Curent Amount"}, {dataField:"stexCurrentAmount", text:"Stex Current Amount"}] } />
        </body>
      </div>
      
      <div className="Platform-header">
        <PlatformHeaderElement/>
        <body className='Platform-info-body'>
          <div className="Gate-table">
            <BootstrapTable bootstrap4 keyField="id" data={ gateData } caption={<GateCaptionElement/>} columns={ columns } />
          </div>
          <div className="Stex-table">
            <BootstrapTable bootstrap4 keyField="id" data={ stexData } caption={<StexCaptionElement />} columns={ columns } />
          </div>
        </body>

      </div>
    
    </div>
  );
}

export default App;
