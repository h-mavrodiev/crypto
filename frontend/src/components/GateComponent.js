import { useEffect, useState } from 'react'
import BootstrapTable from 'react-bootstrap-table-next';
import 'bootstrap/dist/css/bootstrap.min.css';
import 'react-bootstrap-table-next/dist/react-bootstrap-table2.min.css';

const GateCaptionElement = () => <h3 className='Gate-caption-element'>Gate</h3>;

const sleep = ms => new Promise(
  resolve => setTimeout(resolve, ms)
);

const columns = [{
    dataField: 'IWannaBuyFor',
    text: '$',
    style: {color:'#c678dd'},
    headerStyle: {
      color: '#c678dd',
      border: '2px solid #c678dd'
    }
  },{
    dataField: 'TheySellForWeightedUSD',
    text: 'Price per $ (W)',
    style: {color:'#c678dd'},
    headerStyle: {
      color: '#c678dd',
      border: '2px solid #c678dd'
    }
  },  {
    dataField: 'ICanBuy',
    text: 'I Can Buy',
    style: {color:'#c678dd'},
    headerStyle: {
      color: '#c678dd',
      border: '2px solid #c678dd'
    }
  }, {
    dataField: 'TheySellForWeighted',
    text: 'They Sell For (W)',
    style: {color:'#c678dd'},
    headerStyle: {
      color: '#c678dd',
      border: '2px solid #c678dd'
    }
  }, {
    dataField: 'ICanSellFromStex',
    text: 'I Can Sell From Stex',
    style: {color:'#c678dd'},
    headerStyle: {
      color: '#c678dd',
      border: '2px solid #c678dd'
    }
  },{
    dataField: 'ICanSellFromStexForWeighted',
    text: 'I Can Sell From Stex For',
    style: {color:'#c678dd'},
    headerStyle: {
      color: '#c678dd',
      border: '2px solid #c678dd'
    }
  }];

// Use effect - effect hook
function GateInfo(){
    async function loadGateInfo() {
      await sleep(1000)
      fetch("http://localhost:8080/gate")
      .then((response) => response.json())
      .then((data) => setGateInfo([data]));
    }

    const [gateInfo, setGateInfo] = useState([]);
    useEffect(() => {loadGateInfo()});

    return <BootstrapTable 
    className='Gate-Table'
    bootstrap4 
    keyField="AskFixedUSDDemand" 
    data={ gateInfo } 
    condensed
    striped 
    caption={<GateCaptionElement/>}
    columns={ columns }
    rowStyle= {{border: '2px solid #c678dd'}}/>
}

export default GateInfo;