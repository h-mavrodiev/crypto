import './App.css';
import PriceInfo from './components/PricesComponent'
import BalanceInfo from './components/BalanceComponent';
import ArbitrageInfo from './components/ArbitrageComponent';
import 'bootstrap/dist/css/bootstrap.min.css';
import 'react-bootstrap-table-next/dist/react-bootstrap-table2.min.css';


const App = () => {
  return  <div className="App">
            <div className='More-panel'>
              <h1 className="More-panel-header">MORE</h1>
                <BalanceInfo />
                <ArbitrageInfo />
            </div>
            <div className="Platform-panel">
              <h1 className="Platform-panel-header">PLATFORMS</h1>
                  <PriceInfo/>
            </div>
          </div>
};

export default App;