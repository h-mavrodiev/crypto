import { useEffect, useState } from 'react'
import BootstrapTable from 'react-bootstrap-table-next';
import 'react-bootstrap-table-next/dist/react-bootstrap-table2.min.css';

const StexCaptionElement = () => <h3 className='Stex-caption-element'>Stex</h3>;

const sleep = ms => new Promise(
    resolve => setTimeout(resolve, ms)
  );

const columns = [{
    dataField: 'IWannaBuyFor',
    text: '$',
    style: {color:'#98c379'},
    headerStyle: {
        color: '#98c379',
        border: '2px solid #98c379'
      }
  }, {
    dataField: 'TheySellForWeightedUSD',
    text: 'Price per $ (W)',
    style: {color:'#98c379'},
    headerStyle: {
        color: '#98c379',
        border: '2px solid #98c379'
      }
  },{
    dataField: 'ICanBuy',
    text: 'I Can Buy',
    style: {color:'#98c379'},
    headerStyle: {
        color: '#98c379',
        border: '2px solid #98c379'
      }
  },{
    dataField: 'TheySellForWeighted',
    text: 'They Sell For (W)',
    style: {color:'#98c379'},
    headerStyle: {
        color: '#98c379',
        border: '2px solid #98c379'
      }
  }, {
    dataField: 'ICanSellFromGate',
    text: 'I Can Sell From Gate',
    style: {color:'#98c379'},
    headerStyle: {
        color: '#98c379',
        border: '2px solid #98c379'
      }
  }, {
    dataField: 'ICanSellFromGateForWeighted',
    text: 'I Can Sell From Gate For',
    style: {color:'#98c379'},
    headerStyle: {
        color: '#98c379',
        border: '2px solid #98c379'
      }
  }];

// State is
function StexInfo(){
    async function loadGateInfo() {
        await sleep(1000)
        fetch("http://localhost:8080/stex")
        .then((response) => response.json())
        .then((data) => setStexInfo([data]));
    }

    const [stexInfo, setStexInfo] = useState([]);
    useEffect(() => {loadGateInfo()});

    return <BootstrapTable 
            bootstrap4
            keyField="IWannaBuyFor" 
            data={ stexInfo }
            caption={<StexCaptionElement/>}
            condensed
            striped
            columns={ columns }
            rowStyle= {{border: '2px solid #98c379'}} />
}

export default StexInfo;