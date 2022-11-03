import { useEffect, useState } from 'react'
import BootstrapTable from 'react-bootstrap-table-next';
import 'bootstrap/dist/css/bootstrap.min.css';
import 'react-bootstrap-table-next/dist/react-bootstrap-table2.min.css';

const GateCaptionElement = () => <h3 className='Gate-caption-element'>Gate</h3>;

const sleep = ms => new Promise(
  resolve => setTimeout(resolve, ms)
);

const columns = [{
    dataField: 'AskFixedUSDDemand',
    text: 'Ask Fixed USD Demand',
    style: {color:'#c678dd'},
    headerStyle: {
      color: '#c678dd',
      border: '2px solid #c678dd'
    }
  },{
    dataField: 'AskWeightedPrice',
    text: 'Ask Weighted Price',
    style: {color:'#c678dd'},
    headerStyle: {
      color: '#c678dd',
      border: '2px solid #c678dd'
    }
  },  {
    dataField: 'AskWeightedUSDPrice',
    text: 'Ask Weighted USD Price',
    style: {color:'#c678dd'},
    headerStyle: {
      color: '#c678dd',
      border: '2px solid #c678dd'
    }
  }, {
    dataField: 'AskAmount',
    text: 'Ask Amount',
    style: {color:'#c678dd'},
    headerStyle: {
      color: '#c678dd',
      border: '2px solid #c678dd'
    }
  }, {
    dataField: 'BidFixedUSDDemand',
    text: 'Bid Fixed USD Demand',
    style: {color:'#c678dd'},
    headerStyle: {
      color: '#c678dd',
      border: '2px solid #c678dd'
    }
  },{
    dataField: 'BidWeightedPrice',
    text: 'Bid Weighted Price',
    style: {color:'#c678dd'},
    headerStyle: {
      color: '#c678dd',
      border: '2px solid #c678dd'
    }
  }, {
    dataField: 'BidWeightedUSDPrice',
    text: 'Bid Weighted USD Price',
    style: {color:'#c678dd'},
    headerStyle: {
      color: '#c678dd',
      border: '2px solid #c678dd'
    }
  }, {
    dataField: 'BidAmount',
    text: 'Bid Amount',
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