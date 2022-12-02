import { useEffect, useState } from 'react'
import BootstrapTable from 'react-bootstrap-table-next';
import 'bootstrap/dist/css/bootstrap.min.css';
import 'react-bootstrap-table-next/dist/react-bootstrap-table2.min.css';

const PricesCaptionElement = () => <h3 className='Prices-caption-element'>Prices</h3>;

const sleep = ms => new Promise(
  resolve => setTimeout(resolve, ms)
);

const columns = [{
  dataField: 'Platform',
  text: 'Platform',
  style: {color:'#c678dd'},
  headerStyle: {
    color: '#c678dd',
    border: '2px solid #c678dd'
  }
},{
    dataField: 'Sells',
    text: 'Sells',
    style: {color:'#c678dd'},
    headerStyle: {
      color: '#c678dd',
      border: '2px solid #c678dd'
    }
  },{
    dataField: 'SellsVolume',
    text: 'Sells Volume',
    style: {color:'#c678dd'},
    headerStyle: {
      color: '#c678dd',
      border: '2px solid #c678dd'
    }
  },  {
    dataField: 'Buys',
    text: 'Buys',
    style: {color:'#c678dd'},
    headerStyle: {
      color: '#c678dd',
      border: '2px solid #c678dd'
    }
  }, {
    dataField: 'BuysVolume',
    text: 'Buys Volume',
    style: {color:'#c678dd'},
    headerStyle: {
      color: '#c678dd',
      border: '2px solid #c678dd'
    }
  }];

// Use effect - effect hook
function PriceInfo(){
    async function loadGateInfo() {
      await sleep(2000)
      fetch("http://localhost:8080/prices")
      .then((response) => response.json())
      .then((data) => setGateInfo(data));
    }

    const [gateInfo, setGateInfo] = useState([]);
    useEffect(() => {loadGateInfo()});

    return <BootstrapTable
            bootstrap4 
            keyField="Sells" 
            data={ gateInfo } 
            condensed
            striped 
            caption={<PricesCaptionElement/>}
            columns={ columns }
            rowStyle= {{border: '2px solid #c678dd'}}/>
}

export default PriceInfo;