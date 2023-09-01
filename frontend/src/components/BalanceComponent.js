// import { useEffect, useState } from 'react'
import { useEffect, useState } from 'react'
import BootstrapTable from 'react-bootstrap-table-next';
import 'react-bootstrap-table-next/dist/react-bootstrap-table2.min.css';

const BalanceCaptionElement = () => <h3 className='Balance-Info-element'>Balance</h3>;

const sleep = ms => new Promise(
  resolve => setTimeout(resolve, ms)
);

const rowStyle2 = (row) => {
  const style = {};
  style.border =  '2px solid #e06c75';
  if (row.id === "Gate") {
    style.color = '#C678DD';
  }
  if (row.Platform === "Stex") {
    style.color = '#98C379';
  }
  return style;
};

const columns = [
    {
dataField:"Platform", 
text:"Platform",
headerStyle: {
    color: '#E06C75',
    border: '2px solid #E06C75'
    }
},
    {
  dataField:"USDT", 
  text:"USDT",
  headerStyle: {
      color: '#E06C75',
      border: '2px solid #E06C75'
    }
  },
  {
  dataField:"ETH", 
  text:"ETH",
  headerStyle: {
      color: '#E06C75',
      border: '2px solid #E06C75'
  }
  },
  {
  dataField:"BTC", 
  text:"BTC",
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
            keyField="Platform"
            data={ BalanceInfo } 
            condensed
            caption={<BalanceCaptionElement/>} 
            columns={ columns }
            rowStyle= { rowStyle2 }/>
}

export default BalanceInfo;