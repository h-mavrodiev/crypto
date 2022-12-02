// import { useEffect, useState } from 'react'
import { useEffect, useState } from 'react'
import BootstrapTable from 'react-bootstrap-table-next';
import 'react-bootstrap-table-next/dist/react-bootstrap-table2.min.css';

const BalanceCaptionElement = () => <h3 className='Balance-Info-element'>Balance</h3>;

const sleep = ms => new Promise(
  resolve => setTimeout(resolve, ms)
);

const columns = [
    {
dataField:"Platform", 
text:"Platform",
style: {color:'#E06C75'},
headerStyle: {
    color: '#E06C75',
    border: '2px solid #E06C75'
    }
},
    {
  dataField:"USDT", 
  text:"USDT",
  style: {color:'#E06C75'},
  headerStyle: {
      color: '#E06C75',
      border: '2px solid #E06C75'
    }
  },
  {
  dataField:"ETH", 
  text:"ETH",
  style: {color:'#E06C75'},
  headerStyle: {
      color: '#E06C75',
      border: '2px solid #E06C75'
  }
  },
  {
  dataField:"BTC", 
  text:"BTC",
  style: {color:'#E06C75'},
  headerStyle: {
      color: '#E06C75',
      border: '2px solid #E06C75'
    }
  }
]

function BalanceInfo() {
    async function loadBalanceInfo() {
      await sleep(10000)
      fetch("http://localhost:8080/balance")
      .then((response) => response.json())
      .then((data) => setBalanceInfo(data));
    }

    const [BalanceInfo, setBalanceInfo] = useState([]);
    useEffect(() => {loadBalanceInfo()});


    return <BootstrapTable
            bootstrap4 
            keyField="USDT"
            data={ BalanceInfo } 
            condensed
            striped 
            caption={<BalanceCaptionElement/>} 
            columns={ columns }
            rowStyle= {{border: '2px solid #E06C75'}}/>
}

export default BalanceInfo;