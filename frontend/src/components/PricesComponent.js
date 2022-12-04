import { useEffect, useState } from 'react'
import BootstrapTable from 'react-bootstrap-table-next';
import 'bootstrap/dist/css/bootstrap.min.css';
import 'react-bootstrap-table-next/dist/react-bootstrap-table2.min.css';

const PricesCaptionElement = () => <h3 className='Prices-caption-element'>Prices</h3>;

const sleep = ms => new Promise(
  resolve => setTimeout(resolve, ms)
);

const rowStyle2 = (row) => {
  const style = {};
  style.border =  '2px solid #e06c75';
  
  if (row.Platform === 'Gate') {
    style.color = '#C678DD';
  }
  
  if (row.Platform === 'Stex') {
    style.color = '#98C379';
  }
  return style;
};

const columns = [{
  dataField: 'Platform',
  text: 'Platform',
  headerStyle: {
    color: '#e06c75',
    border: '2px solid #e06c75'
  }
},{
    dataField: 'Sells',
    text: 'Sells',
    headerStyle: {
      color: '#e06c75',
      border: '2px solid #e06c75'
    }
  },{
    dataField: 'SellsVolume',
    text: 'Sells Volume',
    headerStyle: {
      color: '#e06c75',
      border: '2px solid #e06c75'
    }
  },  {
    dataField: 'Buys',
    text: 'Buys',
    headerStyle: {
      color: '#e06c75',
      border: '2px solid #e06c75'
    }
  }, {
    dataField: 'BuysVolume',
    text: 'Buys Volume',
    headerStyle: {
      color: '#e06c75',
      border: '2px solid #e06c75'
    }
  }];

// Use effect - effect hook
function PriceInfo(){
    async function loadPricesInfo() {
      await sleep(2000)
      fetch("http://localhost:8080/prices")
      .then((response) => response.json())
      .then((data) => setPricesInfo(data));
    }

    const [pricesInfo, setPricesInfo] = useState([]);
    useEffect(() => {loadPricesInfo()});

    return <BootstrapTable
            bootstrap4 
            keyField="Platform" 
            data={ pricesInfo }
            condensed
            caption={<PricesCaptionElement/>}
            columns={ columns }
            rowStyle= { rowStyle2 }/>
}

export default PriceInfo;