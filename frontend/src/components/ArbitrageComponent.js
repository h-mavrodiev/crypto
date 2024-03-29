import { useEffect, useState } from 'react'
import BootstrapTable from 'react-bootstrap-table-next';
import 'bootstrap/dist/css/bootstrap.min.css';
import 'react-bootstrap-table-next/dist/react-bootstrap-table2.min.css';

const ArbitrageCaptionElement = () => <h3 className='Arbitrage-caption-element'>Arbitrage</h3>;

const sleep = ms => new Promise(
  resolve => setTimeout(resolve, ms)
);

const rowStyle2 = (row) => {
  const style = {};
  style.border =  '2px solid #e06c75'
  if (row.Platforms === "Gate to Stex") {
    style.color = '#C678DD';
  }
  if (row.Platforms === "Stex to Gate") {
    style.color = '#98C379';
  }
  return style;
};

const columns = [{
    dataField: 'Platforms',
    text: 'Platforms',
    headerStyle: {
      color: '#e06c75',
      border: '2px solid #e06c75'
    }
  },{
    dataField: 'Arbitrage',
    text: 'Arbitrage',
    headerStyle: {
      color: '#e06c75',
      border: '2px solid #e06c75'
    }
  }];

// Use effect - effect hook
function ArbitrageInfo(){
    async function loadArbitrageInfo() {
      await sleep(2000)
      fetch("http://localhost:8080/arbitrage")
      .then((response) => response.json())
      .then((data) => setArbitrageInfo(data));
    }

    const [ArbitrageInfo, setArbitrageInfo] = useState([]);
    useEffect(() => {loadArbitrageInfo()});

    return <BootstrapTable
            bootstrap4 
            keyField="Platforms" 
            data={ ArbitrageInfo } 
            condensed
            caption={<ArbitrageCaptionElement/>}
            columns={ columns }
            rowStyle= { rowStyle2 }/>
}

export default ArbitrageInfo;