import './App.css';
import GateInfo from './components/GateComponent'
import StexInfo from './components/StexComponent';
import WalletInfo from './components/WalletComponent';
import 'bootstrap/dist/css/bootstrap.min.css';
import 'react-bootstrap-table-next/dist/react-bootstrap-table2.min.css';


const App = () => {
  return  <div className="App">
            <div className="Platform-panel">
              <h1 className="Platform-panel-header">PLATFORMS</h1>
                  <GateInfo/>
                  <StexInfo />
            </div>
                <div className='More-panel'>
                  <h1 className="More-panel-header">MORE</h1>
                    <WalletInfo />
                </div>
          </div>
};

export default App;