// import { useEffect, useState } from 'react'
import BootstrapTable from 'react-bootstrap-table-next';
import 'react-bootstrap-table-next/dist/react-bootstrap-table2.min.css';

const WalletCaptionElement = () => <h3 className='Wallet-caption-element'>Wallet</h3>;

function WalletInfo() {
return <BootstrapTable
bootstrap4 
keyField="id"
data={ [{Platform: "Gate", Amount: 300}, {Platform: "Stex", Amount: 200}] } 
caption={<WalletCaptionElement/>} 
rowStyle= {{border: '2px solid #E06C75'}}
columns={ [{dataField:"Platform", 
            text:"Platform",
            style: {color:'#E06C75'},
            headerStyle: {
                color: '#E06C75',
                border: '2px solid #E06C75'
              }},
            {dataField:"Amount", 
            text:"Amount",
            style: {color:'#E06C75'},
            headerStyle: {
                color: '#E06C75',
                border: '2px solid #E06C75'
              }}] } />
            }

export default WalletInfo;