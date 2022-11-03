import './App.css';
import GateInfo from './components/GateComponent'
import StexInfo from './components/StexComponent';
import BootstrapTable from 'react-bootstrap-table-next';
import 'bootstrap/dist/css/bootstrap.min.css';
import 'react-bootstrap-table-next/dist/react-bootstrap-table2.min.css';

const WalletCaptionElement = () => <h3 className='Wallet-caption-element'>Wallet</h3>;


const App = () => {
  return  <div className="App">
  <div className='Wallet-panel'>
    <BootstrapTable bootstrap4 
    keyField="id"
    data={ [{gateCurrentAmount: 200, stexCurrentAmount: 100}] } 
    caption={<WalletCaptionElement/>} 
    rowStyle= {{border: '2px solid #E06C75'}}
    columns={ [{dataField:"gateCurrentAmount", 
                text:"Gate Curent Amount",
                style: {color:'#E06C75'},
                headerStyle: {
                    color: '#E06C75',
                    border: '2px solid #E06C75'
                  }}, {dataField:"stexCurrentAmount", 
                  text:"Stex Current Amount",
                  style: {color:'#E06C75'},
                  headerStyle: {
                      color: '#E06C75',
                      border: '2px solid #E06C75'
                    }}] } />
  </div>
  
  <div className="Platform-pane">
      <div className="Gate-table">
        <GateInfo/>
      </div>
      <div className="Stex-table">
        <StexInfo />
      </div>
  </div>

</div>
};

export default App;