import { useEffect, useState } from 'react'
import BootstrapTable from 'react-bootstrap-table-next';
import 'react-bootstrap-table-next/dist/react-bootstrap-table2.min.css';

const StexCaptionElement = () => <h3 className='Stex-caption-element'>Stex</h3>;

const sleep = ms => new Promise(
    resolve => setTimeout(resolve, ms)
  );

const columns = [{
    dataField: 'AskFixedUSDDemand',
    text: 'Ask Fixed USD Demand',
    style: {color:'#98c379'},
    headerStyle: {
        color: '#98c379',
        border: '2px solid #98c379'
      }
  },{
    dataField: 'AskWeightedPrice',
    text: 'Ask Weighted Price',
    style: {color:'#98c379'},
    headerStyle: {
        color: '#98c379',
        border: '2px solid #98c379'
      }
  },  {
    dataField: 'AskWeightedUSDPrice',
    text: 'Ask Weighted USD Price',
    style: {color:'#98c379'},
    headerStyle: {
        color: '#98c379',
        border: '2px solid #98c379'
      }
  }, {
    dataField: 'AskAmount',
    text: 'Ask Amount',
    style: {color:'#98c379'},
    headerStyle: {
        color: '#98c379',
        border: '2px solid #98c379'
      }
  }, {
    dataField: 'BidFixedUSDDemand',
    text: 'Bid Fixed USD Demand',
    style: {color:'#98c379'},
    headerStyle: {
        color: '#98c379',
        border: '2px solid #98c379'
      }
  },{
    dataField: 'BidWeightedPrice',
    text: 'Bid Weighted Price',
    style: {color:'#98c379'},
    headerStyle: {
        color: '#98c379',
        border: '2px solid #98c379'
      }
  }, {
    dataField: 'BidWeightedUSDPrice',
    text: 'Bid Weighted USD Price',
    style: {color:'#98c379'},
    headerStyle: {
        color: '#98c379',
        border: '2px solid #98c379'
      }
  }, {
    dataField: 'BidAmount',
    text: 'Bid Amount',
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

    return <BootstrapTable bootstrap4
    keyField="AskFixedUSDDemand" 
    data={ stexInfo }
    caption={<StexCaptionElement/>}
    condensed
    striped
    columns={ columns }
    rowStyle= {{border: '2px solid #98c379'}} />
}

export default StexInfo;