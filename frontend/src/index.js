import React from 'react';
import ReactDOM from 'react-dom/client';
import './index.css';
import App from './App';
import reportWebVitals from './reportWebVitals';

const root = ReactDOM.createRoot(document.getElementById('root'));
root.render(
  <React.StrictMode>
    <App />
  </React.StrictMode>
);

// If you want to start measuring performance in your app, pass a function
// to log results (for example: reportWebVitals(console.log))
// or send to an analytics endpoint. Learn more: https://bit.ly/CRA-vitals
reportWebVitals();


// var wsGateSocket = new WebSocket("ws://localhost:8080/wsGate")

// let connect = () => {
//   console.log("Attempting Conncetion...");

//   wsGateSocket.onopen = () => {
//     console.log("Successfully Connected"); 
//   };

//   wsGateSocket.onmessage = msg => {
//     var obj = JSON.parse(msg.data)
//     document.write("<br>onmessage: " + JSON.stringify(obj, null, 4) + "</br>");
//     console.log(JSON.stringify(obj, null, 4))
//   };

//   wsGateSocket.onclose = event => {
//     console.log("Socket Closed Connection: ", event);
//   };

//   wsGateSocket.onerror = error => {
//     console.log("Socket Error: ", error);
//   };
// };

// export { connect }
